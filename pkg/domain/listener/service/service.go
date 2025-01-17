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

func (l *Listener) Listen() error {
	ch := make(chan models.ManagerCommand, 10)

	go l.managerInterface.Listen(ch)

	for cmd := range ch {
		switch cmd.CommandType() {
		case "start":
			l.startHandler(cmd.(models.StartCommand))
		case "turn":
			l.turnHandler(cmd.(models.TurnCommand))
		case "begin":
			l.beginHandler()
		case "board":
			l.boardHandler()
		case "board_turn":
			l.boardTurnHandler(cmd.(models.BoardTurnCommand))
		case "board_done":
			l.boardDoneHandler()
		case "end":
			return nil
		case "info":
			l.infoHandler(cmd.(models.InfoCommand))
		case "about":
			l.aboutHandler()
		case "unknown":
			l.UnknownHandler()
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
