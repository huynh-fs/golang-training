package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/huynh-fs/gin-api/internal/dto"
	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/internal/service"
)

type TodoHandler struct {
	todoService *service.TodoService // Đổi tên cho nhất quán (camelCase)
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: s,
	}
}

// Helper function để lấy userID từ context một cách an toàn
func getUserIDFromContext(c *gin.Context) (uint, error) {
	val, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("userID not found in context")
	}
	userID, ok := val.(uint)
	if !ok {
		return 0, errors.New("userID is of invalid type")
	}
	return userID, nil
}

// Helper function để chuyển đổi model.Todo sang dto.TodoResponse
func toTodoResponse(todo *model.Todo) dto.TodoResponse {
	return dto.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		UserID:      todo.UserID,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}

// CreateTodo godoc
// @Summary      Tạo một công việc mới
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body   dto.CreateTodoRequest  true  "Thông tin công việc mới"
// @Success      201  {object}  dto.TodoResponse
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /todos [post]
// @Security     BearerAuth
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req dto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo, err := h.todoService.CreateTodo(&req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, toTodoResponse(newTodo))
}

// ListTodos godoc
// @Summary Lấy danh sách công việc của người dùng
// @Tags todos
// @Produce  json
// @Success 200 {array} dto.TodoResponse
// @Failure 500 {object} map[string]string
// @Router /todos [get]
// @Security BearerAuth
func (h *TodoHandler) ListTodos(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	todos, err := h.todoService.ListTodos(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
		return
	}

	// Chuyển đổi một slice của model.Todo sang slice của dto.TodoResponse
	var todoResponses []dto.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, toTodoResponse(&todo))
	}

	c.JSON(http.StatusOK, todoResponses)
}

// GetTodo godoc
// @Summary Lấy một công việc theo ID
// @Tags todos
// @Produce  json
// @Param   id   path      int  true  "ID Công việc"
// @Success 200 {object} dto.TodoResponse
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [get]
// @Security BearerAuth
func (h *TodoHandler) GetTodo(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	todo, err := h.todoService.GetTodo(uint(id), userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found or you don't have permission"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todo"})
		return
	}

	c.JSON(http.StatusOK, toTodoResponse(todo))
}

// UpdateTodo godoc
// @Summary Cập nhật một công việc
// @Tags todos
// @Accept  json
// @Produce  json
// @Param   id   path      int  true  "ID Công việc"
// @Param   todo  body   dto.UpdateTodoRequest  true  "Thông tin cập nhật (chỉ cần gửi các trường muốn thay đổi)"
// @Success 200 {object} dto.TodoResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [put]
// @Security BearerAuth
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req dto.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	updatedTodo, err := h.todoService.UpdateTodo(uint(id), &req, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found or you don't have permission"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, toTodoResponse(updatedTodo))
}

// DeleteTodo godoc
// @Summary Xóa một công việc
// @Tags todos
// @Produce  json
// @Param   id   path      int  true  "ID Công việc"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [delete]
// @Security BearerAuth
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.todoService.DeleteTodo(uint(id), userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found or you don't have permission"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}
	
	// Trả về 204 No Content là một chuẩn RESTful phổ biến cho việc xóa thành công.
	c.Status(http.StatusNoContent)
}