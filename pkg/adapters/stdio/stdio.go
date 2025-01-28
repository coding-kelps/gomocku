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
		{"START", 	parseStart,			regexp.MustCompile(`^START`)},
		{"TURN", 	parseTurn, 			regexp.MustCompile(`^TURN`)},
		{"BEGIN", 	parseBegin, 		regexp.MustCompile(`^BEGIN`)},
		{"BOARD", 	parseBoardBegin,	regexp.MustCompile(`^BOARD`)},
		{"BOARD", 	parseBoardTurn,		regexp.MustCompile(`^\d+,\d+,\d`)},
		{"BOARD", 	parseBoardDone,		regexp.MustCompile(`^DONE`)},
		{"INFO", 	parseInfo,			regexp.MustCompile(`^INFO`)},
		{"END", 	parseEnd,			regexp.MustCompile(`^END`)},
		{"ABOUT",	parseAbout,			regexp.MustCompile(`^ABOUT`)},
	}

	return &s
}
