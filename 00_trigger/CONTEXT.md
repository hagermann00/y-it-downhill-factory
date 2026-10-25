# Stage 00: Topic Trigger (`00_trigger/CONTEXT.md`)

## 1. Role in the Downhill Pipeline
Stage 00 is the **Topic-Only Entry Point** (modeled on `FFyit/_topic.md` and `_status.md`, upgraded to typed YAML + live Markdown status table).
An operator or autonomous dispatcher drops a single topic definition into `00_trigger/_topic.yaml` and runs `yit-factory.exe run`.

## 2. Input & Output Artifacts
- **Primary Input:** `00_trigger/_topic.yaml`
  - Defines `topic_id` (`T01`), `slug` (`forex-prop-firms`), `title` (`Forex Trading & Prop Firms`), `subtitle`, `blueprint_cover_noun` (`FOREX PROP CHALLENGE`), `category`, `quarantined_topics_excluded` (`[dropshipping]`), and target KDP geometry (`6x9`, `50-80 pages`).
- **Live State Output:** `00_trigger/_status.md`
  - Automatically updated by `yit-factory.exe` as each downhill stage (`00` $\rightarrow$ `09`) and programmatic gate (`G1` $\rightarrow$ `G7`) executes.

## 3. Gate Requirement
- Verified at startup by `yit-factory.exe`: `topic_id`, `slug`, and `title` must be non-empty and must not match any quarantined topic (`dropshipping`).
