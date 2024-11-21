package main

import (
	"fmt"
	"log"

	"github.com/coding-kelps/gomocku/pkg/domain/mock"
	"github.com/coding-kelps/gomocku/pkg/inbound/stdin"
	"github.com/coding-kelps/gomocku/pkg/outbound"
)

func main() {
	fmt.Println("ça gomock")

	stdout := outbound.NewStdout()

	mock := mock.NewMock(stdout)

	stdin := stdin.NewStdin(mock)

	if err := stdin.Run(); err != nil {
		log.Fatal(err)
	}
}