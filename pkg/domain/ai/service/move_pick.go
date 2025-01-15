package service

import (
	"github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func (m *AI) pickRandomMove() (models.Position, error) {
	availables, err := m.board.GetAvailableCells()
	if err != nil {
		return models.Position{}, err
	}

	return availables[m.rng.Intn(len(availables))], nil
}
