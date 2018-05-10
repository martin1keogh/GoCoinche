package models

import (
	"errors"
)

type Board struct {
	bids   []Bid
	tricks [8]Trick
}

func (board *Board) WinningBid() (*Bid, error) {
	nbBids := len(board.bids)
	if nbBids == 0 {
		return nil, errors.New("No winning bid when nobody bid yet.")
	}

	winning := board.bids[0]
	for _, bid := range board.bids {
		if bid.value > winning.value {
			winning = bid
		}
	}
	return &winning, nil
}
