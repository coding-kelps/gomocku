package mock

import (
	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
	"github.com/coding-kelps/gomocku/pkg/domain/mock/ports"
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
			"name":    "gomocku",
			"version": "0.1",
			"author":  "Coding Kelps",
			"desc":    "A mock AI for manager testing",
		},
	}
}

func (m *Mock) RespondStart(size uint8) error {
	m.board = models.NewBoard(size)

	return nil
}

func (m *Mock) RespondTurn(p models.Position) (models.Position, error) {
	if m.board == nil {
		return models.Position{}, &models.BoardUnsetError{}
	}

	err := m.board.SetCell(p, models.OpponentStone)
	if err != nil {
		return models.Position{}, err
	}

	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	return move, nil
}

func (m *Mock) RespondBegin() (models.Position, error) {
	if m.board == nil {
		return models.Position{}, &models.BoardUnsetError{}
	}

	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	return move, nil
}

func playerToCellStatus(p models.Player) (models.CellStatus, error) {
	if p == models.Us {
		return models.OwnStone, nil
	} else if p == models.Opponent {
		return models.OpponentStone, nil
	} else {
		return 0, &models.InvalidPlayerError{PlayerValue: p}
	}
}

func (m *Mock) RespondBoard(turns []models.Turn) (models.Position, error) {
	if m.board == nil {
		return models.Position{}, &models.BoardUnsetError{}
	}
	// Set corresponding cells in board
	for _, t := range turns {
		status, err := playerToCellStatus(t.Player)
		if err != nil {
			return models.Position{}, err
		}

		err = m.board.SetCell(t.Position, status)
		if err != nil {
			return models.Position{}, err
		}
	}

	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	return move, nil
}

func (m *Mock) RespondInfo() error {
	return nil
}

func (m *Mock) RespondEnd() error {
	if m.board == nil {
		return &models.BoardUnsetError{}
	}

	return nil
}

func (m *Mock) RespondAbout() (map[string]string, error) {
	return m.about, nil
}
