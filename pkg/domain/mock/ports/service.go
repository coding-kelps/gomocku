package ports

import (
	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

type Mock interface {
	RespondStart(size uint8) error

	RespondTurn(p models.Position) (models.Position, error)

	RespondBegin() (models.Position, error)

	RespondBoard(p []models.Turn) (models.Position, error)

	RespondInfo() error

	RespondEnd() error

	RespondAbout() (map[string]string, error)
}
