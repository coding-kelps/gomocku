package stdio

import (
	"fmt"
	"strings"
	"regexp"

	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

func (std *Stdio) handleStart(input string) {
	size, err := std.parseStartArgs(input)
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}

	err = std.mock.RespondStart(size)
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}
}

func (std *Stdio) handleTurn(input string) {
	p, err := std.parseTurnArgs(input)
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}

	move, err := std.mock.RespondTurn(p)
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}

	fmt.Printf("%d,%d\n", move.X, move.Y)
}

func (std *Stdio) handleBegin(input string) {
	move, err := std.mock.RespondBegin()
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}

	fmt.Printf("%d,%d\n", move.X, move.Y)
}

func (std *Stdio) handleBoard(input string) {
	turns := []models.Turn{}
	done := false
	r := regexp.MustCompile(`^DONE$`)
	
	for !done {
		if std.scanner.Scan() {
			i := std.scanner.Text()
			
			if r.MatchString(i) {
				done = true
			} else {
				turn, err := std.parseBoardTurn(i)
				if err != nil {
					fmt.Printf("ERROR %s\n", err)
					continue
				}

				turns = append(turns, turn)
			}
		}
	}

	move, err := std.mock.RespondBoard(turns)
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}

	fmt.Printf("%d,%d\n", move.X, move.Y)
}

func (std *Stdio) handleInfo(input string) {
	err := std.mock.RespondInfo()
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}
}

func (std *Stdio) handleEnd(input string) {
	err := std.mock.RespondEnd()
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}

	std.running = false
}

func (std *Stdio) handleAbout(input string) {
	about, err := std.mock.RespondAbout()
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
		return
	}

	infos := make([]string, 0, len(about))
	for k, v := range about {
		infos = append(
			infos,
			fmt.Sprintf("%s=\"%s\"", k, v),
		)
	}

	fmt.Println(strings.Join(infos, ", "))
}

func (std *Stdio) handleUnknown(input string) {
	fmt.Printf("UNKNOWN unknown command \"%s\"\n", input)
}
