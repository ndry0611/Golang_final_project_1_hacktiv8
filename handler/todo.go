package handler

import (
	"final_project_1/dto"
	"final_project_1/pkg/errors_response"
	"final_project_1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) todoHandler {
	return todoHandler{todoService: todoService}
}

// GetTodos godoc
// @Tags Todos
// @Description Get all todos
// @ID get-all-todo
// @Accept json
// @Produce json
// @Success 200 {object} dto.NewTodoResponse
// @Router /todos [get]
func (th *todoHandler) GetTodos(ctx *gin.Context) {
	response, err := th.todoService.GetTodos()
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetTodo godoc
// @Tags Todos
// @Description Get todo
// @ID get-todo
// @Accept json
// @Produce json
// @Param todoId path int true "ID of the todo"
// @Success 200 {object} dto.NewTodoResponse
// @Router /todos/{todoId} [get]
func (th *todoHandler) GetTodo(ctx *gin.Context) {
	params, _ := strconv.Atoi(ctx.Param("todoId"))

	response, err := th.todoService.GetTodo(params)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// CreateTodo godoc
// @Tags Todos
// @Description Create a todo list
// @ID create-todo
// @Accept json
// @Produce json
// @Param Todo body dto.NewTodoRequest true "Create a todo request body"
// @Success 201 {object} dto.NewTodoResponse
// @Router /todos [post]
func (th *todoHandler) CreateTodo(ctx *gin.Context) {
	var todoRequest dto.NewTodoRequest

	if err := ctx.ShouldBindJSON(&todoRequest); err != nil {
		errBind := errors_response.NewUnprocessableEntityResponse("invalid request body")
		ctx.JSON(errBind.Status(), errBind)
		return
	}

	response, err := th.todoService.CreateTodo(&todoRequest)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// UpdateTodo godoc
// @Tags Todos
// @Description Update a todo list
// @ID update-todo
// @Accept json
// @Produce json
// @Param todoId path int true "ID of the todo"
// @Param Todo body dto.NewTodoRequest true "Update todo request body"
// @Success 200 {object} dto.NewTodoResponse
// @Router /todos/{todoId} [put]
func (th *todoHandler) UpdateTodo(ctx *gin.Context) {
	var todoRequest dto.NewTodoRequest

	if err := ctx.ShouldBindJSON(&todoRequest); err != nil {
		errBind := errors_response.NewUnprocessableEntityResponse("invalid request body")
		ctx.JSON(errBind.Status(), errBind)
		return
	}

	params, _ := strconv.Atoi(ctx.Param("todoId"))
	response, err := th.todoService.UpdateTodo(params, &todoRequest)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// DeleteTodo godoc
// @Tags Todos
// @Description Delete a todo list
// @ID delete-todo
// @Accept json
// @Produce json
// @Param todoId path int true "ID of the todo"
// @Success 200 {object} dto.NewTodoResponse
// @Router /todos/{todoId} [delete]
func (th *todoHandler) DeleteTodo(ctx *gin.Context) {
	params, _ := strconv.Atoi(ctx.Param("todoId"))
	response, err := th.todoService.DeleteTodo(params)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
