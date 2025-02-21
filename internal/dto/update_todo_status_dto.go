package dto

import "errors"

// UpdateTodoStatusDTO representa os dados para atualização de status
// @Description DTO para atualização de status de uma tarefa
type UpdateTodoStatusDTO struct {
	// Novo status da tarefa (0: PENDING, 1: IN_PROGRESS, 2: PAUSED, 3: COMPLETED, 4: CANCELED)
	// @Example 1
	Status int `json:"status" example:"1"`
}

func (u *UpdateTodoStatusDTO) Validate() error {
	if u.Status < 0 || u.Status > 4 {
		return errors.New("status inválido")
	}
	return nil
}
