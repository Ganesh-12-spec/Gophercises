package deck

import "testing"

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))

	if len(cards) != 55 {
		t.Errorf("Expected 55 cards, got %d", len(cards))
	}
}

func TestMultipleDecks(t *testing.T) {
	cards := New(MultipleDecks(3))

	if len(cards) != 156 {
		t.Errorf("Expected 156 cards, got %d", len(cards))
	}
}

func TestFilter(t *testing.T) {
	cards := New(Filter(func(c Card) bool {
		return c.Rank != Ace
	}))

	if len(cards) != 48 {
		t.Errorf("Expected 48 cards, got %d", len(cards))
	}
}