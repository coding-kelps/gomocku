package main

import (
	"fmt"

	"github.com/coding-kelps/gomocku/pkg/adapters/tcp"
	"github.com/coding-kelps/gomocku/pkg/domain/ai"
	"github.com/coding-kelps/gomocku/pkg/domain/listener"
)

func main() {
	tcp_interface, err := tcp.NewTCP("localhost:9000")
	if err != nil {
		fmt.Printf("%e\n", err)
	}

	ai := ai.NewRandomAI()
	listener := listener.NewListener(tcp_interface, ai)

	err = listener.Listen()
	if err != nil {
		fmt.Printf("%e\n", err)
	}
}
