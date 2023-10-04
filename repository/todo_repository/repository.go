package todo_repository

import (
	"final_project_1/entity"
	"final_project_1/pkg/errors_response"
)

type Repository interface {
	CreateTodo(todoPayload *entity.Todo) (*entity.Todo, errors_response.ErrorResponse)
	GetTodos() (*[]entity.Todo, errors_response.ErrorResponse)
	GetTodo(id int) (*entity.Todo, errors_response.ErrorResponse)
	UpdateTodo(id int, todoPayload *entity.Todo) (*entity.Todo, errors_response.ErrorResponse)
	DeleteTodo(id int) errors_response.ErrorResponse
}
