package models

import (
	"math/rand"
)

type Deck struct {
	cards [32]Card
}

func (d *Deck) Cards() []Card { return d.cards[:] }

func SortedDeck() Deck {
	var deck Deck
	for i, suit := range AllSuits() {
		for j, card := range AllCardNames() {
			deck.cards[i*8+j] = Card{suit, card}
		}
	}

	return deck
}

func (d *Deck) Shuffle() Deck {
	var newDeck Deck
	for i, j := range rand.Perm(32) {
		newDeck.cards[i] = d.cards[j]
	}

	return newDeck
}
