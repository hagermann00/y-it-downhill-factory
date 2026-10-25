package pipeline

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"yit-downhill-factory/internal/catalog"
)

// BuildStage06And08And09 populates 06_covers_and_matter/, 08_marketing_and_freebie/, and 09_podcap_stub/.
func BuildStage06And08And09(rootDir string) error {
	if err := buildStage06CoversAndMatter(rootDir); err != nil {
		return err
	}
	if err := buildStage08Marketing(rootDir); err != nil {
		return err
	}
	if err := buildStage09Podcap(rootDir); err != nil {
		return err
	}
	return nil
}

func buildStage06CoversAndMatter(rootDir string) error {
	stageDir := filepath.Join(rootDir, "06_covers_and_matter")
	coversDir := filepath.Join(stageDir, "covers")
	frontDir := filepath.Join(stageDir, "front_matter")
	backDir := filepath.Join(stageDir, "back_matter")
	pamphletDir := filepath.Join(stageDir, "pamphlet")
	for _, d := range []string{stageDir, coversDir, frontDir, backDir, pamphletDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	// Render 5 Cover Concepts into 06_covers_and_matter/covers/
	var coverRecs []catalog.AssetRecord
	concepts := []struct {
		slug   string
		title  string
		sub    string
		accent string
		badge  string
	}{
		{"concept1-blueprint-grid", "why your FOREX PROP CHALLENGE probably will Fail", "EVEN IF YOUR AI CHAT BOT LIKES YOUR CHANCES!", "#1E3A8A", "We reveal the secret method in the first 2 chapters... First 2 chapters free!"},
		{"concept2-sinking-ship-posibot", "FOREX & PROP FIRMS: THE $200K MIRAGE", "Why It Probably Won't Work — Yeah… About That.", "#991B1B", "POSIBOT ON THE BURNING BOW: '90% PROFIT SPLIT!'"},
		{"concept3-flatpack-manual", "MODEL: FUNDED-TRADER-200K™", "Assembly Time: 14 Months · Step 7 Not Included", "#0369A1", "IKEA-STYLE PARODY EDITION · 0 OF 4 GATES PASSED"},
		{"concept4-forensic-autopsy", "CASE FILE #FX-2026: THE B-BOOK LEDGER", "98.4% Simulated Execution · 2.4% 12-Month Survival", "#92400E", "CORONER'S DOSSIER EDITION · AUDITED CFTC & BIS DATA"},
		{"concept5-neon-cliff", "THE 5% TRAILING DRAWDOWN TRAP", "How $497 Challenges Harvest $310 Million in Reset Fees", "#0F172A", "QUANTITATIVE REALITY CHECK EDITION"},
	}

	for idx, c := range concepts {
		svg := renderCoverConceptSVG(idx+1, c.title, c.sub, c.badge, c.accent)
		rec, err := catalog.WriteCatalogAsset(coversDir, 1, 0, idx+1, "cover", c.slug, 1, "svg", svg,
			fmt.Sprintf("Cover Concept %d: %s — %s (%s)", idx+1, c.title, c.sub, c.badge), "cover", "", []string{"F-007", "F-008"})
		if err == nil {
			coverRecs = append(coverRecs, rec)
		}
	}
	if err := catalog.WriteSnapshot(coversDir, "forex-prop-firms", coverRecs); err != nil {
		return err
	}

	planJSON, _ := json.MarshalIndent(map[string]interface{}{
		"topic_slug":             "forex-prop-firms",
		"primary_concept":        "T01-P0-C01-cover-concept1-blueprint-grid_v001",
		"kdp_trim_in":            "6.00x9.00",
		"kdp_wrap_in":            "12.405x9.25",
		"spine_width_in":         0.144,
		"spine_text_allowed":     false,
		"spine_rule_note":        "KDP paperbacks under 80 pages (64 pages) prohibit spine text.",
		"cover_sticker_verdict":  "PROBABLY WON'T WORK (2.4% ODDS AT 12 MONTHS)",
		"freebie_hook_banner":    "We reveal the secret method in the first 2 chapters... First 2 chapters free at itiswhatitis.biz",
		"concepts_generated":     5,
	}, "", "  ")
	if err := os.WriteFile(filepath.Join(coversDir, "plan.json"), planJSON, 0644); err != nil {
		return err
	}

	frontMD := `# Front Matter (Pages 1–6)

## Page 1: Series Half-Title & Canonical Tagline
**Why It** ~~Could.~~ ~~Would.~~ ~~Has.~~ ~~Might.~~ **Probably Won't.** **WORK**  
*Yeah… about that.*

## Page 3: Full Title Page
**WHY YOUR FOREX PROP CHALLENGE PROBABLY WILL FAIL**  
*Even If Your AI Chat Bot Likes Your Chances!*  
Funded Trader Challenges, 100x Leverage, and the B-Book Margin Machine  
**An IIWII / Y-IT Forensic Field Guide**

## Page 4: Copyright & Colophon
Copyright © 2026 It Is What It Is (` + "`IIWII`" + `) / Y-IT Imprint. All rights reserved.  
Print Geometry: KDP ` + "`6.00\" × 9.00\"`" + ` Trim (` + "`6.125\" × 9.25\"`" + ` Bleed) · 64 Pages · No Spine Text (` + "`<80p`" + ` Rule).

## Page 5: Dedication
*For everyone who ever calculated 1% daily compounding on a cocktail napkin at 1:30 AM, priced out a matte-black Porsche configurator by 1:45 AM, and got stopped out on GBP/JPY by 2:10 AM.*
`
	if err := os.WriteFile(filepath.Join(frontDir, "front_matter.md"), []byte(frontMD), 0644); err != nil {
		return err
	}

	backMD := `# Back Matter & Appendices (Pages 57–64)

## Page 57: Appendix A — Bulk Source List & Evidence Tiers (Part I)
- **Tier 1 (Academic, Statutory & Central Bank):** U.S. CFTC Retail Foreign Exchange Advisories & Enforcement Accounting (SRC-01, SRC-04), ESMA & UK FCA Product Intervention Disclosures (SRC-02), Bank for International Settlements Triennial Central Bank Survey (SRC-03), U.S. Bureau of Labor Statistics OEWS (SRC-06), Federal Reserve SHED Report (SRC-07).
- **Tier 2 (Industry & Regulatory Cohort Studies):** Prop Trading Evaluation Cohort Study across 420,000 Challenges (SRC-05).
- **Tier 3 (Captured Platform & Funnel Telemetry):** SHA-256 Archived Prop Firm & Guru Checkout Funnels S-01..S-06 and Affiliate CPA Schedules (SRC-08).
- **Tier 4 (Corroborated Practitioner Dossiers):** Corroborated Trader Case Files K-001..K-015 across Molds E2, H4, H5, C1, C3 (SRC-09).

## Pages 58–59: Appendix A Exhibits I & II — Dual-Page Research Machine Visualization Spread (Sequential with Bulk Source List)
- **Page 58 (Verso — Left Plate):** *How We Checked the Receipts (Part I: 5-Round Live Ingestion & Cryptographic Lock)* — ` + "`T01-P3-C08-diagram-research-machine-left_v001.svg`" + `. Depicts the 5-Round Live Harvest across Statutory & Regulatory Dockets, Peer-Reviewed Academic & Central Bank Indices, On-Demand Columnar OLAP Statistical Lakehouse, 12GB Dedicated Local VRAM + Elastic Cloud GPU Compute Array, and Temporal DOM / CDX Snapshot Archiver locked under ` + "`SHA-256`" + `.
- **Page 59 (Recto — Right Plate):** *How We Checked the Receipts (Part II: Convergent Ranges, 14-Mold Sieve & Compiler Gates)* — ` + "`T01-P3-C08-diagram-research-machine-right_v001.svg`" + `. Depicts the ` + "`RM-10v2`" + ` Multi-Source Magnitude Convergence Sieve (` + "`>= 3`" + ` independent sources, range-first integration + audited cohort median), the ` + "`13+1`" + ` Empirical Failure Mold Census, the 3-Layer Affiliate & Beneficial Mechanism Audit (` + "`BM-01..BM-08`" + `), and Deterministic Compiler Gates ` + "`G1–G7`" + `.

## Page 60: Appendix A (Continued) — Bulk Source List Part II & Chapter-by-Chapter Row Citation Trail
- Live harvested U.S. Federal Register dockets, Crossref DOI & OpenAlex academic studies (` + "`140+`" + ` validated records in ` + "`live_harvest_sources.csv`" + `), quantitative finance absorbing-barrier preprints, and the complete 56-row Chapter-by-Chapter Citation Trail (` + "`Ch 1..Ch 8`" + `).

## Page 61: Appendix B — The Unvarnished Glossary (Complete 20-Term Forensic Lexicon)
All 20 industry/guru terms (` + "`01..20`" + `) translated into plain forensic arithmetic with cited Row IDs.

## Page 62: Tear-Out Field Contract & 8 Beneficial Mechanisms (` + "`BM-01..BM-08`" + `)
- Personal Circuit Breaker contract (90-Day / 100-Trade zero-fee rule, $500 lifetime Stake ceiling, 6-month / $29.76/hr tripwires) plus the table of 8 Beneficial Mechanisms (` + "`BM-01..BM-08`" + `) adapted into Modules ` + "`P1..P7`" + ` of the $1.99 Companion Pamphlet.

## Page 63: *Results Typical Disclosure & Built at the IIWII Desk Colophon
- Every financial figure in Phases 2 and 3 prints both the **3-source convergent range** and the **audited cohort median** (` + "`$1,400–$2,200`" + ` Stake / ` + "`$1,840`" + ` median; ` + "`12–18 mo`" + ` Wait / ` + "`14.0 mo`" + ` median; ` + "`1.5%–3.5%`" + ` 12-mo Odds / ` + "`2.4%`" + ` median; ` + "`$4.50–$8.50/hr`" + ` Payoff / ` + "`$6.40/hr`" + ` median).
- Includes the Brian IIWII Architect Avatar and the PosiBot Canonical Blank Sign Pose Sheet (` + "`ADR 0044`" + `).

## Page 64: Last-Page Issue Index & Component Directory
Complete directory of metrics, Chapter 4 composites, exhibits, and back-matter schematics.
`
	if err := os.WriteFile(filepath.Join(backDir, "appendices_and_index.md"), []byte(backMD), 0644); err != nil {
		return err
	}

	pamphletMD := `# Y-IT Companion Pocket Pamphlet ($1.99): The Pre-Deposit Prop & Forex Survival Manual (Modules P1–P7 / BM-01..BM-08)

1. **Module P1 (BM-01 — True-Risk Drawdown De-Masking Worksheet):** Re-denominates every "$100,000 funded account" into its actual $5,000 daily / $10,000 max drawdown allowance ([F-007], [F-023]).
2. **Module P2 (BM-02 — 0.25% Quarter-Kelly Position Sizer & 2-Strike Intraday Lock):** Caps risk at 0.25% of nominal equity (2.5% of real drawdown buffer) and locks the terminal after 2 losing trades in 24 hours (-0.50% max day loss).
3. **Module P3 (BM-03 — Regulated CME Micro FX M6E vs Offshore B-Book Gate):** Routes capital to centralized exchange clearing (CME Micro FX Futures M6E/M6B with IRC Section 1256 60/40 tax treatment) instead of unlicensed offshore B-book simulators.
4. **Module P4 (BM-04 — 30-Minute News-Window Blackout & T-Cost Filter):** Mandatory flat window 15 minutes before and 15 minutes after Tier-1 macro releases (CPI, NFP, FOMC) to defeat asymmetric slippage ([K-015]).
5. **Module P5 (BM-05 — 90-Day / 100-Trade Zero-Fee Paper Expectancy Ledger):** Log 100 consecutive trades over 90 days on a free demo account before spending $1 ([F-023], [F-026]).
6. **Module P6 (BM-06 & BM-07 — 120-Day Chargeback Dispute & Contract Fine-Print Audit Kit):** Pre-purchase red-flag audit for 30% consistency rules + Visa Reason Code 13.1 / Mastercard 4853 evidence packet within 120 days ([F-021], [K-015]).
7. **Module P7 (BM-08 — 30-Day Quantitative Skill Translation Matrix):** Translates retail trading habits into $38.90–$55.00/hr corporate and contract roles ([F-015..F-018]).
`
	return os.WriteFile(filepath.Join(pamphletDir, "companion_pamphlet.md"), []byte(pamphletMD), 0644)
}

func renderCoverConceptSVG(num int, title, subtitle, badge, accent string) []byte {
	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 600 900" width="600" height="900">
  <defs>
    <pattern id="grid%d" width="24" height="24" patternUnits="userSpaceOnUse">
      <path d="M 24 0 L 0 0 0 24" fill="none" stroke="#93C5FD" stroke-width="0.6" stroke-opacity="0.35"/>
    </pattern>
  </defs>
  <rect x="0" y="0" width="600" height="900" fill="%s"/>
  <rect x="0" y="0" width="600" height="900" fill="url(#grid%d)"/>
  <rect x="28" y="28" width="544" height="844" rx="8" fill="none" stroke="#F8FAFC" stroke-width="3"/>

  <!-- Top Series Tagline Banner -->
  <rect x="48" y="48" width="504" height="54" rx="4" fill="#0F172A"/>
  <text x="300" y="81" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="15" font-weight="700" fill="#F8FAFC">Why It Probably Won&apos;t WORK — Yeah… about that.</text>

  <!-- Blueprint Callout Label -->
  <text x="56" y="138" font-family="'IBM Plex Mono', monospace" font-size="12" font-weight="700" fill="#FDE047">[COVER CONCEPT 0%d // BLUEPRINT SPEC]</text>

  <!-- Main Title Block -->
  <rect x="48" y="154" width="504" height="170" rx="6" fill="#FFFFFF" stroke="#0F172A" stroke-width="4"/>
  <text x="300" y="218" text-anchor="middle" font-family="'Big Shoulders Display', 'Impact', sans-serif" font-size="30" font-weight="900" fill="#0F172A">%s</text>
  <rect x="76" y="246" width="448" height="52" rx="4" fill="#DC2626"/>
  <text x="300" y="278" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="14" font-weight="700" fill="#FFFFFF">%s</text>

  <!-- Center Illustration Frame with PosiBot -->
  <rect x="68" y="352" width="464" height="340" rx="10" fill="#0F172A" stroke="#38BDF8" stroke-width="3"/>
  <!-- Candlestick cliff drop -->
  <polyline points="96,420 160,395 220,430 280,385 330,580 410,635 495,655" fill="none" stroke="#EF4444" stroke-width="6"/>
  <!-- PosiBot in Center -->
  <g transform="translate(215, 445)">
    <rect x="25" y="10" width="110" height="84" rx="12" fill="#C8A27A" stroke="#785A3C" stroke-width="3"/>
    <rect x="37" y="22" width="86" height="60" rx="8" fill="#0F172A"/>
    <circle cx="62" cy="45" r="7" fill="#00F0FF"/>
    <circle cx="98" cy="45" r="7" fill="#00F0FF"/>
    <path d="M 56 64 Q 80 80 104 64" fill="none" stroke="#00F0FF" stroke-width="4.5" stroke-linecap="round"/>
    <rect x="5" y="102" width="150" height="56" rx="6" fill="#FEF08A" stroke="#0F172A" stroke-width="3"/>
    <text x="80" y="136" text-anchor="middle" font-family="'Big Shoulders Display', sans-serif" font-size="22" font-weight="900" fill="#0F172A">90%% PROFIT SPLIT!*</text>
  </g>
  <text x="300" y="672" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="11.5" fill="#94A3B8">*97.6%% of starters net-unprofitable at Month 12 [F-008]</text>

  <!-- Bottom Hook Banner (from ME ODDEN CAR TOON_files) -->
  <rect x="48" y="724" width="504" height="118" rx="6" fill="#FEF08A" stroke="#0F172A" stroke-width="3.5"/>
  <text x="300" y="764" text-anchor="middle" font-family="'Work Sans', sans-serif" font-size="16" font-weight="800" fill="#0F172A">%s</text>
  <text x="300" y="796" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="13" font-weight="700" fill="#B91C1C">AUDITED VERDICT: 0 OF 4 GATES PASSED · $1,840 MEDIAN STAKE</text>
  <text x="300" y="824" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="11" fill="#1E293B">Y-IT IMPRINT · ITISWHATITIS.BIZ</text>
</svg>`, num, accent, num, num, escapeSVG(title), escapeSVG(subtitle), escapeSVG(badge))
	return []byte(svg)
}

func escapeSVG(s string) string {
	return strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&apos;",
	).Replace(s)
}

func buildStage08Marketing(rootDir string) error {
	stageDir := filepath.Join(rootDir, "08_marketing_and_freebie")
	freebieDir := filepath.Join(stageDir, "freebie")
	landingDir := filepath.Join(stageDir, "landing")
	emailDir := filepath.Join(stageDir, "email")
	adsDir := filepath.Join(stageDir, "ads")
	for _, d := range []string{stageDir, freebieDir, landingDir, emailDir, adsDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	variants := map[string]interface{}{
		"topic_slug":      "forex-prop-firms",
		"canon_lock_date": "2026-10-12",
		"freebie_hook":    "We reveal the secret method in the first 2 chapters — download Chapters 1 & 2 free before you wire another $340 reset fee.",
		"variants": []map[string]string{
			{
				"id":       "v1_value",
				"angle":    "Direct ROI Math",
				"headline": "Spend $9.99 Before You Wire $1,840 in Prop Challenge Reset Fees.",
				"body":     "Across 420,000 tracked retail prop-firm evaluations, the median trader buys 4.2 challenges ($1,428) plus $412 in indicator and VPS subscriptions before walking away with $0. Read the CFTC B-book accounting first.",
				"cta":      "Download Chapters 1 & 2 Free (Or Grab the Full 64-Page Autopsy for $9.99)",
			},
			{
				"id":       "v2_risk",
				"angle":    "Downside & Drawdown Protection",
				"headline": "A '$100,000 Funded Account' With a 5% Daily Drawdown Is Just a $5,000 Account in a Trench Coat.",
				"body":     "98.4% of prop evaluation trades never leave an internal simulated B-book server, and 91.6% of funded accounts get terminated before a second payout. See how the trapdoor is wired.",
				"cta":      "Read the Forensic Autopsy Before Your Next Reset",
			},
			{
				"id":       "v3_smart_buyer",
				"angle":    "For the Analytical Skeptic",
				"headline": "You Already Suspected the Dubai Balcony Laptop Was Rented. Here Is the CFTC Ledger.",
				"body":     "Why do prop trading gurus give you a 20% discount code? Because they make $54,000 a year on reset referrals while netting $0 on MetaTrader 5. We pulled every filing.",
				"cta":      "Get the Numbers the Discord Server Deleted",
			},
			{
				"id":       "v4_time",
				"angle":    "Opportunity Cost & Hourly Wage",
				"headline": "The Top 2.4% of Prop Survivors Work 26.5 Hours a Week for $6.40 an Hour. Airport Pretzel Kiosks Pay Triple.",
				"body":     "Even if you pass Phase 1 and Phase 2, the median Wait to net profitability is 14.0 months—for $6.40/hr. Chapter 7 shows 4 ways to sell the exact same SQL/PineScript brain for $44.50 to $55.00/hr.",
				"cta":      "Trade the B-Book Casino for a $49.20/Hr Paycheck",
			},
			{
				"id":       "v5_price",
				"angle":    "Impulse Price Anchor",
				"headline": "Costs Less Than a Single 2-Pip Gold Slippage Spike.",
				"body":     "One month of a VIP Discord signal room is $149. One prop reset is $340. This entire 64-page forensic teardown is $9.99—and the first two chapters are free right now.",
				"cta":      "Read Chapters 1 & 2 Free Today",
			},
		},
	}
	vBytes, _ := json.MarshalIndent(variants, "", "  ")
	if err := os.WriteFile(filepath.Join(landingDir, "variants.json"), vBytes, 0644); err != nil {
		return err
	}

	landingMD := `# Y-IT Landing Page Copy Pack — All 5 Canonical Angles (08_marketing_and_freebie/landing/)

1. **v1_value (Direct ROI):** Spend $9.99 before you wire $1,840 (4.2 challenges + $412 tools, [F-001]).
2. **v2_risk (Downside Protection):** A $100,000 account with a 5% daily trailing drawdown is a $5,000 account in a Halloween costume ([F-007], [F-011]).
3. **v3_smart_buyer (Analytical Skeptic):** You already knew the G-Wagon was rented ([C-013]). See how Mold E2 affiliates make $54,000 on referral codes while netting $0 on trades ([K-003]).
4. **v4_time (Opportunity Cost):** 14.0 months and 26.5 hours/week for $6.40/hr ([F-003], [F-009])—or pivot the same quantitative skill into $44.50–$55.00/hr ([F-016], [F-017]).
5. **v5_price (Impulse Anchor):** Cheaper than one month of TradingView Premium ($59/mo, [K-007])—and Chapters 1 & 2 are free.
`
	if err := os.WriteFile(filepath.Join(landingDir, "landing_copy_all_5_angles.md"), []byte(landingMD), 0644); err != nil {
		return err
	}

	dripMD := `# 3-Step Email Drip Sequence (08_marketing_and_freebie/email/drip_sequence.md)

## Email 1 (Immediate Delivery): Your Free 2-Chapter Roadmap (And What Happens on Page 14)
- **Subject:** Here is the $200,000 Funded Trader Roadmap they charge $997 for (Look at Step 7)
- **Body:** Attached is your standalone PDF of **Chapters 1 & 2** (*Vibrant Hope*). Read the 6-step flat-pack manual for unlocking a $200,000 prop account ([C-001], [C-008]). Then look at the bottom of the final page where **Step 7 was not provided** ([C-010]). Tomorrow, we open the CFTC enforcement ledger.

## Email 2 (+24 Hours): How to Lose a Prop Challenge During a Winning Week
- **Subject:** 98.4% of prop trades never leave the building ([F-011])
- **Body:** Meet *Devon Precision* (Mold H5) and *Victor Discipline* (Mold C3) from Chapter 4. One spent $2,923 on indicators, VPS boxes, and resets ([K-007]); the other passed Phase 1 and Phase 2, built a $5,100 funded profit cushion, and got frozen overnight by a retroactive consistency rule ([K-013], [K-015]).

## Email 3 (+48 Hours): The $44.50/Hr Salvage Operation (Plus the 120-Day Chargeback Window)
- **Subject:** Stop paying $340 cover charges. Make a CFO pay you $49.20/hr instead.
- **Body:** In Chapter 7, we score 4 concrete pivots ([PIVOT: THE PRACTICAL PIVOT]) that pass 4 of 4 gates ($38.90/hr to $55.00/hr, [F-016..F-018]). And in Chapter 8, we show you how *Victor Discipline* clawed back $540 using the **120-day Visa/Mastercard chargeback window** ([F-021], [K-015]). Grab the full 64-page book ($9.99) or the complete Book + 25-Min Podcap + $1.99 Pocket Pamphlet Bundle.
`
	if err := os.WriteFile(filepath.Join(emailDir, "drip_sequence.md"), []byte(dripMD), 0644); err != nil {
		return err
	}

	adsJSON, _ := json.MarshalIndent(map[string]interface{}{
		"topic_slug": "forex-prop-firms",
		"campaigns": []map[string]string{
			{"id": "MON-01", "channel": "High-Intent Search Paid / Organic", "hook": "Before you type that 20% prop firm discount code, see why the influencer makes $108–$162 every time you breach your 5% daily drawdown."},
			{"id": "MON-02", "channel": "Chargeback & Frozen Payout Recovery", "hook": "Prop Firm Denied Your Payout Under a Retroactive 30% Consistency Rule? Read the 120-Day Visa/Mastercard Dispute Protocol (Reason Code 13.1 / 4853)."},
			{"id": "MON-03", "channel": "B-Book vs Regulated CME Micro FX (M6E)", "hook": "98.4% of retail prop trades never leave an internal B-book simulator. Compare offshore challenges vs regulated CME Micro FX Futures (IRC Section 1256)."},
			{"id": "MON-04", "channel": "Quantitative Career Pivot ($38.90–$55.00/hr)", "hook": "Stop donating 26.5 hours/week to a $6.40/hr prop simulator. Translate your Python, PineScript & SQL habits into a $44.50–$55.00/hr analyst role."},
			{"id": "MON-05", "channel": "Subscription & Indicator Detox", "hook": "Spending $276/month on TradingView Premium, LuxAlgo, VPS boxes, and VIP Discord rooms? Run the 90-Day Zero-Fee Paper Gate first."},
			{"id": "MON-06", "channel": "Spouse / Family Pre-Deposit Intervention", "hook": "A '$100,000 Funded Account' with a 5% daily drawdown is a $5,000 account in a Halloween costume. First 2 chapters free at itiswhatitis.biz."},
			{"id": "MON-07", "channel": "Tax & Accounting Reality (IRC 988 vs 1256)", "hook": "Prop evaluation fees are non-recoverable service expenses—not Section 1256 capital losses. Read the forensic accounting breakdown."},
			{"id": "MON-08", "channel": "25-Min Dual-Persona Commuter Podcap", "hook": "Listen to the Forensic Accountant vs. the 2:00 AM ICT Believer dissect 420,000 prop challenges in 25 minutes."},
		},
	}, "", "  ")
	return os.WriteFile(filepath.Join(adsDir, "ad_copy_pack.json"), adsJSON, 0644)
}

func buildStage09Podcap(rootDir string) error {
	stageDir := filepath.Join(rootDir, "09_podcap_stub")
	if err := os.MkdirAll(stageDir, 0755); err != nil {
		return err
	}

	scriptMD := `# Y-IT Podcap Episode 01: Forex Trading & Prop Firms (25-Minute Dual-Persona Script)

> **Format Spec:** ADR 0038 (*Podcap — Audio-First Product Format & Tiered Sales Bundle*)  
> **Runtime:** 25:00 (6 Act Segments)  
> **Host A (ANALYST):** Dry, surgical, sardonic forensic accountant carrying the CFTC/BIS/BLS files.  
> **Host B (BELIEVER):** Recovering 2:00 AM chart-watcher who knows every ICT buzzword by heart.

---

## [00:00 – 04:00] ACT I: THE 11:42 PM SEDUCTION (PHASE 1 — VIBRANT HOPE)
- **BELIEVER:** Picture this. It's 11:42 PM on a Tuesday ([C-005]). Your neighbor Greg is ironing his khaki prison uniform for his 8:00 AM commute. You're in flannel pajama pants on the sofa, tapping three buttons on MetaTrader 5, bracketing a 40-pip Gold displacement candle, and waking up to $2,400 before your Pop-Tart pops ([C-005]).
- **ANALYST:** And best of all, you didn't even need $100,000 of your own money! According to Seller S-01, you just pay a refundable $497 audition fee ([C-003]), pass a 2-step test in 14 days, and trade a $200,000 institutional account while keeping 90% of the profits ([C-001]). No warehouses, no angry Amazon buyers, no bosses ([C-006]).
- **BELIEVER:** Plus the flat-pack manual in Chapter 2 only has six steps ([C-008])! Snap on the green-and-red Smart Money indicator template, wait for the 9:30 AM New York liquidity sweep, fire a 1:5 bracket, and frame your gold JPEG Funded Trader Certificate ([C-008], [C-013]).
- **ANALYST:** Except for that tiny dashed box at the bottom of page 18: **Step 7 was not provided** ([C-010]).

## [04:00 – 09:30] ACT II: THE $1,840 B-BOOK CASINO (CHAPTER 3 — THE MONEY PIT)
- **ANALYST:** Let's open the CFTC enforcement file ([F-002], [F-011]). One retail prop firm collected **$310 million in evaluation and reset fees from 135,000 customers** in 22 months—and routed **98.4% of every single trade to an internal simulated B-book server** ([F-002], [F-011]).
- **BELIEVER:** Wait. So when the bio says "Institutional Funded Trader"...
- **ANALYST:** They are playing Mario Kart inside a private database in St. Vincent. Across 420,000 tracked challenges, the median trader buys **4.2 evaluations or resets** at $340 a pop ($1,428) plus $412 in TradingView, VPS, and neon indicators—a median **Stake of $1,840** ([F-001], [F-004]), which is 84% more than the typical American family's entire $1,000 emergency fund ([F-014]).

## [09:30 – 15:00] ACT III: FIVE CORONER FILES ON AMBER PAPER (CHAPTER 4 — TALES FROM THE TRENCHES)
- **ANALYST:** In Chapter 4, PosiBot leaves the room and the jokes stop. Five composites built from 15 documented cases ([K-001..K-015]):
  1. **Marcus Conviction™ (Mold E2):** Netted $0 trading currencies over 7 months—and made **$54,000** from his 20% prop-firm affiliate referral link ([K-003], [F-024]).
  2. **Elena Diligence (Mold H4):** Paid $4,110 up the course-and-Discord-signal ladder over 9 months. Withdrew $0 ([K-004..K-006]).
  3. **Devon Precision (Mold H5):** Paid $2,923 in TradingView Premium, LuxAlgo, VPS hosting, a $499 Gold EA bot, and resets. Account terminated for shared EA IP ([K-007..K-009]).
  4. **Hannah Patience (Mold C1):** Did everything right. Passed Phase 1 and Phase 2. Took 14 months and 1,120 hours to net $1,250—which works out to **$1.11 an hour** ([K-010..K-012]).
  5. **Victor Discipline (Mold C3):** Built $5,100 in funded profit over 5 months, requested a payout, and got hit with a retroactive 30% consistency rule and an offshore MT5 license freeze ([K-013..K-015]).

## [15:00 – 19:00] ACT IV: THE 5% DRAWDOWN WOODCHIPPER & THE VERDICT (CHAPTERS 5 & 6)
- **BELIEVER:** Explain the Halloween costume math from Chapter 5.
- **ANALYST:** A "$100,000 account" with a 5% daily trailing drawdown only gives you $5,000 of daily room ([F-007]). Asking you to make +8% ($8,000) and +5% ($5,000) without losing $5,000 requires a **64% raw win rate** after spread drag ([C-010]). Only **6.6%** pass Phase 1+2 ([F-007]), **91.6%** of funded accounts get terminated before Payout #2 ([F-010]), and only **2.4%** are net-profitable at Month 12 ([F-008]), earning a median of **$6.40 an hour** across 26.5 hours a week ([F-003]).
- **BELIEVER:** Against a $29.76/hr US median wage ([F-013]). That's **0 of 4 gates passed**. Verdict: **Probably won't work.**

## [19:00 – 25:00] ACT V: THE $55/HR PIVOT & THE STUBBORN FOOL'S ARMOR (CHAPTERS 7 & 8)
- **ANALYST:** Chapter 7 is where the news gets good. If you learned time-series data cleaning, PineScript, or macro variance tracking, employers actually pay for that ([F-015]):
  - **Freelance PineScript/Python Backtest Dev:** $150 Stake, 2.5 mo Wait, 62% Odds, **$55.00/hr** (4 of 4 Gates, [F-017]).
  - **SQL / Financial BI Analyst (SOC 15-2051):** $320 Stake, 4.0 mo Wait, 78% Odds, **$44.50/hr ($92,560/yr)** (4 of 4 Gates, [F-016]).
  - **Corporate FX Treasury Analyst (SOC 13-2051):** **$49.20/hr ($102,336/yr)** (4 of 4 Gates, [F-017]).
- **BELIEVER:** And for the stubborn fool in Chapter 8 who still won't close the chart: run the **90-day zero-fee 100-trade test** ([F-023], [F-026]), hard-cap your lifetime Stake at **$500**, cancel the $149/mo Discord rooms ([F-019]), and if a prop firm freezes your payout, file a Visa/Mastercard dispute inside the **120-day chargeback window** just like Victor Discipline did to claw back his $540 ([F-021], [K-015]).
`
	if err := os.WriteFile(filepath.Join(stageDir, "podcap_25min_script.md"), []byte(scriptMD), 0644); err != nil {
		return err
	}

	bundleJSON, _ := json.MarshalIndent(map[string]interface{}{
		"adr_reference": "ADR 0038 — Podcap Audio-First Product Format & Tiered Sales Bundle",
		"topic_slug":    "forex-prop-firms",
		"tiers": []map[string]interface{}{
			{
				"tier_id":      "TIER-0-FREEBIE",
				"name":         "The Sucker's Roadmap Freebie Pack (Chapters 1 & 2 + 3-Min Audio Teaser)",
				"price_usd":    0.00,
				"deliverables": []string{"08_marketing_and_freebie/freebie/ch1_ch2_free_giveaway.pdf"},
			},
			{
				"tier_id":      "TIER-1-CORE-BOOK",
				"name":         "Y-IT Core 64-Page Forensic Book (Digital PDF / KDP 6x9 Paperback)",
				"price_ebook":  9.99,
				"price_print":  14.99,
				"deliverables": []string{"07_assembly_kdp/build/book_complete_kdp.pdf"},
			},
			{
				"tier_id":      "TIER-2-COMPANION-PAMPHLET",
				"name":         "Y-IT Standalone Companion Pocket Pamphlet (90-Day Ledger & Chargeback Guide)",
				"price_usd":    1.99,
				"deliverables": []string{"06_covers_and_matter/pamphlet/companion_pamphlet.pdf"},
			},
			{
				"tier_id":      "TIER-3-FULL-AUTOPSY-PODCAP-BUNDLE",
				"name":         "Complete 64-Page Book + 25-Min Dual-Persona Podcap + Companion Pamphlet",
				"price_usd":    19.99,
				"deliverables": []string{
					"07_assembly_kdp/build/book_complete_kdp.pdf",
					"06_covers_and_matter/pamphlet/companion_pamphlet.pdf",
					"09_podcap_stub/podcap_25min_script.md",
				},
			},
		},
	}, "", "  ")
	if err := os.WriteFile(filepath.Join(stageDir, "tiered_sales_bundle.json"), bundleJSON, 0644); err != nil {
		return err
	}

	ttsJSON, _ := json.MarshalIndent(map[string]interface{}{
		"topic_slug":          "forex-prop-firms",
		"singleton_lock_rule": "D:\\hq\\_record\\tts.lock (Sequential queue regulation required)",
		"voices": map[string]string{
			"ANALYST":  "en-US-Forensic-Baritone",
			"BELIEVER": "en-US-Conversational-Tenor",
		},
		"chunks_count": 5,
	}, "", "  ")
	return os.WriteFile(filepath.Join(stageDir, "tts_chunk_manifest.json"), ttsJSON, 0644)
}
