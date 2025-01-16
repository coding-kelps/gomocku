package service

import (
	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
	ai_models "github.com/coding-kelps/gomocku/pkg/domain/ai/models"
)

func (l Listener) startHandler(c models.StartCommand) error {
	l.lock.Lock()
	err := l.ai.Init(c.Size)
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}
	l.lock.Unlock()

	err = l.managerInterface.NotifyReadiness()
	if err != nil {
		return err
	}

	return nil
}

func (l Listener) turnHandler(c models.TurnCommand) error {
	l.lock.Lock()
	err := l.ai.RegisterMove(c.Position, ai_models.OpponentStone)
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	pos, err := l.ai.PickMove()
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = l.ai.RegisterMove(pos, ai_models.OwnStone)
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}
	l.lock.Unlock()

	err = l.managerInterface.NotifyMove(pos)
	if err != nil {
		return err
	}

	return nil
}

func (l Listener) beginHandler() error {
	l.lock.Lock()
	pos, err := l.ai.PickMove()
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = l.ai.RegisterMove(pos, ai_models.OwnStone)
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}
	l.lock.Unlock()

	err = l.managerInterface.NotifyMove(pos)
	if err != nil {
		return err
	}
	
	return nil
}

func (l Listener) boardHandler() error {
	return nil
}

func (l Listener) boardTurnHandler(c models.BoardTurnCommand) error {
	l.lock.Lock()
	err := l.ai.RegisterMove(c.Turn.Position, ai_models.CellStatus(c.Turn.Player))
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}
	l.lock.Unlock()

	return nil
}

func (l Listener) boardDoneHandler() error {
	l.lock.Lock()
	pos, err := l.ai.PickMove()
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}

	err = l.ai.RegisterMove(pos, ai_models.OwnStone)
	if err != nil {
		err2 := l.managerInterface.NotifyError(err.Error())
		if err2 != nil {
			return err2
		}
		return nil
	}
	l.lock.Unlock()

	err = l.managerInterface.NotifyMove(pos)
	if err != nil {
		return err
	}

	return nil
}

func (l Listener) infoHandler(_ models.InfoCommand) error {
	return nil
}

func (l Listener) aboutHandler() error {
	err := l.managerInterface.NotifyMetadata(l.metadata)
	if err != nil {
		return err
	}

	return nil
}

func (l Listener) UnknownHandler() error {
	err := l.managerInterface.NotifyUnknown()
	if err != nil {
		return err
	}

	return nil
}
