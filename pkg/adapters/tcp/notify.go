package tcp

import (
	"encoding/binary"
	"fmt"
	"strings"

	aiModels "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func (tcp *TcpManagerInterface) NotifyReadiness() error {
	data := []byte{ReadyPlayerActionID}
	_, err := tcp.conn.Write(data)
	
	return err
}

func (tcp *TcpManagerInterface) NotifyMove(p aiModels.Position) error {	
	data := []byte{PlayPlayerActionID, p.X, p.Y}
	_, err := tcp.conn.Write(data)

	return err
}

func (tcp *TcpManagerInterface) NotifyMetadata(metadata map[string]string) error {
	infos := make([]string, 0, len(metadata))
	for k, v := range metadata {
		infos = append(
			infos,
			fmt.Sprintf("%s=\"%s\"", k, v),
		)
	}

	ascii_metadata := []byte(strings.Join(infos, ", "))
	metadata_len := make([]byte, 4)
    binary.BigEndian.PutUint32(metadata_len, uint32(len(ascii_metadata)))

	data := append([]byte{PlayerMetadataActionID}, append(metadata_len, ascii_metadata...)...)
	_, err := tcp.conn.Write(data)
	
	return err
}

func (tcp *TcpManagerInterface) NotifyUnknown() error {
	data := []byte{UnknownPlayerActionID}
	_, err := tcp.conn.Write(data)
	
	return err
}

func (tcp *TcpManagerInterface) NotifyError(str string) error {
	str_len := make([]byte, 4)
    binary.BigEndian.PutUint32(str_len, uint32(len(str)))

	data := append([]byte{ErrorPlayerActionID}, append(str_len, []byte(str)...)...)
	_, err := tcp.conn.Write(data)
	
	return err
}

func (tcp *TcpManagerInterface) NotifyMessage(str string) error {
	str_len := make([]byte, 4)
    binary.BigEndian.PutUint32(str_len, uint32(len(str)))

	data := append([]byte{MessagePlayerActionID}, append(str_len, []byte(str)...)...)
	_, err := tcp.conn.Write(data)
	
	return err
}

func (tcp *TcpManagerInterface) NotifyDebug(str string) error {
	str_len := make([]byte, 4)
    binary.BigEndian.PutUint32(str_len, uint32(len(str)))

	data := append([]byte{DebugPlayerActionID}, append(str_len, []byte(str)...)...)
	_, err := tcp.conn.Write(data)
	
	return err
}

func (tcp *TcpManagerInterface) NotifySuggestion(p aiModels.Position) error {
	data := []byte{SuggestionPlayerActionID, p.X, p.Y}
	_, err := tcp.conn.Write(data)
	
	return err
}
