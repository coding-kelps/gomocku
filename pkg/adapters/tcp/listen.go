package tcp

import (
	"fmt"
	"io"

	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (tcp *TcpManagerInterface) Listen(actionsCh chan<-coordModels.ManagerAction, errorsCh chan<-error) {
	defer close(actionsCh)

	handlers := map[byte]func()(coordModels.ManagerAction, error){
		StartManagerActionID: 			tcp.StartHandler,
		RestartManagerActionID:			tcp.RestartHandler,
		TurnManagerActionID: 			tcp.TurnHandler,
		BeginManagerActionID: 			tcp.BeginHandler,
		BoardManagerActionID:			tcp.BoardHandler,
		InfoManagerActionID: 			tcp.InfoHandler,
		ResultManagerActionID:			tcp.ResultHandler,
		EndManagerActionID: 			tcp.EndHandler,
		AboutManagerActionID: 			tcp.AboutHandler,
		UnknownManagerActionID:			tcp.UnknownHandler,
		ErrorManagerActionID:			tcp.ErrorHandler,
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

		action, err := handler()
        if err != nil {
			msg := fmt.Sprintf("processing failed of manager action with ID 0x%X: %v", actionID[0], err)

            errorsCh <- NewManagerActionError(msg)
        }

		actionsCh <- action
	}
}
