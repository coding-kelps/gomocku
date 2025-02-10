package stdio

import (
	"fmt"
	"regexp"

	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (std *StdioManagerInterface) parseStart(input string) (coordModels.ManagerAction, error) {
	size, err := parseStartArgs(input)
	if err != nil {
		return coordModels.StartAction{}, err
	}

	return coordModels.StartAction{
		Size: size,
	}, nil
}

func (std *StdioManagerInterface) parseRestart(input string) (coordModels.ManagerAction, error) {
	return coordModels.ResultAction{}, nil
}

func (std *StdioManagerInterface) parseTurn(input string) (coordModels.ManagerAction, error) {
	p, err := parseTurnArgs(input)
	if err != nil {
		return coordModels.TurnAction{}, err
	}

	return coordModels.TurnAction{
		Position: p,
	}, err
}

func (std *StdioManagerInterface) parseBegin(input string) (coordModels.ManagerAction, error) {
	return coordModels.BeginAction{}, nil
}

func (std *StdioManagerInterface) parseBoardBegin(input string) (coordModels.ManagerAction, error) {
	turnRe	:= 	regexp.MustCompile(`^\d+,\d+,\d`)
	doneRe	:= 	regexp.MustCompile(`^DONE`)
	board	:= 	coordModels.BoardAction{}

	for {
		if std.scanner.Scan() {
			input := std.scanner.Text()

			if turnRe.MatchString(input) {
				t, err := parseBoardTurnArgs(input)
				if err != nil {
					return nil, err
				}

				board.Turns = append(board.Turns, t)
			} else if doneRe.MatchString(input) {
				return board, nil
			} else {
				msg := fmt.Sprintf("unexpected manager action while registering board action: %s", input)

				return nil, NewManagerActionError(msg)
			}
		}
	}
}

func (std *StdioManagerInterface) parseInfo(input string) (coordModels.ManagerAction, error) {
	return coordModels.InfoAction{}, nil
}

func (std *StdioManagerInterface) parseResult(input string) (coordModels.ManagerAction, error) {
	result, err := parseResultArgs(input)
	if err != nil {
		return nil, err
	}

	return coordModels.ResultAction{
		Result: result,
	}, nil
}

func (std *StdioManagerInterface) parseEnd(input string) (coordModels.ManagerAction, error) {
	return coordModels.EndAction{}, nil
}

func (std *StdioManagerInterface) parseAbout(input string) (coordModels.ManagerAction, error) {
	return coordModels.AboutAction{}, nil
}
