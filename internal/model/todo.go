package model

import "time"

const (
	TodoTableName = "todos"
)

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
	TimeSpent   int64      `json:"time_spent"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

func NewTodo(title, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
	}
}
