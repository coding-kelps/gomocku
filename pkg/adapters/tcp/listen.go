package tcp

import (
	"fmt"
	"io"
	"net"

	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (tcp *TcpManagerInterface) Listen(ch chan<-coordModels.ManagerAction) error {
	handlers := map[byte]func(conn net.Conn)(coordModels.ManagerAction, error){
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
		if _, err := io.ReadFull(tcp.conn, actionID[:]); err != nil {
			return err
		}

		handler, ok := handlers[actionID[0]]
		if !ok {
            fmt.Printf("Unknown command ID %d received, closing connection", actionID[0])

            return nil
		}

		action, err := handler(tcp.conn)
        if err != nil {
            fmt.Printf("Error handling command %d: %v", actionID[0], err)

            continue
        }

		ch <- action
	}
}
