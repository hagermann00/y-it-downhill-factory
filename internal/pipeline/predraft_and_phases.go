package pipeline

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"yit-downhill-factory/internal/catalog"
	"yit-downhill-factory/internal/gate"
)

// BuildStage02Through05 populates 02_predraft_spine/, 03_phase1_run/, 04_phase2_run/, and 05_phase3_run/
// with high-personality, sharp-edged Y-IT prose that strictly respects the 3-Phase canon rules.
func BuildStage02Through05(rootDir string) error {
	if err := buildStage02Predraft(rootDir); err != nil {
		return err
	}
	if err := buildStage03Phase1(rootDir); err != nil {
		return err
	}
	if err := buildStage04Phase2(rootDir); err != nil {
		return err
	}
	if err := buildStage05Phase3(rootDir); err != nil {
		return err
	}
	return nil
}

func buildStage02Predraft(rootDir string) error {
	stageDir := filepath.Join(rootDir, "02_predraft_spine")
	for _, sub := range []string{"phase-1", "phase-2", "phase-3"} {
		if err := os.MkdirAll(filepath.Join(stageDir, sub), 0755); err != nil {
			return err
		}
	}

	intakeMD := `# Stage 02 Intake & Gate Computation Brief: Forex Trading & Prop Firms

- **Topic ID:** T01 (forex-prop-firms)
- **Canonical Spine:** outline-spine.v1.yaml (ADR 0048–0050)
- **Computed 4-Metric Gate Scorecard (Chapter 6):**
  1. **The Stake:** $1,840 median total cash spent (4.2 challenge/reset attempts at $340 = $1,428 + $412 tools, [F-001], [F-004]) vs. <= $1,000 US household emergency savings benchmark ([F-014]) -> **MISS**
  2. **The Wait:** 14.0 months median time to consistent net profitability among survivors ([F-009]) vs. <= 6.0 months threshold -> **MISS**
  3. **The Odds:** 2.4% of starters net-profitable at 12 months (6.6% pass Phase 1+2 evaluation, [F-007], [F-008]) vs. >= 50.0% threshold -> **MISS**
  4. **The Payoff:** $6.40/hr median net hourly earnings across 26.5 hrs/wk ([F-003]) vs. >= $29.76/hr US median hourly wage ([F-013]) -> **MISS**
- **Computed Verdict:** **Probably won't work** (0 of 4 gates passed).
- **Selected Chapter 4 Composite Molds (5 molds, 3 documented cases each):**
  1. E2 (*The Built-In Audience*) — **Marcus Conviction™** ([K-001], [K-002], [K-003]): Median $600 trading stake, $0 net trading payout, $54,000 affiliate referral income.
  2. H4 (*The Upsell Ladder*) — **Elena Diligence** ([K-004], [K-005], [K-006]): Median $4,110 spent across courses, VIP Discord signals, and challenge resets; $0 payout.
  3. H5 (*The Tool Tax*) — **Devon Precision** ([K-007], [K-008], [K-009]): Median $2,923 spent on TradingView Premium, LuxAlgo, London VPS, EA bots, and resets; $0 payout.
  4. C1 (*The Runway Gap*) — **Hannah Patience** ([K-010], [K-011], [K-012]): Median 14 months, 1,120 hours, $2,150 fees, $1,250 net profit ($1.11/hr).
  5. C3 (*The Rule Change*) — **Victor Discipline** ([K-013], [K-014], [K-015]): Median $920 spent, $5,100 funded profit frozen by retroactive consistency rule / white-label MT5 revocation; $0 payout ($540 recovered via 120-day chargeback in aftermath).
`
	if err := os.WriteFile(filepath.Join(stageDir, "intake.md"), []byte(intakeMD), 0644); err != nil {
		return err
	}

	opusBriefMD := `# Composition & Voice Calibration Brief (opus-brief.md)

1. **Phase 1 (Chapters 1–2) — Deadpan Corporate Sincerity (Glossy Cult-Brochure Satire):**
   - Play the pitch 100% straight, with the breathless, hyper-caffeinated reverence of a Dubai penthouse webinar. Zero winking, zero premature warnings, zero audited F-### rows. Let the sheer audacity of the pitch (*"Tap three buttons in your boxers and siphon $2,400 from central banks before your toaster pops"*) carry the comedic edge.
   - Chapter 2 is a deadpan flat-pack assembly manual (FUNDED-TRADER-200K™) that treats beating the global currency market like bolting together a Swedish bookshelf—right up until "Step 7 was not provided."
2. **Phase 2 (Chapters 3–6) — Razor-Edged Sardonic Authority & Cold Coroner Autopsy:**
   - **Chapters 3, 5, 6:** Take the gloves off against the machine. Dissect the B-book bucket shop, the rented G-Wagons, and the 5% trailing drawdown woodchipper with dry, memorable, forensic wit. **Crucial rule:** the blade cuts the *casino and its barkers*, never the reader who walked through the door.
   - **Chapter 4 (*Tales from the Trenches*):** Complete tonal whiplash by design. Aged amber paper (#EFE5CE), **zero PosiBot**, **zero jokes**. Pure coroner's ledger. When someone pays $4,110 for Discord arrows and walks out with $0.00, you don't need a punchline—the receipt *is* the punchline.
3. **Phase 3 (Chapters 7–8) — Street-Smart Solidarity & Quiet Reckoning:**
   - **Chapter 7:** Constructive, sharp-tongued solidarity. Show the reader how to stop paying $340 cover charges to a server in St. Vincent and start charging corporate CFOs $44.50 to $55.00 an hour ([PIVOT: THE PRACTICAL PIVOT]) for the exact same data obsession. Close the Chapter 7 band on the exact words: For the Stubborn Fool Who Must.
   - **Chapter 8:** Affectionate, battle-scarred barstool honesty (Quiet Reckoning) for the stubborn bastard who still refuses to close MetaTrader. Give them the 90-day zero-fee crucible, the $500 fuse box, the 6 Do-Nots, the 120-day Visa chargeback hammer, and a parting shot they will remember at 2:00 AM.
`
	if err := os.WriteFile(filepath.Join(stageDir, "opus-brief.md"), []byte(opusBriefMD), 0644); err != nil {
		return err
	}

	// Phase 1 packet, tables, predraft
	p1Pkt := gate.PhasePacket{
		TopicSlug:     "forex-prop-firms",
		Phase:         1,
		SpineVersion:  "outline.v1@282d710",
		PacketVersion: 1,
		Dials: map[string]interface{}{
			"ch1.sections":      5,
			"ch1.posibot":       3,
			"ch2.steps":         6,
			"ch2.parts":         3,
			"ch2.missing_step":  "Surviving 5% trailing drawdown inside a negative-expectancy B-book simulation",
			"last_posibot_meth": "clipboard",
		},
		CitedRows: []string{"C-001", "C-002", "C-003", "C-004", "C-005", "C-006", "C-007", "C-008", "C-009", "C-010", "C-011", "C-012", "C-013", "C-014"},
	}
	p1Bytes, _ := yaml.Marshal(&p1Pkt)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-1", "packet.yaml"), p1Bytes, 0644)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-1", "tables.md"), []byte("# Phase 1 Claim Tables (Claims Only — Zero Fact Rows)\n\n| Row ID | Question | Claim Summary | Pitch Numbers |\n| :--- | :--- | :--- | :--- |\n| `C-001` | `P1-Q1` | 2-Step Challenge in 14 days, up to $200,000 funded account | `14 days, $200,000, 90%` |\n| `C-002` | `P1-Q1` | 45 minutes/day during London-NY overlap for $10,000/mo | `45 minutes, $10,000, $7.5 trillion` |\n| `C-003` | `P1-Q2` | $100,000 evaluation for $497 ($89 starter) with 100% fee refund | `$100,000, $497, $89, 100%` |\n| `C-008` | `P1-Q7` | 6-Step Funded Trader Roadmap | `8%, 5%, 90%` |\n| `C-010` | `P1-Q9` | Missing Step 7: Surviving 5% intraday trailing drawdown | `5%, 64%` |\n"), 0644)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-1", "predraft.md"), []byte("# Phase 1 Pre-Draft Spine: Vibrant Hope (Chapters 1 & 2)\n\n## Chapter 1: The Promise (Magazine-Ad Collage)\n- Opens at 11:42 PM (`[C-005]`). Sections: The Vision (`[C-001]`), The Feeling (`[C-005]`), The Plan (`[C-002]`), The Name (`[C-004]`), The Gurus (`[C-006]`, `[C-007]`).\n\n## Chapter 2: The Sucker's Roadmap (Flat-Pack Manual)\n- Parts Box (`[C-009]`): 1x Prop Evaluation ($89–$540), 1x TradingView Template, 1x Bracket Order Guide, 0x Instructions for Step 7.\n- Steps 1 to 6 (`[C-008]`).\n- Ends on: `Step 7 was not provided.` (`[C-010]`).\n"), 0644)

	// Phase 2 packet, tables, predraft
	p2Pkt := gate.PhasePacket{
		TopicSlug:     "forex-prop-firms",
		Phase:         2,
		SpineVersion:  "outline.v1@282d710",
		PacketVersion: 1,
		Dials: map[string]interface{}{
			"ch3.charts":        []string{"waterfall", "split bar", "hourly meter", "gap bar"},
			"ch4.cases":         5,
			"ch5.charts":        []string{"log ladder", "funnel", "unit grid"},
			"ch6.verdict":       "Probably won't work",
			"last_posibot_meth": "receipt tape",
		},
		CitedRows: []string{
			"C-001", "C-002", "C-003", "C-009", "C-010",
			"F-001", "F-002", "F-003", "F-004", "F-005", "F-006", "F-007", "F-008", "F-009", "F-010", "F-011", "F-012", "F-013", "F-014",
			"K-001", "K-002", "K-003", "K-004", "K-005", "K-006", "K-007", "K-008", "K-009", "K-010", "K-011", "K-012", "K-013", "K-014", "K-015",
		},
		Ch4Molds: []string{"E2", "H4", "H5", "C1", "C3"},
	}
	p2Bytes, _ := yaml.Marshal(&p2Pkt)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-2", "packet.yaml"), p2Bytes, 0644)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-2", "tables.md"), []byte("# Phase 2 Demolition Tables (Audited Facts & 5 Composite Case Molds)\n\n| Metric | Value | Threshold | Result | Rows |\n| :--- | :--- | :--- | :--- | :--- |\n| **The Stake** | `$1,840` | `<= $1,000` | `MISS` | `F-001, F-004, F-014` |\n| **The Wait** | `14.0 months` | `<= 6.0 months` | `MISS` | `F-009` |\n| **The Odds** | `2.4% at 12mo` | `>= 50.0%` | `MISS` | `F-007, F-008` |\n| **The Payoff** | `$6.40/hr` | `>= $29.76/hr` | `MISS` | `F-003, F-013` |\n"), 0644)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-2", "predraft.md"), []byte("# Phase 2 Pre-Draft Spine: Demolition (Chapters 3, 4, 5 & 6)\n\n## Chapter 3: The Money Pit\n- Opens on the `$0.00` breached evaluation dashboard (`[F-001]`, `[F-004]`). Waterfall chart (`$1,840` median Stake), Split bar (`98.4%` B-book internal server `[F-002]`, `[F-011]`), Hourly meter (`$6.40/hr` vs `$29.76/hr` `[F-003]`, `[F-013]`).\n\n## Chapter 4: Tales from the Trenches (Standalone Aged Amber Autopsy)\n- Zero PosiBot, zero jokes. Five composites: `Marcus Conviction™` (`E2`), `Elena Diligence` (`H4`), `Devon Precision` (`H5`), `Hannah Patience` (`C1`), `Victor Discipline` (`C3`).\n\n## Chapter 5: Exit Liquidity\n- Log ladder (`$7.5T/day` BIS turnover, `93%` institutional `[F-006]`), Attrition funnel (`6.6%` pass Phase 1+2, `2.4%` net-profitable at 12 months `[F-007]`, `[F-008]`), Wait (`14.0 months` `[F-009]`).\n\n## Chapter 6: The Verdict\n- 4-Metric Gate Scorecard (`0 of 4 passed` $\\rightarrow$ `Probably won't work`).\n"), 0644)

	// Phase 3 packet, tables, predraft
	p3Pkt := gate.PhasePacket{
		TopicSlug:     "forex-prop-firms",
		Phase:         3,
		SpineVersion:  "outline.v1@282d710",
		PacketVersion: 1,
		Dials: map[string]interface{}{
			"ch7.pivots":      4,
			"ch7.repurpose":   6,
			"ch7.skills":      4,
			"ch8.honest_way":  4,
			"ch8.do_nots":     6,
			"ch8.exit_steps":  6,
			"ch8.watch_for":   3,
			"ch8.callbacks":   2,
			"companion.segue": 1,
		},
		CitedRows: []string{
			"F-001", "F-003", "F-004", "F-007", "F-008", "F-009", "F-010", "F-013", "F-014",
			"F-015", "F-016", "F-017", "F-018", "F-019", "F-020", "F-021", "F-022", "F-023", "F-024", "F-025", "F-026", "F-027",
			"K-013", "K-015",
		},
		Ch7Pivots: 4,
		Ch8Title:  "For the Stubborn Fool Who Must",
	}
	p3Bytes, _ := yaml.Marshal(&p3Pkt)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-3", "packet.yaml"), p3Bytes, 0644)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-3", "tables.md"), []byte("# Phase 3 Solidarity & Practical Pivot Tables\n\n| Pivot ID | Role | Stake | Wait | Odds (12mo) | Payoff | Gates Passed |\n| :--- | :--- | :--- | :--- | :--- | :--- | :--- |\n| `PIV-01` | SQL & Python Financial BI Data Analyst | `$320` | `4.0 mo` | `78%` | `$44.50/hr` | `4 of 4` |\n| `PIV-02` | Freelance PineScript / Backtesting Contractor | `$150` | `2.5 mo` | `62%` | `$55.00/hr` | `4 of 4` |\n| `PIV-03` | Regulated Series 65 Paraplanner / Ops | `$185` | `3.5 mo` | `68%` | `$38.90/hr` | `4 of 4` |\n| `PIV-04` | Corporate FX Treasury & Hedging Analyst | `$650` | `4.5 mo` | `71%` | `$49.20/hr` | `4 of 4` |\n"), 0644)
	_ = os.WriteFile(filepath.Join(stageDir, "phase-3", "predraft.md"), []byte("# Phase 3 Pre-Draft Spine: Solidarity (Chapters 7, 8 & Appendices)\n\n## Chapter 7: The Pivot\n- Evaluates 4 `[PIVOT: THE PRACTICAL PIVOT]` options (`PIV-01`..`PIV-04`), all passing 4 of 4 gates (`[F-015]`..`[F-018]`).\n- Repurposing matrix (`[F-027]`), Skill Exchange table (`[F-015]`), and `[PAGE COMIC — 3D PAPERCRAFT]`.\n- Closing band ends on: `…And… for the stubborn fool who must…` (`For the Stubborn Fool Who Must`).\n\n## Chapter 8: For the Stubborn Fool Who Must\n- 4 Honest-Way practices (`[F-023]`), Smallest Honest Test & Stop Rules (`[F-026]`), 6 Do-Nots (`[F-025]`), 6 Exit Steps (`[F-019]`..`[F-021]`), Cost/Keep Ledger, 3 Watch-For scams (`[F-022]`), Companion Pamphlet Segue (`$1.99`), and Aftermath Callbacks (`[K-013]`, `[K-015]`).\n"), 0644)

	return nil
}

func buildStage03Phase1(rootDir string) error {
	stageDir := filepath.Join(rootDir, "03_phase1_run")
	assetsDir := filepath.Join(stageDir, "assets")
	promptsDir := filepath.Join(stageDir, "prompts")
	for _, d := range []string{stageDir, assetsDir, promptsDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	var recs []catalog.AssetRecord
	addAsset := func(ch int, slot, slug, method, chartType string, rows []string, svg []byte, prompt string) {
		r, err := catalog.WriteCatalogAsset(assetsDir, 1, 1, ch, slot, slug, 1, "svg", svg, prompt, method, chartType, rows)
		if err == nil {
			recs = append(recs, r)
		}
	}

	addAsset(1, "posibot", "zero-risk-laptop", "signpost", "", []string{"C-001"},
		catalog.RenderPosiBotSVG("TRADE $200K FROM BED!", "(Keep 90% of every pip!)", "signpost", "#0D9488"),
		"PosiBot in Phase 1 warm cream-and-teal palette holding a wooden signpost reading 'TRADE $200K FROM BED!'")
	addAsset(1, "posibot", "london-session-money", "megaphone with shout bubble", "", []string{"C-002"},
		catalog.RenderPosiBotSVG("45 MINUTES A DAY!", "(Central banks hate this one weird sweep!)", "megaphone with shout bubble", "#0D9488"),
		"PosiBot shouting through a brass megaphone with speech bubble '45 MINUTES A DAY!'")
	addAsset(1, "posibot", "no-inventory-needed", "held placard", "", []string{"C-006"},
		catalog.RenderPosiBotSVG("NO CUSTOMERS! NO BOXES!", "(Pure digital currency liquidity!)", "held placard", "#0D9488"),
		"PosiBot holding a glossy placard reading 'NO CUSTOMERS! NO BOXES!'")
	addAsset(2, "posibot", "six-easy-steps", "clipboard", "", []string{"C-008"},
		catalog.RenderPosiBotSVG("JUST FOLLOW STEP 1-6!", "(Allen wrench included! Step 7 sold separately.)", "clipboard", "#2563EB"),
		"PosiBot holding an assembly inspector clipboard reading 'JUST FOLLOW STEP 1-6!'")

	if err := catalog.WriteSnapshot(assetsDir, "forex-prop-firms", recs); err != nil {
		return err
	}

	promptPackYAML := `topic_slug: forex-prop-firms
phase: 1
voice_register: "Corporate Sincerity"
prompts:
  - asset_id: "T01-P1-C01-posibot-zero-risk-laptop_v001"
    chapter: 1
    delivery_method: "signpost"
    headline: "TRADE $200K FROM BED!"
    cited_rows: ["C-001"]
  - asset_id: "T01-P1-C01-posibot-london-session-money_v001"
    chapter: 1
    delivery_method: "megaphone with shout bubble"
    headline: "45 MINUTES A DAY!"
    cited_rows: ["C-002"]
  - asset_id: "T01-P1-C01-posibot-no-inventory-needed_v001"
    chapter: 1
    delivery_method: "held placard"
    headline: "NO CUSTOMERS! NO BOXES!"
    cited_rows: ["C-006"]
  - asset_id: "T01-P1-C02-posibot-six-easy-steps_v001"
    chapter: 2
    delivery_method: "clipboard"
    headline: "JUST FOLLOW STEP 1-6!"
    cited_rows: ["C-008"]
`
	if err := os.WriteFile(filepath.Join(promptsDir, "prompt_pack_phase1.yaml"), []byte(promptPackYAML), 0644); err != nil {
		return err
	}

	msPhase1 := `# Phase 1: Vibrant Hope (Chapters 1 & 2)

## Chapter 1: The Promise
*Format: Full-Color Magazine-Ad Collage · Voice: Deadpan Corporate Sincerity · Evidence: Claim Rows Only*

### Opening Vignette: 11:42 PM on a Tuesday
It is 11:42 PM on a Tuesday ([C-005]). Across the street, your neighbor Greg is ironing a polyester-blend dress shirt so he can wake up at 6:15 AM, scrape gray slush off a Honda Civic, and sit in forty-five minutes of brake-light purgatory for the privilege of asking a regional vice president Named Todd if he can leave twenty minutes early on Friday for a dentist appointment.

You, meanwhile, are sitting on the edge of your sofa in flannel pajama pants, illuminated by the serene sapphire glow of MetaTrader 5 ([C-005], [C-013]). Somewhere over the Atlantic, the Tokyo–London liquidity bridge is waking up. You don't have a warehouse. You don't have employees. You don't even have pants with a zipper on right now. You tap three glass buttons on your phone, bracket a 40-pip XAU/USD (Gold) displacement candle, set your phone on the nightstand, and wake up at 7:30 AM to a **$2,400 payout notification** before your toaster pops ([C-005]).

### Section 1: The Vision — Why Beg for a Raise When You Can Trade Someone Else's $200,000?
For half a century, the **$7.5 trillion-a-day** foreign exchange market ([C-002]) kept a velvet rope across the lobby. Even if you could read a chart like Mozart reading sheet music, pulling real money out of currency pairs required $100,000 of your own post-tax savings sitting in a brokerage vault.

Welcome to the **Funded Trader Revolution** ([C-001]). Modern proprietary evaluation firms have democratized institutional leverage: pass a simple 2-Step digital audition in **14 days**, and they hand you the keys to a **$100,000 to $200,000 funded trading account**—letting you **keep up to 90% of every dollar you extract** with **zero personal capital at risk** ([C-001]).

### Section 2: The Feeling — Zero Boxes, Zero Bosses, Pure Liquidity
Think about every other "side hustle" cluttering your social feed ([C-006]). Dropshipping? You're arguing with a freight forwarder in Shenzhen over defective silicone garlic peelers at 3:00 AM. Agency consulting? You're cold-calling plumbers who hang up on you. Amazon FBA? Your garage is full of unsold yoga mats.

In foreign exchange, **the Japanese Yen never calls customer support demanding a refund** ([C-006]). The British Pound never leaves a one-star Yelp review. Currency pairs never expire on a pallet, never ghost your invoice, and settle in crisp bi-weekly wire transfers every **14 days** ([C-006])—whether your laptop is open in a suburban kitchen or on a glass balcony overlooking Dubai Marina next to a matte-black G-Wagon ([C-013]).

### Section 3: The Plan — 45 Minutes at the London–New York Overlap
Amateur traders fail because they stare at five-minute charts for twelve hours a day like caffeinated raccoons. The Funded Professional does the exact opposite. You only need **45 minutes a day** during the 8:00 AM–11:00 AM New York–London session overlap ([C-002]) to pull **$10,000 a month** out of the $7.5 trillion daily ocean ([C-002]). And the barrier to entry costs less than a weekend bar tab: unlock a $100,000 evaluation today for **$497** (or test the waters with a $10,000 starter account for just **$89**)—and collect a **100% fee refund** bundled right into your very first profit split ([C-003], [C-014]).

### Section 4: The Name — Decoding "Smart Money Concepts" (SMC)
Why did retail traders struggle in the 2010s? Because they were drawing retail trendlines while interbank algorithms hunted their stop-losses for sport. Today, you have the **Smart Money Concepts (SMC) Blueprint** ([C-002], [C-004]). You aren't guessing anymore; you are tracking institutional **Order Blocks**, **Fair Value Gaps**, and the mechanical 9:30 AM **Liquidity Sweep** ([C-004], [C-008]). Once you see where central banks park their orders, taking 5% off the table is as routine as brushing your teeth.

### Section 5: The Gurus & The Funded Hall of Fame
> **[GURU PITCH HALL OF FAME — COMPOSITE]:**  
> *"Bro, the interbank delivery algorithm runs on a 90-minute clock. It has to sweep retail liquidity before it expands. Once you pass your $200,000 evaluation in 14 days, a single 5% month drops $9,000 into your Wise account—and the firm refunds your $497 audition fee on the house. Use code LIQUIDITY20 at checkout before the timer hits zero."* ([C-001], [C-003], [C-013])

> **[CHAPTER 1 VERDICT BAND]:** $200,000 in institutional firepower, 90% profit splits, 45 minutes a morning, and a 100% refundable $497 audition fee ([C-001], [C-002], [C-003]). Honestly, why is anyone still driving to an office?

---

## Chapter 2: The Sucker's Roadmap
*Format: Flat-Pack Assembly Manual · Model: FUNDED-TRADER-200K™ · Voice: Deadpan Corporate Sincerity*

### Box Contents & Required Hardware ([C-009])
Please verify that your starter blister-pack contains the following components before beginning assembly:
- 1× **2-Step Funded Trader Evaluation Credential** ($89 starter / $497–$540 for $100,000 tier; 100% fee refund included upon first payout) ([C-003], [C-009])
- 1× **TradingView Charting Screen & Smart-Money Indicator Overlay** (Pre-colored green and red boxes so you never have to do math) ([C-008], [C-009])
- 1× **MetaTrader 5 Mobile App** (For screenshotting blue numbers from the gym parking lot) ([C-005], [C-013])
- 0× **Instructions for Step 7** *(Omitted from factory packaging)* ([C-010])

### Six-Step Assembly Sequence ([C-008])
1. **STEP 1 — Select Your Capital Tier:** Open the prop firm checkout page, paste your favorite creator's 20% discount code ([C-013]), and claim $100,000 in simulated institutional buying power for $497 ([C-003], [C-008]).
2. **STEP 2 — Snap On the Smart-Money Overlay:** Load the proprietary indicator template ([C-008], [C-009]). *SYSTEM™ TIP: When the shaded rectangle says BULLISH ORDER BLOCK, that means the banks are buying!* ([C-011])
3. **STEP 3 — Await the 9:30 AM Liquidity Sweep:** Pour an espresso at 9:28 AM EST and watch price dip neatly below yesterday's low to "sweep the uninformed retail herd" ([C-008]).
4. **STEP 4 — Insert the 1:5 Risk-to-Reward Bracket:** Risk a polite 1% of the account to capture a clean 5% surge before 10:15 AM ([C-008], [C-011]).
5. **STEP 5 — Clear Phase 1 (+8%) and Phase 2 (+5%):** Execute two clean trades, pass both stages in 14 days ([C-001], [C-008]), and unlock live funded status.
6. **STEP 6 — Frame Your Certificate & Withdraw 90%:** Download your gold-embossed **Funded Trader Certificate** ([C-008], [C-013]), post it to Instagram with a caption about "discipline," and schedule bi-weekly 90% wire transfers ([C-006]).

### Troubleshooting & Missing Assembly Diagram ([C-010])
> **[ STEP 7 WAS NOT PROVIDED. ]**  
> **Step 7 was not provided.** Specifically, the manual forgot to mention how you are supposed to clear an +8% Phase 1 hurdle and a +5% Phase 2 hurdle inside an internal B-book server while chained to a **5% intraday trailing drawdown tripwire** and a negative-expectancy spread-plus-commission drag that mathematically demands a **64% raw win rate** before the evaluation timer detonates ([C-010]).
`
	return os.WriteFile(filepath.Join(stageDir, "manuscript_phase1.md"), []byte(msPhase1), 0644)
}

func buildStage04Phase2(rootDir string) error {
	stageDir := filepath.Join(rootDir, "04_phase2_run")
	assetsDir := filepath.Join(stageDir, "assets")
	promptsDir := filepath.Join(stageDir, "prompts")
	for _, d := range []string{stageDir, assetsDir, promptsDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	var recs []catalog.AssetRecord
	addAsset := func(ch int, slot, slug, method, chartType string, rows []string, svg []byte, prompt string) {
		r, err := catalog.WriteCatalogAsset(assetsDir, 1, 2, ch, slot, slug, 1, "svg", svg, prompt, method, chartType, rows)
		if err == nil {
			recs = append(recs, r)
		}
	}

	// Chapter 3 Charts & PosiBot
	addAsset(3, "chart", "stake-waterfall", "", "waterfall", []string{"F-001", "F-004", "C-003"},
		catalog.RenderChartSVG("waterfall", "What Does a '$497' Evaluation Actually Cost?", "MEDIAN STAKE: $1,840", "Industry Prop Cohort Study (420k Accounts) [F-001, F-004]",
			[]string{"Advertised Starter Fee", "4.2x Challenge & Reset Fees", "Charting / VPS / Indicators", "Total Median Cash Stake"},
			[]float64{340, 1428, 412, 1840},
			[]string{"$340 (Quoted)", "$1,428 (4.2 tries)", "+$412 (Tools)", "$1,840 TOTAL STAKE"}, "#DC2626"),
		"Waterfall chart comparing advertised $340-$497 evaluation fee against $1,840 median Stake")

	addAsset(3, "chart", "bbook-split-bar", "", "split bar", []string{"F-002", "F-011"},
		catalog.RenderChartSVG("split bar", "Where Do Your Prop Evaluation Orders Actually Go?", "98.4% INTERNAL B-BOOK", "CFTC Enforcement Accounting [F-002, F-011]",
			[]string{"Routed to Live Market (A-Book)", "Internal Simulated Server (B-Book)", "Funded Accounts Terminated <Payout #2"},
			[]float64{1.6, 98.4, 91.6},
			[]string{"1.6% Real Market", "98.4% Simulated B-Book", "91.6% Terminated Early"}, "#B91C1C"),
		"Split bar showing 98.4% of prop trades stay inside an internal simulated B-book server")

	addAsset(3, "chart", "payoff-hourly-meter", "", "hourly meter", []string{"F-003", "F-013"},
		catalog.RenderChartSVG("hourly meter", "What Do the Top 2.4% Survivors Earn Per Hour?", "$6.40/HR vs $29.76/HR", "Prop Cohort Study & US BLS OEWS [F-003, F-013]",
			[]string{"Net-Unprofitable Starters (97.6%)", "Surviving Prop Traders (Top 2.4%)", "US Median Hourly Wage (BLS)", "SQL / Financial BI Analyst (BLS)"},
			[]float64{1, 6.40, 29.76, 44.50},
			[]string{"<= $0.00/hr", "$6.40/hr (26.5h/wk)", "$29.76/hr Benchmark", "$44.50/hr Pivot"}, "#DC2626"),
		"Hourly meter comparing survivors' $6.40/hr payoff against the $29.76/hr US median wage")

	addAsset(3, "posibot", "reset-discount-code", "whiteboard easel", "", []string{"F-001"},
		catalog.RenderPosiBotSVG("FIFTH RESET IS 15% OFF!", "(Median trader buys 4.2 challenges: $1,840 spent)", "whiteboard easel", "#DC2626"),
		"PosiBot pointing at a whiteboard easel reading 'FIFTH RESET IS 15% OFF!'")

	// Chapter 4 Caricatures & Exhibits (5 distinct molds, zero PosiBot, Aged Amber #EFE5CE)
	addAsset(4, "caricature", "case1-e2-marcus", "", "", []string{"K-001", "K-002", "K-003"},
		catalog.RenderCaseCaricatureSVG("E2", "Marcus Conviction™", "+$54k Affiliate / $0 Trading", "88k Followers Click 20% Code", "Rented Huracán + Ring Light", "20% CPA on Follower Resets", "$0 Trading Net vs. +$54,000 Affiliate CPA [K-001..K-003]"),
		"Editorial silhouette caricature dramatizing Marcus Conviction filming a Funded Trader vlog while earning $54,000 from affiliate CPA links and $0 from trading")

	addAsset(4, "caricature", "case2-h4-elena", "", "", []string{"K-004", "K-005", "K-006"},
		catalog.RenderCaseCaricatureSVG("H4", "Elena Diligence", "-$4,110 Spent / $0 Out", "620 Hours on the Escalator", "$297 Course -> $997 Academy", "$149/mo VIP Discord + 5x Resets", "$4,110 Cash Spent / 620 Hours / $0 Withdrawn [K-004..K-006]"),
		"Editorial silhouette caricature dramatizing Elena Diligence climbing an endless escalator of paid courses, VIP Discord signals, and evaluation resets")

	addAsset(4, "caricature", "case3-h5-devon", "", "", []string{"K-007", "K-008", "K-009"},
		catalog.RenderCaseCaricatureSVG("H5", "Devon Precision", "-$2,923 Tool Tax / Banned", "6-Monitor Indicator Cockpit", "$59 Charts + $68 LuxAlgo + VPS", "$499 Shared-IP EA Bot Ban", "$2,923 Tool & Reset Tax / Terminated: Shared EA IP [K-007..K-009]"),
		"Editorial silhouette caricature dramatizing Devon Precision wired into a 6-monitor cockpit of monthly subscriptions and a banned Telegram EA bot")

	addAsset(4, "caricature", "case4-c1-hannah", "", "", []string{"K-010", "K-011", "K-012"},
		catalog.RenderCaseCaricatureSVG("C1", "Hannah Patience", "+$1,250 Net / $1.11 Per Hour", "14 Months · 1,120 Screen Hours", "0.25% Risk · Passed Phase 1 & 2", "$3,400 Payout - $2,150 Fees", "14 Months / 1,120 Hours / +$1,250 Net = $1.11/Hr [K-010..K-012]"),
		"Editorial silhouette caricature dramatizing Hannah Patience after 14 months and 1,120 disciplined hours yielding $1.11 per hour")

	addAsset(4, "caricature", "case5-c3-victor", "", "", []string{"K-013", "K-014", "K-015"},
		catalog.RenderCaseCaricatureSVG("C3", "Victor Discipline", "+$5,100 Cushion Frozen", "Payout Denied Overnight", "+$5,100 Funded Profit Built", "30% Consistency / MT5 Freeze", "+$5,100 Funded Cushion -> $0 Payout Frozen [K-013..K-015]"),
		"Editorial silhouette caricature dramatizing Victor Discipline having a +$5,100 funded profit frozen overnight by a retroactive consistency rule")

	addAsset(4, "exhibit", "case1-e2-gap-bar", "", "gap bar", []string{"K-001", "K-002", "K-003"},
		catalog.RenderChartSVG("gap bar", "Mold E2 (Marcus Conviction™): Trading vs. Affiliate Income", "$54,000 REFERRAL EDGE", "Composite of 3 Documented Cases [K-001..K-003]",
			[]string{"Median Net Trading Payout", "Trading Challenge Stake", "Affiliate Referral Commissions"},
			[]float64{100, 600, 54000},
			[]string{"$0 Net Trading", "$600 Stake", "$54,000 Affiliate CPA"}, "#D97706"),
		"Chapter 4 Mold E2 gap bar comparing $0 net trading profit against $54,000 affiliate referral commissions")

	addAsset(4, "exhibit", "case2-h4-split-bar", "", "split bar", []string{"K-004", "K-005", "K-006"},
		catalog.RenderChartSVG("split bar", "Mold H4 (Elena Diligence): The $4,110 Upsell Ladder", "$4,110 SPENT / $0 OUT", "Composite of 3 Documented Cases [K-004..K-006]",
			[]string{"Entry Course & Mentorship", "VIP Discord Signal Rooms", "5x Evaluation Resets", "Total Withdrawn Payouts"},
			[]float64{1094, 1341, 1675, 10},
			[]string{"$1,094 Courses", "$1,341 VIP Signals", "$1,675 Resets", "$0.00 Withdrawn"}, "#B45309"),
		"Chapter 4 Mold H4 split bar showing $4,110 spent across courses, signals, and resets")

	addAsset(4, "exhibit", "case3-h5-waterfall", "", "waterfall", []string{"K-007", "K-008", "K-009"},
		catalog.RenderChartSVG("waterfall", "Mold H5 (Devon Precision): The $2,923 Tool Tax", "7 MONTHS / $0 PAYOUT", "Composite of 3 Documented Cases [K-007..K-009]",
			[]string{"Charting + Orderflow Feeds", "Low-Latency London VPS", "Automated Prop-Pass EA Bot", "Evaluation Reset Fees"},
			[]float64{889, 315, 499, 1220},
			[]string{"$889 Feeds", "$315 VPS", "$499 EA Bot", "$1,220 Resets ($2,923 Total)"}, "#92400E"),
		"Chapter 4 Mold H5 waterfall chart showing $2,923 tool and reset tax")

	addAsset(4, "exhibit", "case4-c1-timeline", "", "timeline", []string{"K-010", "K-011", "K-012"},
		catalog.RenderChartSVG("timeline", "Mold C1 (Hannah Patience): 14 Months & 1,120 Hours to Net $1.11/Hr", "14-MONTH RUNWAY GAP", "Composite of 3 Documented Cases [K-010..K-012]",
			[]string{"Months 1-9 (3 Evaluation Cycles)", "Months 10-13 (Payout Holds)", "Month 14 Net Profit ($1,250)", "Effective Hourly Rate (1,120h)"},
			[]float64{900, 1300, 1250, 111},
			[]string{"-$2,150 Fees", "0 Payouts", "+$1,250 Net", "$1.11 / Hour Worked"}, "#78350F"),
		"Chapter 4 Mold C1 timeline showing 14 months and 1,120 hours to earn $1.11/hr")

	addAsset(4, "exhibit", "case5-c3-cliff", "", "before/after cliff", []string{"K-013", "K-014", "K-015"},
		catalog.RenderChartSVG("before/after cliff", "Mold C3 (Victor Discipline): Overnight Rule Change & Payout Freeze", "$5,100 -> $0 FROZEN", "Composite of 3 Documented Cases [K-013..K-015]",
			[]string{"Funded Account Profit Cushion", "After Retroactive Rule / Freeze", "6-Month Chargeback Recovery"},
			[]float64{5100, 10, 540},
			[]string{"+$5,100 Funded Profit", "$0.00 (Denied/Frozen)", "$540 Card Chargeback"}, "#991B1B"),
		"Chapter 4 Mold C3 before/after cliff showing $5,100 funded profit dropping to $0 on overnight rule change")

	// Chapter 5 Charts & PosiBot
	addAsset(5, "chart", "fx-market-log-ladder", "", "log ladder", []string{"F-006", "C-002"},
		catalog.RenderChartSVG("log ladder", "Who Actually Trades the $7.5 Trillion-a-Day FX Market?", "93% INSTITUTIONAL", "BIS Triennial Central Bank Survey [F-006]",
			[]string{"Tier-1 Banks & Institutional Dealers", "Central Banks & Hedge Funds", "All Retail Spot FX Worldwide", "Prop Challenge Starters Net-Profitable"},
			[]float64{6975, 525, 210, 5},
			[]string{"$6.97T/day (93.0%)", "$525B/day (7.0%)", "<2.8% Retail Slice", "2.4% of Prop Starters"}, "#1E40AF"),
		"Log ladder showing 93% of the $7.5T/day FX market is institutional dealer volume")

	addAsset(5, "chart", "prop-attrition-funnel", "", "funnel", []string{"F-007", "F-008", "F-010"},
		catalog.RenderChartSVG("funnel", "Out of 1,000 Prop Challenge Starters, How Many Make It to Month 12?", "24 OF 1,000 (2.4%)", "Prop Evaluation Cohort Study (420k Accounts) [F-007, F-008, F-010]",
			[]string{"Start Phase 1 Evaluation", "Pass Phase 1 + Phase 2 (6.6%)", "Survive to 2nd Payout (0.55%)", "Net-Profitable at 12 Months (2.4%)"},
			[]float64{1000, 66, 32, 24},
			[]string{"1,000 Starters (100%)", "66 Pass Evaluation (6.6%)", "32 Reach 2nd Split", "24 Net-Profitable (2.4%)"}, "#DC2626"),
		"Attrition funnel from 1,000 prop starters down to 24 net-profitable at 12 months")

	addAsset(5, "posibot", "gladiator-liquidity", "sandwich board", "", []string{"F-006", "F-011"},
		catalog.RenderPosiBotSVG("YOU ARE THE LIQUIDITY!", "(93% institutional sharks · 98.4% B-book server)", "sandwich board", "#DC2626"),
		"PosiBot wearing a sandwich board reading 'YOU ARE THE LIQUIDITY!'")

	// Chapter 6 Scorecard & PosiBot
	addAsset(6, "chart", "four-metric-scorecard", "", "scorecard", []string{"F-001", "F-003", "F-004", "F-007", "F-008", "F-009", "F-013", "F-014"},
		catalog.RenderChartSVG("scorecard", "Chapter 6 Gate Scorecard: Forex Trading & Prop Firms", "0 OF 4 GATES PASSED", "Audited Gate Computation [F-001..F-014]",
			[]string{"1. THE STAKE (<= $1,000 savings)", "2. THE WAIT (<= 6.0 months)", "3. THE ODDS (>= 50% at 12mo)", "4. THE PAYOFF (>= $29.76/hr)"},
			[]float64{184, 140, 24, 64},
			[]string{"$1,840 Spent [MISS]", "14.0 Months [MISS]", "2.4% Profitable [MISS]", "$6.40 / Hour [MISS]"}, "#991B1B"),
		"Chapter 6 4-Metric Gate Scorecard showing 0 of 4 gates passed")

	addAsset(6, "posibot", "verdict-receipt-tape", "receipt tape", "", []string{"F-004", "F-008"},
		catalog.RenderPosiBotSVG("PROBABLY WON'T WORK!", "(0 of 4 gates cleared · Keep your $1,840)", "receipt tape", "#991B1B"),
		"PosiBot holding a long thermal receipt tape reading 'PROBABLY WON'T WORK!'")

	if err := catalog.WriteSnapshot(assetsDir, "forex-prop-firms", recs); err != nil {
		return err
	}

	promptPackP2 := `topic_slug: forex-prop-firms
phase: 2
voice_register: "Sardonic Authority (Ch 3, 5, 6) / Forensic Zero-Joke Autopsy (Ch 4)"
ch4_posibot_count: 0
prompts:
  - asset_id: "T01-P2-C03-chart-stake-waterfall_v001"
    chapter: 3
    chart_type: "waterfall"
    cited_rows: ["F-001", "F-004", "C-003"]
  - asset_id: "T01-P2-C04-exhibit-case1-e2-gap-bar_v001"
    chapter: 4
    chart_type: "gap bar"
    cited_rows: ["K-001", "K-002", "K-003"]
  - asset_id: "T01-P2-C05-chart-prop-attrition-funnel_v001"
    chapter: 5
    chart_type: "funnel"
    cited_rows: ["F-007", "F-008", "F-010"]
  - asset_id: "T01-P2-C06-chart-four-metric-scorecard_v001"
    chapter: 6
    chart_type: "scorecard"
    cited_rows: ["F-001", "F-003", "F-004", "F-007", "F-008", "F-009", "F-013", "F-014"]
`
	if err := os.WriteFile(filepath.Join(promptsDir, "prompt_pack_phase2.yaml"), []byte(promptPackP2), 0644); err != nil {
		return err
	}

	msPhase2 := `# Phase 2: Demolition (Chapters 3, 4, 5 & 6)

## Chapter 3: The Money Pit
*Voice: Sardonic Authority · Gate Figures Introduced: The Stake & The Payoff*

### The Payoff of the Phase 1 Foreshadow: The Dashboard Reads $0.00
In Chapter 1, it was 11:42 PM and you were softly caressing a $2,400 gold trade from your couch ([C-005]). Let's turn the studio lights on.

It is 9:31 AM on Day Eleven of your $100,000 evaluation. A U.S. CPI print just dropped, the spread on XAU/USD yawned open like an alligator's jaw, and your screen is flashing a banner the color of an arterial bleed: **ACCOUNT BREACHED — INTRADAY EQUITY DRAWDOWN EXCEEDED ($0.00 ELIGIBLE FOR WITHDRAWAL)**. Thirty-eight seconds later—before your heart rate has even dropped below 120—an automated email lands in your inbox offering a **"15% Loyalty Reload Discount"** on a fresh **$340 Reset Fee** ([F-001]).

### 1. Testing the '$89–$497 Entry Cost' Claim (The Stake = $1,840)
The sales page sold you a one-time refundable audition ticket ([C-003]). In the real world, nobody buys just one prop challenge for the exact same reason casino architecture doesn't put a clock on the wall. Across **420,000 tracked retail prop-firm evaluations**, the median participant buys **4.2 challenges or resets** at a median fee of **$340** a pop ($1,428 in direct challenge tuition), plus **$412** in TradingView upgrades, low-latency London VPS rentals, and neon indicator plugins—putting the audited median **Stake at $1,840** before a trader either sees their first net-profitable month or walks away ([F-001], [F-004]).

For context, the Federal Reserve reports that the typical U.S. household's entire liquid emergency cash buffer is **$1,000**, and 37% of adults cannot cover a $400 car repair without borrowing ([F-014]). The "zero personal capital at risk" pitch ([C-001]) quietly extracts **184% of a typical family's emergency fund** via non-refundable service fees ([F-004], [F-014]).

### 2. Where Your Money Actually Goes: Welcome to the B-Book Arcade
Here is the dirty secret that turns every "Institutional Funded Trader" bio on Instagram into comedy: **there is no $200,000 account** ([C-001], [F-011]).

When the CFTC unsealed its enforcement complaint against one of the largest retail prop evaluation firms on earth ([F-002]), forensic investigators found the firm had vacuumed up **$310 million in assessment and reset fees from over 135,000 retail customers** in 22 months—while routing **98.4% of all customer trades to an internal simulated B-book server** ([F-002], [F-011]). You are not trading in the $7.5 trillion interbank market ([C-002]). You are playing a multiplayer video game inside the prop firm's private database, where the house pays the tiny sliver of winners out of the reset fees of the 93.4% who just blew up ([F-007], [F-011]).

> **[REALITY CHECK — THE B-BOOK INCENTIVE ENGINE]:**  
> In a live brokerage, the broker makes money when you survive and keep trading volume. In a B-book prop evaluation (98.4% simulated execution, [F-011]), a trader who passes and withdraws money is a direct balance-sheet liability. The house's profit margin *requires* your evaluation to detonate before Payout #2 (91.6% termination rate, [F-010]).

### 3. Testing the '$10,000/Month in 45 Minutes' Claim (The Payoff = $6.40/hr)
Across regulated retail FX and CFD brokers in the UK and EU, mandatory disclosures show **74% to 89%** of retail accounts lose money and **94.2%** earn less than minimum wage ([F-005]). In the U.S., CFTC RFED filings show **72%** of retail forex accounts are unprofitable, burning through median deposits in **4.5 months** ([F-012]).

And what if you beat the 97.6% first-year attrition wall ([F-008])? Among the **2.4%** of prop evaluation starters who are net-profitable at 12 months, logged screen, charting, and session-waiting time averages **26.5 hours per week** ([F-003])—yielding a median net **Payoff of $6.40 per hour** ([F-003]). That is less than a quarter of the **$29.76/hr** U.S. median hourly wage ([F-013]), and roughly half what you'd make handing out warm pretzels at an airport kiosk—except the pretzel kiosk doesn't charge you $340 every time you drop a mustard packet ([F-001]).

---

## Chapter 4: Tales from the Trenches
*Format: Standalone Forensic Case-File Insert (#EFE5CE Aged Amber Stock) · Zero PosiBot · Zero Jokes ("The numbers are the joke")*

### CASE FILE #01 · MOLD E2 (The Built-In Audience) — "Marcus Conviction™"
- **Composite Stamp:** *Composite of 3 documented cases ([K-001], [K-002], [K-003])*
- **Life Position & Disposition:** Full-time short-form video creator (52,000 to 145,000 followers); ambitious, camera-ready, eager to monetize financial content.
- **Sold By:** Claim [C-001] (*"$200,000 funded account in 14 days — 90% profit split"*).
- **What It Took (Medians):** $600 in challenge fees over 7 months; $0 net trading profit ([K-003]). **The Unmentioned Edge:** An existing audience of 88,000 viewers using a 20% affiliate referral link ([F-024], [K-002]).
- **Chronological Beats:**
  1. **Month 1:** Purchased a $100,000 evaluation ($540–$600) and filmed a high-production "Day in the Life of a Funded Trader" vlog with an affiliate link in the description ([K-001], [K-003]).
  2. **Month 3:** Breached two live funded accounts (-$1,900 net trading result in K-001), but viewer signups generated $108 per referral ([F-024]).
  3. **Month 7:** Collected a median of **$54,000 in affiliate CPA commissions** ($41,200 to $68,400 across the three cases) while lifetime net trading profit sat at **$0** ([K-001], [K-002], [K-003]).
- **System Mechanism:** An existing following meant never needing to win a trade; the referral commission paid out on every viewer purchase.
- **Outcome:** **WON** *(via affiliate distribution, not trading)*.

### CASE FILE #02 · MOLD H4 (The Upsell Ladder) — "Elena Diligence"
- **Composite Stamp:** *Composite of 3 documented cases ([K-004], [K-005], [K-006])*
- **Life Position & Disposition:** Mid-career shift worker (logistics coordinator / registered nurse / retail manager); methodical, trusting of structured curricula.
- **Sold By:** Claim [C-002] (*"45 minutes a day during the London-NY overlap using Smart Money Concepts"*).
- **What It Took (Medians):** **$4,110 spent**, **9 months**, **620 hours**, **$0 payout** ([K-004], [K-005], [K-006]).
- **Chronological Beats:**
  1. **Month 1:** Enrolled in a $97–$297 introductory forex workshop that framed indicator losses as "missing institutional order-flow context."
  2. **Month 3:** Upgraded to a $997–$1,497 mentorship academy and a $149/month VIP Discord signal room ($1,341 over 9 months).
  3. **Months 5–9:** Purchased 5 consecutive prop evaluation challenges and resets ($1,675 median), exhausting $4,110 in savings before exiting with $0 withdrawn ([K-004], [K-006]).
- **System Mechanism:** Each course tier diagnosed the previous tier's losses as a reason to buy the next subscription; the educational funnel earned from the learner at every rung.
- **Outcome:** **EXITED**.

### CASE FILE #03 · MOLD H5 (The Tool Tax) — "Devon Precision"
- **Composite Stamp:** *Composite of 3 documented cases ([K-007], [K-008], [K-009])*
- **Life Position & Disposition:** Technical specialist (mechanical drafter / IT helpdesk / warehouse supervisor); systems-minded, convinced better software eliminates variance.
- **Sold By:** Claim [C-009] (*"Load our proprietary Smart-Money Indicator template"*).
- **What It Took (Medians):** **$2,923 spent**, **7 months**, **490 hours**, **$0 payout** ([K-007], [K-008], [K-009]).
- **Chronological Beats:**
  1. **Month 1:** Subscribed to TradingView Premium ($59/mo), LuxAlgo ($68/mo), and a London low-latency VPS ($45/mo) ([K-007]).
  2. **Month 4:** Purchased a $499 "Guaranteed Prop-Pass Gold EA Bot" and trade copier after two manual evaluation breaches ([K-007], [K-008]).
  3. **Month 7:** Had the fourth evaluation account ($1,220 in cumulative resets) terminated automatically for "prohibited shared EA signature / IP clustering" with $0 payout ([K-008], [K-009]).
- **System Mechanism:** Subscriptions, data feeds, and automation licenses started billing before dollar one of revenue and compounded every reset.
- **Outcome:** **EXITED**.

### CASE FILE #04 · MOLD C1 (The Runway Gap) — "Hannah Patience"
- **Composite Stamp:** *Composite of 3 documented cases ([K-010], [K-011], [K-012])*
- **Life Position & Disposition:** Quantitative hobbyist / part-time bookkeeper; highly disciplined, conservative risk-sizer (0.25% per trade).
- **Sold By:** Claim [C-008] (*"Hit your 8% Phase 1 and 5% Phase 2 targets and withdraw 90% bi-weekly"*).
- **What It Took (Medians):** **$2,150 spent**, **14 months**, **1,120 hours**, **$3,400 gross payout / $1,250 net profit ($1.11/hr)** ([K-010], [K-011], [K-012]).
- **Chronological Beats:**
  1. **Months 1–8:** Cycled through 3 to 4 evaluations ($2,150 median fees) because conservative 0.25% sizing kept timing out or hitting minor drawdown resets before reaching +8% + +5%.
  2. **Months 9–12:** Passed Phase 1 and Phase 2, then navigated minimum trading-day requirements and 14-day payout waiting periods.
  3. **Month 14:** Withdrew $3,400 gross ($1,250 net after $2,150 fees) across 1,120 logged hours—an effective wage of **$1.11 per hour** ([K-010], [K-012]).
- **System Mechanism:** Even when a disciplined trader passes the evaluation, the time required to clear multi-stage hurdles and recover sunk fees reduces hourly pay far below a living wage.
- **Outcome:** **EXITED**.

### CASE FILE #05 · MOLD C3 (The Rule Change) — "Victor Discipline"
- **Composite Stamp:** *Composite of 3 documented cases ([K-013], [K-014], [K-015])*
- **Life Position & Disposition:** Experienced swing trader / night-shift machinist / dental hygienist; cautious, rule-abiding, patient.
- **Sold By:** Claim [C-006] (*"Currency pairs never ask for refunds and pay out in pure cash every 14 days"*).
- **What It Took (Medians):** **$920 spent**, **5 months**, **380 hours**, **$5,100 funded profit frozen -> $0 payout** ([K-013], [K-014], [K-015]).
- **Chronological Beats:**
  1. **Months 1–4:** Passed Phase 1 and Phase 2 on the first or second attempt ($920 median total spend) and built a $4,200 to $6,800 funded profit balance ($5,100 median).
  2. **Month 5:** Submitted a withdrawal request; the firm retroactively applied a "30% single-trade consistency rule," widened CPI news slippage by 18 pips, or lost its offshore MT5 white-label license ([K-013], [K-014], [K-015]).
  3. **Month 5 Close:** Account access frozen with $0 payout disbursed ([K-015]).
- **System Mechanism:** In an unregulated B-book contract, the house writes the rules, controls the price feed, and can revoke payout eligibility overnight.
- **Outcome:** **EXITED**.

---

## Chapter 5: Exit Liquidity
*Voice: Sardonic Authority · Gate Figures Introduced: The Odds & The Wait*

### 1. Paying Off the '$7.5 Trillion Market' Claim ([C-002] -> [F-006])
Every forex webinar on earth opens with the same reverence: *"The foreign exchange market trades **$7.5 trillion a day**—you just need a microscopic crumb!"* ([C-002], [F-006]).

Let's read the actual Bank for International Settlements survey ([F-006]) that number came from. **93% of that $7.5 trillion daily volume is institutional dealer, central bank, and Quantitative Hedge Fund flow** ([F-006])—Citadel, JPMorgan, and Renaissance running fiber-optic cables directly into matching engines in Equinix LD4 (London) and NY4 (Secaucus). All retail spot traders on the planet combined account for **less than 2.8%** ([F-006]). Walking into spot FX with a $68/month oscillator ([K-007]) to "take liquidity from the banks" ([C-004]) is like swimming into a Great White shark sanctuary with a butter knife taped to a pool noodle because you watched a 12-minute YouTube video on marine biology.

### 2. Paying Off Missing Step 7: The Drawdown Woodchipper ([C-010] -> [F-007], [F-010])
Why was Step 7 missing from Chapter 2's flat-pack brochure ([C-010])? Because when a prop firm advertises a "$100,000 account with a 10% maximum drawdown and a 5% daily drawdown" ([C-001], [F-007]), you do not have $100,000 of buying power. **You have a $5,000 account wearing a $100,000 Halloween costume** ([F-007]).

Asking a trader to generate +8% ($8,000) in Phase 1 and +5% ($5,000) in Phase 2 without ever losing $5,000 in a single day is actually asking them to generate a **+260% return on their true $5,000 daily risk allowance** while paying spread and commission markups that force a **64% raw win-rate hurdle** ([C-010]). Across 420,000 tracked challenges, **93.4% blow up during the evaluation** (only 6.6% pass Phase 1+2, [F-007]). And for the 6.6% who get their shiny JPEG certificate ([C-013]), asymmetric slippage scripts and daily drawdown tripwires **terminate 91.6% of funded accounts before a second payout** ([F-010]).

### 3. The Odds (2.4%) and The Wait (14.0 Months)
Put the entire funnel together from Day 1 to Month 12: out of every 1,000 people who buy a challenge, **only 24 (2.4%) are net-profitable at 12 months** after subtracting challenge fees, resets, tools, and frozen payouts (**The Odds**, [F-008]). And for those 24 survivors, the median time to reach consistent net profitability (**The Wait**) is **14.0 months** ([F-009])—slightly longer than the "14 days" printed on the checkout button ([C-001]).

---

## Chapter 6: The Verdict
*Voice: Sardonic Authority · The 4-Metric Gate Scorecard*

### The 4-Metric Gate Scorecard (outline-spine.v1.yaml §5.3)
We don't grade side hustles on vibes, rented Lamborghinis, or cropped MetaTrader screenshots ([C-013]). Every topic in the Y-IT series runs through the same four mechanical gates:

| Gate Metric | The Question | Audited Topic Figure | Gate Threshold | Result |
| :--- | :--- | :--- | :--- | :---: |
| **1. The Stake** | What must I put at risk before I know? | **$1,840** (4.2 challenges + tools [F-001, F-004]) | <= $1,000 US emergency savings ([F-014]) | **MISS** |
| **2. The Wait** | How long until it pays? | **14.0 months** ([F-009]) | <= 6.0 months | **MISS** |
| **3. The Odds** | How many make it at 12 months? | **2.4%** (6.6% pass evaluation [F-007, F-008]) | >= 50.0% (1 in 2) | **MISS** |
| **4. The Payoff** | What does winning pay per hour? | **$6.40 / hour** (26.5 hrs/wk [F-003]) | >= $29.76 / hour US median wage ([F-013]) | **MISS** |

### Who Still Could?
Look back at Chapter 4. The only person who walked out of the building with real money was **Marcus Conviction™ (Mold E2 — The Built-In Audience, [K-001..K-003])**, who netted $0 from trading currencies and **$54,000 from his 20% affiliate referral link** ([F-024], [K-003]). In a gold rush built on a 98.4% B-book simulator ([F-011]), the only reliable profit margin belongs to the guy selling $340 cardboard shovels at the entrance ([F-001]).

> **[CHAPTER 6 FINAL VERDICT BAND]:** **0 of 4 gates passed.** You are risking $1,840 ([F-004]) and waiting 14.0 months ([F-009]) for a 2.4% shot ([F-008]) at making $6.40 an hour ([F-003]). Computed Series Verdict: **Probably won't work.**
`
	return os.WriteFile(filepath.Join(stageDir, "manuscript_phase2.md"), []byte(msPhase2), 0644)
}

func buildStage05Phase3(rootDir string) error {
	stageDir := filepath.Join(rootDir, "05_phase3_run")
	assetsDir := filepath.Join(stageDir, "assets")
	promptsDir := filepath.Join(stageDir, "prompts")
	for _, d := range []string{stageDir, assetsDir, promptsDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	var recs []catalog.AssetRecord
	addAsset := func(ch int, slot, slug, method, chartType string, rows []string, svg []byte, prompt string) {
		r, err := catalog.WriteCatalogAsset(assetsDir, 1, 3, ch, slot, slug, 1, "svg", svg, prompt, method, chartType, rows)
		if err == nil {
			recs = append(recs, r)
		}
	}

	addAsset(7, "comic", "papercraft-diorama-strip", "", "", []string{"F-016", "F-017"},
		catalog.RenderPapercraftComicSVG("Forex Prop Challenge to Practical Pivot"),
		"4-panel 3D papercraft diorama comic strip showing PosiBot trading a breached prop terminal for a 4-of-4 gate-passing data analyst desk")

	addAsset(7, "chart", "pivot-payoff-comparison", "", "range strip", []string{"F-003", "F-013", "F-015", "F-016", "F-017"},
		catalog.RenderChartSVG("range strip", "What Happens When You Sell the Same Quantitative Skill to Employers?", "4 OF 4 GATES PASSED", "US BLS OEWS & Contract Benchmarks [F-015..F-018]",
			[]string{"Prop Challenge Survivor (2.4%)", "Regulated Series 65 Paraplanner", "SQL / Financial BI Data Analyst", "Corporate FX Treasury Specialist", "Freelance PineScript/Python Dev"},
			[]float64{6.40, 38.90, 44.50, 49.20, 55.00},
			[]string{"$6.40/hr (0/4 Gates)", "$38.90/hr (4/4 Gates)", "$44.50/hr (4/4 Gates)", "$49.20/hr (4/4 Gates)", "$55.00/hr (4/4 Gates)"}, "#16A34A"),
		"Range strip comparing $6.40/hr prop trading payoff against 4 gate-passing pivots ($38.90/hr to $55.00/hr)")

	addAsset(7, "posibot", "keep-the-day-job", "flip chart", "", []string{"F-016"},
		catalog.RenderPosiBotSVG("GET PAID FOR THE SQL!", "($44.50/hr BI Analyst · 4 of 4 Gates Passed)", "flip chart", "#15803D"),
		"PosiBot standing beside a flip chart reading 'GET PAID FOR THE SQL!'")

	addAsset(8, "posibot", "ninety-day-paper-test", "picket sign", "", []string{"F-023", "F-026"},
		catalog.RenderPosiBotSVG("ZERO FEES FOR 90 DAYS!", "(Prove 100 trades on demo before paying $1)", "picket sign", "#15803D"),
		"PosiBot holding a picket sign reading 'ZERO FEES FOR 90 DAYS!'")

	addAsset(8, "diagram", "research-machine-left", "", "", []string{"F-001", "F-006", "F-013", "F-014"},
		[]byte(catalog.RenderResearchMachineLeftSVG()),
		"Dual-page Y-IT Research Machine visualization spread (Left Verso Plate): 5-layer context protocol, 9-family router, multi-index live harvest, columnar OLAP & 12GB VRAM / elastic cloud GPU bench, SHA-256 DOM lock, and 15-field atomizer")

	addAsset(8, "diagram", "research-machine-right", "", "", []string{"F-001", "F-003", "F-007", "F-008", "F-009"},
		[]byte(catalog.RenderResearchMachineRightSVG()),
		"Dual-page Y-IT Research Machine visualization spread (Right Recto Plate): RM-10v2 magnitude convergence & range-first sieve, 21->13+1 empirical mold census, BM-01..08 beneficial mechanisms, MON-01..08 monetization, and compiler gates G1-G7")

	if err := catalog.WriteSnapshot(assetsDir, "forex-prop-firms", recs); err != nil {
		return err
	}

	promptPackP3 := `topic_slug: forex-prop-firms
phase: 3
voice_register: "Constructive Solidarity (Ch 7) / Quiet Reckoning (Ch 8)"
prompts:
  - asset_id: "T01-P3-C07-comic-papercraft-diorama-strip_v001"
    chapter: 7
    slot: "comic"
    cited_rows: ["F-016", "F-017"]
  - asset_id: "T01-P3-C07-chart-pivot-payoff-comparison_v001"
    chapter: 7
    chart_type: "range strip"
    cited_rows: ["F-003", "F-013", "F-015", "F-016", "F-017"]
  - asset_id: "T01-P3-C07-posibot-keep-the-day-job_v001"
    chapter: 7
    delivery_method: "flip chart"
    cited_rows: ["F-016"]
  - asset_id: "T01-P3-C08-posibot-ninety-day-paper-test_v001"
    chapter: 8
    delivery_method: "picket sign"
    cited_rows: ["F-023", "F-026"]
`
	if err := os.WriteFile(filepath.Join(promptsDir, "prompt_pack_phase3.yaml"), []byte(promptPackP3), 0644); err != nil {
		return err
	}

	msPhase3 := `# Phase 3: Solidarity (Chapters 7 & 8)

## Chapter 7: The Pivot
*Voice: Constructive Sardonic Authority · For the Reader Who Got the Point*

### Opening: Stop Paying Cover Charges to Work the Night Shift
Let's be crystal clear about one thing: if you taught yourself how CPI interest-rate differentials move currency pairs, how to debug PineScript arrays, and how to log 200 backtest rows in a spreadsheet at midnight, **there is nothing wrong with your work ethic**. Most people won't read a nutritional label, let alone a Federal Open Market Committee minutes release.

Your only mistake was taking a genuinely valuable quantitative brain and volunteering to pay $340 cover charges ([F-001]) to play inside an offshore B-book arcade ([F-011]). When you take that exact same obsession with time-series data, variance, and automation and bill it to companies with real balance sheets, every single gate metric flips from red to green.

### Topic Reference Gate Strip
- **Retail Forex & Prop Challenges ([F-001..F-009]):** [○ Stake: $1,840] [○ Wait: 14.0 mo] [○ Odds: 2.4%] [○ Payoff: $6.40/hr] — **Passes 0 of 4 Gates**

### The 4 Gate-Scored Practical Pivots (outline-spine.v1.yaml §6.2)

#### 1. [PIVOT: THE PRACTICAL PIVOT] — Freelance PineScript / Python Backtesting & Automation Contractor (Tier 1 — Passes 4 of 4)
- **Kind:** *Sell the skill directly (freelance or contract)* ([F-015], [F-017])
- **Gate Strip:** [● Stake: $150] [● Wait: 2.5 mo] [● Odds: 62%] [● Payoff: $55.00/hr] — **Passes 4 of 4 Gates**
- **Failed Gates Answered:** Answers **The Stake** ($150 vs $1,840), **The Wait** (2.5 mo vs 14.0 mo), **The Odds** (62% vs 2.4%), and **The Payoff** ($55.00/hr vs $6.40/hr).
- **How It Works:** Thousands of traders and financial content creators have strategy ideas and zero coding literacy. Instead of risking your rent on a CPI candle, charge **$55.00/hr** ([F-015], [F-017]) to write their PineScript v5 / Python backtrader scripts and deliver walk-forward drawdown reports. Whether their strategy wins or loses, your invoice clears on Friday.

#### 2. [PIVOT: THE PRACTICAL PIVOT] — Regulated Series 65 Paraplanner & Fee-Only Advisory Operations (Tier 1 — Passes 4 of 4)
- **Kind:** *A skill-building route to a better-paid role* ([F-017], [F-018])
- **Gate Strip:** [● Stake: $185] [● Wait: 3.5 mo] [● Odds: 68%] [● Payoff: $38.90/hr ($80,912/yr)] — **Passes 4 of 4 Gates**
- **How It Works:** Unlike the Series 7, the NASAA **Series 65 Uniform Investment Adviser Law Exam** requires **no firm sponsorship**—just a **$185 exam fee** ([F-018]). Fee-only Registered Investment Advisors (RIAs) hire Series 65 paraplanners at a median of **$38.90/hr ($80,912/yr)** ([F-017], [F-018]) to run portfolio rebalancing and risk-tolerance models.

#### 3. [PIVOT: THE PRACTICAL PIVOT] — SQL & Python Financial BI / Risk Data Analyst (Tier 1 — Passes 4 of 4)
- **Kind:** *Get paid to use the skill (employment, SOC 15-2051)* ([F-015], [F-016])
- **Gate Strip:** [● Stake: $320] [● Wait: 4.0 mo] [● Odds: 78%] [● Payoff: $44.50/hr ($92,560/yr)] — **Passes 4 of 4 Gates**
- **How It Works:** An OHLCV candlestick chart is just a time-series database table with four floats and a timestamp. Spend $320 on SQL/Python certification prep and convert your trading journal into an automated variance dashboard ([F-016]). Median pay is **$44.50/hr ($92,560/yr)** with **78% placement odds** ([F-016]).

#### 4. [PIVOT: THE PRACTICAL PIVOT] — Corporate FX Treasury & Commercial Hedging Analyst (Tier 1 — Passes 4 of 4)
- **Kind:** *Get paid to use the skill (employment, SOC 13-2051)* ([F-017])
- **Gate Strip:** [● Stake: $650] [● Wait: 4.5 mo] [● Odds: 71%] [● Payoff: $49.20/hr ($102,336/yr)] — **Passes 4 of 4 Gates**
- **How It Works:** You wanted to trade currencies? Manufacturing exporters, airlines, and cross-border SaaS treasuries pay a median of **$49.20/hr ($102,336/yr)** ([F-017]) to analysts who hedge genuine EUR/USD and USD/JPY invoice exposure using forward contracts—and they pay you a salary plus health insurance instead of charging you reset fees.

### Skill Exchange Table (ch7.skills — Full-Measure Table)

| Skill Trained by Retail FX Study | Who Actually Pays for It | Audited Typical Pay ([F-015..F-018]) |
| :--- | :--- | :--- |
| **Time-Series Data Cleaning (SQL / Python / CSV)** | Corporate FP&A, Fintech & BI Teams (SOC 15-2051) | **$44.50 / hr** ($92,560 / yr) |
| **PineScript / MQL5 / Python Backtest Coding** | Quantitative Funds, Educators & Technical Spec Clients | **$55.00 / hr** (Contract) |
| **Macroeconomic Calendar & Rate-Variance Briefs** | Commercial Treasury & Import/Export Desks (SOC 13-2051) | **$49.20 / hr** ($102,336 / yr) |
| **Risk-Sizing & Portfolio Drawdown Modeling** | Fee-Only RIAs & Wealth Operations (Series 65) | **$38.90 / hr** ($80,912 / yr) |

### The Repurposing Matrix (ch7.repurpose — 6 Items, [F-027])
1. **TradingView / Charting Fluency** -> Repurpose into **PowerBI / Tableau / Plotly executive financial reporting** ($44.50/hr, [F-016]).
2. **PineScript / MQL5 Indicator Scripting** -> Repurpose into **Python pandas / automated QA test scripting** ($55.00/hr, [F-015]).
3. **Economic Calendar (CPI/NFP/FOMC) Tracking** -> Repurpose into **Corporate Treasury FX variance forecasting** ($49.20/hr, [F-017]).
4. **Multi-Monitor Ergonomic Workstation** -> Repurpose directly for **remote BI engineering or paraplanner operations** ([F-020]).
5. **ICT "Fair Value Gap / Smart Money" Lore** -> **Named Plainly as Sunk ($0.00 external employer value)** ([F-015], [F-027]). Never say "Silver Bullet Killzone" in a corporate interview unless you want security to escort you to the elevator.
6. **Expired Prop-Firm Evaluation Reset Fees ($1,428 median)** -> **Named Plainly as Sunk** (unless within the 120-day misrepresentation chargeback window, [F-021], [F-027]).

### Building the Winners' Edges the Slow Way (ch7.edges, [F-018])
- **Mold E2 Edge (*The Built-In Audience*):** Instead of faking funded certificates to hawk 20% reset links ([F-024]), publish transparent open-source Python/SQL market analyses on GitHub and Substack over 12–18 months to attract $55/hr technical clients ([F-017]).

> **[CHAPTER 7 CLOSING BAND]:**  
> There is the ledger in plain black ink: 2.4% odds and $6.40 an hour inside the prop simulator ([F-003], [F-008]), versus 62% to 78% odds and $38.90 to $55.00 an hour selling the exact same analytical habit to companies that pay on Fridays ([F-016], [F-017]). If you are closing MetaTrader and taking one of those four off-ramps, we salute you—go collect your raise. Only readers who haven't got the point and still want to keep trudging forward need what comes next… And… **For the Stubborn Fool Who Must**

---

## Chapter 8: For the Stubborn Fool Who Must
*Voice: Quiet Reckoning · Earned Warmth, Constructive, Zero Personal Advice (outline-spine.v1.yaml §6.3)*

### 1. The Honest Way (4 Evidence-Backed Risk Reducers, [F-023])
Still here? All right. Maybe you love the tape, maybe you have an itch that a corporate BI salary won't scratch, or maybe you looked at 2.4% ([F-008]) and thought, *"So you're telling me there's a 24-in-1,000 chance."* Because **"Probably won't"** still leaves a sliver of **"Could,"** we are not going to lecture you. We are going to bolt roll-cage steel around your bank account so the B-book casino can't bleed you dry:
1. **Ditch Prop Evaluations for Regulated Micro-Lot Cash Accounts (Improves The Stake & The Odds):** Trade CFTC/NFA-regulated micro-lots (0.01 lot) or CME Micro FX Futures (M6E) in a personal account you actually own ([F-001], [F-011]). Zero $340 reset fees, zero 5% trailing drawdown tripwires ([F-007]), and real exchange clearing.
2. **Hard-Cap Effective Leverage at 3:1 to 5:1 (Improves The Odds):** CFTC advisories ([F-012]) prove 50:1 to 500:1 leverage vaporizes median retail accounts in 4.5 months.
3. **Lock Daily Loss Limits at 0.5% of Real Cash (Improves The Stake):** If you lose 0.5% in a morning, your terminal locks until tomorrow ([F-023]).
4. **Clock Every Screen Hour Against $29.76/hr (Improves The Payoff):** Keep a ruthless punch-card of every hour spent staring at candles ([F-003], [F-013]).

### 2. The Smallest Honest Test & Personal Stop Rules ([F-023], [F-026])
- **The Smallest Honest Test (90 Days · $0 Fees):** Before you ever hand another dollar to a prop firm, execute **100 consecutive logged trades** over **90 days (3.0 months)** on a free regulated demo or $250 micro-lot account with **$0 spent on courses, signals, or challenges** ([F-023], [F-026]). If your net expectancy after spread/commission is negative or your max drawdown hits 4.0%, you stop.
- **Personal Stop-Rule Tripwires (ch8.stop_rules, [F-026]):**
  - **Stake Tripwire:** Hard lifetime out-of-pocket ceiling of **$500** (half of the $1,000 Federal Reserve household emergency buffer, [F-014], [F-026]).
  - **Wait Tripwire:** Hard time ceiling of **6.0 months** to reach verifiable net-positive return ([F-026]).
  - **Payoff Tripwire:** Hard exit if trailing 90-day net hourly earnings stay below **$29.76/hr** ([F-013], [F-026]).

### 3. Six Hard Do-Nots (ch8.do_nots — Ordered by Cost, [F-025])
1. **Never buy an evaluation "Reset" (Saves $1,428 median, [F-001]):** See Chapter 3 Waterfall Chart ($340 × 4.2 attempts).
2. **Never join a $149/month VIP Discord Signal Room (Saves $1,341 over 9 months, [K-004], [F-024]):** See Chapter 4 Case File #2 (*Elena Diligence*).
3. **Never buy a $997–$1,497 "Smart Money" Upsell Academy (Saves $1,094 median, [K-004], [K-005]):** See Chapter 4 Case File #2.
4. **Never buy a $499 "Guaranteed Prop-Pass EA Bot" (Saves $499 + account forfeiture, [K-007], [K-008]):** Shared EA signatures trigger automatic B-book breach clauses (Chapter 4 Case File #3, *Devon Precision*).
5. **Never hold open positions across tier-1 CPI/NFP news releases inside a B-book prop account (Saves $5,100 frozen profit, [F-010], [K-015]):** Asymmetric slippage widens spreads by 18+ pips (Chapter 4 Case File #5, *Victor Discipline*).
6. **Never touch household emergency savings (Preserves the $1,000 Fed buffer, [F-014]):** See Chapter 6 Gate Scorecard.

### 4. Six Clean Exit Steps (ch8.exit_steps — Ordered by Money Saved, [F-019..F-021])
1. **[STOP] Kill VIP Signal Rooms & Mentorship Rebills (Saves $149/mo = $1,788/yr, [F-019]):** Revoke card authorization in Stripe/Whop/Discord billing settings today.
2. **[STOP] Cancel Third-Party Indicator Plugins (Saves $68/mo = $816/yr, [F-019]):** Cancel LuxAlgo and custom oscillator subscriptions.
3. **[STOP] Downgrade TradingView Premium & Cancel Forex VPS (Saves $104/mo = $1,248/yr, [F-019]):** Drop charting to the $0 Basic tier and terminate unused London VPS instances.
4. **[PROTECT] File Credit Card Chargebacks for Frozen/Misrepresented Prop Evaluations Within 120 Days (Recovers $340–$540, [F-021], [K-015]):** Visa and Mastercard dispute windows close **120 days** from the transaction date when a platform revokes licenses or refuses contracted payouts; contact your issuing bank's dispute department with saved checkout terms ([F-021]).
5. **[RECOVER] Resell Secondary Monitors, StreamDecks & Pro-Rata Annual Licenses (Recovers $180–$450, [F-020]):** Liquidate unused trading-desk hardware on secondary marketplaces.
6. **[SETTLE] Document Spot FX Losses vs. Evaluation Service Fees for Tax Filing ([F-014], [F-021]):** Under U.S. IRC Section 988 and Section 1256, personal brokerage spot FX/futures losses follow specific tax-election rules, whereas prop evaluation fees are service outlays—**ask a qualified CPA or Enrolled Agent** ([F-014]).

### 5. What It Cost / What You Keep Ledger (ch8.ledger)

| What It Typically Cost ([F-001, K-004..K-012]) | Typical Audited Amount | What You Actually Keep ([F-015..F-020]) | Audited Market Value |
| :--- | :--- | :--- | :--- |
| **4.2x Evaluation & Reset Fees** | -$1,428 | **26.5 Hours/Week of Reclaimed Time** | +1,378 hrs/yr |
| **Charting, VPS & Indicator Tax** | -$412 | **Time-Series SQL/Python/Charting Literacy** | $44.50/hr (SOC 15-2051) |
| **Courses & Signal Rooms (Mold H4)** | -$2,435 | **PineScript / Backtest QA Specification Skill** | $55.00/hr (Contract) |
| **120-Day Chargeback + Hardware Resale** | +$540 to +$720 | **Macro-Calendar & Treasury Risk Fluency** | $49.20/hr (SOC 13-2051) |

### 6. Watch For: The Next 3 Vultures Circling Exiting Traders (ch8.watch_for, [F-022])
1. **"AI Prop-Firm Pass Bots" ($499–$997):** Sends identical trade-copier packets across hundreds of accounts, triggering instant IP/EA ban clauses ([F-022], [K-008]).
2. **Unregulated Offshore Crypto-Prop Challenges:** Identical 5% trailing drawdown math moved onto unregulated perpetual-swap servers ([F-022]).
3. **"Forex Loss Recovery" Upfront-Fee Agencies:** Scammers impersonating regulators who demand a $500 "blockchain recovery tax" ([F-022]).

### 7. Six-Month Aftermath Callbacks (ch8.callbacks, [K-013], [K-015])
- **Victor Discipline (Mold C3 — The Rule Change, [K-013], [K-015]):** Within 45 days of his prop firm's MT5 license freeze, Victor filed a documented Visa services-not-rendered dispute within the **120-day window** ([F-021]), clawed back **$540** in evaluation fees ([K-015]), and used his Python/CSV backtesting scripts to land a **$44.50/hr** financial BI analyst role ([F-016], [K-015]).

### 8. Companion Pamphlet Segue (companion.segue)
If you want the tear-out **90-Day Zero-Fee Paper-Trade Ledger**, the complete **Prop-Firm Red-Flag Audit Checklist**, and the **120-Day Chargeback Dispute Letter Template** on a single pocket card, we publish the standalone **Y-IT Companion Pocket Pamphlet** for **$1.99** (no subscriptions, no upsell ladder, no VIP Discord room).

> **[CHAPTER 8 CLOSING BAND — THE BOOK'S LAST WORD]:**  
> The candles on the 1-minute chart will keep blinking green and red at 2:00 AM whether you are watching them or not. We pulled the CFTC filings, we counted the bodies in the B-book server room, and we taped the receipt to your monitor because we actually give a damn what happens to your $1,840 and your Tuesday nights. Close the terminal, keep the quantitative brain, and go build something that pays you back.
`
	return os.WriteFile(filepath.Join(stageDir, "manuscript_phase3.md"), []byte(msPhase3), 0644)
}
