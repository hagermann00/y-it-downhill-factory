# Stage 02: Pre-Draft Spine (`02_predraft_spine/CONTEXT.md`)

## 1. Role in the Downhill Pipeline
Stage 02 bridges raw audited research (`01_research_backplane/`) and phased manuscript writing (`03_phase1_run` through `05_phase3_run`).
It slices the research stockpile into three strict phase compartments (`phase-1`, `phase-2`, `phase-3`) so that Phase 1 never sees audited failure rates and Phases 2 & 3 only cite verified row IDs.

## 2. Output Artifacts
- `intake.md` (Topic brief, selected 5 Ch 4 case molds, and 4-Metric Gate computation summary)
- `opus-brief.md` (Voice & pacing calibration directive across all 8 chapters)
- `phase-1/tables.md`, `phase-1/packet.yaml`, `phase-1/predraft.md` (Phase 1 Claim-Only Spine)
- `phase-2/tables.md`, `phase-2/packet.yaml`, `phase-2/predraft.md` (Phase 2 Demolition & 4-Metric Gate Spine)
- `phase-3/tables.md`, `phase-3/packet.yaml`, `phase-3/predraft.md` (Phase 3 Practical Pivots & Stubborn Fool Spine)

## 3. Programmatic Gates (`G5`–`G6`)
- **Gate G5 (`Phase Packet Checker`):** Verifies `phase-{1,2,3}/packet.yaml` structure, checks that Phase 1 cites only `claim` rows (zero audited failure stats), and checks that Phases 2 & 3 cite valid `row_id` entries from `01_research_backplane/rows.yaml`.
- **Gate G6 (`Pre-Draft Structure & Number Trace`):** Verifies all required section headings in `phase-{1,2,3}/predraft.md`, ensures Chapter 4 contains 5 composite profiles from 5 distinct molds, ensures Chapter 7 contains 3–5 gate-scored pivots, and verifies that every statistic in Phases 2 & 3 traces to an audited row ID.
