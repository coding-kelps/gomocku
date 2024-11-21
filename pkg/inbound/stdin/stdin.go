package stdin

import (
	"regexp"
	"bufio"
	"os"

	"github.com/coding-kelps/gomocku/pkg/domain/mock/ports"
)

type Stdin struct {
	mock ports.Mock
	handlers []handler
}

type handler struct {
	name string
	caller func (s string) error
	regex *regexp.Regexp
}

func NewStdin(m ports.Mock) *Stdin {
	s := Stdin{mock: m}

	s.handlers = []handler{
		{"START", s.HandleStart, regexp.MustCompile(`START (\d+)`)},
		{"TURN", s.HandleTurn, regexp.MustCompile(`TURN (\d+),(\d+)`)},
		{"BEGIN", s.HandleBegin, regexp.MustCompile(`BEGIN`)},
		{"BOARD", s.HandleBoard, regexp.MustCompile(`BOARD`)},
		{"INFO", s.HandleInfo, regexp.MustCompile(`INFO`)},
		{"END", s.HandleEnd, regexp.MustCompile(`END`)},
		{"ABOUT", s.HandleAbout, regexp.MustCompile(`ABOUT`)},
	}

	return &s
}

func (s *Stdin) Run() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			input := scanner.Text()

			for _, h := range s.handlers {
				if h.regex.MatchString(input) {
					err := h.caller(input)
					if err != nil {
						return err
					}
				}
			}
		} else {
			break
		}
	}

	return nil
}
