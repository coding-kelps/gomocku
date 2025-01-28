package stdio

import (
	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func parseStart(input string) (coordModels.ManagerAction, error) {
	size, err := parseStartArgs(input)
	if err != nil {
		return coordModels.StartAction{}, err
	}

	return coordModels.StartAction{
		Size: size,
	}, nil
}

func parseTurn(input string) (coordModels.ManagerAction, error) {
	p, err := parseTurnArgs(input)
	if err != nil {
		return coordModels.TurnAction{}, err
	}

	return coordModels.TurnAction{
		Position: p,
	}, err
}

func parseBegin(input string) (coordModels.ManagerAction, error) {
	return coordModels.BeginAction{}, nil
}

func parseBoardBegin(input string) (coordModels.ManagerAction, error) {
	return coordModels.BoardBeginAction{}, nil
}

func parseBoardTurn(input string) (coordModels.ManagerAction, error) {
	t, err := parseBoardTurnArgs(input)
	if err != nil {
		return coordModels.BoardTurnAction{}, err
	}

	return coordModels.BoardTurnAction{
		Turn: t,
	}, nil
}

func parseBoardDone(input string) (coordModels.ManagerAction, error) {
	return coordModels.BoardDoneAction{}, nil
}

func parseInfo(input string) (coordModels.ManagerAction, error) {
	return coordModels.InfoAction{}, nil
}

func parseEnd(input string) (coordModels.ManagerAction, error) {
	return coordModels.EndAction{}, nil
}

func parseAbout(input string) (coordModels.ManagerAction, error) {
	return coordModels.AboutAction{}, nil
}
