package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"yit-downhill-factory/internal/gate"
	"yit-downhill-factory/internal/pipeline"
	"yit-downhill-factory/internal/render"
)

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting working directory: %v\n", err)
		os.Exit(1)
	}

	cmd := "run"
	if len(os.Args) > 1 {
		cmd = strings.ToLower(strings.TrimSpace(os.Args[1]))
	}

	switch cmd {
	case "run":
		if err := runAll(rootDir); err != nil {
			fmt.Fprintf(os.Stderr, "run error: %v\n", err)
			os.Exit(1)
		}
	case "verify":
		if err := runVerify(rootDir); err != nil {
			fmt.Fprintf(os.Stderr, "verify error: %v\n", err)
			os.Exit(1)
		}
	case "render":
		if err := render.RenderAllDeliverables(rootDir); err != nil {
			fmt.Fprintf(os.Stderr, "render error: %v\n", err)
			os.Exit(1)
		}
		printPDFSummary(rootDir)
	case "status":
		b, err := os.ReadFile(filepath.Join(rootDir, "00_trigger", "_status.md"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "status read error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q (expected: run | verify | render | status)\n", cmd)
		os.Exit(1)
	}
}

func runAll(rootDir string) error {
	start := time.Now()
	fmt.Println("=== Y-IT AUTONOMOUS DOWNHILL BOOK FACTORY (Go 1.26) ===")
	fmt.Println("Stage 01: Building Research Backplane & Stockpile...")
	if err := pipeline.BuildStage01Research(rootDir); err != nil {
		return fmt.Errorf("stage 01 failed: %w", err)
	}

	fmt.Println("Stage 02-05: Building Pre-Draft Spine & Phase 1/2/3 Manuscripts + SHA-256 Assets...")
	if err := pipeline.BuildStage02Through05(rootDir); err != nil {
		return fmt.Errorf("stage 02-05 failed: %w", err)
	}

	fmt.Println("Stage 06, 08, 09: Building 5 Cover Concepts, Front/Back Matter, Companion Pamphlet, Marketing Pack & Podcap Stub...")
	if err := pipeline.BuildStage06And08And09(rootDir); err != nil {
		return fmt.Errorf("stage 06/08/09 failed: %w", err)
	}

	fmt.Println("Stage 07: Rendering Full-Bleed KDP Book, Phase PDFs, Cover Wrap, Companion Pamphlet & Freebie PDF via Headless Edge...")
	if err := render.RenderAllDeliverables(rootDir); err != nil {
		return fmt.Errorf("stage 07 render failed: %w", err)
	}

	fmt.Println("Running Programmatic Gates G1-G7...")
	report, err := gate.RunAllGates(rootDir)
	if err != nil {
		return fmt.Errorf("gate execution error: %w", err)
	}
	if err := gate.WriteGateReports(rootDir, report); err != nil {
		return fmt.Errorf("writing gate reports error: %w", err)
	}

	if err := updateTriggerStatus(rootDir, report); err != nil {
		return fmt.Errorf("updating trigger status error: %w", err)
	}

	printGateSummary(report.Gates)
	printPDFSummary(rootDir)
	fmt.Printf("Downhill factory run finished in %s.\n", time.Since(start).Round(time.Millisecond))

	for _, r := range report.Gates {
		if r.ExitCode != 0 {
			return fmt.Errorf("gate %s exited with code %d (%d failures)", r.GateID, r.ExitCode, len(r.Failures))
		}
	}
	return nil
}

func runVerify(rootDir string) error {
	report, err := gate.RunAllGates(rootDir)
	if err != nil {
		return err
	}
	if err := gate.WriteGateReports(rootDir, report); err != nil {
		return err
	}
	printGateSummary(report.Gates)
	for _, r := range report.Gates {
		if r.ExitCode != 0 {
			return fmt.Errorf("gate %s exited with code %d (%d failures)", r.GateID, r.ExitCode, len(r.Failures))
		}
	}
	return nil
}

func printGateSummary(results []gate.GateResult) {
	fmt.Println("\n--- PROGRAMMATIC GATE SUMMARY (G1-G7) ---")
	for _, r := range results {
		fmt.Printf("[%s] %-38s | ExitCode: %d | Checks: %3d | Failures: %d\n",
			r.GateID, r.Name, r.ExitCode, r.ChecksRun, len(r.Failures))
		for _, f := range r.Failures {
			fmt.Printf("   ! FAILURE: %s\n", f)
		}
	}
}

func printPDFSummary(rootDir string) {
	fmt.Println("\n--- RENDERED PDF DELIVERABLES ---")
	pdfs := []string{
		"03_phase1_run/phase1_interior.pdf",
		"04_phase2_run/phase2_interior.pdf",
		"05_phase3_run/phase3_interior.pdf",
		"06_covers_and_matter/covers/cover_wrap_kdp.pdf",
		"06_covers_and_matter/pamphlet/companion_pamphlet.pdf",
		"07_assembly_kdp/build/book_complete_kdp.pdf",
		"07_assembly_kdp/build/research_machine_dual_page_spread.pdf",
		"08_marketing_and_freebie/freebie/ch1_ch2_free_giveaway.pdf",
	}
	for _, rel := range pdfs {
		full := filepath.Join(rootDir, filepath.FromSlash(rel))
		info, err := os.Stat(full)
		if err != nil {
			fmt.Printf("  %-58s | MISSING (%v)\n", rel, err)
			continue
		}
		pages, err := gate.CountPDFPages(full)
		if err != nil {
			fmt.Printf("  %-58s | %8d bytes | pages err: %v\n", rel, info.Size(), err)
			continue
		}
		fmt.Printf("  %-58s | %8d bytes | %2d pages\n", rel, info.Size(), pages)
	}
}

func updateTriggerStatus(rootDir string, report gate.FullGateReport) error {
	bookPDF := filepath.Join(rootDir, "07_assembly_kdp", "build", "book_complete_kdp.pdf")
	bookPages, _ := gate.CountPDFPages(bookPDF)
	spreadPDF := filepath.Join(rootDir, "07_assembly_kdp", "build", "research_machine_dual_page_spread.pdf")
	spreadPages, _ := gate.CountPDFPages(spreadPDF)
	coverPDF := filepath.Join(rootDir, "06_covers_and_matter", "covers", "cover_wrap_kdp.pdf")
	coverPages, _ := gate.CountPDFPages(coverPDF)
	pamphletPDF := filepath.Join(rootDir, "06_covers_and_matter", "pamphlet", "companion_pamphlet.pdf")
	pamphletPages, _ := gate.CountPDFPages(pamphletPDF)
	freebiePDF := filepath.Join(rootDir, "08_marketing_and_freebie", "freebie", "ch1_ch2_free_giveaway.pdf")
	freebiePages, _ := gate.CountPDFPages(freebiePDF)

	var sb strings.Builder
	sb.WriteString("# Downhill Pipeline Execution Ledger (`00_trigger/_status.md`)\n\n")
	sb.WriteString("- **Topic ID:** `T01` (`forex-prop-firms`)\n")
	sb.WriteString("- **Title:** *Y-IT: Forex Trading & Prop Firms — Why You're Exit Liquidity in a $7.5 Trillion Casino*\n")
	sb.WriteString("- **Chapter 5 Title:** *Exit Liquidity — Who Actually Owns the $7.5 Trillion Casino—and Why Your Stop-Loss Is Their Lunch*\n")
	sb.WriteString(fmt.Sprintf("- **Last Run Timestamp:** `%s`\n\n", time.Now().Format(time.RFC3339)))

	sb.WriteString("## Stage Gate Results (`G1`–`G7`)\n\n")
	sb.WriteString("| Gate | Gate Name | Exit Code | Checks Run | Failures | Summary |\n")
	sb.WriteString("|---|---|---|---|---|---|\n")
	for _, r := range report.Gates {
		sb.WriteString(fmt.Sprintf("| `%s` | %s | `%d` | `%d` | `%d` | `%s` |\n",
			r.GateID, r.Name, r.ExitCode, r.ChecksRun, len(r.Failures), r.SummaryLine))
	}

	sb.WriteString("\n## Rendered PDF Deliverables\n\n")
	sb.WriteString("| Deliverable | Path | Page Count |\n")
	sb.WriteString("|---|---|---|\n")
	sb.WriteString(fmt.Sprintf("| Full 64-Page KDP Book (`6.125\" x 9.25\"` Full Bleed) | `07_assembly_kdp/build/book_complete_kdp.pdf` | `%d` |\n", bookPages))
	sb.WriteString(fmt.Sprintf("| Dual-Page Research Machine Spread (`12.25\" x 9.25\"` Open Spread) | `07_assembly_kdp/build/research_machine_dual_page_spread.pdf` | `%d` |\n", spreadPages))
	sb.WriteString(fmt.Sprintf("| KDP Full Cover Wrap (`12.405\" x 9.25\"`) | `06_covers_and_matter/covers/cover_wrap_kdp.pdf` | `%d` |\n", coverPages))
	sb.WriteString(fmt.Sprintf("| Standalone Companion Pocket Pamphlet (`$1.99`, Modules `P1..P7`) | `06_covers_and_matter/pamphlet/companion_pamphlet.pdf` | `%d` |\n", pamphletPages))
	sb.WriteString(fmt.Sprintf("| 2-Chapter Freebie Giveaway PDF (Ch 1 + Ch 2) | `08_marketing_and_freebie/freebie/ch1_ch2_free_giveaway.pdf` | `%d` |\n", freebiePages))

	return os.WriteFile(filepath.Join(rootDir, "00_trigger", "_status.md"), []byte(sb.String()), 0644)
}
