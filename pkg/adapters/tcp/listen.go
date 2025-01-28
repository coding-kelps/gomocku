package tcp

import (
	"fmt"
	"io"
	"net"

	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
)

func (tcp *Tcp) Listen(ch chan<-models.ManagerCommand) error {
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
		if _, err := io.ReadFull(tcp.connection, actionID[:]); err != nil {
			return err
		}

		handler, ok := handlers[actionID[0]]
		if !ok {
            fmt.Printf("Unknown command ID %d received, closing connection", actionID[0])

            return nil
		}

		action, err := handler(tcp.connection)
        if err != nil {
            fmt.Printf("Error handling command %d: %v", actionID[0], err)

            continue
        }

		ch <- action
	}
}
