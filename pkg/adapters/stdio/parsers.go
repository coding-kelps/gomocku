package stdio

import (
	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
)

func (std *Stdio) parseStart(input string) (models.ManagerCommand, error) {
	size, err := std.parseStartArgs(input)
	if err != nil {
		return models.StartCommand{}, err
	}

	return &models.StartCommand{
		Size: size,
	}, nil
}

func (std *Stdio) parseTurn(input string) (models.ManagerCommand, error) {
	p, err := std.parseTurnArgs(input)
	if err != nil {
		return &models.TurnCommand{}, err
	}

	return &models.TurnCommand{
		Position: p,
	}, err
}

func (std *Stdio) parseBegin(input string) (models.ManagerCommand, error) {
	return &models.BeginCommand{}, nil
}

func (std *Stdio) parseBoard(input string) (models.ManagerCommand, error) {
	return &models.BoardCommand{}, nil
}

func (std *Stdio) parseBoardTurn(input string) (models.ManagerCommand, error) {
	t, err := std.parseBoardTurnArgs(input)
	if err != nil {
		return &models.BoardTurnCommand{}, err
	}

	return &models.BoardTurnCommand{
		Turn: t,
	}, nil
}

func (std *Stdio) parseBoardDone(input string) (models.ManagerCommand, error) {
	return &models.BoardDoneCommand{}, nil
}

func (std *Stdio) parseInfo(input string) (models.ManagerCommand, error) {
	return &models.InfoCommand{}, nil
}

func (std *Stdio) parseEnd(input string) (models.ManagerCommand, error) {
	return &models.EndCommand{}, nil
}

func (std *Stdio) parseAbout(input string) (models.ManagerCommand, error) {
	return &models.AboutCommand{}, nil
}
