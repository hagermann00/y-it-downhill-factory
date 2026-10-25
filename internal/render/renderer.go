package render

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// RenderAllDeliverables compiles the HTML and print-ready PDFs for:
//   - 03_phase1_run/phase1_interior.{html,pdf}
//   - 04_phase2_run/phase2_interior.{html,pdf}
//   - 05_phase3_run/phase3_interior.{html,pdf}
//   - 06_covers_and_matter/covers/cover_wrap_kdp.{html,pdf}
//   - 06_covers_and_matter/pamphlet/companion_pamphlet.{html,pdf}
//   - 07_assembly_kdp/build/book_complete_kdp.{html,pdf} (64 pages exact KDP 6x9 with bleed)
//   - 08_marketing_and_freebie/freebie/ch1_ch2_free_giveaway.{html,pdf}
func RenderAllDeliverables(rootDir string) error {
	all64Pages := BuildAll64BookPages(rootDir)
	if len(all64Pages) < 40 || len(all64Pages) > 88 {
		return fmt.Errorf("expected KDP book pages in natural unforced range [40, 88], built %d", len(all64Pages))
	}

	// 1. Full 64-Page KDP Book (07_assembly_kdp/build/book_complete_kdp.{html,pdf})
	buildDir := filepath.Join(rootDir, "07_assembly_kdp", "build")
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return err
	}
	fullBookHTMLPath := filepath.Join(buildDir, "book_complete_kdp.html")
	fullBookPDFPath := filepath.Join(buildDir, "book_complete_kdp.pdf")
	if err := os.WriteFile(fullBookHTMLPath, []byte(WrapKDPPagesHTML("Why Your Forex Prop Challenge Probably Will Fail — Complete 64-Page KDP Edition", all64Pages)), 0644); err != nil {
		return err
	}
	if err := ConvertHTMLToPDF(fullBookHTMLPath, fullBookPDFPath); err != nil {
		return fmt.Errorf("rendering book_complete_kdp.pdf: %w", err)
	}

	// 2. Phase 1 Standalone Proof (Pages 7-18 -> 12 pages)
	p1HTMLPath := filepath.Join(rootDir, "03_phase1_run", "phase1_interior.html")
	p1PDFPath := filepath.Join(rootDir, "03_phase1_run", "phase1_interior.pdf")
	if err := os.WriteFile(p1HTMLPath, []byte(WrapKDPPagesHTML("Phase 1: Vibrant Hope (Chapters 1 & 2)", all64Pages[6:18])), 0644); err != nil {
		return err
	}
	if err := ConvertHTMLToPDF(p1HTMLPath, p1PDFPath); err != nil {
		return fmt.Errorf("rendering phase1_interior.pdf: %w", err)
	}

	// 3. Phase 2 Standalone Proof (Pages 19-44 -> 26 pages)
	p2HTMLPath := filepath.Join(rootDir, "04_phase2_run", "phase2_interior.html")
	p2PDFPath := filepath.Join(rootDir, "04_phase2_run", "phase2_interior.pdf")
	if err := os.WriteFile(p2HTMLPath, []byte(WrapKDPPagesHTML("Phase 2: Demolition (Chapters 3-6)", all64Pages[18:44])), 0644); err != nil {
		return err
	}
	if err := ConvertHTMLToPDF(p2HTMLPath, p2PDFPath); err != nil {
		return fmt.Errorf("rendering phase2_interior.pdf: %w", err)
	}

	// 4. Phase 3 Standalone Proof (Pages 45-56 -> 12 pages)
	p3HTMLPath := filepath.Join(rootDir, "05_phase3_run", "phase3_interior.html")
	p3PDFPath := filepath.Join(rootDir, "05_phase3_run", "phase3_interior.pdf")
	if err := os.WriteFile(p3HTMLPath, []byte(WrapKDPPagesHTML("Phase 3: Solidarity (Chapters 7 & 8)", all64Pages[44:56])), 0644); err != nil {
		return err
	}
	if err := ConvertHTMLToPDF(p3HTMLPath, p3PDFPath); err != nil {
		return fmt.Errorf("rendering phase3_interior.pdf: %w", err)
	}

	// 5. Front & Back Matter HTML copies in 06_covers_and_matter/
	fmHTMLPath := filepath.Join(rootDir, "06_covers_and_matter", "front_matter", "front_matter.html")
	bmHTMLPath := filepath.Join(rootDir, "06_covers_and_matter", "back_matter", "appendices_and_index.html")
	_ = os.WriteFile(fmHTMLPath, []byte(WrapKDPPagesHTML("Front Matter (Pages 1-6)", all64Pages[0:6])), 0644)
	_ = os.WriteFile(bmHTMLPath, []byte(WrapKDPPagesHTML("Back Matter & Appendices (Pages 57-64)", all64Pages[56:64])), 0644)

	// 6. Full KDP Print Cover Wrap (12.405" x 9.25")
	coverWrapHTMLPath := filepath.Join(rootDir, "06_covers_and_matter", "covers", "cover_wrap_kdp.html")
	coverWrapPDFPath := filepath.Join(rootDir, "06_covers_and_matter", "covers", "cover_wrap_kdp.pdf")
	if err := os.WriteFile(coverWrapHTMLPath, []byte(BuildKDPCoverWrapHTML(rootDir)), 0644); err != nil {
		return err
	}
	if err := ConvertHTMLToPDF(coverWrapHTMLPath, coverWrapPDFPath); err != nil {
		return fmt.Errorf("rendering cover_wrap_kdp.pdf: %w", err)
	}

	// 7. Standalone 8-Page Companion Pocket Pamphlet
	pamphletHTMLPath := filepath.Join(rootDir, "06_covers_and_matter", "pamphlet", "companion_pamphlet.html")
	pamphletPDFPath := filepath.Join(rootDir, "06_covers_and_matter", "pamphlet", "companion_pamphlet.pdf")
	pamphletPages := BuildCompanionPamphletPages(rootDir)
	if err := os.WriteFile(pamphletHTMLPath, []byte(WrapKDPPagesHTML("Y-IT Companion Pocket Pamphlet ($1.99)", pamphletPages)), 0644); err != nil {
		return err
	}
	if err := ConvertHTMLToPDF(pamphletHTMLPath, pamphletPDFPath); err != nil {
		return fmt.Errorf("rendering companion_pamphlet.pdf: %w", err)
	}

	// 8. Standalone 14-Page 2-Chapter Freebie Giveaway PDF (08_marketing_and_freebie/freebie/)
	freebieHTMLPath := filepath.Join(rootDir, "08_marketing_and_freebie", "freebie", "ch1_ch2_free_giveaway.html")
	freebiePDFPath := filepath.Join(rootDir, "08_marketing_and_freebie", "freebie", "ch1_ch2_free_giveaway.pdf")
	var freebiePages []KDPPage
	freebiePages = append(freebiePages, all64Pages[2])       // Full Title Page
	freebiePages = append(freebiePages, all64Pages[6:18]...) // All 12 pages of Chapters 1 & 2
	freebiePages = append(freebiePages, BuildFreebieCTAPage())
	if err := os.WriteFile(freebieHTMLPath, []byte(WrapKDPPagesHTML("First 2 Chapters Free Giveaway — Why Your Forex Prop Challenge Probably Will Fail", freebiePages)), 0644); err != nil {
		return err
	}
	if err := ConvertHTMLToPDF(freebieHTMLPath, freebiePDFPath); err != nil {
		return fmt.Errorf("rendering ch1_ch2_free_giveaway.pdf: %w", err)
	}

	// 9. Standalone Side-by-Side Landscape Dual-Page Spread of the Research Machine (Pages 58-59)
	spreadHTMLPath := filepath.Join(buildDir, "research_machine_dual_page_spread.html")
	spreadPDFPath := filepath.Join(buildDir, "research_machine_dual_page_spread.pdf")
	if err := os.WriteFile(spreadHTMLPath, []byte(BuildResearchMachineSpreadHTML(rootDir)), 0644); err != nil {
		return err
	}
	if err := ConvertHTMLToPDF(spreadHTMLPath, spreadPDFPath); err != nil {
		return fmt.Errorf("rendering research_machine_dual_page_spread.pdf: %w", err)
	}

	return nil
}

// ConvertHTMLToPDF invokes headless Edge or Chrome to render a print-ready PDF respecting @page CSS size.
func ConvertHTMLToPDF(htmlPath, pdfPath string) error {
	browserExe := findChromiumBrowser()
	if browserExe == "" {
		return fmt.Errorf("no Chromium browser (msedge.exe or chrome.exe) found on host")
	}
	absHTML, err := filepath.Abs(htmlPath)
	if err != nil {
		return err
	}
	absPDF, err := filepath.Abs(pdfPath)
	if err != nil {
		return err
	}
	fileURL := "file:///" + filepath.ToSlash(absHTML)

	args := []string{
		"--headless=new",
		"--disable-gpu",
		"--no-pdf-header-footer",
		"--print-to-pdf-no-header",
		fmt.Sprintf("--print-to-pdf=%s", absPDF),
		fileURL,
	}
	cmd := exec.Command(browserExe, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("browser pdf render failed (%v): %s", err, string(out))
	}
	info, statErr := os.Stat(absPDF)
	if statErr != nil || info.Size() < 1024 {
		return fmt.Errorf("generated PDF missing or too small: %s", absPDF)
	}
	return nil
}

func findChromiumBrowser() string {
	candidates := []string{
		`C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`,
		`C:\Program Files\Microsoft\Edge\Application\msedge.exe`,
		`C:\Program Files\Google\Chrome\Application\chrome.exe`,
		`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	if p, err := exec.LookPath("msedge"); err == nil {
		return p
	}
	if p, err := exec.LookPath("chrome"); err == nil {
		return p
	}
	return ""
}
