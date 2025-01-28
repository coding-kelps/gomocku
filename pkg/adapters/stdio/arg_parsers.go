package stdio

import (
	"fmt"
	"regexp"
	"strconv"

	aiModels "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

type InvalidFormatError struct {
	s string
	r *regexp.Regexp
}

func (e *InvalidFormatError) Error() string {
	return fmt.Sprintf(
		"input string \"%s\" doesn't match expected regular expression: \"%s\"",
		e.s,
		e.r.String(),
	)
}

func parseStartArgs(s string) (uint8, error) {
	r := regexp.MustCompile(`^START (\d+)$`)
	m := r.FindStringSubmatch(s)

	size, err := strconv.ParseUint(m[1], 0, 8)
	if err != nil {
		return 0, err
	}

	return uint8(size), nil
}

func parseTurnArgs(s string) (aiModels.Position, error) {
	r := regexp.MustCompile(`^TURN (\d+,\d+)$`)
	m := r.FindStringSubmatch(s)
	if len(m) != 2 {
		return aiModels.Position{}, &InvalidFormatError{s: s, r: r}
	}

	return parsePosition(m[1])
}

func parsePosition(s string) (aiModels.Position, error) {
	r := regexp.MustCompile(`^(\d+),(\d+)$`)
	m := r.FindStringSubmatch(s)
	if len(m) != 3 {
		return aiModels.Position{}, &InvalidFormatError{s: s, r: r}
	}

	x, err := strconv.ParseUint(m[1], 0, 8)
	if err != nil {
		return aiModels.Position{}, err
	}

	y, err := strconv.ParseUint(m[2], 0, 8)
	if err != nil {
		return aiModels.Position{}, err
	}

	p := aiModels.Position{
		X: uint8(x),
		Y: uint8(y),
	}

	return p, nil
}

type WrongFieldError struct {
}

func (e *WrongFieldError) Error() string {
	return "playing field can only be either 1 (Us) or 2 (Opponent)"
}

func parseBoardTurnArgs(s string) (aiModels.Turn, error) {
	r := regexp.MustCompile(`^(\d+,\d+),(\d)$`)
	m := r.FindStringSubmatch(s)
	if len(m) != 3 {
		return aiModels.Turn{}, &InvalidFormatError{s: s, r: r}
	}

	position, err := parsePosition(m[1])
	if err != nil {
		return aiModels.Turn{}, err
	}

	nb, err := strconv.ParseUint(m[2], 0, 8)
	if err != nil {
		return aiModels.Turn{}, err
	}

	player := aiModels.Player(nb - 1)
	if !(player == aiModels.Us || player == aiModels.Opponent) {
		return aiModels.Turn{}, &WrongFieldError{}
	}

	return aiModels.Turn{
		Position: position,
		Player:   player,
	}, nil
}
