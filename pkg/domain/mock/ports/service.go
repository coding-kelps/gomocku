package ports

import (
	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

type Mock interface {
	RespondStart(size uint8) error

	RespondTurn(p models.Position) error

	RespondBegin() error

	RespondBoard(p []models.Position) error

	RespondInfo() error

	RespondEnd() error

	RespondAbout() error
}
