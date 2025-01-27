package tcp

import (
	"net"

	"github.com/coding-kelps/gomocku/pkg/domain/listener/ports"
)

const (
	StartActionID 				byte	= 0x01
	TurnActionID 				byte	= 0x02
	BeginActionID 				byte	= 0x03
	BoardBeginActionID 			byte	= 0x04
	BoardTurnActionID 			byte	= 0x05
	BoardEndActionID 			byte	= 0x06
	InfoActionID 				byte	= 0x07
	EndActionID 				byte	= 0x08
	AboutActionID 				byte	= 0x09
	ReadyActionID 				byte	= 0x0A
	PlayActionID 				byte	= 0x0B
	PlayerDescriptionActionID 	byte	= 0x0C
	UnknownActionID				byte	= 0x0D
	ErrorActionID				byte	= 0x0E
	MessageActionID				byte	= 0x0F
	DebugActionID				byte	= 0x10
	SuggestionActionID			byte	= 0x11
)

type Tcp struct {
	listener 			net.Listener
	managerConnection	net.Conn

	ports.ManagerInterface
}

func NewTCP(localAddress string, managerAddress string) (*Tcp, error) {
	listener, err := net.Listen("tcp", localAddress)
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("tcp", managerAddress)
	if err != nil {
		return nil, err
	}

	return &Tcp{
		listener:			listener,
		managerConnection:	conn,
	}, nil
}

func (tcp *Tcp) Close() error {
	if tcp.managerConnection != nil {
		err := tcp.managerConnection.Close()
		if err != nil {
			return err
		}
	}

	if tcp.listener != nil {
		err := tcp.listener.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
