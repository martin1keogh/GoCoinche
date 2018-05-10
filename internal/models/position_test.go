package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expected [4]Direction = [4]Direction{North, East, South, West}

func TestGetPlayerOrder(t *testing.T) {
	for shiftBy, direction := range expected {
		r := GetPlayerOrder(direction)

		for i := shiftBy; i < (12 + shiftBy); i++ {
			assert.Equal(t, expected[i%4], r.Value)
			r = *r.Next()
		}
	}
}
