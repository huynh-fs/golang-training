package repository

import "github.com/huynh-fs/gin-api/internal/model"

type TodoRepository interface {
	Create(todo *model.Todo) error
	FindAllByUserID(userID uint) ([]model.Todo, error)
	FindByIDAndUserID(id uint, userID uint) (*model.Todo, error)
	Update(todo *model.Todo) error
	Delete(todo *model.Todo) error
}