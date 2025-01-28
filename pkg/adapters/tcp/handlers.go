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

func BoardBeginHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	return coordModels.BoardBeginAction{}, nil
}

func BoardTurnHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	var player aiModels.Player

	payload := make([]byte, 3)
	if _, err := io.ReadFull(conn, payload); err != nil {
		return nil, err
	}

	if payload[2] == 0 {
		player = aiModels.Us
	} else {
		player = aiModels.Opponent
	}

	return coordModels.BoardTurnAction{
		Turn: aiModels.Turn{
			Position: aiModels.Position{
				X: payload[0],
				Y: payload[1],
			},
			Player: player,
		},
	}, nil
}

func BoardDoneHandler(conn net.Conn) (coordModels.ManagerAction, error) {
	return coordModels.BoardDoneAction{}, nil
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
