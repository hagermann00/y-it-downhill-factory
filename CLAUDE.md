# CLAUDE.md — Downhill Factory Agent Protocol

Read `00_BRAND_AND_SYSTEM_CANON.md` first, then read `00_trigger/_status.md` to determine the active topic and downhill stage.

---

## 1. Core Operating Invariants

1. **Models Propose, Programs Decide (`y-it-engine` Spec §0):**
   - Never write words like "verified", "passed", "complete", or "clean" by hand to claim your own output is valid.
   - A downhill stage is complete **only** when `yit-factory.exe verify --stage <N>` exits `0` and writes a deterministic JSON/Markdown gate receipt.
2. **Downhill One-Way Data Flow (`FFyit` Architecture):**
   - Flow moves strictly downhill from `00_trigger` $\rightarrow$ `01_research_backplane` $\rightarrow$ `02_predraft_spine` $\rightarrow$ `03_phase1_run` / `04_phase2_run` / `05_phase3_run` $\rightarrow$ `06_covers_and_matter` $\rightarrow$ `07_assembly_kdp` $\rightarrow$ `08_marketing_and_freebie` $\rightarrow$ `09_podcap_stub`.
   - Every stage directory contains its own `CONTEXT.md` defining exact input files, output files, schema rules, and gate checks.
   - Never invent a statistic in Phases 2 or 3 (`04_phase2_run`, `05_phase3_run`). Every numeric figure in Phases 2 and 3 must trace to an audited `row_id` in `01_research_backplane/rows.yaml` backed by a saved page in `01_research_backplane/pages/` and a verbatim passage in `01_research_backplane/evidence.yaml`.
3. **Phase 1 Claim Boundary (`yit_phase1_universal_research_prompt.v1.txt` & ADR 0046):**
   - Phase 1 (`03_phase1_run`, Chapters 1 & 2) uses **Corporate Sincerity** voice and **Claim Rows Only**.
   - Audited failure statistics are **forbidden** in Phase 1; pitch numbers (e.g. "$10,000/month", "$200,000 funded account", "keep 90% of profits") are permitted only when framed as the guru/vendor pitch.
   - Chapter 2 (*The Sucker's Roadmap*) must end abruptly on the missing step: `"Step N+1 was not provided."`
4. **Chapter 4 Forensic Autopsy Boundary (ADR 0047):**
   - Chapter 4 (*Tales from the Trenches*) is rendered on aged amber case-file paper (`#EFE5CE`), has **zero PosiBot appearances**, **zero jokes** (*"The numbers are the joke"*), and presents **5 composite profiles** drawn from 5 distinct molds of the 13-mold menu (`_canon/contracts/mold-menu.yaml`), each synthesized from $\ge 3$ documented real-world cases.
5. **Go-Only Tooling:**
   - All orchestration, gate verification, asset cataloging, and PDF compilation run via the pure-Go binary `yit-factory.exe` (`cmd/yit-factory`). Do not introduce Python scripts.
6. **Vocabulary Rule:**
   - Never use the deprecated word `Sovereign` anywhere in manuscripts, prompts, code, or metadata. Use `[PIVOT: THE PRACTICAL PIVOT]` for Pivot callout boxes.
