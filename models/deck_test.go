package models

import (
	"testing"
)

func TestFullDeckHas32Cards(t *testing.T) {
	if len(SortedDeck().cards) != 32 {
		t.Errorf("Unexpected number of cards in SortedDeck(). Expected 32, got %d", len(SortedDeck().cards))
	}
}

func TestFullDeckOnlyHasDistinctCards(t *testing.T) {
	viewed := map[Card]bool{}
	for _, card := range SortedDeck().cards {
		if viewed[card] {
			t.Errorf("Duplicate card found in SortedDeck(): %+v", card)
		}

		viewed[card] = true
	}
}

func TestShuffleDeck(t *testing.T) {
	deck := SortedDeck()
	newDeck := deck.Shuffle()

	if deck == newDeck {
		t.Errorf("Got two identical decks after shuffling: \n%+v\n%+v", deck.cards, newDeck.cards)
	}

	// Check that reshuffling `deck` does not yield `newDeck` twice
	secondShuffle := deck.Shuffle()

	if newDeck == secondShuffle {
		t.Errorf("Shuffling the same deck twice should not result in the exact same order.")
	}
}
