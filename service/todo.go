package service

import (
	"final_project_1/dto"
	"final_project_1/entity"
	"final_project_1/pkg/errors_response"
	"final_project_1/repository/todo_repository"
	"net/http"
	"strconv"
)

type todoService struct {
	TodoRepo todo_repository.Repository
}

type TodoService interface {
	CreateTodo(todoPayload *dto.NewTodoRequest) (*dto.NewTodoResponse, errors_response.ErrorResponse)
	GetTodos() (*dto.NewTodoResponse, errors_response.ErrorResponse)
	GetTodo(id int) (*dto.NewTodoResponse, errors_response.ErrorResponse)
	UpdateTodo(id int, todoPayload *dto.NewTodoRequest) (*dto.NewTodoResponse, errors_response.ErrorResponse)
	DeleteTodo(id int) (*dto.NewTodoResponse, errors_response.ErrorResponse)
}

func NewTodoService(todoRepo todo_repository.Repository) TodoService {
	return &todoService{TodoRepo: todoRepo}
}

func (ts *todoService) GetTodos() (*dto.NewTodoResponse, errors_response.ErrorResponse) {
	findAllTodo, err := ts.TodoRepo.GetTodos()
	if err != nil {
		return nil, err
	}
	response := dto.NewTodoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "found all todo list",
		Data:       findAllTodo,
	}
	return &response, nil
}

func (ts *todoService) GetTodo(id int) (*dto.NewTodoResponse, errors_response.ErrorResponse) {
	findTodo, err := ts.TodoRepo.GetTodo(id)
	if err != nil {
		return nil, err
	}
	msg := "find todo list with id: " + strconv.Itoa(id)
	response := dto.NewTodoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    msg,
		Data:       findTodo,
	}
	return &response, nil
}

func (ts *todoService) CreateTodo(todoPayload *dto.NewTodoRequest) (*dto.NewTodoResponse, errors_response.ErrorResponse) {
	todo := entity.Todo{
		Title: todoPayload.Title,
		Done:  todoPayload.Done,
	}

	createdTodo, err := ts.TodoRepo.CreateTodo(&todo)
	if err != nil {
		return nil, err
	}

	response := dto.NewTodoResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "todo successfully created",
		Data:       createdTodo,
	}

	return &response, nil
}

func (ts *todoService) UpdateTodo(id int, todoPayload *dto.NewTodoRequest) (*dto.NewTodoResponse, errors_response.ErrorResponse) {
	todo := entity.Todo{
		Title: todoPayload.Title,
		Done:  todoPayload.Done,
	}

	updatedTodo, err := ts.TodoRepo.UpdateTodo(id, &todo)
	if err != nil {
		return nil, err
	}

	msg := "todo with id: " + strconv.Itoa(id) + " successfully updated"
	response := dto.NewTodoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    msg,
		Data:       updatedTodo,
	}

	return &response, nil
}

func (ts *todoService) DeleteTodo(id int) (*dto.NewTodoResponse, errors_response.ErrorResponse) {
	err := ts.TodoRepo.DeleteTodo(id)
	if err != nil {
		return nil, err
	}

	msg := "todo with id: " + strconv.Itoa(id) + " successfully deleted"
	response := dto.NewTodoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    msg,
		Data:       nil,
	}
	return &response, nil
}
