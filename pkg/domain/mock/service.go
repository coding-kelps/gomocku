package mock

import (
	"github.com/coding-kelps/gomocku/pkg/domain/mock/ports"
	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

type Mock struct {
	sender ports.ResponseSender
	board *models.Board

	ports.Mock
}

func NewMock(s ports.ResponseSender) ports.Mock {
	return &Mock{
		sender: s,
		board: nil,
	}
}

func (m *Mock) RespondStart(size uint8) error {
	m.board = models.NewBoard(size)

	return nil
}

func (m *Mock) RespondTurn(p models.Position) error {
	// Set corresponding cell in board

	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	m.sender.SendMove(move)

	return nil
}

func (m *Mock) RespondBegin() error {
	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	m.sender.SendMove(move)

	return nil
}

func (m *Mock) RespondBoard(p []models.Position) error {
	// Set corresponding cells in board

	// Arbitrary Choose cell in board as move
	move := models.Position{X: 0, Y: 0}

	m.sender.SendMove(move)

	return nil
}

func (m *Mock) RespondInfo() error {
	return nil
}

func (m *Mock) RespondEnd() error {
	return nil
}

func (m *Mock) RespondAbout() error {
	return nil
}
