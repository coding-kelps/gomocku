package tcp

import (
	"net"

	"github.com/coding-kelps/gomocku/pkg/domain/listener/ports"
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
