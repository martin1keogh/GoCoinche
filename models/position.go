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
	direction Direction
	hand      [8]Card
}

func GetPlayerOrder() ring.Ring {
	r := &ring.Ring{Value: North}
	r.Link(&ring.Ring{Value: West})
	r.Link(&ring.Ring{Value: South})
	r.Link(&ring.Ring{Value: East})

	return *r
}
