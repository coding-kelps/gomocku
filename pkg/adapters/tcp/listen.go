package tcp

import (
	"fmt"
	"io"
	"net"

	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
)

func (tcp *Tcp) Listen(ch chan<-models.ManagerCommand) error {
	for {
        conn, err := tcp.listener.Accept()
        if err != nil {
			fmt.Println("Error:", err)
			continue
        }

		go handleConnection(conn, ch)
	}
}

func handleConnection(conn net.Conn, ch chan<-models.ManagerCommand) {
	defer conn.Close()

	handlers := map[byte]func(conn net.Conn)(models.ManagerCommand, error){
		StartActionID: 			StartHandler,
		TurnActionID: 			TurnHandler,
		BeginActionID: 			BeginHandler,
		BoardBeginActionID: 	BoardBeginHandler,
		BoardTurnActionID: 		BoardTurnHandler,
		BoardEndActionID: 		BoardDoneHandler,
		InfoActionID: 			InfoHandler,	
		EndActionID: 			EndHandler,
		AboutActionID: 			AboutHandler,
	}

	for {
		var actionID [1]byte
		if _, err := io.ReadFull(conn, actionID[:]); err != nil {
			if err == io.EOF {
				fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
			} else {
				fmt.Printf("Error reading command ID from %s: %v", conn.RemoteAddr(), err)
			}

			return
		}

		handler, ok := handlers[actionID[0]]
		if !ok {
            fmt.Printf("Unknown command ID %d received, closing connection", actionID[0])

            return
		}

		action, err := handler(conn)
        if err != nil {
            fmt.Printf("Error handling command %d: %v", actionID[0], err)

            continue
        }

		ch <- action
	}
}
