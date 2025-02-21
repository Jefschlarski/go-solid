package model

import (
	"errors"
	"time"
)

type pausedState struct{}

func NewPausedState() ITodoState {
	return &pausedState{}
}

func (s *pausedState) ChangeStatus(todo *Todo, newStatus TodoStatus) error {
	if newStatus != StatusInProgress && newStatus != StatusCompleted && newStatus != StatusCanceled {
		return errors.New("tarefa pausada só pode ser retomada, completada ou cancelada")
	}

	now := time.Now()
	todo.Status = newStatus
	todo.UpdatedAt = &now

	if newStatus == StatusCompleted {
		todo.CompletedAt = &now
	}
	return nil
}

func (s *pausedState) AddTimeSpent(todo *Todo, minutes int64) error {
	return errors.New("não é possível adicionar tempo em uma tarefa pausada")
}
