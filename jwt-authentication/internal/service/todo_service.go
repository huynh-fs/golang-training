package service

import (
	"gorm.io/gorm"
	"github.com/huynh-fs/gin-api/internal/model"
)

type TodoService struct {
	DB *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{DB: db}
}

func (s *TodoService) GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	err := s.DB.Find(&todos).Error
	return todos, err
}

func (s *TodoService) GetTodoByID(id uint) (model.Todo, error) {
	var todo model.Todo
	err := s.DB.First(&todo, id).Error
	return todo, err
}

func (s *TodoService) CreateTodo(title string) (model.Todo, error) {
	todo := model.Todo{
		Title:     title,
		Completed: false, 
	}
	err := s.DB.Create(&todo).Error
	return todo, err
}

func (s *TodoService) UpdateTodo(id uint, title string, completed bool) (model.Todo, error) {
	todo, err := s.GetTodoByID(id)
	if err != nil {
		return todo, err
	}

	todo.Title = title
	todo.Completed = completed

	err = s.DB.Save(&todo).Error
	return todo, err
}

func (s *TodoService) DeleteTodo(id uint) error {
	todo, err := s.GetTodoByID(id)
	if err != nil {
		return err
	}
	return s.DB.Delete(&todo).Error
}