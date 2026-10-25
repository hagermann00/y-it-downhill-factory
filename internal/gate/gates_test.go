package gate

import (
	"testing"
)

func TestCanonical38QuestionsCount(t *testing.T) {
	qs := Canonical38Questions()
	if len(qs) != 38 {
		t.Fatalf("expected 38 canonical questions, got %d", len(qs))
	}
	if qs[0] != "P1-Q1" || qs[11] != "P1-Q12" || qs[12] != "P2-Q1" || qs[23] != "P2-Q12" || qs[24] != "P3-Q1" || qs[37] != "P3-Q14" {
		t.Fatalf("unexpected question sequence boundaries: %v", qs)
	}
}

func TestNormalizeText(t *testing.T) {
	a := "Between  74% and 89%\n of retail  investor accounts lose money."
	b := "between 74% and 89% of retail investor accounts lose money."
	if normalizeText(a) != normalizeText(b) {
		t.Fatalf("expected normalized strings to match")
	}
}
