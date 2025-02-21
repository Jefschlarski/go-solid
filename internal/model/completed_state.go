package model

import "errors"

type completedState struct{}

func NewCompletedState() ITodoState {
	return &completedState{}
}

func (s *completedState) ChangeStatus(todo *Todo, newStatus TodoStatus) error {
	return errors.New("não é possível alterar o status de uma tarefa concluída")
}

func (s *completedState) AddTimeSpent(todo *Todo, minutes int64) error {
	return errors.New("não é possível adicionar tempo em uma tarefa concluída")
}
