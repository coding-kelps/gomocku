package stdio

import (
	"fmt"

	coordModels "github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (std *StdioManagerInterface) Listen(actionsCh chan<-coordModels.ManagerAction, errorsCh chan<-error) {
	defer close(actionsCh)

	for {
		if std.scanner.Scan() {
			input := std.scanner.Text()
			matched := false

			for _, p := range std.parsers {
				if p.regex.MatchString(input) {
					cmd, err := p.caller(input)
					if err != nil {
						errorsCh <- err

						return
					}
					actionsCh <- cmd
					matched = true
				}
			}

			if !matched {
				msg := fmt.Sprintf("unknown manager action: %s", input)

				errorsCh <- NewManagerActionError(msg)				
				return
			}
		} else {
			return
		}
	} 
}
