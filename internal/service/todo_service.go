package service

import (
	"errors"

	"github.com/Jefschlarski/go-solid/internal/model"
	"github.com/Jefschlarski/go-solid/internal/repository"
)

type ITodoService interface {
	CreateTodo(todo *model.Todo) error
	GetAllTodos() ([]model.Todo, error)
	UpdateTodoStatus(id int, newStatus model.TodoStatus) (*model.Todo, error)
	AddTimeSpent(id int, minutes int64) (*model.Todo, error)
}

type todoService struct {
	repo repository.ITodoRepository
}

func NewTodoService(repo repository.ITodoRepository) ITodoService {
	return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(todo *model.Todo) error {
	return s.repo.Create(todo)
}

func (s *todoService) GetAllTodos() ([]model.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) UpdateTodoStatus(id int, newStatus model.TodoStatus) (*model.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("tarefa não encontrada")
	}

	strategy := model.GetTodoState(todo.Status)
	if strategy == nil {
		return nil, errors.New("status atual inválido")
	}

	if err := strategy.ChangeStatus(todo, newStatus); err != nil {
		return nil, err
	}

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoService) AddTimeSpent(id int, minutes int64) (*model.Todo, error) {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if todo == nil {
		return nil, errors.New("tarefa não encontrada")
	}

	todoState := model.GetTodoState(todo.Status)
	if todoState == nil {
		return nil, errors.New("status atual inválido")
	}

	if err := todoState.AddTimeSpent(todo, minutes); err != nil {
		return nil, err
	}

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}
