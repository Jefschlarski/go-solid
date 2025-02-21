package repository

import (
	"database/sql"

	"github.com/Jefschlarski/go-solid/internal/model"
)

type ITodoRepository interface {
	Create(todo *model.Todo) error
	GetAll() ([]model.Todo, error)
	GetByID(id int) (*model.Todo, error)
	Update(todo *model.Todo) error
}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) ITodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	query := `INSERT INTO todos (title, description, status, time_spent, created_at) 
             VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, todo.Title, todo.Description, todo.Status, todo.TimeSpent, todo.CreatedAt)
	return err
}

func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	query := `SELECT id, title, description, status, time_spent, created_at, 
             updated_at, completed_at FROM todos`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Description, &todo.Status,
			&todo.TimeSpent, &todo.CreatedAt, &todo.UpdatedAt,
			&todo.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) GetByID(id int) (*model.Todo, error) {
	query := `SELECT id, title, description, status, time_spent, created_at, 
             updated_at, completed_at FROM todos WHERE id = ?`

	todo := &model.Todo{}
	err := r.db.QueryRow(query, id).Scan(
		&todo.ID, &todo.Title, &todo.Description, &todo.Status,
		&todo.TimeSpent, &todo.CreatedAt, &todo.UpdatedAt,
		&todo.CompletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) Update(todo *model.Todo) error {
	query := `UPDATE todos SET 
             title = ?, description = ?, status = ?, time_spent = ?,
             updated_at = ?, completed_at = ? WHERE id = ?`

	_, err := r.db.Exec(query,
		todo.Title, todo.Description, todo.Status, todo.TimeSpent,
		todo.UpdatedAt, todo.CompletedAt, todo.ID,
	)
	return err
}
