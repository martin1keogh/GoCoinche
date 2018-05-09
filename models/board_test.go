package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWinningBidErrorWhenEmpty(t *testing.T) {
	board := Board{
		bids:   []Bid{},
		tricks: [8]Trick{},
	}
	bid, err := board.WinningBid()

	assert.Nil(t, bid)
	assert.Error(t, err)
}

func TestWinningBid(t *testing.T) {
	bid1 := Bid{Heart, 90}
	bid2 := Bid{Heart, 130}
	bid3 := Bid{Spade, 120}

	board := Board{[]Bid{bid1, bid2, bid3}, [8]Trick{}}

	bid, err := board.WinningBid()

	assert.Nil(t, err)
	assert.Equal(t, bid2, *bid)
}
