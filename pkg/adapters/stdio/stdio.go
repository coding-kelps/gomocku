package stdio

import (
	"bufio"
	"os"
	"regexp"

	coordModels	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
	coordPorts	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/ports"	
)

type StdioManagerInterface struct {
	parsers []parser
	scanner  *bufio.Scanner

	coordPorts.ManagerInterface
}

type parser struct {
	name   string
	caller func(s string)(coordModels.ManagerAction, error)
	regex  *regexp.Regexp
}	

func NewStdioManagerInterface() *StdioManagerInterface {
	s := StdioManagerInterface{
		scanner: bufio.NewScanner(os.Stdin),
	}

	s.parsers = []parser{
		{"START", 	s.parseStart,			regexp.MustCompile(`^START`)},
		{"TURN", 	s.parseTurn, 			regexp.MustCompile(`^TURN`)},
		{"BEGIN", 	s.parseBegin, 			regexp.MustCompile(`^BEGIN`)},
		{"BOARD", 	s.parseBoardBegin,		regexp.MustCompile(`^BOARD`)},
		{"INFO", 	s.parseInfo,			regexp.MustCompile(`^INFO`)},
		{"END", 	s.parseEnd,				regexp.MustCompile(`^END`)},
		{"ABOUT",	s.parseAbout,			regexp.MustCompile(`^ABOUT`)},
	}

	return &s
}
