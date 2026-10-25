package swarm

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// SubstrateSpec defines a free or ultra-low-cost LLM/reasoning worker in the research swarm.
type SubstrateSpec struct {
	ID               string  `yaml:"id" json:"id"`
	Role             string  `yaml:"role" json:"role"`
	ModelIdentifier  string  `yaml:"model_identifier" json:"model_identifier"`
	AuthRoute        string  `yaml:"auth_route" json:"auth_route"`
	EnvKeyName       string  `yaml:"env_key_name" json:"env_key_name"`
	CostTier         string  `yaml:"cost_tier" json:"cost_tier"`
	PerCallHardCapUSD float64 `yaml:"per_call_hard_cap_usd" json:"per_call_hard_cap_usd"`
	MechanismsServed []string `yaml:"mechanisms_served" json:"mechanisms_served"`
	EnvReady         bool    `yaml:"env_ready" json:"env_ready"`
}

// SearchProviderSpec defines a search or neural discovery engine in the multi-engine fan-out.
type SearchProviderSpec struct {
	ID          string   `yaml:"id" json:"id"`
	Name        string   `yaml:"name" json:"name"`
	Endpoint    string   `yaml:"endpoint" json:"endpoint"`
	EnvKeyName  string   `yaml:"env_key_name" json:"env_key_name"`
	CostTier    string   `yaml:"cost_tier" json:"cost_tier"`
	QueryTypes  []string `yaml:"query_types" json:"query_types"`
	EnvReady    bool     `yaml:"env_ready" json:"env_ready"`
}

// CrawlerSpec defines a scraping, DOM-archiving, or reranking tool.
type CrawlerSpec struct {
	ID         string `yaml:"id" json:"id"`
	Name       string `yaml:"name" json:"name"`
	Endpoint   string `yaml:"endpoint" json:"endpoint"`
	EnvKeyName string `yaml:"env_key_name" json:"env_key_name"`
	CostTier   string `yaml:"cost_tier" json:"cost_tier"`
	Purpose    string `yaml:"purpose" json:"purpose"`
	EnvReady   bool   `yaml:"env_ready" json:"env_ready"`
}

// FreeDatasetSpec defines a free public dataset from Google BigQuery, Federal APIs, or global registries.
type FreeDatasetSpec struct {
	ID                string   `yaml:"id" json:"id"`
	Name              string   `yaml:"name" json:"name"`
	Provider          string   `yaml:"provider" json:"provider"`
	DatasetOrEndpoint string   `yaml:"dataset_or_endpoint" json:"dataset_or_endpoint"`
	SourceTier        int      `yaml:"source_tier" json:"source_tier"`
	QuestionsAnswered []string `yaml:"questions_answered" json:"questions_answered"`
	Description       string   `yaml:"description" json:"description"`
}

// ExhaustiveSwarmManifest records the complete multi-engine, multi-dataset research backplane configuration.
type ExhaustiveSwarmManifest struct {
	TopicSlug            string               `yaml:"topic_slug" json:"topic_slug"`
	GeneratedAt          string               `yaml:"generated_at" json:"generated_at"`
	ExecutionStandard    string               `yaml:"execution_standard" json:"execution_standard"`
	PerCallHardCapUSD    float64              `yaml:"per_call_hard_cap_usd" json:"per_call_hard_cap_usd"`
	MinimumSourcesTarget int                  `yaml:"minimum_sources_target" json:"minimum_sources_target"`
	Substrates           []SubstrateSpec      `yaml:"substrates" json:"substrates"`
	SearchEngines        []SearchProviderSpec `yaml:"search_engines" json:"search_engines"`
	CrawlersAndRerankers []CrawlerSpec        `yaml:"crawlers_and_rerankers" json:"crawlers_and_rerankers"`
	FreePublicDatasets   []FreeDatasetSpec    `yaml:"free_public_datasets" json:"free_public_datasets"`
	QueryFanoutMatrix    map[string][]string  `yaml:"query_fanout_matrix" json:"query_fanout_matrix"`
}

// BuildExhaustiveSwarmManifest constructs the topic-universal, multi-engine research swarm manifest
// tapping Flash-Lite, Flash, Gemini Deep Research, Haiku, Local GPU, Exa, Tavily, Parallel, Brave,
// DuckDuckGo, Firecrawl, Jina, Wayback CDX, Google BigQuery Public Datasets, FRED, BLS, SEC EDGAR, and CFTC/FTC.
func BuildExhaustiveSwarmManifest(topicSlug string) ExhaustiveSwarmManifest {
	hasEnv := func(key string) bool {
		if key == "NONE_FREE_PUBLIC" || key == "LOCAL_MSEDGE_BIN" || key == "CLAUDE_CLI_OAUTH" || key == "GEMINI_WEBUI_MCP" || key == "LOCAL_OLLAMA" {
			return true
		}
		return os.Getenv(key) != ""
	}

	substrates := []SubstrateSpec{
		{
			ID:                "SUB-01-FLASH-LITE",
			Role:              "High-Volume Bulk Evidence Extractor & Affiliate Link Harvester",
			ModelIdentifier:   "gemini-2.5-flash-lite",
			AuthRoute:         "Direct REST API (brihag8 free/low-cost project only)",
			EnvKeyName:        "GEMINI_API_KEY",
			CostTier:          "FREE_TIER_OR_SUB_CENT ($0.10/1M in)",
			PerCallHardCapUSD: 0.15,
			MechanismsServed:  []string{"RM-01", "RM-02", "RM-05", "RM-06", "RM-12", "AFFILIATE_EXTRACTOR"},
			EnvReady:          hasEnv("GEMINI_API_KEY") || hasEnv("GOOGLE_API_KEY"),
		},
		{
			ID:                "SUB-02-FLASH-REASONER",
			Role:              "Claim-to-Test Linker, Contradiction Engine & Beneficial Mechanism Synthesizer",
			ModelIdentifier:   "gemini-2.5-flash",
			AuthRoute:         "Direct REST API (brihag8 free/low-cost project only)",
			EnvKeyName:        "GEMINI_API_KEY",
			CostTier:          "FREE_TIER_OR_LOW_COST",
			PerCallHardCapUSD: 0.15,
			MechanismsServed:  []string{"RM-03", "RM-04", "RM-07", "RM-08", "RM-09", "RM-13", "BM-01..BM-08"},
			EnvReady:          hasEnv("GEMINI_API_KEY") || hasEnv("GOOGLE_API_KEY"),
		},
		{
			ID:                "SUB-03-GEMINI-DEEP-RESEARCH",
			Role:              "5-Track Autonomous Deep Research Sweep & NotebookLM 38-Question Assurance Audit",
			ModelIdentifier:   "gemini-deep-research + notebooklm-38q-auditor",
			AuthRoute:         "gemini-webui MCP + Scribe Reduction Pipeline",
			EnvKeyName:        "GEMINI_WEBUI_MCP",
			CostTier:          "INCLUDED_SUBSCRIPTION ($0 marginal)",
			PerCallHardCapUSD: 0.00,
			MechanismsServed:  []string{"TRACK-1-REGULATORY", "TRACK-2-COHORT", "TRACK-3-SELLER-FUNNELS", "TRACK-4-CASE-AUTOPSIES", "TRACK-5-AFFILIATES-AND-PIVOTS"},
			EnvReady:          true,
		},
		{
			ID:                "SUB-04-CLAUDE-HAIKU-AUDITOR",
			Role:              "Independent 20-Fact / 10-Table-Cell Verbatim Trace Checker & Tone Gatekeeper",
			ModelIdentifier:   "claude-haiku-4-5",
			AuthRoute:         "Headless Claude Code CLI OAuth Subscription Route (never ANTHROPIC_API_KEY)",
			EnvKeyName:        "CLAUDE_CLI_OAUTH",
			CostTier:          "INCLUDED_OAUTH_SUBSCRIPTION ($0 API spend)",
			PerCallHardCapUSD: 0.00,
			MechanismsServed:  []string{"TRACE_AUDIT_9_4", "CONTRADICTION_VERIFIER", "TONE_AND_SCHEMA_GATE"},
			EnvReady:          true,
		},
		{
			ID:                "SUB-05-LOCAL-AND-OPEN-WEIGHT-DEV-MODELS",
			Role:              "On-Demand 12GB Dedicated Local VRAM + Elastic Cloud GPU Open-Weight Inference Array (Never Forced; Tapped for Heavy Parsing, Math & Code Verification)",
			ModelIdentifier:   "12b-local-vram-quant | 32b-code-verifier | 70b-open-weight-reasoner",
			AuthRoute:         "Tier 0 Local 12GB VRAM GPU + Tier 1 Open-Weight Inference Pool + Tier 2 Elastic Cloud GPU Array + Tier 3 Open Model Hub",
			EnvKeyName:        "LOCAL_OLLAMA",
			CostTier:          "FREE ($0.00 Local 12GB VRAM) or Sub-Cent On-Demand",
			PerCallHardCapUSD: 0.15,
			MechanismsServed:  []string{"HTML_CLEANUP", "DEDUPLICATION", "INITIAL_RELEVANCE_TRIAGE", "ON_DEMAND_QUANT_SIMULATION", "ADVERSARIAL_MATH_CHECK"},
			EnvReady:          true,
		},
	}

	searchEngines := []SearchProviderSpec{
		{
			ID:         "SEARCH-01-EXA",
			Name:       "Neural Semantic Vector Search Index",
			Endpoint:   "https://api.exa.ai/search",
			EnvKeyName: "EXA_API_KEY",
			CostTier:   "FREE_CREDITS_OR_LOW_COST",
			QueryTypes: []string{"Practitioner post-mortems", "Court & bankruptcy dockets", "Quantitative autopsies", "Non-SEO technical breakdowns"},
			EnvReady:   hasEnv("EXA_API_KEY"),
		},
		{
			ID:         "SEARCH-02-TAVILY",
			Name:       "Agentic Grounding & Fact-Verification Search Index",
			Endpoint:   "https://api.tavily.com/search",
			EnvKeyName: "TAVILY_API_KEY",
			CostTier:   "FREE_TIER_1000_MO_OR_LOW_COST",
			QueryTypes: []string{"RM-04 adversarial fact queries", "Multi-source claim testing (RM-07)", "Picks-and-shovels vendor audits (P3-Q14)"},
			EnvReady:   hasEnv("TAVILY_API_KEY"),
		},
		{
			ID:         "SEARCH-03-PARALLEL",
			Name:       "High-Concurrency Batch Entity-Graph Search API",
			Endpoint:   "https://api.parallel.ai/v1/search",
			EnvKeyName: "PARALLEL_API_KEY",
			CostTier:   "HIGH_CONCURRENCY_BATCH",
			QueryTypes: []string{"Simultaneous 50-URL entity/actor graph expansion (RM-02)", "Cross-competitor pricing & affiliate commission sweeps"},
			EnvReady:   hasEnv("PARALLEL_API_KEY"),
		},
		{
			ID:         "SEARCH-04-BRAVE",
			Name:       "Independent Non-Ad-Biased Web Search Index",
			Endpoint:   "https://api.search.brave.com/res/v1/web/search",
			EnvKeyName: "BRAVE_API_KEY",
			CostTier:   "FREE_TIER_2000_Q_OR_LOW_COST",
			QueryTypes: []string{"Independent web index", "Prop firm payout denial threads", "Uncensored trader complaint logs", "Affiliate partner terms"},
			EnvReady:   hasEnv("BRAVE_API_KEY"),
		},
		{
			ID:         "SEARCH-05-DUCKDUCKGO",
			Name:       "Zero-Cost Statutory & Document Dork Sweeper",
			Endpoint:   "https://html.duckduckgo.com/html/",
			EnvKeyName: "NONE_FREE_PUBLIC",
			CostTier:   "100%_FREE_UNLIMITED",
			QueryTypes: []string{"site:.gov enforcement dorks", "site:cftc.gov / site:ftc.gov / site:sec.gov", "filetype:pdf cohort disclosures", "Zero-key fallback"},
			EnvReady:   true,
		},
	}

	crawlers := []CrawlerSpec{
		{
			ID:         "CRAWL-01-FIRECRAWL",
			Name:       "Headless Dynamic-JS Checkout Funnel & Pricing Extractor",
			Endpoint:   "https://api.firecrawl.dev/v1/scrape",
			EnvKeyName: "FIRECRAWL_API_KEY",
			CostTier:   "FREE_TIER_OR_LOW_COST",
			Purpose:    "Renders dynamic JS checkout pages (S-01..S-06), upsell timers, and affiliate terms into clean Markdown + raw HTML.",
			EnvReady:   hasEnv("FIRECRAWL_API_KEY"),
		},
		{
			ID:         "CRAWL-02-JINA-READER",
			Name:       "Zero-Cost URL-to-Markdown Document Streamer",
			Endpoint:   "https://r.jina.ai/",
			EnvKeyName: "NONE_FREE_PUBLIC",
			CostTier:   "100%_FREE",
			Purpose:    "Zero-cost clean Markdown conversion of regulatory advisories, academic papers, and documentation pages.",
			EnvReady:   true,
		},
		{
			ID:         "CRAWL-03-MSEDGE-DOM",
			Name:       "Native Headless Chromium DOM Archiver",
			Endpoint:   "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe --headless --dump-dom",
			EnvKeyName: "LOCAL_MSEDGE_BIN",
			CostTier:   "100%_FREE_LOCAL",
			Purpose:    "Saves immutable local HTML snapshots into 01_research_backplane/pages/*.html with SHA-256 verification.",
			EnvReady:   true,
		},
		{
			ID:         "CRAWL-04-WAYBACK-CDX",
			Name:       "Temporal CDX Snapshot & Historical DOM Registry",
			Endpoint:   "https://web.archive.org/cdx/search/cdx",
			EnvKeyName: "NONE_FREE_PUBLIC",
			CostTier:   "100%_FREE_PUBLIC",
			Purpose:    "Recovers historical seller pricing pages, deleted guru earnings claims, and Z1-L4 discontinuous rule changes.",
			EnvReady:   true,
		},
		{
			ID:         "CRAWL-05-COHERE-RERANK",
			Name:       "Cross-Encoder Semantic Quantitative Signal Scorer",
			Endpoint:   "https://api.cohere.com/v2/rerank",
			EnvKeyName: "COHERE_API_KEY",
			CostTier:   "TRIAL_OR_SUB_CENT",
			Purpose:    "Scores 100+ harvested pages for quantitative density before passing top passages to extraction.",
			EnvReady:   hasEnv("COHERE_API_KEY"),
		},
	}

	freeDatasets := []FreeDatasetSpec{
		{
			ID:                "DATA-01-BQ-BLS",
			Name:              "Columnar OLAP Public Dataset: U.S. Bureau of Labor Statistics (BLS OEWS & CPI)",
			Provider:          "Columnar OLAP Statistical Lakehouse + BLS.gov",
			DatasetOrEndpoint: "olap-public-data.bls.* / https://api.bls.gov/publicAPI/v2/timeseries/data/",
			SourceTier:        1,
			QuestionsAnswered: []string{"P2-Q12", "P3-Q1", "P3-Q2", "P3-Q3"},
			Description:       "Authoritative U.S. median hourly wage ($29.76/hr) and occupational wage distributions for all Chapter 7 Career/Skill Pivots ($38.90-$55.00/hr).",
		},
		{
			ID:                "DATA-02-BQ-SEC",
			Name:              "Columnar OLAP Public Dataset: SEC EDGAR Quarterly & Annual Financials",
			Provider:          "Columnar OLAP Statistical Lakehouse + SEC EDGAR EFTS",
			DatasetOrEndpoint: "olap-public-data.sec_quarterly_financials / https://efts.sec.gov/LATEST/search-index",
			SourceTier:        1,
			QuestionsAnswered: []string{"P2-Q2", "P2-Q4", "P3-Q14"},
			Description:       "Audited 10-K/10-Q filings for public platforms, brokers, and tool vendors (revenue share, take-rates, customer acquisition cost, ARPU).",
		},
		{
			ID:                "DATA-03-BQ-CFPB-FTC",
			Name:              "Columnar OLAP Public Dataset: CFPB Consumer Complaints & Federal Enforcement Dockets",
			Provider:          "Columnar OLAP Statistical Lakehouse + FTC.gov / CFTC.gov",
			DatasetOrEndpoint: "olap-public-data.cfpb_complaints / https://www.cftc.gov/PressRoom/PressReleases",
			SourceTier:        1,
			QuestionsAnswered: []string{"P2-Q1", "P2-Q5", "P2-Q9", "P3-Q7", "P3-Q15"},
			Description:       "Federal enforcement actions, B-book counterparty disclosures, receivership accounting, and consumer refund/dispute recovery rates.",
		},
		{
			ID:                "DATA-04-BQ-GOOGLE-TRENDS",
			Name:              "Columnar OLAP Public Dataset: Search Demand Velocity Time-Series",
			Provider:          "Columnar OLAP Statistical Lakehouse",
			DatasetOrEndpoint: "olap-public-data.search_trends.top_terms",
			SourceTier:        2,
			QuestionsAnswered: []string{"P1-Q1", "P2-Q9", "P3-Q8"},
			Description:       "Empirical search demand velocity curves to classify topics across Hot (40%), Classic (40%), and Evergreen (20%) bands and track contagion into 'Next Pitches'.",
		},
		{
			ID:                "DATA-05-FRED-SHED",
			Name:              "Federal Reserve Economic Data (FRED) & SHED Household Well-Being Survey",
			Provider:          "Board of Governors of the Federal Reserve System",
			DatasetOrEndpoint: "https://api.stlouisfed.org/fred/ / https://www.federalreserve.gov/consumerscommunities/shed.htm",
			SourceTier:        1,
			QuestionsAnswered: []string{"P2-Q12", "P3-Q13"},
			Description:       "Household liquid emergency savings baseline ($1,000 Stake threshold) and risk-free Treasury/index benchmarks for 'The Boring Benchmark'.",
		},
		{
			ID:                "DATA-06-BIS-ESMA-FCA",
			Name:              "BIS Central Bank Triennial Survey & ESMA/FCA Mandatory Broker Loss Registries",
			Provider:          "Bank for International Settlements (BIS), ESMA, UK FCA",
			DatasetOrEndpoint: "https://stats.bis.org/api/v1/ / https://www.esma.europa.eu/ / https://www.fca.org.uk/",
			SourceTier:        1,
			QuestionsAnswered: []string{"P2-Q1", "P2-Q2", "P2-Q5", "P2-Q6"},
			Description:       "Global OTC FX turnover ($7.5T/day, 93% institutional) and mandated European/UK broker disclosures showing 74%-89% retail account loss rates.",
		},
		{
			ID:                "DATA-07-OPENALEX-ARXIV",
			Name:              "OpenAlex, Crossref, SSRN & Academic Longitudinal Cohort Registry",
			Provider:          "OpenAlex REST API + Crossref DOI Registry + SSRN",
			DatasetOrEndpoint: "https://api.openalex.org/works?search=",
			SourceTier:        1,
			QuestionsAnswered: []string{"P2-Q3", "P2-Q5", "P2-Q6", "P3-Q10"},
			Description:       "Peer-reviewed longitudinal day-trading, e-commerce, and creator-economy cohort studies (N >= 1,000 participants).",
		},
		{
			ID:                "DATA-08-KAGGLE-AND-HF-DATASETS",
			Name:              "On-Demand Elastic GPU & Parquet Cohort Datasets (12GB Local VRAM + Cloud GPU Array)",
			Provider:          "Elastic Cloud GPU & Parquet Dataset Registry + Open-Weight Hub + Columnar OLAP Sandbox",
			DatasetOrEndpoint: "https://huggingface.co/api/datasets",
			SourceTier:        2,
			QuestionsAnswered: []string{"P2-Q3", "P2-Q5", "P2-Q6", "P3-Q10", "P3-Q13"},
			Description:       "On-demand empirical CSV/Parquet datasets (e.g., e-commerce order logs, historical OHLCV tick data, SaaS churn tables, creator RPM distributions) tapped on 12GB Local VRAM or Elastic Cloud GPUs whenever a topic needs custom cohort crunching without forcing unnecessary overhead.",
		},
	}

	queryMatrix := map[string][]string{
		"vector_1_regulatory_tier1_ddg_and_edgar": {
			fmt.Sprintf(`"%s" enforcement action OR complaint OR receivership site:cftc.gov OR site:ftc.gov OR site:sec.gov`, topicSlug),
			fmt.Sprintf(`"retail foreign exchange" OR "funded trader" loss percentage OR unprofitable site:.gov`),
			`bigquery-public-data.bls.wm_series SOC 15-2051 13-2051 median hourly wage`,
		},
		"vector_2_cohort_studies_exa_and_openalex": {
			fmt.Sprintf(`"%s" empirical cohort survival rate net profit after fees N=`, topicSlug),
			`"proprietary trading" evaluation pass rate Phase 1 Phase 2 payout attrition cohort`,
			`retail day trading effective hourly wage after spread commission swap drag`,
		},
		"vector_3_seller_funnels_firecrawl_and_wayback": {
			fmt.Sprintf(`"%s" "funded account" "keep 90%%" challenge fee checkout pricing`, topicSlug),
			`"use code" "off your challenge" prop firm affiliate commission rate CPA`,
			`wayback cdx prop firm terms of service consistency rule news trading slippage`,
		},
		"vector_4_case_autopsies_brave_and_tavily": {
			fmt.Sprintf(`"%s" "spent $" "months" "hours" payout denied OR breached post-mortem`, topicSlug),
			`prop firm affiliate dashboard revenue vs trading profit screenshot OR disclosure`,
			`Automated EA prop pass bot account terminated shared IP signature`,
		},
		"vector_5_affiliates_and_beneficial_mechanisms_parallel": {
			`top financial data analytics SQL Python certification affiliate program CPA commission`,
			`QuantConnect TradingView partner affiliate commission structure cookie window`,
			`institutional FX position sizing drawdown ruin boundary walk-forward backtest protocol`,
		},
	}

	return ExhaustiveSwarmManifest{
		TopicSlug:            topicSlug,
		GeneratedAt:          "2026-10-24T10:25:00Z",
		ExecutionStandard:    "SPEC-2026-EXHAUSTIVE-SWARM-V2 (Multi-Engine + Free Public Datasets + 14-Mold Empirical Population)",
		PerCallHardCapUSD:    0.15,
		MinimumSourcesTarget: 100,
		Substrates:           substrates,
		SearchEngines:        searchEngines,
		CrawlersAndRerankers: crawlers,
		FreePublicDatasets:   freeDatasets,
		QueryFanoutMatrix:    queryMatrix,
	}
}

// NewSafeHTTPClient returns a timeout-bounded pure-Go net/http client for direct REST API calls.
func NewSafeHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 120 * time.Second,
	}
}
