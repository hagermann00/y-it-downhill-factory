# 00_BRAND_AND_SYSTEM_CANON.md — The Y-IT Master Specification (Temporal Recency Lock)

> **Effective Date:** October 24, 2026  
> **Precedence Rule:** Newer dated specifications (`2026-10-17` > `2026-10-12` > `2026-10-06` > `2026-09-26` > `2026-09-18` > `2026-06-20`) strictly override older drafts where they conflict, while preserving additive visual and structural components.

---

## 1. Brand Identity & Series Spine

- **Imprint Name:** **Y-IT** (pronounced *"Why-It"*) — Parent Studio: **It Is What It Is (`IIWII`)**
- **Canonical Page-1 Series Tagline:**
  > **Why It** ~~Could.~~ ~~Would.~~ ~~Has.~~ ~~Might.~~ **Probably Won't.** **WORK**  
  > *Yeah… about that.*
- **Core Premise:** Every Y-IT book is a field guide for someone standing at the edge of an over-hyped modern side-hustle or financial dream with their credit card in their hand. It gives the reader the real math, the real cost structure, the real attrition curve, and a dignified exit before they pay tuition to a guru or platform.
- **Target Format:**
  - **Print Trim:** `6.00" × 9.00"` (`6.125" × 9.25"` with `0.125"` bleed on top, bottom, and outside edges).
  - **Target Page Count:** **50 to 80 pages** (sweet spot: **64 pages**, 8 chapters + front matter + appendices + index).
  - **Spine Rule:** KDP paperbacks under 80 pages have **no spine text** (spine width at 64 pages white paper is `64 × 0.002252" = 0.1441"`).
  - **Recto Opening Rule:** Every chapter (`Ch 1` through `Ch 8`) and major section opens on a **recto (odd, right-hand) page** (`break-before: right`).

---

## 2. The 3-Phase Emotional Arc & 8-Chapter Spine (`outline-spine.v1.yaml`)

The book is strictly partitioned into **three phases** with distinct emotional registers, visual palettes, and evidence rules:

### Phase 1 · Vibrant Hope (Pages ~7–18 · Also Published Standalone as the Free 2-Chapter Giveaway)
- **Voice Register:** **Corporate Sincerity** — Warm, breathless, pitch-perfect recreation of the dream as sold by top gurus, landing pages, and lifestyle influencers. No winking, no premature cynicism, no audited failure rates.
- **Evidence Rule:** **Claim rows only** (`claim_row`). Pitch numbers (*"$10,000/month"*, *"$200,000 funded account in 14 days"*, *"90% profit split"*) are allowed. Audited failure statistics are strictly banned in Phase 1.
- **Chapter 1: *The Promise***
  - **Visual Treatment:** **Full-Color Magazine-Ad Collage (`#FFFDF7` warm glossy stock, gold/emerald prosperity accents `#0D9488` / `#F59E0B`)**.
  - **Structure:** Opens with an unnamed everyman second-person vignette (`"It is 11:42 PM on a Tuesday..."`), displays the freedom pitch, the income claim, the social proof carousel, the lifestyle upgrade promise, and the *"Why Not You?"* hook.
- **Chapter 2: *The Sucker's Roadmap***
  - **Visual Treatment:** **IKEA-Style Flat-Pack Assembly Manual (`#F4F7FB` blueprint tint, numbered assembly steps, isometric tool diagrams)**.
  - **Structure:** Presents the 5–7 "simple" steps every beginner is told to follow, complete with "Parts Included," "Estimated Assembly Time," and "No Experience Required" badges.
  - **Mandatory Cliffhanger Ending:** Ends abruptly at the bottom of the final page with a dashed empty box:
    > **`[ STEP N+1 WAS NOT PROVIDED IN YOUR STARTER KIT. TURN THE PAGE FOR THE AUTOPSY. ]`**

### Phase 2 · Demolition (Pages ~19–44 · The Cold Ledger)
- **Voice Register:** **Sardonic Authority** — Forensic, dry, unsparing, mathematically exact. Every number carries an audited source citation (`[R-###]`).
- **Chapter 3: *The Money Pit***
  - **Structure:** Dismantles the advertised entry cost (`$49–$199`) against the real 6-month TCO (`$1,800–$6,400+`): reset fees, data feeds, TradingView Pro, EA subscriptions, course upsells, Discord VIP rooms, and tax/spread drag. Includes the **Iceberg Cost Diagram** and **Advertising vs. Reality Ledger Table**.
- **Chapter 4: *Tales from the Trenches* (Standalone Forensic Autopsy)**
  - **Visual Treatment:** **Aged Amber Case-File Paper (`#EFE5CE`)**, black redacted dossier headers, rubber-stamp classifications, silhouette avatars (`silhouette_01`..`05`), and monospace financial exhibits.
  - **Strict Boundary Rules (ADR 0047):**
    - **Zero PosiBot appearances** anywhere in Chapter 4.
    - **Zero jokes or punchlines.** (*"The numbers are the joke."*)
    - **5 Composite Profiles** drawn from **5 distinct molds** of the 13-Mold Menu (`_canon/contracts/mold-menu.yaml`), where each composite is backed by $\ge 3$ documented real-world accounts in `01_research_backplane/rows.yaml`.
- **Chapter 5: *Exit Liquidity***
  - **Structure:** Maps who the beginner is actually competing against (institutional liquidity providers, HFT market makers, B-book synthetic execution engines, Drawdown-trap rule designers, and affiliate educators who make 100% of their income from challenge referral codes). Includes the **Ecosystem Food-Chain Diagram** and **Attrition Funnel Chart**.
- **Chapter 6: *The Verdict***
  - **Structure:** Runs the venture through the **4-Metric Gate Scorecard** (`outline-spine.v1.yaml`):
    1. **The Stake:** Is true 6-month capital at risk $\le$ median US emergency savings?
    2. **The Wait:** Is median time to sustainable net profit $\le 6$ months?
    3. **The Odds:** Are $\ge 50\%$ of participants net-profitable at 12 months?
    4. **The Payoff:** Does median net hourly return exceed the US median hourly wage (`~$29.76/hr`)?
  - **Output:** Renders the **4-Metric Radar/Scorecard**, the **House Edge Comparison Table** (comparing the venture against Blackjack, Roulette, S&P 500 indexing, and skilled trades), and the **Final Verdict Stamp** (`FAIL — 0 OF 4 GATES CLEARED`).

### Phase 3 · Solidarity (Pages ~45–64 · The Salvage Operation)
- **Voice Register:** **Warm Solidarity** — You are not stupid for wanting out of a rigged schedule; the machine was engineered by teams of behavioral economists to extract your deposit. Now let's salvage the underlying drive and skill.
- **Chapter 7: *The Pivot***
  - **Structure:**
    - **3 to 5 Gate-Scored Pivots** (`[PIVOT: THE PRACTICAL PIVOT]`) that use the exact same analytical/quantitative appetite (` Quantitative Data Automation`, `B2B FX Treasury / Hedging Ops`, `Commission-Backed Technical Sales`, `Regulated Series 3 / 65 / CFA Advisory Track`) — each scored against the same 4-Metric Gate.
    - **Repurposing Ledger:** What to keep from what you already bought/learned.
    - **Skill Translation Matrix:** How to rewrite "failed prop trader" into high-value operational resume bullets.
    - **3D Papercraft Comic Strip (`[PAGE COMIC — 3D PAPERCRAFT]`):** Visual relief sequence showing PosiBot trading his broken prop-challenge terminal for a real cash-flowing toolset.
- **Chapter 8: *For the Stubborn Fool Who Must***
  - **Structure:** For the 5% who read Chapters 1–7 and still refuse to walk away:
    - **The Honest Way:** Stripping away leverage illusions and challenge resets.
    - **The Smallest Honest Test:** 90-day paper/micro-lot statistical proof requirement before paying a single dollar in evaluation fees.
    - **Hard Stop Rules:** Pre-signed capital and time circuit-breakers.
    - **5 to 7 Hard Do-Nots:** (e.g., never buy a reset within 14 days of a blown account; never buy an algorithmic EA sold via Telegram).
    - **4 to 9 Clean Exit Steps:** How to cancel recurring subscriptions, request clawbacks/refunds where applicable, and close the ledger cleanly.
    - **Cost/Keep Ledger & Watch-For Signals.**
    - **Companion Pamphlet Segue & Callbacks** to Chapters 1–6.
- **Back Matter & Appendices:**
  - **Appendix A: Audited Sources & Evidence Tiers** (`Tier 1` Regulatory/Academic, `Tier 2` Platform Disclosures/Audited Filings, `Tier 3` Practitioner/Community Aggregations).
  - **Appendix B: The Unvarnished Glossary** (20 industry terms translated from guru-speak into plain English).
  - **Results Typical Disclosure & IIWII Author Note** (featuring Brian's illustrated IIWII avatar from `_canon/assets/author_iiwii/`).
  - **Last-Page Issue Index & Field Notes Grid.**

---

## 3. Visual Architecture & `ME ODDEN CAR TOON_files` Integration

All four visual pillars from `C:\Users\dell3630\Documents\Obsidian-Brain\00_Inbox\ME ODDEN CAR TOON_files` and `y-it-auto-designer` are locked into `_canon/assets/` and rendered automatically by `internal/catalog` and `internal/render`:

1. **Cover Concept A — The Blueprint Grid Cover (`covers_blueprint/blueprint_cover_populated_ref.jpg`):**
   - Engineering grid background with callout annotations (`[TITLE]`, `[SUB-HEADER]`, `[IMAGE]`).
   - Top Hook: **`why your [TOPIC] probably will Fail`**
   - Sub-Hook: **`EVEN IF YOUR AI CHAT BOT LIKES YOUR CHANCES!`**
   - Bottom Banner: **`We reveal the secret method in the first 2 chapters...`** + **`First 2 chapters free at itiswhatitis.biz`**.
2. **Cover Concept B — The Sinking-Ship PosiBot Cover (`covers_blueprint/sinking_ship_posibot_cover_v1.jpg`):**
   - High-contrast oil-illustration of a burning paper/wooden galleon sinking in a storm while PosiBot stands on the bow holding a cheerful sign (`"PASSIVE INCOME AHEAD!"`).
3. **PosiBot Sign System (`posibot/posibot_blank_sign_pose_sheet.jpg` + cutouts):**
   - Knee-height kraft-tan robot (`#C8A27A`), dark rounded CRT screen face (`#111827`), glowing cyan eyes and wide fixed grin (`#00F0FF`) that **never changes expression**, even while standing in financial wreckage.
   - Holds wooden/cardboard signs with **$\le 6$ words shouted in bold caps** + a tiny parenthetical reality undercut underneath.
   - Appears in **Chapters 1, 2, 3, 5, 6, 7, 8** (strictly **excluded from Chapter 4**).
4. **3D Papercraft Comic Strip (`comics_papercraft/comic_3d_papercraft_strip_ref.jpg`):**
   - Multi-panel tactile papercraft diorama comic (`[PAGE COMIC — 3D PAPERCRAFT]`) placed at major phase transitions to provide visual breathing room between dense financial tables.
5. **IIWII Author Bio & Stamp Assets (`author_iiwii/brian_iiwii_architect_avatar.jpg` & `brian_iiwii_badge_avatar.jpg`):**
   - Used in the Back Matter Author Colophon (*"Built at the IIWII Desk"*).

---

## 4. Universal Phase 1 Research Architecture (`yit_phase1_universal_research_prompt.v1.txt`)

Stage `01_research_backplane` executes both the **4-Layer Universal Phase 1 Research Specification** (Oct 17, 2026) and the **38-Question Stockpile** (`y-it-engine` Spec v1):

- **Layer A — Promotional Presentation (`promotional_presentation`):** Headline claims, narrative hooks, lifestyle imagery, authority signals, and pricing ladders as presented to prospects.
- **Layer B — Operational Pathway (`operational_pathway`):** The exact sequence of steps (`Step 1`..`Step N`), tools, accounts, capital outlays, and skill assumptions required in practice.
- **Layer C — Real-World Friction (`friction_register.yaml`):** Structured register of every divergence between Layer A and Layer B (`entry_id`, `stage`, `promised_state`, `documented_reality`, `cost_delta`, `attrition_rate`, `tier`, `row_id`).
- **Layer D — Presentational Encoding (`visual_persona_atlas`):** Archetypes, vocabulary tokens, color/status codes, and psychological triggers used by the industry's marketing funnel.
- **Programmatic Gates (`G-01`..`G-14` mapped into Go Gates `G1`..`G4`):** Verified by `yit-factory.exe verify`.
