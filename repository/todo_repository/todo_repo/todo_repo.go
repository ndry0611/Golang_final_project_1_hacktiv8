package todo_repository

import (
	"errors"
	"final_project_1/entity"
	"final_project_1/pkg/errors_response"
	"final_project_1/repository/todo_repository"
	"strconv"

	"gorm.io/gorm"
)

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) todo_repository.Repository {
	return &todoRepo{db: db}
}

func (tr *todoRepo) CreateTodo(todoPayload *entity.Todo) (*entity.Todo, errors_response.ErrorResponse) {
	var Todo = *todoPayload
	err := tr.db.Create(&Todo).Error
	if err != nil {
		return nil, errors_response.NewInternalServerError("something went wrong")
	}
	return &Todo, nil
}

func (tr *todoRepo) GetTodos() (*[]entity.Todo, errors_response.ErrorResponse) {
	var Todos []entity.Todo
	err := tr.db.Find(&Todos).Error

	if err != nil {
		return nil, errors_response.NewInternalServerError("something went wrong")
	}
	return &Todos, nil
}

func (tr *todoRepo) GetTodo(id int) (*entity.Todo, errors_response.ErrorResponse) {
	var Todo entity.Todo
	err := tr.db.First(&Todo, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "todo with id: " + strconv.Itoa(id) + " not found"
			return nil, errors_response.NewNotFoundError(msg)
		}
		return nil, errors_response.NewInternalServerError("something went wrong")
	}
	return &Todo, nil
}

func (tr *todoRepo) UpdateTodo(id int, todoPayload *entity.Todo) (*entity.Todo, errors_response.ErrorResponse) {
	var Todo = *todoPayload

	err := tr.db.Model(&Todo).Where("id = ?", id).Updates(entity.Todo{Title: todoPayload.Title, Done: todoPayload.Done}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "todo with id: " + strconv.Itoa(id) + " not found"
			return nil, errors_response.NewNotFoundError(msg)
		}
		return nil, errors_response.NewInternalServerError("something went wrong")
	}
	return &Todo, nil
}

func (tr *todoRepo) DeleteTodo(id int) errors_response.ErrorResponse {
	err := tr.db.Where("id = ?", id).Delete(&entity.Todo{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "todo with id: " + strconv.Itoa(id) + " not found"
			return errors_response.NewNotFoundError(msg)
		}
		return errors_response.NewInternalServerError("something went wrong")
	}
	return nil
}
