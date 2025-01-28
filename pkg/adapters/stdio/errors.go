package stdio

type ManagerActionError struct {
	msg	string
}

func NewManagerActionError(msg string) *ManagerActionError {
	return &ManagerActionError{
		msg: msg,
	}
}

func (e *ManagerActionError) Error() string {
	return e.msg
}
