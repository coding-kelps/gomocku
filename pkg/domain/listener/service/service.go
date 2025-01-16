package service

import (
	"sync"

	ai_ports "github.com/coding-kelps/gomocku/pkg/domain/ai/ports"
	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
	"github.com/coding-kelps/gomocku/pkg/domain/listener/ports"
)

type Listener struct {
	managerInterface 	ports.ManagerInterface
	metadata 			map[string]string
	ai 					ai_ports.AI
	lock				*sync.RWMutex

	ports.Listener
}

func (l Listener) Listen() error {
	ch := make(chan models.ManagerCommand, 10)

	go l.managerInterface.Listen(ch)
	defer close(ch)

	for cmd := range ch {
		switch cmd.CommandType() {
		case "start":
			go l.startHandler(cmd.(models.StartCommand))
		case "turn":
			go l.turnHandler(cmd.(models.TurnCommand))
		case "begin":
			go l.beginHandler()
		case "board":
			go l.boardHandler()
		case "board_turn":
			go l.boardTurnHandler(cmd.(models.BoardTurnCommand))
		case "board_done":
			go l.boardDoneHandler()
		case "end":
			return nil
		case "info":
			go l.infoHandler(cmd.(models.InfoCommand))
		case "about":
			go l.aboutHandler()
		case "unknown":
			go l.UnknownHandler()
		}
	}

	return nil
}

func NewListener(managerInterface ports.ManagerInterface, ai ai_ports.AI) ports.Listener {
	return &Listener{
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
