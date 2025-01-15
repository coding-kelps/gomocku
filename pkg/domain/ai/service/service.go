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

func NewMock() ports.AI {
	return NewMockWithSeed(time.Now().UnixNano())
}

func NewMockWithSeed(seed int64) ports.AI {
	source := rand.NewSource(seed)

	return &AI{
		board: nil,
		rng: rand.New(source),
	}
}
