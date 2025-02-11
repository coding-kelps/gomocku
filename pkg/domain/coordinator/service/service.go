package service

import (
    "fmt"
	"sync"

	"github.com/rs/zerolog"

	aiPorts "github.com/coding-kelps/gomocku/pkg/domain/ai/ports"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/ports"
)

type Coordinator struct {
	managerInterface 	ports.ManagerInterface
	metadata 			map[string]string
	ai 					aiPorts.AI
	lock				*sync.RWMutex
    logger              zerolog.Logger
    endCh               chan struct{}

	ports.Coordinator
}

func (c *Coordinator) Serve() error {
	actionsCh := make(chan models.ManagerAction, 10)
	errorsCh := make(chan error)

	go c.managerInterface.Listen(actionsCh, errorsCh)

	for {
        select {  
        case _, ok := <-c.endCh:
            if !ok {
                c.logger.Warn().
                    Msg("end channel closed")

                    return fmt.Errorf("end channel unexpectedly closed")
            }

            return nil
        
        case action, ok := <-actionsCh:
            if !ok {
                select {
                case <-c.endCh:
                    return nil
                default:
                    return fmt.Errorf("connection to manager unexpectedly interrupted")
                }
            }
            err := c.actionHandler(action)
            if err != nil {
                return err
            }

        case err, ok := <-errorsCh:
            if !ok {
                return fmt.Errorf("connection to manager unexpectedly interrupted")

            }

            c.logger.Error().
                Err(err).
                Msg("coordinator's interface error")

            return err
        }
    }
}

func NewCoordinator(managerInterface ports.ManagerInterface, ai aiPorts.AI, logger zerolog.Logger) ports.Coordinator {
	return &Coordinator{
		managerInterface: managerInterface,
		ai:         ai,
		lock:       &sync.RWMutex{},
		metadata:   map[string]string{
			"name":    "gomocku",
			"version": "0.1",
			"author":  "Coding Kelps",
			"desc":    "A mock AI for manager testing",
		},
        logger:     logger,
        endCh:      make(chan struct{}, 1),
	}
}
