package stdio

import (
	"fmt"
	"strings"

	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

func (std *Stdio) handleStart(s string) error {
	size, err := std.parseStartArgs(s)
	if err != nil {
		return err
	}

	err = std.mock.RespondStart(size)
	if err != nil {
		return err
	}

	return nil
}

func (std *Stdio) handleTurn(s string) error {
	p, err := std.parseTurnArgs(s)
	if err != nil {
		return err
	}

	move, err := std.mock.RespondTurn(p)
	if err != nil {
		return err
	}

	fmt.Printf("%d,%d\n", move.X, move.Y)

	return nil
}

func (std *Stdio) handleBegin(input string) error {
	move, err := std.mock.RespondBegin()
	if err != nil {
		return err
	}

	fmt.Printf("%d,%d\n", move.X, move.Y)

	return nil
}

func (std *Stdio) handleBoard(input string) error {
	move, err := std.mock.RespondBoard([]models.Position{})
	if err != nil {
		return err
	}

	fmt.Printf("%d,%d\n", move.X, move.Y)

	return nil
}

func (std *Stdio) handleInfo(input string) error {
	std.mock.RespondInfo()

	return nil
}

func (std *Stdio) handleEnd(input string) error {
	std.mock.RespondEnd()

	return nil
}

func (std *Stdio) handleAbout(input string) error {
	about, err := std.mock.RespondAbout()
	if err != nil {
		return err
	}

	infos := make([]string, 0, len(about))
	for k, v := range about {
		infos = append(
			infos,
			fmt.Sprintf("%s=\"%s\"", k, v),
		)
	}

	fmt.Println(strings.Join(infos, ", "))

	return nil
}
