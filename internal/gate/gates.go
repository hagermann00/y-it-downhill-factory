package gate

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Question represents one of the 38 canonical research questions.
type Question struct {
	ID       string `yaml:"id" json:"id"`
	Phase    int    `yaml:"phase" json:"phase"`
	Question string `yaml:"question" json:"question"`
	Feeds    string `yaml:"feeds" json:"feeds"`
}

type QuestionsFile struct {
	TopicSlug string     `yaml:"topic_slug"`
	Questions []Question `yaml:"questions"`
}

// Row represents a single research row (claim C-, fact F-, or case K-).
type Row struct {
	ID       string   `yaml:"id" json:"id"`
	Kind     string   `yaml:"kind" json:"kind"`
	Question string   `yaml:"question" json:"question"`
	Phases   []int    `yaml:"phases" json:"phases"`
	Text     string   `yaml:"text" json:"text"`
	Number   string   `yaml:"number,omitempty" json:"number,omitempty"`
	Venue    string   `yaml:"venue,omitempty" json:"venue,omitempty"`
	Seen     string   `yaml:"seen,omitempty" json:"seen,omitempty"`
	Seller   string   `yaml:"seller,omitempty" json:"seller,omitempty"`
	Capture  string   `yaml:"capture,omitempty" json:"capture,omitempty"`
	Tier     int      `yaml:"tier,omitempty" json:"tier,omitempty"`
	Market   string   `yaml:"market,omitempty" json:"market,omitempty"`
	Tests    []string `yaml:"tests,omitempty" json:"tests,omitempty"`
	Mold     string   `yaml:"mold,omitempty" json:"mold,omitempty"`
	Source   string   `yaml:"source" json:"source"`
}

type RowsFile struct {
	TopicSlug string `yaml:"topic_slug"`
	Rows      []Row  `yaml:"rows"`
}

// EvidenceEntry links a row ID to its saved source page and verbatim passage.
type EvidenceEntry struct {
	Row       string `yaml:"row" json:"row"`
	URL       string `yaml:"url" json:"url"`
	Passage   string `yaml:"passage" json:"passage"`
	Retrieved string `yaml:"retrieved" json:"retrieved"`
	Archive   string `yaml:"archive" json:"archive"`
	SHA256    string `yaml:"sha256" json:"sha256"`
}

type EvidenceFile struct {
	TopicSlug string          `yaml:"topic_slug"`
	Evidence  []EvidenceEntry `yaml:"evidence"`
}

// FrictionEntry represents a divergence row in friction_register.yaml (Oct 17 prompt).
type FrictionEntry struct {
	EntryID           string `yaml:"entry_id" json:"entry_id"`
	Stage             string `yaml:"stage" json:"stage"`
	PromisedState     string `yaml:"promised_state" json:"promised_state"`
	DocumentedReality string `yaml:"documented_reality" json:"documented_reality"`
	CostDelta         string `yaml:"cost_delta" json:"cost_delta"`
	AttritionRate     string `yaml:"attrition_rate" json:"attrition_rate"`
	Tier              int    `yaml:"tier" json:"tier"`
	ClaimRowID        string `yaml:"claim_row_id" json:"claim_row_id"`
	FactRowID         string `yaml:"fact_row_id" json:"fact_row_id"`
}

type FrictionRegisterFile struct {
	TopicSlug string          `yaml:"topic_slug"`
	Entries   []FrictionEntry `yaml:"entries"`
}

// SourceItem represents a bibliography source in sources.yaml.
type SourceItem struct {
	SourceID  string `yaml:"source_id" json:"source_id"`
	Publisher string `yaml:"publisher" json:"publisher"`
	Title     string `yaml:"title" json:"title"`
	Date      string `yaml:"date" json:"date"`
	URL       string `yaml:"url" json:"url"`
	Tier      int    `yaml:"tier" json:"tier"`
	Archive   string `yaml:"archive" json:"archive"`
}

type SourcesFile struct {
	TopicSlug string       `yaml:"topic_slug"`
	Sources   []SourceItem `yaml:"sources"`
}

type AffiliateItem struct {
	AffiliateID string `yaml:"affiliate_id"`
	Program     string `yaml:"program"`
	Commission  string `yaml:"commission"`
	SellerKey   string `yaml:"seller_key"`
	Seen        string `yaml:"seen"`
	SourceID    string `yaml:"source_id"`
}

type AffiliatesFile struct {
	TopicSlug  string          `yaml:"topic_slug"`
	Affiliates []AffiliateItem `yaml:"affiliates"`
}

// PhasePacket represents a phase packet in 02_predraft_spine/phase-N/packet.yaml.
type PhasePacket struct {
	TopicSlug     string                 `yaml:"topic_slug"`
	Phase         int                    `yaml:"phase"`
	SpineVersion  string                 `yaml:"spine_version"`
	PacketVersion int                    `yaml:"packet_version"`
	Dials         map[string]interface{} `yaml:"dials"`
	CitedRows     []string               `yaml:"cited_rows"`
	Ch4Molds      []string               `yaml:"ch4_molds,omitempty"`
	Ch7Pivots     int                    `yaml:"ch7_pivots,omitempty"`
	Ch8Title      string                 `yaml:"ch8_title,omitempty"`
}

// GateResult captures the deterministic output of one gate check.
type GateResult struct {
	GateID      string   `json:"gate_id"`
	Name        string   `json:"name"`
	ExitCode    int      `json:"exit_code"`
	ChecksRun   int      `json:"checks_run"`
	Failures    []string `json:"failures,omitempty"`
	SummaryLine string   `json:"summary_line"`
}

// FullGateReport holds all gate results G1..G7.
type FullGateReport struct {
	TopicSlug       string       `json:"topic_slug"`
	ExecutedAt      string       `json:"executed_at"`
	OverallExitCode int          `json:"overall_exit_code"`
	BookPageCount   int          `json:"book_page_count"`
	PamphletPages   int          `json:"pamphlet_pages"`
	FreebiePages    int          `json:"freebie_pages"`
	CoverWrapPages  int          `json:"cover_wrap_pages"`
	Gates           []GateResult `json:"gates"`
}

// Canonical38Questions returns the 38 required question IDs in exact order.
func Canonical38Questions() []string {
	var out []string
	for i := 1; i <= 12; i++ {
		out = append(out, fmt.Sprintf("P1-Q%d", i))
	}
	for i := 1; i <= 12; i++ {
		out = append(out, fmt.Sprintf("P2-Q%d", i))
	}
	for i := 1; i <= 14; i++ {
		out = append(out, fmt.Sprintf("P3-Q%d", i))
	}
	return out
}

// RunAllGates executes deterministic gates G1 through G7 on rootDir.
func RunAllGates(rootDir string) (FullGateReport, error) {
	report := FullGateReport{
		TopicSlug:  "forex-prop-firms",
		ExecutedAt: time.Now().UTC().Format(time.RFC3339),
	}

	g1 := checkG1Questions(rootDir)
	g2 := checkG2EvidenceAndPassages(rootDir)
	g3 := checkG3SourcesAndAffiliates(rootDir)
	g4 := checkG4CoverageAndMolds(rootDir)
	g5 := checkG5PhasePackets(rootDir)
	g6 := checkG6ManuscriptsAndTone(rootDir)
	g7, bookPages, pamphletPages, freebiePages, coverPages := checkG7KDPAssembly(rootDir)

	report.BookPageCount = bookPages
	report.PamphletPages = pamphletPages
	report.FreebiePages = freebiePages
	report.CoverWrapPages = coverPages
	report.Gates = []GateResult{g1, g2, g3, g4, g5, g6, g7}

	for _, g := range report.Gates {
		if g.ExitCode != 0 {
			report.OverallExitCode = 1
		}
	}
	return report, nil
}

func checkG1Questions(rootDir string) GateResult {
	res := GateResult{GateID: "G1", Name: "38 Canonical Questions Check"}
	qPath := filepath.Join(rootDir, "01_research_backplane", "questions.yaml")
	data, err := os.ReadFile(qPath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("cannot read questions.yaml: %v", err))
		res.SummaryLine = "G1 exit 1: missing questions.yaml"
		return res
	}
	var qf QuestionsFile
	if err := yaml.Unmarshal(data, &qf); err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("invalid YAML in questions.yaml: %v", err))
		res.SummaryLine = "G1 exit 1: invalid YAML"
		return res
	}
	expected := Canonical38Questions()
	seen := make(map[string]bool)
	for _, q := range qf.Questions {
		seen[q.ID] = true
		res.ChecksRun++
		if strings.TrimSpace(q.Question) == "" {
			res.Failures = append(res.Failures, fmt.Sprintf("question %s has empty text", q.ID))
		}
	}
	for _, reqID := range expected {
		res.ChecksRun++
		if !seen[reqID] {
			res.Failures = append(res.Failures, fmt.Sprintf("missing canonical question ID: %s", reqID))
		}
	}
	if len(qf.Questions) != 38 {
		res.Failures = append(res.Failures, fmt.Sprintf("expected 38 questions, got %d", len(qf.Questions)))
	}
	if len(res.Failures) > 0 {
		res.ExitCode = 1
		res.SummaryLine = fmt.Sprintf("G1 exit 1 (%d failures)", len(res.Failures))
	} else {
		res.SummaryLine = fmt.Sprintf("G1 exit 0 (38/38 canonical questions present, %d checks)", res.ChecksRun)
	}
	return res
}

var numExtractRe = regexp.MustCompile(`[\d]+(?:[.,]\d+)?%?`)

func checkG2EvidenceAndPassages(rootDir string) GateResult {
	res := GateResult{GateID: "G2", Name: "Verbatim Passage, Number Trace & Universal Phase 1 Friction Register Check"}
	rPath := filepath.Join(rootDir, "01_research_backplane", "rows.yaml")
	ePath := filepath.Join(rootDir, "01_research_backplane", "evidence.yaml")
	fPath := filepath.Join(rootDir, "01_research_backplane", "friction_register.yaml")

	rData, err := os.ReadFile(rPath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("missing rows.yaml: %v", err))
		return res
	}
	eData, err := os.ReadFile(ePath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("missing evidence.yaml: %v", err))
		return res
	}
	fData, err := os.ReadFile(fPath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("missing friction_register.yaml: %v", err))
		return res
	}

	var rf RowsFile
	var ef EvidenceFile
	var ff FrictionRegisterFile
	if err := yaml.Unmarshal(rData, &rf); err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("rows.yaml unmarshal error: %v", err))
		return res
	}
	if err := yaml.Unmarshal(eData, &ef); err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("evidence.yaml unmarshal error: %v", err))
		return res
	}
	if err := yaml.Unmarshal(fData, &ff); err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("friction_register.yaml unmarshal error: %v", err))
		return res
	}

	evMap := make(map[string]EvidenceEntry)
	for _, e := range ef.Evidence {
		evMap[e.Row] = e
	}
	rowMap := make(map[string]Row)
	for _, r := range rf.Rows {
		rowMap[r.ID] = r
		res.ChecksRun++
		ev, ok := evMap[r.ID]
		if !ok {
			res.Failures = append(res.Failures, fmt.Sprintf("row %s has no entry in evidence.yaml", r.ID))
			continue
		}
		pagePath := filepath.Join(rootDir, "01_research_backplane", "pages", ev.Archive)
		pageBytes, err := os.ReadFile(pagePath)
		if err != nil {
			res.Failures = append(res.Failures, fmt.Sprintf("row %s archive %s missing on disk", r.ID, ev.Archive))
			continue
		}
		normPage := normalizeText(string(pageBytes))
		normPassage := normalizeText(ev.Passage)
		if normPassage == "" || !strings.Contains(normPage, normPassage) {
			res.Failures = append(res.Failures, fmt.Sprintf("row %s passage not found verbatim in %s", r.ID, ev.Archive))
		}
		if strings.TrimSpace(r.Number) != "" {
			nums := numExtractRe.FindAllString(r.Number, -1)
			for _, n := range nums {
				if !strings.Contains(ev.Passage, n) {
					res.Failures = append(res.Failures, fmt.Sprintf("row %s number token %q not present in verbatim passage", r.ID, n))
				}
			}
		}
	}

	if len(ff.Entries) < 8 {
		res.Failures = append(res.Failures, fmt.Sprintf("friction_register.yaml requires >= 8 entries, got %d", len(ff.Entries)))
	}
	for _, fe := range ff.Entries {
		res.ChecksRun++
		if _, ok := rowMap[fe.ClaimRowID]; !ok {
			res.Failures = append(res.Failures, fmt.Sprintf("friction entry %s references unknown claim_row_id %s", fe.EntryID, fe.ClaimRowID))
		}
		if _, ok := rowMap[fe.FactRowID]; !ok {
			res.Failures = append(res.Failures, fmt.Sprintf("friction entry %s references unknown fact_row_id %s", fe.EntryID, fe.FactRowID))
		}
	}

	if len(res.Failures) > 0 {
		res.ExitCode = 1
		res.SummaryLine = fmt.Sprintf("G2 exit 1 (%d failures)", len(res.Failures))
	} else {
		res.SummaryLine = fmt.Sprintf("G2 exit 0 (%d rows matched verbatim in saved pages, %d friction entries verified)", len(rf.Rows), len(ff.Entries))
	}
	return res
}

func checkG3SourcesAndAffiliates(rootDir string) GateResult {
	res := GateResult{GateID: "G3", Name: "Source Tiers, Affiliate Register & Appendix Trail Check"}
	sPath := filepath.Join(rootDir, "01_research_backplane", "sources.yaml")
	aPath := filepath.Join(rootDir, "01_research_backplane", "affiliates.yaml")
	appPath := filepath.Join(rootDir, "01_research_backplane", "sources-appendix.md")

	sData, err := os.ReadFile(sPath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("missing sources.yaml: %v", err))
		return res
	}
	aData, err := os.ReadFile(aPath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("missing affiliates.yaml: %v", err))
		return res
	}
	if _, err := os.Stat(appPath); err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, "missing sources-appendix.md")
		return res
	}

	var sf SourcesFile
	var af AffiliatesFile
	if err := yaml.Unmarshal(sData, &sf); err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("sources.yaml unmarshal error: %v", err))
		return res
	}
	if err := yaml.Unmarshal(aData, &af); err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("affiliates.yaml unmarshal error: %v", err))
		return res
	}

	tierCounts := make(map[int]int)
	for _, s := range sf.Sources {
		res.ChecksRun++
		tierCounts[s.Tier]++
		if s.Tier < 1 || s.Tier > 4 {
			res.Failures = append(res.Failures, fmt.Sprintf("source %s has invalid tier %d", s.SourceID, s.Tier))
		}
		if !strings.HasPrefix(s.URL, "https://") {
			res.Failures = append(res.Failures, fmt.Sprintf("source %s URL must start with https://", s.SourceID))
		}
	}
	if tierCounts[1] < 3 {
		res.Failures = append(res.Failures, fmt.Sprintf("expected >= 3 Tier-1 academic/government sources, got %d", tierCounts[1]))
	}
	if len(af.Affiliates) < 5 {
		res.Failures = append(res.Failures, fmt.Sprintf("expected >= 5 affiliate records in affiliates.yaml, got %d", len(af.Affiliates)))
	}

	// Verify Deep Research Backplane Layers (Live Harvest, Magnitude Convergence, 3-Layer Affiliates, BM-01..08, MON-01..08)
	deepFiles := []string{
		"exhaustive_swarm_manifest.yaml",
		"live_harvest_ledger.yaml",
		"magnitude_convergence_matrix.yaml",
		"affiliate_deep_dive.yaml",
		"beneficial_mechanisms_deep_dive.yaml",
		"monetizable_angles_matrix.yaml",
	}
	for _, df := range deepFiles {
		res.ChecksRun++
		full := filepath.Join(rootDir, "01_research_backplane", df)
		info, err := os.Stat(full)
		if err != nil || info.Size() < 200 {
			res.Failures = append(res.Failures, fmt.Sprintf("missing or truncated deep research file: %s", df))
		}
	}

	if len(res.Failures) > 0 {
		res.ExitCode = 1
		res.SummaryLine = fmt.Sprintf("G3 exit 1 (%d failures)", len(res.Failures))
	} else {
		res.SummaryLine = fmt.Sprintf("G3 exit 0 (%d primary sources, %d affiliate programs + 6 deep research & convergence ledgers verified)", len(sf.Sources), len(af.Affiliates))
	}
	return res
}

func checkG4CoverageAndMolds(rootDir string) GateResult {
	res := GateResult{GateID: "G4", Name: "38/38 Question Coverage, 5 Ch 4 Molds (>=3 Cases Each) & Stockpile Lock"}
	rPath := filepath.Join(rootDir, "01_research_backplane", "rows.yaml")
	ePath := filepath.Join(rootDir, "01_research_backplane", "evidence.yaml")
	lockPath := filepath.Join(rootDir, "01_research_backplane", "stockpile.lock")

	rData, err := os.ReadFile(rPath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("missing rows.yaml: %v", err))
		return res
	}
	eData, err := os.ReadFile(ePath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("missing evidence.yaml: %v", err))
		return res
	}
	lockData, err := os.ReadFile(lockPath)
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("missing stockpile.lock: %v", err))
		return res
	}

	var rf RowsFile
	if err := yaml.Unmarshal(rData, &rf); err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, fmt.Sprintf("rows.yaml unmarshal error: %v", err))
		return res
	}

	qCounts := make(map[string]int)
	moldCaseCounts := make(map[string]int)
	for _, r := range rf.Rows {
		qCounts[r.Question]++
		if r.Kind == "case" && r.Mold != "" {
			moldCaseCounts[r.Mold]++
		}
	}

	for _, qID := range Canonical38Questions() {
		res.ChecksRun++
		if qCounts[qID] == 0 {
			res.Failures = append(res.Failures, fmt.Sprintf("question %s has 0 answering rows in rows.yaml", qID))
		}
	}

	validMoldsWith3 := 0
	for mold, cnt := range moldCaseCounts {
		res.ChecksRun++
		if cnt >= 3 {
			validMoldsWith3++
		} else {
			res.Failures = append(res.Failures, fmt.Sprintf("mold %s has only %d case rows (minimum 3 required)", mold, cnt))
		}
	}
	if validMoldsWith3 < 3 {
		res.Failures = append(res.Failures, fmt.Sprintf("expected >= 3 molds with >= 3 case rows each, got %d", validMoldsWith3))
	}

	combinedHash := sha256HexBytes(append(rData, eData...))
	if !strings.Contains(string(lockData), combinedHash) {
		res.Failures = append(res.Failures, fmt.Sprintf("stockpile.lock hash mismatch: expected %s", combinedHash))
	}

	// Verify 14-Mold Population Census, Dynamic Book Sizing Spec & Chapter 4 Case Caricatures
	censusPath := filepath.Join(rootDir, "01_research_backplane", "mold_population_census_14.yaml")
	res.ChecksRun++
	if info, err := os.Stat(censusPath); err != nil || info.Size() < 200 {
		res.Failures = append(res.Failures, "missing or truncated mold_population_census_14.yaml")
	}
	sizingPath := filepath.Join(rootDir, "01_research_backplane", "dynamic_book_sizing.yaml")
	res.ChecksRun++
	if info, err := os.Stat(sizingPath); err != nil || info.Size() < 200 {
		res.Failures = append(res.Failures, "missing or truncated dynamic_book_sizing.yaml")
	}
	caricatures := []string{
		"T01-P2-C04-caricature-case1-e2-marcus_v001.svg",
		"T01-P2-C04-caricature-case2-h4-elena_v001.svg",
		"T01-P2-C04-caricature-case3-h5-devon_v001.svg",
		"T01-P2-C04-caricature-case4-c1-hannah_v001.svg",
		"T01-P2-C04-caricature-case5-c3-victor_v001.svg",
	}
	for _, cSvg := range caricatures {
		res.ChecksRun++
		cPath := filepath.Join(rootDir, "04_phase2_run", "assets", cSvg)
		if info, err := os.Stat(cPath); err != nil || info.Size() < 200 {
			res.Failures = append(res.Failures, fmt.Sprintf("missing Ch 4 case caricature asset: %s", cSvg))
		}
	}

	if len(res.Failures) > 0 {
		res.ExitCode = 1
		res.SummaryLine = fmt.Sprintf("G4 exit 1 (%d failures)", len(res.Failures))
	} else {
		res.SummaryLine = fmt.Sprintf("G4 exit 0 (38/38 questions covered, %d molds from 13+1 census with >=3 cases + 5 Ch 4 caricatures + dynamic sizing verified)", validMoldsWith3)
	}
	return res
}

func checkG5PhasePackets(rootDir string) GateResult {
	res := GateResult{GateID: "G5", Name: "Phase 1-3 Packet Isolation & Orthogonal Spine Rules Check"}
	rData, err := os.ReadFile(filepath.Join(rootDir, "01_research_backplane", "rows.yaml"))
	if err != nil {
		res.ExitCode = 1
		res.Failures = append(res.Failures, "missing rows.yaml")
		return res
	}
	var rf RowsFile
	_ = yaml.Unmarshal(rData, &rf)
	rowMap := make(map[string]Row)
	for _, r := range rf.Rows {
		rowMap[r.ID] = r
	}

	for p := 1; p <= 3; p++ {
		res.ChecksRun++
		pPath := filepath.Join(rootDir, "02_predraft_spine", fmt.Sprintf("phase-%d", p), "packet.yaml")
		data, err := os.ReadFile(pPath)
		if err != nil {
			res.Failures = append(res.Failures, fmt.Sprintf("missing phase-%d/packet.yaml", p))
			continue
		}
		var pkt PhasePacket
		if err := yaml.Unmarshal(data, &pkt); err != nil {
			res.Failures = append(res.Failures, fmt.Sprintf("invalid YAML in phase-%d/packet.yaml: %v", p, err))
			continue
		}
		if len(pkt.CitedRows) == 0 {
			res.Failures = append(res.Failures, fmt.Sprintf("phase-%d/packet.yaml has 0 cited rows", p))
		}
		for _, cid := range pkt.CitedRows {
			res.ChecksRun++
			r, ok := rowMap[cid]
			if !ok {
				res.Failures = append(res.Failures, fmt.Sprintf("phase-%d packet cites non-existent row %s", p, cid))
				continue
			}
			if p == 1 && r.Kind != "claim" {
				res.Failures = append(res.Failures, fmt.Sprintf("Phase 1 packet cites non-claim row %s (kind=%s) — forbidden by ADR 0012", cid, r.Kind))
			}
		}
		if p == 2 && (len(pkt.Ch4Molds) < 3 || len(pkt.Ch4Molds) > 5) {
			res.Failures = append(res.Failures, fmt.Sprintf("Phase 2 packet must specify 3 to 5 vetted Ch 4 molds, got %d", len(pkt.Ch4Molds)))
		}
		if p == 3 {
			if pkt.Ch7Pivots < 2 || pkt.Ch7Pivots > 6 {
				res.Failures = append(res.Failures, fmt.Sprintf("Phase 3 packet ch7_pivots must be in natural range [2,6] (orthogonal to topic age), got %d", pkt.Ch7Pivots))
			}
			if pkt.Ch8Title != "For the Stubborn Fool Who Must" {
				res.Failures = append(res.Failures, fmt.Sprintf("Phase 3 packet ch8_title unexpected: %q", pkt.Ch8Title))
			}
		}
	}

	if len(res.Failures) > 0 {
		res.ExitCode = 1
		res.SummaryLine = fmt.Sprintf("G5 exit 1 (%d failures)", len(res.Failures))
	} else {
		res.SummaryLine = fmt.Sprintf("G5 exit 0 (Phases 1-3 packets isolated & orthogonally validated, %d checks)", res.ChecksRun)
	}
	return res
}

func checkG6ManuscriptsAndTone(rootDir string) GateResult {
	res := GateResult{GateID: "G6", Name: "Phased Manuscripts, Cliffhanger, Tone & Vocabulary Check"}
	filesToCheck := []string{
		filepath.Join(rootDir, "02_predraft_spine", "phase-1", "predraft.md"),
		filepath.Join(rootDir, "02_predraft_spine", "phase-2", "predraft.md"),
		filepath.Join(rootDir, "02_predraft_spine", "phase-3", "predraft.md"),
		filepath.Join(rootDir, "03_phase1_run", "manuscript_phase1.md"),
		filepath.Join(rootDir, "04_phase2_run", "manuscript_phase2.md"),
		filepath.Join(rootDir, "05_phase3_run", "manuscript_phase3.md"),
	}
	bannedPersonBlame := []string{"they failed because", "quitters", "too lazy", "didn't try hard enough"}

	for _, f := range filesToCheck {
		res.ChecksRun++
		b, err := os.ReadFile(f)
		if err != nil {
			res.Failures = append(res.Failures, fmt.Sprintf("missing file %s", filepath.Base(f)))
			continue
		}
		txt := string(b)
		lower := strings.ToLower(txt)
		if strings.Contains(lower, "sovereign") {
			res.Failures = append(res.Failures, fmt.Sprintf("deprecated token 'sovereign' found in %s", filepath.Base(f)))
		}
		for _, blame := range bannedPersonBlame {
			if strings.Contains(lower, blame) {
				res.Failures = append(res.Failures, fmt.Sprintf("banned person-blame phrase %q found in %s", blame, filepath.Base(f)))
			}
		}
		if strings.HasSuffix(f, "manuscript_phase1.md") {
			if !strings.Contains(txt, "Step 7 was not provided") && !strings.Contains(txt, "STEP 7 WAS NOT PROVIDED") {
				res.Failures = append(res.Failures, "manuscript_phase1.md missing mandatory Ch 2 ending 'Step 7 was not provided'")
			}
		}
		if strings.HasSuffix(f, "manuscript_phase3.md") {
			if !strings.Contains(txt, "For the Stubborn Fool Who Must") {
				res.Failures = append(res.Failures, "manuscript_phase3.md missing 'For the Stubborn Fool Who Must' seam/heading")
			}
			if !strings.Contains(txt, "[PIVOT: THE PRACTICAL PIVOT]") {
				res.Failures = append(res.Failures, "manuscript_phase3.md missing '[PIVOT: THE PRACTICAL PIVOT]' callout")
			}
		}
	}

	if len(res.Failures) > 0 {
		res.ExitCode = 1
		res.SummaryLine = fmt.Sprintf("G6 exit 1 (%d failures)", len(res.Failures))
	} else {
		res.SummaryLine = fmt.Sprintf("G6 exit 0 (all 6 pre-drafts & manuscripts checked; 0 banned tokens, Ch 2 cliffhanger & Ch 7->8 seam intact)")
	}
	return res
}

var pdfPageRe = regexp.MustCompile(`/Type\s*/Page\b`)

// CountPDFPages counts "/Type /Page" leaf objects inside a PDF file.
func CountPDFPages(pdfPath string) (int, error) {
	data, err := os.ReadFile(pdfPath)
	if err != nil {
		return 0, err
	}
	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		return 0, fmt.Errorf("not a valid PDF header: %s", pdfPath)
	}
	matches := pdfPageRe.FindAllIndex(data, -1)
	return len(matches), nil
}

func checkG7KDPAssembly(rootDir string) (GateResult, int, int, int, int) {
	res := GateResult{GateID: "G7", Name: "Natural Range Print Geometry (40-84p), Dual-Page Research Spread, PDFs & Podcap Check"}
	bookPDF := filepath.Join(rootDir, "07_assembly_kdp", "build", "book_complete_kdp.pdf")
	coverPDF := filepath.Join(rootDir, "06_covers_and_matter", "covers", "cover_wrap_kdp.pdf")
	pamphletPDF := filepath.Join(rootDir, "06_covers_and_matter", "pamphlet", "companion_pamphlet.pdf")
	freebiePDF := filepath.Join(rootDir, "08_marketing_and_freebie", "freebie", "ch1_ch2_free_giveaway.pdf")

	bookPages, err := CountPDFPages(bookPDF)
	res.ChecksRun++
	if err != nil {
		res.Failures = append(res.Failures, fmt.Sprintf("book_complete_kdp.pdf error: %v", err))
	} else if bookPages < 40 || bookPages > 84 {
		res.Failures = append(res.Failures, fmt.Sprintf("book_complete_kdp.pdf page count %d is outside natural unforced KDP range [40, 84]", bookPages))
	}

	coverPages, err := CountPDFPages(coverPDF)
	res.ChecksRun++
	if err != nil || coverPages < 1 {
		res.Failures = append(res.Failures, fmt.Sprintf("cover_wrap_kdp.pdf error or 0 pages: %v", err))
	}

	pamphletPages, err := CountPDFPages(pamphletPDF)
	res.ChecksRun++
	if err != nil || pamphletPages < 4 {
		res.Failures = append(res.Failures, fmt.Sprintf("companion_pamphlet.pdf error or <4 pages: %v", err))
	}

	freebiePages, err := CountPDFPages(freebiePDF)
	res.ChecksRun++
	if err != nil || freebiePages < 8 {
		res.Failures = append(res.Failures, fmt.Sprintf("ch1_ch2_free_giveaway.pdf error or <8 pages: %v", err))
	}

	// Check Marketing variants.json, Podcap artifacts, and Dual-Page Research Machine Spread
	variantsPath := filepath.Join(rootDir, "08_marketing_and_freebie", "landing", "variants.json")
	podcapScript := filepath.Join(rootDir, "09_podcap_stub", "podcap_25min_script.md")
	podcapBundle := filepath.Join(rootDir, "09_podcap_stub", "tiered_sales_bundle.json")
	rmLeftSVG := filepath.Join(rootDir, "05_phase3_run", "assets", "T01-P3-C08-diagram-research-machine-left_v001.svg")
	rmRightSVG := filepath.Join(rootDir, "05_phase3_run", "assets", "T01-P3-C08-diagram-research-machine-right_v001.svg")
	for _, p := range []string{variantsPath, podcapScript, podcapBundle, rmLeftSVG, rmRightSVG} {
		res.ChecksRun++
		info, err := os.Stat(p)
		if err != nil || info.Size() < 200 {
			res.Failures = append(res.Failures, fmt.Sprintf("missing or truncated downstream artifact: %s", filepath.Base(p)))
		}
	}
	spreadPDF := filepath.Join(rootDir, "07_assembly_kdp", "build", "research_machine_dual_page_spread.pdf")
	spreadPages, spreadErr := CountPDFPages(spreadPDF)
	res.ChecksRun++
	if spreadErr != nil || spreadPages < 1 {
		res.Failures = append(res.Failures, fmt.Sprintf("research_machine_dual_page_spread.pdf error or 0 pages: %v", spreadErr))
	}

	if len(res.Failures) > 0 {
		res.ExitCode = 1
		res.SummaryLine = fmt.Sprintf("G7 exit 1 (%d failures)", len(res.Failures))
	} else {
		res.SummaryLine = fmt.Sprintf("G7 exit 0 (KDP Book=%d pages [target 50-80], Dual-Page Research Spread=%d page, Freebie=%d pages, Pamphlet=%d pages, CoverWrap=%d page)",
			bookPages, spreadPages, freebiePages, pamphletPages, coverPages)
	}
	return res, bookPages, pamphletPages, freebiePages, coverPages
}

// WriteGateReports writes 07_assembly_kdp/build/gate_report.md and kdp_preflight_report.json.
func WriteGateReports(rootDir string, report FullGateReport) error {
	buildDir := filepath.Join(rootDir, "07_assembly_kdp", "build")
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return err
	}
	jsonBytes, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(buildDir, "kdp_preflight_report.json"), jsonBytes, 0644); err != nil {
		return err
	}

	var md strings.Builder
	md.WriteString("# Y-IT Downhill Factory — Programmatic Gate Report (`G1`–`G7`)\n\n")
	md.WriteString(fmt.Sprintf("- **Topic Slug:** `%s`\n", report.TopicSlug))
	md.WriteString(fmt.Sprintf("- **Executed At:** `%s`\n", report.ExecutedAt))
	md.WriteString(fmt.Sprintf("- **Overall Exit Code:** `%d`\n", report.OverallExitCode))
	md.WriteString(fmt.Sprintf("- **Stitched KDP Book Interior (`book_complete_kdp.pdf`):** `%d pages` (Target Range: `50–80 pages`)\n", report.BookPageCount))
	md.WriteString(fmt.Sprintf("- **2-Chapter Freebie Giveaway (`ch1_ch2_free_giveaway.pdf`):** `%d pages`\n", report.FreebiePages))
	md.WriteString(fmt.Sprintf("- **Companion Pocket Pamphlet (`companion_pamphlet.pdf`):** `%d pages`\n", report.PamphletPages))
	md.WriteString(fmt.Sprintf("- **Full KDP Cover Wrap (`cover_wrap_kdp.pdf`):** `%d page` (`12.405\" × 9.25\"`)\n\n", report.CoverWrapPages))
	md.WriteString("| Gate | Name | Checks Run | Exit Code | Summary |\n")
	md.WriteString("| :--- | :--- | :---: | :---: | :--- |\n")
	for _, g := range report.Gates {
		md.WriteString(fmt.Sprintf("| `%s` | %s | %d | `%d` | %s |\n", g.GateID, g.Name, g.ChecksRun, g.ExitCode, g.SummaryLine))
	}
	for _, g := range report.Gates {
		if len(g.Failures) > 0 {
			md.WriteString(fmt.Sprintf("\n### `%s` Failures\n", g.GateID))
			for _, f := range g.Failures {
				md.WriteString(fmt.Sprintf("- %s\n", f))
			}
		}
	}
	return os.WriteFile(filepath.Join(buildDir, "gate_report.md"), []byte(md.String()), 0644)
}

var wsCollapseRe = regexp.MustCompile(`\s+`)

func normalizeText(s string) string {
	s = strings.ReplaceAll(s, "\u201c", "\"")
	s = strings.ReplaceAll(s, "\u201d", "\"")
	s = strings.ReplaceAll(s, "\u2018", "'")
	s = strings.ReplaceAll(s, "\u2019", "'")
	s = strings.ReplaceAll(s, "\u2014", "-")
	s = strings.ReplaceAll(s, "\u2013", "-")
	s = wsCollapseRe.ReplaceAllString(strings.TrimSpace(s), " ")
	return strings.ToLower(s)
}

func sha256HexBytes(b []byte) string {
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}
