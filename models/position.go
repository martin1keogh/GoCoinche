package models

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
