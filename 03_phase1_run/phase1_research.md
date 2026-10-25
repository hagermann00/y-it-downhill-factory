# Y-IT — PHASE 1 RESEARCH & EDITORIAL ENGINE (v2.0 Output)
## Topic: Retail Forex & Proprietary Trading ("Funded Trader") Evaluation Challenges

---

## 1. SCOPE & VARIANT ARCHITECTURE
- **Activity Scoped:** Retail proprietary trading ("prop firm") evaluation challenges and funded trader programs across two primary operational variants:
  - **Variant A (`CFD / Spot FX Multi-Step Evaluation` — e.g., `SRC-001` FTMO, `SRC-002` FundedNext, `SRC-003` The5ers):** Two-phase (or instant/1-step) evaluation on simulated CFD/FX accounts (`MT4`, `MT5`, `cTrader`, `TradingView`) with advertised scaling up to `$2,000,000`–`$4,000,000` (`SRC-001`, `SRC-003`) and `85%`–`90%` reward shares (`SRC-002`).
  - **Variant B (`Regulated Exchange Futures Combine` — e.g., `SRC-004` Topstep):** Simulated futures evaluation (`Trading Combine` -> `Express Funded Account`) advertising up to `90%` profit keep (`SRC-004`).
- **Transformation Promised:** Transition from undercapitalized retail trader to funded trader managing six- to seven-figure notional capital without risking personal savings beyond the evaluation fee (`SRC-001..SRC-004`).
- **Authentic vs. Synthetic Boundary (`v2.0 Rule 3`):**
  - Only live HTTP captures in `03_phase1_run/authentic_captures/` (`SRC-001..SRC-004`, `SRC-006`) and live API probes in `01_research_backplane/live_harvest_sources.csv` (`LIVE-SRC-001..206`) are classified `AUTHENTIC`.
  - Earlier local HTML fixtures (`src_01..src_09` in `01_research_backplane/pages/`) are explicitly classified `SYNTHETIC` (`SYNTH-EXCL-001`) and excluded from real-source claims.

---

## 2. SOURCE-BACKED CLAIM BANK (`CLM-001..CLM-008`)

| Claim ID | Wording / Faithful Paraphrase | Type | Speaker / Source ID | Conditions Stated | Evidence Status |
|---|---|---|---|---|---|
| `CLM-001` | `"Grow & Monetize Your Demo Trading"` / `"Prove your skills. Earn real-money rewards."` | Explicit | FTMO (`SRC-001`) | Must pass evaluation objectives | `SELLER-CLAIM ONLY` (Verified that seller states it) |
| `CLM-002` | `"Scaling plan: Grow up to $2M"` | Explicit | FTMO (`SRC-001`) | Requires recurring multi-month consistency targets | `SELLER-CLAIM ONLY` |
| `CLM-003` | `"We sponsor your trading journey"` / `"85% Reward Share"` / `"2 Minimum Profitable Days in each Phase"` | Explicit | FundedNext (`SRC-002`) | Subject to `12% Maximum Loss Limit` and phase rules | `SELLER-CLAIM ONLY` |
| `CLM-004` | `"99.99% processed within 24 hours, no delays — or get $1,000 extra"` | Explicit | FundedNext (`SRC-002`) | Applies to approved reward disbursements | `SELLER-CLAIM ONLY` |
| `CLM-005` | `"Get qualified and manage up to $4,000,000!"` / `"Turning your passion into a trading career"` | Explicit | The5ers (`SRC-003`) | Qualification and progressive scaling milestones required | `SELLER-CLAIM ONLY` |
| `CLM-006` | `"Payouts of $98,038"`, `"Payouts of $35,000"`, `"Payouts of $21,769"`, `"Payouts of $20,206"` | Explicit | The5ers (`SRC-003`) | Individual highlighted trader payouts on home page ticker | `ANECDOTAL / SELLER-CURATED` (Denominator unknown) |
| `CLM-007` | `"Learn and trade futures with our prop firm capital and keep up to 90% of your profits."` | Explicit | Topstep (`SRC-004`) | Subject to Trading Combine & Express Funded rules | `SELLER-CLAIM ONLY` |
| `CLM-008` | Global OTC foreign exchange turnover averaged `$7.5 trillion per day` in April 2022 (`6%` non-financial customers). | Explicit | BIS (`SRC-006`) | Wholesale & institutional OTC FX survey | `VERIFIED` (Tier-1 Central Bank data; frequently cited by marketers to imply retail opportunity) |

---

## 3. VISUAL ATLAS (`VIS-001..VIS-005`)

- **`VIS-001` (`Simulated Account Scaling & Reward Hero Banners` — `SRC-001..SRC-004`, `4/4` inspected seller home pages):** High-contrast dark-mode or clean corporate hero blocks pairing six/seven-figure account ceilings (`"$2M"`, `"$4,000,000"`, `"$1.5B+ paid"`) with one-click `"Start Challenge"` / `"Free Trial"` CTAs.
- **`VIS-002` (`Curated Payout Ticker Cards` — `SRC-003` The5ers):** Scrolling carousel of individual trader portraits paired with exact dollar payout badges (`"Payouts of $98,038"`, `"Payouts of $35,000"`, `"Payouts of $20,116"`). Observation: no failed-attempt counter or cohort pass-rate denominator appears adjacent to the cards.
- **`VIS-003` (`Award & Trustpilot Laurel Grids` — `SRC-001..SRC-003`, `3/4` pages):** Repeated badge clusters (`"MOST USER FRIENDLY PROP FIRM"`, `"BEST SCALING STRUCTURE"`, `"HIGHEST RATED PROP FIRM"`, `"95% customer satisfaction score"`).
- **`VIS-004` (`Platform Integration Logos` — `SRC-001`, `SRC-002`):** Prominent `TradingView`, `MetaTrader 5`, and `FX Blue App Suite` badges signaling institutional-grade tooling.
- **`VIS-005` (`Fine-Print Simulated Program Footer` — `SRC-004` Topstep):** Small-print legal footer at bottom of page disclosing: *"Simulated trading program services, comprising the Trading Combine, and the Express Funded Account... are operated by TopstepTrader, LLC."*

---

## 4. COMPLETE STEP TABLE — ADVERTISED vs. REAL-WORLD COUNTERPART (`STEP-001..STEP-007`)

*(All costs, prerequisites, restrictions, and friction links live here in the backplane table so the public Chapter 2 Micro-Manuscript stays feather-light.)*

| Step ID | Promoted Action (Idealized) | Real-World Operational Counterpart | Requirement Class | Variants & Branches | Dependencies & Skills | Friction Links |
|---|---|---|---|---|---|---|
| `STEP-001` | **Choose Your Challenge Size** (Pick `$10K`–`$200K` tier or Free Trial) | Select CFD (`SRC-001..003`) or Futures (`SRC-004`) evaluation plan; pay upfront evaluation fee (or start limited demo trial) | `MANDATORY` | Variant A: 2-Step CFD (`SRC-001`, `SRC-002`); Variant B: 1-Step/Instant (`SRC-001`, `SRC-003`); Variant C: Futures Combine (`SRC-004`) | Valid payment method; country eligibility check (U.S. restrictions apply on many offshore CFD firms) | `FRIC-001` |
| `STEP-002` | **Connect Your Trading Platform** (`TradingView` / `MT5` login) | Link evaluation credentials to charting & execution terminal (`TradingView`, `MT5`, `cTrader`) | `MANDATORY` | Platform choice varies by firm (`SRC-001` supports TradingView/MT5; `SRC-004` uses futures platforms) | Platform configuration, order-entry literacy | `FRIC-002` |
| `STEP-003` | **Follow the Trading Objectives** (Stay within daily and max loss limits) | Execute trades while strictly respecting maximum daily loss (commonly `4%–5%`) and maximum overall loss (`8%–12%`, `SRC-002`) | `MANDATORY` | Instant breach termination if equity/balance crosses daily or max loss threshold | Position sizing math, stop-loss discipline, volatility awareness | `FRIC-003` |
| `STEP-004` | **Hit Phase 1 Profit Target** (`Prove your skills` — minimum 2 trading days) | Achieve Phase 1 percentage gain target while logging at least `2 minimum profitable days` (`SRC-002`) without breaching loss limits | `MANDATORY` (for 2-step) | Skipped in Instant Funding branch (which trades lower initial split/tighter drawdown for no Phase 1) | Positive expectancy under spread/commission drag | `FRIC-003`, `FRIC-004` |
| `STEP-005` | **Complete Phase 2 Verification** (Confirm consistency) | Repeat profitable execution in Phase 2 verification stage under identical drawdown limits | `MANDATORY` (for 2-step) | Single-step combines (`SRC-004`) merge evaluation into one stage with consistency rules | Sustained risk control across consecutive weeks | `FRIC-003`, `FRIC-004` |
| `STEP-006` | **Sign Contract & Activate Funded Demo Account** (`Become a Funded Trader`) | Complete KYC/identity verification and sign contractor agreement for a **simulated funded account** (`"Monetize Your Demo Trading"`, `SRC-001`, `SRC-004`) | `MANDATORY` | KYC rejection or prohibited trading-style review (news straddling, tick scalping, copy trading) can block activation | Clean KYC documentation; compliance with prohibited-strategy clauses | `FRIC-004`, `FRIC-005` |
| `STEP-007` | **Request Reward Payout & Scale Up** (`Keep up to 85%–90%` / `Grow up to $2M–$4M`) | Maintain net profit above high-water mark, request reward split (`85%–90%`, `SRC-002`, `SRC-004`), and hit multi-month profit milestones to qualify for capital scaling (`$2M` `SRC-001`; `$4M` `SRC-003`) | `CONDITIONAL` (requires net profit without prior breach) | Payout processed in `24h` (`SRC-002` claim); scaling requires consecutive profitable cycles | Multi-month survival above drawdown floor | `FRIC-003`, `FRIC-005` |

---

# PUBLIC PREVIEW 1: CHAPTER 1 — COLLAGE MANIFEST (`v2.0` Format)

- **Scene-Setting Context (`26 words total` — well under the `45-word` maximum):**
  *Every day, $7.5 trillion moves through global currency markets (`SRC-006`). Modern prop firms have opened the vault: pass a digital audition, trade their capital, keep 90%.*

| Artifact ID | Visible Microcopy | Source ID & Location | Verbatim or Paraphrase | Authentic Visual Reference / Proposed Treatment | Size / Placement Priority |
|---|---|---|---|---|---|
| `ART-001` | `"Grow & Monetize Your Demo Trading"` | `SRC-001` (`ftmo.com/en/` Hero H2) | Verbatim (`6 words`) | Dark-navy institutional hero banner over upward candlestick silhouette | **P1 — Lead Hero Headline** (Top of Page 1) |
| `ART-002` | `"We sponsor your trading journey"` | `SRC-002` (`fundednext.com/` Hero H2) | Verbatim (`5 words`) | Clean corporate sub-banner with glowing cyan accent border | **P1 — Primary Sub-Banner** |
| `ART-003` | `"Keep up to 90% of your profits."` | `SRC-004` (`topstep.com/` Hero Lead) | Verbatim (`7 words`) | Tilted starburst / signpost sticker held by PosiBot (`posibot_pose1_tilted_sign.jpg`) | **P1 — Center Callout Sticker** |
| `ART-004` | `"Get qualified and manage up to $4,000,000!"` | `SRC-003` (`the5ers.com/` Hero Lead) | Verbatim (`7 words`) | Gold-foil account tier display box (`$10K -> $100K -> $2M -> $4M`) | **P1 — Capital Ceiling Display** |
| `ART-005` | `"Prove your skills. Earn real-money rewards."` | `SRC-001` (`ftmo.com/en/` How It Works H2s) | Verbatim (`6 words`) | Two-step ribbon badge across facing spread | **P2 — Spread Bridge Banner** |
| `ART-006` | `"$7.5 Trillion Daily Global FX Turnover"` | `SRC-006` (`bis.org` Triennial Survey Highlights) | Faithful Paraphrase (`6 words`) | PosiBot presenting silver whiteboard (`posibot_pose2_whiteboard.jpg`) | **P2 — Market Scale Callout** |
| `ART-007` | `"Get rewarded in 3 days · 99.99% processed within 24 hours"` | `SRC-002` (`fundednext.com/` Plan Block) | Verbatim (`10 words`) | Stopwatch speed-payout seal (`"OR GET $1,000 EXTRA"`) | **P2 — Speed Guarantee Seal** |
| `ART-008` | `"2 Minimum Profitable Days in each Phase"` | `SRC-002` (`fundednext.com/` Plan Card) | Verbatim (`7 words`) | Checklist coupon stamp with green checkmarks | **P2 — Low-Barrier Badge** |
| `ART-009` | `"Payouts of $98,038 · Payouts of $35,000 · Payouts of $21,769"` | `SRC-003` (`the5ers.com/` Real Traders Carousel) | Verbatim (`10 words`) | 3-card horizontal trader payout strip with 5-star rating header | **P2 — Social Proof Strip** |
| `ART-010` | `"Turning your passion into a trading career"` | `SRC-003` (`the5ers.com/` Hero H1) | Verbatim (`7 words`) | Beach hammock laptop lifestyle photo (`ad_ch1_hero_beach.jpg`) + Before/After split (`ad_ch1_before_after.jpg`) | **P1 — Dominant Lifestyle Visuals** |
| `ART-011` | `"MOST USER FRIENDLY PROP FIRM · BEST SCALING STRUCTURE"` | `SRC-003` (`the5ers.com/` Award Grid H3s) | Verbatim (`9 words`) | Gold laurel award crest row | **P3 — Trust Badge Footer** |
| `ART-012` | `"Your challenge is open now."` | `SRC-002` (`fundednext.com/` Closing CTA) | Verbatim (`5 words`) | Bottom-right tear-out action card with PosiBot megaphone (`posibot_pose3_megaphone.jpg`) | **P1 — Closing Page-Turn CTA** |

---

# PUBLIC PREVIEW 2: CHAPTER 2 — THE OPERATIONAL PATHWAY (`v2.0` Micro-Manuscript)

> **Banner Label:** `THE SELLER'S IDEALIZED PATHWAY — MODEL FOREX-PROP-5000 (7 EASY STEPS)`

### STEP 1: PICK YOUR ACCOUNT SIZE
- **Action (`16 words`):** Select your desired evaluation balance—from $10,000 up to $200,000—and activate your challenge dashboard.
- **Tiny Tip (`6 words`):** Start with the free practice trial.
- **Diagram Cue:** `[FIG-01: Isometric tier selector slider clicked onto "$100,000 Challenge" with a cheerful checkmark.]`

### STEP 2: CONNECT YOUR CHARTING PLATFORM
- **Action (`15 words`):** Log into TradingView or MetaTrader 5 using the credentials emailed to your inbox.
- **Tiny Tip (`6 words`):** All major platforms work out-of-the-box.
- **Diagram Cue:** `[FIG-02: Laptop screen snapping neatly onto a glowing server plug labeled "MT5 / TradingView".]`

### STEP 3: STAY INSIDE THE LOSS GUARDRAILS
- **Action (`17 words`):** Keep your daily dip above the daily loss line and your overall balance above maximum drawdown.
- **Tiny Tip (`7 words`):** Built-in dashboard meters track everything automatically.
- **Diagram Cue:** `[FIG-03: Smooth rising line gliding comfortably above two padded bumper rails.]`

### STEP 4: LOG TWO PROFITABLE DAYS (PHASE 1)
- **Action (`16 words`):** Hit the Phase 1 profit target across at least two green trading days to advance.
- **Tiny Tip (`5 words`):** No maximum time limit pressure.
- **Diagram Cue:** `[FIG-04: Two calendar squares marked with bright stars and a "PHASE 1 COMPLETE" stamp.]`

### STEP 5: CONFIRM IN PHASE 2
- **Action (`15 words`):** Repeat your routine to clear the smaller Phase 2 verification target under the same rules.
- **Tiny Tip (`6 words`):** Instant-funding plans skip right past this.
- **Diagram Cue:** `[FIG-05: PosiBot with clipboard (posibot_pose4_clipboard.jpg) checking off box #2.]`

### STEP 6: ACTIVATE YOUR FUNDED DEMO ACCOUNT
- **Action (`16 words`):** Upload your ID, sign the digital trader agreement, and unlock your monetized funded account.
- **Tiny Tip (`6 words`):** Approval takes just a few hours.
- **Diagram Cue:** `[FIG-06: Gold-embossed "OFFICIAL FUNDED TRADER" certificate sliding out of an envelope.]`

### STEP 7: COLLECT 90% REWARDS & SCALE TO $4M
- **Action (`18 words`):** Click withdraw to receive up to 90% of your simulated profits—processed within 24 hours—and scale to $4,000,000.
- **Tiny Tip (`7 words`):** Initial challenge fee refunded with first reward.
- **Diagram Cue:** `[FIG-07: Ascending staircase ($100K -> $500K -> $2M -> $4M) leading to a direct-deposit bank icon.]`

---

## 5. BACKPLANE FRICTION SUMMARY & PHASE 2 HANDOFF QUESTIONS (`Held Back from Public Copy`)
- **`FRIC-001` (`Simulated Demo Environment vs. Live Market Execution`):** Even after passing (`STEP-006`), traders remain in a *"monetized demo"* / *"simulated trading program"* (`SRC-001` `"Monetize Your Demo Trading"`, `SRC-004` footer disclosure).
- **`FRIC-002` (`Asymmetric Drawdown Math`):** A `10%` maximum loss limit (`$10,000` on a `$100,000` nominal account) means the trader's true risk capital is `$10,000`, not `$100,000` (`STEP-003`).
- **`FRIC-003` (`Recurring Reset Economics`):** Breaching the daily or maximum loss rail (`STEP-003..005`) terminates the challenge and requires purchasing a new challenge or reset fee.
- **`FRIC-004` (`Unstated Cohort Denominator`):** Seller pages showcase top dollar payouts (`$98,038`, `SRC-003`) and aggregate payouts (`$316.1M+`, `SRC-002`), but omit the denominator of total challenges purchased and the percentage of starters who reach `STEP-007`.
- **Handoff Questions for Phase 2 (`Chapters 3–6`):**
  1. What percentage of paid evaluation starters pass both phases (`STEP-004` + `STEP-005`) and receive at least one payout (`STEP-007`)?
  2. How much revenue do prop firms derive from evaluation/reset fees (`STEP-001`) versus actual market profits?
  3. How do intraday trailing drawdown rules and news-window restrictions affect survival probability in `STEP-006` and `STEP-007`?
