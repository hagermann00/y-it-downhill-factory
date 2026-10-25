# Stage 07: KDP 50–80 Page Assembly & Stitching (`07_assembly_kdp/CONTEXT.md`)

## 1. Role in the Downhill Pipeline
Stage 07 stitches Front Matter (`06`), Phase 1 (`03`), Phase 2 (`04`), Phase 3 (`05`), and Back Matter/Appendices (`06`) into a single unified KDP `6" × 9"` print-ready interior (`6.125" × 9.25"` with bleed), synchronizes Table of Contents page numbers, enforces recto (odd-page) chapter openings, compiles the final PDF, and runs **Gate G7**.

## 2. Output Artifacts
- `build/book_complete_kdp.html` (Unified print-ready HTML/CSS interior with `@page` KDP bleed geometry)
- `build/book_complete_kdp.pdf` (Final KDP print-ready interior PDF, strictly **50 to 80 pages**)
- `build/kdp_preflight_report.json` (Machine-readable preflight verification: exact page count, trim dimensions, chapter recto page numbers, asset hash manifest, font & bleed checks)
- `build/gate_report.md` (Full programmatic Gate `G1`–`G7` execution transcript)

## 3. Programmatic Gate (`G7`)
- **Gate G7 (`KDP Print Geometry & Page Count Gate`):**
  1. Verifies `build/book_complete_kdp.pdf` exists and has a valid `/Type /Pages /Count N` (or counted page objects) where **`50 <= N <= 80`**.
  2. Verifies `06_covers_and_matter/covers/cover_wrap_kdp.pdf` and `06_covers_and_matter/pamphlet/companion_pamphlet.pdf` exist and are non-empty PDFs.
  3. Verifies every chapter (`Ch 1`..`Ch 8`) starts on an odd (recto) page number.
  4. Verifies all `internal/catalog` assets referenced in the HTML resolve on disk with matching SHA-256 sidecars.
