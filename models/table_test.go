package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTableSetTeams(t *testing.T) {
	table := NewTable()

	assert.Equal(t, table.NorthSouth.position1.direction, North)
	assert.Equal(t, table.NorthSouth.position2.direction, South)
	assert.Equal(t, table.NorthSouth.score, Score(0))

	assert.Equal(t, table.EastWest.position1.direction, East)
	assert.Equal(t, table.EastWest.position2.direction, West)
	assert.Equal(t, table.EastWest.score, Score(0))
}

func TestDealDoesntShuffle(t *testing.T) {
	table := NewTable()
	table.Deal()
	h1 := table.deck
	table.Deal()
	h2 := table.deck

	assert.Equal(t, h1, h2)
}

func TestDealDealsToAll(t *testing.T) {
	table := NewTable()
	table.Deal()
	var seen []Direction

	table.dealer.Next().Do(func(d interface{}) {
		direction := d.(Direction)
		position := table.findPosition(direction)
		assert.NotEmpty(t, position.hand)
		assert.NotContains(t, position.direction, seen)
		seen = append(seen, position.direction)
	})
}
