package mock

import (
	"github.com/coding-kelps/gomocku/pkg/domain/mock/ports"
	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

type Mock struct {
	board *models.Board
	about map[string]string

	ports.Mock
}

func NewMock() ports.Mock {
	return &Mock{
		board: nil,
		about: map[string]string{
			"name": "gomocku",
			"version": "0.1",
			"author": "Coding Kelps",
		},
	}
}

func (m *Mock) RespondStart(size uint8) error {
	m.board = models.NewBoard(size)

	return nil
}

func (m *Mock) RespondTurn(p models.Position) (models.Position, error) {
	// Set corresponding cell in board

	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	return move, nil
}

func (m *Mock) RespondBegin() (models.Position, error) {
	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	return move, nil
}

func (m *Mock) RespondBoard(p []models.Position) (models.Position, error) {
	// Set corresponding cells in board

	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	return move, nil
}

func (m *Mock) RespondInfo() error {
	return nil
}

func (m *Mock) RespondEnd() error {
	return nil
}

func (m *Mock) RespondAbout() (map[string]string, error) {
	return m.about, nil
}
