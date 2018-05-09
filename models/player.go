package models

type Player struct {
	hand     *[]Card
	table    Table
	position Position
}
