package render

import (
	"fmt"
	"path/filepath"
)

func buildPhase2Pages19To44(rootDir string) []KDPPage {
	p2Assets := filepath.Join(rootDir, "04_phase2_run", "assets")
	ch3Waterfall := fileURI(filepath.Join(p2Assets, "T01-P2-C03-chart-stake-waterfall_v001.svg"))
	ch3Split := fileURI(filepath.Join(p2Assets, "T01-P2-C03-chart-bbook-split-bar_v001.svg"))
	ch3Hourly := fileURI(filepath.Join(p2Assets, "T01-P2-C03-chart-payoff-hourly-meter_v001.svg"))
	ch3PosiBot := fileURI(filepath.Join(p2Assets, "T01-P2-C03-posibot-reset-discount-code_v001.svg"))

	ch4Car1 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-caricature-case1-e2-marcus_v001.svg"))
	ch4Car2 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-caricature-case2-h4-elena_v001.svg"))
	ch4Car3 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-caricature-case3-h5-devon_v001.svg"))
	ch4Car4 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-caricature-case4-c1-hannah_v001.svg"))
	ch4Car5 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-caricature-case5-c3-victor_v001.svg"))

	ch4Ex1 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-exhibit-case1-e2-gap-bar_v001.svg"))
	ch4Ex2 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-exhibit-case2-h4-split-bar_v001.svg"))
	ch4Ex3 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-exhibit-case3-h5-waterfall_v001.svg"))
	ch4Ex4 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-exhibit-case4-c1-timeline_v001.svg"))
	ch4Ex5 := fileURI(filepath.Join(p2Assets, "T01-P2-C04-exhibit-case5-c3-cliff_v001.svg"))

	ch5Ladder := fileURI(filepath.Join(p2Assets, "T01-P2-C05-chart-fx-market-log-ladder_v001.svg"))
	ch5Funnel := fileURI(filepath.Join(p2Assets, "T01-P2-C05-chart-prop-attrition-funnel_v001.svg"))
	ch5PosiBot := fileURI(filepath.Join(p2Assets, "T01-P2-C05-posibot-gladiator-liquidity_v001.svg"))

	ch6Scorecard := fileURI(filepath.Join(p2Assets, "T01-P2-C06-chart-four-metric-scorecard_v001.svg"))
	ch6PosiBot := fileURI(filepath.Join(p2Assets, "T01-P2-C06-posibot-verdict-receipt-tape_v001.svg"))

	return []KDPPage{
		// Page 19: Ch 3 Opener (Recto)
		{
			PageNum:     19,
			ThemeClass:  "theme-phase2",
			RunningHead: "PHASE 2: DEMOLITION · CHAPTER 3: THE MONEY PIT",
			FooterNote:  "VOICE: SARDONIC AUTHORITY · FORESHADOW PAYOFF [F-001, F-004]",
			BodyHTML: `
<div class="phase-badge" style="background:#DC2626;">PHASE 2 · DEMOLITION // CHAPTER 03 · RECTO OPENER</div>
<h1>CHAPTER 3: THE MONEY PIT</h1>
<p style="font-family:'Consolas',monospace; font-size:8.4pt; color:#991B1B; font-weight:700;">GATE METRICS INTRODUCED: THE STAKE ($1,840) &amp; THE PAYOFF ($6.40/HR)</p>
<h3>The Foreshadow Pays Off: The Dashboard Reads $0.00</h3>
<p>In Chapter 1, it was 11:42 PM and you were softly caressing a <code>$2,400</code> gold trade from your sofa in flannel pajama pants (<code>[C-005]</code>). Let's turn the fluorescent overhead lights on.</p>
<p>It is 9:31 AM on Day Eleven of your <code>$100,000</code> evaluation. A U.S. CPI inflation print just hit the wire, the bid-ask spread on XAU/USD yawned open like an alligator's jaw, and your screen is flashing a banner the color of an arterial bleed:</p>
<div class="callout-rc">
  <strong>🚨 ACCOUNT #FT-100K BREACHED — INTRADAY EQUITY DRAWDOWN EXCEEDED</strong><br>
  <code>ELIGIBLE PAYOUT BALANCE: $0.00 USD · ACCOUNT STATUS: TERMINATED</code><br>
  <em>Automated Inbox Notification (38 seconds later): "Don't give up on your dream! Use code RELOAD15 for 15% off a fresh $340 Evaluation Reset!"</em> (<code>[F-001]</code>)
</div>`,
		},
		// Page 20: Ch 3 The Stake Waterfall (Verso)
		{
			PageNum:     20,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 3: THE MONEY PIT · THE STAKE ($1,400–$2,200 RANGE / $1,840 MEDIAN)",
			FooterNote:  "AUDITED SOURCES: [F-001, F-002, F-004, F-014, C-003]",
			BodyHTML: fmt.Sprintf(`
<h2>1. TESTING THE "$89–$497 ENTRY FEE" CLAIM (THE STAKE = $1,400–$2,200 RANGE · $1,840 MEDIAN)</h2>
<p>The checkout page sold you a one-time refundable audition ticket (<code>[C-003]</code>). In the wild, nobody buys just one prop challenge for the exact same reason casino architecture doesn't put a clock above the craps table. Across multiple independent cohorts and enforcement filings, total Year-1 starter outlay consistently lands in the <strong><code>$1,400–$2,200</code> range</strong> (<code>[F-001, F-002, K-004..K-015]</code>). In the <strong><code>420,000</code>-account prop evaluation benchmark</strong>, the median participant purchases <strong><code>4.2</code> challenges or resets</strong> at a median fee of <strong><code>$340</code></strong> per attempt (<code>$1,428</code> in direct challenge tuition), plus <strong><code>$412</code></strong> in TradingView upgrades, low-latency London VPS rentals, and indicator plugins—putting the audited cohort median <strong>Stake at <code>$1,840</code></strong> before a trader either sees a net-profitable month or walks away (<code>[F-001]</code>, <code>[F-004]</code>).</p>
<div class="asset-frame"><img src="%s" alt="Stake Waterfall Chart"></div>
<p>The Federal Reserve reports that the typical U.S. household's entire liquid emergency cash buffer is <strong><code>$1,000</code></strong> (<code>[F-014]</code>). The "zero personal capital at risk" pitch quietly extracts <strong>140%% to 220%% (<code>184%%</code> at the median) of a typical family's emergency fund</strong> in non-refundable fees.</p>
`, ch3Waterfall),
		},
		// Page 21: Ch 3 The B-Book Arcade (Recto)
		{
			PageNum:     21,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 3: THE MONEY PIT · THE 98.4% B-BOOK ARCADE",
			FooterNote:  "CFTC ENFORCEMENT ACCOUNTING [F-002, F-010, F-011]",
			BodyHTML: fmt.Sprintf(`
<h2>2. WELCOME TO THE B-BOOK ARCADE: THERE IS NO $200,000 ACCOUNT</h2>
<p>Here is the structural truth that turns every "Institutional Funded Trader" bio on Instagram into performance art: <strong>there is no <code>$200,000</code> institutional account</strong> (<code>[C-001]</code>, <code>[F-011]</code>).</p>
<p>When the CFTC unsealed its enforcement complaint against one of the largest retail prop evaluation operators on earth (<code>[F-002]</code>), investigators documented that the firm collected <strong><code>$310 million</code> in assessment and reset fees from more than <code>135,000</code> retail customers</strong> in 22 months—while routing <strong><code>98.4%%</code> of all customer trades to an internal simulated B-book server</strong> (<code>[F-002]</code>, <code>[F-011]</code>). You are not trading in the <code>$7.5 trillion</code> interbank market. You are playing a video game inside the operator's private SQLite database, where the house pays the tiny sliver of winners out of the reset fees of the <code>93.4%%</code> who just blew up (<code>[F-007]</code>, <code>[F-011]</code>).</p>
<div class="asset-frame"><img src="%s" alt="B-Book Split Bar Chart"></div>
`, ch3Split),
		},
		// Page 22: Ch 3 Reality Check & PosiBot (Verso)
		{
			PageNum:     22,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 3: THE MONEY PIT · B-BOOK INCENTIVES",
			FooterNote:  "AUDITED SOURCES: [F-001, F-010, F-011]",
			BodyHTML: fmt.Sprintf(`
<div class="callout-rc">
  <strong>[REALITY CHECK — THE B-BOOK COUNTERPARTY ENGINE]:</strong><br>
  In a regulated live brokerage, the broker earns commissions when you survive and keep trading volume. In a B-book prop evaluation (<code>98.4%%</code> simulated execution, <code>[F-011]</code>), a trader who passes and withdraws money is a direct balance-sheet expense. Forensic accounting showed that asymmetric slippage scripts, commission markups, and daily drawdown tripwires terminated <strong><code>91.6%%</code> of funded accounts before a second payout</strong> could be completed (<code>[F-010]</code>).
</div>
<table class="yit-table">
  <tr><th>Expense Category</th><th>Brochure Pitch (<code>[C-003, C-009]</code>)</th><th>Audited 6-Month Reality (<code>[F-001, K-004..K-009]</code>)</th></tr>
  <tr><td><strong>Evaluation &amp; Reset Fees</strong></td><td>$89–$497 (100%% refunded!)</td><td><strong>$1,428 median</strong> (4.2 attempts × $340; 93.4%% unrefunded)</td></tr>
  <tr><td><strong>Charting, VPS &amp; Indicators</strong></td><td>"Included template"</td><td><strong>$412 median</strong> (up to $1,703 in Mold H5 Tool Tax)</td></tr>
  <tr><td><strong>Courses &amp; VIP Signal Rooms</strong></td><td>"Free YouTube blueprint"</td><td><strong>$0 to $2,435</strong> ($997 academy + $149/mo Discord signals)</td></tr>
</table>
<div class="asset-frame"><img src="%s" alt="PosiBot Whiteboard Fifth Reset"></div>
`, ch3PosiBot),
		},
		// Page 23: Ch 3 The Payoff Hourly Meter (Recto)
		{
			PageNum:     23,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 3: THE MONEY PIT · THE PAYOFF ($4.50–$8.50/HR RANGE · $6.40/HR MEDIAN)",
			FooterNote:  "AUDITED SOURCES: [F-003, F-005, F-008, F-013]",
			BodyHTML: fmt.Sprintf(`
<h2>3. TESTING THE "$10,000/MO IN 45 MINUTES" CLAIM (THE PAYOFF = $4.50–$8.50/HR · $6.40/HR MEDIAN)</h2>
<p>What if you actually beat the <code>96.5%%–98.5%%</code> (<code>97.6%%</code> cohort median) first-year attrition wall (<code>[F-008]</code>)? Across surviving retail FX and prop cohorts (the top <strong><code>1.5%%–3.5%%</code></strong>, with a <strong><code>2.4%%</code></strong> cohort midpoint), logged screen, backtesting, and session-waiting time spans <strong><code>20 to 30 hours per week</code></strong> (<code>26.5 hrs/wk</code> median)—yielding an effective net <strong>Payoff in the <code>$4.50 to $8.50 per hour</code> range, with an audited cohort median of <code>$6.40 per hour</code></strong> (<code>[F-003, F-005]</code>).</p>
<div class="asset-frame"><img src="%s" alt="Hourly Meter Chart"></div>
<p>That <code>$4.50–$8.50/hr</code> (<code>$6.40/hr</code> median) survivor wage is less than a quarter of the <strong><code>$29.76/hr</code></strong> U.S. median hourly wage (<code>[F-013]</code>), and corroborates multi-year regulatory cohorts showing <strong><code>94.2%%</code> of leveraged retail FX traders earn below prevailing minimum wage</strong> (<code>[F-005]</code>).</p>
`, ch3Hourly),
		},
		// Page 24: Ch 3 Global Regulatory Benchmarks & Verdict Band (Verso)
		{
			PageNum:     24,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 3: THE MONEY PIT · GLOBAL REGULATORY LEDGER",
			FooterNote:  "AUDITED SOURCES: [F-005, F-012]",
			BodyHTML: `
<h2>4. THE GLOBAL REGULATORY CONSENSUS: CFTC, ESMA &amp; FCA</h2>
<p>Think the prop-challenge numbers are an anomaly? Look at regulated cash forex accounts where traders aren't even fighting a 5% trailing drawdown tripwire:</p>
<table class="yit-table">
  <tr><th>Regulatory Authority</th><th>Mandated Cohort Finding</th><th>Citation</th></tr>
  <tr><td><strong>U.S. CFTC (RFED Filings)</strong></td><td><strong>72%</strong> of active U.S. retail forex accounts unprofitable across trailing 4 quarters; median deposit exhausted in <strong>4.5 months</strong>.</td><td><code>[F-012]</code></td></tr>
  <tr><td><strong>ESMA &amp; UK FCA Disclosures</strong></td><td>Between <strong>74% and 89%</strong> of retail investor accounts lose money trading CFDs and spot FX.</td><td><code>[F-005]</code></td></tr>
  <tr><td><strong>Longitudinal Day-Trader Study</strong></td><td><strong>94.2%</strong> of leveraged retail FX participants earn less than prevailing minimum wage after spread and swap drag.</td><td><code>[F-005]</code></td></tr>
</table>
<div class="verdict-band">
  <strong>CHAPTER 3 VERDICT BAND:</strong> The brochure promised <code>$10,000/month</code> for <code>$497</code> down (<code>[C-002, C-003]</code>). The ledger shows a <strong><code>$1,840</code> median Stake</strong> (<code>[F-004]</code>), a <strong><code>98.4%</code> simulated B-book server</strong> (<code>[F-011]</code>), and a <strong><code>$6.40/hr</code> survivor Payoff</strong> (<code>[F-003]</code>).
</div>`,
		},
		// Page 25: Chapter 4 Opener — Aged Amber Case-File Insert (Recto)
		{
			PageNum:     25,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CHAPTER 4: TALES FROM THE TRENCHES · FORENSIC INSERT",
			FooterNote:  "AGED AMBER DOSSIER STOCK (#EFE5CE) · ZERO POSIBOT · ZERO JOKES",
			BodyHTML: `
<div class="stamp-badge">FORENSIC AUTOPSY INSERT · STANDALONE CORONER'S LEDGER</div>
<h1 style="color:#451A03; margin-top:8px;">CHAPTER 4: TALES FROM THE TRENCHES</h1>
<p><strong>PROTOCOL NOTICE (<code>outline-spine.v1.yaml §5.2</code>):</strong> PosiBot does not appear in this chapter. There are no punchlines in this chapter. The numbers are the joke.</p>
<div class="dossier-box">
  <strong>SELECTION &amp; COMPOSITE RULE:</strong> Every persona in this insert is a composite synthesized from <strong>at all times at least three documented real-world cases</strong> (<code>[K-001]</code> through <code>[K-015]</code>), grouped by the structural mechanism that decided their outcome. All financial figures printed below are the <strong>exact medians</strong> of the underlying case rows.
</div>
<table class="yit-table">
  <tr><th>File #</th><th>Mold Code &amp; Name</th><th>Composite Persona</th><th>Cases</th><th>Outcome</th></tr>
  <tr><td><code>#01</code></td><td><code>E2</code> · The Built-In Audience</td><td><strong>Marcus Conviction™</strong></td><td><code>K-001..K-003</code> (3)</td><td><strong>WON</strong> (Affiliate Edge)</td></tr>
  <tr><td><code>#02</code></td><td><code>H4</code> · The Upsell Ladder</td><td><strong>Elena Diligence</strong></td><td><code>K-004..K-006</code> (3)</td><td><strong>EXITED</strong> (-$4,110)</td></tr>
  <tr><td><code>#03</code></td><td><code>H5</code> · The Tool Tax</td><td><strong>Devon Precision</strong></td><td><code>K-007..K-009</code> (3)</td><td><strong>EXITED</strong> (-$2,923)</td></tr>
  <tr><td><code>#04</code></td><td><code>C1</code> · The Runway Gap</td><td><strong>Hannah Patience</strong></td><td><code>K-010..K-012</code> (3)</td><td><strong>EXITED</strong> ($1.11/hr)</td></tr>
  <tr><td><code>#05</code></td><td><code>C3</code> · The Rule Change</td><td><strong>Victor Discipline</strong></td><td><code>K-013..K-015</code> (3)</td><td><strong>EXITED</strong> ($5.1k Frozen)</td></tr>
</table>`,
		},
		// Page 26: Case File #01 Dossier (Verso)
		{
			PageNum:     26,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #01 · MOLD E2: THE BUILT-IN AUDIENCE",
			FooterNote:  "COMPOSITE OF 3 DOCUMENTED CASES [K-001, K-002, K-003]",
			BodyHTML: fmt.Sprintf(`
<div class="stamp-badge">CASE FILE #01 · MOLD E2: THE BUILT-IN AUDIENCE</div>
<h2>PERSONA: "MARCUS CONVICTION™"</h2>
<div class="asset-frame"><img src="%s" alt="Marcus Conviction Caricature"></div>
<div class="dossier-box">
  • <strong>Life Position:</strong> Short-form lifestyle &amp; finance creator (52,000 to 145,000 existing subscribers across YouTube/TikTok/Newsletter, <code>[K-001..K-003]</code>).<br>
  • <strong>Sold By Claim:</strong> <code>[C-001]</code> (<em>"Pass our 2-Step Challenge in 14 days and trade up to a $200,000 funded account — keep 90%% of the profits!"</em>).<br>
  • <strong>What It Took (Range &amp; Medians):</strong> <code>$540–$900</code> (<code>$600</code> median) in challenge fees over <code>7 months</code>; <strong><code>-$1,900 to +$1,400</code> (<code>$0</code> median) net trading profit</strong>; <strong><code>$41,200–$68,400</code> (<code>$54,000</code> median) in affiliate CPA commissions</strong> from <code>88,000</code> followers (<code>[K-001..K-003]</code>, <code>[F-024]</code>).
</div>
<p><strong>1. Month 1:</strong> Purchased a <code>$100,000</code> evaluation (<code>$540–$600</code>) and filmed "Day in the Life of a Funded Trader" vlogs with a <code>20%%</code> discount code paying <code>$108</code> CPA per viewer checkout (<code>[K-001, F-024]</code>). <strong>2. Month 3:</strong> Breached two live evaluations while viewer referrals paid <code>$6,000–$9,500/mo</code>. <strong>3. Month 7:</strong> Collected a median of <strong><code>$54,000</code> in affiliate commissions</strong> while cumulative net trading profit sat at <strong><code>$0</code></strong> (<code>[K-001..K-003]</code>).</p>`, ch4Car1),
		},
		// Page 27: Case File #01 Exhibit (Recto)
		{
			PageNum:     27,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #01 EXHIBIT · MOLD E2 GAP BAR",
			FooterNote:  "STAMP: COMPOSITE OF 3 DOCUMENTED CASES [K-001..K-003]",
			BodyHTML: fmt.Sprintf(`
<h3>FORENSIC EXHIBIT #01: TRADING PAYOUT VS. AFFILIATE REFERRAL REVENUE</h3>
<div class="asset-frame"><img src="%s" alt="Case 1 E2 Gap Bar"></div>
<div class="dossier-box">
  <strong>SYSTEM MECHANISM (MOLD E2):</strong> An existing following meant never paying for customers—and never needing to win a trade. The referral link paid 20%%–25%% on every viewer's evaluation purchase regardless of whether the creator or the viewer ever passed Phase 2.<br><br>
  <strong>OUTCOME:</strong> <span class="stamp-badge">WON (VIA AFFILIATE DISTRIBUTION, NOT TRADING)</span><br>
  <strong>STAMP:</strong> <em>Composite of 3 documented cases (<code>K-001, K-002, K-003</code>).</em>
</div>
`, ch4Ex1),
		},
		// Page 28: Case File #02 Dossier (Verso)
		{
			PageNum:     28,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #02 · MOLD H4: THE UPSELL LADDER",
			FooterNote:  "COMPOSITE OF 3 DOCUMENTED CASES [K-004, K-005, K-006]",
			BodyHTML: fmt.Sprintf(`
<div class="stamp-badge">CASE FILE #02 · MOLD H4: THE UPSELL LADDER</div>
<h2>PERSONA: "ELENA DILIGENCE"</h2>
<div class="asset-frame"><img src="%s" alt="Elena Diligence Caricature"></div>
<div class="dossier-box">
  • <strong>Life Position:</strong> Full-time shift supervisor / healthcare / logistics worker (<code>[K-004..K-006]</code>).<br>
  • <strong>Sold By Claim:</strong> <code>[C-002]</code> (<em>"45 minutes a day during the London-NY overlap using our institutional Smart Money Concepts blueprint."</em>).<br>
  • <strong>What It Took (Range &amp; Medians):</strong> <strong><code>$3,680–$4,650</code> (<code>$4,110</code> median) cash spent</strong>, <strong><code>9 months</code></strong>, <strong><code>620 hours</code> logged</strong>, <strong><code>$0.00</code> withdrawn</strong> (<code>[K-004..K-006]</code>).
</div>
<p><strong>1. Month 1:</strong> Purchased a <code>$49–$297</code> intro workshop; when demo trades failed during choppy London opens, the curriculum diagnosed the loss as "missing advanced institutional order-flow mentorship." <strong>2. Month 3:</strong> Upgraded to a <code>$997–$1,497</code> private academy plus a <code>$149/mo</code> VIP Discord signal room (<code>$1,341</code> median spend, <code>[K-004]</code>). <strong>3. Months 5–9:</strong> Bought <code>5</code> challenges/resets (<code>$1,675</code> median) mirroring delayed Discord alerts; exited at Month 9 down <strong><code>$4,110</code></strong> with <strong><code>$0</code></strong> withdrawn (<code>[K-004..K-006]</code>).</p>`, ch4Car2),
		},
		// Page 29: Case File #02 Exhibit (Recto)
		{
			PageNum:     29,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #02 EXHIBIT · MOLD H4 SPLIT BAR",
			FooterNote:  "STAMP: COMPOSITE OF 3 DOCUMENTED CASES [K-004..K-006]",
			BodyHTML: fmt.Sprintf(`
<h3>FORENSIC EXHIBIT #02: THE $4,110 CURRICULUM &amp; SIGNAL ESCALATOR</h3>
<div class="asset-frame"><img src="%s" alt="Case 2 H4 Split Bar"></div>
<div class="dossier-box">
  <strong>SYSTEM MECHANISM (MOLD H4):</strong> Each course sells the next one; when a student loses a challenge, the curriculum frames the loss as a missing tier of paid secret knowledge rather than a negative-expectancy B-book rule.<br><br>
  <strong>OUTCOME:</strong> <span class="stamp-badge">EXITED ($4,110 SPENT · $0 WITHDRAWN)</span><br>
  <strong>STAMP:</strong> <em>Composite of 3 documented cases (<code>K-004, K-005, K-006</code>).</em>
</div>
`, ch4Ex2),
		},
		// Page 30: Case File #03 Dossier (Verso)
		{
			PageNum:     30,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #03 · MOLD H5: THE TOOL TAX",
			FooterNote:  "COMPOSITE OF 3 DOCUMENTED CASES [K-007, K-008, K-009]",
			BodyHTML: fmt.Sprintf(`
<div class="stamp-badge">CASE FILE #03 · MOLD H5: THE TOOL TAX</div>
<h2>PERSONA: "DEVON PRECISION"</h2>
<div class="asset-frame"><img src="%s" alt="Devon Precision Caricature"></div>
<div class="dossier-box">
  • <strong>Life Position:</strong> Technical operator (mechanical drafter, IT helpdesk tech, warehouse systems lead, <code>[K-007..K-009]</code>).<br>
  • <strong>Sold By Claim:</strong> <code>[C-009]</code> (<em>"Load our proprietary Smart-Money Indicator template on TradingView."</em>).<br>
  • <strong>What It Took (Range &amp; Medians):</strong> <strong><code>$2,590–$3,280</code> (<code>$2,923</code> median) cash spent</strong>, <strong><code>7 months</code></strong>, <strong><code>490 hours</code> logged</strong>, <strong><code>$0.00</code> withdrawn</strong> (<code>[K-007..K-009]</code>).
</div>
<p><strong>1. Month 1:</strong> Activated TradingView Premium (<code>$59/mo</code>), LuxAlgo/footprint overlays (<code>$68/mo</code>), FXReplay backtesting, and a London VPS (<code>$45/mo</code>) before placing a single evaluation trade (<code>[K-007, K-009]</code>). <strong>2. Month 4:</strong> After two manual breaches, paid <code>$499</code> for an "Automated Prop-Pass Gold EA Bot" on Telegram (<code>[K-007, K-008]</code>). <strong>3. Month 7:</strong> Fourth evaluation (<code>$1,220</code> cumulative resets) terminated automatically by the firm's risk server for "prohibited shared EA signature / IP clustering" with <code>$0</code> payout (<code>[K-008]</code>).</p>`, ch4Car3),
		},
		// Page 31: Case File #03 Exhibit (Recto)
		{
			PageNum:     31,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #03 EXHIBIT · MOLD H5 WATERFALL",
			FooterNote:  "STAMP: COMPOSITE OF 3 DOCUMENTED CASES [K-007..K-009]",
			BodyHTML: fmt.Sprintf(`
<h3>FORENSIC EXHIBIT #03: THE $2,923 SUBSCRIPTION &amp; AUTOMATION WATERFALL</h3>
<div class="asset-frame"><img src="%s" alt="Case 3 H5 Waterfall"></div>
<div class="dossier-box">
  <strong>SYSTEM MECHANISM (MOLD H5):</strong> Subscriptions, data feeds, VPS servers, and bot licenses begin billing on Day 1 and never stop—turning a "$340 evaluation" into a <code>$2,923</code> software tax before a single dollar of payout is eligible.<br><br>
  <strong>OUTCOME:</strong> <span class="stamp-badge">EXITED ($2,923 SPENT · $0 WITHDRAWN)</span><br>
  <strong>STAMP:</strong> <em>Composite of 3 documented cases (<code>K-007, K-008, K-009</code>).</em>
</div>
`, ch4Ex3),
		},
		// Page 32: Case File #04 Dossier (Verso)
		{
			PageNum:     32,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #04 · MOLD C1: THE RUNWAY GAP",
			FooterNote:  "COMPOSITE OF 3 DOCUMENTED CASES [K-010, K-011, K-012]",
			BodyHTML: fmt.Sprintf(`
<div class="stamp-badge">CASE FILE #04 · MOLD C1: THE RUNWAY GAP</div>
<h2>PERSONA: "HANNAH PATIENCE"</h2>
<div class="asset-frame"><img src="%s" alt="Hannah Patience Caricature"></div>
<div class="dossier-box">
  • <strong>Life Position:</strong> Part-time bookkeeper / civil engineering tech / quantitative hobbyist (<code>[K-010..K-012]</code>).<br>
  • <strong>Sold By Claim:</strong> <code>[C-008]</code> (<em>"Hit your 8%% Phase 1 and 5%% Phase 2 targets and withdraw 90%% bi-weekly payouts."</em>).<br>
  • <strong>What It Took (Range &amp; Medians):</strong> <strong><code>$1,850–$2,450</code> (<code>$2,150</code> median) in fees</strong>, <strong><code>12–16 months</code> (<code>14 mo</code> median)</strong>, <strong><code>1,120 hours</code></strong>, <strong><code>$1,250</code> net profit = <code>$1.11/hr</code></strong> (<code>[K-010..K-012]</code>).
</div>
<p><strong>1. Months 1–8:</strong> Cycled through 3 to 5 evaluations (<code>$2,150</code> median fees) because responsible <code>0.25%%</code> risk sizing required dozens of winning trades to reach <code>+8%%</code> and <code>+5%%</code>. <strong>2. Months 9–12:</strong> Passed Phase 1 and Phase 2, then waited through minimum trading-day windows and bi-weekly payout holds. <strong>3. Month 14:</strong> Withdrew <code>$3,400</code> gross (<code>$1,250</code> net after <code>$2,150</code> fees) across <code>1,120</code> logged hours—an effective wage of <strong><code>$1.11 per hour</code></strong> (<code>[K-010..K-012]</code>).</p>`, ch4Car4),
		},
		// Page 33: Case File #04 Exhibit (Recto)
		{
			PageNum:     33,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #04 EXHIBIT · MOLD C1 TIMELINE",
			FooterNote:  "STAMP: COMPOSITE OF 3 DOCUMENTED CASES [K-010..K-012]",
			BodyHTML: fmt.Sprintf(`
<h3>FORENSIC EXHIBIT #04: 14 MONTHS &amp; 1,120 HOURS FOR $1.11 AN HOUR</h3>
<div class="asset-frame"><img src="%s" alt="Case 4 C1 Timeline"></div>
<div class="dossier-box">
  <strong>SYSTEM MECHANISM (MOLD C1):</strong> Profit, even when it arrives for a disciplined top-3%% survivor, arrives later than an ordinary household can wait—and once sunk evaluation fees and screen hours are deducted, pays <code>$1.11/hr</code>.<br><br>
  <strong>OUTCOME:</strong> <span class="stamp-badge">EXITED ($1,250 NET OVER 1,120 HOURS = $1.11/HR)</span><br>
  <strong>STAMP:</strong> <em>Composite of 3 documented cases (<code>K-010, K-011, K-012</code>).</em>
</div>
`, ch4Ex4),
		},
		// Page 34: Case File #05 Dossier (Verso)
		{
			PageNum:     34,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #05 · MOLD C3: THE RULE CHANGE",
			FooterNote:  "COMPOSITE OF 3 DOCUMENTED CASES [K-013, K-014, K-015]",
			BodyHTML: fmt.Sprintf(`
<div class="stamp-badge">CASE FILE #05 · MOLD C3: THE RULE CHANGE</div>
<h2>PERSONA: "VICTOR DISCIPLINE"</h2>
<div class="asset-frame"><img src="%s" alt="Victor Discipline Caricature"></div>
<div class="dossier-box">
  • <strong>Life Position:</strong> Experienced swing trader / night-shift machinist / dental hygienist (<code>[K-013..K-015]</code>).<br>
  • <strong>Sold By Claim:</strong> <code>[C-006]</code> (<em>"Currency pairs never ask for refunds, and pay out in pure cash every 14 days."</em>).<br>
  • <strong>What It Took (Range &amp; Medians):</strong> <strong><code>$680–$1,180</code> (<code>$920</code> median) spent</strong>, <strong><code>5 months</code></strong>, <strong><code>380 hours</code></strong>, <strong><code>+$4,200 to +$6,800</code> (<code>+$5,100</code> median) funded profit frozen $\rightarrow$ <code>$0</code> payout</strong> (<code>[K-013..K-015]</code>).
</div>
<p><strong>1. Months 1–4:</strong> Passed Phase 1 and Phase 2 cleanly (<code>$920</code> median spend) and built a <code>$4,200–$6,800</code> funded profit cushion (<code>$5,100</code> median, <code>[K-013..K-015]</code>). <strong>2. Month 5 (Withdrawal Request):</strong> In <code>K-013</code>, the firm retroactively invoked a "30%% single-trade consistency rule"; in <code>K-014</code>, MetaQuotes revoked the offshore broker's white-label MT5 license; in <code>K-015</code>, news-event spread slippage widened by 18 pips during CPI. <strong>3. Month 5 Close:</strong> All three accounts were frozen or breached with <strong><code>$0</code> disbursed from trading profits</strong> (<code>[K-013..K-015]</code>).</p>`, ch4Car5),
		},
		// Page 35: Case File #05 Exhibit (Recto)
		{
			PageNum:     35,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CASE FILE #05 EXHIBIT · MOLD C3 BEFORE/AFTER CLIFF",
			FooterNote:  "STAMP: COMPOSITE OF 3 DOCUMENTED CASES [K-013..K-015]",
			BodyHTML: fmt.Sprintf(`
<h3>FORENSIC EXHIBIT #05: THE OVERNIGHT RULE CHANGE &amp; LICENSE FREEZE</h3>
<div class="asset-frame"><img src="%s" alt="Case 5 C3 Before After Cliff"></div>
<div class="dossier-box">
  <strong>SYSTEM MECHANISM (MOLD C3):</strong> Even when a trader beats the <code>6.6%%</code> evaluation hurdle, an unregulated offshore B-book operator controls the price feed, writes retroactive consistency clauses, and carries <code>0%%</code> statutory SIPC protection (<code>[F-010, F-021]</code>).<br><br>
  <strong>OUTCOME:</strong> <span class="stamp-badge">EXITED ($5,100 FROZEN · $540 RECOVERED VIA CHARGEBACK IN AFTERMATH)</span><br>
  <strong>STAMP:</strong> <em>Composite of 3 documented cases (<code>K-013, K-014, K-015</code>).</em>
</div>
`, ch4Ex5),
		},
		// Page 36: Chapter 4 Cross-Case Synthesis & Verdict Band (Verso)
		{
			PageNum:     36,
			ThemeClass:  "theme-amber-ch4",
			RunningHead: "CHAPTER 4: TALES FROM THE TRENCHES · CORONER'S SUMMARY",
			FooterNote:  "END OF STANDALONE AMBER INSERT · 15 DOCUMENTED CASES [K-001..K-015]",
			BodyHTML: `
<h2>CROSS-CASE SYNTHESIS: FIVE MOLDS, FIFTEEN LEDGERS</h2>
<table class="yit-table">
  <tr><th>Persona &amp; Mold</th><th>Median Cash In</th><th>Median Hours</th><th>Net Trading Payout</th><th>System Mechanism</th></tr>
  <tr><td><strong>Marcus Conviction™ (E2)</strong></td><td><code>$600</code></td><td><code>140 hrs</code></td><td><code>$0</code> (+$54k Affiliate)</td><td>Referral link monetized viewer churn</td></tr>
  <tr><td><strong>Elena Diligence (H4)</strong></td><td><code>$4,110</code></td><td><code>620 hrs</code></td><td><code>$0</code></td><td>Course &amp; VIP signal upsell escalator</td></tr>
  <tr><td><strong>Devon Precision (H5)</strong></td><td><code>$2,923</code></td><td><code>490 hrs</code></td><td><code>$0</code></td><td>Software, VPS, EA bot &amp; reset tax</td></tr>
  <tr><td><strong>Hannah Patience (C1)</strong></td><td><code>$2,150</code></td><td><code>1,120 hrs</code></td><td><code>+$1,250</code> ($1.11/hr)</td><td>14-month multi-stage runway gap</td></tr>
  <tr><td><strong>Victor Discipline (C3)</strong></td><td><code>$920</code></td><td><code>380 hrs</code></td><td><code>$0</code> ($5.1k Frozen)</td><td>Retroactive rule / MT5 license freeze</td></tr>
</table>
<div class="verdict-band">
  <strong>CHAPTER 4 VERDICT BAND:</strong> Across 15 documented practitioner files, the only participant who cleared a living wage did so by selling discount codes to his subscribers (<code>[K-001..K-003]</code>). Everyone who relied on the trading terminal itself either lost <code>$920–$4,110</code> or worked 14 months for <code>$1.11</code> an hour.
</div>`,
		},
		// Page 37: Chapter 5 Opener — Exit Liquidity (Recto)
		{
			PageNum:     37,
			ThemeClass:  "theme-phase2",
			RunningHead: "PHASE 2: DEMOLITION · CHAPTER 5: EXIT LIQUIDITY",
			FooterNote:  "GATE METRICS INTRODUCED: THE ODDS (2.4%) & THE WAIT (14.0 MO) [F-006]",
			BodyHTML: fmt.Sprintf(`
<div class="phase-badge" style="background:#DC2626;">PHASE 2 · DEMOLITION // CHAPTER 05 · RECTO OPENER</div>
<h1>CHAPTER 5: EXIT LIQUIDITY</h1>
<p style="font-family:'Consolas',monospace; font-size:8.4pt; color:#991B1B; font-weight:700;">WHO ACTUALLY OWNS THE $7.5 TRILLION CASINO—AND WHY YOUR STOP-LOSS IS THEIR LUNCH</p>
<h3>1. Paying Off the "$7.5 Trillion Market" Claim (<code>[C-002]</code> $\rightarrow$ <code>[F-006]</code>)</h3>
<p>Every forex webinar on earth opens with the same breathless statistic from the Bank for International Settlements: <em>"The foreign exchange market trades <strong><code>$7.5 trillion per day</code></strong>—you just need a microscopic crumb!"</em> (<code>[C-002]</code>, <code>[F-006]</code>).</p>
<p>Let's read the actual BIS Triennial Central Bank Survey (<code>[F-006]</code>). <strong><code>93%%</code> of that <code>$7.5 trillion</code> daily turnover is executed between institutional reporting dealers, Tier-1 investment banks, quantitative hedge funds, and central banks</strong>, while all retail spot traders worldwide account for <strong>less than <code>2.8%%</code></strong> (<code>[F-006]</code>).</p>
<div class="asset-frame"><img src="%s" alt="FX Market Log Ladder"></div>
`, ch5Ladder),
		},
		// Page 38: Ch 5 The Drawdown Woodchipper (Verso)
		{
			PageNum:     38,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 5: EXIT LIQUIDITY · THE DRAWDOWN WOODCHIPPER",
			FooterNote:  "PAYING OFF MISSING STEP 7 [C-010 -> F-007, F-010]",
			BodyHTML: `
<h2>2. PAYING OFF MISSING STEP 7: THE $5,000 ACCOUNT IN A $100,000 HALLOWEEN COSTUME</h2>
<p>Why was Step 7 missing from Chapter 2's flat-pack assembly brochure (<code>[C-010]</code>)? Because when a prop firm sells you a "$100,000 account with a 10% maximum drawdown and a 5% daily trailing drawdown" (<code>[C-001]</code>, <code>[F-007]</code>), you do not have <code>$100,000</code> of capital. <strong>You have a <code>$5,000</code> account wearing a <code>$100,000</code> Halloween costume</strong> (<code>[F-007]</code>).</p>
<p>If your balance drops by <code>$5,000</code> in a single day, your account is executed on the spot. Asking a trader to make <code>+8%</code> (<code>$8,000</code>) in Phase 1 and <code>+5%</code> (<code>$5,000</code>) in Phase 2 without ever touching a <code>-$5,000</code> daily floor is actually asking them to generate a <strong>+260% return on their true <code>$5,000</code> daily risk allowance</strong> while paying synthetic spread and commission markups that force a <strong><code>64%</code> raw win-rate hurdle</strong> just to break even (<code>[C-010]</code>).</p>
<div class="callout-rc">
  <strong>[REALITY CHECK — THE HIGH-WATER-MARK RATCHET]:</strong><br>
  In many prop evaluations, the 5% trailing drawdown locks to your <em>intraday unrealized high-water mark</em>. If your open trade spikes up +$4,000 during a news candle and pulls back to +$0, you just lost $4,000 of your $5,000 daily allowance on a trade that never even went red (<code>[C-010]</code>, <code>[F-007]</code>).
</div>`,
		},
		// Page 39: Ch 5 The Odds & The Wait Funnel (Recto)
		{
			PageNum:     39,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 5: EXIT LIQUIDITY · THE ODDS (2.4%) & THE WAIT (14 MO)",
			FooterNote:  "COHORT TELEMETRY (420,000 ACCOUNTS) [F-007, F-008, F-009, F-010]",
			BodyHTML: fmt.Sprintf(`
<h2>3. THE ODDS (2.4%%) AND THE WAIT (14.0 MONTHS)</h2>
<p>Let's follow <code>1,000</code> hopeful starters walking into the evaluation funnel with credit cards in hand (<code>[F-007..F-010]</code>):</p>
<p>• Out of <code>1,000</code> starters, <strong><code>934</code> blow up during Phase 1 or Phase 2</strong> (only <strong><code>6.6%%</code></strong> pass both stages, <code>[F-007]</code>).<br>
• Of the <code>66</code> who receive a funded certificate, asymmetric slippage and drawdown tripwires <strong>terminate <code>91.6%%</code> before a second payout</strong> (<code>[F-010]</code>).<br>
• At the 12-month mark, after subtracting all challenge fees, resets, and tool subscriptions, <strong>only <code>24</code> out of <code>1,000</code> (<code>2.4%%</code>) are net-profitable</strong> (<strong>The Odds</strong>, <code>[F-008]</code>)—and their median time to reach consistent net profitability (<strong>The Wait</strong>) is <strong><code>14.0 months</code></strong> (<code>[F-009]</code>).</p>
<div class="asset-frame"><img src="%s" alt="Prop Attrition Funnel"></div>
`, ch5Funnel),
		},
		// Page 40: Ch 5 Ecosystem Food Chain & PosiBot (Verso)
		{
			PageNum:     40,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 5: EXIT LIQUIDITY · THE ECOSYSTEM FOOD CHAIN",
			FooterNote:  "AUDITED SOURCES: [F-002, F-006, F-011, F-024]",
			BodyHTML: fmt.Sprintf(`
<h2>4. THE ECOSYSTEM FOOD CHAIN: WHO EATS WHOM</h2>
<table class="yit-table">
  <tr><th>Tier</th><th>Participant</th><th>Revenue Source</th><th>12-Mo Win Rate</th></tr>
  <tr><td><strong>1. Top Apex</strong></td><td>B-Book Prop Operators (<code>[F-002]</code>)</td><td>$310M+ in challenge &amp; reset fees (98.4%% simulated)</td><td><strong>~100%% House Margin</strong></td></tr>
  <tr><td><strong>2. Barkers</strong></td><td>Affiliate Gurus (Mold E2, <code>[F-024]</code>)</td><td>20%%–25%% CPA ($108/sale) + $149/mo Discord rooms</td><td><strong>$54,000/yr Median</strong></td></tr>
  <tr><td><strong>3. Shovel Sellers</strong></td><td>TradingView / VPS / EA Vendors (<code>[K-007]</code>)</td><td>$45–$68/mo recurring SaaS &amp; $499 bot licenses</td><td><strong>Guaranteed Monthly Billing</strong></td></tr>
  <tr><td><strong>4. Exit Liquidity</strong></td><td>Retail Challenge Buyers (<code>[F-008]</code>)</td><td>Attempting to beat a 5%% trailing drawdown barrier</td><td><strong>2.4%% Net-Profitable</strong></td></tr>
</table>
<div class="asset-frame"><img src="%s" alt="PosiBot Sandwich Board Liquidity"></div>
<div class="verdict-band">
  <strong>CHAPTER 5 VERDICT BAND:</strong> In a <code>$7.5 trillion</code> ocean controlled <code>93%%</code> by institutional sharks (<code>[F-006]</code>), the prop challenge locks you inside a 5%% drawdown fishbowl where <strong><code>97.6%%</code> of starters end the year in the red</strong> (<code>[F-008]</code>).
</div>
`, ch5PosiBot),
		},
		// Page 41: Chapter 6 Opener — The Verdict (Recto)
		{
			PageNum:     41,
			ThemeClass:  "theme-phase2",
			RunningHead: "PHASE 2: DEMOLITION · CHAPTER 6: THE VERDICT",
			FooterNote:  "THE 4-METRIC GATE SCORECARD [F-001..F-014]",
			BodyHTML: fmt.Sprintf(`
<div class="phase-badge" style="background:#991B1B;">PHASE 2 · DEMOLITION // CHAPTER 06 · RECTO OPENER</div>
<h1>CHAPTER 6: THE VERDICT</h1>
<p style="font-family:'Consolas',monospace; font-size:8.4pt; color:#991B1B; font-weight:700;">THE 4-METRIC MECHANICAL GATE SCORECARD · ZERO DIALS, PURE ARITHMETIC</p>
<p>We don't grade business models on vibes, rented Lamborghinis, or cropped screenshots of blue MetaTrader numbers (<code>[C-013]</code>). Every venture in the Y-IT canon runs through the exact same four mechanical gates defined in <code>outline-spine.v1.yaml §5.3</code>. Unless all four pass, the verdict is automatic.</p>
<div class="asset-frame"><img src="%s" alt="Four Metric Gate Scorecard"></div>
`, ch6Scorecard),
		},
		// Page 42: Ch 6 Gate-by-Gate Forensic Breakdown (Verso)
		{
			PageNum:     42,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 6: THE VERDICT · GATE-BY-GATE BREAKDOWN",
			FooterNote:  "AUDITED GATE COMPUTATION [F-001, F-003, F-004, F-007, F-008, F-009, F-013, F-014]",
			BodyHTML: `
<h2>THE FOUR GATES: LINE BY LINE</h2>
<table class="yit-table">
  <tr><th>Gate Metric</th><th>The Question &amp; Passing Rule</th><th>Audited Figure</th><th>Gate Status</th></tr>
  <tr><td><strong>1. THE STAKE</strong></td><td>What must I put at risk before I know?<br><em>Passes when &le; $1,000 US emergency savings (<code>[F-014]</code>).</em></td><td><strong><code>$1,840</code></strong><br>(4.2 tries + $412 tools, <code>[F-001, F-004]</code>)</td><td><span class="stamp-badge">MISS (+84% OVER)</span></td></tr>
  <tr><td><strong>2. THE WAIT</strong></td><td>How long until it pays?<br><em>Passes when &le; 6.0 months to net profit.</em></td><td><strong><code>14.0 months</code></strong><br>(Median survivor runway, <code>[F-009]</code>)</td><td><span class="stamp-badge">MISS (2.3x LONGER)</span></td></tr>
  <tr><td><strong>3. THE ODDS</strong></td><td>How many make it at 12 months?<br><em>Passes when &ge; 50.0% (1 in 2 starters).</em></td><td><strong><code>2.4%</code></strong><br>(6.6% pass Phase 1+2, <code>[F-007, F-008]</code>)</td><td><span class="stamp-badge">MISS (97.6% LOSE)</span></td></tr>
  <tr><td><strong>4. THE PAYOFF</strong></td><td>What does winning pay per hour?<br><em>Passes when &ge; $29.76/hr US median wage (<code>[F-013]</code>).</em></td><td><strong><code>$6.40 / hr</code></strong><br>(Across 26.5 hrs/wk, <code>[F-003]</code>)</td><td><span class="stamp-badge">MISS (-78% BELOW)</span></td></tr>
</table>
<div class="callout-rc">
  <strong>COMPUTED VERDICT RULE (<code>outline-spine.v1.yaml §5.3</code>):</strong><br>
  • <strong>Gates Passed:</strong> <code>0 of 4</code>.<br>
  • <strong>Mandatory Series Verdict:</strong> <strong>PROBABLY WON'T WORK.</strong>
</div>`,
		},
		// Page 43: Ch 6 House Edge Comparison & PosiBot Receipt Tape (Recto)
		{
			PageNum:     43,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 6: THE VERDICT · HOUSE EDGE COMPARISON",
			FooterNote:  "POSIBOT DELIVERY METHOD: RECEIPT TAPE [F-004, F-008]",
			BodyHTML: fmt.Sprintf(`
<h2>HOW PROP CHALLENGES STACK UP AGAINST LAS VEGAS &amp; REAL WORK</h2>
<table class="yit-table">
  <tr><th>Activity / Venture</th><th>Median Stake</th><th>12-Month / Session Odds</th><th>Median Hourly / Annual Return</th></tr>
  <tr><td><strong>Retail Prop Challenge (<code>[F-001..F-009]</code>)</strong></td><td><code>$1,840</code></td><td><strong>2.4%%</strong> net-profitable (12mo)</td><td><strong>$6.40 / hr</strong> (for the top 2.4%%)</td></tr>
  <tr><td><strong>Single-Zero European Roulette</strong></td><td><code>$100</code> table</td><td><strong>48.6%%</strong> single-spin win</td><td>-2.70%% house edge (Free drinks included)</td></tr>
  <tr><td><strong>Basic-Strategy Blackjack</strong></td><td><code>$100</code> table</td><td><strong>~49.5%%</strong> hand odds</td><td>-0.50%% house edge (No $149/mo Discord)</td></tr>
  <tr><td><strong>SQL / Financial BI Analyst (<code>[F-016]</code>)</strong></td><td><code>$320</code> prep</td><td><strong>78.0%%</strong> placement (12mo)</td><td><strong>$44.50 / hr ($92,560 / yr)</strong></td></tr>
</table>
<div class="asset-frame"><img src="%s" alt="PosiBot Receipt Tape Verdict"></div>
`, ch6PosiBot),
		},
		// Page 44: Ch 6 Who Still Could & End of Phase 2 (Verso)
		{
			PageNum:     44,
			ThemeClass:  "theme-phase2",
			RunningHead: "CHAPTER 6: THE VERDICT · WHO STILL COULD?",
			FooterNote:  "END OF PHASE 2 (DEMOLITION) · PHASE 3 OPENS ON PAGE 45",
			BodyHTML: `
<h2>WHO STILL COULD? (NAMING THE WINNERS' EDGES)</h2>
<p>Why does the cover say <strong>"Probably Won't Work"</strong> instead of "Never Works"? Because <code>2.4%</code> is not zero (<code>[F-008]</code>), and because Chapter 4 showed us the exact structural edges the system actually rewards:</p>
<p><strong>1. The Built-In Media Audience (Mold E2, <code>Marcus Conviction™</code>, <code>[K-001..K-003]</code>):</strong> If you already command <code>50,000+</code> followers, the prop-firm affiliate program pays you <code>$108</code> every time a viewer buys a <code>$540</code> challenge (<code>[F-024]</code>)—yielding <code>$54,000</code> in commissions even when your own trading net is <code>$0</code>.</p>
<p><strong>2. Pre-Existing Quantitative &amp; Emotional Capital (Mold C1, <code>Hannah Patience</code>, <code>[K-010..K-012]</code>):</strong> If you already have 5+ years of institutional risk-sizing discipline and 14 months of spare runway, you can extract a positive payout—though the hourly math still caps out near <code>$1.11</code> to <code>$6.40/hr</code> (<code>[F-003, K-010]</code>).</p>
<div class="verdict-band">
  <strong>PHASE 2 FINAL VERDICT BAND:</strong> <strong>0 of 4 gates passed.</strong> You are risking <code>$1,840</code> and waiting <code>14.0 months</code> for a <code>2.4%</code> shot at <code>$6.40/hr</code>. Turn to Page 45: let's salvage the skill and leave the casino behind.
</div>`,
		},
	}
}

func buildPhase3Pages45To56(rootDir string) []KDPPage {
	p3Assets := filepath.Join(rootDir, "05_phase3_run", "assets")
	ch7RangeChart := fileURI(filepath.Join(p3Assets, "T01-P3-C07-chart-pivot-payoff-comparison_v001.svg"))
	ch7PosiBot := fileURI(filepath.Join(p3Assets, "T01-P3-C07-posibot-keep-the-day-job_v001.svg"))
	ch7ComicSVG := fileURI(filepath.Join(p3Assets, "T01-P3-C07-comic-papercraft-diorama-strip_v001.svg"))
	ch7ComicRef := fileURI(filepath.Join(rootDir, "_canon", "assets", "comics_papercraft", "comic_3d_papercraft_strip_ref.jpg"))
	ch8PosiBot := fileURI(filepath.Join(p3Assets, "T01-P3-C08-posibot-ninety-day-paper-test_v001.svg"))

	return []KDPPage{
		// Page 45: Chapter 7 Opener — The Pivot (Recto)
		{
			PageNum:     45,
			ThemeClass:  "theme-phase3",
			RunningHead: "PHASE 3: SOLIDARITY · CHAPTER 7: THE PIVOT",
			FooterNote:  "VOICE: CONSTRUCTIVE SARDONIC AUTHORITY [F-015..F-018]",
			BodyHTML: `
<div class="phase-badge" style="background:#15803D;">PHASE 3 · SOLIDARITY // CHAPTER 07 · RECTO OPENER</div>
<h1>CHAPTER 7: THE PIVOT</h1>
<p style="font-family:'Consolas',monospace; font-size:8.4pt; color:#166534; font-weight:700;">FOR THE READER WHO GOT THE POINT · TURNING "PROBABLY WON'T" INTO WHAT PROBABLY WILL</p>
<h3>Stop Paying $340 Cover Charges to Work the Night Shift</h3>
<p>Let's get one thing straight before we look at the alternatives: if you spent six months teaching yourself how CPI interest-rate differentials move currency pairs, how to debug PineScript arrays, and how to log 200 backtest rows in a spreadsheet at midnight, <strong>there is nothing wrong with your drive</strong>. Most people won't read a cereal box, let alone an FOMC minutes release.</p>
<p>Your only error was taking a genuinely valuable quantitative brain and volunteering to pay <code>$340</code> cover charges (<code>[F-001]</code>) to play inside an offshore B-book arcade (<code>[F-011]</code>). When you take that exact same obsession with time-series data, variance, and automation and bill it to employers or B2B clients with real balance sheets, all four gate metrics flip from red to green.</p>
<div class="gate-strip" style="border-left-color:#DC2626;">
  TOPIC REFERENCE STRIP (PROP CHALLENGES): [○ Stake: $1,840] [○ Wait: 14.0m] [○ Odds: 2.4%] [○ Payoff: $6.40/hr] ──► 0 OF 4
</div>`,
		},
		// Page 46: Ch 7 Practical Pivots #1 & #2 (Verso)
		{
			PageNum:     46,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 7: THE PIVOT · PRACTICAL PIVOTS #1 & #2",
			FooterNote:  "TIER 1 PIVOTS (PASS 4 OF 4 GATES) [F-015, F-017, F-018]",
			BodyHTML: `
<h2>THE 4 GATE-SCORED PRACTICAL PIVOTS (PART I)</h2>
<div class="callout-pivot">
  <strong>[PIVOT: THE PRACTICAL PIVOT] #01 — FREELANCE PINESCRIPT / PYTHON BACKTESTING CONTRACTOR</strong><br>
  • <strong>Kind:</strong> Sell the skill directly (freelance or contract, <code>[F-015, F-017]</code>) · <strong>Tier 1 (Passes 4 of 4 Gates)</strong><br>
  • <strong>Gate Strip:</strong> <code>[● Stake: $150] [● Wait: 2.5 mo] [● Odds: 62%] [● Payoff: $55.00/hr]</code><br>
  • <strong>Failed Gates Answered:</strong> Cuts Stake by 92% ($150 vs $1,840), cuts Wait to 2.5 months, lifts Odds 25x (62% vs 2.4%), and lifts Payoff 8.5x ($55.00/hr vs $6.40/hr).<br>
  • <strong>First Step:</strong> Package 3 clean Python/PineScript strategy backtest templates with walk-forward Sharpe/drawdown reports and sell fixed-price specification testing to retail educators and funds. Whether their strategy wins or loses, your invoice clears on Friday.
</div>
<div class="callout-pivot">
  <strong>[PIVOT: THE PRACTICAL PIVOT] #02 — REGULATED SERIES 65 PARAPLANNER &amp; ADVISORY OPS</strong><br>
  • <strong>Kind:</strong> A skill-building route to a better-paid role (<code>[F-017, F-018]</code>) · <strong>Tier 1 (Passes 4 of 4 Gates)</strong><br>
  • <strong>Gate Strip:</strong> <code>[● Stake: $185] [● Wait: 3.5 mo] [● Odds: 68%] [● Payoff: $38.90/hr ($80,912/yr)]</code><br>
  • <strong>Failed Gates Answered:</strong> Replaces unregulated B-book challenges with a <code>$185</code> NASAA Series 65 exam credential (no firm sponsorship required) hired by fee-only Registered Investment Advisors (RIAs) for portfolio rebalancing and risk modeling (<code>[F-018]</code>).
</div>`,
		},
		// Page 47: Ch 7 Practical Pivots #3 & #4 + Range Strip Chart (Recto)
		{
			PageNum:     47,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 7: THE PIVOT · PRACTICAL PIVOTS #3 & #4",
			FooterNote:  "TIER 1 PIVOTS (PASS 4 OF 4 GATES) [F-015, F-016, F-017]",
			BodyHTML: fmt.Sprintf(`
<div class="callout-pivot">
  <strong>[PIVOT: THE PRACTICAL PIVOT] #03 — SQL &amp; PYTHON FINANCIAL BI / RISK DATA ANALYST (SOC 15-2051)</strong><br>
  • <strong>Kind:</strong> Get paid to use the skill (employment, <code>[F-015, F-016]</code>) · <strong>Tier 1 (Passes 4 of 4 Gates)</strong><br>
  • <strong>Gate Strip:</strong> <code>[● Stake: $320] [● Wait: 4.0 mo] [● Odds: 78%%] [● Payoff: $44.50/hr ($92,560/yr)]</code><br>
  • <strong>How It Works:</strong> An OHLCV candlestick chart is just a time-series SQL table with four floats and a timestamp. Convert your trading journal into an automated SQL/PowerBI variance dashboard (<code>[F-016]</code>).
</div>
<div class="callout-pivot">
  <strong>[PIVOT: THE PRACTICAL PIVOT] #04 — CORPORATE FX TREASURY &amp; HEDGING ANALYST (SOC 13-2051)</strong><br>
  • <strong>Kind:</strong> Get paid to use the skill (employment, <code>[F-017]</code>) · <strong>Tier 1 (Passes 4 of 4 Gates)</strong><br>
  • <strong>Gate Strip:</strong> <code>[● Stake: $650] [● Wait: 4.5 mo] [● Odds: 71%%] [● Payoff: $49.20/hr ($102,336/yr)]</code><br>
  • <strong>How It Works:</strong> Real import/export corporations pay <code>$102,336/yr</code> to analysts who hedge genuine EUR/USD invoice exposure using forward contracts—and they give you dental insurance instead of a 5%% drawdown trapdoor.
</div>
<div class="asset-frame"><img src="%s" alt="Pivot Payoff Comparison Range Strip"></div>
`, ch7RangeChart),
		},
		// Page 48: Ch 7 Skill Exchange Table, Repurposing Matrix & PosiBot (Verso)
		{
			PageNum:     48,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 7: THE PIVOT · SKILL EXCHANGE & REPURPOSING",
			FooterNote:  "FULL-MEASURE TABLE (ADR 0026) [F-015..F-018, F-027]",
			BodyHTML: fmt.Sprintf(`
<h2>SKILL EXCHANGE TABLE (<code>ch7.skills</code> — ADR 0026)</h2>
<table class="yit-table">
  <tr><th>Skill Trained by Retail FX Study</th><th>Who Actually Pays for It</th><th>Audited Typical Pay (<code>[F-015..F-018]</code>)</th></tr>
  <tr><td><strong>Time-Series Data Cleaning (SQL/Python)</strong></td><td>Corporate FP&amp;A, Fintech &amp; BI Teams (<code>SOC 15-2051</code>)</td><td><strong><code>$44.50 / hr</code></strong> ($92,560/yr)</td></tr>
  <tr><td><strong>PineScript / MQL5 / Backtest QA</strong></td><td>Quant Educators, Funds &amp; Technical Spec Clients</td><td><strong><code>$55.00 / hr</code></strong> (Contract)</td></tr>
  <tr><td><strong>Macro Calendar &amp; Rate Variance</strong></td><td>Commercial FX Treasury &amp; Hedging Desks (<code>SOC 13-2051</code>)</td><td><strong><code>$49.20 / hr</code></strong> ($102,336/yr)</td></tr>
  <tr><td><strong>Drawdown &amp; Position-Sizing Math</strong></td><td>Fee-Only RIAs &amp; Wealth Paraplanners (<code>Series 65</code>)</td><td><strong><code>$38.90 / hr</code></strong> ($80,912/yr)</td></tr>
</table>
<h3>The 6-Item Repurposing Matrix (<code>[F-027]</code>)</h3>
<p><strong>1. Charting Fluency</strong> $\rightarrow$ Executive PowerBI/Tableau Dashboards (<code>$44.50/hr</code>). <strong>2. PineScript</strong> $\rightarrow$ Python QA Automation (<code>$55.00/hr</code>). <strong>3. CPI/FOMC Tracking</strong> $\rightarrow$ Corporate Treasury Variance Briefs (<code>$49.20/hr</code>). <strong>4. Multi-Monitor Desk</strong> $\rightarrow$ Remote BI Workstation (<code>[F-020]</code>). <strong>5. ICT "Fair Value Gap" Lore</strong> $\rightarrow$ <strong>SUNK (<code>$0.00</code> employer value)</strong>. <strong>6. Expired Reset Fees</strong> $\rightarrow$ <strong>SUNK</strong> (unless inside 120-day chargeback window, <code>[F-021]</code>).</p>
<div class="asset-frame"><img src="%s" alt="PosiBot Flip Chart SQL"></div>
`, ch7PosiBot),
		},
		// Page 49: Ch 7 Full-Page 3D Papercraft Comic Strip (Recto)
		{
			PageNum:     49,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 7: THE PIVOT · [PAGE COMIC — 3D PAPERCRAFT]",
			FooterNote:  "VISUAL RELIEF PLATE · ME ODDEN CAR TOON PAPERCRAFT ARCHIVE",
			BodyHTML: fmt.Sprintf(`
<h2>[PAGE COMIC — 3D PAPERCRAFT DIORAMA]</h2>
<p>When PosiBot folds up the <code>$200,000</code> simulated paper crown and walks across the hall to a real quantitative desk:</p>
<div class="asset-frame"><img src="%s" alt="3D Papercraft Diorama Strip SVG" style="max-height: 2.95in;"></div>
<div class="asset-frame"><img src="%s" alt="Canonical 3D Papercraft Comic Reference" style="max-height: 2.75in;"></div>
`, ch7ComicSVG, ch7ComicRef),
		},
		// Page 50: Ch 7 Slow-Built Edges & Closing Band Seam into Ch 8 (Verso)
		{
			PageNum:     50,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 7: THE PIVOT · SLOW-BUILT EDGES & SEAM HANDOFF",
			FooterNote:  "MANDATORY SEAM: ENDS WORD-FOR-WORD ON CH 8 TITLE (OUTLINE-SPINE §6.1)",
			BodyHTML: `
<h2>BUILDING THE WINNERS' EDGES THE SLOW WAY (<code>ch7.edges</code>, <code>[F-018]</code>)</h2>
<p>What if you looked at the winners' structural edges in Chapter 4 and want to build them honestly without buying reset tickets?</p>
<p>• <strong>The Built-In Audience Edge (Mold E2, <code>[K-001..K-003]</code>):</strong> Instead of renting a balcony in Dubai to hawk <code>20%</code> prop reset codes (<code>[F-024]</code>), publish transparent open-source Python/SQL market studies on GitHub and Substack over 12–18 months to attract <code>$55.00/hr</code> engineering and analytics retainers (<code>[F-017]</code>).</p>
<p>• <strong>The Institutional Seat Edge (Mold E1 / E3):</strong> Instead of paying <code>$1,840</code> into unregulated B-book challenges (<code>[F-001]</code>), spend <code>$185–$650</code> on regulated Series 65 / Series 3 / CFA credentials and let a licensed firm pay you <code>$80,912–$102,336/yr</code> while you sit on the institutional side of the desk (<code>[F-017, F-018]</code>).</p>
<div class="verdict-band" style="margin-top: 1.1in;">
  <strong>CHAPTER 7 CLOSING BAND (<code>outline-spine.v1.yaml §6.1</code>):</strong><br>
  There is the ledger in plain black ink: <code>2.4%</code> odds and <code>$6.40</code> an hour inside the prop simulator (<code>[F-003, F-008]</code>), versus <code>62%</code> to <code>78%</code> odds and <code>$38.90</code> to <code>$55.00</code> an hour selling the exact same analytical habit to companies that pay on Fridays (<code>[F-016, F-017]</code>). If you are closing MetaTrader and taking one of those four off-ramps, we wish you every dollar of that raise; you don't need another page of this book. Only readers who haven't got the point and still want to keep trudging forward need what comes next… And… <strong>For the Stubborn Fool Who Must</strong>
</div>`,
		},
		// Page 51: Chapter 8 Opener — For the Stubborn Fool Who Must (Recto)
		{
			PageNum:     51,
			ThemeClass:  "theme-phase3",
			RunningHead: "PHASE 3: SOLIDARITY · CHAPTER 8: FOR THE STUBBORN FOOL WHO MUST",
			FooterNote:  "VOICE: QUIET RECKONING · EARNED WARMTH & ARMOR [F-023]",
			BodyHTML: `
<div class="phase-badge" style="background:#15803D;">PHASE 3 · SOLIDARITY // CHAPTER 08 · RECTO OPENER</div>
<h1>CHAPTER 8: FOR THE STUBBORN FOOL WHO MUST</h1>
<p style="font-family:'Consolas',monospace; font-size:8.4pt; color:#166534; font-weight:700;">FOR THE STUBBORN READER WHO KEEPS GOING ANYWAY · ZERO PERSONAL ADVICE, PURE ROLL-CAGE STEEL</p>
<h3>1. The Honest Way: Four Evidence-Backed Risk Reducers (<code>[F-023]</code>)</h3>
<p>Still here? All right. Maybe you love the tape, maybe you have an itch a corporate BI salary won't scratch, or maybe you looked at <code>2.4%</code> (<code>[F-008]</code>) and muttered, <em>"So you're telling me 24 out of 1,000 pull it off."</em> Because <strong>"Probably won't"</strong> still leaves a sliver of <strong>"Could,"</strong> we aren't going to preach at you. We are going to weld a roll cage around your checking account so the B-book casino can't bleed you out:</p>
<p><strong>1. Ditch Prop Evaluations for Regulated Micro-Lot Cash Accounts (Improves The Stake &amp; The Odds):</strong> Trade CFTC/NFA-regulated micro-lots (<code>0.01</code> lot) or CME Micro FX Futures (<code>M6E</code>) in a personal brokerage account you own (<code>[F-001, F-011]</code>). Zero <code>$340</code> reset fees, zero 5% trailing drawdown tripwires (<code>[F-007]</code>), and real exchange clearing.</p>
<p><strong>2. Hard-Cap Effective Leverage at 3:1 to 5:1 (Improves The Odds):</strong> CFTC filings (<code>[F-012]</code>) prove 50:1 to 500:1 leverage exhausts median retail accounts in <code>4.5 months</code>.</p>
<p><strong>3. Lock Daily Loss Circuit Breakers at 0.5% of Real Cash (Improves The Stake, <code>[F-023]</code>).</strong></p>
<p><strong>4. Clock Every Screen Hour Against <code>$29.76/hr</code> (Improves The Payoff, <code>[F-003, F-013]</code>).</strong></p>`,
		},
		// Page 52: Ch 8 Smallest Honest Test, Stop Rules & PosiBot (Verso)
		{
			PageNum:     52,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 8: FOR THE STUBBORN FOOL · 90-DAY TEST & STOP RULES",
			FooterNote:  "AUDITED SOURCES: [F-013, F-014, F-023, F-026]",
			BodyHTML: fmt.Sprintf(`
<h2>2. THE SMALLEST HONEST TEST &amp; PERSONAL STOP-RULE TRIPWIRES (<code>[F-023, F-026]</code>)</h2>
<div class="callout-pivot">
  <strong>THE SMALLEST HONEST TEST (90 DAYS · 100 TRADES · $0 FEES):</strong><br>
  Before you ever wire another dollar to an evaluation portal, execute <strong>100 consecutive logged trades over 90 days (3.0 months)</strong> on a free regulated paper account or <code>$250</code> micro-lot cash account with <strong><code>$0.00</code> spent on courses, signals, or prop challenges</strong> (<code>[F-023, F-026]</code>). If net expectancy after realistic spread/commission is negative or max drawdown hits <code>4.0%%</code>, the experiment stops automatically.
</div>
<h3>Personal Stop-Rule Tripwires (Decided in Advance, <code>[F-026]</code>)</h3>
<p>• <strong>The Stake Tripwire:</strong> Hard lifetime out-of-pocket ceiling of <strong><code>$500</code></strong> (half of the <code>$1,000</code> Federal Reserve household emergency buffer, <code>[F-014, F-026]</code>).<br>
• <strong>The Wait Tripwire:</strong> Hard time ceiling of <strong><code>6.0 months</code></strong> to reach verified net-positive return (<code>[F-026]</code>).<br>
• <strong>The Payoff Tripwire:</strong> Hard exit if trailing 90-day net hourly earnings stay below <strong><code>$29.76/hr</code></strong> at Month 12 (<code>[F-013, F-026]</code>).</p>
<div class="asset-frame"><img src="%s" alt="PosiBot Picket Sign 90 Days"></div>
`, ch8PosiBot),
		},
		// Page 53: Ch 8 Six Hard Do-Nots (Recto)
		{
			PageNum:     53,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 8: FOR THE STUBBORN FOOL · SIX HARD DO-NOTS",
			FooterNote:  "ORDERED BY COST · CROSS-REFERENCED TO CH 3–6 [F-025]",
			BodyHTML: `
<h2>3. SIX HARD DO-NOTS (ORDERED BY COST, <code>[F-025]</code>)</h2>
<p>Every trap below names what the <em>system</em> engineered to harvest your deposit—and points back to the exact exhibit in Chapters 3–6 that proved it:</p>
<table class="yit-table">
  <tr><th>#</th><th>The System Trap to Refuse</th><th>Median Cost Saved</th><th>Book Cross-Reference</th></tr>
  <tr><td><code>01</code></td><td><strong>Never buy an evaluation "Reset"</strong> after a drawdown breach.</td><td><strong><code>$1,428</code></strong><br>(4.2 tries × $340, <code>[F-001]</code>)</td><td>Ch 3 Waterfall Chart (p. 20)</td></tr>
  <tr><td><code>02</code></td><td><strong>Never join a $149/month VIP Discord Signal Room.</strong></td><td><strong><code>$1,341</code></strong><br>(9 months, <code>[K-004, F-024]</code>)</td><td>Ch 4 Case File #02: <em>Elena Diligence</em> (p. 28)</td></tr>
  <tr><td><code>03</code></td><td><strong>Never buy a $997–$1,497 "Smart Money" Upsell Academy.</strong></td><td><strong><code>$1,094</code></strong><br>(Median course rung, <code>[K-004]</code>)</td><td>Ch 4 Exhibit #02: Mold H4 Split Bar (p. 29)</td></tr>
  <tr><td><code>04</code></td><td><strong>Never buy a $499 "Guaranteed Prop-Pass EA Bot."</strong></td><td><strong><code>$499</code> + Account</strong><br>(Shared IP ban, <code>[K-007, K-008]</code>)</td><td>Ch 4 Case File #03: <em>Devon Precision</em> (p. 30)</td></tr>
  <tr><td><code>05</code></td><td><strong>Never hold B-book prop positions across CPI/NFP news releases.</strong></td><td><strong><code>$5,100</code> Frozen</strong><br>(18-pip slippage, <code>[F-010, K-015]</code>)</td><td>Ch 4 Case File #05: <em>Victor Discipline</em> (p. 34)</td></tr>
  <tr><td><code>06</code></td><td><strong>Never touch household emergency savings for margin or fees.</strong></td><td><strong><code>$1,000</code> Buffer</strong><br>(Fed SHED benchmark, <code>[F-014]</code>)</td><td>Ch 6 Gate Scorecard (p. 41)</td></tr>
</table>`,
		},
		// Page 54: Ch 8 Six Clean Exit Steps (Verso)
		{
			PageNum:     54,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 8: FOR THE STUBBORN FOOL · SIX CLEAN EXIT STEPS",
			FooterNote:  "ORDERED BY MONEY SAVED · STOP, PROTECT, RECOVER, SETTLE [F-019..F-021]",
			BodyHTML: `
<h2>4. SIX CLEAN EXIT STEPS (ORDERED BY MONEY SAVED, <code>[F-019..F-021]</code>)</h2>
<p><strong>1. <code>[STOP]</code> Kill VIP Signal Rooms &amp; Mentorship Rebills (Saves <code>$149/mo = $1,788/yr</code>, <code>[F-019]</code>):</strong> Revoke recurring card authorizations inside Whop, Stripe, and Discord billing settings today.</p>
<p><strong>2. <code>[STOP]</code> Cancel Third-Party Indicator Plugins (Saves <code>$68/mo = $816/yr</code>, <code>[F-019]</code>):</strong> Cancel LuxAlgo and custom oscillator subscriptions.</p>
<p><strong>3. <code>[STOP]</code> Downgrade TradingView Premium &amp; Terminate Forex VPS (Saves <code>$104/mo = $1,248/yr</code>, <code>[F-019]</code>):</strong> Drop charting to the <code>$0</code> Basic tier and shut down unused London VPS boxes.</p>
<p><strong>4. <code>[PROTECT]</code> File Credit Card Chargebacks for Frozen/Misrepresented Prop Evaluations Within 120 Days (Recovers <code>$340–$540</code>, <code>[F-021, K-015]</code>):</strong> When a prop firm shuts down, loses its MT5 license, or denies payouts via retroactive consistency rules, Visa and Mastercard dispute windows remain open for <strong>120 days</strong> from the transaction date. Contact your issuing bank's chargeback department with saved checkout promises.</p>
<p><strong>5. <code>[RECOVER]</code> Resell Secondary Monitors, StreamDecks &amp; Pro-Rata Annual Licenses (Recovers <code>$180–$450</code>, <code>[F-020]</code>).</strong></p>
<p><strong>6. <code>[SETTLE]</code> Separate Spot FX Losses from Evaluation Service Fees at Tax Time (<code>[F-014, F-021]</code>):</strong> Under U.S. IRC Section 988 and Section 1256, personal brokerage spot FX/futures losses follow specific tax-election rules, whereas prop evaluation fees are service outlays—<strong>ask a qualified CPA or Enrolled Agent</strong>.</p>`,
		},
		// Page 55: Ch 8 Cost/Keep Ledger & Watch-For Scams (Recto)
		{
			PageNum:     55,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 8: FOR THE STUBBORN FOOL · COST/KEEP LEDGER & WATCH-FOR",
			FooterNote:  "AUDITED SOURCES: [F-001, F-015..F-022]",
			BodyHTML: `
<h2>5. WHAT IT COST / WHAT YOU KEEP LEDGER (<code>ch8.ledger</code>)</h2>
<table class="yit-table">
  <tr><th>What It Typically Cost (<code>[F-001, K-004..K-012]</code>)</th><th>Typical Audited Amount</th><th>What You Actually Keep (<code>[F-015..F-020]</code>)</th><th>Audited Market Value</th></tr>
  <tr><td><strong>4.2× Evaluation &amp; Reset Fees</strong></td><td><code>-$1,428</code></td><td><strong>26.5 Hours/Week of Reclaimed Life</strong></td><td><code>+1,378 hrs / yr</code></td></tr>
  <tr><td><strong>Charting, VPS &amp; Indicator Tax</strong></td><td><code>-$412</code></td><td><strong>Time-Series SQL / Python / Charting Skill</strong></td><td><code>$44.50 / hr</code> (SOC 15-2051)</td></tr>
  <tr><td><strong>Courses &amp; VIP Signal Rooms (H4)</strong></td><td><code>-$2,435</code></td><td><strong>PineScript / Backtest QA Spec Skill</strong></td><td><code>$55.00 / hr</code> (Contract)</td></tr>
  <tr><td><strong>120-Day Chargeback + Gear Resale</strong></td><td><code>+$540 to +$720</code></td><td><strong>Macro-Calendar &amp; Treasury Fluency</strong></td><td><code>$49.20 / hr</code> (SOC 13-2051)</td></tr>
</table>
<h2>6. WATCH FOR: THE NEXT 3 VULTURES CIRCLING EXITING TRADERS (<code>[F-022]</code>)</h2>
<p>• <strong>1. "Guaranteed Prop-Pass AI Bots" (<code>$499–$997</code>):</strong> Routes identical trade-copier packets across hundreds of accounts, triggering automatic shared-IP termination (<code>[F-022, K-008]</code>).<br>
• <strong>2. Unregulated Offshore Crypto-Prop Challenges:</strong> The exact same 5% trailing drawdown woodchipper moved onto perpetual-swap tokens (<code>[F-022]</code>).<br>
• <strong>3. "Forex Loss Recovery" Upfront-Fee Agencies:</strong> Imposters demanding a <code>$500</code> "recovery tax" to unlock frozen accounts (<code>[F-022]</code>).</p>`,
		},
		// Page 56: Ch 8 Callbacks, Companion Pamphlet Segue & The Book's Last Word (Verso)
		{
			PageNum:     56,
			ThemeClass:  "theme-phase3",
			RunningHead: "CHAPTER 8: FOR THE STUBBORN FOOL · CALLBACKS & LAST WORD",
			FooterNote:  "END OF PHASE 3 · APPENDICES OPEN ON PAGE 57 [K-013, K-015]",
			BodyHTML: `
<h2>7. SIX-MONTH AFTERMATH CALLBACKS (<code>ch8.callbacks</code>, <code>[K-013, K-015]</code>)</h2>
<p>Remember <strong>Victor Discipline (Mold C3 — The Rule Change, <code>[K-013, K-015]</code>)</strong> from Chapter 4, whose <code>$5,100</code> funded profit cushion was frozen when the firm widened CPI slippage and lost its offshore MT5 license? Within 45 days of the freeze, Victor submitted his saved checkout terms to his credit card issuer inside the <strong>120-day Visa dispute window</strong> (<code>[F-021]</code>), clawed back <strong><code>$540</code> in evaluation fees</strong> (<code>[K-015]</code>), and used his Python/CSV backtesting scripts as a portfolio sample to land a <strong><code>$44.50/hr</code></strong> financial BI analyst role (<code>[F-016, K-015]</code>).</p>
<h2>8. THE COMPANION POCKET PAMPHLET SEGUE (<code>companion.segue</code>)</h2>
<p>If you want the printable <strong>90-Day Zero-Fee Paper-Trade Ledger (100 Trades)</strong>, the <strong>Prop-Firm Red-Flag Audit Checklist</strong>, and the <strong>120-Day Chargeback Dispute Letter Template</strong> on a single pocket reference, we publish the standalone <strong>Y-IT Companion Pocket Pamphlet</strong> for <strong><code>$1.99</code></strong> (no recurring rebills, no upsell ladder, no VIP Discord room).</p>
<div class="verdict-band">
  <strong>CHAPTER 8 CLOSING BAND — THE BOOK'S LAST WORD:</strong><br>
  The candles on the 1-minute chart will keep blinking green and red at 2:00 AM whether you are watching them or not. We pulled the CFTC filings, we counted the bodies in the B-book server room, and we taped the receipt to your monitor because we actually give a damn what happens to your <code>$1,840</code> and your Tuesday nights. Close the terminal, keep the quantitative brain, and go build something that pays you back.
</div>`,
		},
	}
}

func buildBackMatterPages57To64(rootDir string) []KDPPage {
	brianAvatar := fileURI(filepath.Join(rootDir, "_canon", "assets", "author_iiwii", "brian_iiwii_architect_avatar.jpg"))
	posibotSheet := fileURI(filepath.Join(rootDir, "_canon", "assets", "posibot", "posibot_blank_sign_pose_sheet.jpg"))
	rmLeftSVG := fileURI(filepath.Join(rootDir, "05_phase3_run", "assets", "T01-P3-C08-diagram-research-machine-left_v001.svg"))
	rmRightSVG := fileURI(filepath.Join(rootDir, "05_phase3_run", "assets", "T01-P3-C08-diagram-research-machine-right_v001.svg"))

	return []KDPPage{
		// Page 57: Appendix A Part 1 — Bulk Source List & Evidence Tiers (Recto)
		{
			PageNum:     57,
			ThemeClass:  "theme-back",
			RunningHead: "APPENDIX A · BULK SOURCE LIST & BIBLIOGRAPHY (PART I)",
			FooterNote:  "RECTO OPENER · TIERS 1–4 STATUTORY, CENTRAL BANK & COHORT BIBLIOGRAPHY",
			BodyHTML: `
<div class="phase-badge">BACK MATTER // APPENDIX A · RECTO OPENER</div>
<h1>APPENDIX A: BULK SOURCE LIST &amp; BIBLIOGRAPHY</h1>
<p>Per <code>outline-spine.v1.yaml §3</code> and <code>RM-10v2</code>, every core figure in this book is iterated across <strong>at least three independent sources of matching magnitude</strong> before integration, and classified into four evidentiary tiers (publisher/outlet, title, date, link, tier—never a person's name):</p>
<table class="yit-table">
  <tr><th>Tier</th><th>Canonical Definition</th><th>Role in Multi-Round Verification</th></tr>
  <tr><td><strong>Tier 1</strong></td><td>Academic, Statutory &amp; Central Bank</td><td>Primary federal dockets, central bank surveys, labor wage series &amp; peer-reviewed DOIs.</td></tr>
  <tr><td><strong>Tier 2</strong></td><td>Industry &amp; Regulatory Cohort Audits</td><td>Large-sample empirical cohorts (300,000–420,000+ evaluation accounts).</td></tr>
  <tr><td><strong>Tier 3</strong></td><td>Captured Platform &amp; Funnel Telemetry</td><td>SHA-256 archived checkout funnels, fee schedules &amp; affiliate CPA structures.</td></tr>
  <tr><td><strong>Tier 4</strong></td><td>Corroborated Practitioner Dossiers</td><td>14-Mold practitioner loss &amp; payout records corroborated by Tiers 1–3.</td></tr>
</table>
<h3>Bulk Source List — Primary Anchors (<code>01_research_backplane/sources.yaml</code>)</h3>
<p style="font-size:7.6pt; font-family:'Consolas',monospace; line-height:1.34;">
• <strong>[SRC-01]</strong> U.S. Commodity Futures Trading Commission (CFTC), <em>Retail Foreign Exchange Customer Advisory &amp; RFED Obligation Reports</em>, 2026-03-15, https://www.cftc.gov/LearnAndProtect/AdvisoriesAndArticles/fraudadv_forex.html (Tier 1)<br>
• <strong>[SRC-02]</strong> European Securities and Markets Authority (ESMA) &amp; UK FCA, <em>Product Intervention Measures on CFDs and Retail FX Loss Disclosures (74%–89% Loss Mandate)</em>, 2025-11-20, https://www.esma.europa.eu/press-news/esma-news/esma-adopts-final-product-intervention-measures-cfds-and-binary-options (Tier 1)<br>
• <strong>[SRC-03]</strong> Bank for International Settlements (BIS), <em>Triennial Central Bank Survey of Foreign Exchange and OTC Derivatives Markets ($7.5T/day)</em>, 2025-12-01, https://www.bis.org/statistics/rpfx22_fx.htm (Tier 1)<br>
• <strong>[SRC-04]</strong> U.S. CFTC Division of Enforcement, <em>Enforcement Action: Proprietary Trading Evaluation B-Book (98.4%) &amp; Asymmetric Slippage Findings</em>, 2025-09-01, https://www.cftc.gov/PressRoom/PressReleases/8771-23 (Tier 1)<br>
• <strong>[SRC-05]</strong> FCA &amp; Multi-Platform Cohort Benchmark, <em>Retail Funded Trader Evaluation Pass-Rate (5%–8%) &amp; Cost Cohort Study (420k Accounts)</em>, 2026-04-10, https://www.fca.org.uk/publications/research/retail-prop-trading-cohort-benchmarks (Tier 2)<br>
• <strong>[SRC-06]</strong> U.S. Bureau of Labor Statistics (BLS), <em>Occupational Employment and Wage Statistics (OEWS) National Benchmarks ($29.76/hr Median)</em>, 2026-05-01, https://www.bls.gov/oes/current/oes_nat.htm (Tier 1)<br>
• <strong>[SRC-07]</strong> Board of Governors of the Federal Reserve System, <em>Report on the Economic Well-Being of U.S. Households (SHED $1,000 Emergency Buffer)</em>, 2026-05-20, https://www.federalreserve.gov/publications/report-economic-well-being-us-households.htm (Tier 1)<br>
• <strong>[Turn the page for the Dual-Page Schematic of the Y-IT Research &amp; Verification Machine (Pages 58–59), followed immediately by Bulk Source List Part II &amp; the Chapter Citation Trail on Page 60 &rarr;]</strong>
</p>`,
		},
		// Page 58: Dual-Page Research Machine Spread — LEFT PLATE (Verso)
		{
			PageNum:     58,
			ThemeClass:  "theme-back",
			RunningHead: "APPENDIX A · EXHIBIT I (LEFT PLATE · VERSO): THE Y-IT RESEARCH & PROVENANCE MACHINE",
			FooterNote:  "DUAL-PAGE TECHNICAL SPREAD (LEFT PLATE · OPEN FLAT WITH PAGE 59) · ZERO BRAND FLUFF",
			BodyHTML: fmt.Sprintf(`
<div class="phase-badge" style="background:#0F172A;">APPENDIX A // DUAL-PAGE TECHNICAL SPREAD · LEFT PLATE (VERSO)</div>
<h2 style="margin-bottom:2px;">HOW WE CHECKED THE RECEIPTS (PART I: INGESTION &amp; CRYPTOGRAPHIC LOCK)</h2>
<p style="font-size:8.0pt; margin-bottom:4px;">We don't ask an AI chatbot what it "thinks" about a side hustle. Every Y-IT volume runs a 5-round live empirical harvest across statutory registries, peer-reviewed DOI indices, temporal DOM archives, a dedicated <strong>12GB Local VRAM + Elastic Cloud GPU Compute Array</strong>, and an on-demand <strong>Columnar OLAP Statistical Lakehouse</strong>, locking every raw byte under a <code>SHA-256</code> hash before a single sentence is drafted:</p>
<div class="asset-frame" style="padding:3px; margin: 2px -0.18in; background:#0F172A; border:2px solid #0F172A;">
  <img src="%s" alt="Y-IT Research Machine Left Plate — 5-Round Live Ingestion & Cryptographic Provenance" style="max-height:5.35in; width:100%%;">
</div>
<div style="font-family:'Consolas',monospace; font-size:7.4pt; background:#F8FAFC; border:1px solid #94A3B8; padding:4px 8px; margin: 2px -0.18in;">
  <strong>VERIFICATION TELEMETRY (ISSUE #01):</strong> <code>36</code> Live Multi-Index Queries Executed &bull; <code>140+</code> Primary Records Harvested &amp; Hashed &bull; <code>35</code> Offline Snapshot Captures Locked &bull; <strong>Continues across the gutter on Page 59 &rarr;</strong>
</div>`, rmLeftSVG),
		},
		// Page 59: Dual-Page Research Machine Spread — RIGHT PLATE (Recto)
		{
			PageNum:     59,
			ThemeClass:  "theme-back",
			RunningHead: "APPENDIX A · EXHIBIT II (RIGHT PLATE · RECTO): CONVERGENT RANGE SIEVE & COMPILER GATES",
			FooterNote:  "DUAL-PAGE TECHNICAL SPREAD (RIGHT PLATE · OPEN FLAT WITH PAGE 58) · DETERMINISTIC GATES G1–G7",
			BodyHTML: fmt.Sprintf(`
<div class="phase-badge" style="background:#0F172A;">APPENDIX A // DUAL-PAGE TECHNICAL SPREAD · RIGHT PLATE (RECTO)</div>
<h2 style="margin-bottom:2px;">HOW WE CHECKED THE RECEIPTS (PART II: CONVERGENCE, 14-MOLD SIEVE &amp; GATES)</h2>
<p style="font-size:8.0pt; margin-bottom:4px;"><strong>&larr; Continued from Page 58.</strong> Raw sources enter the <code>RM-10v2</code> Multi-Source Magnitude Convergence Sieve: every figure must be reiterated across <code>&ge;3</code> independent sources of matching magnitude and expressed as a <strong>defensible convergent range + audited cohort median</strong>. Practitioner cases are sorted across our <code>13+1</code> empirical failure molds, and the manuscript is compiled through seven deterministic Go gates (<code>G1–G7</code>):</p>
<div class="asset-frame" style="padding:3px; margin: 2px -0.18in; background:#0F172A; border:2px solid #0F172A;">
  <img src="%s" alt="Y-IT Research Machine Right Plate — Convergent Range Sieve, 14-Mold Taxonomy & Compiler Gates" style="max-height:5.35in; width:100%%;">
</div>
<div style="font-family:'Consolas',monospace; font-size:7.4pt; background:#F8FAFC; border:1px solid #94A3B8; padding:4px 8px; margin: 2px -0.18in;">
  <strong>WHY THIS MATTERS TO YOU:</strong> No single-source cherry-picking. No invented guru outliers. If a number cannot survive 3-source convergence and deterministic gate compilation, it cannot print in a Y-IT book.
</div>`, rmRightSVG),
		},
		// Page 60: Appendix A Part 2 — Bulk Source List Part II & Chapter Row Citation Trail (Verso)
		{
			PageNum:     60,
			ThemeClass:  "theme-back",
			RunningHead: "APPENDIX A (CONT.) · BULK SOURCE LIST PART II & CHAPTER CITATION TRAIL",
			FooterNote:  "LIVE HARVESTED DOCKETS, DOI REGISTRY & 56 AUDITED ROWS BY CHAPTER",
			BodyHTML: `
<h2>APPENDIX A (CONT.): BULK SOURCE LIST PART II &amp; CITATION TRAIL</h2>
<p style="font-size:7.3pt; font-family:'Consolas',monospace; line-height:1.28; margin-bottom:6px;">
• <strong>[LIVE-FEDREG]</strong> U.S. Federal Register Statutory Dockets (CFTC Part 5 Retail Foreign Exchange &amp; Commodity Pool Leverage Rules), <code>api/v1/documents.json</code>, SHA-256 verified in <code>01_research_backplane/live_harvest_ledger.yaml</code> (Tier 1)<br>
• <strong>[LIVE-OPENALEX / DOI]</strong> Peer-Reviewed Empirical Studies on Retail FX &amp; Day Trading Profitability (Barber, Lee, Liu &amp; Odean; Heimer &amp; Simsek Leverage Studies; Crossref DOI &amp; OpenAlex Works Registry, <code>140+</code> indexed records in <code>live_harvest_sources.csv</code>) (Tier 1)<br>
• <strong>[LIVE-ARXIV]</strong> Quantitative Finance Microstructure &amp; Absorbing Barrier Drawdown Probability Preprints (Ornstein-Uhlenbeck &amp; First-Passage Time Models on 5% Daily Trailing Locks) (Tier 1)<br>
• <strong>[LIVE-CDX / DOM]</strong> Temporal CDX &amp; Headless DOM Captures of Seller Funnels <code>S-01..S-06</code>, Prop Evaluation Pricing Tables, 20%–25% Affiliate CPA Portals &amp; 30% Consistency Rule Clauses (Tiers 3–4)
</p>
<h3>Chapter-by-Chapter Row Citation Trail (56 Audited Rows)</h3>
<table class="yit-table" style="font-size:7.4pt;">
  <tr><th>Chapter</th><th>Cited Row IDs</th><th>Convergent Range + Cohort Midpoint &amp; Evidentiary Scope</th></tr>
  <tr><td><strong>Ch 1: The Promise</strong></td><td><code>C-001..C-007, C-011..C-014</code></td><td>Pitch claims only (Sellers <code>S-01..S-06</code>, Tier 3 archive). Zero fact rows.</td></tr>
  <tr><td><strong>Ch 2: Sucker's Roadmap</strong></td><td><code>C-001, C-003, C-006, C-008..C-011, C-013</code></td><td>6-step brochure assembly + omitted Step 7 setup (<code>C-010</code>, Tier 3).</td></tr>
  <tr><td><strong>Ch 3: The Money Pit</strong></td><td><code>F-001..F-005, F-007, F-008, F-010..F-014</code></td><td><code>$1,400–$2,200</code> Stake (<code>$1,840</code> median), 98.4% B-book, <code>$4.50–$8.50/hr</code> Payoff (<code>$6.40/hr</code> median).</td></tr>
  <tr><td><strong>Ch 4: The Trenches</strong></td><td><code>K-001..K-015, F-021, F-024</code></td><td>5 empirical molds selected from 13+1 census (3 documented cases + caricature + exhibit each).</td></tr>
  <tr><td><strong>Ch 5: Exit Liquidity</strong></td><td><code>F-006..F-011, C-002, C-010</code></td><td><code>$7.5T/day</code> BIS survey, <code>5%–8%</code> pass (<code>6.6%</code> median), <code>1.5%–3.5%</code> 12mo Odds (<code>2.4%</code> median), <code>12–18mo</code> Wait.</td></tr>
  <tr><td><strong>Ch 6: The Verdict</strong></td><td><code>F-001, F-003, F-004, F-007..F-009, F-013, F-014</code></td><td>4-Metric Gate Scorecard (0 of 4 passed across full convergent range &amp; median).</td></tr>
  <tr><td><strong>Ch 7: The Pivot</strong></td><td><code>F-015..F-018, F-020, F-024, F-027</code></td><td>4 Tier-1 Practical Pivots (<code>$38.90/hr</code> to <code>$55.00/hr</code>, BLS OEWS Tier 1).</td></tr>
  <tr><td><strong>Ch 8: Stubborn Fool</strong></td><td><code>F-013, F-014, F-019..F-026, K-013, K-015</code></td><td>90-day test, <code>$500</code> Stake cap, 6 Do-Nots, 6 Exit Steps, 120-day chargeback, <code>BM-01..BM-08</code>.</td></tr>
</table>`,
		},
		// Page 61: Appendix B — Complete 20-Term Unvarnished Glossary (Recto)
		{
			PageNum:     61,
			ThemeClass:  "theme-back",
			RunningHead: "APPENDIX B · THE UNVARNISHED GLOSSARY (COMPLETE 20-TERM LEXICON)",
			FooterNote:  "RECTO OPENER · TRANSLATING GURU-SPEAK INTO PLAIN ARITHMETIC (TERMS 01–20)",
			BodyHTML: `
<div class="phase-badge">BACK MATTER // APPENDIX B · COMPLETE 20-TERM FORENSIC LEXICON</div>
<h1>APPENDIX B: THE UNVARNISHED GLOSSARY</h1>
<table class="yit-table" style="font-size:7.1pt; margin:4px -0.18in;">
  <tr><th>#</th><th>Guru / Industry Term</th><th>The Unvarnished Forensic Translation</th></tr>
  <tr><td><code>01</code></td><td><strong>"$200k Funded Account"</strong></td><td>A simulated demo login with a $10,000 daily loss tripwire wearing a $200,000 trench coat (<code>[F-007, F-011]</code>).</td></tr>
  <tr><td><code>02</code></td><td><strong>B-Book Execution</strong></td><td>Internal simulator where 98.4% of prop trades never touch a live market and reset fees fund payouts (<code>[F-011]</code>).</td></tr>
  <tr><td><code>03</code></td><td><strong>5% Daily Trailing Lock</strong></td><td>Absorbing barrier that ratchets up with unrealized open profit and liquidates you on a normal pullback (<code>[C-010, F-007]</code>).</td></tr>
  <tr><td><code>04</code></td><td><strong>Evaluation Reset ($340)</strong></td><td>The actual flagship revenue product of the retail prop industry (3–5 attempts = $1,400–$2,200, <code>[F-001]</code>).</td></tr>
  <tr><td><code>05</code></td><td><strong>Smart Money Concepts</strong></td><td>1930s Wyckoff chart patterns renamed with sci-fi acronyms so an influencer can charge $997 for a PDF (<code>[C-004, K-004]</code>).</td></tr>
  <tr><td><code>06</code></td><td><strong>Liquidity Sweep</strong></td><td>When your stop-loss gets triggered at 9:31 AM, explained as a personal vendetta by central banks (<code>[C-008]</code>).</td></tr>
  <tr><td><code>07</code></td><td><strong>90% Profit Split</strong></td><td>90% of zero, for the 96.5%–98.5% (97.6% median) of starters who are net-unprofitable at Month 12 (<code>[F-008]</code>).</td></tr>
  <tr><td><code>08</code></td><td><strong>30% Consistency Rule</strong></td><td>Retroactive fine-print clause used to deny payouts when one winning trade was "too large a share" of profit (<code>[K-013]</code>).</td></tr>
  <tr><td><code>09</code></td><td><strong>Funded Certificate</strong></td><td>Gold-embossed dark-mode JPEG with a referral QR code designed to recruit your friends into Phase 1 (<code>[C-013, F-024]</code>).</td></tr>
  <tr><td><code>10</code></td><td><strong>"Zero Capital at Risk"</strong></td><td>You risk $1,400–$2,200 ($1,840 median) in non-refundable fees—184% of the US emergency buffer (<code>[F-004, F-014]</code>).</td></tr>
  <tr><td><code>11</code></td><td><strong>A-Book Routing</strong></td><td>Sending orders to an external liquidity provider; happens on only 1.6% of investigated prop trades (<code>[F-011]</code>).</td></tr>
  <tr><td><code>12</code></td><td><strong>Asymmetric Slippage</strong></td><td>Server plugins that fill your stop-loss 18 pips worse during CPI while filling take-profits to the tick (<code>[F-010, K-015]</code>).</td></tr>
  <tr><td><code>13</code></td><td><strong>Prop-Pass EA Bot ($499)</strong></td><td>Automated script sold to 500 buyers at once that gets all 500 accounts banned for shared IP/EA signatures (<code>[K-007, K-008]</code>).</td></tr>
  <tr><td><code>14</code></td><td><strong>Affiliate CPA (20%–30%)</strong></td><td>The $108–$162 kickback a YouTuber collects every time you use their discount code on a $540 challenge (<code>[F-024, K-001]</code>).</td></tr>
  <tr><td><code>15</code></td><td><strong>VIP Signal Room ($149/mo)</strong></td><td>Chatroom where delayed entry alerts turn your account into the moderator's exit liquidity (<code>[K-004, F-019]</code>).</td></tr>
  <tr><td><code>16</code></td><td><strong>SIPC Protection</strong></td><td>U.S. securities account insurance that covers <strong>0%</strong> of offshore spot forex and simulated prop challenges (<code>[SRC-01]</code>).</td></tr>
  <tr><td><code>17</code></td><td><strong>120-Day Chargeback</strong></td><td>Visa/Mastercard dispute clock for recovering evaluation fees when a firm freezes payouts or shuts down (<code>[F-021, K-015]</code>).</td></tr>
  <tr><td><code>18</code></td><td><strong>IRC Section 988 / 1256</strong></td><td>U.S. tax code sections governing spot FX vs regulated futures losses (prop fees are service expenses, <code>[F-014]</code>).</td></tr>
  <tr><td><code>19</code></td><td><strong>The Four-Metric Gate</strong></td><td>Y-IT scorecard: Stake (&le;$1,000), Wait (&le;6mo), Odds (&ge;50%), Payoff (&ge;$29.76/hr) (<code>[F-001..F-014]</code>).</td></tr>
  <tr><td><code>20</code></td><td><strong>The Practical Pivot</strong></td><td>Billing employers $38.90–$55.00/hr for the exact same SQL, Python, and risk-modeling habits (<code>[F-015..F-018]</code>).</td></tr>
</table>`,
		},
		// Page 62: Tear-Out Stop-Loss Contract & 8 Beneficial Mechanisms (Verso)
		{
			PageNum:     62,
			ThemeClass:  "theme-back",
			RunningHead: "TEAR-OUT FIELD CONTRACT · CIRCUIT BREAKER & BENEFICIAL MECHANISMS",
			FooterNote:  "TAPE THIS PAGE TO YOUR MONITOR · MODULES P1–P7 IN THE $1.99 COMPANION PAMPHLET",
			BodyHTML: `
<h2>TEAR-OUT CONTRACT: MY PERSONAL CIRCUIT BREAKER</h2>
<p style="font-size:8.6pt;">If you are marching into Chapter 8 anyway, sign this box in ink and tape it to your monitor before entering a credit card:</p>
<div style="border: 2.5px solid #0F172A; padding: 10px; margin: 6px -0.18in; background: #FFFFFF; font-family: 'Consolas', monospace; font-size: 8.0pt;">
  <strong>1. THE 90-DAY ZERO-FEE RULE (<code>BM-05 / [F-023, F-026]</code>):</strong> I will log <code>100</code> consecutive trades across <code>90 days</code> on a free regulated demo or micro-lot account with <code>$0.00</code> spent on challenges, courses, or Discord rooms.<br>
  <strong>2. MY HARD STAKE CEILING (<code>$500 MAX</code>, <code>[F-014, F-026]</code>):</strong> I will never spend more than <code>$500</code> lifetime out-of-pocket (half of the $1,000 US emergency buffer). I will never purchase an evaluation "Reset" (<code>[F-001]</code>).<br>
  <strong>3. MY HARD WAIT &amp; PAYOFF TRIPWIRES (<code>[F-013, F-026]</code>):</strong> If I am not net-profitable by Month <code>6.0</code>, or if my net hourly return divided by total screen hours is below <code>$29.76/hr</code> at Month <code>12.0</code>, I will execute the 6 Exit Steps on Page 54.<br>
  <strong>SIGNED:</strong> ___________________________ &nbsp;&nbsp; <strong>DATE:</strong> _____________ &nbsp;&nbsp; <strong>WITNESS:</strong> ___________________
</div>
<h3>8 Beneficial Mechanisms Adapted into the $1.99 Companion Pamphlet (<code>BM-01..BM-08</code>)</h3>
<table class="yit-table" style="font-size:7.3pt;">
  <tr><th>ID</th><th>Beneficial Mechanism (In-The-Wild Proven)</th><th>Pamphlet Module</th><th>What It Replaces</th></tr>
  <tr><td><code>BM-01</code></td><td><strong>True-Risk Drawdown De-Masking Worksheet</strong></td><td>Module P1</td><td>"10% of $100k" nominal leverage illusion</td></tr>
  <tr><td><code>BM-02</code></td><td><strong>0.25% Quarter-Kelly Position Sizer + Lock</strong></td><td>Module P2</td><td>2% per-trade sizing into a 5% daily barrier</td></tr>
  <tr><td><code>BM-03</code></td><td><strong>CME Micro FX (M6E) Section 1256 Gate</strong></td><td>Module P3</td><td>Offshore unlicensed MT5 B-book portals</td></tr>
  <tr><td><code>BM-04</code></td><td><strong>T-Cost &amp; News-Window Blackout Filter</strong></td><td>Module P4</td><td>Trading across CPI/NFP spread-widening traps</td></tr>
  <tr><td><code>BM-05</code></td><td><strong>90-Day / 100-Trade Expectancy Gate</strong></td><td>Module P5</td><td>Buying $540 challenges on Day 1</td></tr>
  <tr><td><code>BM-06</code></td><td><strong>120-Day Chargeback &amp; Evidence Kit</strong></td><td>Module P6</td><td>Walking away when a prop firm freezes payouts</td></tr>
  <tr><td><code>BM-07</code></td><td><strong>B-Book Contract Fine-Print Forensics</strong></td><td>Module P6B</td><td>Blindly clicking "I Agree" on 30% consistency rules</td></tr>
  <tr><td><code>BM-08</code></td><td><strong>30-Day Quantitative Skill Translation Matrix</strong></td><td>Module P7</td><td>Donating SQL/Python hours to a $6.40/hr simulator</td></tr>
</table>`,
		},
		// Page 63: *Results Typical & Built at the IIWII Desk Colophon (Recto)
		{
			PageNum:     63,
			ThemeClass:  "theme-back",
			RunningHead: "RESULTS TYPICAL & IIWII COLOPHON · BRIAN AVATAR & POSIBOT SHEET",
			FooterNote:  "CANONICAL REQUIREMENT: OUTLINE-SPINE.V1.YAML §6.3 & ADR 0044",
			BodyHTML: fmt.Sprintf(`
<div class="phase-badge" style="background:#0F172A;">BACK MATTER // *RESULTS TYPICAL &amp; IIWII IMPRINT COLOPHON</div>
<h1>*RESULTS TYPICAL (AND WHO BUILT THIS)</h1>
<div class="callout-rc" style="background:#FFFFFF; border-color:#0F172A; border-left-width:7px; padding:7px 10px; margin:4px -0.18in;">
  <p style="font-size:8.4pt; margin:0;">Every sales page in Chapter 1 buries <em>"*Results not typical"</em> in 6-point gray type. <strong>In this book, every figure in Phases 2 and 3 IS the typical result:</strong> we print both the <strong>3-source convergent range</strong> and the <strong>audited cohort median</strong> (<code>$1,400–$2,200</code> Stake / <code>$1,840</code> median; <code>12–18 mo</code> Wait / <code>14.0 mo</code> median; <code>1.5%%–3.5%%</code> 12-mo Odds / <code>2.4%%</code> median; <code>$4.50–$8.50/hr</code> Payoff / <code>$6.40/hr</code> median)—never the lottery ticket.</p>
</div>
<p style="font-size:8.2pt; margin:4px 0;"><strong>Zero Guru Sponsorship &amp; Clean Monetization:</strong> Nobody selling a prop challenge, course, or Discord room paid a cent for this book. We sell the standalone <strong><code>$1.99</code> Y-IT Companion Pocket Pamphlet</strong> (<code>Modules P1–P7</code>). Any Chapter 7 training or certification links are vetted under <code>ADR 0050 Decision 11</code> and disclosed in 10-point bold type.</p>
<div style="display:flex; gap:10px; align-items:center; margin: 6px -0.18in; background:#FFFFFF; border:2px solid #0F172A; padding:8px; border-radius:6px;">
  <img src="%s" alt="Brian IIWII Architect Avatar" style="width:1.25in; height:1.25in; object-fit:cover; border-radius:6px; border:2px solid #0F172A;">
  <div style="font-size:8.1pt;">
    <strong>BUILT AT THE IIWII DESK (IT IS WHAT IT IS):</strong><br>
    <em>"Every guru on the internet wants <code>$997</code> to sell you the brochure in Chapter 1. We give you Chapters 1 and 2 for free—and then we hand you the CFTC subpoena, the coroner's ledger, the Dual-Page Research Machine schematic, and four real ways to get paid <code>$38.90</code> to <code>$55.00</code> an hour for the exact same analytical brain."</em>
  </div>
</div>
<h3 style="margin-top:4px; margin-bottom:2px;">Meet PosiBot (Canonical Pose Sheet · Appears in Phases 1 &amp; 3 Only)</h3>
<div class="asset-frame" style="padding:3px; margin:2px -0.18in;"><img src="%s" alt="PosiBot Blank Sign Pose Sheet" style="max-height:1.75in;"></div>
`, brianAvatar, posibotSheet),
		},
		// Page 64: Last-Page Issue Index & Series Catalog Grid (Verso)
		{
			PageNum:     64,
			ThemeClass:  "theme-back",
			RunningHead: "LAST-PAGE ISSUE INDEX · Y-IT SERIES CATALOG",
			FooterNote:  "FINAL PAGE (64 OF 64) · KDP PREFLIGHT LOCKED",
			BodyHTML: `
<h2>LAST-PAGE ISSUE INDEX &amp; COMPONENT DIRECTORY</h2>
<table class="yit-table" style="font-size:7.6pt;">
  <tr><th>Component / Metric / Composite / Exhibit</th><th>Primary Chapter &amp; Page</th><th>Audited Row / Spec IDs</th></tr>
  <tr><td><strong>The Stake ($1,400–$2,200 Range · $1,840 Median)</strong></td><td>Ch 3 (p. 20), Ch 6 (p. 41)</td><td><code>F-001, F-004, F-014</code></td></tr>
  <tr><td><strong>The Wait (12–18 Months Range · 14.0 Mo Median)</strong></td><td>Ch 5 (p. 39), Ch 6 (p. 41)</td><td><code>F-009</code></td></tr>
  <tr><td><strong>The Odds (1.5%–3.5% 12Mo Range · 2.4% Median)</strong></td><td>Ch 5 (p. 39), Ch 6 (p. 41)</td><td><code>F-007, F-008, F-010</code></td></tr>
  <tr><td><strong>The Payoff ($4.50–$8.50/hr Range · $6.40/hr Median)</strong></td><td>Ch 3 (p. 23), Ch 6 (p. 41)</td><td><code>F-003, F-013</code></td></tr>
  <tr><td><strong>B-Book Simulator (98.4% / $310M Enforcement)</strong></td><td>Ch 3 (pp. 21–22)</td><td><code>F-002, F-011</code></td></tr>
  <tr><td><strong>Marcus Conviction™ (Mold E2 + Caricature + Curve)</strong></td><td>Ch 4 Case #01 (pp. 26–27)</td><td><code>K-001..K-003, F-024</code></td></tr>
  <tr><td><strong>Elena Diligence (Mold H4 + Caricature + Invoice)</strong></td><td>Ch 4 Case #02 (pp. 28–29)</td><td><code>K-004..K-006</code></td></tr>
  <tr><td><strong>Devon Precision (Mold H5 + Caricature + Ban Modal)</strong></td><td>Ch 4 Case #03 (pp. 30–31)</td><td><code>K-007..K-009</code></td></tr>
  <tr><td><strong>Hannah Patience (Mold C1 + Caricature + Time Log)</strong></td><td>Ch 4 Case #04 (pp. 32–33)</td><td><code>K-010..K-012</code></td></tr>
  <tr><td><strong>Victor Discipline (Mold C3 + Caricature + Denial)</strong></td><td>Ch 4 Case #05 (pp. 34–35), Ch 8 (p. 56)</td><td><code>K-013..K-015, F-021</code></td></tr>
  <tr><td><strong>4 Practical Pivots ($38.90–$55.00/hr BLS OEWS)</strong></td><td>Ch 7 (pp. 46–48)</td><td><code>F-015..F-018, F-027</code></td></tr>
  <tr><td><strong>Bulk Source List Bibliography (Parts I &amp; II)</strong></td><td>Appendix A (p. 57 &amp; p. 60)</td><td><code>SRC-01..07, LIVE-140+</code></td></tr>
  <tr><td><strong>Dual-Page Research Machine Visualization Spread</strong></td><td>Appendix A Exhibits I–II (pp. 58–59)</td><td><code>RM-10v2, RM-12, G1–G7</code></td></tr>
  <tr><td><strong>90-Day Test, 120-Day Chargeback &amp; BM-01..BM-08</strong></td><td>Ch 8 (pp. 52–54), Tear-Out (p. 62)</td><td><code>F-019..F-023, F-026</code></td></tr>
</table>
<div class="verdict-band" style="text-align:center;">
  <strong>END OF ISSUE #01 (FOREX TRADING &amp; PROP FIRMS) · 64 PAGES · ITISWHATITIS.BIZ</strong>
</div>`,
		},
	}
}

// BuildKDPCoverWrapHTML builds the full 12.405" x 9.25" KDP Print Cover Wrap HTML (Back Cover + 0.144" Spine with no spine text + Front Cover + Barcode Safe Zone).
func BuildKDPCoverWrapHTML(rootDir string) string {
	blueprintCover := fileURI(filepath.Join(rootDir, "_canon", "assets", "covers_blueprint", "blueprint_cover_populated_ref.jpg"))
	posibotCut := fileURI(filepath.Join(rootDir, "_canon", "assets", "posibot", "pose_smiling_cut.png"))
	brianBadge := fileURI(filepath.Join(rootDir, "_canon", "assets", "author_iiwii", "brian_iiwii_badge_avatar.jpg"))

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>KDP Full Cover Wrap (12.405in x 9.25in) — Forex Trading &amp; Prop Firms</title>
<style>
  @page {
    size: 12.405in 9.25in;
    margin: 0;
  }
  * {
    box-sizing: border-box;
    -webkit-print-color-adjust: exact !important;
    print-color-adjust: exact !important;
  }
  html, body {
    margin: 0;
    padding: 0;
    width: 12.405in;
    height: 9.25in;
    overflow: hidden;
    background: #0F172A;
    font-family: 'Segoe UI', 'Work Sans', sans-serif;
    color: #F8FAFC;
  }
  .wrap-canvas {
    width: 12.405in;
    height: 9.25in;
    display: flex;
    flex-direction: row;
    position: relative;
    background-image:
      linear-gradient(to right, rgba(56,189,248,0.12) 1px, transparent 1px),
      linear-gradient(to bottom, rgba(56,189,248,0.12) 1px, transparent 1px);
    background-size: 0.25in 0.25in;
  }
  /* Left Back Cover: 0.125" bleed + 6.00" back trim = 6.125" */
  .back-cover {
    width: 6.125in;
    height: 9.25in;
    padding: 0.55in 0.48in 0.55in 0.60in;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border-right: 1px dashed rgba(148,163,184,0.4);
  }
  /* Center Spine: 64 pages * 0.002252" = 0.144" (KDP <80 pages rule: NO SPINE TEXT) */
  .spine-strip {
    width: 0.155in;
    height: 9.25in;
    background: #DC2626;
  }
  /* Right Front Cover: 6.00" front trim + 0.125" outer bleed = 6.125" */
  .front-cover {
    width: 6.125in;
    height: 9.25in;
    padding: 0.50in 0.58in 0.50in 0.48in;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    border-left: 1px dashed rgba(148,163,184,0.4);
  }
  .barcode-zone {
    width: 2.00in;
    height: 1.20in;
    background: #FFFFFF;
    color: #0F172A;
    border: 2px solid #0F172A;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    font-family: 'Consolas', monospace;
    font-size: 8pt;
    font-weight: 700;
  }
</style>
</head>
<body>
<div class="wrap-canvas">
  <!-- BACK COVER -->
  <div class="back-cover">
    <div>
      <div style="background:#DC2626; display:inline-block; padding:4px 10px; font-family:'Consolas',monospace; font-weight:800; font-size:9pt; border-radius:4px;">
        Y-IT FORENSIC FIELD GUIDE · ISSUE #01
      </div>
      <h2 style="font-family:'Impact',sans-serif; font-size:22pt; color:#FEF08A; margin:12px 0 8px 0; text-transform:uppercase;">
        A "$100,000 FUNDED ACCOUNT" WITH A 5%% DAILY DRAWDOWN IS JUST A $5,000 ACCOUNT IN A HALLOWEEN COSTUME.
      </h2>
      <p style="font-size:10pt; line-height:1.45; color:#E2E8F0;">
        Every night at 11:42 PM, thousands of smart, ambitious people wire <code>$497</code> to an offshore proprietary trading portal dreaming of <code>90%%</code> profit splits from a laptop in Dubai.
      </p>
      <p style="font-size:10pt; line-height:1.45; color:#E2E8F0;">
        Then the CFTC enforcement filings and cohort telemetry arrive: <strong><code>98.4%%</code></strong> of evaluation trades never leave an internal simulated B-book server. <strong><code>93.4%%</code></strong> of challenges breach Phase 1 or Phase 2. The median participant buys <strong><code>4.2</code> evaluations or resets (<code>$1,840</code> total Stake)</strong>, and the <strong><code>2.4%%</code></strong> who survive to Month 12 earn a median of <strong><code>$6.40 an hour</code></strong>.
      </p>
      <div style="background:#1E293B; border-left:5px solid #38BDF8; padding:10px 12px; margin:12px 0; font-family:'Consolas',monospace; font-size:8.6pt;">
        <strong>INSIDE THIS 3-PHASE TEARDOWN:</strong><br>
        • <strong>Phase 1 (Vibrant Hope):</strong> The 6-step brochure as sold—and the missing Step 7.<br>
        • <strong>Phase 2 (Demolition):</strong> 5 Amber Case Files &amp; the 4-Metric Gate Scorecard.<br>
        • <strong>Phase 3 (Solidarity):</strong> Four <code>$38.90–$55.00/hr</code> Practical Pivots + the 120-day Visa/Mastercard chargeback manual.
      </div>
    </div>
    <div style="display:flex; justify-content:space-between; align-items:flex-end;">
      <div style="display:flex; align-items:center; gap:10px;">
        <img src="%s" alt="IIWII Badge" style="width:0.85in; height:0.85in; border-radius:50%%; border:2px solid #FEF08A;">
        <div style="font-family:'Consolas',monospace; font-size:8pt; color:#94A3B8;">
          <strong>IIWII / Y-IT IMPRINT</strong><br>
          Models Propose, Programs Decide.<br>
          itiswhatitis.biz
        </div>
      </div>
      <div class="barcode-zone">
        <span>KDP BARCODE ZONE</span>
        <span>2.00" × 1.20" SAFE AREA</span>
        <span>ISBN 979-8-YIT-FX001</span>
      </div>
    </div>
  </div>

  <!-- SPINE (0.155" — NO SPINE TEXT UNDER 80 PAGES PER KDP SPEC) -->
  <div class="spine-strip"></div>

  <!-- FRONT COVER -->
  <div class="front-cover">
    <div style="background:#FFFFFF; color:#0F172A; padding:8px 12px; border:3px solid #0F172A; text-align:center; font-family:'Consolas',monospace; font-size:9.5pt; font-weight:800;">
      Why It <span style="text-decoration:line-through; color:#64748B;">Could. Would. Has. Might.</span> <span style="color:#DC2626;">Probably Won't.</span> WORK — <em>Yeah… about that.</em>
    </div>

    <div style="background:#FFFFFF; color:#0F172A; border:4px solid #0F172A; padding:14px; text-align:center; box-shadow:6px 6px 0px #DC2626;">
      <div style="font-family:'Impact','Big Shoulders Display',sans-serif; font-size:25pt; line-height:1.05; text-transform:uppercase;">
        why your FOREX PROP CHALLENGE probably will Fail
      </div>
      <div style="background:#DC2626; color:#FFFFFF; padding:5px 8px; margin-top:8px; font-family:'Consolas',monospace; font-size:10.5pt; font-weight:900;">
        EVEN IF YOUR AI CHAT BOT LIKES YOUR CHANCES!
      </div>
    </div>

    <div style="position:relative; text-align:center; background:#FFFFFF; border:3px solid #38BDF8; border-radius:6px; padding:6px;">
      <img src="%s" alt="Blueprint Cover Artwork" style="width:100%%; max-height:3.65in; object-fit:contain;">
      <img src="%s" alt="PosiBot Cutout" style="position:absolute; bottom:8px; right:10px; width:1.15in; height:auto;">
    </div>

    <div style="background:#FEF08A; color:#0F172A; border:3.5px solid #0F172A; padding:10px 12px; text-align:center;">
      <div style="font-weight:900; font-size:11pt; text-transform:uppercase;">
        We reveal the secret method in the first 2 chapters...
      </div>
      <div style="font-family:'Consolas',monospace; font-size:8.8pt; font-weight:700; color:#991B1B; margin-top:3px;">
        First 2 chapters free at itiswhatitis.biz · Audited Verdict: 0 of 4 Gates Passed
      </div>
    </div>
  </div>
</div>
</body>
</html>`, brianBadge, blueprintCover, posibotCut)
}

// BuildCompanionPamphletPages returns an 8-page standalone Companion Pocket Pamphlet incorporating BM-01..BM-08 (Modules P1..P7).
func BuildCompanionPamphletPages(rootDir string) []KDPPage {
	p3PosiBot := fileURI(filepath.Join(rootDir, "05_phase3_run", "assets", "T01-P3-C08-posibot-ninety-day-paper-test_v001.svg"))
	return []KDPPage{
		{PageNum: 1, ThemeClass: "theme-phase3", RunningHead: "Y-IT COMPANION POCKET PAMPHLET ($1.99)", FooterNote: "PAMPHLET PAGE 1 OF 8 · MODULES P1–P7 (BM-01..BM-08)", BodyHTML: `<h1>THE PRE-DEPOSIT PROP &amp; FOREX SURVIVAL MANUAL</h1><p><strong>Standalone Post-Sale Companion Pocket Pamphlet ($1.99) · Zero Subscriptions · Zero Guru Upsells</strong></p><div class="callout-pivot">Contains the 8 Proven In-The-Wild Beneficial Mechanisms (<code>BM-01..BM-08</code>) engineered into 7 Field Modules (<code>P1..P7</code>):<br>• <strong>Module P1 (<code>BM-01</code>):</strong> The True-Risk Drawdown De-Masking Worksheet<br>• <strong>Module P2 (<code>BM-02</code>):</strong> The 0.25% Quarter-Kelly Position Sizer &amp; 2-Strike Lock<br>• <strong>Module P3 (<code>BM-03</code>):</strong> Regulated CME Micro FX (<code>M6E</code>) vs Offshore B-Book Gate<br>• <strong>Module P4 (<code>BM-04</code>):</strong> The 30-Minute News-Window Blackout &amp; T-Cost Filter<br>• <strong>Module P5 (<code>BM-05</code>):</strong> The 90-Day / 100-Trade Zero-Fee Paper Expectancy Ledger<br>• <strong>Module P6 (<code>BM-06, BM-07</code>):</strong> The 120-Day Chargeback Dispute &amp; Fine-Print Audit Kit<br>• <strong>Module P7 (<code>BM-08</code>):</strong> The 30-Day Quantitative Skill Translation Matrix (<code>$38.90–$55.00/hr</code>)</div>`},
		{PageNum: 2, ThemeClass: "theme-phase3", RunningHead: "PAMPHLET · MODULES P1 & P2 (TRUE-RISK & QUARTER-KELLY)", FooterNote: "PAMPHLET PAGE 2 OF 8 · BM-01 & BM-02", BodyHTML: `<h2>MODULES P1 &amp; P2: DE-MASKING LEVERAGE &amp; QUARTER-KELLY SIZING (<code>[F-023]</code>)</h2><p><strong>Module P1 (<code>BM-01</code> — True-Risk De-Masking):</strong> Never size against nominal headline capital (<code>$100,000</code>). Your real capital is the <strong>Maximum Trailing Drawdown Allowance</strong> (<code>$5,000</code> daily / <code>$10,000</code> max).<br>• <strong>Module P2 (<code>BM-02</code> — 0.25% Nominal / 2.5% True-Buffer Cap):</strong> Cap per-trade risk at <code>0.25%</code> of nominal (<code>$250</code> on a $100k challenge = <code>2.5%</code> of real drawdown buffer).<br>• <strong>The 2-Strike Intraday Lock:</strong> Hard stop at 2 losing trades in a 24-hour window (max <code>-0.50%</code> intraday loss—10x safer than the 5% absorbing barrier).</p>`},
		{PageNum: 3, ThemeClass: "theme-phase3", RunningHead: "PAMPHLET · MODULES P3 & P4 (CME M6E & NEWS BLACKOUT)", FooterNote: "PAMPHLET PAGE 3 OF 8 · BM-03 & BM-04", BodyHTML: `<h2>MODULES P3 &amp; P4: REGULATED VENUE GATE &amp; T-COST FILTER (<code>[F-023, F-025]</code>)</h2><p><strong>Module P3 (<code>BM-03</code> — Regulated Exchange Gate):</strong> Trade regulated cash micro-lots (<code>0.01</code>) or CME Micro FX Futures (<code>M6E / M6B</code>, centralized clearing, transparent depth-of-book, IRC Section 1256 60/40 tax treatment) instead of offshore unlicensed B-book simulators.<br>• <strong>Module P4 (<code>BM-04</code> — 30-Minute News Blackout):</strong> Flatten all positions 15 minutes before and 15 minutes after Tier-1 macro releases (CPI, NFP, FOMC) to defeat asymmetric stop-loss slippage (<code>[K-015]</code>).</p>`},
		{PageNum: 4, ThemeClass: "theme-phase3", RunningHead: "PAMPHLET · MODULE P5: 90-DAY 100-TRADE LEDGER (PART I)", FooterNote: "PAMPHLET PAGE 4 OF 8 · BM-05", BodyHTML: fmt.Sprintf(`<h2>MODULE P5 (<code>BM-05</code>): THE 90-DAY ZERO-FEE PAPER LEDGER (TRADES 001–050)</h2><p>Log 100 consecutive trades on a free demo account over 90 days with realistic slippage before spending $1 (<code>[F-026]</code>):</p><div class="asset-frame"><img src="%s" alt="90 Day Test PosiBot"></div>`, p3PosiBot)},
		{PageNum: 5, ThemeClass: "theme-phase3", RunningHead: "PAMPHLET · MODULE P5: 90-DAY 100-TRADE LEDGER (PART II)", FooterNote: "PAMPHLET PAGE 5 OF 8 · BM-05 TRACKING GRID", BodyHTML: `<h2>MODULE P5 (CONT.): EXPECTANCY &amp; SHARPE TRACKING GRID (TRADES 051–100)</h2><table class="yit-table"><tr><th>TradeBatch</th><th>Win % (&ge;55%?)</th><th>Max Intraday DD (&lt;2.5%?)</th><th>Net R-Multiple (&gt;+0.25R?)</th><th>Hours Logged</th></tr><tr><td><code>#001–025</code></td><td>______%</td><td>______%</td><td>________ R</td><td>______ hrs</td></tr><tr><td><code>#026–050</code></td><td>______%</td><td>______%</td><td>________ R</td><td>______ hrs</td></tr><tr><td><code>#051–075</code></td><td>______%</td><td>______%</td><td>________ R</td><td>______ hrs</td></tr><tr><td><code>#076–100</code></td><td>______%</td><td>______%</td><td>________ R</td><td>______ hrs</td></tr></table>`},
		{PageNum: 6, ThemeClass: "theme-phase3", RunningHead: "PAMPHLET · MODULE P6: 120-DAY CHARGEBACK & FINE-PRINT KIT", FooterNote: "PAMPHLET PAGE 6 OF 8 · BM-06 & BM-07", BodyHTML: `<h2>MODULE P6 (<code>BM-06, BM-07</code>): THE 120-DAY CHARGEBACK &amp; CONTRACT AUDIT (<code>[F-021, K-015]</code>)</h2><p><strong>1. Pre-Purchase Fine-Print Red Flags (<code>BM-07</code>):</strong> Walk away immediately if the contract contains: (a) Equity-high intraday trailing drawdown, (b) Retroactive 30% single-trade consistency rules, (c) Mandatory minimum trading days combined with news blackouts, or (d) Unilateral B-book execution clauses.<br><strong>2. Visa Reason Code 13.1 / Mastercard 4853 Dispute Kit (<code>BM-06</code>):</strong> File within <strong>120 days</strong> of transaction date when a firm freezes withdrawals or revokes its MT5 license. Attach: (1) Checkout receipt, (2) SHA-256 archived Terms PDF, (3) Withdrawal denial screenshot.</p>`},
		{PageNum: 7, ThemeClass: "theme-phase3", RunningHead: "PAMPHLET · MODULE P7: 30-DAY SKILL TRANSLATION MATRIX", FooterNote: "PAMPHLET PAGE 7 OF 8 · BM-08 & P3-Q14", BodyHTML: `<h2>MODULE P7 (<code>BM-08</code>): 30-DAY QUANTITATIVE SKILL TRANSLATION MATRIX (<code>[F-015..F-018, F-024]</code>)</h2><table class="yit-table" style="font-size:7.6pt;"><tr><th>Retail Trading Habit</th><th>Corporate / Contract Equivalent</th><th>Audited BLS Pay</th></tr><tr><td>Coding MT5 / PineScript indicators</td><td><strong>Quant Backtest &amp; Execution Engineer</strong></td><td><code>$55.00/hr</code></td></tr><tr><td>Tracking FX macro &amp; rates calendars</td><td><strong>Corporate FX Treasury Risk Analyst</strong></td><td><code>$49.20/hr</code></td></tr><tr><td>Building equity-curve &amp; win-rate sheets</td><td><strong>Financial Data / SQL BI Analyst</strong></td><td><code>$44.50/hr</code></td></tr><tr><td>Explaining drawdown &amp; risk rules</td><td><strong>Series 65 Registered Paraplanner</strong></td><td><code>$38.90/hr</code></td></tr></table>`},
		{PageNum: 8, ThemeClass: "theme-phase3", RunningHead: "PAMPHLET · PERSONAL TRIPWIRE SCORECARD", FooterNote: "PAMPHLET PAGE 8 OF 8", BodyHTML: `<h2>PERSONAL STOP-RULE TRIPWIRES (<code>[F-026]</code>)</h2><div class="verdict-band">STAKE CAP: $500 · WAIT CAP: 6.0 MONTHS · PAYOFF FLOOR: $29.76/HR · ITISWHATITIS.BIZ</div>`},
	}
}

// BuildResearchMachineSpreadHTML builds a standalone 12.25" x 9.25" side-by-side landscape Dual-Page Spread of Pages 58-59 so the Research Machine visualization can also be inspected as an open two-page spread.
func BuildResearchMachineSpreadHTML(rootDir string) string {
	rmLeftSVG := fileURI(filepath.Join(rootDir, "05_phase3_run", "assets", "T01-P3-C08-diagram-research-machine-left_v001.svg"))
	rmRightSVG := fileURI(filepath.Join(rootDir, "05_phase3_run", "assets", "T01-P3-C08-diagram-research-machine-right_v001.svg"))

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Y-IT Dual-Page Research Machine Spread (Pages 58–59)</title>
<style>
  @page {
    size: 12.25in 9.25in;
    margin: 0;
  }
  * { box-sizing: border-box; -webkit-print-color-adjust: exact !important; print-color-adjust: exact !important; }
  html, body {
    margin: 0; padding: 0;
    width: 12.25in; height: 9.25in;
    background: #0F172A;
    font-family: 'Inter', 'Segoe UI', Arial, sans-serif;
    color: #F8FAFC;
  }
  .spread-canvas {
    width: 12.25in;
    height: 9.25in;
    display: grid;
    grid-template-columns: 6.125in 6.125in;
    position: relative;
  }
  .spread-page {
    padding: 0.32in 0.35in;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    background: #F1F5F9;
    color: #0F172A;
  }
  .spread-left {
    border-right: 2px dashed #475569;
  }
  .spread-right {
    border-left: 2px dashed #475569;
  }
  .badge {
    background: #0F172A;
    color: #38BDF8;
    font-family: 'Consolas', monospace;
    font-size: 7.5pt;
    font-weight: 800;
    padding: 3px 8px;
    display: inline-block;
    border-radius: 3px;
  }
  h2 {
    font-family: 'Impact', 'Big Shoulders Display', sans-serif;
    font-size: 14.5pt;
    margin: 4px 0;
    text-transform: uppercase;
    color: #0F172A;
  }
  p {
    font-size: 8.0pt;
    line-height: 1.32;
    margin: 2px 0 6px 0;
    color: #1E293B;
  }
  .svg-frame {
    background: #0F172A;
    border: 2px solid #0F172A;
    border-radius: 6px;
    padding: 4px;
    text-align: center;
  }
  .svg-frame img {
    width: 100%%;
    max-height: 6.65in;
    object-fit: contain;
    display: block;
  }
  .footer-bar {
    font-family: 'Consolas', monospace;
    font-size: 7.2pt;
    color: #334155;
    border-top: 1.5px solid #94A3B8;
    padding-top: 4px;
    display: flex;
    justify-content: space-between;
  }
</style>
</head>
<body>
<div class="spread-canvas">
  <div class="spread-page spread-left">
    <div>
      <div class="badge">APPENDIX A // EXHIBIT I · LEFT PLATE (VERSO PAGE 58) · SEQUENTIAL WITH BULK SOURCE LIST</div>
      <h2>HOW WE CHECKED THE RECEIPTS (PART I: 5-ROUND LIVE INGESTION &amp; SHA-256 LOCK)</h2>
      <p>Every Y-IT volume executes a multi-round empirical harvest across statutory federal dockets, central bank surveys, peer-reviewed DOI registries, temporal DOM archives, a <strong>12GB Dedicated Local VRAM + Elastic Cloud GPU Compute Array</strong>, and a <strong>Columnar OLAP Statistical Lakehouse</strong>:</p>
      <div class="svg-frame">
        <img src="%s" alt="Left Plate — Y-IT Research & Provenance Machine">
      </div>
    </div>
    <div class="footer-bar">
      <span>PAGE 58 (VERSO) · PRECEDED BY BULK SOURCE LIST PART I (PAGE 57)</span>
      <span>CONTINUES ACROSS CENTER GUTTER &rarr;</span>
    </div>
  </div>
  <div class="spread-page spread-right">
    <div>
      <div class="badge">APPENDIX A // EXHIBIT II · RIGHT PLATE (RECTO PAGE 59) · DETERMINISTIC COMPILER GATES</div>
      <h2>HOW WE CHECKED THE RECEIPTS (PART II: CONVERGENT RANGES, 14-MOLD SIEVE &amp; GATES)</h2>
      <p><strong>&larr; Continued from Page 58.</strong> Raw records pass through the <code>RM-10v2</code> Multi-Source Magnitude Convergence Sieve (requiring <code>&ge;3</code> independent sources of matching magnitude, integrated as a <strong>convergent range + cohort median</strong>), the <code>13+1</code> Empirical Failure Mold Census, and 7 Go Compiler Gates (<code>G1–G7</code>):</p>
      <div class="svg-frame">
        <img src="%s" alt="Right Plate — Convergent Range Sieve & Compiler Gates">
      </div>
    </div>
    <div class="footer-bar">
      <span>&larr; OPEN-FLAT DUAL-PAGE SPREAD</span>
      <span>PAGE 59 (RECTO) · FOLLOWED BY BULK SOURCE LIST PART II (PAGE 60)</span>
    </div>
  </div>
</div>
</body>
</html>`, rmLeftSVG, rmRightSVG)
}

// BuildFreebieCTAPage builds Page 14 of the 2-Chapter Freebie Giveaway PDF.
func BuildFreebieCTAPage() KDPPage {
	return KDPPage{
		PageNum:     14,
		ThemeClass:  "theme-phase2",
		RunningHead: "FREE PREVIEW COMPLETE · WHAT HAPPENED TO STEP 7?",
		FooterNote:  "GET THE FULL 64-PAGE FORENSIC AUTOPSY AT ITISWHATITIS.BIZ",
		BodyHTML: `
<div style="margin-top:0.6in; text-align:center; border:4px solid #DC2626; padding:20px; margin-left:-0.18in; margin-right:-0.18in; background:#FFFFFF; box-shadow:5px 6px 0px #991B1B;">
  <div class="phase-badge" style="background:#DC2626;">END OF FREE 2-CHAPTER PREVIEW</div>
  <h1 style="font-size:22pt; color:#0F172A; margin-top:10px;">YOU JUST READ THE BROCHURE.<br>NOW READ THE AUTOPSY.</h1>
  <p style="text-align:center; font-size:10pt; margin:12px 0;">
    In Chapters 1 and 2, we gave you the exact 6-step Funded Trader roadmap the gurus sell for <code>$997</code>—right up to the bottom of Page 13 where <strong>Step 7 was not provided</strong> (<code>[C-010]</code>).
  </p>
  <p style="text-align:center; font-size:9.8pt; font-weight:700; color:#991B1B; margin:12px 0;">
    In Chapters 3 through 8 of the full 64-page edition, we open the CFTC enforcement ledger ($310M in fees, 98.4% simulated B-book execution), examine 5 coroner case files on aged amber paper, and hand you 4 concrete pivots paying $38.90 to $55.00/hr.
  </p>
  <div class="verdict-band" style="text-align:center;">
    UNLOCK THE FULL 64-PAGE BOOK ($9.99) OR THE COMPLETE BOOK + 25-MIN PODCAP + PAMPHLET BUNDLE ($19.99) AT ITISWHATITIS.BIZ
  </div>
</div>`,
	}
}
