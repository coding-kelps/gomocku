package stdio

import (
	"bufio"
	"os"
	"regexp"
	"fmt"
	"strings"

	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
	"github.com/coding-kelps/gomocku/pkg/domain/listener/ports"
	
	ai_models "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
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
		{"START", parseStart, regexp.MustCompile(`^START`)},
		{"TURN", parseTurn, regexp.MustCompile(`^TURN`)},
		{"BEGIN", parseBegin, regexp.MustCompile(`^BEGIN`)},
		{"BOARD", parseBoard, regexp.MustCompile(`^BOARD`)},
		{"BOARD", parseBoardTurn, regexp.MustCompile(`^\d+,\d+,\d`)},
		{"BOARD", parseBoardDone, regexp.MustCompile(`^DONE`)},
		{"INFO", parseInfo, regexp.MustCompile(`^INFO`)},
		{"END", parseEnd, regexp.MustCompile(`^END`)},
		{"ABOUT", parseAbout, regexp.MustCompile(`^ABOUT`)},
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
					cmd, err := p.caller(input)
					if err != nil {
						return err
					}
					ch <- cmd
					matched = true
				}
			}

			if !matched {
				ch <- models.UnknownCommand{}
			}
		} else {
			break
		}
	}

	return nil
}

func (std *Stdio) NotifyMove(p ai_models.Position) error {
	_, err := fmt.Printf("%d,%d\n", p.X, p.Y)
	if err != nil {
		return err
	}

	return nil
}

func (std *Stdio) NotifyReadiness() error {
	_, err := fmt.Printf("OK\n")
	if err != nil {
		return err
	}

	return nil
}

func (std *Stdio) NotifyUnknown() error {
	_, err := fmt.Printf("UNKNOWN unknown command\n")
	if err != nil {
		return err
	}

	return nil
}

func (std *Stdio) NotifyError(str string) error {
	_, err := fmt.Printf("ERROR %s\n", str)
	if err != nil {
		return err
	}

	return nil
}

func (std *Stdio) NotifyMessage(str string) error {
	_, err := fmt.Printf("MESSAGE %s\n", str)
	if err != nil {
		return err
	}

	return nil
}


func (std *Stdio) NotifyDebug(str string) error {
	_, err := fmt.Printf("DEBUG %s\n", str)
	if err != nil {
		return err
	}

	return nil
}

func (std *Stdio) NotifySuggestion(p ai_models.Position) error {
	_, err := fmt.Printf("SUGGEST %d,%d\n", p.X, p.Y)
	if err != nil {
		return err
	}

	return nil
}

func (std *Stdio) NotifyMetadata(metadata map[string]string) error {
	infos := make([]string, 0, len(metadata))
	for k, v := range metadata {
		infos = append(
			infos,
			fmt.Sprintf("%s=\"%s\"", k, v),
		)
	}

	_, err := fmt.Println(strings.Join(infos, ", "))
	if err != nil {
		return err
	}

	return nil
}
