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

func (d *Deck) deal() [4][8]Card {
	c := d.cards
	newHands := [4][8]Card{}

	for i := 0; i < 4; i++ {
		hand := [8]Card{}
		hand[0] = c[i*3]
		hand[1] = c[i*3+1]
		hand[2] = c[i*3+2]
		hand[3] = c[i*3+12]
		hand[4] = c[i*3+13]
		hand[5] = c[i*3+20]
		hand[6] = c[i*3+21]
		hand[7] = c[i*3+22]
		newHands[i] = hand
	}

	return newHands
}
