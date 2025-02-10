package tcp

import (
	"fmt"
	"io"
	"encoding/binary"

	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
	aiModels 	"github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func (tcp *TcpManagerInterface) StartHandler() (coordModels.ManagerAction, error) {
	payload := make([]byte, 1)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	return coordModels.StartAction{
		Size: payload[0],
	}, nil
}

func (tcp *TcpManagerInterface) RestartHandler() (coordModels.ManagerAction, error) {
	return coordModels.RestartAction{}, nil
}

func (tcp *TcpManagerInterface) TurnHandler() (coordModels.ManagerAction, error) {
	payload := make([]byte, 2)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	return coordModels.TurnAction{
		Position: aiModels.Position{
			X: payload[0],
			Y: payload[1],
		},
	}, nil
}

func (tcp *TcpManagerInterface) BeginHandler() (coordModels.ManagerAction, error) {
	return coordModels.BeginAction{}, nil
}

func (tcp *TcpManagerInterface) BoardHandler() (coordModels.ManagerAction, error) {
	board := coordModels.BoardAction{}

	payload := make([]byte, 4)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	nbTurn := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, nbTurn * 3)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	for i := uint32(0); i < nbTurn; i += 3 {
		turn := aiModels.Turn{
			Position: aiModels.Position{
				X: payload[0],
				Y: payload[1],
			},
		}

		if payload[2] == 0 {
			turn.Player = aiModels.Us
		} else {
			turn.Player = aiModels.Opponent
		}

		board.Turns = append(board.Turns, turn)
	}

	return board, nil
}

func (tcp *TcpManagerInterface) InfoHandler() (coordModels.ManagerAction, error) {
	payload := make([]byte, 4)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}


	infoSize := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, infoSize)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	return coordModels.InfoAction{
		Str: string(payload),
	}, nil
}

func (tcp *TcpManagerInterface) ResultHandler() (coordModels.ManagerAction, error) {
	payload := make([]byte, 1)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	var result aiModels.GameEnd

	switch payload[0] {
	case 0:
		result = aiModels.Draw
	case 1:
		result = aiModels.Win
	case 2:
		result = aiModels.Loose
	default:
		return nil, fmt.Errorf("invalid result")
	}

	return coordModels.ResultAction{
		Result: result,
	}, nil
}

func (tcp *TcpManagerInterface) EndHandler() (coordModels.ManagerAction, error) {
	return coordModels.EndAction{}, nil
}

func (tcp *TcpManagerInterface) AboutHandler() (coordModels.ManagerAction, error) {
	return coordModels.AboutAction{}, nil
}

func (tcp *TcpManagerInterface) UnknownHandler() (coordModels.ManagerAction, error) {
	payload := make([]byte, 4)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	msgSize := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, msgSize)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	return coordModels.UnknownAction{
		Msg: string(payload),
	}, nil
}

func (tcp *TcpManagerInterface) ErrorHandler() (coordModels.ManagerAction, error) {
	payload := make([]byte, 4)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	msgSize := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, msgSize)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return nil, err
	}

	return coordModels.ErrorAction{
		Msg: string(payload),
	}, nil
}
