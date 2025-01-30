package service

import (
	"math/rand"
	"time"

    "github.com/rs/zerolog"

	"github.com/coding-kelps/gomocku/pkg/domain/ai/models"
	"github.com/coding-kelps/gomocku/pkg/domain/ai/ports"
)

type AI struct {
	board 				*models.Board
	rng   				*rand.Rand
	logger              zerolog.Logger

	ports.AI
}

func (a *AI) Init(size uint8) error {
	a.board = models.NewBoard(size)
	a.logger.Debug().
		Uint8("size", size).
		Msg("init board")

	return nil
}

func (a *AI) RegisterMove(pos models.Position, player models.CellStatus) error {
	if a.board == nil {
		return &models.BoardUnsetError{}
	}

	err := a.board.SetCell(pos, models.OpponentStone)
	if err != nil {
		return err
	}

	a.logger.Debug().
		Uint8("turn_x", pos.X).
		Uint8("turn_y", pos.Y).
		Str("", "").
		Msg("registered move")

	return nil
}

func (a *AI) PickMove() (models.Position, error) {
	if a.board == nil {
		return models.Position{}, &models.BoardUnsetError{}
	}
	
	move, err := a.pickRandomMove()
	if err != nil {
		return models.Position{}, err
	}

	a.logger.Debug().
		Uint8("move_x", move.X).
		Uint8("move_y", move.Y).
		Msg("pick move")

	return move, nil
}

func NewRandomAI(logger zerolog.Logger) ports.AI {
	return NewRandomAIWithSeed(logger, time.Now().UnixNano())
}

func NewRandomAIWithSeed(logger zerolog.Logger, seed int64) ports.AI {
	source := rand.NewSource(seed)

	return &AI{
		board:	nil,
		rng:	rand.New(source),
		logger:	logger,
	}
}
