package tcp

type IncompatibleProtocolError struct {
	managerErrorMsg	string
}

func NewIncompatibleProtocolError(msg string) *IncompatibleProtocolError {
	return &IncompatibleProtocolError{
		managerErrorMsg: msg,
	}
}

func (e *IncompatibleProtocolError) Error() string {
	return e.managerErrorMsg
}

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
