package ports

import (
	"github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

type AI interface {
	Init(size uint8) error

	Reset() error

	RegisterMove(pos models.Position, player models.CellStatus) error

	PickMove() (models.Position, error)
}
