package models

import (
	aiModels "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

type ManagerAction interface {
	ActionType() string
}

type StartAction struct {
	Size uint8
}

func (c StartAction) ActionType() string {
	return "start"
}

type TurnAction struct {
	Position aiModels.Position
}

func (c TurnAction) ActionType() string {
	return "turn"
}

type BeginAction struct {
}

func (c BeginAction) ActionType() string {
	return "begin"
}

type BoardBeginAction struct {
}


func (c BoardBeginAction) ActionType() string {
	return "board_begin"
}

type BoardTurnAction struct {
	Turn aiModels.Turn
}

func (c BoardTurnAction) ActionType() string {
	return "board_turn"
}

type BoardDoneAction struct {
}

func (c BoardDoneAction) ActionType() string {
	return "board_done"
}

type InfoAction struct {
	Str string
}

func (c InfoAction) ActionType() string {
	return "info"
}

type EndAction struct {
}

func (c EndAction) ActionType() string {
	return "end"
}

type AboutAction struct {
}

func (c AboutAction) ActionType() string {
	return "about"
}

type UnknownAction struct {
}

func (c UnknownAction) ActionType() string {
	return "unknown"
}
