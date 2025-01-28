package stdio

import (
	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (std *StdioManagerInterface) Listen(ch chan<-coordModels.ManagerAction) error {
	defer close(ch)

	for {
		if std.scanner.Scan() {
			input := std.scanner.Text()
			matched := false

			for _, p := range std.parsers {
				if p.regex.MatchString(input) {
					cmd, err := p.caller(input)
					if err != nil {
						return err
					}
					ch <- cmd
					matched = true
				}
			}

			if !matched {
				ch <- coordModels.UnknownAction{}
			}
		} else {
			break
		}
	}
 
	return nil
}
