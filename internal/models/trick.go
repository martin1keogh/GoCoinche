package models

type Trick struct {
	bidSuit BidSuit
	opener  Player
	North   *Card
	East    *Card
	South   *Card
	West    *Card
}
