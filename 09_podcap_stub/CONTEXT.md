# Stage 09: Podcap 25-Minute Dual-Persona Script & Tiered Bundle Stub (`09_podcap_stub/CONTEXT.md`)

## 1. Role in the Downhill Pipeline
Stage 09 implements the **ADR 0038 ("Podcap — Audio-First Product Format & Tiered Sales Bundle")** downhill compartment:
- **Podcap Format:** A 25-minute, two-voice conversational breakdown (*The Forensic Analyst* + *The Recovering Believer / Storyteller*) walking through the book's exact audited numbers (`[R-001]`..`[R-042]`), the 5 Chapter 4 Trench Cases, the 4-Metric Gate Scorecard, and the 4 Chapter 7 Practical Pivots.
- **Tiered Sales Bundle Architecture (`tiered_sales_bundle.json`):**
  - **Tier 0 (Free Lead Magnet):** First 2 Chapters PDF (`08_marketing_and_freebie/freebie/ch1_ch2_free_giveaway.pdf`) + 3-minute Podcap Teaser clip.
  - **Tier 1 (Core Digital / Print — `$9.99` eBook / `$14.99` KDP Paperback):** Full 64-page Y-IT Book.
  - **Tier 2 (The Autopsy + Podcap Bundle — `$19.99`):** Full Book + 25-Minute Podcap Audio Episode + Companion Pocket Pamphlet PDF + 90-Day Paper-Trade Spreadsheet Template.

## 2. Output Artifacts
- `podcap_25min_script.md` (Full timestamped `00:00`–`25:00` dual-host production script citing real research rows)
- `tiered_sales_bundle.json` (Machine-readable product SKU, pricing, and deliverable manifest per ADR 0038)
- `tts_chunk_manifest.json` (Pre-segmented voice synthesis turn queue for deterministic dual-voice TTS rendering)
