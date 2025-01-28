package tcp

import (
	"net"

	coordPorts "github.com/coding-kelps/gomocku/pkg/domain/coordinator/ports"
)

const (
	StartManagerActionID 				byte	= 0x01
	TurnManagerActionID 				byte	= 0x02
	BeginManagerActionID 				byte	= 0x03
	BoardBeginManagerActionID 			byte	= 0x04
	BoardTurnManagerActionID 			byte	= 0x05
	BoardEndManagerActionID 			byte	= 0x06
	InfoManagerActionID 				byte	= 0x07
	EndManagerActionID					byte	= 0x08
	AboutManagerActionID 				byte	= 0x09
	ReadyPlayerActionID 				byte	= 0x0A
	PlayPlayerActionID					byte	= 0x0B
	PlayerDescriptionPlayerActionID		byte	= 0x0C
	UnknownPlayerActionID				byte	= 0x0D
	ErrorPlayerActionID					byte	= 0x0E
	MessagePlayerActionID				byte	= 0x0F
	DebugPlayerActionID					byte	= 0x10
	SuggestionPlayerActionID			byte	= 0x11
)

type TcpManagerInterface struct {
	conn	net.Conn

	coordPorts.ManagerInterface
}

func NewTCPManagerInterface(conn net.Conn) *TcpManagerInterface {
	return &TcpManagerInterface{
		conn:	conn,
	}
}
