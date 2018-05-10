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
	for _, card := range arr.Cards() {
		total += card.PointsAt(bs)
	}
	return total
}
