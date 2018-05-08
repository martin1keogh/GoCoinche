package models

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
