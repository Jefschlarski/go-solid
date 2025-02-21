package model

import (
	"errors"
	"time"
)

type pendingState struct{}

func NewPendingState() ITodoState {
	return &pendingState{}
}

func (s *pendingState) ChangeStatus(todo *Todo, newStatus TodoStatus) error {
	if newStatus != StatusInProgress && newStatus != StatusCanceled {
		return errors.New("tarefa pendente só pode ser iniciada ou cancelada")
	}

	now := time.Now()
	todo.Status = newStatus
	todo.UpdatedAt = &now

	return nil
}

func (s *pendingState) AddTimeSpent(todo *Todo, minutes int64) error {
	return errors.New("não é possível adicionar tempo em uma tarefa pendente")
}
