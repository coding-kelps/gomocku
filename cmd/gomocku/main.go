package main

import (
	"fmt"

	"github.com/coding-kelps/gomocku/pkg/adapters"
	"github.com/coding-kelps/gomocku/pkg/domain/ai"
	"github.com/coding-kelps/gomocku/pkg/domain/listener"
)

func main() {
	stdio := adapters.NewStdio()
	ai := ai.NewRandomAI()
	listener := listener.NewListener(stdio, ai)

	err := listener.Listen()
	if err != nil {
		fmt.Printf("%e\n", err)
	}
}
