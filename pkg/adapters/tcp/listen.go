package tcp

import (
	"fmt"
	"io"
	"net"

	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (tcp *TcpManagerInterface) Listen(actionsCh chan<-coordModels.ManagerAction, errorsCh chan<-error) {
	defer close(actionsCh)

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
			errorsCh <- err

			return
		}

		handler, ok := handlers[actionID[0]]
		if !ok {
			msg := fmt.Sprintf("unknown manager action with ID 0x%X", actionID[0])

            errorsCh <- NewManagerActionError(msg)
		}

		action, err := handler(tcp.conn)
        if err != nil {
			msg := fmt.Sprintf("processing failed of manager action with ID 0x%X: %v", actionID[0], err)

            errorsCh <- NewManagerActionError(msg)
        }

		actionsCh <- action
	}
}
