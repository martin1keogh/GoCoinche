package models

type cardArray interface {
	Cards() []Card
}

func Contains(arr cardArray, card Card) bool {
	for i := range arr.Cards() {
		if arr.Cards()[i] == card {
			return true
		}
	}

	return false
}

func PointsAt(arr cardArray, bs BidSuit) int {
	total := 0
	switch bs {
	case AllTrump:
		for _, card := range arr.Cards() {
			total += whenAllTrumpsCardValue[card.name]
		}

	case NoTrump:
		for _, card := range arr.Cards() {
			total += whenNoTrumpsCardValue[card.name]
		}

	default:
		for _, card := range arr.Cards() {
			if card.IsTrumpAt(bs) {
				total += whenIsTrumpCardValue[card.name]
			} else {
				total += whenIsNotTrumpCardValue[card.name]
			}
		}

	}

	return total
}
