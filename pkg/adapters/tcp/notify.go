package tcp

import (
	"fmt"

	ai_models "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func (tcp *Tcp) NotifyMove(p ai_models.Position) error {
	data := []byte("Hello, manager!")
	_, err := tcp.managerConnection.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return nil
}

func (tcp *Tcp) NotifyReadiness() error {
	data := []byte("Hello, manager!")
	_, err := tcp.managerConnection.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return nil
}

func (tcp *Tcp) NotifyUnknown(str string) error {
	data := []byte("Hello, manager!")
	_, err := tcp.managerConnection.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return nil
}

func (tcp *Tcp) NotifyError(str string) error {
	data := []byte("Hello, manager!")
	_, err := tcp.managerConnection.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return nil
}

func (tcp *Tcp) NotifyMessage(str string) error {
	data := []byte("Hello, manager!")
	_, err := tcp.managerConnection.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return nil
}

func (tcp *Tcp) NotifyDebug(str string) error {
	data := []byte("Hello, manager!")
	_, err := tcp.managerConnection.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return nil
}

func (tcp *Tcp) NotifySuggestion(p ai_models.Position) error {
	data := []byte("Hello, manager!")
	_, err := tcp.managerConnection.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	
	return nil
}


func (tcp *Tcp) NotifyMetadata(metadata map[string]string) error {
	data := []byte("Hello, manager!")
	_, err := tcp.managerConnection.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	
	return nil
}

