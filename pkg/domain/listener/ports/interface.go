package ports

import (
	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
	ai_models "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

type ManagerNotifier interface {
	NotifyMove(p ai_models.Position) error

	NotifyReadiness() error

	NotifyUnknown() error
	
	NotifyError(str string) error

	NotifyMessage(str string) error

	NotifyDebug(str string) error

	NotifySuggestion(p ai_models.Position) error
	
	NotifyMetadata(metadata map[string]string) error
}

type ManagerListener interface {
	Listen(ch chan<-models.ManagerCommand) error
}

type ManagerInterface interface {
	ManagerListener
	ManagerNotifier
}
