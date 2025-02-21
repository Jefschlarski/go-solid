package dto

import (
	"github.com/Jefschlarski/go-solid/internal/model"
)

// ReturnTodoDTO representa os dados de retorno de uma tarefa
// @Description DTO para retorno de uma tarefa
type ReturnTodoDTO struct {
	// ID único da tarefa
	// @Example 1
	ID int `json:"id" example:"1"`

	// Título da tarefa
	// @Example "Fazer compras"
	Title string `json:"title" example:"Fazer compras"`

	// Descrição detalhada da tarefa
	// @Example "Comprar leite, pão e ovos"
	Description string `json:"description" example:"Comprar leite, pão e ovos"`

	// Status da tarefa
	// @Example "PENDING"
	Status string `json:"status" example:"PENDING"`

	// Tempo gasto em minutos
	// @Example 30
	TimeSpent int64 `json:"time_spent" example:"30"`

	// Data e hora de criação da tarefa
	// @Example "10/02/2025 10:00"
	CreatedAt string `json:"created_at" example:"10/02/2025 10:00"`

	// Data e hora da última atualização
	// @Example "10/02/2025 10:00"
	UpdatedAt *string `json:"updated_at,omitempty" example:"10/02/2025 10:00"`

	// Data e hora da conclusão
	// @Example "10/02/2025 10:00"
	CompletedAt *string `json:"completed_at,omitempty" example:"10/02/2025 10:00"`
}

func (r *ReturnTodoDTO) FromModel(todo *model.Todo) {
	r.ID = todo.ID
	r.Title = todo.Title
	r.Description = todo.Description
	r.Status = todo.Status.String()
	r.TimeSpent = todo.TimeSpent
	r.CreatedAt = todo.CreatedAt.Format("02/01/2006 15:04")

	if todo.UpdatedAt != nil {
		formatted := todo.UpdatedAt.Format("02/01/2006 15:04")
		r.UpdatedAt = &formatted
	}
	if todo.CompletedAt != nil {
		formatted := todo.CompletedAt.Format("02/01/2006 15:04")
		r.CompletedAt = &formatted
	}
}
