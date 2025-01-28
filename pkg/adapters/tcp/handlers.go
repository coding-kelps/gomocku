package tcp

import (
	"io"
	"net"
	"encoding/binary"

	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
	aiModels 	"github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func StartHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	payload := make([]byte, 1)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	return coordModels.StartAction{
		Size: payload[0],
	}, nil
}

func TurnHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	payload := make([]byte, 2)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	return coordModels.TurnAction{
		Position: aiModels.Position{
			X: payload[0],
			Y: payload[1],
		},
	}, nil
}

func BeginHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	return coordModels.BeginAction{}, nil
}

func BoardHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	board := coordModels.BoardAction{}

	payload := make([]byte, 4)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	nbTurn := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, nbTurn * 3)
	if _, err := io.ReadFull(conn, payload); err != nil {
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

func InfoHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	payload := make([]byte, 4)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}


	infoSize := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, infoSize)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	return coordModels.InfoAction{
		Str: string(payload),
	}, nil
}

func EndHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	return coordModels.EndAction{}, nil
}

func AboutHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	return coordModels.AboutAction{}, nil
}

func UnknownHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	payload := make([]byte, 4)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	msgSize := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, msgSize)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	return coordModels.UnknownAction{
		Msg: string(payload),
	}, nil
}

func ErrorHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	payload := make([]byte, 4)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	msgSize := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, msgSize)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	return coordModels.ErrorAction{
		Msg: string(payload),
	}, nil
}
