package dto

import (
	"errors"
)

// AddTimeSpentDTO representa os dados para adicionar tempo gasto
// @Description DTO para adicionar tempo gasto em uma tarefa
type AddTimeSpentDTO struct {
	// Tempo gasto em minutos
	// @Example 30
	Minutes int64 `json:"minutes" example:"30"`
}

func (a *AddTimeSpentDTO) Validate() error {
	if a.Minutes <= 0 {
		return errors.New("tempo deve ser maior que zero")
	}
	return nil
}
