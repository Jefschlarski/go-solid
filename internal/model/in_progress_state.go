package model

import (
	"errors"
	"time"
)

type inProgressState struct{}

func NewInProgressState() ITodoState {
	return &inProgressState{}
}

func (s *inProgressState) ChangeStatus(todo *Todo, newStatus TodoStatus) error {
	if newStatus != StatusCompleted && newStatus != StatusCanceled && newStatus != StatusPaused {
		return errors.New("tarefa em progresso só pode ser completada, cancelada ou pausada")
	}

	now := time.Now()
	todo.Status = newStatus
	todo.UpdatedAt = &now

	if newStatus == StatusCompleted {
		todo.CompletedAt = &now
	}
	return nil
}

func (s *inProgressState) AddTimeSpent(todo *Todo, minutes int64) error {
	if minutes <= 0 {
		return errors.New("tempo deve ser maior que zero")
	}
	todo.TimeSpent += minutes
	now := time.Now()
	todo.UpdatedAt = &now
	return nil
}
