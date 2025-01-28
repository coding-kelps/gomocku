package ports

import (
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
	aiModels "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

type ManagerNotifier interface {
	NotifyMove(p aiModels.Position) error

	NotifyReadiness() error

	NotifyUnknown() error
	
	NotifyError(str string) error

	NotifyMessage(str string) error

	NotifyDebug(str string) error

	NotifySuggestion(p aiModels.Position) error
	
	NotifyMetadata(metadata map[string]string) error
}

type ManagerListener interface {
	Listen(ch chan<-models.ManagerAction) error
}

type ManagerInterface interface {
	ManagerListener
	ManagerNotifier
}
