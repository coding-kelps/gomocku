package stdio

import (
	"regexp"
	"bufio"
	"os"

	"github.com/coding-kelps/gomocku/pkg/domain/mock/ports"
)

type Stdio struct {
	mock ports.Mock
	handlers []handler
}

type handler struct {
	name string
	caller func (s string) error
	regex *regexp.Regexp
}

func NewStdio(m ports.Mock) *Stdio {
	s := Stdio{mock: m}

	s.handlers = []handler{
		{"START", s.handleStart, regexp.MustCompile(`^START (\d+)`)},
		{"TURN", s.handleTurn, regexp.MustCompile(`^TURN (\d+),(\d+)`)},
		{"BEGIN", s.handleBegin, regexp.MustCompile(`^BEGIN`)},
		{"BOARD", s.handleBoard, regexp.MustCompile(`^BOARD`)},
		{"INFO", s.handleInfo, regexp.MustCompile(`^INFO`)},
		{"END", s.handleEnd, regexp.MustCompile(`^END`)},
		{"ABOUT", s.handleAbout, regexp.MustCompile(`^ABOUT`)},
	}

	return &s
}

func (s *Stdio) Run() error {
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
