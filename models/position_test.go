package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPlayerOrderNoArg(t *testing.T) {
	r := GetPlayerOrder()
	expected := [4]Direction{North, East, South, West}

	for i := 0; i < 12; i++ {
		assert.Equal(t, expected[i%4], r.Value)
		r = *r.Next()
	}
}
