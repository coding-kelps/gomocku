package models

type Position struct {
	Y uint8
	X uint8
}

type Board struct {
	cells [][]uint8
	size uint8
}

type Player uint8

const (
	White Player = 0
	Black Player = 1
)

type Turn struct {
	Position Position
	Player Player
}

func NewBoard(size uint8) *Board {
	cells := make([][]uint8, size)
	for i := range cells {
		cells[i] = make([]uint8, size)
	}

	return &Board{
		cells: cells,
		size: size,
	}
}
