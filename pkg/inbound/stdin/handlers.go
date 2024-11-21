package stdin

import (
	"regexp"
	"strconv"

	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

func (s *Stdin) HandleStart(input string) error {
	r := regexp.MustCompile(`START (\d+)`)
	m := r.FindStringSubmatch(input)
	
	size, err := strconv.ParseUint(m[1], 0, 8)
	if err != nil {
		return err
	}

	s.mock.RespondStart(uint8(size))

	return nil
}

func (s *Stdin) HandleTurn(input string) error {
	r := regexp.MustCompile(`TURN (\d+),(\d+)`)
	m := r.FindStringSubmatch(input)
	
	x, err := strconv.ParseUint(m[1], 0, 8)
	if err != nil {
		return err
	}

	y, err := strconv.ParseUint(m[2], 0, 8)
	if err != nil {
		return err
	}

	s.mock.RespondTurn(models.Position{
		X: uint8(x),
		Y: uint8(y),
	})

	return nil
}

func (s *Stdin) HandleBegin(input string) error {
	s.mock.RespondBegin()

	return nil
}

func (s *Stdin) HandleBoard(input string) error {
	s.mock.RespondBoard([]models.Position{})

	return nil
}

func (s *Stdin) HandleInfo(input string) error {
	s.mock.RespondInfo()

	return nil
}

func (s *Stdin) HandleEnd(input string) error {
	s.mock.RespondEnd()

	return nil
}

func (s *Stdin) HandleAbout(input string) error {
	s.mock.RespondAbout()

	return nil
}
