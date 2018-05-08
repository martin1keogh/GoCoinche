package models

type Suit string

const (
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
	Spades   Suit = "Spades"
)

func AllSuits() [4]Suit {
	return [...]Suit{Hearts, Diamonds, Clubs, Spades}
}

type CardName string

const (
	_7  CardName = "Seven"
	_8  CardName = "Eight"
	_9  CardName = "Nine"
	_10 CardName = "Ten"
	_J  CardName = "Jack"
	_Q  CardName = "Queen"
	_K  CardName = "King"
	_As CardName = "As"
)

func AllCardNames() [8]CardName {
	return [8]CardName{_7, _8, _9, _10, _J, _Q, _K, _As}
}

type Card struct {
	suit Suit
	name CardName
}

func (c *Card) IsTrumpAt(bs BidSuit) bool {
	switch bs {
	case Heart:
		return c.suit == Hearts
	case Diamond:
		return c.suit == Diamonds
	case Club:
		return c.suit == Clubs
	case Spade:
		return c.suit == Spades
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
