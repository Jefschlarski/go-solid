package handler

import (
	"net/http"
	"strconv"

	"github.com/Jefschlarski/go-solid/internal/dto"
	"github.com/Jefschlarski/go-solid/internal/model"
	"github.com/Jefschlarski/go-solid/internal/service"
	"github.com/labstack/echo/v4"
)

type ITodoHandler interface {
	CreateTodo(c echo.Context) error
	GetAllTodos(c echo.Context) error
	UpdateTodoStatus(c echo.Context) error
	AddTimeSpent(c echo.Context) error
}

type todoHandler struct {
	service service.ITodoService
}

func NewTodoHandler(service service.ITodoService) ITodoHandler {
	return &todoHandler{service: service}
}

// @Summary Criar uma nova tarefa
// @Description Cria uma nova tarefa na lista
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body dto.CreateTodoDTO true "Dados da tarefa"
// @Success 201 {object} dto.CreateTodoDTO
// @Failure 400 {string} string "Payload inválido"
// @Failure 500 {string} string "Erro ao criar tarefa"
// @Router /todos [post]
func (h *todoHandler) CreateTodo(c echo.Context) error {
	var todo dto.CreateTodoDTO
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, "Payload inválido")
	}

	if err := todo.Validate(); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	todoModel := model.NewTodo(todo.Title, todo.Description)

	if err := h.service.CreateTodo(todoModel); err != nil {
		return c.JSON(http.StatusInternalServerError, "Erro ao criar tarefa")
	}
	return c.JSON(http.StatusCreated, todo)
}

// @Summary Listar todas as tarefas
// @Description Retorna todas as tarefas cadastradas
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} dto.ReturnTodoDTO
// @Failure 500 {string} string "Erro ao buscar tarefas"
// @Router /todos [get]
func (h *todoHandler) GetAllTodos(c echo.Context) error {
	todos, err := h.service.GetAllTodos()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Erro ao buscar tarefas")
	}

	var todosDTO []dto.ReturnTodoDTO

	for _, todo := range todos {
		var todoDTO dto.ReturnTodoDTO
		todoDTO.FromModel(&todo)
		todosDTO = append(todosDTO, todoDTO)
	}

	return c.JSON(http.StatusOK, todosDTO)
}

// @Summary Atualizar status da tarefa
// @Description Atualiza o status de uma tarefa
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "ID da tarefa"
// @Param status body dto.UpdateTodoStatusDTO true "Novo status"
// @Success 200 {object} dto.ReturnTodoDTO
// @Failure 400 {string} string "Payload inválido"
// @Failure 404 {string} string "Tarefa não encontrada"
// @Router /todos/{id}/status [patch]
func (h *todoHandler) UpdateTodoStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID inválido")
	}

	var statusDTO dto.UpdateTodoStatusDTO
	if err := c.Bind(&statusDTO); err != nil {
		return c.JSON(http.StatusBadRequest, "Payload inválido")
	}

	if err := statusDTO.Validate(); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	todo, err := h.service.UpdateTodoStatus(id, model.TodoStatus(statusDTO.Status))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var todoDTO dto.ReturnTodoDTO
	todoDTO.FromModel(todo)
	return c.JSON(http.StatusOK, todoDTO)
}

// @Summary Adicionar tempo gasto na tarefa
// @Description Adiciona tempo gasto em uma tarefa
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "ID da tarefa"
// @Param time body dto.AddTimeSpentDTO true "Tempo gasto"
// @Success 200 {object} dto.ReturnTodoDTO
// @Failure 400 {string} string "Payload inválido"
// @Failure 404 {string} string "Tarefa não encontrada"
// @Router /todos/{id}/time [patch]
func (h *todoHandler) AddTimeSpent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "ID inválido")
	}

	var timeDTO dto.AddTimeSpentDTO
	if err := c.Bind(&timeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, "Payload inválido")
	}

	if err := timeDTO.Validate(); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	todo, err := h.service.AddTimeSpent(id, timeDTO.Minutes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var todoDTO dto.ReturnTodoDTO
	todoDTO.FromModel(todo)
	return c.JSON(http.StatusOK, todoDTO)
}
