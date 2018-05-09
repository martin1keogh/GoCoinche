package models

import (
	"container/ring"
)

type Table struct {
	NorthSouth Team
	EastWest   Team
	Board      *Board

	deck   Deck
	dealer ring.Ring
}

func NewTable() Table {
	var NS Team = Team{
		position1: Position{North, [8]Card{}},
		position2: Position{South, [8]Card{}},
		score:     0,
	}
	var EW Team = Team{
		position1: Position{East, [8]Card{}},
		position2: Position{West, [8]Card{}},
		score:     0,
	}

	deck := SortedDeck()
	shuffled := deck.Shuffle()

	dealer := GetPlayerOrder()

	table := Table{
		NorthSouth: NS,
		EastWest:   EW,
		Board:      nil,

		deck:   shuffled,
		dealer: dealer,
	}

	table.Deal()
	return table
}

func (t *Table) Deal() {
	newHands := t.deck.deal()
	i := 0

	// Pass the dealer chip to the next player
	// TODO see if it's not better to call this at the end of
	// a board instead of here
	t.dealer = *t.dealer.Next()

	// We first deal cards to the player after the current dealer,
	// so we have to call Next() one more time.
	t.dealer.Next().Do(func(d interface{}) {
		direction := d.(Direction)
		position := t.findPosition(direction)
		position.hand = newHands[i]
		i += 1
	})
}

func (t *Table) findPosition(d Direction) Position {
	switch d {
	case North:
		return t.NorthSouth.position1
	case South:
		return t.NorthSouth.position2
	case East:
		return t.EastWest.position1
	default: // O real enums how I miss thee
		return t.EastWest.position2
	}
}
