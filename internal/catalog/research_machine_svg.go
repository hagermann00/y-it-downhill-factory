package catalog

// RenderResearchMachineLeftSVG generates the Verso (Left Page, Page 58) vector diagram
// of the Dual-Page Y-IT Empirical Research Machine Spread.
// Uses pure technical & architectural terminology (Columnar OLAP Lakehouse, 12GB VRAM + Elastic
// Cloud GPU Compute, Multi-Index Neural Retrieval, Cryptographic SHA-256 DOM Vault) with zero
// third-party vendor brand name-dropping in reader-facing copy.
func RenderResearchMachineLeftSVG() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 820 1060" width="100%" height="100%">
  <defs>
    <linearGradient id="bgLeft" x1="0" y1="0" x2="0" y2="1">
      <stop offset="0%" stop-color="#0F172A"/>
      <stop offset="100%" stop-color="#1E293B"/>
    </linearGradient>
    <linearGradient id="accentCyan" x1="0" y1="0" x2="1" y2="0">
      <stop offset="0%" stop-color="#0284C7"/>
      <stop offset="100%" stop-color="#38BDF8"/>
    </linearGradient>
    <linearGradient id="accentAmber" x1="0" y1="0" x2="1" y2="0">
      <stop offset="0%" stop-color="#D97706"/>
      <stop offset="100%" stop-color="#F59E0B"/>
    </linearGradient>
  </defs>

  <!-- Background -->
  <rect x="0" y="0" width="820" height="1060" rx="10" fill="url(#bgLeft)" stroke="#334155" stroke-width="3"/>

  <!-- Top Header Banner -->
  <rect x="16" y="16" width="788" height="74" rx="8" fill="#020617" stroke="#38BDF8" stroke-width="2"/>
  <text x="32" y="42" font-family="Arial Black, sans-serif" font-size="15.0" fill="#38BDF8">HOW WE CHECKED THE RECEIPTS · PART I: THE EMPIRICAL INGESTION ENGINE</text>
  <text x="32" y="62" font-family="Consolas, monospace" font-size="10.2" fill="#E2E8F0">WHY TRUST THIS LEDGER? Most $997 courses are researched in the back of a rented supercar. Here is our actual engine:</text>
  <text x="32" y="79" font-family="Consolas, monospace" font-size="9.5" fill="#94A3B8">ARCHITECTURE: 5-Layer Interpretable Context Protocol (L0–L4) · Mechanisms RM-01..RM-06 · Zero-Mock Gate</text>

  <!-- STAGE 1: TOPIC ROUTER & DUAL HYPOTHESIS ENGINE -->
  <rect x="16" y="104" width="788" height="142" rx="8" fill="#0F172A" stroke="#475569" stroke-width="1.8"/>
  <rect x="16" y="104" width="788" height="26" rx="6" fill="#1E3A8A"/>
  <text x="28" y="121" font-family="Arial Black, sans-serif" font-size="11.5" fill="#FFFFFF">STAGE I · UNIVERSAL 54-TOPIC ROUTER &amp; DUAL-HYPOTHESIS PROTOCOL (RM-01 / RM-02 / RM-03)</text>
  <text x="615" y="121" font-family="Consolas, monospace" font-size="10" fill="#93C5FD">9 FAMILIES (F1..F9) -&gt; 44 BOOKS</text>

  <!-- Left box: 9 Topic Families -->
  <rect x="28" y="138" width="250" height="96" rx="6" fill="#1E293B" stroke="#64748B" stroke-width="1.2"/>
  <text x="38" y="155" font-family="Arial, sans-serif" font-size="10.5" font-weight="bold" fill="#38BDF8">1. Dynamic Topic Classification</text>
  <text x="38" y="171" font-family="Consolas, monospace" font-size="9" fill="#CBD5E1">• F1 Ecommerce  • F2 Agencies  • F3 Content</text>
  <text x="38" y="185" font-family="Consolas, monospace" font-size="9" fill="#CBD5E1">• F4 Trading    • F5 Guru/MLM  • F6 Brand</text>
  <text x="38" y="199" font-family="Consolas, monospace" font-size="9" fill="#CBD5E1">• F7 Real Estate• F8 Flipping  • F9 AI/Tech</text>
  <rect x="38" y="207" width="230" height="20" rx="4" fill="#0284C7"/>
  <text x="46" y="221" font-family="Consolas, monospace" font-size="9.5" font-weight="bold" fill="#FFFFFF">ACTIVE: F4 / ARCH-1 SPECULATION_BBOOK</text>

  <!-- Middle box: RM-02 Entity & Actor Graph -->
  <rect x="290" y="138" width="240" height="96" rx="6" fill="#1E293B" stroke="#64748B" stroke-width="1.2"/>
  <text x="300" y="155" font-family="Arial, sans-serif" font-size="10.5" font-weight="bold" fill="#FBBF24">2. RM-02 Structural Actor Graph</text>
  <text x="300" y="171" font-family="Consolas, monospace" font-size="9" fill="#E2E8F0">We don't query seller buzzwords; we map</text>
  <text x="300" y="184" font-family="Consolas, monospace" font-size="9" fill="#E2E8F0">the institutions that hold the ledger:</text>
  <text x="300" y="199" font-family="Consolas, monospace" font-size="8.8" fill="#FDE68A">• Regulators: CFTC, SEC, FTC, ESMA, FCA</text>
  <text x="300" y="212" font-family="Consolas, monospace" font-size="8.8" fill="#FDE68A">• Central Banks &amp; Labor: BIS, Fed, BLS</text>
  <text x="300" y="225" font-family="Consolas, monospace" font-size="8.8" fill="#FDE68A">• Clearing &amp; Server Rails: B-Book, Card Nets</text>

  <!-- Right box: Dual Hypotheses -->
  <rect x="542" y="138" width="250" height="96" rx="6" fill="#1E293B" stroke="#64748B" stroke-width="1.2"/>
  <text x="552" y="155" font-family="Arial, sans-serif" font-size="10.5" font-weight="bold" fill="#4ADE80">3. Anti-Bias Dual Hypotheses</text>
  <text x="552" y="172" font-family="Consolas, monospace" font-size="8.8" fill="#86EFAC">H0 (Bull Test): Under what exact skill,</text>
  <text x="552" y="184" font-family="Consolas, monospace" font-size="8.8" fill="#86EFAC">capital, or audience edge does it work?</text>
  <text x="552" y="201" font-family="Consolas, monospace" font-size="8.8" fill="#FCA5A5">H1 (Bear Test): Where do drawdown</text>
  <text x="552" y="213" font-family="Consolas, monospace" font-size="8.8" fill="#FCA5A5">barriers, B-books &amp; fees break starters?</text>
  <text x="552" y="227" font-family="Consolas, monospace" font-size="8.5" fill="#94A3B8">Both tested across 38 canonical questions.</text>

  <!-- Down Arrow 1 -> 2 -->
  <polygon points="410,248 400,258 420,258" fill="#38BDF8"/>
  <line x1="410" y1="246" x2="410" y2="262" stroke="#38BDF8" stroke-width="3"/>

  <!-- STAGE 2: EXHAUSTIVE MULTI-INDEX & STATISTICAL LAKEHOUSE SWARM -->
  <rect x="16" y="262" width="788" height="270" rx="8" fill="#0F172A" stroke="#38BDF8" stroke-width="1.8"/>
  <rect x="16" y="262" width="788" height="26" rx="6" fill="url(#accentCyan)"/>
  <text x="28" y="279" font-family="Arial Black, sans-serif" font-size="11.5" fill="#020617">STAGE II · MULTI-INDEX RETRIEVAL &amp; STATISTICAL LAKEHOUSE SWARM (RM-04 / RM-05 · 36 LIVE PROBES)</text>
  <text x="615" y="279" font-family="Consolas, monospace" font-size="10" font-weight="bold" fill="#020617">140+ VALIDATED SOURCES</text>

  <!-- 4 Columns of Source Providers (Technical Names Only) -->
  <!-- Col 1: Tier 1 Gov & Statutory -->
  <rect x="28" y="298" width="184" height="172" rx="6" fill="#1E293B" stroke="#0EA5E9" stroke-width="1.5"/>
  <rect x="28" y="298" width="184" height="22" rx="4" fill="#0369A1"/>
  <text x="36" y="313" font-family="Arial Black, sans-serif" font-size="9.5" fill="#FFFFFF">TIER 1 · STATUTORY &amp; GOV</text>
  <text x="36" y="332" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• U.S. Federal Register API</text>
  <text x="36" y="346" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• CFTC RFED Loss Filings</text>
  <text x="36" y="360" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• CFTC Enforcement Dockets</text>
  <text x="36" y="374" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• BIS Triennial FX ($7.5T)</text>
  <text x="36" y="388" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• ESMA &amp; UK FCA Registries</text>
  <text x="36" y="402" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• SEC EDGAR 10-K/10-Q Index</text>
  <text x="36" y="416" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• FTC Business Opp Dockets</text>
  <rect x="36" y="428" width="168" height="34" rx="4" fill="#0F172A"/>
  <text x="42" y="442" font-family="Consolas, monospace" font-size="8.3" fill="#38BDF8">Rule: Tier 1 stands alone on</text>
  <text x="42" y="454" font-family="Consolas, monospace" font-size="8.3" fill="#38BDF8">statutory &amp; macro facts.</text>

  <!-- Col 2: Columnar OLAP Statistical Lakehouse & Parquet Clusters -->
  <rect x="221" y="298" width="184" height="172" rx="6" fill="#1E293B" stroke="#10B981" stroke-width="1.5"/>
  <rect x="221" y="298" width="184" height="22" rx="4" fill="#047857"/>
  <text x="229" y="313" font-family="Arial Black, sans-serif" font-size="9.5" fill="#FFFFFF">OLAP &amp; PARQUET LAKEHOUSE</text>
  <text x="229" y="332" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Columnar OLAP Warehouse:</text>
  <text x="229" y="346" font-family="Consolas, monospace" font-size="8.5" fill="#A7F3D0">  - National Labor Wage Tables</text>
  <text x="229" y="360" font-family="Consolas, monospace" font-size="8.5" fill="#A7F3D0">  - Quarterly Corporate Filings</text>
  <text x="229" y="374" font-family="Consolas, monospace" font-size="8.5" fill="#A7F3D0">  - Federal Complaint Ledgers</text>
  <text x="229" y="388" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• U.S. BLS OEWS ($29.76/hr)</text>
  <text x="229" y="402" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Fed Reserve SHED ($1,000)</text>
  <text x="229" y="416" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Elastic Parquet Cohort Sets</text>
  <rect x="229" y="428" width="168" height="34" rx="4" fill="#0F172A"/>
  <text x="235" y="442" font-family="Consolas, monospace" font-size="8.3" fill="#34D399">Elastic: Tapped on demand,</text>
  <text x="235" y="454" font-family="Consolas, monospace" font-size="8.3" fill="#34D399">never forced when not needed.</text>

  <!-- Col 3: Tier 1/2 Academic & Large Cohorts -->
  <rect x="414" y="298" width="184" height="172" rx="6" fill="#1E293B" stroke="#F59E0B" stroke-width="1.5"/>
  <rect x="414" y="298" width="184" height="22" rx="4" fill="#B45309"/>
  <text x="422" y="313" font-family="Arial Black, sans-serif" font-size="9.5" fill="#FFFFFF">TIER 1/2 · COHORT STUDIES</text>
  <text x="422" y="332" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Academic Graph API (80+)</text>
  <text x="422" y="346" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Crossref DOI Registry (36)</text>
  <text x="422" y="360" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Quantitative Finance Archive</text>
  <text x="422" y="374" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Empirical Cohort Index</text>
  <text x="422" y="388" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• 420,000-Account Prop</text>
  <text x="422" y="402" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">  Evaluation Cohort Study</text>
  <text x="422" y="416" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Multi-Year Day-Trade Logs</text>
  <rect x="422" y="428" width="168" height="34" rx="4" fill="#0F172A"/>
  <text x="428" y="442" font-family="Consolas, monospace" font-size="8.3" fill="#FBBF24">Rule: Tier 2 requires &gt;=2</text>
  <text x="428" y="454" font-family="Consolas, monospace" font-size="8.3" fill="#FBBF24">independent non-seller sources.</text>

  <!-- Col 4: Tier 3/4 Neural Search, DOM Crawl & Funnels -->
  <rect x="607" y="298" width="185" height="172" rx="6" fill="#1E293B" stroke="#EC4899" stroke-width="1.5"/>
  <rect x="607" y="298" width="185" height="22" rx="4" fill="#BE185D"/>
  <text x="615" y="313" font-family="Arial Black, sans-serif" font-size="9.5" fill="#FFFFFF">TIER 3/4 · DOM &amp; FUNNELS</text>
  <text x="615" y="332" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Neural Semantic Vector Index</text>
  <text x="615" y="346" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Parallel Multi-Index Search</text>
  <text x="615" y="360" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Federal .gov Query Sweeper</text>
  <text x="615" y="374" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Headless JS &amp; DOM Renderer</text>
  <text x="615" y="388" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Temporal CDX Snapshot API</text>
  <text x="615" y="402" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Captured Sellers S-01..S-06</text>
  <text x="615" y="416" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Open-Source Code Registries</text>
  <rect x="615" y="428" width="169" height="34" rx="4" fill="#0F172A"/>
  <text x="621" y="442" font-family="Consolas, monospace" font-size="8.3" fill="#F472B6">Rule: Seller pages = Claims</text>
  <text x="621" y="454" font-family="Consolas, monospace" font-size="8.3" fill="#F472B6">only (never cited as Facts).</text>

  <!-- Bottom bar inside Stage II: Dedicated Local VRAM + Elastic Cloud GPU & Reasoning Array -->
  <rect x="28" y="478" width="764" height="44" rx="6" fill="#020617" stroke="#475569" stroke-width="1.2"/>
  <text x="38" y="495" font-family="Arial Black, sans-serif" font-size="9.5" fill="#38BDF8">HYBRID COMPUTE &amp; INFERENCE ARRAY (12GB DEDICATED LOCAL VRAM + ELASTIC CLOUD GPU CLUSTER):</text>
  <text x="38" y="512" font-family="Consolas, monospace" font-size="8.5" fill="#CBD5E1">[High-Throughput Bulk Atomizer] · [Cross-Source Claim Linker] · [38-Question Assurance Auditor] · [Independent Trace Verifier] · [Local 12GB VRAM &amp; Elastic GPU Bench]</text>

  <!-- Down Arrow 2 -> 3 -->
  <polygon points="410,536 400,546 420,546" fill="#38BDF8"/>
  <line x1="410" y1="532" x2="410" y2="550" stroke="#38BDF8" stroke-width="3"/>

  <!-- STAGE 3: IMMUTABLE DOM ARCHIVE, SHA-256 FINGERPRINTING & PRE-TRIM AFFILIATE HARVEST -->
  <rect x="16" y="550" width="788" height="196" rx="8" fill="#0F172A" stroke="#F59E0B" stroke-width="1.8"/>
  <rect x="16" y="550" width="788" height="26" rx="6" fill="url(#accentAmber)"/>
  <text x="28" y="567" font-family="Arial Black, sans-serif" font-size="11.5" fill="#020617">STAGE III · CRYPTOGRAPHIC DOM VAULT &amp; PRE-TRIM AFFILIATE HARVESTER (01_research_backplane/)</text>
  <text x="630" y="567" font-family="Consolas, monospace" font-size="10" font-weight="bold" fill="#020617">SHA-256 CONTENT LOCK</text>

  <!-- Left box: Raw HTML / Probe Vault -->
  <rect x="28" y="586" width="372" height="148" rx="6" fill="#1E293B" stroke="#F59E0B" stroke-width="1.2"/>
  <text x="38" y="604" font-family="Arial, sans-serif" font-size="10.5" font-weight="bold" fill="#FBBF24">1. Immutable Local Source Archive (pages/*.html)</text>
  <text x="38" y="620" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">Every source page &amp; API payload is saved locally before</text>
  <text x="38" y="633" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">extraction so URLs can never rot or change behind our backs:</text>
  <text x="38" y="650" font-family="Consolas, monospace" font-size="8.5" fill="#FDE68A">• pages/src_01..src_09.html (9 Primary Canonical Archives)</text>
  <text x="38" y="664" font-family="Consolas, monospace" font-size="8.5" fill="#FDE68A">• pages/live_probes/probe-01..36.json (36 Live Sweeps)</text>
  <text x="38" y="678" font-family="Consolas, monospace" font-size="8.5" fill="#FDE68A">• stockpile.lock -&gt; SHA-256 digest over rows + evidence</text>
  <rect x="38" y="688" width="352" height="36" rx="4" fill="#020617"/>
  <text x="46" y="703" font-family="Consolas, monospace" font-size="8.3" fill="#4ADE80">GATE G2 LAW: Every cited passage in evidence.yaml must</text>
  <text x="46" y="716" font-family="Consolas, monospace" font-size="8.3" fill="#4ADE80">match character-for-character inside pages/*.html.</text>

  <!-- Right box: Pre-Trim Affiliate & Funnel Extraction -->
  <rect x="414" y="586" width="378" height="148" rx="6" fill="#1E293B" stroke="#F59E0B" stroke-width="1.2"/>
  <text x="424" y="604" font-family="Arial, sans-serif" font-size="10.5" font-weight="bold" fill="#FBBF24">2. Pre-Trim Affiliate &amp; Seller Funnel Capture</text>
  <text x="424" y="620" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">Before stripping HTML boilerplate, the harvester extracts</text>
  <text x="424" y="633" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">every referral code, IB rebate, and upsell rung:</text>
  <text x="424" y="650" font-family="Consolas, monospace" font-size="8.5" fill="#FCA5A5">• Layer 1A (Exposed): 20%-25% Prop Challenge CPA loops,</text>
  <text x="424" y="663" font-family="Consolas, monospace" font-size="8.5" fill="#FCA5A5">  $6-$15/lot B-book IB rebates, $68/mo indicator kickbacks</text>
  <text x="424" y="678" font-family="Consolas, monospace" font-size="8.5" fill="#93C5FD">• Layer 1B (P3-Q14): Picks-and-Shovels Honesty/Ethics Audit</text>
  <text x="424" y="693" font-family="Consolas, monospace" font-size="8.5" fill="#86EFAC">• Layer 1C (Own-Profit): Vetted SQL/Python BI &amp; Exit Tools</text>
  <text x="424" y="716" font-family="Consolas, monospace" font-size="8.3" fill="#94A3B8">Anonymized Seller Keys (S-01..S-06): System, never person.</text>

  <!-- Down Arrow 3 -> 4 -->
  <polygon points="410,750 400,760 420,760" fill="#F59E0B"/>
  <line x1="410" y1="746" x2="410" y2="764" stroke="#F59E0B" stroke-width="3"/>

  <!-- STAGE 4: 15-FIELD EVIDENCE ATOMIZER & 38-QUESTION MATRIX -->
  <rect x="16" y="764" width="788" height="276" rx="8" fill="#0F172A" stroke="#10B981" stroke-width="2"/>
  <rect x="16" y="764" width="788" height="26" rx="6" fill="#059669"/>
  <text x="28" y="781" font-family="Arial Black, sans-serif" font-size="11.5" fill="#FFFFFF">STAGE IV · 15-FIELD EVIDENCE ATOMIZER &amp; 38-QUESTION MATRIX (RM-06 · 53 LOCKED ROWS)</text>
  <text x="615" y="781" font-family="Consolas, monospace" font-size="10" font-weight="bold" fill="#A7F3D0">38 / 38 QUESTIONS ANSWERED</text>

  <!-- 3 Atomized Row Streams -->
  <rect x="28" y="800" width="242" height="172" rx="6" fill="#1E293B" stroke="#F43F5E" stroke-width="1.5"/>
  <text x="38" y="819" font-family="Arial Black, sans-serif" font-size="10" fill="#FDA4AF">CLAIM ROWS (C-001..C-014)</text>
  <text x="38" y="835" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• 14 Seller Pitch Claims</text>
  <text x="38" y="850" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Answers P1-Q1 .. P1-Q12</text>
  <text x="38" y="865" font-family="Consolas, monospace" font-size="8.8" fill="#FDA4AF">• Feeds Phase 1 (Ch 1 &amp; Ch 2)</text>
  <text x="38" y="882" font-family="Consolas, monospace" font-size="8.3" fill="#CBD5E1">Examples: "$200k Funded Account",</text>
  <text x="38" y="895" font-family="Consolas, monospace" font-size="8.3" fill="#CBD5E1">"$10k/mo in 45 mins", "Keep 90%"</text>
  <rect x="38" y="908" width="222" height="52" rx="4" fill="#020617"/>
  <text x="44" y="924" font-family="Consolas, monospace" font-size="8.2" fill="#F43F5E">QUARANTINE RULE: Phase 1 cites</text>
  <text x="44" y="937" font-family="Consolas, monospace" font-size="8.2" fill="#F43F5E">ONLY C-### rows. Every numbered</text>
  <text x="44" y="950" font-family="Consolas, monospace" font-size="8.2" fill="#F43F5E">claim links to a testing F-### row.</text>

  <rect x="284" y="800" width="246" height="172" rx="6" fill="#1E293B" stroke="#38BDF8" stroke-width="1.5"/>
  <text x="294" y="819" font-family="Arial Black, sans-serif" font-size="10" fill="#7DD3FC">FACT ROWS (F-001..F-027)</text>
  <text x="294" y="835" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• 24 Audited Primary Facts</text>
  <text x="294" y="850" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Answers P2-Q1..Q12 &amp; P3-Q1..Q14</text>
  <text x="294" y="865" font-family="Consolas, monospace" font-size="8.8" fill="#7DD3FC">• Feeds Phases 2 &amp; 3 (Ch 3,5,6,7,8)</text>
  <text x="294" y="882" font-family="Consolas, monospace" font-size="8.3" fill="#CBD5E1">Examples: $1,840 Stake, 98.4% B-book,</text>
  <text x="294" y="895" font-family="Consolas, monospace" font-size="8.3" fill="#CBD5E1">2.4% 12mo Odds, $6.40/hr vs $29.76/hr</text>
  <rect x="294" y="908" width="226" height="52" rx="4" fill="#020617"/>
  <text x="300" y="924" font-family="Consolas, monospace" font-size="8.2" fill="#38BDF8">CONVERGENCE RULE: Every core</text>
  <text x="300" y="937" font-family="Consolas, monospace" font-size="8.2" fill="#38BDF8">metric requires &gt;=3 corroborating</text>
  <text x="300" y="950" font-family="Consolas, monospace" font-size="8.2" fill="#38BDF8">sources of matching magnitude.</text>

  <rect x="544" y="800" width="248" height="172" rx="6" fill="#1E293B" stroke="#FBBF24" stroke-width="1.5"/>
  <text x="554" y="819" font-family="Arial Black, sans-serif" font-size="10" fill="#FDE68A">CASE ROWS (K-001..K-015)</text>
  <text x="554" y="835" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• 15 Corroborated Case Dossiers</text>
  <text x="554" y="850" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">• Tagged across 13+1 Molds (RM-12)</text>
  <text x="554" y="865" font-family="Consolas, monospace" font-size="8.8" fill="#FDE68A">• Feeds Ch 4 Forensic Insert &amp; Ch 8</text>
  <text x="554" y="882" font-family="Consolas, monospace" font-size="8.3" fill="#CBD5E1">35 Raw Cases Evaluated -&gt; Top 5</text>
  <text x="554" y="895" font-family="Consolas, monospace" font-size="8.3" fill="#CBD5E1">Molds Locked (3 Cases Per Mold)</text>
  <rect x="554" y="908" width="228" height="52" rx="4" fill="#020617"/>
  <text x="560" y="924" font-family="Consolas, monospace" font-size="8.2" fill="#FBBF24">COMPOSITE LAW: No real person's</text>
  <text x="560" y="937" font-family="Consolas, monospace" font-size="8.2" fill="#FBBF24">name printed; every persona is the</text>
  <text x="560" y="950" font-family="Consolas, monospace" font-size="8.2" fill="#FBBF24">exact median of &gt;=3 real cases.</text>

  <!-- Bottom Gutter Bridge Banner pointing right to Page 59 -->
  <rect x="28" y="982" width="764" height="46" rx="6" fill="#064E3B" stroke="#34D399" stroke-width="1.8"/>
  <text x="42" y="1002" font-family="Arial Black, sans-serif" font-size="11" fill="#A7F3D0">PIPELINE BRIDGE ACROSS THE GUTTER: 53 ATOMIZED ROWS + 140+ LIVE SOURCES ───────────────►</text>
  <text x="42" y="1019" font-family="Consolas, monospace" font-size="9.2" fill="#ECFDF5">Continues onto Right Page (p. 59): Magnitude Convergence Sieve (RM-10v2) -&gt; 14-Mold Empirical Census -&gt; Compiler Gates G1-G7</text>
</svg>`
}

// RenderResearchMachineRightSVG generates the Recto (Right Page, Page 59) vector diagram
// of the Dual-Page Y-IT Empirical Research Machine Spread.
func RenderResearchMachineRightSVG() string {
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 820 1060" width="100%" height="100%">
  <defs>
    <linearGradient id="bgRight" x1="0" y1="0" x2="0" y2="1">
      <stop offset="0%" stop-color="#0F172A"/>
      <stop offset="100%" stop-color="#1E293B"/>
    </linearGradient>
    <linearGradient id="accentGreen" x1="0" y1="0" x2="1" y2="0">
      <stop offset="0%" stop-color="#059669"/>
      <stop offset="100%" stop-color="#10B981"/>
    </linearGradient>
    <linearGradient id="accentPurple" x1="0" y1="0" x2="1" y2="0">
      <stop offset="0%" stop-color="#7C3AED"/>
      <stop offset="100%" stop-color="#A855F7"/>
    </linearGradient>
  </defs>

  <!-- Background -->
  <rect x="0" y="0" width="820" height="1060" rx="10" fill="url(#bgRight)" stroke="#334155" stroke-width="3"/>

  <!-- Top Header Banner -->
  <rect x="16" y="16" width="788" height="74" rx="8" fill="#020617" stroke="#10B981" stroke-width="2"/>
  <text x="32" y="42" font-family="Arial Black, sans-serif" font-size="15.0" fill="#34D399">HOW WE CHECKED THE RECEIPTS · PART II: CONVERGENCE, MOLD SIEVE &amp; GATES</text>
  <text x="32" y="62" font-family="Consolas, monospace" font-size="10.2" fill="#E2E8F0">PLAIN ENGLISH: Why we use multi-source ranges instead of hype, how cases pick themselves, and how code checks every page.</text>
  <text x="32" y="79" font-family="Consolas, monospace" font-size="9.5" fill="#94A3B8">METHODOLOGY: RM-07..RM-13 · RM-10v2 Range-First Convergence · 21-&gt;13+1 Mold Census · Compiler Gates G1–G7</text>

  <!-- STAGE 5: RM-10v2 MAGNITUDE CONVERGENCE & RANGE-FIRST SIEVE -->
  <rect x="16" y="104" width="788" height="232" rx="8" fill="#0F172A" stroke="#38BDF8" stroke-width="1.8"/>
  <rect x="16" y="104" width="788" height="26" rx="6" fill="#0284C7"/>
  <text x="28" y="121" font-family="Arial Black, sans-serif" font-size="11.5" fill="#FFFFFF">◄── STAGE V · MAGNITUDE CONVERGENCE &amp; RANGE-FIRST SIEVE (RM-07 / RM-08 / RM-10v2)</text>
  <text x="595" y="121" font-family="Consolas, monospace" font-size="10" font-weight="bold" fill="#E0F2FE">&gt;= 3 INDEPENDENT SOURCES</text>

  <!-- Left box: Disqualified Outliers -->
  <rect x="28" y="138" width="235" height="186" rx="6" fill="#1E293B" stroke="#EF4444" stroke-width="1.5"/>
  <rect x="28" y="138" width="235" height="22" rx="4" fill="#991B1B"/>
  <text x="36" y="153" font-family="Arial Black, sans-serif" font-size="9.2" fill="#FFFFFF">QUARANTINED (disqualified.yaml)</text>
  <text x="36" y="172" font-family="Consolas, monospace" font-size="8.6" fill="#FCA5A5">✗ Single-source outlier claims</text>
  <text x="36" y="186" font-family="Consolas, monospace" font-size="8.6" fill="#FCA5A5">✗ Cropped Demo MT4/MT5 screens</text>
  <text x="36" y="200" font-family="Consolas, monospace" font-size="8.6" fill="#FCA5A5">✗ Unverified blog roundups</text>
  <text x="36" y="214" font-family="Consolas, monospace" font-size="8.6" fill="#FCA5A5">✗ Hyperbolic worst-case anecdotes</text>
  <rect x="36" y="226" width="219" height="88" rx="4" fill="#020617"/>
  <text x="42" y="242" font-family="Consolas, monospace" font-size="8.2" fill="#F87171">ANTI-HYPERBOLE LAW:</text>
  <text x="42" y="255" font-family="Consolas, monospace" font-size="8.0" fill="#CBD5E1">We never state a single extreme</text>
  <text x="42" y="267" font-family="Consolas, monospace" font-size="8.0" fill="#CBD5E1">number as a rigid fact unless</text>
  <text x="42" y="279" font-family="Consolas, monospace" font-size="8.0" fill="#CBD5E1">multiple independent sources</text>
  <text x="42" y="291" font-family="Consolas, monospace" font-size="8.0" fill="#CBD5E1">converge on the same magnitude.</text>
  <text x="42" y="305" font-family="Consolas, monospace" font-size="8.0" fill="#34D399">Ranges + Medians = Unassailable.</text>

  <!-- Right box: The 4 Convergent Gate Ranges + Midpoints -->
  <rect x="275" y="138" width="517" height="186" rx="6" fill="#1E293B" stroke="#38BDF8" stroke-width="1.5"/>
  <rect x="275" y="138" width="517" height="22" rx="4" fill="#0369A1"/>
  <text x="285" y="153" font-family="Arial Black, sans-serif" font-size="9.5" fill="#FFFFFF">AUDITED CONVERGENT RANGES &amp; COHORT MEDIANS (magnitude_convergence_matrix.yaml)</text>

  <text x="285" y="174" font-family="Consolas, monospace" font-size="9.2" font-weight="bold" fill="#38BDF8">1. THE STAKE (Cash Outlay):</text>
  <text x="470" y="174" font-family="Consolas, monospace" font-size="9.2" fill="#FFFFFF">$1,400–$2,200 Range (Median: $1,840 vs $1,000 Fed Buffer) [MISS]</text>
  <text x="285" y="187" font-family="Consolas, monospace" font-size="8.2" fill="#94A3B8">   Iterated in: 420k Cohort ($1,840 [F-001]), CFTC Receivership (~$2,296 avg [F-002]), Cases ($920–$4,110 [K-004..K-015])</text>

  <text x="285" y="207" font-family="Consolas, monospace" font-size="9.2" font-weight="bold" fill="#38BDF8">2. THE WAIT (Time to Net +):</text>
  <text x="470" y="207" font-family="Consolas, monospace" font-size="9.2" fill="#FFFFFF">12.0–18.0 Months Range (Median: 14.0 mo vs &lt;=6.0 mo Gate) [MISS]</text>
  <text x="285" y="220" font-family="Consolas, monospace" font-size="8.2" fill="#94A3B8">   Iterated in: 420k Survivor Cohort (14.0m [F-009]), Mold C1 Cases (12–16m [K-010..K-012]), RFED 4.5m Reloads [F-012]</text>

  <text x="285" y="240" font-family="Consolas, monospace" font-size="9.2" font-weight="bold" fill="#38BDF8">3. THE ODDS (12-Mo Survival):</text>
  <text x="470" y="240" font-family="Consolas, monospace" font-size="9.2" fill="#FFFFFF">5%–8% Pass (6.6%) · 1.5%–3.5% Net Profitable at 12m (2.4%) [MISS]</text>
  <text x="285" y="253" font-family="Consolas, monospace" font-size="8.2" fill="#94A3B8">   Iterated in: 420k Cohort (2.4% [F-008]), CFTC B-Book (91.6% cut &lt;Payout #2 [F-010]), CFTC 72% &amp; ESMA/FCA 74%–89% Loss [F-005]</text>

  <text x="285" y="273" font-family="Consolas, monospace" font-size="9.2" font-weight="bold" fill="#38BDF8">4. THE PAYOFF (Survivor $/Hr):</text>
  <text x="470" y="273" font-family="Consolas, monospace" font-size="9.2" fill="#FFFFFF">$4.50–$8.50/Hr Range (Median: $6.40/hr vs $29.76/hr BLS) [MISS]</text>
  <text x="285" y="286" font-family="Consolas, monospace" font-size="8.2" fill="#94A3B8">   Iterated in: Survivor Time-Logs ($6.40/hr @ 26.5h/wk [F-003]), Day-Trade Study (94.2% &lt; min wage [F-005]), Mold C1 ($1.11/hr)</text>

  <rect x="285" y="296" width="497" height="20" rx="4" fill="#020617"/>
  <text x="293" y="310" font-family="Consolas, monospace" font-size="8.5" fill="#FBBF24">STATUTORY POINT FACTS: $7.5T/day BIS FX Volume (93% Institutional) · 98.4% B-Book Server · $29.76/hr BLS Wage</text>

  <!-- Down Arrow 5 -> 6 -->
  <polygon points="410,340 400,350 420,350" fill="#A855F7"/>
  <line x1="410" y1="336" x2="410" y2="354" stroke="#A855F7" stroke-width="3"/>

  <!-- STAGE 6: RM-12 EMPIRICAL MOLD POPULATION CENSUS (21 -> 13+1 -> TOP 5 SELECTED) -->
  <rect x="16" y="354" width="788" height="236" rx="8" fill="#0F172A" stroke="#A855F7" stroke-width="1.8"/>
  <rect x="16" y="354" width="788" height="26" rx="6" fill="url(#accentPurple)"/>
  <text x="28" y="371" font-family="Arial Black, sans-serif" font-size="11.5" fill="#FFFFFF">STAGE VI · EMPIRICAL MOLD POPULATION SIEVE (21 ARCHETYPES -&gt; 13+1 MOLDS -&gt; TOP 5 SELECTED)</text>
  <text x="615" y="371" font-family="Consolas, monospace" font-size="10" font-weight="bold" fill="#F3E8FF">THRESHOLD: &gt;= 3 CASES</text>

  <text x="28" y="396" font-family="Consolas, monospace" font-size="9.0" fill="#E2E8F0">Research tags all 35 discovered practitioner cases across the 14 canonical mold codes without pre-selection. Only molds</text>
  <text x="28" y="409" font-family="Consolas, monospace" font-size="9.0" fill="#E2E8F0">reaching &gt;=3 corroborated cases qualify; Stage 02 selects the 5 most densely populated molds for Chapter 4 Caricatures + Exhibits:</text>

  <!-- Bar chart of all 14 molds -->
  <!-- Selected Top 5 (Amber/Gold bars) -->
  <rect x="28" y="420" width="380" height="158" rx="6" fill="#1E293B" stroke="#F59E0B" stroke-width="1.5"/>
  <text x="38" y="438" font-family="Arial Black, sans-serif" font-size="9.5" fill="#FBBF24">TOP 5 POPULATED MOLDS (&gt;=3 CASES -&gt; CH 4 SPREADS)</text>

  <text x="38" y="458" font-family="Consolas, monospace" font-size="8.8" fill="#FFFFFF">H4 Upsell Ladder (Elena)</text>
  <rect x="195" y="447" width="144" height="13" rx="3" fill="#D97706"/>
  <text x="345" y="458" font-family="Consolas, monospace" font-size="8.5" fill="#FDE68A">6 raw / 3 locked</text>

  <text x="38" y="480" font-family="Consolas, monospace" font-size="8.8" fill="#FFFFFF">E2 Built-In Audience (Marcus)</text>
  <rect x="195" y="469" width="120" height="13" rx="3" fill="#F59E0B"/>
  <text x="321" y="480" font-family="Consolas, monospace" font-size="8.5" fill="#FDE68A">5 raw / 3 locked</text>

  <text x="38" y="502" font-family="Consolas, monospace" font-size="8.8" fill="#FFFFFF">H5 The Tool Tax (Devon)</text>
  <rect x="195" y="491" width="120" height="13" rx="3" fill="#D97706"/>
  <text x="321" y="502" font-family="Consolas, monospace" font-size="8.5" fill="#FDE68A">5 raw / 3 locked</text>

  <text x="38" y="524" font-family="Consolas, monospace" font-size="8.8" fill="#FFFFFF">C3 The Rule Change (Victor)</text>
  <rect x="195" y="513" width="120" height="13" rx="3" fill="#B45309"/>
  <text x="321" y="524" font-family="Consolas, monospace" font-size="8.5" fill="#FDE68A">5 raw / 3 locked</text>

  <text x="38" y="546" font-family="Consolas, monospace" font-size="8.8" fill="#FFFFFF">C1 The Runway Gap (Hannah)</text>
  <rect x="195" y="535" width="96" height="13" rx="3" fill="#B45309"/>
  <text x="297" y="546" font-family="Consolas, monospace" font-size="8.5" fill="#FDE68A">4 raw / 3 locked</text>

  <text x="38" y="568" font-family="Consolas, monospace" font-size="8.2" fill="#86EFAC">Each gets: 1 Dramatic Caricature (Verso) + 1 Data Exhibit (Recto)</text>

  <!-- Non-selected 9 molds (Slate bars below threshold) -->
  <rect x="420" y="420" width="372" height="158" rx="6" fill="#1E293B" stroke="#64748B" stroke-width="1.2"/>
  <text x="430" y="438" font-family="Arial Black, sans-serif" font-size="9.5" fill="#94A3B8">REMAINING 9 MOLD CODES (BELOW &gt;=3 THRESHOLD)</text>
  <text x="430" y="457" font-family="Consolas, monospace" font-size="8.5" fill="#CBD5E1">• E4 Full-Time Grind (2 raw)   • H3 Hidden Hours (2 raw)</text>
  <text x="430" y="473" font-family="Consolas, monospace" font-size="8.5" fill="#CBD5E1">• C4 Concentration Trap (2 raw)• E1 Second Attempt (1 raw)</text>
  <text x="430" y="489" font-family="Consolas, monospace" font-size="8.5" fill="#CBD5E1">• E3 Early Mover (1 raw)       • H1 Capital Floor (1 raw)</text>
  <text x="430" y="505" font-family="Consolas, monospace" font-size="8.5" fill="#CBD5E1">• C2 Margin Squeeze (1 raw)    • H2 The Ad Tax (0 raw in FX)</text>
  <text x="430" y="521" font-family="Consolas, monospace" font-size="8.5" fill="#CBD5E1">• new? Unclassified Anomaly (0 raw — all 35 fit canon menu)</text>
  <rect x="430" y="532" width="352" height="36" rx="4" fill="#020617"/>
  <text x="438" y="547" font-family="Consolas, monospace" font-size="8.2" fill="#38BDF8">ORTHOGONAL INDEPENDENCE: A young topic with sparse data</text>
  <text x="438" y="560" font-family="Consolas, monospace" font-size="8.2" fill="#38BDF8">may still yield 5+ strong Pivots; we use what clears the bar.</text>

  <!-- Down Arrow 6 -> 7 -->
  <polygon points="410,594 400,604 420,604" fill="#10B981"/>
  <line x1="410" y1="590" x2="410" y2="608" stroke="#10B981" stroke-width="3"/>

  <!-- STAGE 7: BENEFICIAL MECHANISMS (BM-01..08), ETHICAL AFFILIATES & DETERMINISTIC COMPILER GATES G1-G7 -->
  <rect x="16" y="608" width="788" height="432" rx="8" fill="#0F172A" stroke="#10B981" stroke-width="2"/>
  <rect x="16" y="608" width="788" height="26" rx="6" fill="url(#accentGreen)"/>
  <text x="28" y="625" font-family="Arial Black, sans-serif" font-size="11.5" fill="#020617">STAGE VII · BENEFICIAL MECHANISM EXTRACTION, MONETIZATION &amp; DETERMINISTIC COMPILER GATES (G1–G7)</text>
  <text x="650" y="625" font-family="Consolas, monospace" font-size="10" font-weight="bold" fill="#020617">ALL 7 GATES EXIT 0</text>

  <!-- Top split: Beneficial Mechanisms (BM-01..08) + 8 Monetizable Angles (MON-01..08) -->
  <rect x="28" y="644" width="374" height="174" rx="6" fill="#1E293B" stroke="#10B981" stroke-width="1.4"/>
  <text x="38" y="662" font-family="Arial Black, sans-serif" font-size="9.5" fill="#4ADE80">1. Beneficial Mechanisms (BM-01..08 -&gt; $1.99 Pamphlet)</text>
  <text x="38" y="678" font-family="Consolas, monospace" font-size="8.4" fill="#E2E8F0">While exposing the 97.6% failure rate, Research extracts</text>
  <text x="38" y="690" font-family="Consolas, monospace" font-size="8.4" fill="#E2E8F0">the 8 practices that genuinely protect &amp; pay the reader:</text>
  <text x="38" y="706" font-family="Consolas, monospace" font-size="8.3" fill="#A7F3D0">• BM-01: 0.20%-0.25% Position Sizing (&lt;0.5% ruin odds)</text>
  <text x="38" y="719" font-family="Consolas, monospace" font-size="8.3" fill="#A7F3D0">• BM-02: 100-Trade / 90-Day Walk-Forward Paper Protocol</text>
  <text x="38" y="732" font-family="Consolas, monospace" font-size="8.3" fill="#A7F3D0">• BM-03: Regulated Micro-Cash Routing (Zero Reset Tax)</text>
  <text x="38" y="745" font-family="Consolas, monospace" font-size="8.3" fill="#A7F3D0">• BM-04: CPI/FOMC/NFP 15-Min Spread Blackout Filter</text>
  <text x="38" y="758" font-family="Consolas, monospace" font-size="8.3" fill="#A7F3D0">• BM-05..07: $500 Cap + Zero-Stack Audit + 120d Chargeback</text>
  <text x="38" y="771" font-family="Consolas, monospace" font-size="8.3" fill="#A7F3D0">• BM-08: SQL/Python Skill Translation ($38.90–$55.00/hr)</text>
  <rect x="38" y="780" width="354" height="30" rx="4" fill="#020617"/>
  <text x="44" y="799" font-family="Consolas, monospace" font-size="8.3" fill="#34D399">-&gt; Assembled into Ch 7/8 &amp; $1.99 Companion Pamphlet (P1–P7)</text>

  <rect x="414" y="644" width="378" height="174" rx="6" fill="#1E293B" stroke="#38BDF8" stroke-width="1.4"/>
  <text x="424" y="662" font-family="Arial Black, sans-serif" font-size="9.5" fill="#38BDF8">2. Ethical Affiliate &amp; Post-Sale Ecosystem (MON-01..08)</text>
  <text x="424" y="678" font-family="Consolas, monospace" font-size="8.4" fill="#E2E8F0">Every research pass simultaneously populates 8 ethical</text>
  <text x="424" y="690" font-family="Consolas, monospace" font-size="8.4" fill="#E2E8F0">downstream products &amp; vetted partner slots:</text>
  <text x="424" y="706" font-family="Consolas, monospace" font-size="8.3" fill="#BAE6FD">• MON-01: Cleared Skill-Pivot Affiliates (SQL/Python/Series 65)</text>
  <text x="424" y="719" font-family="Consolas, monospace" font-size="8.3" fill="#BAE6FD">• MON-02: $1.99 Post-Sale Success &amp; Exit Pamphlet PDF</text>
  <text x="424" y="732" font-family="Consolas, monospace" font-size="8.3" fill="#BAE6FD">• MON-03: Live 'Run the 4 Gates' Web Calculator &amp; Chatbot</text>
  <text x="424" y="745" font-family="Consolas, monospace" font-size="8.3" fill="#BAE6FD">• MON-04: De-Guru'd Spreadsheet &amp; Backtest Template Vault</text>
  <text x="424" y="758" font-family="Consolas, monospace" font-size="8.3" fill="#BAE6FD">• MON-05..07: 14p Freebie PDF + Podcap + 3-Book Bundles</text>
  <text x="424" y="771" font-family="Consolas, monospace" font-size="8.3" fill="#BAE6FD">• MON-08: $49.99/mo Y-IT Pro 54-Topic Lakehouse Access</text>
  <rect x="424" y="780" width="358" height="30" rx="4" fill="#020617"/>
  <text x="430" y="799" font-family="Consolas, monospace" font-size="8.3" fill="#7DD3FC">-&gt; Disclosed plainly on *Results Typical Page (ADR 0050)</text>

  <!-- Bottom Full-Width Box: The 7 Deterministic Compiler Verification Gates (G1-G7) -->
  <rect x="28" y="828" width="764" height="198" rx="6" fill="#020617" stroke="#10B981" stroke-width="1.8"/>
  <text x="40" y="848" font-family="Arial Black, sans-serif" font-size="10.5" fill="#34D399">3. DETERMINISTIC COMPILER VERIFICATION GATES (yit-factory.exe verify · internal/gate/gates.go)</text>
  <text x="40" y="864" font-family="Consolas, monospace" font-size="8.6" fill="#94A3B8">No human or AI agent is allowed to self-declare "PASS." Compiled code parses every YAML, Markdown, SVG, and PDF byte:</text>

  <text x="40" y="884" font-family="Consolas, monospace" font-size="8.8" fill="#4ADE80">[EXIT 0] GATE G1 (Trigger &amp; Questions):</text>
  <text x="275" y="884" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">All 38 canonical research questions verified across P1, P2, and P3.</text>

  <text x="40" y="902" font-family="Consolas, monospace" font-size="8.8" fill="#4ADE80">[EXIT 0] GATE G2 (Verbatim Trace):</text>
  <text x="275" y="902" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">53/53 rows &amp; numbers verified verbatim in pages/*.html + SHA-256 lock.</text>

  <text x="40" y="920" font-family="Consolas, monospace" font-size="8.8" fill="#4ADE80">[EXIT 0] GATE G3 (Tiers &amp; Deep Layers):</text>
  <text x="275" y="920" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">Tiers 1–4 sources, RM-10v2 convergence matrix, 3-layer affiliates &amp; BM-01..08 verified.</text>

  <text x="40" y="938" font-family="Consolas, monospace" font-size="8.8" fill="#4ADE80">[EXIT 0] GATE G4 (14-Mold Sieve &amp; Ch 4):</text>
  <text x="275" y="938" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">Zero PosiBot in Ch 4; vetted molds (&gt;=3 cases each) with Caricature + Technical Exhibit.</text>

  <text x="40" y="956" font-family="Consolas, monospace" font-size="8.8" fill="#4ADE80">[EXIT 0] GATE G5 (Phase Isolation):</text>
  <text x="275" y="956" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">Phase 1 cites Claims only; Phase 3 pivots &amp; Ch 7-&gt;Ch 8 seam verified.</text>

  <text x="40" y="974" font-family="Consolas, monospace" font-size="8.8" fill="#4ADE80">[EXIT 0] GATE G6 (Tone &amp; Vocabulary):</text>
  <text x="275" y="974" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">Zero person-blame or deprecated tokens; Ch 2 cliffhanger &amp; callout grammar verified.</text>

  <text x="40" y="992" font-family="Consolas, monospace" font-size="8.8" fill="#4ADE80">[EXIT 0] GATE G7 (Natural Range Print):</text>
  <text x="275" y="992" font-family="Consolas, monospace" font-size="8.8" fill="#E2E8F0">Static front/back bookends + unforced natural body/bibliography range (40–84 pages).</text>

  <rect x="40" y="1001" width="740" height="19" rx="3" fill="#064E3B"/>
  <text x="48" y="1014" font-family="Consolas, monospace" font-size="8.5" font-weight="bold" fill="#A7F3D0">LIVE RUN TELEMETRY: 36 Live Probes · 140+ Validated Sources · 53 Locked Rows · 38/38 Questions · 100% Verbatim Match</text>
</svg>`
}
