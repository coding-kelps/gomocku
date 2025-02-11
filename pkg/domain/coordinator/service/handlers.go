package service

import (
	"fmt"

	aiModels "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (c *Coordinator) actionHandler(action models.ManagerAction) error {
	handlers := map[string]func(a models.ManagerAction)error{
		"START": 			c.startHandler,
		"RESTART":			c.restartHandler,
		"TURN": 			c.turnHandler,
		"BEGIN": 			c.beginHandler,
		"BOARD":			c.boardHandler,
		"INFO": 			c.infoHandler,
		"RESULT":			c.resultHandler,
		"END": 				c.endHandler,
		"ABOUT": 			c.aboutHandler,
	}

	actionType := action.ActionType()
	handler, ok := handlers[actionType]
    if !ok {
		c.unknownActionErrorHandler()

		return fmt.Errorf("received an unknown action from the manager")
	}

    c.logger.Debug().
        Str("action_type", action.ActionType()).
        Msg("manager action received")
	
    err := handler(action)
	if err != nil {
		return err
	}

    return nil
}

func (c *Coordinator) startHandler(a models.ManagerAction) error {
	start := a.(models.StartAction)

	err := c.ai.Init(start.Size)
	if err != nil {
		err2 := c.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = c.managerInterface.NotifyReadiness()
	if err != nil {
		return err
	}


	return nil
}

func (c *Coordinator) restartHandler(_ models.ManagerAction) error {
	err := c.ai.Reset()
	if err != nil {
		return err
	}

	err = c.managerInterface.NotifyReadiness()
	if err != nil {
		return err
	}

	return nil
}

func (c *Coordinator) turnHandler(a models.ManagerAction) error {
	turn := a.(models.TurnAction)

	err := c.ai.RegisterMove(turn.Position, aiModels.OpponentStone)
	if err != nil {
		err2 := c.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	pos, err := c.ai.PickMove()
	if err != nil {
		err2 := c.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = c.ai.RegisterMove(pos, aiModels.OwnStone)
	if err != nil {
		err2 := c.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = c.managerInterface.NotifyMove(pos)
	if err != nil {
		return err
	}

	return nil
}

func (c *Coordinator) beginHandler(_ models.ManagerAction) error {
	c.lock.Lock()
	pos, err := c.ai.PickMove()
	if err != nil {
		err2 := c.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = c.ai.RegisterMove(pos, aiModels.OwnStone)
	if err != nil {
		err2 := c.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.lock.Unlock()

	err = c.managerInterface.NotifyMove(pos)
	if err != nil {
		return err
	}
	
	return nil
}

func (c *Coordinator) boardHandler(a models.ManagerAction) error {
	board := a.(models.BoardAction)

	for _, t := range board.Turns {
		err := c.ai.RegisterMove(t.Position, aiModels.CellStatus(t.Player))
		if err != nil {
			err2 := c.managerInterface.NotifyError(err.Error())
			if err2 != nil {
				return err2
			}
			return nil
		}
	}

	pos, err := c.ai.PickMove()
	if err != nil {
		err2 := c.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = c.ai.RegisterMove(pos, aiModels.OwnStone)
	if err != nil {
		err2 := c.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = c.managerInterface.NotifyMove(pos)
	if err != nil {
		return err
	}

	return nil
}

func (c *Coordinator) resultHandler(a models.ManagerAction) error {
	result := a.(models.ResultAction)

	var resultStr string
	switch result.Result {
	case aiModels.Draw:
		resultStr = "draw"
	case aiModels.Win:
		resultStr = "win"
	case aiModels.Loose:
		resultStr = "loose"
	}

	c.logger.Info().
		Str("result", resultStr).
		Msg("game result")

	return nil
}

func (c *Coordinator) endHandler(_ models.ManagerAction) error {
	c.logger.Info().
		Msg("session termination requested by manager")

	c.endCh <- struct{}{}

	return nil
}

func (c *Coordinator) infoHandler(a models.ManagerAction) error {
	info := a.(models.InfoAction)

	c.logger.Info().
		Str("info", info.Str).
		Msg("manager info")

	return nil
}

func (c *Coordinator) aboutHandler(_ models.ManagerAction) error {
	err := c.managerInterface.NotifyMetadata(c.metadata)
	if err != nil {
		return err
	}

	return nil
}

func (c *Coordinator) unknownActionErrorHandler() error {
	err := c.managerInterface.NotifyUnknown()
	if err != nil {
		return err
	}

	return nil
}
