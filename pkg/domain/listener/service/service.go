package service

import (
	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
	"github.com/coding-kelps/gomocku/pkg/domain/listener/ports"
	ai_models "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
	ai_ports "github.com/coding-kelps/gomocku/pkg/domain/ai/ports"
)

type Listener struct {
	managerInterface ports.ManagerInterface
	ai ai_ports.AI
	metadata map[string]string

	ports.Listener
}

func (l Listener) Listen() error {
	ch := make(chan models.ManagerCommand, 10)

	go l.managerInterface.Listen(ch)
	defer close(ch)

	for cmd := range ch {
		switch c := cmd.(type) {
		case models.StartCommand:
			err := l.ai.Init(c.Size)
			if err != nil {
				err2 := l.managerInterface.NotifyError(err.Error())
				if err2 != nil {
					return err2
				}
				continue
			}

			l.managerInterface.NotifyReadiness()
			continue
		case models.TurnCommand:
			err := l.ai.RegisterMove(c.Position, ai_models.OpponentStone)
			if err != nil {
				err2 := l.managerInterface.NotifyError(err.Error())
				if err2 != nil {
					return err2
				}
				continue
			}

			pos, err := l.ai.PickMove()
			if err != nil {
				err2 := l.managerInterface.NotifyError(err.Error())
				if err2 != nil {
					return err2
				}
				continue
			}

			err = l.managerInterface.NotifyMove(pos)
			if err != nil {
				return err
			}

			continue
		case models.BeginCommand:
			pos, err := l.ai.PickMove()
			if err != nil {
				err2 := l.managerInterface.NotifyError(err.Error())
				if err2 != nil {
					return err2
				}
				continue
			}

			err = l.managerInterface.NotifyMove(pos)
			if err != nil {
				return err
			}
			
			continue
		case models.BoardCommand:
			continue
		case models.BoardTurnCommand:
			continue
		case models.BoardDoneCommand:
			err := l.managerInterface.NotifyMove(ai_models.Position{X: 10, Y: 10})
			if err != nil {
				return err
			}

			continue
		case models.InfoCommand:
			continue
		case models.EndCommand:
			return nil
		case models.AboutCommand:
			err := l.managerInterface.NotifyMetadata(l.metadata)
			if err != nil {
				return err
			}

			continue
		default:
			err := l.managerInterface.NotifyUnknown()
			if err != nil {
				return err
			}

			continue
		}
	}

	return nil
} 

func NewListener(managerInterface ports.ManagerInterface, ai ai_ports.AI) ports.Listener {
	return &Listener{
		managerInterface: managerInterface,
		ai: ai,
		metadata: map[string]string{
			"name":    "gomocku",
			"version": "0.1",
			"author":  "Coding Kelps",
			"desc":    "A mock AI for manager testing",
		},
	}
}
