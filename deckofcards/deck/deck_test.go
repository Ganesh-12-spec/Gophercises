package deck

import "testing"

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}
}