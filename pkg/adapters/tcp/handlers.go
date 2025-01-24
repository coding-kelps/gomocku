package tcp

import (
	"io"
	"net"
	"fmt"
	"encoding/binary"

	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
	ai_models "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func StartHandler(conn net.Conn) (models.ManagerCommand, error) {
	payload := make([]byte, 1)
	if _, err := io.ReadFull(conn, payload); err != nil {
		if err == io.EOF {
			fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
		} else {
			fmt.Printf("Error reading command ID from %s: %v", conn.RemoteAddr(), err)
		}

		return nil, err
	}

	return models.StartCommand{
		Size: payload[0],
	}, nil
}

func TurnHandler(conn net.Conn) (models.ManagerCommand, error) {
	payload := make([]byte, 2)
	if _, err := io.ReadFull(conn, payload); err != nil {
		if err == io.EOF {
			fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
		} else {
			fmt.Printf("Error reading command ID from %s: %v", conn.RemoteAddr(), err)
		}

		return nil, err
	}

	return models.TurnCommand{
		Position: ai_models.Position{
			X: payload[0],
			Y: payload[1],
		},
	}, nil
}

func BeginHandler(conn net.Conn) (models.ManagerCommand, error) {
	return models.BeginCommand{}, nil
}

func BoardHandler(conn net.Conn) (models.ManagerCommand, error) {
	return models.BoardCommand{}, nil
}

func BoardTurnHandler(conn net.Conn) (models.ManagerCommand, error) {
	var player ai_models.Player

	payload := make([]byte, 3)
	if _, err := io.ReadFull(conn, payload); err != nil {
		if err == io.EOF {
			fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
		} else {
			fmt.Printf("Error reading command ID from %s: %v", conn.RemoteAddr(), err)
		}

		return nil, err
	}

	if payload[2] == 0 {
		player = ai_models.Us
	} else {
		player = ai_models.Opponent
	}

	return models.BoardTurnCommand{
		Turn: ai_models.Turn{
			Position: ai_models.Position{
				X: payload[0],
				Y: payload[1],
			},
			Player: player,
		},
	}, nil
}

func BoardDoneHandler(conn net.Conn) (models.ManagerCommand, error) {
	return models.BoardDoneCommand{}, nil
}

func InfoHandler(conn net.Conn) (models.ManagerCommand, error) {
	payload := make([]byte, 4)
	if _, err := io.ReadFull(conn, payload); err != nil {
		if err == io.EOF {
			fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
		} else {
			fmt.Printf("Error reading command ID from %s: %v", conn.RemoteAddr(), err)
		}

		return nil, err
	}


	infoSize := binary.BigEndian.Uint32(payload[:])
	payload = make([]byte, infoSize)
	if _, err := io.ReadFull(conn, payload); err != nil {
		if err == io.EOF {
			fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
		} else {
			fmt.Printf("Error reading command ID from %s: %v", conn.RemoteAddr(), err)
		}

		return nil, err
	}

	return models.InfoCommand{
		Str: string(payload),
	}, nil
}

func EndHandler(conn net.Conn) (models.ManagerCommand, error) {
	return models.EndCommand{}, nil
}

func AboutHandler(conn net.Conn) (models.ManagerCommand, error) {
	return models.AboutCommand{}, nil
}
