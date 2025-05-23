package service

import (
	"sync"

	aiPorts "github.com/coding-kelps/gomocku/pkg/domain/ai/ports"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/ports"
)

type Coordinator struct {
	managerInterface 	ports.ManagerInterface
	metadata 			map[string]string
	ai 					aiPorts.AI
	lock				*sync.RWMutex

	ports.Coordinator
}

func (c *Coordinator) Serve() error {
	actionsCh := make(chan models.ManagerAction, 10)
	errorsCh := make(chan error)

	go c.managerInterface.Listen(actionsCh, errorsCh)

	for {
        select {
        case cmd, ok := <-actionsCh:
            if !ok {
                return nil
            }

            switch cmd.ActionType() {
            case "start":
                c.startHandler(cmd.(models.StartAction))
            case "turn":
                c.turnHandler(cmd.(models.TurnAction))
            case "begin":
                c.beginHandler()
            case "board":
                c.boardHandler()
            case "board_turn":
                c.boardTurnHandler(cmd.(models.BoardTurnAction))
            case "board_done":
                c.boardDoneHandler()
            case "end":
                return nil
            case "info":
                c.infoHandler(cmd.(models.InfoAction))
            case "about":
                c.aboutHandler()
            case "unknown":
                c.UnknownHandler()
            }

        case err, ok := <-errorsCh:
            if !ok {
                return nil
            }
			
            return err
        }
    }
}

func NewCoordinator(managerInterface ports.ManagerInterface, ai aiPorts.AI) ports.Coordinator {
	return &Coordinator{
		managerInterface: managerInterface,
		ai: ai,
		lock: &sync.RWMutex{},
		metadata: map[string]string{
			"name":    "gomocku",
			"version": "0.1",
			"author":  "Coding Kelps",
			"desc":    "A mock AI for manager testing",
		},
	}
}
