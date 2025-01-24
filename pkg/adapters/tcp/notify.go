package tcp

import (
	"encoding/binary"
	"strings"
	"fmt"

	ai_models "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func (tcp *Tcp) NotifyMove(p ai_models.Position) error {
	data := []byte{0x00, p.X, p.Y}
	_, err := tcp.managerConnection.Write(data)

	return err
}

func (tcp *Tcp) NotifyReadiness() error {
	data := []byte{0x00}
	_, err := tcp.managerConnection.Write(data)
	
	return err
}

func (tcp *Tcp) NotifyUnknown() error {
	data := []byte{0x00}
	_, err := tcp.managerConnection.Write(data)
	
	return err
}

func (tcp *Tcp) NotifyError(str string) error {
	str_len := make([]byte, 4)
    binary.BigEndian.PutUint32(str_len, uint32(len(str)))

	data := append([]byte{0x00}, append(str_len, []byte(str)...)...)
	_, err := tcp.managerConnection.Write(data)
	
	return err
}

func (tcp *Tcp) NotifyMessage(str string) error {
	str_len := make([]byte, 4)
    binary.BigEndian.PutUint32(str_len, uint32(len(str)))

	data := append([]byte{0x00}, append(str_len, []byte(str)...)...)
	_, err := tcp.managerConnection.Write(data)
	
	return err
}

func (tcp *Tcp) NotifyDebug(str string) error {
	str_len := make([]byte, 4)
    binary.BigEndian.PutUint32(str_len, uint32(len(str)))

	data := append([]byte{0x00}, append(str_len, []byte(str)...)...)
	_, err := tcp.managerConnection.Write(data)
	
	return err
}

func (tcp *Tcp) NotifySuggestion(p ai_models.Position) error {
	data := []byte{0x00, p.X, p.Y}
	_, err := tcp.managerConnection.Write(data)
	
	return err
}

func (tcp *Tcp) NotifyMetadata(metadata map[string]string) error {
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

	data := append([]byte{0x00}, append(metadata_len, ascii_metadata...)...)
	_, err := tcp.managerConnection.Write(data)
	
	return err
}
