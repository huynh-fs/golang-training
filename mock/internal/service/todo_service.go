package service

import (
	"github.com/huynh-fs/gin-api/internal/dto" // Import DTO
	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/internal/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

// SỬA LẠI: Nhận vào con trỏ đến CreateTodoRequest DTO
func (s *TodoService) CreateTodo(req *dto.CreateTodoRequest, userID uint) (*model.Todo, error) {
	todo := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		Completed:   false, // Một todo mới luôn mặc định là chưa hoàn thành
	}

	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// Hàm này giữ nguyên, nó không nhận DTO
func (s *TodoService) ListTodos(userID uint) ([]model.Todo, error) {
	return s.repo.FindAllByUserID(userID)
}

// Hàm này giữ nguyên
func (s *TodoService) GetTodo(id uint, userID uint) (*model.Todo, error) {
	return s.repo.FindByIDAndUserID(id, userID)
}

// SỬA LẠI: Nhận vào con trỏ đến UpdateTodoRequest DTO
func (s *TodoService) UpdateTodo(id uint, req *dto.UpdateTodoRequest, userID uint) (*model.Todo, error) {
	// 1. Tìm todo và kiểm tra quyền sở hữu
	todo, err := s.repo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, err // Không tìm thấy hoặc không có quyền
	}

	// 2. Chỉ cập nhật các trường được cung cấp trong request (khác nil)
	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}

	// 3. Lưu lại thay đổi
	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}
	
	return todo, nil
}

// Hàm này giữ nguyên
func (s *TodoService) DeleteTodo(id uint, userID uint) error {
	todo, err := s.repo.FindByIDAndUserID(id, userID)
	if err != nil {
		return err // Không tìm thấy hoặc không có quyền
	}
	return s.repo.Delete(todo)
}