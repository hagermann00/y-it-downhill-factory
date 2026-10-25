package catalog

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

var slugCleanRe = regexp.MustCompile(`[^a-z0-9-]+`)

// AssetRecord represents one content-addressed visual asset and its prompt provenance.
type AssetRecord struct {
	AssetID        string    `json:"asset_id" yaml:"asset_id"`
	Topic          int       `json:"topic" yaml:"topic"`
	Phase          int       `json:"phase" yaml:"phase"`
	Chapter        int       `json:"chapter" yaml:"chapter"`
	Slot           string    `json:"slot" yaml:"slot"`
	Slug           string    `json:"slug" yaml:"slug"`
	Version        int       `json:"version" yaml:"version"`
	Extension      string    `json:"extension" yaml:"extension"`
	RelativePath   string    `json:"relative_path" yaml:"relative_path"`
	PromptPath     string    `json:"prompt_path" yaml:"prompt_path"`
	MetaPath       string    `json:"meta_path" yaml:"meta_path"`
	SHA256         string    `json:"sha256" yaml:"sha256"`
	PromptSHA256   string    `json:"prompt_sha256" yaml:"prompt_sha256"`
	DeliveryMethod string    `json:"delivery_method,omitempty" yaml:"delivery_method,omitempty"`
	ChartType      string    `json:"chart_type,omitempty" yaml:"chart_type,omitempty"`
	CitedRows      []string  `json:"cited_rows,omitempty" yaml:"cited_rows,omitempty"`
	CreatedAt      time.Time `json:"created_at" yaml:"created_at"`
}

// Snapshot holds the complete deterministic manifest of all cataloged assets in a stage or build.
type Snapshot struct {
	TopicSlug   string        `json:"topic_slug"`
	GeneratedAt string        `json:"generated_at"`
	TotalAssets int           `json:"total_assets"`
	Assets      []AssetRecord `json:"assets"`
}

// FormatAssetID returns the canonical Y-IT asset identifier:
// T{topic:02}-P{phase}-C{chapter:02}-{slot}-{slug}_v{ver:03}
func FormatAssetID(topic, phase, chapter int, slot, slug string, version int) string {
	cleanSlot := sanitizeToken(slot)
	cleanSlug := sanitizeToken(slug)
	if version < 1 {
		version = 1
	}
	return fmt.Sprintf("T%02d-P%d-C%02d-%s-%s_v%03d", topic, phase, chapter, cleanSlot, cleanSlug, version)
}

func sanitizeToken(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, "_", "-")
	s = strings.ReplaceAll(s, " ", "-")
	s = slugCleanRe.ReplaceAllString(s, "")
	s = strings.Trim(s, "-")
	if s == "" {
		return "item"
	}
	return s
}

// WriteCatalogAsset writes the asset bytes, its .prompt.md sidecar, and its .json metadata sidecar.
func WriteCatalogAsset(baseDir string, topic, phase, chapter int, slot, slug string, version int, ext string, content []byte, promptText string, deliveryMethod, chartType string, citedRows []string) (AssetRecord, error) {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return AssetRecord{}, err
	}
	ext = strings.TrimPrefix(strings.ToLower(ext), ".")
	if ext == "" {
		ext = "svg"
	}
	assetID := FormatAssetID(topic, phase, chapter, slot, slug, version)
	fileName := fmt.Sprintf("%s.%s", assetID, ext)
	promptName := fmt.Sprintf("%s.prompt.md", assetID)
	metaName := fmt.Sprintf("%s.json", assetID)

	fullPath := filepath.Join(baseDir, fileName)
	promptPath := filepath.Join(baseDir, promptName)
	metaPath := filepath.Join(baseDir, metaName)

	contentHash := sha256Hex(content)
	promptBody := fmt.Sprintf("# Prompt Sidecar: %s\n\n- **Asset ID:** `%s`\n- **Phase / Chapter:** Phase %d, Chapter %d\n- **Slot:** `%s`\n- **Content SHA-256:** `%s`\n\n## Canonical Generation Prompt\n\n%s\n",
		assetID, assetID, phase, chapter, slot, contentHash, strings.TrimSpace(promptText))
	promptHash := sha256Hex([]byte(promptBody))

	if err := os.WriteFile(fullPath, content, 0644); err != nil {
		return AssetRecord{}, err
	}
	if err := os.WriteFile(promptPath, []byte(promptBody), 0644); err != nil {
		return AssetRecord{}, err
	}

	rec := AssetRecord{
		AssetID:        assetID,
		Topic:          topic,
		Phase:          phase,
		Chapter:        chapter,
		Slot:           slot,
		Slug:           slug,
		Version:        version,
		Extension:      ext,
		RelativePath:   fileName,
		PromptPath:     promptName,
		MetaPath:       metaName,
		SHA256:         contentHash,
		PromptSHA256:   promptHash,
		DeliveryMethod: deliveryMethod,
		ChartType:      chartType,
		CitedRows:      citedRows,
		CreatedAt:      time.Date(2026, 10, 24, 9, 15, 0, 0, time.UTC),
	}

	metaBytes, err := json.MarshalIndent(rec, "", "  ")
	if err != nil {
		return AssetRecord{}, err
	}
	if err := os.WriteFile(metaPath, metaBytes, 0644); err != nil {
		return AssetRecord{}, err
	}
	return rec, nil
}

// WriteSnapshot writes a deterministic snapshot.json in baseDir.
func WriteSnapshot(baseDir, topicSlug string, records []AssetRecord) error {
	sorted := make([]AssetRecord, len(records))
	copy(sorted, records)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].AssetID < sorted[j].AssetID
	})
	snap := Snapshot{
		TopicSlug:   topicSlug,
		GeneratedAt: "2026-10-24T09:15:00Z",
		TotalAssets: len(sorted),
		Assets:      sorted,
	}
	data, err := json.MarshalIndent(snap, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(baseDir, "snapshot.json"), data, 0644)
}

func sha256Hex(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

// RenderPosiBotSVG creates a high-contrast vector illustration of PosiBot (kraft-tan body,
// dark CRT face, glowing cyan fixed smile/eyes) using an approved delivery method and <=6 word sign.
func RenderPosiBotSVG(headline, subline, method, phaseAccent string) []byte {
	if phaseAccent == "" {
		phaseAccent = "#0D9488"
	}
	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 560 220" width="560" height="220">
  <defs>
    <linearGradient id="bgGrad" x1="0" y1="0" x2="1" y2="1">
      <stop offset="0%%" stop-color="#FAF8F5"/>
      <stop offset="100%%" stop-color="#F1ECE1"/>
    </linearGradient>
    <filter id="shadow" x="-5%%" y="-5%%" width="110%%" height="110%%">
      <feDropShadow dx="2" dy="3" stdDeviation="2" flood-color="#000000" flood-opacity="0.18"/>
    </filter>
  </defs>
  <rect x="4" y="4" width="552" height="212" rx="10" fill="url(#bgGrad)" stroke="%s" stroke-width="2.5"/>
  <rect x="14" y="14" width="150" height="22" rx="4" fill="%s"/>
  <text x="89" y="29" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="10" font-weight="700" fill="#FFFFFF">POSIBOT // %s</text>

  <!-- PosiBot Kraft-Tan Body & Fixed Cyan CRT Face -->
  <g transform="translate(38, 46)" filter="url(#shadow)">
    <!-- Antenna -->
    <line x1="55" y1="0" x2="55" y2="14" stroke="#374151" stroke-width="3"/>
    <circle cx="55" cy="0" r="5" fill="#00F0FF"/>
    <!-- Kraft-Tan Head -->
    <rect x="15" y="14" width="80" height="62" rx="10" fill="#C8A27A" stroke="#785A3C" stroke-width="2.5"/>
    <!-- Dark CRT Screen -->
    <rect x="23" y="22" width="64" height="46" rx="6" fill="#0F172A" stroke="#334155" stroke-width="1.5"/>
    <!-- Glowing Cyan Eyes (never change) -->
    <circle cx="41" cy="39" r="5.5" fill="#00F0FF"/>
    <circle cx="69" cy="39" r="5.5" fill="#00F0FF"/>
    <!-- Glowing Cyan Fixed Wide Smile -->
    <path d="M 37 52 Q 55 64 73 52" fill="none" stroke="#00F0FF" stroke-width="3.5" stroke-linecap="round"/>
    <!-- Torso -->
    <rect x="22" y="80" width="66" height="54" rx="6" fill="#B9936C" stroke="#785A3C" stroke-width="2.5"/>
    <rect x="34" y="92" width="42" height="24" rx="3" fill="#E5D3B3" stroke="#785A3C" stroke-width="1.2"/>
    <line x1="40" y1="100" x2="70" y2="100" stroke="#785A3C" stroke-width="1.5"/>
    <line x1="40" y1="108" x2="62" y2="108" stroke="#785A3C" stroke-width="1.5"/>
    <!-- Arms holding sign pole -->
    <path d="M 88 94 L 132 76" stroke="#785A3C" stroke-width="6" stroke-linecap="round"/>
    <!-- Legs -->
    <rect x="33" y="134" width="14" height="22" rx="3" fill="#6B7280"/>
    <rect x="63" y="134" width="14" height="22" rx="3" fill="#6B7280"/>
  </g>

  <!-- Sign / Placard Held by PosiBot -->
  <g transform="translate(172, 34)" filter="url(#shadow)">
    <line x1="0" y1="86" x2="20" y2="86" stroke="#785A3C" stroke-width="5"/>
    <rect x="16" y="8" width="352" height="148" rx="8" fill="#FFFBEB" stroke="#1E293B" stroke-width="3"/>
    <rect x="24" y="16" width="336" height="132" rx="4" fill="none" stroke="#D97706" stroke-width="1.2" stroke-dasharray="5,3"/>
    <text x="192" y="68" text-anchor="middle" font-family="'Big Shoulders Display', 'Impact', sans-serif" font-size="28" font-weight="900" fill="#111827" letter-spacing="0.8">%s</text>
    <line x1="48" y1="86" x2="336" y2="86" stroke="#CBD5E1" stroke-width="1.5"/>
    <text x="192" y="114" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="11.5" font-weight="600" fill="#475569">%s</text>
    <text x="192" y="136" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="9.5" fill="#94A3B8">[DELIVERY: %s · FIXED CYAN SMILE ACTIVE]</text>
  </g>
</svg>`, phaseAccent, phaseAccent, strings.ToUpper(method), escapeXML(headline), escapeXML(subline), strings.ToUpper(method))
	return []byte(svg)
}

// RenderChartSVG renders one of the 14 canonical Y-IT chart types with question title,
// answer callout, data bars/curves, and bottom source citation line.
func RenderChartSVG(chartType, questionTitle, answerCallout, sourceLine string, labels []string, values []float64, displayVals []string, accentColor string) []byte {
	if accentColor == "" {
		accentColor = "#DC2626"
	}
	var body strings.Builder
	maxVal := 1.0
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	for i, label := range labels {
		if i >= 6 {
			break
		}
		y := 86 + i*34
		val := 0.0
		if i < len(values) {
			val = values[i]
		}
		disp := ""
		if i < len(displayVals) {
			disp = displayVals[i]
		}
		barW := int((val / maxVal) * 250.0)
		if barW < 8 {
			barW = 8
		}
		barColor := accentColor
		if i == 0 && (chartType == "gap bar" || chartType == "waterfall" || chartType == "before/after cliff") {
			barColor = "#0D9488"
		}
		body.WriteString(fmt.Sprintf(`
    <text x="18" y="%d" font-family="'Work Sans', sans-serif" font-size="11.5" font-weight="600" fill="#1E293B">%s</text>
    <rect x="195" y="%d" width="%d" height="20" rx="3" fill="%s"/>
    <text x="%d" y="%d" font-family="'IBM Plex Mono', monospace" font-size="11" font-weight="700" fill="#0F172A">%s</text>`,
			y+14, escapeXML(label), y, barW, barColor, 203+barW, y+14, escapeXML(disp)))
	}

	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 560 320" width="560" height="320">
  <rect x="2" y="2" width="556" height="316" rx="8" fill="#FFFFFF" stroke="#1E293B" stroke-width="2"/>
  <rect x="2" y="2" width="556" height="58" rx="8" fill="#0F172A"/>
  <text x="18" y="24" font-family="'IBM Plex Mono', monospace" font-size="10" font-weight="700" fill="#38BDF8">CHART EXHIBIT // %s</text>
  <text x="18" y="45" font-family="'Big Shoulders Display', sans-serif" font-size="18" font-weight="800" fill="#FFFFFF">%s</text>
  <rect x="372" y="14" width="172" height="34" rx="5" fill="%s"/>
  <text x="458" y="35" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="11" font-weight="700" fill="#FFFFFF">%s</text>
  %s
  <line x1="16" y1="286" x2="544" y2="286" stroke="#CBD5E1" stroke-width="1"/>
  <text x="18" y="304" font-family="'IBM Plex Mono', monospace" font-size="9.5" fill="#475569">Source: %s</text>
</svg>`, strings.ToUpper(chartType), escapeXML(questionTitle), accentColor, escapeXML(answerCallout), body.String(), escapeXML(sourceLine))
	return []byte(svg)
}

// RenderPapercraftComicSVG renders the 4-panel [PAGE COMIC — 3D PAPERCRAFT] visual relief strip
// inspired by ME ODDEN CAR TOON_files/comic_3d_papercraft_strip_ref.jpg.
func RenderPapercraftComicSVG(topicTitle string) []byte {
	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 560 360" width="560" height="360">
  <rect x="2" y="2" width="556" height="356" rx="10" fill="#F8FAFC" stroke="#0F172A" stroke-width="2.5"/>
  <rect x="14" y="12" width="532" height="32" rx="5" fill="#0F172A"/>
  <text x="280" y="33" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="12" font-weight="700" fill="#F8FAFC">[PAGE COMIC — 3D PAPERCRAFT DIORAMA] // %s</text>

  <!-- Panel 1 -->
  <g transform="translate(16, 54)">
    <rect x="0" y="0" width="256" height="136" rx="6" fill="#FEF3C7" stroke="#92400E" stroke-width="2"/>
    <text x="10" y="18" font-family="'IBM Plex Mono', monospace" font-size="9.5" font-weight="700" fill="#92400E">PANEL 01 · THE $100K PAPER CROWN</text>
    <polygon points="40,115 80,55 120,115" fill="#F59E0B" stroke="#78350F" stroke-width="1.5"/>
    <rect x="135" y="48" width="105" height="66" rx="4" fill="#FFFBEB" stroke="#78350F" stroke-width="1.5"/>
    <text x="187" y="75" text-anchor="middle" font-family="'Work Sans', sans-serif" font-size="10.5" font-weight="700" fill="#1E293B">"I MANAGE $200K!"</text>
    <text x="187" y="95" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="8.5" fill="#78350F">(Simulated Demo Server #4)</text>
  </g>

  <!-- Panel 2 -->
  <g transform="translate(288, 54)">
    <rect x="0" y="0" width="256" height="136" rx="6" fill="#FEE2E2" stroke="#991B1B" stroke-width="2"/>
    <text x="10" y="18" font-family="'IBM Plex Mono', monospace" font-size="9.5" font-weight="700" fill="#991B1B">PANEL 02 · THE 5%% TRAILING TRAPDOOR</text>
    <path d="M 24 55 L 95 45 L 115 112 L 220 112" fill="none" stroke="#DC2626" stroke-width="4"/>
    <rect x="122" y="38" width="120" height="56" rx="4" fill="#FFFFFF" stroke="#991B1B" stroke-width="1.5"/>
    <text x="182" y="62" text-anchor="middle" font-family="'Work Sans', sans-serif" font-size="10" font-weight="700" fill="#991B1B">ACCOUNT BREACHED</text>
    <text x="182" y="80" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="8.5" fill="#475569">Reset Fee: $540.00</text>
  </g>

  <!-- Panel 3 -->
  <g transform="translate(16, 204)">
    <rect x="0" y="0" width="256" height="136" rx="6" fill="#E0F2FE" stroke="#075985" stroke-width="2"/>
    <text x="10" y="18" font-family="'IBM Plex Mono', monospace" font-size="9.5" font-weight="700" fill="#075985">PANEL 03 · POSIBOT FOLDS THE BROCHURE</text>
    <rect x="24" y="42" width="58" height="72" rx="6" fill="#C8A27A" stroke="#785A3C" stroke-width="2"/>
    <rect x="32" y="50" width="42" height="28" rx="4" fill="#0F172A"/>
    <circle cx="44" cy="61" r="3.5" fill="#00F0FF"/>
    <circle cx="62" cy="61" r="3.5" fill="#00F0FF"/>
    <path d="M 42 70 Q 53 76 64 70" fill="none" stroke="#00F0FF" stroke-width="2.5"/>
    <rect x="96" y="44" width="145" height="72" rx="4" fill="#FFFFFF" stroke="#075985" stroke-width="1.5"/>
    <text x="168" y="72" text-anchor="middle" font-family="'Big Shoulders Display', sans-serif" font-size="15" font-weight="800" fill="#0F172A">"SELL THE SHOVELS!"</text>
    <text x="168" y="94" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="8.5" fill="#0369A1">(Or get paid $42/hr for the SQL)</text>
  </g>

  <!-- Panel 4 -->
  <g transform="translate(288, 204)">
    <rect x="0" y="0" width="256" height="136" rx="6" fill="#DCFCE7" stroke="#166534" stroke-width="2"/>
    <text x="10" y="18" font-family="'IBM Plex Mono', monospace" font-size="9.5" font-weight="700" fill="#166534">PANEL 04 · THE PRACTICAL PIVOT DESK</text>
    <rect x="22" y="42" width="212" height="74" rx="6" fill="#FFFFFF" stroke="#166534" stroke-width="1.8"/>
    <text x="128" y="68" text-anchor="middle" font-family="'Work Sans', sans-serif" font-size="11" font-weight="700" fill="#14532D">4 OF 4 GATES PASSED</text>
    <text x="128" y="88" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="9" fill="#166534">Stake: $0 · Payoff: $44.50/hr</text>
    <text x="128" y="104" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="8.5" fill="#475569">Sleep Schedule: Intact</text>
  </g>
</svg>`, escapeXML(strings.ToUpper(topicTitle)))
	return []byte(svg)
}

// RenderCaseCaricatureSVG renders a dignified editorial silhouette caricature (ADR 0015 / ADR 0039)
// dramatizing both the composite persona and the structural situation of their Chapter 4 mold.
func RenderCaseCaricatureSVG(moldCode, personaName, disposition, sceneTitle, propLeft, propRight, captionLine string) []byte {
	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 560 155" width="560" height="155">
  <rect x="2" y="2" width="556" height="151" rx="6" fill="#F6EEDB" stroke="#78350F" stroke-width="2.2"/>
  <rect x="10" y="10" width="540" height="22" rx="3" fill="#451A03"/>
  <text x="18" y="25" font-family="'IBM Plex Mono', monospace" font-size="10" font-weight="700" fill="#FDE68A">[CASE CARICATURE · MOLD %s · %s]</text>
  <text x="542" y="25" text-anchor="end" font-family="'IBM Plex Mono', monospace" font-size="9.5" fill="#FCD34D">DISPOSITION: %s</text>

  <!-- Left Dramatic Situation Prop Box -->
  <rect x="16" y="40" width="168" height="84" rx="5" fill="#FFFBEB" stroke="#92400E" stroke-width="1.8" stroke-dasharray="4,2"/>
  <text x="100" y="58" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="8.5" font-weight="700" fill="#92400E">THE PITCHED SCENE</text>
  <text x="100" y="82" text-anchor="middle" font-family="'Work Sans', sans-serif" font-size="10.5" font-weight="700" fill="#1E293B">%s</text>
  <line x1="36" y1="96" x2="164" y2="96" stroke="#D97706" stroke-width="1.2"/>
  <text x="100" y="112" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="8.5" fill="#78350F">%s</text>

  <!-- Center Dignified Silhouette Caricature (Person + Dramatic Situation) -->
  <g transform="translate(205, 36)">
    <rect x="0" y="0" width="150" height="90" rx="6" fill="#EFE5CE" stroke="#451A03" stroke-width="2"/>
    <polygon points="75,2 22,88 128,88" fill="#FDE68A" fill-opacity="0.35"/>
    <circle cx="75" cy="30" r="14" fill="#1E293B" stroke="#451A03" stroke-width="1.5"/>
    <rect x="65" y="26" width="8" height="5" rx="1" fill="none" stroke="#F8FAFC" stroke-width="1.2"/>
    <rect x="77" y="26" width="8" height="5" rx="1" fill="none" stroke="#F8FAFC" stroke-width="1.2"/>
    <line x1="73" y1="28" x2="77" y2="28" stroke="#F8FAFC" stroke-width="1.2"/>
    <path d="M 46 86 C 46 54, 104 54, 104 86 Z" fill="#1E293B"/>
    <polygon points="75,50 70,64 75,72 80,64" fill="#B45309"/>
    <line x1="28" y1="62" x2="54" y2="58" stroke="#1E293B" stroke-width="4" stroke-linecap="round"/>
    <line x1="96" y1="58" x2="122" y2="62" stroke="#1E293B" stroke-width="4" stroke-linecap="round"/>
    <line x1="18" y1="62" x2="132" y2="62" stroke="#991B1B" stroke-width="2" stroke-dasharray="3,2"/>
  </g>

  <!-- Right Structural Reality Trap Box -->
  <rect x="376" y="40" width="168" height="84" rx="5" fill="#FEF2F2" stroke="#991B1B" stroke-width="1.8"/>
  <text x="460" y="58" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="8.5" font-weight="700" fill="#991B1B">THE SYSTEM MECHANISM</text>
  <text x="460" y="82" text-anchor="middle" font-family="'Work Sans', sans-serif" font-size="10.5" font-weight="700" fill="#7F1D1D">%s</text>
  <line x1="396" y1="96" x2="524" y2="96" stroke="#DC2626" stroke-width="1.2"/>
  <text x="460" y="112" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="8.5" font-weight="700" fill="#991B1B">ZERO POSIBOT · CORONER FILE</text>

  <!-- Bottom Caricature Caption Band -->
  <text x="280" y="142" text-anchor="middle" font-family="'IBM Plex Mono', monospace" font-size="9.5" font-weight="700" fill="#451A03">%s</text>
</svg>`,
		escapeXML(moldCode),
		escapeXML(strings.ToUpper(personaName)),
		escapeXML(strings.ToUpper(disposition)),
		escapeXML(propLeft),
		escapeXML(sceneTitle),
		escapeXML(propRight),
		escapeXML(captionLine),
	)
	return []byte(svg)
}

func escapeXML(s string) string {
	r := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&apos;",
	)
	return r.Replace(s)
}
