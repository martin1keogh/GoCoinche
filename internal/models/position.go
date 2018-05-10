package models

import (
	"container/ring"
)

type Direction string

const (
	North Direction = "North"
	East  Direction = "East"
	South Direction = "South"
	West  Direction = "West"
)

type Position struct {
	Direction Direction
	Hand      [8]Card
}

func GetPlayerOrder(startingDirection Direction) ring.Ring {
	r := &ring.Ring{Value: North}
	r.Link(&ring.Ring{Value: West})
	r.Link(&ring.Ring{Value: South})
	r.Link(&ring.Ring{Value: East})

	var shiftBy int
	switch startingDirection {
	case East:
		shiftBy = 1
	case South:
		shiftBy = 2
	case West:
		shiftBy = 3
	default:
		shiftBy = 0
	}

	for i := 0; i < shiftBy; i++ {
		r = r.Next()
	}

	return *r
}
