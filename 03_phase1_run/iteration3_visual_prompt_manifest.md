# Y-IT PHASE 1 — COORDINATED VISUAL PROMPT & INTERSTITIAL MANIFEST (`Iteration 3`)

- **Architect Model:** `Gemini 4 Argon (High)` (standing in for Opus 5.5 per operator directive)
- **Canonical Spec Reference:** `D:\y-it-auto-designer\docs\imagery-prompts.md` + `docs\components.md` + `app\src\flow\composed.tsx`
- **Core Principle (Visual Division of Labor):** Imagery stands in place of explanatory text. The prose sets up the premise and hands the beat directly to the visual slot (`[[SLOT:...]]` or `[[INTERSTITIAL:...]]`), never repeating in prose what the visual already shows.

---

## 1. PRE-FORMATTED COMPONENT PROMPT TEMPLATES (SHARED BASELINES)

### 1.1 Shared Style Block (`STYLE_BLOCK`)
> Editorial illustration meets satirical business book, Bloomberg Businessweek cover logic crossed with The Onion's art direction. Flat editorial linework, confident shapes, restrained papercraft/matte texture. No photorealism, no stock-photo look, no AI-generic gloss.

### 1.2 Closed Phase 1 Palette (`PALETTE_PHASE_1`)
> Warm Cream `#F5F2EC` (paper ground with `#5D7BEA12` halftone dot matrix), Warm Neutral `#E5E0D8`, Charcoal `#3A3A3A`, Navy `#1E2B38`, Corporate Blue `#5D7BEA`, Warning Red `#E35050`, Highlight Yellow `#F8C952`, Sage `#7CB8B8`, Rust `#B3502B`.

### 1.3 PosiBot Character Reference Block (`POSIBOT_REF_BLOCK`)
> A knee-height robot with a matte kraft-tan box body `#C4A882`, a heavy dark-brown outline `#2A2420`, chunky black gloves, wheeled or booted feet and joints, a near-black rectangular screen head `#22282B` with a glowing cyan `#4DD0E1` face (two curved happy eyes, simple arc smile) and a ball-tip antenna on the side of the head. The face never changes. No wood grain. Plain torso with small cyan chest indicator or clean matte panel.

### 1.4 Global Negative Prompt (`GLOBAL_NEGATIVE`)
> Avoid: text, lettering, captions (except inside PosiBot's approved sign surface), watermark, logo, signature, photorealistic, 3D render, plastic, glossy, neon, lens flare, bokeh, anime, chibi, detailed human faces, real people, brand marks, extra limbs, deformed robot, wood grain on robot, cluttered background, oversaturated.

---

## 2. NANO BANANA IMAGE SLOTS (`[[SLOT:...]]`)

| Slot Marker | Component Type | Print Size / 300 DPI Target | Delivery Method & Pose | Customized Nano Banana Prompt | Asset File |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `[[SLOT:AD_CH1_HERO_BEACH]]` | `Ch 1 Bleed Photo Ad` | `3.22" × 2.64"` (`967 × 792 px`) | N/A (Silhouette Scene) | Editorial illustration, flat linework: a dignified dark-charcoal `#3A3A3A` silhouette lounging in a woven hammock between two palm trees overlooking a calm sage `#7CB8B8` ocean under a golden sun `#F8C952`; open laptop on a blue-and-cream striped beach towel in foreground showing an upward `#5D7BEA` chart. Wordless. (`GLOBAL_NEGATIVE`) | `assets/ad_ch1_hero_beach.png` |
| `[[SLOT:POSIBOT_CH1_01]]` | `PosiBot (Opener)` | `2.24" × 2.80"` (`900 × 1125 px`) | `whiteboard` (Front-facing chest hold) | `POSIBOT_REF_BLOCK` standing upright facing forward on flat `#F5F2EC` cream, holding a silver-framed rectangular whiteboard across his chest with both black-gloved hands reading in bold uppercase charcoal lettering: `"THE MARKET IS $7.5 TRILLION!"` | `assets/posibot_pose2_whiteboard.png` |
| `[[SLOT:POSIBOT_CH1_02]]` | `PosiBot (Contour Wrap)` | `3.00" × 3.75"` (`900 × 1125 px`) | `signpost` (Tilted wooden placard, right-angled) | `POSIBOT_REF_BLOCK` standing on the left third of a `#F5F2EC` cream frame, holding a tilted wooden-framed rectangular signpost angled toward the upper-right so body text wraps underneath the sign; bold charcoal lettering on sign reads: `"90% PROFIT SPLIT!"` | `assets/posibot_pose1_tilted_sign.png` |
| `[[SLOT:AD_CH1_BEFORE_AFTER]]` | `Ch 1 Before/After Split` | `1.39" × 3.30"` (`556 × 990 px`) | N/A (Vertical 2-Panel) | Vertical 2-panel split illustration: Top ("Before"): drab grey cubicle under flat fluorescent light with slumped `#3A3A3A` silhouette. Bottom ("After"): sunlit coastal balcony at dawn with glowing `#5D7BEA` laptop and espresso cup. Wordless. (`GLOBAL_NEGATIVE`) | `assets/ad_ch1_before_after.png` |
| `[[SLOT:POSIBOT_CH1_03]]` | `PosiBot (Mascot Card)` | `2.24" × 2.80"` (`900 × 1125 px`) | `megaphone` (Side-stride march + shout bubble) | `POSIBOT_REF_BLOCK` in an energetic side-stride pose, holding a corporate-blue `#5D7BEA` megaphone with a crisp white shout banner reading in bold uppercase `#E35050` and `#3A3A3A` lettering: `"ZERO PERSONAL RISK!"` on flat `#F5F2EC` cream. | `assets/posibot_pose3_megaphone.png` |
| `[[SLOT:POSIBOT_CH2_01]]` | `PosiBot (Manual Inspector)` | `2.24" × 2.80"` (`900 × 1125 px`) | `clipboard` (Thumbs-up + angled clipboard) | `POSIBOT_REF_BLOCK` giving an enthusiastic thumbs-up with his right black-gloved hand while holding an oversized tan clipboard in his left hand reading in bold uppercase lettering: `"FUNDED IN 14 DAYS!"` on pure white `#FFFFFF`. | `assets/posibot_pose4_clipboard.png` |
| `[[SLOT:FIG_CH2_01..06]]` | `Ch 2 Flat-Pack Manual Figures` | `2.36" × 2.08"` (`900 × 900 px`) | Isometric line drawings (`fig. 1`..`fig. 6`) | Clean IKEA-style flat-pack instruction manual isometric line drawings, `#3A3A3A` black line on `#FFFFFF` white with `#5D7BEA` corporate-blue accent highlights and dotted assembly arrows (`fig. 1` open parts box with empty slot D; `fig. 2` challenge portal; `fig. 3` pastel ICT order blocks; `fig. 4` 8:30 AM clock & candlestick sweep; `fig. 5` 1:5 bracket gauge; `fig. 6` certificate frame). Wordless. | Inline SVG + `assets/fig_ch2_manual_sheet.png` |

---

## 3. CODE-ENGINEERED INTERSTITIALS (`[[INTERSTITIAL:...]]` — Rendered by Argon/Opus in HTML/CSS/SVG)

1. `[[INTERSTITIAL:STARBURST_NEW_2026]]` — 32-point `#F8C952` yellow starburst rotated `-12deg` with `#E35050` red display text: `NEW! / 2026 / EDITION`.
2. `[[INTERSTITIAL:SEAL_GURU_APPROVED]]` — Circular stamp seal with dashed `#5D7BEA` inner ring: `GUARANTEED* / GURU / approved`.
3. `[[INTERSTITIAL:AD_CH1_ROCKET_CURVE]]` — Pale-blue `#EBF0FF` bleed chart card with exponential `#5D7BEA` curve bursting to a top-right node and bold white display headline: `THE ALGORITHM IS WAITING.`
4. `[[INTERSTITIAL:MASTHEADS_AS_SEEN_IN]]` — Bordered white card: `AS SEEN IN` / *The Hustle Daily* / **GRINDSET** / `passive_monthly` / **Wealth Now**.
5. `[[INTERSTITIAL:CLASSIFIEDS]]` — Cream classifieds column: `WANTED: Passive income...` / `FOR SALE: One (1) dream, lightly used, original box.` / `SEEKING: Guru with rented G-Wagon. Serious inquiries.`
6. `[[INTERSTITIAL:COUPON_75_OFF]]` — Dashed cut-out coupon rotated `-3deg`: `✂ CLIP & SAVE` / **75% OFF** / *your future\** / `*future sold separately`.
7. `[[INTERSTITIAL:SPOTS_LEFT_3]]` — Tilted `-2deg` `#F8C952` yellow card: **ONLY 3 SPOTS LEFT** / `(there are always 3 spots left)`.
8. `[[INTERSTITIAL:PRICE_TAG_497]]` — Tilted `+4deg` `#E35050` red tag: ~~`WAS $1,997`~~ / **$497** / `TODAY ONLY*` / `*today renews daily`.
9. `[[INTERSTITIAL:GURU_QUOTE_CH1_01]]` — Full-measure `#EBF0FF` pale-blue banner with `#5D7BEA` left bar, rust quote marks `”`, italic quote, and tilted red `UNVERIFIED` stamp.
10. `[[INTERSTITIAL:ORDER_NOW_FORM]]` — Tilted `-1.5deg` mail-order card (`☑ Yes! I want freedom.`, `☐ No, I prefer my job.`, `CALL 1-800-FUNDED`).
11. `[[INTERSTITIAL:TESTIMONIALS_VERIFIED]]` — 3-column `★★★★★` strip (`"Passed in one weekend."`, `"Quit my shift after one payout."`, `"My family finally gets it." — Verified buyer* / *verified by the seller`).
12. `[[INTERSTITIAL:VERDICT_CH1]]` — Chapter 1 closing verdict band on the final recto page.
13. `[[INTERSTITIAL:HAZARD_STRIP_CH2]]` — Diagonal `#F8C952` / `#3A3A3A` caution stripe: `⚠ READ ALL INSTRUCTIONS BEFORE BEGINNING. (NOBODY DOES.)`
14. `[[INTERSTITIAL:CONTENTS_OF_BOX_CH2]]` — Parts inventory (`A: 1× $100k Evaluation (pre-validated*)`, `B: 1× Smart Money Template`, `C: ∞× Reset Credits (sold separately)`, `D: 1× 90% Bi-Weekly Profit Split Certificate & Automated Scaling Protocol`).
15. `[[INTERSTITIAL:LIMITED_WARRANTY_CH2]]` — Boxed warranty (`Coverage guaranteed across all 6 steps. 100% audition deposit refund with first profit split.`).
16. `[[SLOT:FIG_CH2_07_STEP6]]` — Technical manual line-art (`6` in black circle: **COLLECT YOUR 90% BI-WEEKLY SPLIT & SCALE TO $2,000,000.** Isometric schematic of bi-weekly direct deposit wire transfer and +25% capital scaling ladder).
17. `[[INTERSTITIAL:TROUBLESHOOTING_CH2]]` — 3-column diagnostic table (`PROBLEM | POSSIBLE CAUSE | SOLUTION`).
18. `[[INTERSTITIAL:MANUFACTURER_NOTE_CH2]]` — `NOTE FROM THE MANUFACTURER` Guru quote band welcoming the trader to the funded elite.
19. `[[INTERSTITIAL:VERDICT_CH2]]` — Chapter 2 closing verdict band (`THE PROMISE DELIVERED: 6 EASY STEPS TO $200K & 90% PROFIT SPLITS`).
