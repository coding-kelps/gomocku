package service

import (
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

	ports.Coordinator
}

func (c *Coordinator) Serve() error {
	actionsCh := make(chan models.ManagerAction, 10)
	errorsCh := make(chan error)

	go c.managerInterface.Listen(actionsCh, errorsCh)

	for {
        select {
        case cmd, ok := <-actionsCh:
            if !ok {
                c.logger.Warn().
                    Msg("action channel closed")

                return nil
            }

            switch cmd.ActionType() {
            case "start":
                c.logger.Debug().
                    Str("action_type", cmd.ActionType()).
                    Msg("manager action received")
                
                c.startHandler(cmd.(models.StartAction))
            case "turn":
                turn := cmd.(models.TurnAction)

                c.logger.Debug().
                    Str("action_type", cmd.ActionType()).
                    Uint8("turn_position_x", turn.Position.X).
                    Msg("manager action received")
                
                c.turnHandler(turn)
            case "begin":
                c.logger.Debug().
                    Str("action_type", cmd.ActionType()).
                    Msg("manager action received")
                
                c.beginHandler()
            case "board":
                board := cmd.(models.BoardAction)

                c.boardHandler(board)
            case "end":
                c.logger.Debug().
                    Str("action_type", cmd.ActionType()).
                    Msg("manager action received")

                c.logger.Info().
                    Msg("game ending requested by manager")
                
                return nil
            case "info":
                info := cmd.(models.InfoAction)

                c.infoHandler(info)
            case "about":
                c.logger.Debug().
                    Str("action_type", cmd.ActionType()).
                    Msg("manager action received")
                
                c.aboutHandler()
            case "unknown":
                c.logger.Debug().
                    Str("action_type", cmd.ActionType()).
                    Msg("manager action received")
                
                c.UnknownHandler()
            }

        case err, ok := <-errorsCh:
            if !ok {
                c.logger.Warn().
                    Msg("error channel closed")

                return nil
            }

            c.logger.Error().
                Err(err).
                Msg("coordinator's listener error")

            return err
        }
    }
}

func NewCoordinator(managerInterface ports.ManagerInterface, ai aiPorts.AI, logger zerolog.Logger) ports.Coordinator {
	return &Coordinator{
		managerInterface: managerInterface,
		ai: ai,
		lock: &sync.RWMutex{},
		metadata: map[string]string{
			"name":    "gomocku",
			"version": "0.1",
			"author":  "Coding Kelps",
			"desc":    "A mock AI for manager testing",
		},
        logger: logger,
	}
}
