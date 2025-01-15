package service

import (
	"github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func (m *AI) Init(size uint8) error {
	m.board = models.NewBoard(size)

	return nil
}

func (m *AI) RegisterMove(pos models.Position, player models.CellStatus) error {
	if m.board == nil {
		return &models.BoardUnsetError{}
	}

	return m.board.SetCell(pos, models.OpponentStone)
}

func (m *AI) PickMove() (models.Position, error) {
	if m.board == nil {
		return models.Position{}, &models.BoardUnsetError{}
	}
	
	move, err := m.pickRandomMove()
	if err != nil {
		return models.Position{}, err
	}

	err = m.board.SetCell(move, models.OwnStone)
	if err != nil {
		return models.Position{}, err
	}

	return move, nil
}
