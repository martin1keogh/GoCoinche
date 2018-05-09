package models

type Table struct {
	NorthSouth   Team
	EastWest     Team
	currentBoard *Board
}

func New() Table {
	var NS Team = Team{
		position1: North,
		position2: South,
		score:     0,
	}
	var EW Team = Team{
		position1: East,
		position2: West,
		score:     0,
	}

	return Table{
		NorthSouth:   NS,
		EastWest:     EW,
		currentBoard: nil,
	}
}
