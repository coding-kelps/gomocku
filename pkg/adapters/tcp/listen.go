package tcp

import (
	"fmt"
	"io"
	"net"

	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (tcp *TcpManagerInterface) Listen(ch chan<-coordModels.ManagerAction) error {
	handlers := map[byte]func(conn net.Conn)(coordModels.ManagerAction, error){
		StartManagerActionID: 			StartHandler,
		TurnManagerActionID: 			TurnHandler,
		BeginManagerActionID: 			BeginHandler,
		BoardBeginManagerActionID:		BoardBeginHandler,
		BoardTurnManagerActionID: 		BoardTurnHandler,
		BoardEndManagerActionID: 		BoardDoneHandler,
		InfoManagerActionID: 			InfoHandler,	
		EndManagerActionID: 			EndHandler,
		AboutManagerActionID: 			AboutHandler,
		UnknownManagerActionID:			UnknownHandler,
		ErrorManagerActionID:			ErrorHandler,
	}

	for {
		var actionID [1]byte
		if _, err := io.ReadFull(tcp.conn, actionID[:]); err != nil {
			return err
		}

		handler, ok := handlers[actionID[0]]
		if !ok {
			msg := fmt.Sprintf("unknown manager action with ID 0x%X", actionID[0])

            return NewManagerActionError(msg)
		}

		action, err := handler(tcp.conn)
        if err != nil {
			msg := fmt.Sprintf("processing failed of manager action with ID 0x%X: %v", actionID[0], err)

            return NewManagerActionError(msg)
        }

		ch <- action
	}
}
