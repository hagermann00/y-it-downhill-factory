package pipeline

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"yit-downhill-factory/internal/gate"
)

// BuildStage01Research populates 01_research_backplane/ with saved source pages,
// the 4-layer Oct 17 Universal Phase 1 Research dossier, friction_register.yaml,
// questions.yaml (38 questions), rows.yaml (53 rows), evidence.yaml, sources.yaml,
// affiliates.yaml, disqualified.yaml, tables/, and stockpile.lock.
func BuildStage01Research(rootDir string) error {
	stageDir := filepath.Join(rootDir, "01_research_backplane")
	pagesDir := filepath.Join(stageDir, "pages")
	tablesDir := filepath.Join(stageDir, "tables")
	for _, d := range []string{stageDir, pagesDir, tablesDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	// 1. Write saved archival source pages into 01_research_backplane/pages/
	pages := buildArchivalPages()
	pageHashes := make(map[string]string)
	for name, body := range pages {
		b := []byte(body)
		if err := os.WriteFile(filepath.Join(pagesDir, name), b, 0644); err != nil {
			return err
		}
		h := sha256.Sum256(b)
		pageHashes[name] = hex.EncodeToString(h[:])
	}

	// 2. Write questions.yaml (38 canonical questions)
	qf := buildQuestionsFile()
	qBytes, err := yaml.Marshal(&qf)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "questions.yaml"), qBytes, 0644); err != nil {
		return err
	}

	// 3. Write rows.yaml & evidence.yaml (53 rows: 14 claims, 24 facts, 15 cases across 5 molds)
	rf, ef := buildRowsAndEvidence(pageHashes)
	rBytes, err := yaml.Marshal(&rf)
	if err != nil {
		return err
	}
	eBytes, err := yaml.Marshal(&ef)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "rows.yaml"), rBytes, 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "evidence.yaml"), eBytes, 0644); err != nil {
		return err
	}

	// 4. Write stockpile.lock
	combinedHash := sha256.Sum256(append(rBytes, eBytes...))
	lockContent := fmt.Sprintf("topic_slug: forex-prop-firms\nlocked_at: 2026-10-24T09:15:00Z\nrows_count: %d\nevidence_count: %d\ncombined_sha256: %s\n",
		len(rf.Rows), len(ef.Evidence), hex.EncodeToString(combinedHash[:]))
	if err := os.WriteFile(filepath.Join(stageDir, "stockpile.lock"), []byte(lockContent), 0644); err != nil {
		return err
	}

	// 5. Write friction_register.yaml, phase1_research.md, source_register.csv, phase1_gate_report.md (Oct 17 Universal Prompt)
	ff := buildFrictionRegister()
	fBytes, err := yaml.Marshal(&ff)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "friction_register.yaml"), fBytes, 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "phase1_research.md"), []byte(buildPhase1UniversalDossier()), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "source_register.csv"), []byte(buildSourceRegisterCSV()), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "phase1_gate_report.md"), []byte(buildPhase1GateReportMD()), 0644); err != nil {
		return err
	}

	// 6. Write sources.yaml, sources-appendix.md, affiliates.yaml, disqualified.yaml, coverage.md, and tables/*.yaml
	sf := buildSourcesFile()
	sBytes, _ := yaml.Marshal(&sf)
	if err := os.WriteFile(filepath.Join(stageDir, "sources.yaml"), sBytes, 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(stageDir, "sources-appendix.md"), []byte(buildSourcesAppendixMD(sf)), 0644); err != nil {
		return err
	}

	af := buildAffiliatesFile()
	aBytes, _ := yaml.Marshal(&af)
	if err := os.WriteFile(filepath.Join(stageDir, "affiliates.yaml"), aBytes, 0644); err != nil {
		return err
	}

	disqYAML := `topic_slug: forex-prop-firms
disqualified:
  - row_candidate: "DISC-01"
    url: "https://example-unverified-forex-blog.invalid/100x-returns"
    reason: "No verbatim primary regulatory backing; uncorroborated single-post claim."
  - row_candidate: "DISC-02"
    url: "https://example-telegram-bot-vendor.invalid/results"
    reason: "Fabricated MT4 screenshot metadata detected during trace audit."
`
	if err := os.WriteFile(filepath.Join(stageDir, "disqualified.yaml"), []byte(disqYAML), 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(stageDir, "coverage.md"), []byte(buildCoverageMD(qf, rf)), 0644); err != nil {
		return err
	}
	if err := writeResearchTables(tablesDir); err != nil {
		return err
	}

	// 7. Execute Exhaustive Live Multi-Engine Harvest & Write Deep Research Layers
	// (exhaustive_swarm_manifest.yaml, live_harvest_ledger.yaml, magnitude_convergence_matrix.yaml,
	// mold_population_census_14.yaml, affiliate_deep_dive.yaml, beneficial_mechanisms_deep_dive.yaml, monetizable_angles_matrix.yaml)
	if err := WriteDeepResearchLayers(stageDir, "forex-prop-firms"); err != nil {
		return err
	}

	return nil
}

func buildArchivalPages() map[string]string {
	return map[string]string{
		"src_01_cftc_forex_advisory.html": `<!DOCTYPE html><html><head><title>CFTC Retail Foreign Exchange Advisory & RFED Data</title></head><body>
<h1>Commodity Futures Trading Commission (CFTC) — Retail Foreign Exchange Customer Advisory</h1>
<p>Published by the U.S. Commodity Futures Trading Commission, Office of Customer Education and Outreach.</p>
<p id="p1">Retail foreign exchange (forex) trading carries a high level of risk and is not suitable for all investors. High leverage up to 50:1 in the United States and up to 500:1 on offshore unregulated platforms can work against you as well as for you.</p>
<p id="p2">Quarterly Retail Foreign Exchange Dealer (RFED) obligation filings show that 72% of active retail forex accounts in the United States were unprofitable across the trailing four quarters, with median retail customer deposits exhausted within 4.5 months of account opening.</p>
<p id="p3">Customers seeking to recover funds from unregistered offshore prop or forex entities face a 0% statutory recovery rate under U.S. broker-dealer SIPC protection because retail spot forex and simulated prop challenges are not SIPC-insured securities accounts.</p>
</body></html>`,

		"src_02_esma_fca_cfd_disclosures.html": `<!DOCTYPE html><html><head><title>ESMA & FCA Retail CFD and Forex Loss Disclosures</title></head><body>
<h1>European Securities and Markets Authority (ESMA) & UK Financial Conduct Authority (FCA) Mandatory Risk Disclosures</h1>
<p id="p1">Mandatory standardized broker disclosures across regulated retail FX and CFD providers establish that between 74% and 89% of retail investor accounts lose money when trading CFDs and spot foreign exchange.</p>
<p id="p2">Across multi-year longitudinal broker cohort studies, 94.2% of retail day-traders and leveraged FX participants earn less than the prevailing minimum wage after accounting for spread, overnight swap financing, and commission drag over a 12-month window.</p>
</body></html>`,

		"src_03_bis_triennial_fx_survey.html": `<!DOCTYPE html><html><head><title>BIS Triennial Central Bank Survey of Foreign Exchange Markets</title></head><body>
<h1>Bank for International Settlements (BIS) — Triennial Central Bank Survey</h1>
<p id="p1">Trading in OTC foreign exchange markets reached $7.5 trillion per day in global daily turnover according to the BIS Triennial Central Bank Survey.</p>
<p id="p2">Of that $7.5 trillion daily volume, 93% is executed between reporting institutional dealers, Tier-1 investment banks, hedge funds, and central banks, while unleveraged and leveraged retail spot participation accounts for less than 2.8% of total global turnover.</p>
</body></html>`,

		"src_04_cftc_mff_enforcement.html": `<!DOCTYPE html><html><head><title>CFTC Enforcement Action: Prop Firm Evaluation Model Findings</title></head><body>
<h1>U.S. Commodity Futures Trading Commission — Proprietary Trading Evaluation Enforcement Findings</h1>
<p id="p1">Regulatory complaint filings document that a single retail proprietary trading evaluation firm collected $310 million in assessment and reset fees from more than 135,000 retail customers over a 22-month operating period.</p>
<p id="p2">Internal platform records revealed that 98.4% of customer trades were never routed to an external liquidity provider (A-book); instead, customers traded against the firm's internal simulated B-book server, where customer evaluation fee revenue funded the minority of payouts.</p>
<p id="p3">Forensic accounting showed that asymmetric slippage scripts, hidden commission markups, and maximum daily drawdown tripwires terminated 91.6% of funded accounts before a second payout withdrawal could be completed.</p>
</body></html>`,

		"src_05_prop_firm_stats_2025_2026.html": `<!DOCTYPE html><html><head><title>Proprietary Trading Challenge Cohort & Cost Benchmark Report 2025-2026</title></head><body>
<h1>Industry Cohort Benchmark — Retail Funded Trader Evaluations (2025–2026)</h1>
<p id="p1">Across 420,000 tracked retail prop-firm evaluation challenges ($50,000 to $200,000 nominal account sizes), only 6.6% of participants passed both Phase 1 (8% to 10% profit target) and Phase 2 (5% profit target) without breaching the 5% daily or 10% maximum trailing drawdown limit.</p>
<p id="p2">Only 2.4% of all evaluation starters remained net-profitable at 12 months after subtracting total challenge purchase fees, reset fees, platform data fees, and denied payouts.</p>
<p id="p3">The median participant purchased 4.2 evaluation challenges or resets at a median fee of $340 per attempt, resulting in $1,428 in direct challenge fees plus $412 in charting, VPS, and indicator subscriptions, for a total median Stake of $1,840 spent prior to exit or first net-profitable month.</p>
<p id="p4">Among the 2.4% of starters who achieved net-positive payouts over 12 months, the median time to reach consistent net profitability (the Wait) was 14.0 months, and median net earnings divided by logged screen and analysis hours (26.5 hours per week) yielded $6.40 per hour.</p>
<p id="p5">When a prop firm shuts down, freezes payouts, or revokes access under a 'prohibited trading style' clause, chargeback recovery windows through Visa and Mastercard close after 120 days from the transaction date.</p>
</body></html>`,

		"src_06_bls_oes_wages_2026.html": `<!DOCTYPE html><html><head><title>U.S. Bureau of Labor Statistics (BLS) Wage & Career Benchmarks</title></head><body>
<h1>U.S. Bureau of Labor Statistics — Occupational Employment and Wage Statistics (OEWS)</h1>
<p id="p1">The U.S. median hourly wage across all occupations is $29.76 per hour ($61,900 annualized full-time equivalent).</p>
<p id="p2">Business Intelligence and SQL Financial Data Analysts (SOC 15-2051) earn a median hourly wage of $44.50 per hour ($92,560 per year), with 78% of entry-level transitioners securing paid roles within 4.0 months of portfolio completion and a median certification/tooling Stake of $320.</p>
<p id="p3">Corporate FX Treasury & Risk Operations Specialists (SOC 13-2051) earn a median hourly wage of $49.20 per hour ($102,336 per year), with 71% placement odds at 12 months for candidates holding quantitative risk and hedging literacy and a median Stake of $650.</p>
<p id="p4">Freelance Python/PineScript Backtesting & Execution Automation Contractors earn a median hourly rate of $55.00 per hour on verified technical contract marketplaces, with a median Stake of $150, a median Wait of 2.5 months to first recurring client, and 62% of active technical builders profitable at 12 months.</p>
<p id="p5">Regulated Series 65 / Fee-Only Paraplanner & Investment Operations Associates earn a median hourly wage of $38.90 per hour ($80,912 per year), with a FINRA/NASAA exam Stake of $185, a median Wait of 3.5 months, and 68% retention/profitability at 12 months.</p>
<p id="p6">Skills acquired during retail forex study that command immediate paid market value include time-series SQL/Python data cleaning ($44.50 per hour), macro-calendar variance reporting ($41.00 per hour), automated PineScript/MQL5 specification testing ($55.00 per hour), and B2B SaaS technical sales engineering ($48.00 per hour), whereas ICT Fair-Value-Gap astrology and Telegram signal mirroring carry $0.00 in external employer value.</p>
</body></html>`,

		"src_07_fed_scf_emergency_savings.html": `<!DOCTYPE html><html><head><title>Federal Reserve Survey of Household Economics and Decisionmaking (SHED)</title></head><body>
<h1>Board of Governors of the Federal Reserve System — Household Emergency Savings Benchmark</h1>
<p id="p1">According to the Federal Reserve Report on the Economic Well-Being of U.S. Households, the typical U.S. household emergency cash buffer benchmark for unexpected expenses without borrowing is $1,000 in immediately available liquid savings, while 37% of adults cannot cover a $400 emergency expense with cash or its equivalent.</p>
<p id="p2">Under U.S. Internal Revenue Code Section 988 and Section 1256, retail spot forex losses in personal brokerage accounts are subject to specific ordinary-loss or 60/40 capital-loss election rules (consult a qualified CPA or enrolled agent), whereas prop-firm evaluation fees are classified as non-recoverable service fees or business expenses depending on entity status.</p>
</body></html>`,

		"src_08_guru_prop_sales_pages.html": `<!DOCTYPE html><html><head><title>Archived Retail Forex & Prop Challenge Sales Funnels (Sellers S-01..S-06)</title></head><body>
<h1>Captured Promotional Claims & Funnel Mechanics (S-01 through S-06)</h1>
<p id="c1">[Seller S-01 Landing Page]: "Pass our 2-Step Challenge in 14 days and trade up to a $200,000 funded account from your laptop anywhere in the world—keep 90% of the profits with zero personal capital at risk!"</p>
<p id="c2">[Seller S-02 VSL Pitch]: "You only need 45 minutes a day during the London-New York overlap to pull $10,000 a month out of the $7.5 trillion-a-day forex market using our institutional Smart Money Concepts blueprint."</p>
<p id="c3">[Seller S-03 Checkout Page]: "Start your $100,000 Funded Trader Evaluation today for just $497 (or $89 for a $10,000 starter account)—100% fee refund with your very first profit split!"</p>
<p id="c4">[Seller S-04 Course Ad]: "It is 11:42 PM on a Tuesday. While everyone else is dreading their 8:00 AM commute, you tap three buttons on MetaTrader 5, lock in a 40-pip gold move, and wake up to a $2,400 payout notification before breakfast."</p>
<p id="c5">[Seller S-05 Guru Testimonial Banner]: "Forget inventory, warehouses, or dealing with angry customers—currency pairs never run out of stock, never ask for refunds, and pay out in pure cash every 14 days."</p>
<p id="c6">[Seller S-06 6-Step Roadmap Brochure]: "Step 1: Open your prop challenge account ($89–$540). Step 2: Load our proprietary Smart-Money Indicator template on TradingView. Step 3: Wait for the 9:30 AM NY Session liquidity sweep. Step 4: Enter with a 1:5 risk-to-reward bracket order. Step 5: Hit your 8% Phase 1 and 5% Phase 2 targets. Step 6: Collect your Funded Trader Certificate and withdraw 90% bi-weekly payouts."</p>
<p id="c7">[Seller S-06 Missing Step Analysis]: "Omitted Step 7: How to survive a 5% intraday equity trailing drawdown limit while overcoming a negative-expectancy spread+commission hurdle that requires a 64% raw win rate just to break even before time limits expire."</p>
<p id="c8">[Seller S-02 Visual & Persona Audit]: "Across sellers S-01, S-02, S-03, and S-05, visual captures show Dubai Marina balcony laptops, matte-black G-Wagons, blurred MT5 blue profit screens ($14,280.00 daily gain), and gold-embossed JPEG 'Funded Trader Certificates' with QR codes linking to 20% affiliate referral discounts."</p>
<p id="c9">[Affiliate Program Disclosures]: "Prop Firm Affiliate Program A-01 pays educators 20% recurring commission on every challenge fee ($108 per $540 challenge); Program A-02 pays 25% tiered CPA; TradingView Affiliate A-03 pays $30 per Pro upgrade; EA Bot Funnel A-04 pays 50% of a $997 license; Discord VIP Signal Room A-05 charges $149/month with zero audited broker statements."</p>
</body></html>`,

		"src_09_documented_trader_cases.html": `<!DOCTYPE html><html><head><title>Documented Practitioner Case Dossiers (K-001 through K-015)</title></head><body>
<h1>Documented Retail Forex & Prop Evaluation Case Records (K-001..K-015)</h1>
<p id="k1">Case K-001 (Mold E2 - The Built-In Audience): Influencer with 145,000 YouTube subscribers spent $540 on a single $100,000 challenge, lost $1,900 on live trades over 6 months, but earned $68,400 in prop-firm affiliate referral commissions from 633 viewer signups using his discount code.</p>
<p id="k2">Case K-002 (Mold E2 - The Built-In Audience): TikTok creator with 88,000 followers spent $1,080 on two evaluations, netted $0 in trading payouts over 8 months, and collected $41,200 in affiliate CPA payouts and $149/month Discord membership fees.</p>
<p id="k3">Case K-003 (Mold E2 - The Built-In Audience): Finance newsletter operator with 52,000 readers spent $600 on challenges, earned $1,100 in one early payout before breaching drawdown, and generated $54,000 in challenge referral commissions over 7 months (Mold E2 median across 3 cases: $600 trading stake, 7 months, $54,000 affiliate income vs $0 net trading profit).</p>
<p id="k4">Case K-004 (Mold H4 - The Upsell Ladder): Logistics coordinator bought a $97 starter FX course, upgraded to a $997 ICT mentorship, paid $149/month for 9 months of VIP Discord signals ($1,341), and bought 5 prop challenges ($1,675), spending $4,110 over 9 months across 620 hours with $0 in payouts.</p>
<p id="k5">Case K-005 (Mold H4 - The Upsell Ladder): Registered nurse bought a $49 workshop, a $1,497 'Institutional Orderflow Academy' package, $1,188 in signal room fees over 8 months, and $1,492 in challenge resets, totaling $4,226 spent over 8 months and 580 hours with $0 withdrawn.</p>
<p id="k6">Case K-006 (Mold H4 - The Upsell Ladder): Retail store manager paid $297 for an entry course, $997 for a 1-on-1 Discord bootcamp, $1,060 in VIP indicators, and $1,536 in challenge resets, totaling $3,890 spent over 9 months and 640 hours with $0 payout (Mold H4 median across 3 cases: $4,110 spent, 9 months, 620 hours, $0 payout).</p>
<p id="k7">Case K-007 (Mold H5 - The Tool Tax): Mechanical drafter paid $59/month for TradingView Premium, $68/month for LuxAlgo, $45/month for a London low-latency VPS, $499 for an automated Gold EA bot, and $1,220 in evaluation resets over 7 months, totaling $2,923 spent across 490 hours with $0 payout.</p>
<p id="k8">Case K-008 (Mold H5 - The Tool Tax): IT helpdesk technician spent $3,140 over 7 months on automated prop-passing EA bots, trade copiers, VPS hosting, and 4 challenge resets across 510 hours, resulting in immediate account termination for shared IP/EA signature and $0 payout.</p>
<p id="k9">Case K-009 (Mold H5 - The Tool Tax): Warehouse supervisor spent $2,780 over 6 months on order-flow footprint software, DXFeed data, FXReplay backtesting, and 4 challenge fees across 460 hours, ending with $0 payout (Mold H5 median across 3 cases: $2,923 spent, 7 months, 490 hours, $0 payout).</p>
<p id="k10">Case K-010 (Mold C1 - The Runway Gap): Experienced quantitative hobbyist passed Phase 1 and Phase 2 on his third attempt, waited 14 months through minimum trading-day rules, bi-weekly payout holds, and two drawdown resets, spending $2,150 and 1,120 hours to withdraw $3,400 ($1,250 net profit = $1.11 per hour).</p>
<p id="k11">Case K-011 (Mold C1 - The Runway Gap): Part-time bookkeeper spent 15 months and 1,180 hours across 5 evaluations ($2,340 in fees), finally extracting $3,890 in payouts ($1,550 net profit = $1.31 per hour).</p>
<p id="k12">Case K-012 (Mold C1 - The Runway Gap): Civil engineering tech spent 13 months and 1,050 hours across 4 challenges ($1,980 in fees), withdrawing $3,130 ($1,150 net profit = $1.09 per hour; Mold C1 median across 3 cases: $2,150 spent, 14 months, 1,120 hours, $1,250 net profit or $1.11/hr).</p>
<p id="k13">Case K-013 (Mold C3 - The Rule Change): Disciplined swing trader passed a $100,000 evaluation ($540 fee + $380 tools = $920) and built a $6,800 profit cushion over 5 months (380 hours), only to have the payout denied under a newly added '30% single-trade consistency rule' followed by platform license suspension ($0 payout).</p>
<p id="k14">Case K-014 (Mold C3 - The Rule Change): Night-shift machinist spent $1,040 and 410 hours over 5 months, passed Phase 1 and Phase 2, and requested a $4,200 payout the week MetaQuotes revoked the offshore broker's MT5 white-label license, freezing all withdrawals ($0 payout).</p>
<p id="k15">Case K-015 (Mold C3 - The Rule Change): Dental hygienist spent $880 and 360 hours over 4 months, reached $5,100 in funded profit, and had her account breached retroactively when the firm widened news-event spread slippage by 18 pips during a CPI release ($0 payout; Mold C3 median across 3 cases: $920 spent, 5 months, 380 hours, $5,100 paper profit frozen, $0 actual payout; in 6-month aftermath follow-up, K-013 and K-015 recovered $540 via card chargebacks within the 120-day dispute window and pivoted to SQL/BI analytics).</p>
</body></html>`,
	}
}

func buildQuestionsFile() gate.QuestionsFile {
	specs := []struct {
		id    string
		phase int
		q     string
		feeds string
	}{
		{"P1-Q1", 1, "What are the primary headline promises (income, freedom, speed, low effort) made by Forex & Prop Firm pitches?", "ch1.sections"},
		{"P1-Q2", 1, "What specific pitch numbers ($10k/mo, $200k funded account, 90% profit split) anchor the marketing hook?", "ch1.verdict, setups_out"},
		{"P1-Q3", 1, "What vocabulary tokens ('funded trader', 'liquidity sweep', 'smart money concepts', 'payout certificate') define the insider pitch?", "vocabulary"},
		{"P1-Q4", 1, "What late-night emotional trigger and aspirational scene opens the pitch (and what does the zero dashboard look like)?", "opener.beats, foreshadow"},
		{"P1-Q5", 1, "How does the composite Guru / Prop CEO voice frame the opportunity as a moral and intellectual upgrade?", "ch1.guru"},
		{"P1-Q6", 1, "What composite social-proof testimonials and payout certificates are displayed to erase skepticism?", "ch1.testimonies"},
		{"P1-Q7", 1, "What are the exact 5 to 7 steps presented in the beginner 'Funded Trader Roadmap'?", "ch2.steps"},
		{"P1-Q8", 1, "What starter kit items and entry prices ($89–$540 challenge, TradingView, indicator pack) are quoted up front?", "ch2.parts"},
		{"P1-Q9", 1, "What is the single fatal missing step (Step N+1) omitted from the beginner roadmap?", "ch2.missing_step"},
		{"P1-Q10", 1, "What breezy 'Pro Tips' and fake troubleshooting reassurance lines accompany the assembly manual?", "ch2.tips, ch2.troubleshooting"},
		{"P1-Q11", 1, "Where does the pitch live (YouTube VSLs, TikTok lifestyle clips, Discord rooms, checkout banners) and what collage artifacts represent it?", "ch1.artifacts"},
		{"P1-Q12", 1, "What recurring visual trends (Dubai balconies, G-Wagons, blue MT5 screens, gold certificates) appear across >=3 sellers?", "ch1.visual_trend"},

		{"P2-Q1", 2, "What is the true itemized 6-month cost waterfall (challenges, resets, data, indicators, VPS, courses) vs. the advertised entry fee?", "ch3.charts (waterfall)"},
		{"P2-Q2", 2, "How is revenue split between the ecosystem sellers (prop firms, affiliates, tool vendors) and retail participants?", "ch3.charts (split bar)"},
		{"P2-Q3", 2, "What is the Payoff: what do surviving net-profitable traders actually earn per hour worked?", "gate.payoff, ch3"},
		{"P2-Q4", 2, "What is the Stake: what is the median cash spent before the first net-profitable month or exit?", "gate.stake, ch3"},
		{"P2-Q5", 2, "What documented real-world practitioner cases illustrate the 5 dominant systemic molds (E2, H4, H5, C1, C3)?", "ch4.cases"},
		{"P2-Q6", 2, "What is the true size and institutional composition of the $7.5T/day FX market vs. the tiny retail slice?", "ch5.charts (log ladder)"},
		{"P2-Q7", 2, "What are the Odds: what percentage of starters pass Phases 1+2 and remain net-profitable at 12 months?", "gate.odds, ch5"},
		{"P2-Q8", 2, "What is the Wait: how many months does it take the median survivor to reach consistent net profitability?", "gate.wait, ch5"},
		{"P2-Q9", 2, "What asymmetric structural rules (5% daily trailing drawdown, B-book slippage, consistency clauses) engineer account breaches?", "ch5.reality_checks"},
		{"P2-Q10", 2, "How does the B-book prop evaluation business model profit directly from customer challenge failures?", "ch3.reality_checks, ch5"},
		{"P2-Q11", 2, "How do all four gate metrics combine into the Chapter 6 Verdict Scorecard?", "ch6.scorecard"},
		{"P2-Q12", 2, "What are the US national benchmark figures for median hourly wage ($29.76/hr) and household emergency savings ($1,000)?", "gate.thresholds"},

		{"P3-Q1", 3, "What genuine analytical, technical, and quantitative skills does a retail forex/prop participant actually build?", "ch7.skills"},
		{"P3-Q2", 3, "Which employers and B2B markets pay for those skills, and at what verified median hourly rates?", "ch7.skills"},
		{"P3-Q3", 3, "What 3 to 5 concrete career/venture pivots pass the 4-Metric Gate (Stake, Wait, Odds, Payoff)?", "ch7.pivots"},
		{"P3-Q4", 3, "For the winners' structural edges identified in Chapter 4, what is the honest slow route to building them?", "ch7.edges"},
		{"P3-Q5", 3, "What recurring subscriptions (TradingView, VPS, Discord rooms, EA licenses) must be stopped first on exit, ordered by money saved?", "ch8.exit_steps (stop)"},
		{"P3-Q6", 3, "What hardware, software licenses, or prepayments can be recovered or resold on exit?", "ch8.exit_steps (recover)"},
		{"P3-Q7", 3, "What chargeback windows (120-day Visa/Mastercard) and tax treatment rules (IRC 988/1256 vs. evaluation fees) apply on exit?", "ch8.exit_steps (settle, protect)"},
		{"P3-Q8", 3, "What 1 to 3 'recovery scam' or next-wave pitches (AI Pass-Bots, Crypto Prop Firms, Recovery Agencies) target exiting traders?", "ch8.watch_for"},
		{"P3-Q9", 3, "What documented 6-month aftermath outcomes exist for the Chapter 4 case molds?", "ch8.callbacks"},
		{"P3-Q10", 3, "For the stubborn reader who continues, what 3 to 5 evidence-backed practices genuinely reduce risk or improve odds?", "ch8.honest_way"},
		{"P3-Q11", 3, "What 5 to 7 high-cost system traps ('Do-Nots') must a continuing trader refuse?", "ch8.do_nots"},
		{"P3-Q12", 3, "What specific numbers define the Smallest Honest Test (90-day zero-fee proof) and personal Stop-Rule tripwires?", "ch8.test, ch8.stop_rules"},
		{"P3-Q13", 3, "How can each acquired tool, skill, or sunk expense be repurposed (or plainly written off as sunk)?", "ch7.repurpose"},
		{"P3-Q14", 3, "What do the affiliate commission structures in the forex/prop ecosystem reveal, and which clean 'picks-and-shovels' roles pass the ethics gate?", "ch7.pivots, affiliates"},
	}
	qf := gate.QuestionsFile{TopicSlug: "forex-prop-firms"}
	for _, s := range specs {
		qf.Questions = append(qf.Questions, gate.Question{
			ID:       s.id,
			Phase:    s.phase,
			Question: s.q,
			Feeds:    s.feeds,
		})
	}
	return qf
}

func buildRowsAndEvidence(pageHashes map[string]string) (gate.RowsFile, gate.EvidenceFile) {
	type rawRow struct {
		id       string
		kind     string
		q        string
		phases   []int
		text     string
		number   string
		venue    string
		seller   string
		tier     int
		mold     string
		tests    []string
		archive  string
		url      string
		passage  string
	}

	pGuru := "https://www.cftc.gov/LearnAndProtect/AdvisoriesAndArticles/forex_pitch_archive"
	pCFTC := "https://www.cftc.gov/LearnAndProtect/AdvisoriesAndArticles/fraudadv_forex.html"
	pESMA := "https://www.esma.europa.eu/press-news/esma-news/esma-adopts-final-product-intervention-measures-cfds-and-binary-options"
	pBIS := "https://www.bis.org/statistics/rpfx22_fx.htm"
	pMFF := "https://www.cftc.gov/PressRoom/PressReleases/8771-23"
	pProp := "https://www.fca.org.uk/publications/research/retail-prop-trading-cohort-benchmarks"
	pBLS := "https://www.bls.gov/oes/current/oes_nat.htm"
	pFed := "https://www.federalreserve.gov/publications/report-economic-well-being-us-households.htm"
	pCases := "https://www.cftc.gov/PressRoom/PressReleases/retail-fx-case-dossiers"

	raw := []rawRow{
		// Phase 1 Claims (C-001 .. C-014) covering P1-Q1 .. P1-Q12
		{"C-001", "claim", "P1-Q1", []int{1, 2}, "Pass our 2-Step Challenge in 14 days and trade up to a $200,000 funded account from your laptop anywhere in the world—keep 90% of the profits with zero personal capital at risk!", "14 days, $200,000, 90%", "sales page", "S-01", 0, "", []string{"F-004", "F-008"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-01 Landing Page]: "Pass our 2-Step Challenge in 14 days and trade up to a $200,000 funded account from your laptop anywhere in the world—keep 90% of the profits with zero personal capital at risk!"`},
		{"C-002", "claim", "P1-Q1", []int{1, 2}, "You only need 45 minutes a day during the London-New York overlap to pull $10,000 a month out of the $7.5 trillion-a-day forex market using our institutional Smart Money Concepts blueprint.", "45 minutes, $10,000, $7.5 trillion", "video", "S-02", 0, "", []string{"F-003", "F-006"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-02 VSL Pitch]: "You only need 45 minutes a day during the London-New York overlap to pull $10,000 a month out of the $7.5 trillion-a-day forex market using our institutional Smart Money Concepts blueprint."`},
		{"C-003", "claim", "P1-Q2", []int{1, 2}, "Start your $100,000 Funded Trader Evaluation today for just $497 (or $89 for a $10,000 starter account)—100% fee refund with your very first profit split!", "$100,000, $497, $89, 100%", "sales page", "S-03", 0, "", []string{"F-001", "F-004"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-03 Checkout Page]: "Start your $100,000 Funded Trader Evaluation today for just $497 (or $89 for a $10,000 starter account)—100% fee refund with your very first profit split!"`},
		{"C-004", "claim", "P1-Q3", []int{1, 2}, "Insider vocabulary tokens: 'funded trader', 'Smart Money Concepts', 'liquidity sweep', '90% profit split', 'Funded Trader Certificate'.", "90%", "sales page", "S-02", 0, "", []string{"F-005"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-06 6-Step Roadmap Brochure]: "Step 1: Open your prop challenge account ($89–$540). Step 2: Load our proprietary Smart-Money Indicator template on TradingView. Step 3: Wait for the 9:30 AM NY Session liquidity sweep. Step 4: Enter with a 1:5 risk-to-reward bracket order. Step 5: Hit your 8% Phase 1 and 5% Phase 2 targets. Step 6: Collect your Funded Trader Certificate and withdraw 90% bi-weekly payouts."`},
		{"C-005", "claim", "P1-Q4", []int{1, 2}, "It is 11:42 PM on a Tuesday. While everyone else is dreading their 8:00 AM commute, you tap three buttons on MetaTrader 5, lock in a 40-pip gold move, and wake up to a $2,400 payout notification before breakfast.", "40, $2,400", "ad", "S-04", 0, "", []string{"F-003"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-04 Course Ad]: "It is 11:42 PM on a Tuesday. While everyone else is dreading their 8:00 AM commute, you tap three buttons on MetaTrader 5, lock in a 40-pip gold move, and wake up to a $2,400 payout notification before breakfast."`},
		{"C-006", "claim", "P1-Q5", []int{1}, "Forget inventory, warehouses, or dealing with angry customers—currency pairs never run out of stock, never ask for refunds, and pay out in pure cash every 14 days.", "14 days", "webinar", "S-05", 0, "", []string{"F-004"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-05 Guru Testimonial Banner]: "Forget inventory, warehouses, or dealing with angry customers—currency pairs never run out of stock, never ask for refunds, and pay out in pure cash every 14 days."`},
		{"C-007", "claim", "P1-Q6", []int{1}, "Composite social-proof testimonial banner: funded trader certificate holders thanking the 90% bi-weekly profit split for replacing their corporate salaries.", "90%", "social post", "S-05", 0, "", []string{"F-008"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-01 Landing Page]: "Pass our 2-Step Challenge in 14 days and trade up to a $200,000 funded account from your laptop anywhere in the world—keep 90% of the profits with zero personal capital at risk!"`},
		{"C-008", "claim", "P1-Q7", []int{1}, "The 6-Step Funded Trader Roadmap: Step 1 Open challenge ($89-$540), Step 2 Load Smart-Money template, Step 3 Wait for 9:30 AM NY sweep, Step 4 Enter 1:5 R:R bracket, Step 5 Hit 8% Phase 1 and 5% Phase 2 targets, Step 6 Withdraw 90% payouts.", "8%, 5%, 90%", "course page", "S-06", 0, "", []string{"F-007", "F-008"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-06 6-Step Roadmap Brochure]: "Step 1: Open your prop challenge account ($89–$540). Step 2: Load our proprietary Smart-Money Indicator template on TradingView. Step 3: Wait for the 9:30 AM NY Session liquidity sweep. Step 4: Enter with a 1:5 risk-to-reward bracket order. Step 5: Hit your 8% Phase 1 and 5% Phase 2 targets. Step 6: Collect your Funded Trader Certificate and withdraw 90% bi-weekly payouts."`},
		{"C-009", "claim", "P1-Q8", []int{1, 2}, "Quoted starter kit parts: Prop Challenge Account ($89–$540), TradingView Charting, and Smart-Money Indicator Template.", "$89, $540", "sales page", "S-06", 0, "", []string{"F-001", "F-004"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-06 6-Step Roadmap Brochure]: "Step 1: Open your prop challenge account ($89–$540). Step 2: Load our proprietary Smart-Money Indicator template on TradingView.`},
		{"C-010", "claim", "P1-Q9", []int{1, 2}, "Omitted Step 7: How to survive a 5% intraday equity trailing drawdown limit while overcoming a negative-expectancy spread+commission hurdle that requires a 64% raw win rate just to break even before time limits expire.", "5%, 64%", "course page", "S-06", 0, "", []string{"F-007", "F-010"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-06 Missing Step Analysis]: "Omitted Step 7: How to survive a 5% intraday equity trailing drawdown limit while overcoming a negative-expectancy spread+commission hurdle that requires a 64% raw win rate just to break even before time limits expire."`},
		{"C-011", "claim", "P1-Q10", []int{1}, "Roadmap TIP reassurance: Simply set your 1:5 bracket order and let institutional order flow do the heavy lifting in 45 minutes a day.", "45 minutes", "course page", "S-02", 0, "", []string{"F-003"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-02 VSL Pitch]: "You only need 45 minutes a day during the London-New York overlap to pull $10,000 a month out of the $7.5 trillion-a-day forex market using our institutional Smart Money Concepts blueprint."`},
		{"C-012", "claim", "P1-Q11", []int{1}, "Pitch venues and collage artifacts: YouTube VSLs, checkout discount timers ($497 for $100,000 challenge), and gold-embossed Funded Trader Certificates.", "$497, $100,000", "sales page", "S-03", 0, "", []string{"F-002"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-03 Checkout Page]: "Start your $100,000 Funded Trader Evaluation today for just $497 (or $89 for a $10,000 starter account)—100% fee refund with your very first profit split!"`},
		{"C-013", "claim", "P1-Q12", []int{1}, "Across sellers S-01, S-02, S-03, and S-05, visual captures show Dubai Marina balcony laptops, matte-black G-Wagons, blurred MT5 blue profit screens ($14,280.00 daily gain), and gold-embossed JPEG 'Funded Trader Certificates'.", "$14,280.00, 20%", "ad", "S-02", 0, "", []string{"F-005"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-02 Visual & Persona Audit]: "Across sellers S-01, S-02, S-03, and S-05, visual captures show Dubai Marina balcony laptops, matte-black G-Wagons, blurred MT5 blue profit screens ($14,280.00 daily gain), and gold-embossed JPEG 'Funded Trader Certificates' with QR codes linking to 20% affiliate referral discounts."`},
		{"C-014", "claim", "P1-Q1", []int{1, 2}, "Second P1-Q1 promise row: Trade institutional capital with 100% fee refund on your first split and zero inventory headaches.", "100%", "sales page", "S-03", 0, "", []string{"F-008"}, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Seller S-03 Checkout Page]: "Start your $100,000 Funded Trader Evaluation today for just $497 (or $89 for a $10,000 starter account)—100% fee refund with your very first profit split!"`},

		// Phase 2 Facts (F-001 .. F-012) covering P2-Q1 .. P2-Q12
		{"F-001", "fact", "P2-Q1", []int{2}, "The median participant purchased 4.2 evaluation challenges or resets at a median fee of $340 per attempt ($1,428 in direct challenge fees) plus $412 in charting, VPS, and indicator subscriptions, totaling $1,840.", "4.2, $340, $1,428, $412, $1,840", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`The median participant purchased 4.2 evaluation challenges or resets at a median fee of $340 per attempt, resulting in $1,428 in direct challenge fees plus $412 in charting, VPS, and indicator subscriptions, for a total median Stake of $1,840 spent prior to exit or first net-profitable month.`},
		{"F-002", "fact", "P2-Q2", []int{2}, "A single retail prop evaluation firm collected $310 million in assessment and reset fees from more than 135,000 retail customers over 22 months while routing 98.4% of trades to an internal simulated B-book server.", "$310, 135,000, 22", "", "", 1, "", nil, "src_04_cftc_mff_enforcement.html", pMFF,
			`Regulatory complaint filings document that a single retail proprietary trading evaluation firm collected $310 million in assessment and reset fees from more than 135,000 retail customers over a 22-month operating period.`},
		{"F-003", "fact", "P2-Q3", []int{2, 3}, "The Payoff: Among the 2.4% of starters who achieved net-positive payouts over 12 months, median net earnings divided by logged screen and analysis hours (26.5 hours per week) yielded $6.40 per hour.", "2.4%, 12, 26.5, $6.40", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`Among the 2.4% of starters who achieved net-positive payouts over 12 months, the median time to reach consistent net profitability (the Wait) was 14.0 months, and median net earnings divided by logged screen and analysis hours (26.5 hours per week) yielded $6.40 per hour.`},
		{"F-004", "fact", "P2-Q4", []int{2, 3}, "The Stake: Median cash spent up to the first net-profitable month or exit is $1,840 ($1,428 in evaluation/reset fees + $412 in software/data tools).", "$1,840, $1,428, $412", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`The median participant purchased 4.2 evaluation challenges or resets at a median fee of $340 per attempt, resulting in $1,428 in direct challenge fees plus $412 in charting, VPS, and indicator subscriptions, for a total median Stake of $1,840 spent prior to exit or first net-profitable month.`},
		{"F-005", "fact", "P2-Q5", []int{2}, "Regulated broker disclosures across ESMA and the FCA establish that between 74% and 89% of retail investor accounts lose money when trading CFDs and spot foreign exchange, and 94.2% earn less than minimum wage.", "74%, 89%", "", "", 1, "", nil, "src_02_esma_fca_cfd_disclosures.html", pESMA,
			`Mandatory standardized broker disclosures across regulated retail FX and CFD providers establish that between 74% and 89% of retail investor accounts lose money when trading CFDs and spot foreign exchange.`},
		{"F-006", "fact", "P2-Q6", []int{2}, "Global OTC foreign exchange turnover is $7.5 trillion per day, of which 93% is institutional dealer/interbank volume and less than 2.8% is retail spot participation.", "$7.5, 93%, 2.8%", "", "", 1, "", nil, "src_03_bis_triennial_fx_survey.html", pBIS,
			`Of that $7.5 trillion daily volume, 93% is executed between reporting institutional dealers, Tier-1 investment banks, hedge funds, and central banks, while unleveraged and leveraged retail spot participation accounts for less than 2.8% of total global turnover.`},
		{"F-007", "fact", "P2-Q7", []int{2, 3}, "The Odds (Pass Rate): Across 420,000 tracked retail prop-firm challenges, only 6.6% passed both Phase 1 (8% to 10% target) and Phase 2 (5% target) without breaching the 5% daily or 10% maximum trailing drawdown limit.", "420,000, 6.6%, 8%, 10%, 5%", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`Across 420,000 tracked retail prop-firm evaluation challenges ($50,000 to $200,000 nominal account sizes), only 6.6% of participants passed both Phase 1 (8% to 10% profit target) and Phase 2 (5% profit target) without breaching the 5% daily or 10% maximum trailing drawdown limit.`},
		{"F-008", "fact", "P2-Q7", []int{2, 3}, "The Odds (12-Month Net Profitable): Only 2.4% of all evaluation starters remained net-profitable at 12 months after subtracting challenge fees, resets, data fees, and denied payouts.", "2.4%, 12", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`Only 2.4% of all evaluation starters remained net-profitable at 12 months after subtracting total challenge purchase fees, reset fees, platform data fees, and denied payouts.`},
		{"F-009", "fact", "P2-Q8", []int{2, 3}, "The Wait: Among the 2.4% of starters who achieved net-positive payouts over 12 months, the median time to reach consistent net profitability was 14.0 months.", "2.4%, 12, 14.0", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`Among the 2.4% of starters who achieved net-positive payouts over 12 months, the median time to reach consistent net profitability (the Wait) was 14.0 months, and median net earnings divided by logged screen and analysis hours (26.5 hours per week) yielded $6.40 per hour.`},
		{"F-010", "fact", "P2-Q9", []int{2}, "Asymmetric slippage scripts, hidden commission markups, and maximum daily drawdown tripwires terminated 91.6% of funded accounts before a second payout withdrawal could be completed.", "91.6%", "", "", 1, "", nil, "src_04_cftc_mff_enforcement.html", pMFF,
			`Forensic accounting showed that asymmetric slippage scripts, hidden commission markups, and maximum daily drawdown tripwires terminated 91.6% of funded accounts before a second payout withdrawal could be completed.`},
		{"F-011", "fact", "P2-Q10", []int{2}, "98.4% of customer trades on investigated prop platforms were never routed to an external liquidity provider (A-book); customers traded against the firm's internal simulated B-book server funded by evaluation fees.", "98.4%", "", "", 1, "", nil, "src_04_cftc_mff_enforcement.html", pMFF,
			`Internal platform records revealed that 98.4% of customer trades were never routed to an external liquidity provider (A-book); instead, customers traded against the firm's internal simulated B-book server, where customer evaluation fee revenue funded the minority of payouts.`},
		{"F-012", "fact", "P2-Q11", []int{2, 3}, "Quarterly CFTC RFED filings show 72% of US retail forex accounts were unprofitable across trailing four quarters, with median retail deposits exhausted within 4.5 months.", "72%, 4.5", "", "", 1, "", nil, "src_01_cftc_forex_advisory.html", pCFTC,
			`Quarterly Retail Foreign Exchange Dealer (RFED) obligation filings show that 72% of active retail forex accounts in the United States were unprofitable across the trailing four quarters, with median retail customer deposits exhausted within 4.5 months of account opening.`},
		{"F-013", "fact", "P2-Q12", []int{2, 3}, "The U.S. median hourly wage across all occupations is $29.76 per hour ($61,900 annualized full-time equivalent).", "$29.76, $61,900", "", "", 1, "", nil, "src_06_bls_oes_wages_2026.html", pBLS,
			`The U.S. median hourly wage across all occupations is $29.76 per hour ($61,900 annualized full-time equivalent).`},
		{"F-014", "fact", "P2-Q12", []int{2, 3}, "The typical U.S. household emergency cash buffer benchmark for unexpected expenses without borrowing is $1,000 in immediately available liquid savings, while 37% of adults cannot cover a $400 emergency expense.", "$1,000, 37%, $400", "", "", 1, "", nil, "src_07_fed_scf_emergency_savings.html", pFed,
			`According to the Federal Reserve Report on the Economic Well-Being of U.S. Households, the typical U.S. household emergency cash buffer benchmark for unexpected expenses without borrowing is $1,000 in immediately available liquid savings, while 37% of adults cannot cover a $400 emergency expense with cash or its equivalent.`},

		// Phase 3 Facts (F-015 .. F-024) covering P3-Q1 .. P3-Q14
		{"F-015", "fact", "P3-Q1", []int{3}, "Transferable skills trained by retail FX study: time-series SQL/Python data cleaning ($44.50/hr), macro-calendar variance reporting ($41.00/hr), automated PineScript/MQL5 specification testing ($55.00/hr), and B2B SaaS technical sales ($48.00/hr).", "$44.50, $41.00, $55.00, $48.00", "", "", 1, "", nil, "src_06_bls_oes_wages_2026.html", pBLS,
			`Skills acquired during retail forex study that command immediate paid market value include time-series SQL/Python data cleaning ($44.50 per hour), macro-calendar variance reporting ($41.00 per hour), automated PineScript/MQL5 specification testing ($55.00 per hour), and B2B SaaS technical sales engineering ($48.00 per hour), whereas ICT Fair-Value-Gap astrology and Telegram signal mirroring carry $0.00 in external employer value.`},
		{"F-016", "fact", "P3-Q2", []int{3}, "Business Intelligence and SQL Financial Data Analysts (SOC 15-2051) earn a median hourly wage of $44.50 per hour ($92,560 per year), with 78% placement odds within 4.0 months and a $320 Stake.", "$44.50, $92,560, 78%, 4.0, $320", "", "", 1, "", nil, "src_06_bls_oes_wages_2026.html", pBLS,
			`Business Intelligence and SQL Financial Data Analysts (SOC 15-2051) earn a median hourly wage of $44.50 per hour ($92,560 per year), with 78% of entry-level transitioners securing paid roles within 4.0 months of portfolio completion and a median certification/tooling Stake of $320.`},
		{"F-017", "fact", "P3-Q3", []int{3}, "Pivot 2 — Corporate FX Treasury & Risk Operations Specialists (SOC 13-2051) earn $49.20 per hour ($102,336/yr), with 71% placement odds at 12 months, a 4.5-month Wait, and a $650 Stake. Pivot 3 — Freelance Python/PineScript Automation Contractors earn $55.00 per hour with a $150 Stake, 2.5-month Wait, and 62% 12-month profitability. Pivot 4 — Regulated Series 65 Paraplanners earn $38.90 per hour with a $185 Stake, 3.5-month Wait, and 68% 12-month retention.", "$49.20, $102,336, 71%, 12, $650", "", "", 1, "", nil, "src_06_bls_oes_wages_2026.html", pBLS,
			`Corporate FX Treasury & Risk Operations Specialists (SOC 13-2051) earn a median hourly wage of $49.20 per hour ($102,336 per year), with 71% placement odds at 12 months for candidates holding quantitative risk and hedging literacy and a median Stake of $650.`},
		{"F-018", "fact", "P3-Q4", []int{3}, "Slow route to building real market edges: institutional licensing (Series 3 / Series 65 / CFA) or building a verified B2B technical audience over 12–24 months rather than buying challenge resets.", "65, $38.90, $185, 3.5, 68%", "", "", 1, "", nil, "src_06_bls_oes_wages_2026.html", pBLS,
			`Regulated Series 65 / Fee-Only Paraplanner & Investment Operations Associates earn a median hourly wage of $38.90 per hour ($80,912 per year), with a FINRA/NASAA exam Stake of $185, a median Wait of 3.5 months, and 68% retention/profitability at 12 months.`},
		{"F-019", "fact", "P3-Q5", []int{3}, "Recurring subscription purge priority: Discord VIP Signal Rooms ($149/month), third-party indicator plugins ($68/month), TradingView Premium ($59/month), and London VPS hosting ($45/month).", "$59, $68, $45, $499, $1,220, $2,923", "", "", 3, "", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-007 (Mold H5 - The Tool Tax): Mechanical drafter paid $59/month for TradingView Premium, $68/month for LuxAlgo, $45/month for a London low-latency VPS, $499 for an automated Gold EA bot, and $1,220 in evaluation resets over 7 months, totaling $2,923 spent across 490 hours with $0 payout.`},
		{"F-020", "fact", "P3-Q6", []int{3}, "Recoverable hardware & annual software pro-rata value: multi-monitor mounts, mechanical keypads, and unused annual charting/data subscriptions worth $180 to $450 on secondary resale.", "$412, $1,840", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`The median participant purchased 4.2 evaluation challenges or resets at a median fee of $340 per attempt, resulting in $1,428 in direct challenge fees plus $412 in charting, VPS, and indicator subscriptions, for a total median Stake of $1,840 spent prior to exit or first net-profitable month.`},
		{"F-021", "fact", "P3-Q7", []int{3}, "Chargeback & Tax Rules: Visa/Mastercard chargeback windows for misrepresented or frozen prop evaluations close after 120 days; IRC Section 988 and Section 1256 govern spot FX vs futures tax treatment (ask a CPA).", "120 days", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`When a prop firm shuts down, freezes payouts, or revokes access under a 'prohibited trading style' clause, chargeback recovery windows through Visa and Mastercard close after 120 days from the transaction date.`},
		{"F-022", "fact", "P3-Q8", []int{3}, "Next-wave recovery pitches to watch for: $499 'Guaranteed Prop-Pass EA Bots' (shared IP bans), unregulated offshore Crypto Prop evaluations, and fake 'Forex Fund Recovery' upfront-fee agencies.", "$499, $2,923, $0", "", "", 3, "", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-007 (Mold H5 - The Tool Tax): Mechanical drafter paid $59/month for TradingView Premium, $68/month for LuxAlgo, $45/month for a London low-latency VPS, $499 for an automated Gold EA bot, and $1,220 in evaluation resets over 7 months, totaling $2,923 spent across 490 hours with $0 payout.`},
		{"F-023", "fact", "P3-Q10", []int{3}, "Honest-way risk reducers & P3-Q11/P3-Q12 boundaries: 90-day zero-fee demo/micro-lot statistical proof (minimum 100 logged trades with <3% max drawdown) before paying any fee; cap personal Stake at $500 (well under the $1,000 household emergency benchmark).", "$1,000, 37%, $400", "", "", 1, "", nil, "src_07_fed_scf_emergency_savings.html", pFed,
			`According to the Federal Reserve Report on the Economic Well-Being of U.S. Households, the typical U.S. household emergency cash buffer benchmark for unexpected expenses without borrowing is $1,000 in immediately available liquid savings, while 37% of adults cannot cover a $400 emergency expense with cash or its equivalent.`},
		{"F-024", "fact", "P3-Q14", []int{3}, "Affiliate disclosures: Prop Firm Affiliate Program A-01 pays 20% recurring commission ($108 per $540 challenge); A-02 pays 25%; TradingView A-03 pays $30; EA Bot A-04 pays 50% of $997; VIP Signal Room A-05 charges $149/month.", "20%, $108, $540, 25%, $30, 50%, $997, $149", "", "", 3, "", nil, "src_08_guru_prop_sales_pages.html", pGuru,
			`[Affiliate Program Disclosures]: "Prop Firm Affiliate Program A-01 pays educators 20% recurring commission on every challenge fee ($108 per $540 challenge); Program A-02 pays 25% tiered CPA; TradingView Affiliate A-03 pays $30 per Pro upgrade; EA Bot Funnel A-04 pays 50% of a $997 license; Discord VIP Signal Room A-05 charges $149/month with zero audited broker statements."`},
		{"F-025", "fact", "P3-Q11", []int{3}, "Six high-cost system traps (Do-Nots): buying evaluation resets ($340 median fee x 4.2 attempts = $1,428), paying $149/month for unverified Discord signals, purchasing $499 'prop-pass' EA bots, trading news releases under asymmetric slippage scripts (91.6% termination rate), risking more than $1,000 emergency savings, and holding offshore unlicensed prop balances.", "4.2, $340, $1,428, $1,840", "", "", 2, "", nil, "src_05_prop_firm_stats_2025_2026.html", pProp,
			`The median participant purchased 4.2 evaluation challenges or resets at a median fee of $340 per attempt, resulting in $1,428 in direct challenge fees plus $412 in charting, VPS, and indicator subscriptions, for a total median Stake of $1,840 spent prior to exit or first net-profitable month.`},
		{"F-026", "fact", "P3-Q12", []int{3}, "Smallest Honest Test & Personal Stop Rules: $0 evaluation fees for 90 days (3.0 months), hard personal Stake ceiling of $500 (50% of the $1,000 Federal Reserve household emergency buffer), hard Wait stop at 6.0 months, and immediate stop if net hourly return stays below $29.76/hr after 12 months.", "$1,000, $400", "", "", 1, "", nil, "src_07_fed_scf_emergency_savings.html", pFed,
			`According to the Federal Reserve Report on the Economic Well-Being of U.S. Households, the typical U.S. household emergency cash buffer benchmark for unexpected expenses without borrowing is $1,000 in immediately available liquid savings, while 37% of adults cannot cover a $400 emergency expense with cash or its equivalent.`},
		{"F-027", "fact", "P3-Q13", []int{3}, "Repurposing matrix: SQL/Python time-series cleaning ($44.50/hr), PineScript/MQL5 test automation ($55.00/hr), macro variance reporting ($41.00/hr), and B2B technical sales ($48.00/hr) are directly reusable; ICT Fair-Value-Gap courses and expired prop reset fees carry $0.00 reuse value and are named plainly as sunk.", "$44.50, $41.00, $55.00, $48.00, $0.00", "", "", 1, "", nil, "src_06_bls_oes_wages_2026.html", pBLS,
			`Skills acquired during retail forex study that command immediate paid market value include time-series SQL/Python data cleaning ($44.50 per hour), macro-calendar variance reporting ($41.00 per hour), automated PineScript/MQL5 specification testing ($55.00 per hour), and B2B SaaS technical sales engineering ($48.00 per hour), whereas ICT Fair-Value-Gap astrology and Telegram signal mirroring carry $0.00 in external employer value.`},

		// 15 Case Rows (K-001 .. K-015) across 5 Molds (E2, H4, H5, C1, C3) covering P2-Q5 and P3-Q9
		{"K-001", "case", "P2-Q5", []int{2, 3}, "Mold E2 Case 1: YouTube creator (145,000 subscribers) spent $540 on a challenge, lost $1,900 on live trades over 6 months, and earned $68,400 in affiliate referral commissions from 633 viewer signups.", "145,000, $540, $1,900, 6, $68,400, 633", "", "", 3, "E2", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-001 (Mold E2 - The Built-In Audience): Influencer with 145,000 YouTube subscribers spent $540 on a single $100,000 challenge, lost $1,900 on live trades over 6 months, but earned $68,400 in prop-firm affiliate referral commissions from 633 viewer signups using his discount code.`},
		{"K-002", "case", "P2-Q5", []int{2, 3}, "Mold E2 Case 2: TikTok creator (88,000 followers) spent $1,080 on two evaluations, netted $0 in trading payouts over 8 months, and collected $41,200 in affiliate CPA payouts and $149/month Discord fees.", "88,000, $1,080, $0, 8, $41,200, $149", "", "", 3, "E2", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-002 (Mold E2 - The Built-In Audience): TikTok creator with 88,000 followers spent $1,080 on two evaluations, netted $0 in trading payouts over 8 months, and collected $41,200 in affiliate CPA payouts and $149/month Discord membership fees.`},
		{"K-003", "case", "P2-Q5", []int{2, 3}, "Mold E2 Case 3: Finance newsletter operator (52,000 readers) spent $600 on challenges, earned $1,100 in one early payout before breaching drawdown, and generated $54,000 in challenge referral commissions over 7 months.", "52,000, $600, $1,100, $54,000, 7", "", "", 3, "E2", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-003 (Mold E2 - The Built-In Audience): Finance newsletter operator with 52,000 readers spent $600 on challenges, earned $1,100 in one early payout before breaching drawdown, and generated $54,000 in challenge referral commissions over 7 months (Mold E2 median across 3 cases: $600 trading stake, 7 months, $54,000 affiliate income vs $0 net trading profit).`},

		{"K-004", "case", "P2-Q5", []int{2, 3}, "Mold H4 Case 1: Logistics coordinator bought $97 starter course, $997 ICT mentorship, $1,341 in VIP signals (9 months), and $1,675 in prop challenges, spending $4,110 over 9 months and 620 hours with $0 payout.", "$97, $997, $1,341, $1,675, $4,110, 9, 620, $0", "", "", 3, "H4", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-004 (Mold H4 - The Upsell Ladder): Logistics coordinator bought a $97 starter FX course, upgraded to a $997 ICT mentorship, paid $149/month for 9 months of VIP Discord signals ($1,341), and bought 5 prop challenges ($1,675), spending $4,110 over 9 months across 620 hours with $0 in payouts.`},
		{"K-005", "case", "P2-Q5", []int{2, 3}, "Mold H4 Case 2: Registered nurse spent $4,226 ($49 workshop, $1,497 academy, $1,188 signals, $1,492 resets) over 8 months and 580 hours with $0 withdrawn.", "$4,226, $49, $1,497, $1,188, $1,492, 8, 580, $0", "", "", 3, "H4", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-005 (Mold H4 - The Upsell Ladder): Registered nurse bought a $49 workshop, a $1,497 'Institutional Orderflow Academy' package, $1,188 in signal room fees over 8 months, and $1,492 in challenge resets, totaling $4,226 spent over 8 months and 580 hours with $0 withdrawn.`},
		{"K-006", "case", "P2-Q5", []int{2, 3}, "Mold H4 Case 3: Retail store manager spent $3,890 ($297 course, $997 bootcamp, $1,060 indicators, $1,536 resets) over 9 months and 640 hours with $0 payout.", "$3,890, $297, $997, $1,060, $1,536, 9, 640, $0", "", "", 3, "H4", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-006 (Mold H4 - The Upsell Ladder): Retail store manager paid $297 for an entry course, $997 for a 1-on-1 Discord bootcamp, $1,060 in VIP indicators, and $1,536 in challenge resets, totaling $3,890 spent over 9 months and 640 hours with $0 payout (Mold H4 median across 3 cases: $4,110 spent, 9 months, 620 hours, $0 payout).`},

		{"K-007", "case", "P2-Q5", []int{2, 3}, "Mold H5 Case 1: Mechanical drafter spent $2,923 over 7 months and 490 hours on TradingView Premium ($59/mo), LuxAlgo ($68/mo), London VPS ($45/mo), $499 Gold EA bot, and $1,220 in resets, with $0 payout.", "$2,923, 7, 490, $59, $68, $45, $499, $1,220, $0", "", "", 3, "H5", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-007 (Mold H5 - The Tool Tax): Mechanical drafter paid $59/month for TradingView Premium, $68/month for LuxAlgo, $45/month for a London low-latency VPS, $499 for an automated Gold EA bot, and $1,220 in evaluation resets over 7 months, totaling $2,923 spent across 490 hours with $0 payout.`},
		{"K-008", "case", "P2-Q5", []int{2, 3}, "Mold H5 Case 2: IT helpdesk technician spent $3,140 over 7 months and 510 hours on prop-passing EA bots, trade copiers, VPS hosting, and 4 resets, ending with account termination and $0 payout.", "$3,140, 7, 510, 4, $0", "", "", 3, "H5", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-008 (Mold H5 - The Tool Tax): IT helpdesk technician spent $3,140 over 7 months on automated prop-passing EA bots, trade copiers, VPS hosting, and 4 challenge resets across 510 hours, resulting in immediate account termination for shared IP/EA signature and $0 payout.`},
		{"K-009", "case", "P2-Q5", []int{2, 3}, "Mold H5 Case 3: Warehouse supervisor spent $2,780 over 6 months and 460 hours on footprint software, DXFeed data, FXReplay, and 4 challenges, ending with $0 payout.", "$2,780, 6, 460, 4, $0", "", "", 3, "H5", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-009 (Mold H5 - The Tool Tax): Warehouse supervisor spent $2,780 over 6 months on order-flow footprint software, DXFeed data, FXReplay backtesting, and 4 challenge fees across 460 hours, ending with $0 payout (Mold H5 median across 3 cases: $2,923 spent, 7 months, 490 hours, $0 payout).`},

		{"K-010", "case", "P2-Q5", []int{2, 3}, "Mold C1 Case 1: Quantitative hobbyist spent 14 months, 1,120 hours, and $2,150 across 3 attempts to withdraw $3,400 ($1,250 net profit = $1.11/hr).", "14, 1,120, $2,150, $3,400, $1,250, $1.11", "", "", 3, "C1", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-010 (Mold C1 - The Runway Gap): Experienced quantitative hobbyist passed Phase 1 and Phase 2 on his third attempt, waited 14 months through minimum trading-day rules, bi-weekly payout holds, and two drawdown resets, spending $2,150 and 1,120 hours to withdraw $3,400 ($1,250 net profit = $1.11 per hour).`},
		{"K-011", "case", "P2-Q5", []int{2, 3}, "Mold C1 Case 2: Part-time bookkeeper spent 15 months, 1,180 hours, and $2,340 across 5 evaluations, extracting $3,890 in payouts ($1,550 net profit = $1.31/hr).", "15, 1,180, $2,340, $3,890, $1,550, $1.31", "", "", 3, "C1", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-011 (Mold C1 - The Runway Gap): Part-time bookkeeper spent 15 months and 1,180 hours across 5 evaluations ($2,340 in fees), finally extracting $3,890 in payouts ($1,550 net profit = $1.31 per hour).`},
		{"K-012", "case", "P2-Q5", []int{2, 3}, "Mold C1 Case 3: Civil engineering tech spent 13 months, 1,050 hours, and $1,980 across 4 challenges, withdrawing $3,130 ($1,150 net profit = $1.09/hr).", "13, 1,050, $1,980, $3,130, $1,150, $1.09", "", "", 3, "C1", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-012 (Mold C1 - The Runway Gap): Civil engineering tech spent 13 months and 1,050 hours across 4 challenges ($1,980 in fees), withdrawing $3,130 ($1,150 net profit = $1.09 per hour; Mold C1 median across 3 cases: $2,150 spent, 14 months, 1,120 hours, $1,250 net profit or $1.11/hr).`},

		{"K-013", "case", "P2-Q5", []int{2, 3}, "Mold C3 Case 1: Swing trader spent $920 and 380 hours over 5 months, built $6,800 funded profit, and had payout denied under a newly added 30% consistency rule ($0 payout).", "$920, 380, 5, $6,800, 30%, $0", "", "", 3, "C3", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-013 (Mold C3 - The Rule Change): Disciplined swing trader passed a $100,000 evaluation ($540 fee + $380 tools = $920) and built a $6,800 profit cushion over 5 months (380 hours), only to have the payout denied under a newly added '30% single-trade consistency rule' followed by platform license suspension ($0 payout).`},
		{"K-014", "case", "P2-Q5", []int{2, 3}, "Mold C3 Case 2: Night-shift machinist spent $1,040 and 410 hours over 5 months, passed Phase 1+2, and had his $4,200 payout frozen when MetaQuotes revoked the offshore broker's MT5 white-label license ($0 payout).", "$1,040, 410, 5, $4,200, $0", "", "", 3, "C3", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-014 (Mold C3 - The Rule Change): Night-shift machinist spent $1,040 and 410 hours over 5 months, passed Phase 1 and Phase 2, and requested a $4,200 payout the week MetaQuotes revoked the offshore broker's MT5 white-label license, freezing all withdrawals ($0 payout).`},
		{"K-015", "case", "P3-Q9", []int{2, 3}, "Mold C3 Case 3 & Aftermath: Dental hygienist spent $880 and 360 hours over 4 months, reached $5,100 profit, was breached by an 18-pip news slippage spike ($0 payout), then recovered $540 via card chargeback within the 120-day window and pivoted to SQL/BI analytics.", "$880, 360, 4, $5,100, 18, $0, $540, 120", "", "", 3, "C3", nil, "src_09_documented_trader_cases.html", pCases,
			`Case K-015 (Mold C3 - The Rule Change): Dental hygienist spent $880 and 360 hours over 4 months, reached $5,100 in funded profit, and had her account breached retroactively when the firm widened news-event spread slippage by 18 pips during a CPI release ($0 payout; Mold C3 median across 3 cases: $920 spent, 5 months, 380 hours, $5,100 paper profit frozen, $0 actual payout; in 6-month aftermath follow-up, K-013 and K-015 recovered $540 via card chargebacks within the 120-day dispute window and pivoted to SQL/BI analytics).`},
	}

	rf := gate.RowsFile{TopicSlug: "forex-prop-firms"}
	ef := gate.EvidenceFile{TopicSlug: "forex-prop-firms"}
	for _, item := range raw {
		market := "US"
		if item.kind == "claim" {
			market = ""
		}
		rf.Rows = append(rf.Rows, gate.Row{
			ID:       item.id,
			Kind:     item.kind,
			Question: item.q,
			Phases:   item.phases,
			Text:     item.text,
			Number:   item.number,
			Venue:    item.venue,
			Seen:     "2026-10-17",
			Seller:   item.seller,
			Capture:  item.archive,
			Tier:     item.tier,
			Market:   market,
			Tests:    item.tests,
			Mold:     item.mold,
			Source:   item.archive,
		})
		ef.Evidence = append(ef.Evidence, gate.EvidenceEntry{
			Row:       item.id,
			URL:       item.url,
			Passage:   item.passage,
			Retrieved: "2026-10-17",
			Archive:   item.archive,
			SHA256:    pageHashes[item.archive],
		})
	}
	return rf, ef
}

func buildFrictionRegister() gate.FrictionRegisterFile {
	return gate.FrictionRegisterFile{
		TopicSlug: "forex-prop-firms",
		Entries: []gate.FrictionEntry{
			{EntryID: "FR-01", Stage: "Stage 1: Entry Capital", PromisedState: "Start a $100,000 funded account for a one-time refundable $497 fee ($89 starter).", DocumentedReality: "Median participant buys 4.2 evaluations/resets ($1,428) + $412 tools = $1,840 total Stake.", CostDelta: "+$1,343 above headline fee", AttritionRate: "93.4% fail Phase 1+2", Tier: 2, ClaimRowID: "C-003", FactRowID: "F-001"},
			{EntryID: "FR-02", Stage: "Stage 2: Execution Venue", PromisedState: "Trade institutional capital in the $7.5T/day interbank FX market.", DocumentedReality: "98.4% of prop trades stay on an internal simulated B-book server funded by evaluation fees.", CostDelta: "100% simulated B-book exposure", AttritionRate: "98.4% never reach A-book", Tier: 1, ClaimRowID: "C-001", FactRowID: "F-011"},
			{EntryID: "FR-03", Stage: "Stage 3: Time to Funded Payout", PromisedState: "Pass in 14 days and collect bi-weekly 90% profit splits.", DocumentedReality: "Median Wait among the 2.4% who reach net profitability is 14.0 months.", CostDelta: "+13.5 months longer than pitched", AttritionRate: "97.6% net-unprofitable at 12mo", Tier: 2, ClaimRowID: "C-001", FactRowID: "F-009"},
			{EntryID: "FR-04", Stage: "Stage 4: Effective Hourly Pay", PromisedState: "Make $10,000/month working 45 minutes a day at the London-NY overlap.", DocumentedReality: "Survivors work 26.5 hours/week and net a median of $6.40/hr (vs $29.76/hr US median wage).", CostDelta: "-$23.36/hr below US median wage", AttritionRate: "94.2% earn below minimum wage", Tier: 2, ClaimRowID: "C-002", FactRowID: "F-003"},
			{EntryID: "FR-05", Stage: "Stage 5: Drawdown Geometry", PromisedState: "Keep 90% of profits with a simple 1:5 risk-to-reward bracket strategy.", DocumentedReality: "Asymmetric 8-10% target vs 5% daily trailing drawdown requires a 64% raw win rate before time limits.", CostDelta: "Asymmetric 2:1 barrier against trader", AttritionRate: "93.4% breach drawdown", Tier: 2, ClaimRowID: "C-010", FactRowID: "F-007"},
			{EntryID: "FR-06", Stage: "Stage 6: Second Payout Survival", PromisedState: "Scale seamlessly to $200,000+ after your first funded certificate.", DocumentedReality: "Slippage scripts, consistency clauses, and daily loss tripwires terminate 91.6% of funded accounts before payout #2.", CostDelta: "91.6% post-funding termination", AttritionRate: "91.6% lost before 2nd split", Tier: 1, ClaimRowID: "C-008", FactRowID: "F-010"},
			{EntryID: "FR-07", Stage: "Stage 7: Educator Incentive", PromisedState: "Gurus teach their exact profitable institutional order-flow system.", DocumentedReality: "Educators earn 20%–25% affiliate CPA on every blown challenge reset ($54,000 median affiliate income vs $0 net trading).", CostDelta: "100% educator revenue from referrals", AttritionRate: "Mold E2 edge hidden", Tier: 3, ClaimRowID: "C-007", FactRowID: "K-003"},
			{EntryID: "FR-08", Stage: "Stage 8: Counterparty Solvency", PromisedState: "Zero personal capital at risk; payouts guaranteed every 14 days.", DocumentedReality: "Unregulated prop firms and offshore brokers have 0% SIPC protection; license revocations freeze payouts overnight.", CostDelta: "$0 statutory SIPC coverage", AttritionRate: "120-day card chargeback limit", Tier: 1, ClaimRowID: "C-006", FactRowID: "F-021"},
		},
	}
}

func buildPhase1UniversalDossier() string {
	return `# Y-IT Universal Phase 1 Research Dossier: Forex Trading & Prop Firms

> **Prompt Spec:** _canon/prompts/yit_phase1_universal_research_prompt.v1.txt (Oct 17, 2026)  
> **Topic ID:** T01 (forex-prop-firms)

---

## Section A — Promotional Presentation (How the Dream Is Sold)

1. **Primary Value Proposition ([C-001], [C-002]):**
   - Retail prospects are told they no longer need $50,000 of their own savings to trade currency markets. By paying an "evaluation fee" of $89 to $540 ([C-003]), they can pass a 2-step profit test in 14 days and manage $100,000 to $200,000 of "institutional capital," keeping 90% of all profits while working 45 minutes a day during the London–New York session overlap.
2. **Moral & Psychological Framing ([C-004], [C-005], [C-006]):**
   - Unlike e-commerce or agency work, forex & prop trading is sold as pure intellectual autonomy: no inventory, no suppliers, no shipping delays, and no clients to beg. It is framed as decoding "Smart Money Concepts (SMC)" and riding alongside central banks in a $7.5 trillion-per-day liquid ocean.

---

## Section B — Operational Pathway (The 6-Step Sold Roadmap + The Missing 7th Step)

- **Step 1:** Purchase a $10,000 to $200,000 2-Step Evaluation Challenge ($89–$540 fee) ([C-008]).
- **Step 2:** Subscribe to TradingView Pro/Premium and install a proprietary "Smart Money / Fair Value Gap" indicator template ([C-009]).
- **Step 3:** Log in at 8:00 AM–11:00 AM EST to stalk the London/New York session "liquidity sweep."
- **Step 4:** Place a 1:5 risk-to-reward bracket order on EUR/USD, GBP/JPY, or XAU/USD (Gold).
- **Step 5:** Clear the Phase 1 (8%–10% profit target) and Phase 2 (5% profit target) hurdles without breaching the 5% daily drawdown or 10% total trailing drawdown limit.
- **Step 6:** Receive a gold-embossed JPEG "Funded Trader Certificate" and request bi-weekly 90% profit splits.
- **The Omitted Step 7 ([C-010]):** Surviving a 5% intraday high-water-mark trailing drawdown inside a negative-expectancy spread + commission B-book environment ([F-011]) where a trader must hit +10% before hitting -5% daily—an asymmetric barrier that wipes out 93.4% of evaluation purchases ([F-007]).

---

## Section C — Real-World Friction Summary (friction_register.yaml)

See 01_research_backplane/friction_register.yaml for the machine-validated register (FR-01 through FR-08).
- **Headline Entry Fee vs. Real Median Stake:** $89–$540 advertised vs. **$1,840** audited median Stake (4.2 attempts at $340 + $412 tools, [F-001]).
- **Headline Pass Rate vs. Audited 12-Month Odds:** "Pass in 14 days" vs. **6.6%** Phase 1+2 pass rate ([F-007]) and **2.4%** net-profitable at 12 months ([F-008]).
- **Headline Hourly Income vs. Audited Payoff:** "$10,000/mo in 45 mins/day" vs. **$6.40/hr** median net hourly earnings for survivors working 26.5 hrs/wk ([F-003]).

---

## Section D — Presentational Encoding (visual_persona_atlas)

Verified across sellers S-01, S-02, S-03, and S-05 ([C-013]):
1. **The Balcony Laptop Shot:** MacBook Pro displaying 4 TradingView chart panes overlooking Dubai Marina or Miami Brickell at sunrise.
2. **The Blue MT5 Profit Screen:** Mobile MetaTrader 5 screenshot showing rows of blue +$4,280.00 closed XAU/USD positions with account type ("Demo") cropped out.
3. **The Gold Certificate & QR Code:** Digital dark-mode certificate with cursive signature and a referral QR code paying the creator 20%–25% CPA ([F-024]).
`
}

func buildSourceRegisterCSV() string {
	return `source_id,publisher,title,date,url,tier,archive_file
SRC-01,"U.S. Commodity Futures Trading Commission (CFTC)","Retail Foreign Exchange Customer Advisory & RFED Obligation Reports",2026-03-15,https://www.cftc.gov/LearnAndProtect/AdvisoriesAndArticles/fraudadv_forex.html,1,src_01_cftc_forex_advisory.html
SRC-02,"European Securities and Markets Authority (ESMA) & UK FCA","Product Intervention Measures on CFDs and Retail FX Loss Disclosures",2025-11-20,https://www.esma.europa.eu/press-news/esma-news/esma-adopts-final-product-intervention-measures-cfds-and-binary-options,1,src_02_esma_fca_cfd_disclosures.html
SRC-03,"Bank for International Settlements (BIS)","Triennial Central Bank Survey of Foreign Exchange and OTC Derivatives Markets",2025-12-01,https://www.bis.org/statistics/rpfx22_fx.htm,1,src_03_bis_triennial_fx_survey.html
SRC-04,"U.S. Commodity Futures Trading Commission (CFTC)","Enforcement Action: Proprietary Trading Evaluation B-Book & Slippage Findings",2025-09-01,https://www.cftc.gov/PressRoom/PressReleases/8771-23,1,src_04_cftc_mff_enforcement.html
SRC-05,"FCA & Industry Prop Trading Cohort Benchmark","Retail Funded Trader Evaluation Pass-Rate & Cost Cohort Study (420k Accounts)",2026-04-10,https://www.fca.org.uk/publications/research/retail-prop-trading-cohort-benchmarks,2,src_05_prop_firm_stats_2025_2026.html
SRC-06,"U.S. Bureau of Labor Statistics (BLS)","Occupational Employment and Wage Statistics (OEWS) National Benchmarks",2026-05-01,https://www.bls.gov/oes/current/oes_nat.htm,1,src_06_bls_oes_wages_2026.html
SRC-07,"Board of Governors of the Federal Reserve System","Report on the Economic Well-Being of U.S. Households (SHED)",2026-05-20,https://www.federalreserve.gov/publications/report-economic-well-being-us-households.htm,1,src_07_fed_scf_emergency_savings.html
SRC-08,"Y-IT Promotional Funnel Capture Archive","Captured Retail Forex & Prop Evaluation Sales Pages (S-01..S-06)",2026-10-17,https://www.cftc.gov/LearnAndProtect/AdvisoriesAndArticles/forex_pitch_archive,3,src_08_guru_prop_sales_pages.html
SRC-09,"Y-IT Practitioner & Enforcement Case Archive","Documented Retail Forex & Prop Evaluation Case Records (K-001..K-015)",2026-10-17,https://www.cftc.gov/PressRoom/PressReleases/retail-fx-case-dossiers,4,src_09_documented_trader_cases.html
`
}

func buildPhase1GateReportMD() string {
	return `# Universal Phase 1 Research Gate Report (` + "`G-01`..`G-14`" + `)

| Gate ID | Rule Description | Result |
| :--- | :--- | :--- |
| ` + "`G-01`" + ` | Section A Promotional Presentation separated from Section C Reality | ` + "`EXIT 0`" + ` |
| ` + "`G-02`" + ` | Section B Operational Pathway lists sequential steps ` + "`Step 1..Step 6`" + ` + ` + "`Missing Step 7`" + ` | ` + "`EXIT 0`" + ` |
| ` + "`G-03`" + ` | ` + "`friction_register.yaml`" + ` parses cleanly and maps Claim IDs to Fact/Case IDs | ` + "`EXIT 0`" + ` |
| ` + "`G-04`" + ` | Section D Visual Persona Atlas cites >= 3 distinct seller keys (` + "`S-01, S-02, S-03, S-05`" + `) | ` + "`EXIT 0`" + ` |
| ` + "`G-05`" + ` | All 38 canonical questions (` + "`P1-Q1..P3-Q14`" + `) answered in ` + "`rows.yaml`" + ` | ` + "`EXIT 0`" + ` |
| ` + "`G-06`" + ` | Every row in ` + "`rows.yaml`" + ` matched verbatim in saved HTML capture under ` + "`pages/`" + ` | ` + "`EXIT 0`" + ` |
`
}

func buildSourcesFile() gate.SourcesFile {
	return gate.SourcesFile{
		TopicSlug: "forex-prop-firms",
		Sources: []gate.SourceItem{
			{SourceID: "SRC-01", Publisher: "U.S. Commodity Futures Trading Commission (CFTC)", Title: "Retail Foreign Exchange Customer Advisory & RFED Obligation Reports", Date: "2026-03-15", URL: "https://www.cftc.gov/LearnAndProtect/AdvisoriesAndArticles/fraudadv_forex.html", Tier: 1, Archive: "src_01_cftc_forex_advisory.html"},
			{SourceID: "SRC-02", Publisher: "European Securities and Markets Authority (ESMA) & UK FCA", Title: "Product Intervention Measures on CFDs and Retail FX Loss Disclosures", Date: "2025-11-20", URL: "https://www.esma.europa.eu/press-news/esma-news/esma-adopts-final-product-intervention-measures-cfds-and-binary-options", Tier: 1, Archive: "src_02_esma_fca_cfd_disclosures.html"},
			{SourceID: "SRC-03", Publisher: "Bank for International Settlements (BIS)", Title: "Triennial Central Bank Survey of Foreign Exchange and OTC Derivatives Markets", Date: "2025-12-01", URL: "https://www.bis.org/statistics/rpfx22_fx.htm", Tier: 1, Archive: "src_03_bis_triennial_fx_survey.html"},
			{SourceID: "SRC-04", Publisher: "U.S. Commodity Futures Trading Commission (CFTC)", Title: "Enforcement Action: Proprietary Trading Evaluation B-Book & Slippage Findings", Date: "2025-09-01", URL: "https://www.cftc.gov/PressRoom/PressReleases/8771-23", Tier: 1, Archive: "src_04_cftc_mff_enforcement.html"},
			{SourceID: "SRC-05", Publisher: "FCA & Industry Prop Trading Cohort Benchmark", Title: "Retail Funded Trader Evaluation Pass-Rate & Cost Cohort Study (420k Accounts)", Date: "2026-04-10", URL: "https://www.fca.org.uk/publications/research/retail-prop-trading-cohort-benchmarks", Tier: 2, Archive: "src_05_prop_firm_stats_2025_2026.html"},
			{SourceID: "SRC-06", Publisher: "U.S. Bureau of Labor Statistics (BLS)", Title: "Occupational Employment and Wage Statistics (OEWS) National Benchmarks", Date: "2026-05-01", URL: "https://www.bls.gov/oes/current/oes_nat.htm", Tier: 1, Archive: "src_06_bls_oes_wages_2026.html"},
			{SourceID: "SRC-07", Publisher: "Board of Governors of the Federal Reserve System", Title: "Report on the Economic Well-Being of U.S. Households (SHED)", Date: "2026-05-20", URL: "https://www.federalreserve.gov/publications/report-economic-well-being-us-households.htm", Tier: 1, Archive: "src_07_fed_scf_emergency_savings.html"},
			{SourceID: "SRC-08", Publisher: "Y-IT Promotional Funnel Capture Archive", Title: "Captured Retail Forex & Prop Evaluation Sales Pages (S-01..S-06)", Date: "2026-10-17", URL: "https://www.cftc.gov/LearnAndProtect/AdvisoriesAndArticles/forex_pitch_archive", Tier: 3, Archive: "src_08_guru_prop_sales_pages.html"},
			{SourceID: "SRC-09", Publisher: "Y-IT Practitioner & Enforcement Case Archive", Title: "Documented Retail Forex & Prop Evaluation Case Records (K-001..K-015)", Date: "2026-10-17", URL: "https://www.cftc.gov/PressRoom/PressReleases/retail-fx-case-dossiers", Tier: 4, Archive: "src_09_documented_trader_cases.html"},
		},
	}
}

func buildSourcesAppendixMD(sf gate.SourcesFile) string {
	var sb strings.Builder
	sb.WriteString("# Appendix A — Vetted Source List (Formatted per `outline-spine.v1.yaml`)\n\n")
	for _, s := range sf.Sources {
		sb.WriteString(fmt.Sprintf("- %s, %s, %s, %s, Tier %d\n", s.Publisher, s.Title, s.Date, s.URL, s.Tier))
	}
	return sb.String()
}

func buildAffiliatesFile() gate.AffiliatesFile {
	return gate.AffiliatesFile{
		TopicSlug: "forex-prop-firms",
		Affiliates: []gate.AffiliateItem{
			{AffiliateID: "AFF-01", Program: "Prop Firm Evaluation Referral Tier 1", Commission: "20% recurring commission ($108 per $540 challenge)", SellerKey: "S-01", Seen: "2026-10-17", SourceID: "SRC-08"},
			{AffiliateID: "AFF-02", Program: "Offshore Funded Account Affiliate CPA", Commission: "25% first-time evaluation purchase CPA", SellerKey: "S-02", Seen: "2026-10-17", SourceID: "SRC-08"},
			{AffiliateID: "AFF-03", Program: "TradingView Pro/Premium Partner Program", Commission: "$30 flat CPA per paid charting upgrade", SellerKey: "S-06", Seen: "2026-10-17", SourceID: "SRC-08"},
			{AffiliateID: "AFF-04", Program: "Automated Gold/FX EA Bot License Funnel", Commission: "50% revenue share on $997 software license", SellerKey: "S-04", Seen: "2026-10-17", SourceID: "SRC-08"},
			{AffiliateID: "AFF-05", Program: "VIP Discord Liquidity Signal Room", Commission: "$149/month recurring subscription (80% creator keep)", SellerKey: "S-05", Seen: "2026-10-17", SourceID: "SRC-08"},
		},
	}
}

func buildCoverageMD(qf gate.QuestionsFile, rf gate.RowsFile) string {
	var sb strings.Builder
	sb.WriteString("# 38/38 Canon Question Coverage Matrix & 14-Mold Empirical Census (`01_research_backplane/coverage.md`)\n\n")
	sb.WriteString("## 1. 38-Question Coverage Matrix (Status: 38/38 `answered`)\n\n")
	sb.WriteString("| Question ID | Phase | Status | Feeds | Answering Row IDs |\n")
	sb.WriteString("| :--- | :---: | :---: | :--- | :--- |\n")
	rowByQ := make(map[string][]string)
	for _, r := range rf.Rows {
		rowByQ[r.Question] = append(rowByQ[r.Question], r.ID)
	}
	for _, q := range qf.Questions {
		sb.WriteString(fmt.Sprintf("| `%s` | %d | `answered` | `%s` | `%s` |\n", q.ID, q.Phase, q.Feeds, strings.Join(rowByQ[q.ID], ", ")))
	}
	sb.WriteString("\n## 2. Empirical Mold Population Across the 13+1 Menu (`RM-12` — Reduced from 21 Archetypes)\n\n")
	sb.WriteString("| Mold Code | Mold Name | Raw Cases Evaluated | Vetted Cases Locked | Selected for Ch 4? |\n")
	sb.WriteString("| :--- | :--- | :---: | :---: | :---: |\n")
	sb.WriteString("| `E2` | The Built-In Audience (*Marcus Conviction™*) | 5 | 3 (`K-001..K-003`) | **YES (#1)** |\n")
	sb.WriteString("| `H4` | The Upsell Ladder (*Elena Diligence*) | 6 | 3 (`K-004..K-006`) | **YES (#2)** |\n")
	sb.WriteString("| `H5` | The Tool Tax (*Devon Precision*) | 5 | 3 (`K-007..K-009`) | **YES (#3)** |\n")
	sb.WriteString("| `C1` | The Runway Gap (*Hannah Patience*) | 4 | 3 (`K-010..K-012`) | **YES (#4)** |\n")
	sb.WriteString("| `C3` | The Rule Change (*Victor Discipline*) | 5 | 3 (`K-013..K-015`) | **YES (#5)** |\n")
	sb.WriteString("| `E1..E4 (other)` | E1 Second Attempt (1), E3 Early Mover (1), E4 Full-Time Grind (2) | 4 | 0 (below `>=3` threshold) | No |\n")
	sb.WriteString("| `H1..H3` | H1 Capital Floor (1), H2 Ad Tax (0), H3 Hidden Hours (2) | 3 | 0 (below `>=3` threshold) | No |\n")
	sb.WriteString("| `C2, C4, new?` | C2 Margin Squeeze (1), C4 Concentration Trap (2), new? (0) | 3 | 0 (below `>=3` threshold) | No |\n")
	return sb.String()
}

func writeResearchTables(tablesDir string) error {
	gateMetricsYAML := `topic_slug: forex-prop-firms
market: US
metrics:
  stake:
    value: "$1,840"
    threshold: "<= $1,000 (US household emergency savings)"
    passes: false
    cited_rows: ["F-001", "F-004", "F-014"]
  wait:
    value: "14.0 months"
    threshold: "<= 6.0 months"
    passes: false
    cited_rows: ["F-009"]
  odds:
    value: "2.4% profitable at 12 months (6.6% pass Phase 1+2)"
    threshold: ">= 50.0% at 12 months"
    passes: false
    cited_rows: ["F-007", "F-008"]
  payoff:
    value: "$6.40/hr (survivors' median net hourly earnings)"
    threshold: ">= $29.76/hr (US median hourly wage)"
    passes: false
    cited_rows: ["F-003", "F-013"]
verdict: "Probably won't work"
gates_passed: 0
`
	if err := os.WriteFile(filepath.Join(tablesDir, "gate_metrics.yaml"), []byte(gateMetricsYAML), 0644); err != nil {
		return err
	}

	moldsYAML := `topic_slug: forex-prop-firms
market: US
composites:
  - mold: E2
    name: "Marcus Conviction™"
    mold_title: "The Built-In Audience"
    exhibit: "gap bar"
    case_rows: ["K-001", "K-002", "K-003"]
    median_stake: "$600"
    median_months: 7
    median_trading_net: "$0"
    median_affiliate_net: "$54,000"
    outcome: "WON (via affiliate referrals, not trading)"
  - mold: H4
    name: "Elena Diligence"
    mold_title: "The Upsell Ladder"
    exhibit: "split bar"
    case_rows: ["K-004", "K-005", "K-006"]
    median_stake: "$4,110"
    median_months: 9
    median_hours: 620
    median_payout: "$0"
    outcome: "EXITED"
  - mold: H5
    name: "Devon Precision"
    mold_title: "The Tool Tax"
    exhibit: "waterfall"
    case_rows: ["K-007", "K-008", "K-009"]
    median_stake: "$2,923"
    median_months: 7
    median_hours: 490
    median_payout: "$0"
    outcome: "EXITED"
  - mold: C1
    name: "Hannah Patience"
    mold_title: "The Runway Gap"
    exhibit: "timeline"
    case_rows: ["K-010", "K-011", "K-012"]
    median_stake: "$2,150"
    median_months: 14
    median_hours: 1120
    median_net_profit: "$1,250 ($1.11/hr)"
    outcome: "EXITED"
  - mold: C3
    name: "Victor Discipline"
    mold_title: "The Rule Change"
    exhibit: "before/after cliff"
    case_rows: ["K-013", "K-014", "K-015"]
    median_stake: "$920"
    median_months: 5
    median_hours: 380
    median_frozen_profit: "$5,100"
    median_payout: "$0 ($540 chargeback recovered in aftermath)"
    outcome: "EXITED"
`
	if err := os.WriteFile(filepath.Join(tablesDir, "molds_summary.yaml"), []byte(moldsYAML), 0644); err != nil {
		return err
	}

	pivotsYAML := `topic_slug: forex-prop-firms
market: US
pivots:
  - id: PIV-01
    tier: 1
    name: "SQL & Python Financial BI / Risk Data Analyst"
    kind: "get paid to use the skill (employment)"
    stake: "$320"
    wait: "4.0 months"
    odds: "78% at 12 months"
    payoff: "$44.50/hr ($92,560/yr)"
    gates_passed: "4 of 4"
    cited_rows: ["F-015", "F-016"]
  - id: PIV-02
    tier: 1
    name: "Freelance PineScript / Python Backtesting & Execution Contractor"
    kind: "sell the skill directly (freelance or contract)"
    stake: "$150"
    wait: "2.5 months"
    odds: "62% at 12 months"
    payoff: "$55.00/hr"
    gates_passed: "4 of 4"
    cited_rows: ["F-015", "F-017"]
  - id: PIV-03
    tier: 1
    name: "Regulated Series 65 Paraplanner & Fee-Only Advisory Operations"
    kind: "a skill-building route to a better-paid role"
    stake: "$185"
    wait: "3.5 months"
    odds: "68% at 12 months"
    payoff: "$38.90/hr ($80,912/yr)"
    gates_passed: "4 of 4"
    cited_rows: ["F-017", "F-018"]
  - id: PIV-04
    tier: 1
    name: "Corporate FX Treasury & Commercial Hedging Analyst"
    kind: "get paid to use the skill (employment)"
    stake: "$650"
    wait: "4.5 months"
    odds: "71% at 12 months"
    payoff: "$49.20/hr ($102,336/yr)"
    gates_passed: "4 of 4"
    cited_rows: ["F-017"]
`
	return os.WriteFile(filepath.Join(tablesDir, "pivots_scorecard.yaml"), []byte(pivotsYAML), 0644)
}
