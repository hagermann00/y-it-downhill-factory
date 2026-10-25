# Stage 01: Research Backplane (`01_research_backplane/CONTEXT.md`)

## 1. Role in the Downhill Pipeline
Stage 01 executes maximal-depth, source-backed research combining:
1. **The Oct 17, 2026 Universal Phase 1 Research Prompt (`_canon/prompts/yit_phase1_universal_research_prompt.v1.txt`):**
   - **Section A:** Promotional Presentation (headline claims, lifestyle pitch, guru pricing ladder).
   - **Section B:** Operational Pathway (`Step 1`..`Step N` execution sequence).
   - **Section C:** Real-World Friction (`friction_register.yaml` mapping every divergence between pitch and reality).
   - **Section D:** Presentational Encoding (`visual_persona_atlas` of guru archetypes and visual cues).
2. **The `y-it-engine` 38-Question Stockpile (`questions.yaml`, `rows.yaml`, `evidence.yaml`, `pages/`):**
   - Answers all 38 canon questions (`P1-Q1..P1-Q12`, `P2-Q1..P2-Q12`, `P3-Q1..P3-Q14`).
   - Every audited row in `rows.yaml` links to a verbatim passage in `evidence.yaml` and a saved source file in `pages/<source_id>.html` or `.txt`.

## 2. Output Artifacts
- `phase1_research.md` (4-Section Universal Research Dossier)
- `friction_register.yaml` (Structured A-vs-B divergence register)
- `source_register.csv` (Tabular source index with tiers and dates)
- `phase1_gate_report.md` (Gate report for `G-01`..`G-14`)
- `questions.yaml` (The 38 canon questions for the active topic)
- `rows.yaml` (All claim rows and audited rows with figures, units, questions, and mold tags)
- `evidence.yaml` (Verbatim quote passages, URLs, SHA-256 page hashes, and dates)
- `sources.yaml` & `sources-appendix.md` (Tiered bibliography)
- `affiliates.yaml` & `disqualified.yaml` (Affiliate conflict register & rejected sources)
- `coverage.md` & `tables/*.yaml` (38/38 question coverage matrix & sliced YAML tables)
- `stockpile.lock` (Deterministic SHA-256 lockfile of the research stockpile)

## 3. Programmatic Gates (`G1`–`G4`)
- **Gate G1 (`Questions`):** Verifies all 38 question IDs (`P1-Q1..P1-Q12`, `P2-Q1..P2-Q12`, `P3-Q1..P3-Q14`) exist and match `_canon/contracts/outline-spine.v1.yaml`.
- **Gate G2 (`Evidence & Verbatim Passage`):** Verifies every row in `rows.yaml` has a matching `evidence.yaml` record, a saved file in `pages/`, a verbatim substring match (after whitespace/quote normalization) inside the saved page, and exact presence of `row.figure` inside the passage. Also validates `friction_register.yaml`.
- **Gate G3 (`Corroboration & Tiers`):** Verifies source tier distribution (`Tier 1`, `Tier 2`, `Tier 3`), `affiliates.yaml`, and `disqualified.yaml`.
- **Gate G4 (`38/38 Coverage & 5 Molds`):** Verifies every one of the 38 questions has $\ge 1$ answering row, and at least 5 distinct molds from `_canon/contracts/mold-menu.yaml` each have $\ge 3$ documented case records for Chapter 4.
