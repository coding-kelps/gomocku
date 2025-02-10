package tcp

import (
	"fmt"
	"io"
	"net"
	"encoding/binary"

	"github.com/rs/zerolog"

	coordPorts "github.com/coding-kelps/gomocku/pkg/domain/coordinator/ports"
)

const (
	ProtocolVersion						string	= "0.2.0"

	// Actions that can be send from the manager to the player.

	ProtocolCompatibleManagerActionID	byte	= 0x00
	StartManagerActionID 				byte	= 0x01
	RestartManagerActionID				byte	= 0x02
	TurnManagerActionID 				byte	= 0x03
	BeginManagerActionID 				byte	= 0x04
	BoardManagerActionID 				byte	= 0x05
	InfoManagerActionID 				byte	= 0x06
	ResultManagerActionID				byte	= 0x07
	EndManagerActionID					byte	= 0x08
	AboutManagerActionID 				byte	= 0x09
	UnknownManagerActionID				byte	= 0x0A
	ErrorManagerActionID				byte	= 0x0B

	// Actions that can be send from the player to the manager.

	ProtocolVersionPlayerActionID		byte	= 0x0C
	ReadyPlayerActionID 				byte	= 0x0D
	PlayPlayerActionID					byte	= 0x0E
	PlayerMetadataActionID				byte	= 0x0F
	UnknownPlayerActionID				byte	= 0x10
	ErrorPlayerActionID					byte	= 0x11
	MessagePlayerActionID				byte	= 0x12
	DebugPlayerActionID					byte	= 0x13
	SuggestionPlayerActionID			byte	= 0x14
)

type TcpManagerInterface struct {
	conn	net.Conn
	logger	zerolog.Logger

	coordPorts.ManagerInterface
}

func NewTCPManagerInterface(conn net.Conn, logger zerolog.Logger) (*TcpManagerInterface, error) {
	tcp := TcpManagerInterface{
		conn:	conn,
		logger:	logger,
	}

	err := tcp.checkProtocolCompatibilty()
	if err != nil {
		return nil, err
	} else {
		return &tcp, nil
	}
}

func (tcp *TcpManagerInterface) checkProtocolCompatibilty() error {
	protocolVersionAsBytes := []byte(ProtocolVersion)
	protocolVersionAsBytesLenAsBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(protocolVersionAsBytesLenAsBytes, uint32(len(protocolVersionAsBytes)))

	data := append([]byte{ProtocolVersionPlayerActionID}, append(protocolVersionAsBytesLenAsBytes, []byte(ProtocolVersion)...)...)
	if _, err := tcp.conn.Write(data); err != nil {
		return err
	}

	payload := make([]byte, 1)
	if _, err := io.ReadFull(tcp.conn, payload); err != nil {
		return err
	}
	
	switch payload[0] {
	case ProtocolCompatibleManagerActionID:
		return nil
	case UnknownManagerActionID:
		return NewManagerActionError("manager does not know protocol compatibility check action")
	case ErrorManagerActionID:
		payload := make([]byte, 4)
		if _, err := io.ReadFull(tcp.conn, payload); err != nil {
			return err
		}
	
		msgSize := binary.BigEndian.Uint32(payload[:])
		payload = make([]byte, msgSize)
		if _, err := io.ReadFull(tcp.conn, payload); err != nil {
			return err
		}

		return NewIncompatibleProtocolError(string(payload))
	default:
		msg := fmt.Sprintf("unexpected manager action with ID 0x%X at protocol compatibility validation", payload[0])

		return NewManagerActionError(msg)
	}
}
