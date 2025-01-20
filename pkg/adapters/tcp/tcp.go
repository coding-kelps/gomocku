package tcp

import (
	"net"
)

type Tcp struct {
	listener 			net.Listener
	managerConnection	net.Conn
}

type CreateTcpConfiguration struct {
	localAddress 	string
	managerAddress 	string
}

func NewTCP(cfg CreateTcpConfiguration) (*Tcp, error) {
	listener, err := net.Listen("tcp", cfg.localAddress)
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("tcp", cfg.managerAddress)
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
