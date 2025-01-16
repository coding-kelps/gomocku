package service

import (
	"math/rand"
	"time"

	"github.com/coding-kelps/gomocku/pkg/domain/ai/models"
	"github.com/coding-kelps/gomocku/pkg/domain/ai/ports"
)

type AI struct {
	board *models.Board
	rng   *rand.Rand

	ports.AI
}

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

	return move, nil
}

func NewRandomAI() ports.AI {
	return NewRandomAIWithSeed(time.Now().UnixNano())
}

func NewRandomAIWithSeed(seed int64) ports.AI {
	source := rand.NewSource(seed)

	return &AI{
		board: nil,
		rng: rand.New(source),
	}
}
