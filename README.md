# Y-IT Downhill Autonomous Book Factory (`D:\y-it-downhill-factory`)

> **Architecture Baseline:** October 24, 2026 (Temporal Recency Synthesis across June 2026 → October 17, 2026)  
> **Engine Bedrock:** Pure Go 1.26 (`cmd/yit-factory`) · Zero Python · Deterministic Gates `G1`–`G7`  
> **Target Output:** Complete 50–80 Page KDP Print-Ready Book (`6" × 9"` trim, `6.125" × 9.25"` with bleed), Full Cover Wrap (`12.405" × 9.25"`), 2-Chapter Freebie Giveaway, Companion Pamphlet, 5-Variant Marketing Pack, and 25-Minute Dual-Persona Podcap Script.

---

## 1. Overview & Purpose

`D:\y-it-downhill-factory` is a standalone, compartmentalized, downhill book-production factory for the **Y-IT** satirical non-fiction imprint (*"Why It ~~Could.~~ ~~Would.~~ ~~Has.~~ ~~Might.~~ Probably Won't. WORK — Yeah… about that."*).

It merges the best architectural patterns across every iteration of the Y-IT estate while strictly enforcing **temporal recency** whenever earlier iterations conflict with later specifications:

| Era / Source | Date | What Was Retained (Canonical) | What Was Superseded |
| :--- | :--- | :--- | :--- |
| **`I:\My Drive\FFyit`** | 2026-06-20 | Orderly **downhill folder architecture** (`00_trigger` $\rightarrow$ `09_podcap_stub`), self-routing `CONTEXT.md` per stage, and single-file topic trigger (`_topic.yaml` + `_status.md`). | Wooden-box PosiBot description; flat single-voice 8-chapter outline; manual copy-paste workflow. |
| **`ME ODDEN CAR TOON_files`** | 2026-09-18 | **Blueprint Grid Cover Layout** (*"why your [TOPIC] probably will Fail / EVEN IF YOUR AI CHAT BOT LIKES YOUR CHANCES! / We reveal the secret method in the first 2 chapters..."*), **PosiBot Blank-Sign Pose Sheet**, **3D Papercraft Comic Strip** (`[PAGE COMIC — 3D PAPERCRAFT]`), and **Brian's IIWII Author Avatars**. | Unstructured image filenames (now content-addressed via `internal/catalog`). |
| **`D:\y-it-auto-designer`** | 2026-10-06 | **3-Phase Emotional Arc** (Phase 1 *Vibrant Hope*, Phase 2 *Demolition*, Phase 3 *Solidarity*), **38 Canon Questions**, **13 Case Study Molds**, **4-Metric Gate Scorecard** (*the Stake, the Wait, the Odds, the Payoff*), **Ch 4 Standalone Amber Autopsy (`#EFE5CE`)**, **KDP 6×9 Print Geometry (50–80 pages)**, and **ADR 0038 Podcap Audio/Text Tiered Bundle**. | Early "Sovereign Pivot" label (replaced with `[PIVOT: THE PRACTICAL PIVOT]` per estate vocabulary rules). |
| **`D:\y-it-promocopy` & `BUILDOUT_STATE`** | 2026-10-12 | Quarantine of `dropshipping` until canon audit clears; promotion of **Topic 01: "Forex Trading & Prop Firms"** as the #1 live candidate; **5 Landing Copy Angles** (`v1 value`, `v2 risk`, `v3 smart buyer`, `v4 time`, `v5 price`), 3-step email drip, and `cover-bundle.v1.yaml`. | Dropshipping placeholder runs. |
| **`D:\y-it-engine` & `yit_phase1_universal_research_prompt.txt`** | 2026-10-17 | **One Law: *"Models propose, programs decide."*** Deterministic Go gates (`G1`–`G7`), verbatim passage & number verification against saved HTML/text pages, **4-Layer Universal Phase 1 Research Architecture** (`Promotional Presentation`, `Operational Pathway`, `Real-World Friction` / `friction_register.yaml`, `Presentational Encoding`), and **SHA-256 Content-Addressed Asset Catalog**. | Self-attested prose claims without programmatic gate checks. |

---

## 2. Downhill Folder Architecture (`00` $\rightarrow$ `09`)

Each stage reads **only** from `_canon/` and upstream lower-numbered compartments, writes its own artifacts, and must clear its deterministic Go gate before the next stage executes:

```text
D:\y-it-downhill-factory\
├── README.md                          # Factory architecture & quickstart guide
├── CLAUDE.md                          # Agent operating protocol & downhill routing rules
├── 00_BRAND_AND_SYSTEM_CANON.md       # Complete Y-IT Brand, Voice, Visual & Temporal Recency Bible
├── go.mod / go.sum                    # Pure Go 1.26 module (zero CGO, zero Python)
├── cmd\yit-factory\main.go            # Autonomous CLI (`run`, `verify`, `render`, `status`)
├── internal\
│   ├── catalog\                       # Content-addressed SHA-256 image/prompt catalog & SVG/PNG renderer
│   ├── gate\                          # Deterministic programmatic gates G1 through G7
│   └── render\                        # KDP 6x9 HTML/CSS/PDF book, cover wrap & pamphlet compiler
├── _canon\
│   ├── prompts\                       # Versioned canonical prompts (incl. Oct 17 Universal Phase 1 Prompt)
│   ├── contracts\                     # outline-spine.v1.yaml, mold-menu.yaml, voice.md, COVER_SPEC.md
│   └── assets\                        # Canonical visual references & cutouts
│       ├── covers_blueprint\          # Blueprint grid covers, sinking-ship covers, KDP press visual
│       ├── posibot\                   # Blank-sign pose sheet + transparent PosiBot PNG cutouts
│       ├── comics_papercraft\         # 3D papercraft comic strip reference
│       └── author_iiwii\              # Brian's IIWII illustrated avatars & badges
├── 00_trigger\                        # Stage 00: Topic-only trigger (_topic.yaml) & live pipeline _status.md
├── 01_research_backplane\             # Stage 01: 4-Layer Universal Research + 38-Q Stockpile + Saved Pages (Gates G1-G4)
├── 02_predraft_spine\                 # Stage 02: Intake, Phase 1/2/3 Packets, Pre-Drafts & Opus Brief (Gates G5-G6)
├── 03_phase1_run\                     # Stage 03: Phase 1 (Vibrant Hope - Ch 1 & Ch 2) Manuscript, Art & PDF
├── 04_phase2_run\                     # Stage 04: Phase 2 (Demolition - Ch 3, Ch 4 Amber Autopsy, Ch 5, Ch 6)
├── 05_phase3_run\                     # Stage 05: Phase 3 (Solidarity - Ch 7 Pivots & Ch 8 Stubborn Fool)
├── 06_covers_and_matter\              # Stage 06: 5 Cover Concepts, KDP Cover Wrap, Front/Back Matter & Pamphlet
├── 07_assembly_kdp\                   # Stage 07: Full 50-80 Page Stitched KDP Book PDF & Gate G7 Preflight
├── 08_marketing_and_freebie\          # Stage 08: Ch 1-2 Freebie PDF, 5 Landing Angles, Email Drip & Ad Pack
└── 09_podcap_stub\                    # Stage 09: 25-Min Dual-Persona Podcap Script & Tiered Sales Bundle Stub
```

---

## 3. Running the Factory CLI

Build and run the pure-Go binary from `D:\y-it-downhill-factory`:

```powershell
# 1. Build the factory binary
go build -o yit-factory.exe ./cmd/yit-factory

# 2. Run the full downhill production pipeline from 00_trigger/_topic.yaml
.\yit-factory.exe run

# 3. Run all programmatic verification gates (G1 through G7)
.\yit-factory.exe verify

# 4. Check downhill stage status and artifact inventory
.\yit-factory.exe status
```
