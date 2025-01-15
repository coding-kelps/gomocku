package stdio

import (
	"bufio"
	"os"
	"regexp"

	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
	"github.com/coding-kelps/gomocku/pkg/domain/listener/ports"
)

type Stdio struct {
	parsers []parser
	scanner  *bufio.Scanner

	ports.ManagerInterface
}

type parser struct {
	name   string
	caller func(s string)(models.ManagerCommand, error)
	regex  *regexp.Regexp
}

func NewStdio() *Stdio {
	s := Stdio{
		scanner: bufio.NewScanner(os.Stdin),
	}

	s.parsers = []parser{
		{"START", s.parseStart, regexp.MustCompile(`^START`)},
		{"TURN", s.parseTurn, regexp.MustCompile(`^TURN`)},
		{"BEGIN", s.parseBegin, regexp.MustCompile(`^BEGIN`)},
		{"BOARD", s.parseBoard, regexp.MustCompile(`^BOARD`)},
		{"BOARD", s.parseBoardTurn, regexp.MustCompile(`^\d+,\d+,\d`)},
		{"BOARD", s.parseBoardDone, regexp.MustCompile(`^DONE`)},
		{"INFO", s.parseInfo, regexp.MustCompile(`^INFO`)},
		{"END", s.parseEnd, regexp.MustCompile(`^END`)},
		{"ABOUT", s.parseAbout, regexp.MustCompile(`^ABOUT`)},
	}

	return &s
}

func (std *Stdio) Listen(ch chan<-models.ManagerCommand) error {
	for {
		if std.scanner.Scan() {
			input := std.scanner.Text()
			matched := false

			for _, p := range std.parsers {
				if p.regex.MatchString(input) {
					p.caller(input)
					matched = true
				}
			}

			if !matched {
				// TO DO: Implement unmatched patterns
				_ = 0
			}
		} else {
			break
		}
	}

	return nil
}
