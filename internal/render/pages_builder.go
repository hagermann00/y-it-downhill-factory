package render

import (
	"fmt"
	"path/filepath"
	"strings"
)

// KDPPage represents one exact 6.125" x 9.25" full-bleed page in the KDP book interior.
type KDPPage struct {
	PageNum     int
	ThemeClass  string // "theme-front", "theme-phase1", "theme-flatpack", "theme-phase2", "theme-amber-ch4", "theme-phase3", "theme-back"
	RunningHead string
	BodyHTML    string
	FooterNote  string
}

// WrapKDPPagesHTML wraps a slice of KDPPage structs into a self-contained print-ready HTML document.
// Regular body text respects strict KDP safe margins (0.60in gutter / 0.54in outside), while
// images, PosiBot cards, interstitials, and verdict bands tastefully break out across the margin line
// or bleed to the trim edge.
func WrapKDPPagesHTML(docTitle string, pages []KDPPage) string {
	var sb strings.Builder
	sb.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>` + docTitle + `</title>
<style>
  @page {
    size: 6.125in 9.25in;
    margin: 0;
  }
  * {
    box-sizing: border-box;
    -webkit-print-color-adjust: exact !important;
    print-color-adjust: exact !important;
  }
  html, body {
    margin: 0;
    padding: 0;
    width: 6.125in;
    background: #FFFFFF;
    font-family: 'Segoe UI', 'Work Sans', -apple-system, sans-serif;
    color: #0F172A;
    font-size: 9.5pt;
    line-height: 1.42;
  }
  /* Full-bleed page canvas (6.125in x 9.25in).
     Safe body text margins are enforced by padding (0.52in top/bottom, 0.62in gutter, 0.54in outer). */
  .kdp-page {
    width: 6.125in;
    height: 9.25in;
    padding: 0.50in 0.54in 0.48in 0.62in;
    position: relative;
    overflow: hidden;
    page-break-after: always;
    break-after: page;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .kdp-page.verso-page {
    padding: 0.50in 0.62in 0.48in 0.54in;
  }
  .kdp-page:last-child {
    page-break-after: auto;
    break-after: auto;
  }

  /* Full-Bleed Phase Theme Palettes & Edge Bleed Bars */
  .theme-front {
    background: radial-gradient(circle at 50% 18%, #FFFFFF 0%, #F5F2EB 100%);
    color: #0F172A;
    border-top: 0.14in solid #0F172A;
    border-bottom: 0.10in solid #0F172A;
  }
  .theme-phase1 {
    background: linear-gradient(180deg, #FFFDF7 0%, #FEFCE8 100%);
    color: #0F172A;
    border-top: 0.14in solid #0D9488;
    border-right: 0.08in solid #F59E0B;
  }
  .theme-flatpack {
    background: linear-gradient(180deg, #F4F7FB 0%, #E0F2FE 100%);
    color: #0F172A;
    border-top: 0.14in solid #2563EB;
    border-left: 0.08in solid #1D4ED8;
  }
  .theme-phase2 {
    background: linear-gradient(180deg, #F8FAFC 0%, #F1F5F9 100%);
    color: #0F172A;
    border-top: 0.14in solid #DC2626;
  }
  .theme-amber-ch4 {
    background: linear-gradient(180deg, #EFE5CE 0%, #E6D5B8 100%);
    color: #1C1917;
    border-top: 0.16in solid #78350F;
    border-bottom: 0.10in solid #78350F;
    font-family: 'Consolas', 'IBM Plex Mono', monospace;
    font-size: 9.0pt;
  }
  .theme-phase3 {
    background: linear-gradient(180deg, #F7FAF7 0%, #ECFDF5 100%);
    color: #0F172A;
    border-top: 0.14in solid #15803D;
  }
  .theme-back {
    background: #F8FAFC;
    color: #0F172A;
    border-top: 0.12in solid #334155;
  }

  .page-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1.5px solid #CBD5E1;
    padding-bottom: 4px;
    margin-bottom: 8px;
    font-family: 'Consolas', 'IBM Plex Mono', monospace;
    font-size: 7.4pt;
    font-weight: 700;
    letter-spacing: 0.6px;
    text-transform: uppercase;
    color: #475569;
  }
  .theme-amber-ch4 .page-header {
    border-bottom: 2px solid #78350F;
    color: #78350F;
  }
  .page-content {
    flex: 1;
    overflow: visible;
  }
  .page-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-top: 1px solid #CBD5E1;
    padding-top: 5px;
    margin-top: 6px;
    font-family: 'Consolas', 'IBM Plex Mono', monospace;
    font-size: 7.4pt;
    color: #475569;
  }
  .theme-amber-ch4 .page-footer {
    border-top: 1.5px solid #78350F;
    color: #78350F;
  }

  /* Regular body text strictly respects the inner measure */
  h1 {
    font-family: 'Impact', 'Big Shoulders Display', 'Arial Black', sans-serif;
    font-size: 20.5pt;
    line-height: 1.06;
    margin: 0 0 7px 0;
    letter-spacing: 0.4px;
    text-transform: uppercase;
  }
  h2 {
    font-family: 'Impact', 'Big Shoulders Display', 'Arial Black', sans-serif;
    font-size: 13pt;
    line-height: 1.14;
    margin: 7px 0 4px 0;
    color: #0F172A;
    text-transform: uppercase;
  }
  h3 {
    font-size: 10pt;
    font-weight: 800;
    margin: 5px 0 3px 0;
    color: #1E293B;
  }
  p {
    margin: 0 0 6px 0;
    text-align: justify;
  }
  .phase-badge {
    display: inline-block;
    padding: 2px 8px;
    border-radius: 3px;
    font-family: 'Consolas', monospace;
    font-size: 7.6pt;
    font-weight: 700;
    color: #FFFFFF;
    background: #0F172A;
    margin-bottom: 5px;
  }

  /* TASTEFUL MARGIN-CROSSING INTERSTITIALS, IMAGES, POSIBOT & VERDICT BANDS:
     These elements step 0.22in to 0.54in outside the body text column into the outer margin
     or full-bleed edge while keeping their internal text safely padded. */
  .callout-guru {
    background: #FEF3C7;
    border: 2px solid #D97706;
    border-left: 8px solid #D97706;
    border-radius: 6px;
    padding: 8px 14px;
    margin: 8px -0.22in;
    font-size: 8.9pt;
    box-shadow: 3px 4px 0px rgba(217, 119, 6, 0.22);
  }
  .callout-rc {
    background: #FEE2E2;
    border: 2px solid #DC2626;
    border-left: 8px solid #991B1B;
    border-radius: 6px;
    padding: 8px 14px;
    margin: 8px -0.22in;
    font-size: 8.9pt;
    box-shadow: 3px 4px 0px rgba(153, 27, 27, 0.20);
  }
  .callout-pivot {
    background: #DCFCE7;
    border: 2px solid #16A34A;
    border-left: 8px solid #15803D;
    border-radius: 6px;
    padding: 8px 14px;
    margin: 7px -0.22in;
    font-size: 8.8pt;
    box-shadow: 3px 4px 0px rgba(21, 128, 61, 0.20);
  }
  /* Verdict band bleeds all the way across the outer trim edge */
  .verdict-band {
    background: #0F172A;
    color: #F8FAFC;
    padding: 9px 0.54in 9px 0.62in;
    margin: 9px -0.54in 0 -0.62in;
    font-family: 'Consolas', 'IBM Plex Mono', monospace;
    font-size: 8.3pt;
    border-top: 3px solid #F59E0B;
    border-bottom: 3px solid #F59E0B;
  }
  .missing-step-box {
    border: 3px dashed #DC2626;
    background: #FEF2F2;
    padding: 12px 16px;
    margin: 10px -0.24in;
    text-align: center;
    box-shadow: 4px 4px 0px rgba(220, 38, 38, 0.18);
  }
  .dossier-box {
    border: 2px solid #78350F;
    background: #F5EBD6;
    padding: 8px 13px;
    margin: 6px -0.18in;
    box-shadow: 3px 3px 0px rgba(120, 53, 15, 0.18);
  }
  .stamp-badge {
    display: inline-block;
    border: 2px solid #991B1B;
    color: #991B1B;
    padding: 1px 7px;
    font-family: 'Consolas', monospace;
    font-weight: 900;
    font-size: 8pt;
    text-transform: uppercase;
    transform: rotate(-1.2deg);
  }
  table.yit-table {
    width: 100%;
    border-collapse: collapse;
    margin: 6px 0;
    font-size: 8.1pt;
  }
  table.yit-table th {
    background: #0F172A;
    color: #FFFFFF;
    text-align: left;
    padding: 4px 6px;
    font-family: 'Consolas', monospace;
    font-size: 7.7pt;
  }
  table.yit-table td {
    border: 1px solid #CBD5E1;
    padding: 4px 6px;
    vertical-align: top;
  }
  .theme-amber-ch4 table.yit-table th {
    background: #78350F;
  }
  .theme-amber-ch4 table.yit-table td {
    border: 1px solid #A16207;
  }
  /* Images and PosiBot cards step tastefully across the margin line (-0.26in) */
  .asset-frame {
    width: calc(100% + 0.52in);
    margin: 7px -0.26in;
    text-align: center;
  }
  .asset-frame img {
    width: 100%;
    max-height: 2.38in;
    object-fit: contain;
    border-radius: 5px;
    border: 1.5px solid #94A3B8;
    background: #FFFFFF;
    box-shadow: 3px 4px 0px rgba(15, 23, 42, 0.16);
  }
  .gate-strip {
    font-family: 'Consolas', monospace;
    font-size: 7.9pt;
    font-weight: 700;
    background: #0F172A;
    color: #F8FAFC;
    padding: 5px 12px;
    border-radius: 4px;
    margin: 5px -0.16in;
    border-left: 5px solid #22C55E;
  }
  .strike {
    text-decoration: line-through;
    color: #94A3B8;
  }
</style>
</head>
<body>
`)

	for _, p := range pages {
		parity := "VERSO · EVEN"
		parityClass := "verso-page"
		if p.PageNum%2 == 1 {
			parity = "RECTO · ODD"
			parityClass = "recto-page"
		}
		sb.WriteString(fmt.Sprintf(`<section class="kdp-page %s %s" data-page="%d">
  <div class="page-header">
    <span>%s</span>
    <span>[%s]</span>
  </div>
  <div class="page-content">
    %s
  </div>
  <div class="page-footer">
    <span>%s</span>
    <span>PAGE %02d OF 64</span>
  </div>
</section>
`, p.ThemeClass, parityClass, p.PageNum, p.RunningHead, parity, p.BodyHTML, p.FooterNote, p.PageNum))
	}

	sb.WriteString("</body>\n</html>\n")
	return sb.String()
}

func fileURI(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return "file:///" + filepath.ToSlash(abs)
}

// BuildAll64BookPages returns all 64 KDP interior pages in exact folio order (Pages 1..64).
func BuildAll64BookPages(rootDir string) []KDPPage {
	var pages []KDPPage
	pages = append(pages, buildFrontMatter6Pages(rootDir)...)
	pages = append(pages, buildPhase1Pages7To18(rootDir)...)
	pages = append(pages, buildPhase2Pages19To44(rootDir)...)
	pages = append(pages, buildPhase3Pages45To56(rootDir)...)
	pages = append(pages, buildBackMatterPages57To64(rootDir)...)
	return pages
}

func buildFrontMatter6Pages(rootDir string) []KDPPage {
	blueprintImg := fileURI(filepath.Join(rootDir, "_canon", "assets", "covers_blueprint", "blueprint_cover_populated_ref.jpg"))
	sinkingImg := fileURI(filepath.Join(rootDir, "_canon", "assets", "covers_blueprint", "sinking_ship_posibot_cover_v1.jpg"))

	return []KDPPage{
		{
			PageNum:     1,
			ThemeClass:  "theme-front",
			RunningHead: "Y-IT SERIES · HALF-TITLE",
			FooterNote:  "IIWII IMPRINT // SERIES TAGLINE LOCK",
			BodyHTML: `
<div style="margin-top: 1.5in; text-align: center;">
  <div class="phase-badge">Y-IT FORENSIC FIELD GUIDE SERIES · ISSUE #01</div>
  <h1 style="font-size: 28pt; line-height: 1.25; margin-top: 18px;">
    Why It <span class="strike">Could.</span> <span class="strike">Would.</span> <span class="strike">Has.</span> <span class="strike">Might.</span><br>
    <span style="color: #DC2626; font-size: 32pt;">Probably Won't.</span><br>
    <span style="font-size: 36pt; letter-spacing: 2px;">WORK</span>
  </h1>
  <p style="text-align: center; font-size: 16pt; font-style: italic; font-weight: 700; color: #475569; margin-top: 14px;">
    Yeah… about that.
  </p>
  <hr style="width: 50%; margin: 28px auto; border: 1px solid #CBD5E1;">
  <p style="text-align: center; font-family: 'Consolas', monospace; font-size: 9pt; color: #334155;">
    VOLUME 01: FOREX TRADING &amp; PROP FIRMS<br>
    Audited CFTC, BIS, ESMA/FCA &amp; BLS Edition
  </p>
</div>`,
		},
		{
			PageNum:     2,
			ThemeClass:  "theme-front",
			RunningHead: "FRONTISPIECE · BLUEPRINT ARCHITECTURE",
			FooterNote:  "CANONICAL COVER PLATE · ME ODDEN CAR TOON ARCHIVE",
			BodyHTML: fmt.Sprintf(`
<h3>FRONTISPIECE EXHIBIT: THE ANATOMY OF A MODERN HUSTLE PITCH</h3>
<p>Before you wire <code>$497</code> to an offshore evaluation portal, study the two visual engines that power the retail prop-trading boom: the high-gloss lifestyle promise and the burning galleon underneath.</p>
<div class="asset-frame"><img src="%s" alt="Blueprint Cover Reference" style="max-height: 2.95in;"></div>
<div class="asset-frame"><img src="%s" alt="Sinking Ship PosiBot Reference" style="max-height: 2.75in;"></div>
`, blueprintImg, sinkingImg),
		},
		{
			PageNum:     3,
			ThemeClass:  "theme-front",
			RunningHead: "TITLE PAGE · TOPIC 01",
			FooterNote:  "ITISWHATITIS.BIZ // Y-IT PRESS",
			BodyHTML: `
<div style="margin-top: 0.6in; text-align: center; border: 4px solid #0F172A; padding: 24px 18px; margin-left: -0.18in; margin-right: -0.18in; background: #FFFFFF; box-shadow: 5px 6px 0px #0F172A;">
  <div class="phase-badge" style="background:#DC2626;">AUDITED VERDICT: PROBABLY WON'T WORK (2.4% ODDS)</div>
  <h1 style="font-size: 24pt; margin-top: 14px; color: #0F172A;">
    WHY YOUR FOREX PROP CHALLENGE PROBABLY WILL FAIL
  </h1>
  <div style="background: #FEF08A; border: 2px solid #0F172A; padding: 6px; margin: 12px -10px; font-family: 'Consolas', monospace; font-weight: 900; font-size: 10.2pt;">
    EVEN IF YOUR AI CHAT BOT LIKES YOUR CHANCES!
  </div>
  <p style="text-align: center; font-size: 10.8pt; font-weight: 700; color: #334155; margin: 14px 0;">
    Funded Trader Evaluations, 100x Leverage, Dubai Balcony Laptops, and the $310 Million B-Book Reset Machine
  </p>
  <hr style="border: 1px dashed #94A3B8; margin: 18px 0;">
  <p style="text-align: center; font-family: 'Consolas', monospace; font-size: 8.8pt; color: #1E293B;">
    <strong>We reveal the secret method in the first 2 chapters...</strong><br>
    (And why Step 7 was mysteriously left out of your starter box.)
  </p>
  <p style="text-align: center; font-family: 'Consolas', monospace; font-size: 8.5pt; margin-top: 24px; color: #64748B;">
    BY THE IIWII / Y-IT FORENSIC DESK · OCTOBER 2026
  </p>
</div>`,
		},
		{
			PageNum:     4,
			ThemeClass:  "theme-front",
			RunningHead: "COPYRIGHT & COLOPHON",
			FooterNote:  "KDP 6.00x9.00 TRIM · 64 PAGES · NO SPINE TEXT (<80P RULE)",
			BodyHTML: `
<h2>COPYRIGHT, EDITION &amp; EVIDENTIARY NOTICE</h2>
<p><strong>Copyright © 2026 It Is What It Is (IIWII) / Y-IT Imprint.</strong> All rights reserved.</p>
<p><strong>No Personal Financial, Legal, or Tax Advice:</strong> This publication presents audited regulatory statistics, industry cohort medians, and statutory citations for educational and satirical analysis only. It never gives personal financial, investment, legal, or tax advice (<code>outline-spine.v1.yaml §6</code>). Where a regulatory or tax rule applies (such as Visa/Mastercard 120-day dispute windows or IRC Sections 988 and 1256), the rule is stated plainly alongside the professional credential to consult (CPA, Enrolled Agent, or issuing bank dispute department).</p>
<div class="callout-rc">
  <strong>THE Y-IT EVIDENTIARY STANDARD:</strong><br>
  • <strong>Phase 1 (Chapters 1–2):</strong> Quotes the industry pitch in its own words and numbers (<code>Claim Rows C-001..C-014</code>). Zero audited failure statistics appear in Phase 1.<br>
  • <strong>Phase 2 (Chapters 3–6):</strong> Tests every claim against government and cohort data (<code>Fact Rows F-001..F-014</code>) and 15 documented practitioner cases grouped into 5 composite molds (<code>Case Rows K-001..K-015</code>).<br>
  • <strong>Phase 3 (Chapters 7–8):</strong> Scores concrete career and freelance pivots against the exact same 4-Metric Gate (<code>Fact Rows F-015..F-027</code>).
</div>
<p style="font-family: 'Consolas', monospace; font-size: 8pt; color: #475569; margin-top: 12px;">
  Print Specification: Amazon KDP Paperback · Trim 6.00" × 9.00" · Full Bleed 6.125" × 9.25" · 64 Pages · Spine Width 0.144" (No Spine Text Permitted Under 80 Pages).
</p>`,
		},
		{
			PageNum:     5,
			ThemeClass:  "theme-front",
			RunningHead: "DEDICATION & HOW THIS BOOK WORKS",
			FooterNote:  "RESPECT THE PHASING · THREE ACTS, THREE RULES",
			BodyHTML: `
<div class="callout-guru" style="background:#FFFFFF; border-color:#0F172A; border-left-width:8px;">
  <p style="font-size: 10.6pt; font-style: italic; font-weight: 700; margin:0; text-align:left;">
    "For everyone who ever calculated 1% daily compounding on a cocktail napkin at 1:30 AM, priced out a matte-black Porsche configurator by 1:45 AM, and got stopped out on GBP/JPY by 2:10 AM."
  </p>
</div>
<h2>HOW TO READ THIS THREE-PHASE FIELD GUIDE</h2>
<p>Most books about side hustles make one of two fatal errors: either they are written by an affiliate marketer trying to sell you a <code>$997</code> course, or they are written by a scolding newspaper columnist who thinks anyone who wants to escape a cubicle is an idiot. We don't do either.</p>
<table class="yit-table">
  <tr><th>Phase</th><th>Chapters</th><th>Voice &amp; Evidence Rule</th></tr>
  <tr><td><strong>Phase 1: Vibrant Hope</strong></td><td>Ch 1 – Ch 2<br>(pp. 07–18)</td><td><strong>Deadpan Corporate Sincerity.</strong> We pitch you the dream at full strength using the gurus' own words (<code>C-001..C-014</code>). Zero cynicism, zero failure stats.</td></tr>
  <tr><td><strong>Phase 2: Demolition</strong></td><td>Ch 3 – Ch 6<br>(pp. 19–44)</td><td><strong>Sardonic Authority &amp; Coroner's Autopsy.</strong> We turn the lights on, open the CFTC B-book filings, and run 5 composite case files on aged amber paper.</td></tr>
  <tr><td><strong>Phase 3: Solidarity</strong></td><td>Ch 7 – Ch 8<br>(pp. 45–56)</td><td><strong>Constructive Solidarity &amp; Quiet Reckoning.</strong> Four <code>$38.90–$55.00/hr</code> pivots that pass all 4 gates—plus a roll-cage manual for the stubborn fool who stays.</td></tr>
</table>`,
		},
		{
			PageNum:     6,
			ThemeClass:  "theme-front",
			RunningHead: "TABLE OF CONTENTS · 64-PAGE SYNCHRONIZED FOLIO",
			FooterNote:  "ALL CHAPTERS OPEN ON RECTO (ODD) PAGES",
			BodyHTML: `
<h2>SYNCHRONIZED TABLE OF CONTENTS</h2>
<table class="yit-table">
  <tr><th>Section / Chapter</th><th>Format &amp; Core Content</th><th>Recto Page</th></tr>
  <tr><td><strong>Front Matter</strong></td><td>Series Tagline, Frontispiece, Title, Evidentiary Standard &amp; TOC</td><td><strong>p. 01</strong></td></tr>
  <tr><td><strong>PHASE 1 · VIBRANT HOPE</strong></td><td><strong>Ch 1: The Promise</strong> (Glossy Magazine-Ad Collage)</td><td><strong>p. 07</strong></td></tr>
  <tr><td><strong>Chapter 2</strong></td><td><strong>The Sucker's Roadmap</strong> (Flat-Pack Manual · Missing Step 7)</td><td><strong>p. 13</strong></td></tr>
  <tr><td><strong>PHASE 2 · DEMOLITION</strong></td><td><strong>Ch 3: The Money Pit</strong> ($1,840 Stake, 98.4% B-Book, $6.40/hr Payoff)</td><td><strong>p. 19</strong></td></tr>
  <tr><td><strong>Chapter 4 (Amber Insert)</strong></td><td><strong>Tales from the Trenches</strong> (5 Composite Case Files · Zero Jokes)</td><td><strong>p. 25</strong></td></tr>
  <tr><td><strong>Chapter 5</strong></td><td><strong>Exit Liquidity</strong> ($7.5T BIS Ladder, 6.6% Pass, 2.4% Odds)</td><td><strong>p. 37</strong></td></tr>
  <tr><td><strong>Chapter 6</strong></td><td><strong>The Verdict</strong> (The 4-Metric Gate Scorecard: 0 of 4 Passed)</td><td><strong>p. 41</strong></td></tr>
  <tr><td><strong>PHASE 3 · SOLIDARITY</strong></td><td><strong>Ch 7: The Pivot</strong> (4 Gate-Passed Roles, 3D Papercraft Comic)</td><td><strong>p. 45</strong></td></tr>
  <tr><td><strong>Chapter 8</strong></td><td><strong>For the Stubborn Fool Who Must</strong> (90-Day Test, 120-Day Chargeback)</td><td><strong>p. 51</strong></td></tr>
  <tr><td><strong>BACK MATTER</strong></td><td><strong>Appendix A:</strong> Audited Sources &amp; Tiers (<code>C-001..K-015</code>)</td><td><strong>p. 57</strong></td></tr>
  <tr><td><strong>Appendix B &amp; Disclosures</strong></td><td>20-Term Unvarnished Glossary, <em>Results Typical</em>, IIWII Colophon &amp; Index</td><td><strong>p. 59</strong></td></tr>
</table>`,
		},
	}
}

func buildPhase1Pages7To18(rootDir string) []KDPPage {
	p1Sign1 := fileURI(filepath.Join(rootDir, "03_phase1_run", "assets", "T01-P1-C01-posibot-zero-risk-laptop_v001.svg"))
	p1Sign2 := fileURI(filepath.Join(rootDir, "03_phase1_run", "assets", "T01-P1-C01-posibot-no-inventory-needed_v001.svg"))
	p1Sign3 := fileURI(filepath.Join(rootDir, "03_phase1_run", "assets", "T01-P1-C01-posibot-london-session-money_v001.svg"))
	p2Sign1 := fileURI(filepath.Join(rootDir, "03_phase1_run", "assets", "T01-P1-C02-posibot-six-easy-steps_v001.svg"))

	return []KDPPage{
		{
			PageNum:     7,
			ThemeClass:  "theme-phase1",
			RunningHead: "PHASE 1: VIBRANT HOPE · CHAPTER 1: THE PROMISE",
			FooterNote:  "VOICE: CORPORATE SINCERITY · CLAIM ROWS ONLY [C-001..C-005]",
			BodyHTML: `
<div class="phase-badge" style="background:#0D9488;">PHASE 1 · VIBRANT HOPE // CHAPTER 01 · RECTO OPENER</div>
<h1>CHAPTER 1: THE PROMISE</h1>
<p style="font-family:'Consolas',monospace; font-size:8.4pt; color:#0F766E; font-weight:700;">FORMAT: MAGAZINE-AD LIFESTYLE COLLAGE · WHAT THEY ARE SELLING YOU</p>
<h3>11:42 PM on a Tuesday: The Flannel Pajama Arbitrage</h3>
<p>It is 11:42 PM on a Tuesday (<code>[C-005]</code>). Across the street, your neighbor Greg is ironing a damp polyester-blend dress shirt so he can wake up at 6:15 AM, scrape gray slush off a Honda Civic, and sit in forty-five minutes of brake-light purgatory for the privilege of asking a regional vice president named Todd if he can leave twenty minutes early on Friday for a root canal.</p>
<p>You, meanwhile, are sitting on the edge of your sofa in flannel pajama pants, bathed in the serene sapphire glow of MetaTrader 5 (<code>[C-005]</code>, <code>[C-013]</code>). Somewhere over the Atlantic, the Tokyo–London liquidity bridge is waking up. You don't have a warehouse. You don't have employees. You don't even have pants with a zipper on right now. You tap three glass buttons on your phone, bracket a 40-pip XAU/USD (Gold) displacement candle, set your phone on the nightstand, and wake up at 7:30 AM to a <strong><code>$2,400</code> payout notification</strong> before your toaster pops (<code>[C-005]</code>).</p>
<div class="callout-guru">
  <strong>[COLLAGE ARTIFACT #01 — THE MIDNIGHT NOTIFICATION BANNER]:</strong><br>
  <em>"PAYOUT APPROVED: $2,400.00 USD · Account #FT-200K · Split: 90% Trader / 10% House · Time Elapsed: 11 Minutes."</em> (<code>[C-001]</code>, <code>[C-005]</code>)
</div>`,
		},
		{
			PageNum:     8,
			ThemeClass:  "theme-phase1",
			RunningHead: "CHAPTER 1: THE PROMISE · SECTION 1: THE VISION",
			FooterNote:  "CLAIM CITATIONS: [C-001, C-002, C-003]",
			BodyHTML: fmt.Sprintf(`
<h2>SECTION 1: THE VISION — WHY BEG FOR A RAISE WHEN YOU CAN TRADE SOMEONE ELSE'S $200,000?</h2>
<p>For half a century, the <strong><code>$7.5 trillion-a-day</code></strong> foreign exchange market (<code>[C-002]</code>) kept a velvet rope across the lobby. Even if you could read a candlestick chart like Mozart reading sheet music, pulling life-changing income out of currency pairs required <code>$100,000</code> of your own post-tax savings sitting in a brokerage vault.</p>
<p>Welcome to the <strong>Funded Trader Revolution</strong> (<code>[C-001]</code>). Modern proprietary evaluation firms have blown the vault doors off institutional leverage: pass a simple 2-Step digital audition in <strong>14 days</strong>, and they hand you the keys to a <strong><code>$100,000</code> to <code>$200,000</code> funded trading account</strong>—letting you <strong>keep up to 90%% of every dollar you extract</strong> with <strong>zero personal capital at risk</strong> (<code>[C-001]</code>).</p>
<div class="asset-frame"><img src="%s" alt="PosiBot Signpost Trade 200K"></div>
<p>Think about the math they put on the whiteboard: a modest 5%% gain on a <code>$200,000</code> account is <code>$10,000</code> in a single month (<code>[C-001]</code>, <code>[C-002]</code>). You don't need to find clients. You just take a tiny sip from a <code>$7.5 trillion</code> river.</p>
`, p1Sign1),
		},
		{
			PageNum:     9,
			ThemeClass:  "theme-phase1",
			RunningHead: "CHAPTER 1: THE PROMISE · SECTION 2: THE FEELING",
			FooterNote:  "CLAIM CITATIONS: [C-006, C-013]",
			BodyHTML: fmt.Sprintf(`
<h2>SECTION 2: THE FEELING — ZERO BOXES, ZERO BOSSES, PURE LIQUIDITY</h2>
<p>Compare funded trading against every other "side hustle" cluttering your social feed (<code>[C-006]</code>). E-commerce dropshipping? You're arguing on WhatsApp with a factory in Yiwu over cracked garlic presses at 3:00 AM. Digital marketing agency? You're cold-calling local roofers who hang up on you while eating lunch. Amazon FBA? Your two-car garage is stacked to the rafters with unsold posture correctors.</p>
<p>In foreign exchange, <strong>the Japanese Yen never calls customer support demanding a refund</strong> (<code>[C-006]</code>). The British Pound never leaves a one-star review because FedEx dented the box. Currency pairs never expire on a warehouse shelf, never ghost your invoice, and settle in crisp bi-weekly wire transfers every <strong>14 days</strong> (<code>[C-006]</code>)—whether your MacBook is open at a kitchen counter in Ohio or on a glass balcony overlooking Dubai Marina next to a matte-black G-Wagon (<code>[C-013]</code>).</p>
<div class="asset-frame"><img src="%s" alt="PosiBot Placard No Customers"></div>
`, p1Sign2),
		},
		{
			PageNum:     10,
			ThemeClass:  "theme-phase1",
			RunningHead: "CHAPTER 1: THE PROMISE · SECTION 3: THE PLAN",
			FooterNote:  "CLAIM CITATIONS: [C-002, C-003, C-014]",
			BodyHTML: fmt.Sprintf(`
<h2>SECTION 3: THE PLAN — 45 MINUTES AT THE LONDON–NEW YORK OVERLAP</h2>
<p>Amateur traders burn out because they stare at one-minute candles for fourteen hours a day like caffeinated raccoons. The Funded Professional does the exact opposite. As every VSL pitch explains, you only need <strong>45 minutes a day</strong> during the 8:00 AM–11:00 AM New York–London session overlap (<code>[C-002]</code>) to pull <strong><code>$10,000</code> a month</strong> out of institutional order flow (<code>[C-002]</code>).</p>
<p>And the price of admission costs less than a Saturday night steak dinner: unlock a <code>$100,000</code> Funded Trader Evaluation today for just <strong><code>$497</code></strong> (or test the waters with a <code>$10,000</code> starter challenge for <strong><code>$89</code></strong>)—with a <strong>100%% fee refund</strong> bundled automatically into your very first profit split (<code>[C-003]</code>, <code>[C-014]</code>).</p>
<div class="asset-frame"><img src="%s" alt="PosiBot Megaphone 45 Minutes"></div>
`, p1Sign3),
		},
		{
			PageNum:     11,
			ThemeClass:  "theme-phase1",
			RunningHead: "CHAPTER 1: THE PROMISE · SECTION 4: THE NAME",
			FooterNote:  "CLAIM CITATIONS: [C-004, C-008, C-011]",
			BodyHTML: `
<h2>SECTION 4: THE NAME — DECING "SMART MONEY CONCEPTS" (SMC)</h2>
<p>Why did old-school retail traders struggle back in 2015? Because they were drawing juvenile retail support lines, double bottoms, and RSI oscillators while interbank algorithms hunted their stop-losses for sport. Today, you are armed with <strong>Smart Money Concepts (<code>SMC</code>)</strong> (<code>[C-002]</code>, <code>[C-004]</code>).</p>
<p>You aren't gambling anymore; you are reading the central bank's own footprints: <strong>Institutional Order Blocks</strong>, <strong>Fair Value Gaps (FVGs)</strong>, and the mechanical 9:30 AM New York <strong>Liquidity Sweep</strong> (<code>[C-004]</code>, <code>[C-008]</code>). Once you install the pre-colored indicator template, the chart practically labels the money for you.</p>
<div class="callout-guru">
  <strong>[COLLAGE ARTIFACT #02 — THE NUMBERLESS "ROCKET CHART" AD]:</strong><br>
  <code>[ RETAIL STOP HUNT ZONE ] ──► [ 9:30 AM LIQUIDITY SWEEP ] ──► [ 1:5 R:R INSTITUTIONAL EXPANSION 🚀 ]</code><br>
  <em>"Stop trading like the 95% retail herd. Trade like the Central Bank Algorithm in 45 minutes a morning!"</em> (<code>[C-002]</code>, <code>[C-011]</code>)
</div>`,
		},
		{
			PageNum:     12,
			ThemeClass:  "theme-phase1",
			RunningHead: "CHAPTER 1: THE PROMISE · SECTION 5: THE GURUS",
			FooterNote:  "END OF CHAPTER 1 · CLAIM ROWS ONLY",
			BodyHTML: `
<h2>SECTION 5: THE GURUS &amp; COMPOSITE TESTIMONIALS</h2>
<div class="callout-guru">
  <strong>[GURU PITCH HALL OF FAME — COMPOSITE]:</strong><br>
  <em>"Bro, the interbank delivery algorithm runs on a 90-minute clock. It has to sweep retail liquidity before it expands. Once you pass your <code>$200,000</code> evaluation in 14 days, a single 5% month drops <code>$9,000</code> straight into your account—and the prop firm refunds your <code>$497</code> audition fee on the house. Use code <code>LIQUIDITY20</code> before the checkout timer hits zero."</em> (<code>[C-001]</code>, <code>[C-003]</code>, <code>[C-013]</code>)
</div>
<h3>Composite Testimonial Carousel (<code>[C-007]</code>)</h3>
<p>• <em>"I used to work 50 hours a week in auto parts. After learning the 9:30 AM Liquidity Sweep, I passed my $100,000 evaluation in two weeks and framed my Funded Trader Certificate above my desk!"</em> — Composite Social Proof Banner (<code>[C-007]</code>, <code>[C-008]</code>)</p>
<p>• <em>"No inventory, no clients, no overhead—just my laptop, a cup of coffee at the London-NY overlap, and a 90% profit split every 14 days."</em> — Composite Webinar Slide (<code>[C-006]</code>, <code>[C-007]</code>)</p>
<div class="verdict-band">
  <strong>CHAPTER 1 VERDICT BAND:</strong> <code>$200,000</code> in institutional buying power, <code>90%</code> profit splits, <code>45 minutes</code> a morning, and a <code>100%</code> refundable <code>$497</code> entry fee (<code>[C-001..C-003]</code>). Honestly, why is Greg across the street still ironing his shirt?
</div>`,
		},
		{
			PageNum:     13,
			ThemeClass:  "theme-flatpack",
			RunningHead: "PHASE 1: VIBRANT HOPE · CHAPTER 2: THE SUCKER'S ROADMAP",
			FooterNote:  "FORMAT: FLAT-PACK ASSEMBLY MANUAL · MODEL: FUNDED-TRADER-200K™",
			BodyHTML: `
<div class="phase-badge" style="background:#2563EB;">PHASE 1 · VIBRANT HOPE // CHAPTER 02 · RECTO OPENER</div>
<h1>CHAPTER 2: THE SUCKER'S ROADMAP</h1>
<p style="font-family:'Consolas',monospace; font-size:8.4pt; color:#1D4ED8; font-weight:700;">ASSEMBLY MANUAL: MODEL #FT-200K · ESTIMATED BUILD TIME: 14 DAYS · DIFFICULTY: EASY</p>
<p>Congratulations on your purchase of the <strong>FUNDED-TRADER-200K™ Financial Freedom System</strong>! Please lay out all components on a clean surface and verify your hardware count before beginning assembly.</p>
<table class="yit-table">
  <tr><th>Part #</th><th>Qty</th><th>Component Description (As Sold, <code>[C-003, C-008, C-009, C-010]</code>)</th><th>Quoted Price</th></tr>
  <tr><td><code>PART-A</code></td><td><code>1×</code></td><td><strong>2-Step Prop Evaluation Credential</strong> ($10k starter or $100k flagship; 100% refundable upon first profit split)</td><td><code>$89 – $540</code></td></tr>
  <tr><td><code>PART-B</code></td><td><code>1×</code></td><td><strong>TradingView Charting Screen + Smart-Money Overlay</strong> (Pre-shaded green &amp; red institutional zones)</td><td><code>Included / Trial</code></td></tr>
  <tr><td><code>PART-C</code></td><td><code>1×</code></td><td><strong>MetaTrader 5 / cTrader Mobile Terminal</strong> (Optimized for screenshotting blue rows from luxury car dealerships)</td><td><code>$0.00 App</code></td></tr>
  <tr><td><code>PART-X</code></td><td><code>0×</code></td><td><strong>Instructions for Step 7</strong> <em>(Not included in starter carton)</em></td><td><code>N/A</code></td></tr>
</table>`,
		},
		{
			PageNum:     14,
			ThemeClass:  "theme-flatpack",
			RunningHead: "CHAPTER 2: THE SUCKER'S ROADMAP · STEPS 1 & 2",
			FooterNote:  "FLAT-PACK STEPS 1–2 [C-003, C-008, C-009, C-013]",
			BodyHTML: `
<h2>STEP 1: SELECT YOUR INSTITUTIONAL CAPITAL TIER (<code>[C-003, C-008]</code>)</h2>
<p><strong>Directive:</strong> Navigate to the evaluation checkout portal, paste your favorite YouTuber's <code>20%</code> discount code into the promo box (<code>[C-013]</code>), and select the <code>$100,000</code> or <code>$200,000</code> account size for <code>$497–$540</code> (<code>[C-001]</code>, <code>[C-003]</code>, <code>[C-008]</code>).</p>
<div class="callout-guru">
  <strong>🔧 ASSEMBLY TIP #1:</strong> Don't waste time starting small! A 5% gain on a $10,000 account is only $500, whereas the exact same mouse click on a $200,000 account pays you $9,000 after your 90% split! (<code>[C-001]</code>, <code>[C-002]</code>)
</div>
<h2>STEP 2: SNAP ON THE SMART-MONEY INDICATOR OVERLAY (<code>[C-008, C-009]</code>)</h2>
<p><strong>Directive:</strong> Open TradingView and attach the proprietary Smart-Money Concepts template (<code>[C-008]</code>, <code>[C-009]</code>). Your screen will automatically populate with pastel boxes labeled <code>BOS</code> (Break of Structure), <code>CHoCH</code> (Change of Character), and <code>FVG</code> (Fair Value Gap).</p>
<div class="callout-guru">
  <strong>🔧 SYSTEM™ INSTRUCTIONS (<code>[C-011]</code>):</strong> No economics degree required! When the indicator draws a green box, that means JPMorgan and Goldman Sachs just left their wallet on the table for you.
</div>`,
		},
		{
			PageNum:     15,
			ThemeClass:  "theme-flatpack",
			RunningHead: "CHAPTER 2: THE SUCKER'S ROADMAP · STEPS 3 & 4",
			FooterNote:  "FLAT-PACK STEPS 3–4 [C-008, C-011]",
			BodyHTML: fmt.Sprintf(`
<h2>STEP 3: AWAIT THE 9:30 AM NEW YORK LIQUIDITY SWEEP (<code>[C-008]</code>)</h2>
<p><strong>Directive:</strong> Sit down at 9:25 AM EST with an artisanal cold brew. Watch price dip briefly below the London session low to "sweep uninformed retail stop-losses" before reversing sharply upward (<code>[C-008]</code>).</p>
<h2>STEP 4: INSERT THE 1:5 RISK-TO-REWARD BRACKET ORDER (<code>[C-008, C-011]</code>)</h2>
<p><strong>Directive:</strong> Risk a tidy <code>1%%</code> of your evaluation balance with a tight 5-pip stop loss and set your take-profit target at <code>+5%%</code> (<code>1:5 R:R</code>). Close your laptop by 10:15 AM and head to the gym (<code>[C-002]</code>, <code>[C-008]</code>).</p>
<div class="asset-frame"><img src="%s" alt="PosiBot Clipboard Steps 1 to 6"></div>
`, p2Sign1),
		},
		{
			PageNum:     16,
			ThemeClass:  "theme-flatpack",
			RunningHead: "CHAPTER 2: THE SUCKER'S ROADMAP · STEPS 5 & 6",
			FooterNote:  "FLAT-PACK STEPS 5–6 [C-001, C-006, C-008, C-013]",
			BodyHTML: `
<h2>STEP 5: CLEAR PHASE 1 (+8%) AND PHASE 2 (+5%) IN 14 DAYS (<code>[C-001, C-008]</code>)</h2>
<p><strong>Directive:</strong> Repeat Steps 3 and 4 three times over two weeks. Your <code>1:5</code> risk-to-reward bracket clears the <code>+8%</code> Phase 1 target on Tuesday and the <code>+5%</code> Phase 2 verification target the following Thursday (<code>[C-001]</code>, <code>[C-008]</code>).</p>
<h2>STEP 6: FRAME YOUR GOLD CERTIFICATE &amp; WITHDRAW 90% BI-WEEKLY (<code>[C-006, C-008, C-013]</code>)</h2>
<p><strong>Directive:</strong> Download your gold-embossed JPEG <strong>Funded Trader Certificate</strong> (complete with cursive CEO signature and QR code, <code>[C-008]</code>, <code>[C-013]</code>), collect your <code>100%</code> evaluation fee refund (<code>[C-003]</code>), and schedule automatic <code>90%</code> profit-split withdrawals every 14 days (<code>[C-006]</code>, <code>[C-008]</code>).</p>
<div class="callout-guru">
  <strong>🔧 TROUBLESHOOTING GUIDE (<code>[C-010, C-011]</code>):</strong><br>
  • <em>Q: What if my trade gets stopped out?</em><br>
  • <em>A: Simply wait for the next 90-minute algorithmic macro cycle! With a 1:5 risk-to-reward ratio, you can be wrong four times in a row and still come out ahead on trade five!</em>
</div>`,
		},
		{
			PageNum:     17,
			ThemeClass:  "theme-flatpack",
			RunningHead: "CHAPTER 2: THE SUCKER'S ROADMAP · THE MISSING STEP",
			FooterNote:  "MANDATORY CLIFFHANGER ENDING [C-010]",
			BodyHTML: `
<h2>STEP 7: [MISSING FROM FACTORY CARTON]</h2>
<div class="missing-step-box">
  <h1 style="color:#DC2626; font-size:20pt; margin-bottom:6px;">Step 7 was not provided.</h1>
  <p style="text-align:center; font-family:'Consolas',monospace; font-size:9.2pt; font-weight:700; color:#7F1D1D;">
    [ PART #07 OMITTED FROM STARTER KIT · SEE CFTC &amp; COHORT AUTOPSY ON PAGE 19 ]
  </p>
</div>
<p><strong>Step 7 was not provided</strong> (<code>[C-010]</code>). Specifically, the assembly brochure omits how a human being is supposed to clear an <code>+8%</code> Phase 1 target and a <code>+5%</code> Phase 2 target inside an internal B-book server while handcuffed to a <strong>5% intraday equity trailing drawdown limit</strong> and a negative-expectancy spread-plus-commission drag that mathematically requires a <strong>64% raw win rate</strong> just to break even before the evaluation clock expires (<code>[C-010]</code>).</p>
<div class="verdict-band">
  <strong>CHAPTER 2 VERDICT BAND:</strong> Six breezy steps on glossy blue paper, one giant missing parachute at the bottom of the box (<code>[C-008, C-010]</code>). Turn the page—we're opening the server room door.
</div>`,
		},
		{
			PageNum:     18,
			ThemeClass:  "theme-flatpack",
			RunningHead: "END OF PHASE 1 · SEAM TRANSITION",
			FooterNote:  "END OF FREE 2-CHAPTER PREVIEW · PHASE 2 OPENS ON PAGE 19",
			BodyHTML: `
<div style="margin-top: 1.4in; text-align: center; border: 3px solid #0F172A; padding: 22px; margin-left: -0.18in; margin-right: -0.18in; background: #FFFFFF; box-shadow: 4px 5px 0px #0F172A;">
  <div class="phase-badge" style="background:#0F172A;">END OF PHASE 1 (VIBRANT HOPE)</div>
  <h2 style="font-size: 18pt; margin-top: 10px;">THE BROCHURE ENDS HERE.</h2>
  <p style="text-align: center; font-size: 10pt; margin: 12px 0;">
    Everything you just read in Chapters 1 and 2 was constructed strictly from the industry's own sales pages, VSL scripts, and checkout banners (<code>Claim Rows C-001 through C-014</code>).
  </p>
  <p style="text-align: center; font-size: 10pt; font-weight: 700; color: #DC2626; margin: 12px 0;">
    On Page 19, we bring in the CFTC enforcement accountants, the 420,000-account cohort telemetry, and five coroner's dossiers from the trenches.
  </p>
  <div class="gate-strip" style="text-align:center;">
    NEXT: PHASE 2 · DEMOLITION (CHAPTERS 3, 4, 5 &amp; 6) ──► PAGE 19
  </div>
</div>`,
		},
	}
}
