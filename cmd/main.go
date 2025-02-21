package main

import (
	"database/sql"
	"net/http"

	_ "github.com/Jefschlarski/go-solid/docs"
	"github.com/Jefschlarski/go-solid/internal/config"
	"github.com/Jefschlarski/go-solid/internal/handler"
	"github.com/Jefschlarski/go-solid/internal/repository"
	"github.com/Jefschlarski/go-solid/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title ToDo List API
// @version 1.0
// @description API para gerenciar uma lista de tarefas.
// @host localhost:8080
// @BasePath /api
func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Configuração do banco de dados SQLite
	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Criar tabela se não existir
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        status INTEGER NOT NULL,
        time_spent INTEGER DEFAULT 0,
        created_at DATETIME NOT NULL,
        updated_at DATETIME,
        completed_at DATETIME
    )`)
	if err != nil {
		panic(err)
	}

	// Inicializar repositório, serviço e handler
	repo := repository.NewTodoRepository(db)
	service := service.NewTodoService(repo)
	handler := handler.NewTodoHandler(service)

	// Configuração do Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Rotas
	e.POST("/api/todos", handler.CreateTodo)
	e.GET("/api/todos", handler.GetAllTodos)
	e.PATCH("/api/todos/:id/status", handler.UpdateTodoStatus)
	e.PATCH("/api/todos/:id/time", handler.AddTimeSpent)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "UP",
		})
	})

	// Iniciar servidor com porta da configuração
	e.Logger.Fatal(e.Start(":" + cfg.ServerPort))
}
