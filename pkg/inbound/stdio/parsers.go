package stdio

import (
	"regexp"
	"strconv"
	"fmt"

	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
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

func (std *Stdio) parseStartArgs(s string) (uint8, error) {
	r := regexp.MustCompile(`^START (\d+)$`)
	m := r.FindStringSubmatch(s)

	size, err := strconv.ParseUint(m[1], 0, 8)
	if err != nil {
		return 0, err
	}

	return uint8(size), nil
}

func (std *Stdio) parseTurnArgs(s string) (models.Position, error) {
	r := regexp.MustCompile(`^TURN (\d+,\d+)$`)
	m := r.FindStringSubmatch(s)
	if len(m) != 2 {
		return models.Position{}, &InvalidFormatError{s, r}
	}

	return std.parsePosition(m[1])
}

func (std *Stdio) parsePosition(s string) (models.Position, error) {
	r := regexp.MustCompile(`^(\d+),(\d+)$`)
	m := r.FindStringSubmatch(s)
	if len(m) != 3 {
		return models.Position{}, &InvalidFormatError{s, r}
	}

	x, err := strconv.ParseUint(m[1], 0, 8)
	if err != nil {
		return models.Position{}, err
	}

	y, err := strconv.ParseUint(m[2], 0, 8)
	if err != nil {
		return models.Position{}, err
	}

	p := models.Position{
		X: uint8(x),
		Y: uint8(y),
	}

	return p, nil
}