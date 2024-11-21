package main

import (
	"log"

	"github.com/coding-kelps/gomocku/pkg/domain/mock"
	"github.com/coding-kelps/gomocku/pkg/inbound/stdio"
)

func main() {
	mock := mock.NewMock()

	stdio := stdio.NewStdio(mock)

	if err := stdio.Run(); err != nil {
		log.Fatal(err)
	}
}
