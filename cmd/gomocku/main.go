package main

import (
	"github.com/coding-kelps/gomocku/pkg/domain/mock"
	"github.com/coding-kelps/gomocku/pkg/adapters"
)

func main() {
	stdio := inbound.NewStdio(mock.NewMock())

	if err := stdio.Run(); err != nil {
		panic(err)
	}
}
