package dto

import "errors"

// CreateTodoDTO representa os dados necessários para criar uma nova tarefa
// @Description DTO para criação de uma tarefa
type CreateTodoDTO struct {
	// Título da tarefa
	// @Example "Fazer compras"
	Title string `json:"title" example:"Fazer compras" binding:"required"`

	// Descrição detalhada da tarefa
	// @Example "Comprar leite, pão e ovos"
	Description string `json:"description" example:"Comprar leite, pão e ovos" binding:"required"`
}

func (c *CreateTodoDTO) Validate() error {
	if c.Title == "" {
		return errors.New("título é obrigatório")
	}
	if c.Description == "" {
		return errors.New("descrição é obrigatória")
	}
	return nil
}
