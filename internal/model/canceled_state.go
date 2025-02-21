package model

import "errors"

type canceledState struct{}

func NewCanceledState() ITodoState {
	return &canceledState{}
}

func (s *canceledState) ChangeStatus(todo *Todo, newStatus TodoStatus) error {
	return errors.New("não é possível alterar o status de uma tarefa cancelada")
}

func (s *canceledState) AddTimeSpent(todo *Todo, minutes int64) error {
	return errors.New("não é possível adicionar tempo em uma tarefa cancelada")
}
