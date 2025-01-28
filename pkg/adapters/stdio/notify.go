package stdio

import (
	"fmt"
	"strings"

	aiModels "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func (std *StdioManagerInterface) NotifyMove(p aiModels.Position) error {
	_, err := fmt.Printf("%d,%d\n", p.X, p.Y)
	if err != nil {
		return err
	}

	return nil
}

func (std *StdioManagerInterface) NotifyReadiness() error {
	_, err := fmt.Printf("OK\n")
	if err != nil {
		return err
	}

	return nil
}

func (std *StdioManagerInterface) NotifyUnknown() error {
	_, err := fmt.Printf("UNKNOWN unknown command\n")
	if err != nil {
		return err
	}

	return nil
}

func (std *StdioManagerInterface) NotifyError(str string) error {
	_, err := fmt.Printf("ERROR %s\n", str)
	if err != nil {
		return err
	}

	return nil
}

func (std *StdioManagerInterface) NotifyMessage(str string) error {
	_, err := fmt.Printf("MESSAGE %s\n", str)
	if err != nil {
		return err
	}

	return nil
}


func (std *StdioManagerInterface) NotifyDebug(str string) error {
	_, err := fmt.Printf("DEBUG %s\n", str)
	if err != nil {
		return err
	}

	return nil
}

func (std *StdioManagerInterface) NotifySuggestion(p aiModels.Position) error {
	_, err := fmt.Printf("SUGGEST %d,%d\n", p.X, p.Y)
	if err != nil {
		return err
	}

	return nil
}

func (std *StdioManagerInterface) NotifyMetadata(metadata map[string]string) error {
	infos := make([]string, 0, len(metadata))
	for k, v := range metadata {
		infos = append(
			infos,
			fmt.Sprintf("%s=\"%s\"", k, v),
		)
	}

	_, err := fmt.Println(strings.Join(infos, ", "))
	if err != nil {
		return err
	}

	return nil
}
