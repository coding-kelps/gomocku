package ports

import (
	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

type ResponseSender interface {
	SendMove(p models.Position) error

	SendUnknown(msg string) error

	SendError(msg string) error

	SendDebug(msg string) error
}
