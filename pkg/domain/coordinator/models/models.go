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

func (a StartAction) ActionType() string {
	return "start"
}

type TurnAction struct {
	Position aiModels.Position
}

func (a TurnAction) ActionType() string {
	return "turn"
}

type BeginAction struct {
}

func (a BeginAction) ActionType() string {
	return "begin"
}

type BoardAction struct {
	Turns []aiModels.Turn
}

func (a BoardAction) ActionType() string {
	return "board"
}

type InfoAction struct {
	Str string
}

func (a InfoAction) ActionType() string {
	return "info"
}

type EndAction struct {
}

func (a EndAction) ActionType() string {
	return "end"
}

type AboutAction struct {
}

func (a AboutAction) ActionType() string {
	return "about"
}

type UnknownAction struct {
	Msg string
}

func (a UnknownAction) ActionType() string {
	return "unknown"
}

type ErrorAction struct {
	Msg string
}

func (a ErrorAction) ActionType() string {
	return "error"
}

