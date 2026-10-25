# Layer 1 Workspace Router (`CONTEXT.md`) — Y-IT Downhill Book Factory

> **ICM Layer:** Layer 1 (`CONTEXT.md` — *"Where do I go and how is this workspace routed?"*)  
> **Protocol:** Interpretable Context Methodology (ICM) / Model Workspace Protocol (MWP, `arXiv:2603.16021`)  
> **Target Token Load:** `~300 tokens` on workspace entry

---

## 1. Workspace Objective
Transform a single triggered Y-IT topic (`00_trigger/_topic.yaml`) into a fully audited, multi-source-validated **64-page Full-Bleed KDP Book**, **KDP Cover Wrap**, **Companion Pocket Pamphlet (`$1.99`)**, **2-Chapter Freebie Giveaway PDF**, **5-Angle Marketing Pack**, and **25-Minute Dual-Persona Podcap Script**.

## 2. The 5-Layer ICM Context Map

| Layer | Location | Role Question | Token Budget | Mutability |
|---|---|---|---|---|
| **L0: Identity** | `CLAUDE.md` | *"Where am I and what are the global laws?"* | `~600 tok` | Static (Factory) |
| **L1: Router** | `CONTEXT.md` *(this file)* | *"Where do I go next in the pipeline?"* | `~300 tok` | Static (Factory) |
| **L2: Stage Contract** | `00_trigger/CONTEXT.md` .. `09_podcap_stub/CONTEXT.md` | *"What is my single job in this folder?"* | `~300–500 tok` | Static (Factory) |
| **L3: References** | `_canon/` (`contracts/`, `prompts/`, `assets/`) | *"What schemas, voice rules, and specs apply?"* | `1,500–4,000 tok` *(loaded selectively)* | Static (Factory) |
| **L4: Working Artifacts** | Files inside `00_trigger/` .. `09_podcap_stub/` | *"What topic data & drafts am I working with?"* | `2,000–6,000 tok` *(scoped to current + upstream stage)* | Dynamic (Product) |

## 3. Sequential Routing Table (One Stage, One Job)

1. **`00_trigger/`** $\rightarrow$ Read `_topic.yaml` & `_status.md`. Verify topic slug, 5 Chapter 4 molds, and KDP geometry.
2. **`01_research_backplane/`** *(Hybrid Pipeline Stage + ICM Knowledge Bundle)* $\rightarrow$ Execute intense, multi-angle primary source research. Save raw HTML/text captures in `pages/` (`SHA-256`), answer all 38 questions (`questions.yaml`), extract verbatim passages (`rows.yaml`, `evidence.yaml`), map pitch-vs-reality deltas (`friction_register.yaml`), and write `deep_validation_audit.md`. **Must pass Gates `G1`–`G4` before Stage 02.**
3. **`02_predraft_spine/`** $\rightarrow$ Compute the 4-Metric Gate Scorecard (`The Stake`, `The Wait`, `The Odds`, `The Payoff`) from `01_research_backplane/rows.yaml` and emit `phase-1..3/packet.yaml`. **Must pass Gate `G5`.**
4. **`03_phase1_run/`** $\rightarrow$ Draft Chapters 1–2 (*Vibrant Hope*, Corporate Sincerity, Claim rows `C-###` only).
5. **`04_phase2_run/`** $\rightarrow$ Draft Chapters 3–6 (*Demolition*, Sardonic Authority in Ch 3/5/6; zero-joke, zero-PosiBot Aged Amber Coroner Dossier in Ch 4).
6. **`05_phase3_run/`** $\rightarrow$ Draft Chapters 7–8 (*Solidarity & Quiet Reckoning*, 4 gate-passing pivots in Ch 7, Ch 7 $\rightarrow$ Ch 8 seam rule, barstool armor in Ch 8). **Stages 03–05 must pass Gate `G6`.**
7. **`06_covers_and_matter/`** $\rightarrow$ Build 5 cover concepts, KDP `12.405" × 9.25"` cover wrap, front/back matter, and `$1.99` Companion Pocket Pamphlet.
8. **`07_assembly_kdp/`** $\rightarrow$ Assemble full-bleed `6.125" × 9.25"` 64-page KDP interior with safe body margins and margin-crossing PosiBot/charts/callouts. **Must pass Gate `G7`.**
9. **`08_marketing_and_freebie/`** $\rightarrow$ Render standalone 2-Chapter Freebie PDF + 5 landing copy angles + 3-email drip.
10. **`09_podcap_stub/`** $\rightarrow$ Produce 25-minute dual-persona (`ANALYST` + `BELIEVER`) Podcap script + ADR 0038 tiered bundle JSON.
