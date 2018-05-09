package models

type BidSuit string

const (
	Heart    BidSuit = "Heart"
	Diamond  BidSuit = "Diamond"
	Club     BidSuit = "Club"
	Spade    BidSuit = "Spade"
	NoTrump  BidSuit = "NoTrump"
	AllTrump BidSuit = "AllTrump"
)

type BidValue int

const (
	_80  BidValue = 80
	_90  BidValue = 90
	_100 BidValue = 100
	_110 BidValue = 110
	_120 BidValue = 120
	_130 BidValue = 130
	_140 BidValue = 140
	_150 BidValue = 150
	_160 BidValue = 160
	_250 BidValue = 250
	_400 BidValue = 400
)

type Bid struct {
	suit  BidSuit
	value BidValue
}
