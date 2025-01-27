package stdio

import (
	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
)

func parseStart(input string) (models.ManagerCommand, error) {
	size, err := parseStartArgs(input)
	if err != nil {
		return models.StartCommand{}, err
	}

	return models.StartCommand{
		Size: size,
	}, nil
}

func parseTurn(input string) (models.ManagerCommand, error) {
	p, err := parseTurnArgs(input)
	if err != nil {
		return models.TurnCommand{}, err
	}

	return models.TurnCommand{
		Position: p,
	}, err
}

func parseBegin(input string) (models.ManagerCommand, error) {
	return models.BeginCommand{}, nil
}

func parseBoardBegin(input string) (models.ManagerCommand, error) {
	return models.BoardBeginCommand{}, nil
}

func parseBoardTurn(input string) (models.ManagerCommand, error) {
	t, err := parseBoardTurnArgs(input)
	if err != nil {
		return models.BoardTurnCommand{}, err
	}

	return models.BoardTurnCommand{
		Turn: t,
	}, nil
}

func parseBoardDone(input string) (models.ManagerCommand, error) {
	return models.BoardDoneCommand{}, nil
}

func parseInfo(input string) (models.ManagerCommand, error) {
	return models.InfoCommand{}, nil
}

func parseEnd(input string) (models.ManagerCommand, error) {
	return models.EndCommand{}, nil
}

func parseAbout(input string) (models.ManagerCommand, error) {
	return models.AboutCommand{}, nil
}
