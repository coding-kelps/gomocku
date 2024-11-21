package outbound

import (
	"fmt"

	"github.com/coding-kelps/gomocku/pkg/domain/mock/ports"
	"github.com/coding-kelps/gomocku/pkg/domain/mock/models"
)

type Stdout struct {
	ports.ResponseSender
}

func NewStdout() ports.ResponseSender {
	return &Stdout{}
}

func (s *Stdout) SendMove(p models.Position) error {
	fmt.Printf("%d,%d\n", p.X, p.Y)

	return nil
}

func (s *Stdout) SendUnknown(msg string) error {
	fmt.Printf("UNKNOWN %s\n", s)

	return nil
}

func (s *Stdout) SendError(msg string) error {
	fmt.Printf("UNKNOWN %s\n", s)

	return nil
}

func (s *Stdout) SendDebug(msg string) error {
	fmt.Printf("UNKNOWN %s\n", s)

	return nil
}
