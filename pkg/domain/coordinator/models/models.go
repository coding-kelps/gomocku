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
	return "START"
}

type RestartAction struct {
}

func (a RestartAction) ActionType() string {
	return "RESTART"
}

type TurnAction struct {
	Position aiModels.Position
}

func (a TurnAction) ActionType() string {
	return "TURN"
}

type BeginAction struct {
}

func (a BeginAction) ActionType() string {
	return "BEGIN"
}

type BoardAction struct {
	Turns []aiModels.Turn
}

func (a BoardAction) ActionType() string {
	return "BOARD"
}

type InfoAction struct {
	Str string
}

func (a InfoAction) ActionType() string {
	return "INFO"
}

type ResultAction struct {
	Result aiModels.GameEnd
}

func (a ResultAction) ActionType() string {
	return "RESULT"
}

type EndAction struct {
}

func (a EndAction) ActionType() string {
	return "END"
}

type AboutAction struct {
}

func (a AboutAction) ActionType() string {
	return "ABOUT"
}

type UnknownAction struct {
	Msg string
}

func (a UnknownAction) ActionType() string {
	return "UNKNOWN"
}

type ErrorAction struct {
	Msg string
}

func (a ErrorAction) ActionType() string {
	return "ERROR"
}

