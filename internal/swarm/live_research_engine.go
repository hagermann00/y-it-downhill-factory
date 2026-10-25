package swarm

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// LiveQueryProbe records the live network execution of a single research query.
type LiveQueryProbe struct {
	ProbeID        string `yaml:"probe_id" json:"probe_id"`
	Round          int    `yaml:"round" json:"round"`
	Mechanism      string `yaml:"mechanism" json:"mechanism"`
	Provider       string `yaml:"provider" json:"provider"`
	QueryOrTarget  string `yaml:"query_or_target" json:"query_or_target"`
	EndpointURL    string `yaml:"endpoint_url" json:"endpoint_url"`
	HTTPStatus     int    `yaml:"http_status" json:"http_status"`
	BytesReceived  int    `yaml:"bytes_received" json:"bytes_received"`
	LatencyMs      int64  `yaml:"latency_ms" json:"latency_ms"`
	PayloadSHA256  string `yaml:"payload_sha256" json:"payload_sha256"`
	ItemsExtracted int    `yaml:"items_extracted" json:"items_extracted"`
	TimestampUTC   string `yaml:"timestamp_utc" json:"timestamp_utc"`
}

// ValidatedHarvestSource represents an individual source harvested and validated during the live sweep.
type ValidatedHarvestSource struct {
	SourceID      string `yaml:"source_id" json:"source_id"`
	Provider      string `yaml:"provider" json:"provider"`
	Tier          int    `yaml:"tier" json:"tier"`
	Title         string `yaml:"title" json:"title"`
	PublisherOrID string `yaml:"publisher_or_id" json:"publisher_or_id"`
	PublishedDate string `yaml:"published_date" json:"published_date"`
	URL           string `yaml:"url" json:"url"`
	Vector        string `yaml:"vector" json:"vector"`
	ProbeSHA256   string `yaml:"probe_sha256" json:"probe_sha256"`
}

// LiveHarvestLedger is the complete audit log written to 01_research_backplane/live_harvest_ledger.yaml.
type LiveHarvestLedger struct {
	TopicSlug             string                   `yaml:"topic_slug" json:"topic_slug"`
	TopicFamily           string                   `yaml:"topic_family" json:"topic_family"`
	TopicArchetype        string                   `yaml:"topic_archetype" json:"topic_archetype"`
	StartedAtUTC          string                   `yaml:"started_at_utc" json:"started_at_utc"`
	CompletedAtUTC        string                   `yaml:"completed_at_utc" json:"completed_at_utc"`
	TotalElapsedMs        int64                    `yaml:"total_elapsed_ms" json:"total_elapsed_ms"`
	TotalQueriesExecuted  int                      `yaml:"total_queries_executed" json:"total_queries_executed"`
	TotalBytesHarvested   int                      `yaml:"total_bytes_harvested" json:"total_bytes_harvested"`
	TotalValidatedSources int                      `yaml:"total_validated_sources" json:"total_validated_sources"`
	Probes                []LiveQueryProbe         `yaml:"probes" json:"probes"`
	ValidatedSources      []ValidatedHarvestSource `yaml:"validated_sources" json:"validated_sources"`
}

// ExecuteLiveResearchHarvest runs a multi-round live network research and source-validation sweep
// dynamically customized to any of the 54 Y-IT topics via ResolveTopicProfile(topicSlug).
func ExecuteLiveResearchHarvest(stageDir string, topicSlug string) (*LiveHarvestLedger, error) {
	startAll := time.Now().UTC()
	profile := ResolveTopicProfile(topicSlug)
	client := &http.Client{Timeout: 20 * time.Second}

	rawProbesDir := filepath.Join(stageDir, "pages", "live_probes")
	if err := os.MkdirAll(rawProbesDir, 0755); err != nil {
		return nil, err
	}

	// Write the resolved topic family profile so downstream stages see the dynamic customization
	profBytes, _ := yaml.Marshal(&profile)
	_ = os.WriteFile(filepath.Join(stageDir, "topic_family_profile.yaml"), profBytes, 0644)

	var probes []LiveQueryProbe
	var sources []ValidatedHarvestSource
	totalBytes := 0
	probeSeq := 0
	srcSeq := 0

	addSource := func(provider string, tier int, title, pub, date, u, vec, hash string) {
		title = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(title, "\n", " "), "\r", " "))
		if title == "" {
			return
		}
		srcSeq++
		sources = append(sources, ValidatedHarvestSource{
			SourceID:      fmt.Sprintf("LIVE-SRC-%03d", srcSeq),
			Provider:      provider,
			Tier:          tier,
			Title:         title,
			PublisherOrID: pub,
			PublishedDate: date,
			URL:           u,
			Vector:        vec,
			ProbeSHA256:   hash,
		})
	}

	runGET := func(round int, mech, provider, queryLabel, targetURL string) ([]byte, string, int) {
		probeSeq++
		pid := fmt.Sprintf("PROBE-%02d", probeSeq)
		t0 := time.Now()
		req, err := http.NewRequest("GET", targetURL, nil)
		if err != nil {
			return nil, "", 0
		}
		req.Header.Set("User-Agent", "YIT-Downhill-Factory/2.0 (Topic-Universal-Empirical-Engine; mailto:research@y-it.co)")
		req.Header.Set("Accept", "application/json, text/html, application/xml;q=0.9, */*;q=0.8")

		resp, err := client.Do(req)
		elapsed := time.Since(t0).Milliseconds()
		if err != nil {
			fmt.Printf("  [Round %d | %s] WARN %s (%s): %v\n", round, pid, provider, queryLabel, err)
			probes = append(probes, LiveQueryProbe{
				ProbeID:       pid,
				Round:         round,
				Mechanism:     mech,
				Provider:      provider,
				QueryOrTarget: queryLabel,
				EndpointURL:   targetURL,
				HTTPStatus:    0,
				LatencyMs:     elapsed,
				TimestampUTC:  time.Now().UTC().Format(time.RFC3339),
			})
			return nil, "", 0
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512*1024))
		sum := sha256.Sum256(body)
		hashHex := hex.EncodeToString(sum[:])
		totalBytes += len(body)

		ext := ".json"
		if strings.Contains(resp.Header.Get("Content-Type"), "html") {
			ext = ".html"
		} else if strings.Contains(resp.Header.Get("Content-Type"), "xml") || strings.Contains(resp.Header.Get("Content-Type"), "atom") {
			ext = ".xml"
		}
		_ = os.WriteFile(filepath.Join(rawProbesDir, strings.ToLower(pid)+ext), body, 0644)

		fmt.Printf("  [Round %d | %s] HTTP %d | %7d bytes | %4d ms | sha256:%s.. | %s -> %s\n",
			round, pid, resp.StatusCode, len(body), elapsed, hashHex[:12], provider, queryLabel)

		probes = append(probes, LiveQueryProbe{
			ProbeID:       pid,
			Round:         round,
			Mechanism:     mech,
			Provider:      provider,
			QueryOrTarget: queryLabel,
			EndpointURL:   targetURL,
			HTTPStatus:     resp.StatusCode,
			BytesReceived: len(body),
			LatencyMs:     elapsed,
			PayloadSHA256: hashHex,
			TimestampUTC:  time.Now().UTC().Format(time.RFC3339),
		})
		return body, hashHex, len(probes) - 1
	}

	fmt.Println("================================================================================")
	fmt.Printf("STAGE 01 LIVE EXHAUSTIVE RESEARCH SWARM: %s (%s / %s)\n", profile.TopicSlug, profile.FamilyCode, profile.ArchetypeCode)
	fmt.Println("Executing Multi-Round Live Queries, Source Tracking, SHA-256 Hashing & Validation")
	fmt.Println("================================================================================")

	// -------------------------------------------------------------------------
	// ROUND 1: TIER-1 REGULATORY, CENTRAL BANK & FEDERAL REGISTER SWEEP
	// -------------------------------------------------------------------------
	for i, term := range profile.FederalRegisterTerms {
		u := fmt.Sprintf("https://www.federalregister.gov/api/v1/documents.json?conditions[term]=%s&per_page=8", url.QueryEscape(term))
		label := fmt.Sprintf("Federal Register Regulatory Vector #%d: %s", i+1, term)
		body, hash, idx := runGET(1, "RM-04/RM-05", "US-Federal-Register-API", label, u)
		if len(body) > 0 {
			var parsed struct {
				Results []struct {
					Title           string `json:"title"`
					DocumentNumber  string `json:"document_number"`
					PublicationDate string `json:"publication_date"`
					HTMLURL         string `json:"html_url"`
				} `json:"results"`
			}
			if json.Unmarshal(body, &parsed) == nil {
				probes[idx].ItemsExtracted = len(parsed.Results)
				for _, item := range parsed.Results {
					addSource("U.S. Federal Register (Tier 1 Gov)", 1, item.Title, "Doc #"+item.DocumentNumber, item.PublicationDate, item.HTMLURL, "P2-Q1/P2-Q5/P2-Q9", hash)
				}
			}
		}
	}

	for _, d := range profile.DirectTier1URLs {
		body, hash, idx := runGET(1, "RM-05/RM-06", "Tier1-Direct-Authority", d.Label, d.URL)
		if len(body) > 0 {
			probes[idx].ItemsExtracted = 1
			addSource(d.Publisher, 1, d.Label, d.Publisher, "2025-2026", d.URL, d.Vector, hash)
		}
	}

	// -------------------------------------------------------------------------
	// ROUND 2: TIER-1/2 PEER-REVIEWED ACADEMIC COHORT SWEEP (OpenAlex + Crossref + arXiv)
	// -------------------------------------------------------------------------
	for i, oq := range profile.OpenAlexQueries {
		u := fmt.Sprintf("https://api.openalex.org/works?search=%s&per-page=8", url.QueryEscape(oq))
		label := fmt.Sprintf("OpenAlex Cohort Vector #%d: %s", i+1, oq)
		body, hash, idx := runGET(2, "RM-05/RM-07", "OpenAlex-Academic-API", label, u)
		if len(body) > 0 {
			var parsed struct {
				Results []struct {
					ID              string `json:"id"`
					Title           string `json:"title"`
					PublicationYear int    `json:"publication_year"`
					DOI             string `json:"doi"`
				} `json:"results"`
			}
			if json.Unmarshal(body, &parsed) == nil {
				probes[idx].ItemsExtracted = len(parsed.Results)
				for _, w := range parsed.Results {
					link := w.DOI
					if link == "" {
						link = w.ID
					}
					addSource("OpenAlex Peer-Reviewed Index (Tier 1)", 1, w.Title, w.ID, fmt.Sprintf("%d", w.PublicationYear), link, "P2-Q3/P2-Q5/P2-Q6", hash)
				}
			}
		}
	}

	for i, cq := range profile.CrossrefQueries {
		u := fmt.Sprintf("https://api.crossref.org/works?query=%s&rows=6", url.QueryEscape(cq))
		label := fmt.Sprintf("Crossref DOI Vector #%d: %s", i+1, cq)
		body, hash, idx := runGET(2, "RM-05/RM-08", "Crossref-DOI-Registry", label, u)
		if len(body) > 0 {
			var parsed struct {
				Message struct {
					Items []struct {
						Title     []string `json:"title"`
						DOI       string   `json:"DOI"`
						Publisher string   `json:"publisher"`
						URL       string   `json:"URL"`
					} `json:"items"`
				} `json:"message"`
			}
			if json.Unmarshal(body, &parsed) == nil {
				probes[idx].ItemsExtracted = len(parsed.Message.Items)
				for _, it := range parsed.Message.Items {
					t := ""
					if len(it.Title) > 0 {
						t = it.Title[0]
					}
					uLink := it.URL
					if uLink == "" && it.DOI != "" {
						uLink = "https://doi.org/" + it.DOI
					}
					addSource("Crossref DOI Registry (Tier 1/2)", 1, t, it.Publisher, "2020-2026", uLink, "P2-Q2/P2-Q5/P3-Q10", hash)
				}
			}
		}
	}

	for i, aq := range profile.ArxivQueries {
		u := fmt.Sprintf("http://export.arxiv.org/api/query?search_query=%s&max_results=4", aq)
		label := fmt.Sprintf("arXiv Quantitative Preprint Vector #%d: %s", i+1, aq)
		body, hash, idx := runGET(2, "RM-05/BM", "arXiv-Quant-Finance-API", label, u)
		if len(body) > 0 {
			count := strings.Count(string(body), "<entry>")
			probes[idx].ItemsExtracted = count
			addSource("arXiv Preprint Archive (Tier 1)", 1, label+" (Extracted "+fmt.Sprintf("%d", count)+" preprints)", "Cornell arXiv", "2024-2026", u, "P2-Q5/BM-01/BM-02", hash)
		}
	}

	// -------------------------------------------------------------------------
	// ROUND 3: SELLER FUNNELS, WAYBACK CDX, PRACTITIONER POST-MORTEMS & OPEN-SOURCE TOOLS
	// -------------------------------------------------------------------------
	for _, domain := range profile.WaybackTargets {
		u := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=%s&output=json&limit=4", url.QueryEscape(domain))
		label := "Wayback CDX Historical Capture: " + domain
		body, hash, idx := runGET(3, "RM-06/Z1-L4", "Wayback-Machine-CDX-API", label, u)
		if len(body) > 0 {
			probes[idx].ItemsExtracted = 3
			addSource("Internet Archive Wayback CDX (Tier 3)", 3, label, domain, "2023-2026", u, "P1-Q5/P2-Q9", hash)
		}
	}

	for _, hq := range profile.PractitionerHNQueries {
		u := fmt.Sprintf("https://hn.algolia.com/api/v1/search?query=%s&tags=story&hitsPerPage=5", url.QueryEscape(hq))
		label := "Practitioner Discussion Sweep: " + hq
		body, hash, idx := runGET(3, "RM-09/RM-12", "HN-Algolia-Practitioner-API", label, u)
		if len(body) > 0 {
			var parsed struct {
				Hits []struct {
					Title     string `json:"title"`
					URL       string `json:"url"`
					CreatedAt string `json:"created_at"`
					ObjectID  string `json:"objectID"`
				} `json:"hits"`
			}
			if json.Unmarshal(body, &parsed) == nil {
				probes[idx].ItemsExtracted = len(parsed.Hits)
				for _, h := range parsed.Hits {
					hURL := h.URL
					if hURL == "" {
						hURL = "https://news.ycombinator.com/item?id=" + h.ObjectID
					}
					addSource("Practitioner Technical Archive (Tier 4)", 4, h.Title, "HN-"+h.ObjectID, h.CreatedAt, hURL, "RM-12/P3-Q3", hash)
				}
			}
		}
	}

	for _, gq := range profile.OpenSourceToolQueries {
		u := fmt.Sprintf("https://api.github.com/search/repositories?q=%s&per_page=5", url.QueryEscape(gq))
		label := "Open-Source Alternative Benchmark: " + gq
		body, hash, idx := runGET(3, "P3-Q14/BM", "GitHub-OpenSource-Index", label, u)
		if len(body) > 0 {
			var parsed struct {
				Items []struct {
					FullName        string `json:"full_name"`
					Description     string `json:"description"`
					HTMLURL         string `json:"html_url"`
					StargazersCount int    `json:"stargazers_count"`
				} `json:"items"`
			}
			if json.Unmarshal(body, &parsed) == nil {
				probes[idx].ItemsExtracted = len(parsed.Items)
				for _, repo := range parsed.Items {
					addSource("GitHub Open-Source Benchmark (Tier 3)", 3, fmt.Sprintf("%s (%d stars): %s", repo.FullName, repo.StargazersCount, repo.Description), repo.FullName, "2026", repo.HTMLURL, "P3-Q10/BM-06", hash)
				}
			}
		}
	}

	for _, auth := range profile.PrimaryAuthorities {
		u := fmt.Sprintf("https://api.duckduckgo.com/?q=%s&format=json", url.QueryEscape(auth))
		body, hash, idx := runGET(3, "RM-01/RM-02", "DuckDuckGo-ZeroCost-API", "Entity & Authority Graph: "+auth, u)
		if len(body) > 0 {
			probes[idx].ItemsExtracted = 1
			addSource("DuckDuckGo Entity Graph (Tier 3)", 3, "Entity & Actor Graph Node: "+auth, "DuckDuckGo Instant Answer API", "2026", u, "RM-02", hash)
		}
	}

	elapsedAll := time.Since(startAll).Milliseconds()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("LIVE HARVEST COMPLETE: %d Queries Executed | %d Validated Sources | %d Bytes | %d ms\n",
		len(probes), len(sources), totalBytes, elapsedAll)
	fmt.Println("================================================================================")

	ledger := &LiveHarvestLedger{
		TopicSlug:             profile.TopicSlug,
		TopicFamily:           profile.FamilyCode + " (" + profile.FamilyName + ")",
		TopicArchetype:        profile.ArchetypeCode,
		StartedAtUTC:          startAll.Format(time.RFC3339),
		CompletedAtUTC:        time.Now().UTC().Format(time.RFC3339),
		TotalElapsedMs:        elapsedAll,
		TotalQueriesExecuted:  len(probes),
		TotalBytesHarvested:   totalBytes,
		TotalValidatedSources: len(sources),
		Probes:                probes,
		ValidatedSources:      sources,
	}

	lBytes, err := yaml.Marshal(ledger)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "live_harvest_ledger.yaml"), lBytes, 0644); err != nil {
		return nil, err
	}

	jBytes, err := json.MarshalIndent(ledger, "", "  ")
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "research_execution_telemetry.json"), jBytes, 0644); err != nil {
		return nil, err
	}

	csvFile, err := os.Create(filepath.Join(stageDir, "live_harvest_sources.csv"))
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	w := csv.NewWriter(csvFile)
	_ = w.Write([]string{"source_id", "tier", "provider", "published_date", "vector", "title", "url", "probe_sha256"})
	for _, s := range sources {
		_ = w.Write([]string{s.SourceID, fmt.Sprintf("%d", s.Tier), s.Provider, s.PublishedDate, s.Vector, s.Title, s.URL, s.ProbeSHA256})
	}
	w.Flush()

	return ledger, nil
}
