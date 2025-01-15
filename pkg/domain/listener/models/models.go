package models

import (
	"github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

type ManagerCommand interface {
	CommandType() string
}

type StartCommand struct {
	Size uint8
}

func (c StartCommand) CommandType() string {
	return "start"
}

type TurnCommand struct {
	Position models.Position
}

func (c TurnCommand) CommandType() string {
	return "turn"
}

type BeginCommand struct {
}

func (c BeginCommand) CommandType() string {
	return "begin"
}

type BoardCommand struct {
}


func (c BoardCommand) CommandType() string {
	return "board"
}

type BoardTurnCommand struct {
	Turn models.Turn
}

func (c BoardTurnCommand) CommandType() string {
	return "board_turn"
}

type BoardDoneCommand struct {
}

func (c BoardDoneCommand) CommandType() string {
	return "board_done"
}

type InfoCommand struct {
	Str string
}

func (c InfoCommand) CommandType() string {
	return "info"
}

type EndCommand struct {
}

func (c EndCommand) CommandType() string {
	return "end"
}

type AboutCommand struct {
}

func (c AboutCommand) CommandType() string {
	return "about"
}
