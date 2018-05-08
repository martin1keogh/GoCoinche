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
