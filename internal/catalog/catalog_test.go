package catalog

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFormatAssetIDAndWriteCatalog(t *testing.T) {
	id := FormatAssetID(1, 2, 4, "exhibit", "E2 Gap Bar", 1)
	if id != "T01-P2-C04-exhibit-e2-gap-bar_v001" {
		t.Fatalf("unexpected asset ID: %s", id)
	}

	tmpDir := t.TempDir()
	svg := RenderPosiBotSVG("KEEP THE DAY JOB!", "(Audited median: -$1,840)", "signpost", "#0D9488")
	rec, err := WriteCatalogAsset(tmpDir, 1, 3, 7, "posibot", "keep-day-job", 1, "svg", svg, "PosiBot holding signpost", "signpost", "", []string{"F-019"})
	if err != nil {
		t.Fatalf("WriteCatalogAsset failed: %v", err)
	}
	if rec.SHA256 == "" || rec.PromptSHA256 == "" {
		t.Fatalf("expected non-empty SHA256 hashes")
	}
	if _, err := os.Stat(filepath.Join(tmpDir, rec.RelativePath)); err != nil {
		t.Fatalf("missing asset file: %v", err)
	}
	if _, err := os.Stat(filepath.Join(tmpDir, rec.PromptPath)); err != nil {
		t.Fatalf("missing prompt sidecar: %v", err)
	}
}
