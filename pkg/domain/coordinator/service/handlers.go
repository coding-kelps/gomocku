package service

import (
	aiModels "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator/models"
)

func (c *Coordinator) startHandler(a models.StartAction) error {
	err := c.ai.Init(a.Size)
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

func (c *Coordinator) restartHandler() error {

	return nil
}

func (c *Coordinator) turnHandler(a models.TurnAction) error {
	err := c.ai.RegisterMove(a.Position, aiModels.OpponentStone)
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

func (c *Coordinator) beginHandler() error {
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

func (c *Coordinator) boardHandler(a models.BoardAction) error {
	for _, t := range a.Turns {
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

func (c *Coordinator) resultHandler(a models.ResultAction) error {
	var resultStr string
	
	switch a.Result {
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

func (c *Coordinator) infoHandler(a models.InfoAction) error {
	c.logger.Info().
		Str("info", a.Str).
		Msg("manager info")

	return nil
}

func (c *Coordinator) aboutHandler() error {
	err := c.managerInterface.NotifyMetadata(c.metadata)
	if err != nil {
		return err
	}

	return nil
}

func (c *Coordinator) UnknownHandler() error {
	err := c.managerInterface.NotifyUnknown()
	if err != nil {
		return err
	}

	return nil
}
