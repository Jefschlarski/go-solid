package model

type ITodoState interface {
	ChangeStatus(todo *Todo, newStatus TodoStatus) error
	AddTimeSpent(todo *Todo, minutes int64) error
}

func GetTodoState(status TodoStatus) ITodoState {
	switch status {
	case StatusPending:
		return NewPendingState()
	case StatusInProgress:
		return NewInProgressState()
	case StatusPaused:
		return NewPausedState()
	case StatusCompleted:
		return NewCompletedState()
	case StatusCanceled:
		return NewCanceledState()
	default:
		return nil
	}
}
