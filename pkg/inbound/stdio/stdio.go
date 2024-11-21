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
	running bool
}

type handler struct {
	name string
	caller func (s string)
	regex *regexp.Regexp
}

func NewStdio(m ports.Mock) *Stdio {
	s := Stdio{mock: m}

	s.handlers = []handler{
		{"START", s.handleStart, regexp.MustCompile(`^START`)},
		{"TURN", s.handleTurn, regexp.MustCompile(`^TURN`)},
		{"BEGIN", s.handleBegin, regexp.MustCompile(`^BEGIN`)},
		{"BOARD", s.handleBoard, regexp.MustCompile(`^BOARD`)},
		{"INFO", s.handleInfo, regexp.MustCompile(`^INFO`)},
		{"END", s.handleEnd, regexp.MustCompile(`^END`)},
		{"ABOUT", s.handleAbout, regexp.MustCompile(`^ABOUT`)},
	}

	return &s
}

func (std *Stdio) Run() error {
	scanner := bufio.NewScanner(os.Stdin)
	std.running = true

	for std.running {
		if scanner.Scan() {
			input := scanner.Text()
			matched := false

			for _, h := range std.handlers {
				if h.regex.MatchString(input) {
					h.caller(input)
					matched = true
				}
			}

			if !matched {
				std.handleUnknown(input)
			}
		} else {
			break
		}
	}

	return nil
}
