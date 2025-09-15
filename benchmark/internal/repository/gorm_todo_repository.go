package repository

import (
	"github.com/huynh-fs/gin-api/internal/model"
	"gorm.io/gorm"
)

type gormTodoRepository struct {
	db *gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) TodoRepository {
	return &gormTodoRepository{db: db}
}

func (r *gormTodoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

func (r *gormTodoRepository) FindAllByUserID(userID uint) ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Where("user_id = ?", userID).Find(&todos).Error
	return todos, err
}

func (r *gormTodoRepository) FindByIDAndUserID(id uint, userID uint) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&todo).Error
	return &todo, err
}

func (r *gormTodoRepository) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

func (r *gormTodoRepository) Delete(todo *model.Todo) error {
	return r.db.Delete(todo).Error
}