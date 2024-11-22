package models

import (
	"fmt"
)

type Position struct {
	Y uint8
	X uint8
}

type Player uint8

const (
	Us Player = 0
	Opponent Player = 1
)

type InvalidPlayerError struct {
	PlayerValue Player
}

func (e *InvalidPlayerError) Error() string {
	return fmt.Sprintf(
		"\"%d\" value doesn't correspond to a valid player",
		e.PlayerValue,
	)
}

type Turn struct {
	Position Position
	Player   Player
}

type Board struct {
	cells [][]CellStatus
	size  uint8
}

type CellStatus uint8

const (
	Available  CellStatus = 0
	OwnStone CellStatus = 1
	OpponentStone CellStatus = 2
)

func NewBoard(size uint8) *Board {
	cells := make([][]CellStatus, size)
	for x := range cells {
		cells[x] = make([]CellStatus, size)
		for y := range cells[x] {
			cells[x][y] = Available
		}
	}

	return &Board{
		cells: cells,
		size:  size,
	}
}

type OutOfBoundError struct {
	position  Position
	boardSize uint8
}

func (e *OutOfBoundError) Error() string {
	return fmt.Sprintf(
		"move (x: %d, y: %d) is out of bound of board (size: %dx%d)",
		e.position.X,
		e.position.Y,
		e.boardSize,
		e.boardSize,
	)
}

type CellUnavailableError struct {
	position Position
}

func (e *CellUnavailableError) Error() string {
	return fmt.Sprintf(
		"cell (x: %d, y: %d) is not available",
		e.position.X,
		e.position.Y,
	)
}

func (b *Board) SetCell(position Position, status CellStatus) error {
	if position.X > b.size || position.Y > b.size {
		return &OutOfBoundError{position, b.size}
	}

	if b.cells[position.X][position.Y] != Available {
		return &CellUnavailableError{position}
	}

	b.cells[position.X][position.Y] = status

	return nil
}



type NoCellAvailableError struct {
}

func (e *NoCellAvailableError) Error() string {
	return "board full, no cell available"
}

func (b *Board) GetAvailableCells() ([]Position, error) {
	availables := []Position{}

	for x := range b.size {
		for y := range b.size {
			if b.cells[x][y] == Available {
				availables = append(availables, Position{x, y})
			}
		}
	}

	if len(availables) == 0 {
		return []Position{}, &NoCellAvailableError{}
	}

	return availables, nil
}

type BoardUnsetError struct {
}

func (e *BoardUnsetError) Error() string {
	return "game hasn't started"
}