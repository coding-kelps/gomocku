package models

type Position struct {
	Y uint8
	X uint8
}

type Board struct {
	cells [][]uint8
	size uint8
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
