package models

import "fmt"

type Suit string

const (
	Hearts   Suit = "♥"
	Diamonds Suit = "♦"
	Clubs    Suit = "♣"
	Spades   Suit = "♠"
)

func AllSuits() [4]Suit {
	return [...]Suit{Hearts, Diamonds, Clubs, Spades}
}

type CardName string

const (
	_7  CardName = "7"
	_8  CardName = "8"
	_9  CardName = "9"
	_10 CardName = "10"
	_J  CardName = "J"
	_Q  CardName = "Q"
	_K  CardName = "K"
	_As CardName = "As"
)

func AllCardNames() [8]CardName {
	return [8]CardName{_7, _8, _9, _10, _J, _Q, _K, _As}
}

type Card struct {
	suit Suit
	name CardName
}

func (card Card) String() string {
	return fmt.Sprintf("%v%v", card.name, card.suit)
}

func (card *Card) IsTrumpAt(bs BidSuit) bool {
	switch bs {
	case Heart:
		return card.suit == Hearts
	case Diamond:
		return card.suit == Diamonds
	case Club:
		return card.suit == Clubs
	case Spade:
		return card.suit == Spades
	case AllTrump:
		return true
	case NoTrump:
		return false
	}

	return false
}

func (card *Card) PointsAt(bs BidSuit) int {
	res := 0
	switch bs {
	case AllTrump:
		res = whenAllTrumpsCardValue[card.name]

	case NoTrump:
		res = whenNoTrumpsCardValue[card.name]

	default:
		if card.IsTrumpAt(bs) {
			res = whenIsTrumpCardValue[card.name]
		} else {
			res = whenIsNotTrumpCardValue[card.name]
		}

	}

	return res
}

type cardToPoint map[CardName]int

var whenIsTrumpCardValue = cardToPoint{
	_J:  20,
	_9:  14,
	_As: 11,
	_10: 10,
	_K:  4,
	_Q:  3}

var whenIsNotTrumpCardValue = cardToPoint{
	_As: 11,
	_10: 10,
	_K:  4,
	_Q:  3,
	_J:  2}

var whenAllTrumpsCardValue = cardToPoint{
	_J:  14,
	_9:  9,
	_As: 6,
	_10: 4,
	_K:  3,
	_Q:  2}

var whenNoTrumpsCardValue = cardToPoint{
	_As: 19,
	_10: 10,
	_K:  4,
	_Q:  3,
	_J:  2}
