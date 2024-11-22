package service

import (
	"math/rand"
	"time"

	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
	"github.com/coding-kelps/gomocku/pkg/domain/mock/ports"
)

type Mock struct {
	board *models.Board
	about map[string]string
	rng   *rand.Rand

	ports.Mock
}

func NewMock() ports.Mock {
	return NewMockWithSeed(time.Now().UnixNano())
}

func NewMockWithSeed(seed int64) ports.Mock {
	source := rand.NewSource(seed)

	return &Mock{
		board: nil,
		about: map[string]string{
			"name":    "gomocku",
			"version": "0.1",
			"author":  "Coding Kelps",
			"desc":    "A mock AI for manager testing",
		},
		rng: rand.New(source),
	}
}
