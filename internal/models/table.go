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

	dealer := GetPlayerOrder(North)

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

	// We first deal cards to the player after the current dealer,
	// so we have to call Next() one more time.
	firstReceiver := t.dealer.Next().Value.(Direction)

	t.ForeachPositionStartingFrom(firstReceiver, func(pos *Position) {
		pos.Hand = newHands[i]
		i += 1
	})
}

func (t *Table) findPosition(d Direction) *Position {
	switch d {
	case North:
		return &t.NorthSouth.position1
	case South:
		return &t.NorthSouth.position2
	case East:
		return &t.EastWest.position1
	default: // O real enums how I miss thee
		return &t.EastWest.position2
	}
}

func (t *Table) ForeachPositionStartingFrom(start Direction, action func(*Position)) {
	ring := GetPlayerOrder(start)
	ring.Next().Do(func(i interface{}) {
		dir := i.(Direction)
		pos := t.findPosition(dir)
		action(pos)
	})
}
