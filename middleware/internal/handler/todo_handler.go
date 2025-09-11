package handler

import (
	"net/http"
	"strconv"
	"github.com/huynh-fs/gin-api/internal/dto"
	"github.com/huynh-fs/gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct{
	Service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{
		Service: s,
	}
}


// CreateTodo godoc
// @Summary      Tạo một công việc mới
// @Description  Thêm một công việc mới vào danh sách với trạng thái mặc định là chưa hoàn thành
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body   dto.CreateTodoDTO  true  "Chỉ cần nhập tiêu đề công việc"
// @Success      200  {object}  model.Todo
// @Router       /todos [post]
// @Security BearerAuth
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var input dto.CreateTodoDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo, err := h.Service.CreateTodo(input.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTodo)
}

// GetTodos godoc
// @Summary Lấy danh sách công việc
// @Description Lấy tất cả công việc trong danh sách
// @Tags todos
// @Produce  json
// @Success 200 {array} model.Todo
// @Router /todos [get]
// @Security BearerAuth
func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.Service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GetTodo godoc
// @Summary Lấy một công việc theo ID
// @Description Lấy thông tin chi tiết của một công việc
// @Tags todos
// @Produce  json
// @Param   id   path      int  true  "ID Công việc"
// @Success 200 {object} model.Todo
// @Failure 404 {object} map[string]interface{}
// @Router /todos/{id} [get]
// @Security BearerAuth
func (h *TodoHandler) GetTodo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	todo, err := h.Service.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo godoc
// @Summary Cập nhật một công việc
// @Description Cập nhật thông tin của một công việc đã có
// @Tags todos
// @Accept  json
// @Produce  json
// @Param   id   path      int  true  "ID Công việc"
// @Param   todo  body   dto.UpdateTodoDTO  true  "Thông tin cập nhật"
// @Success 200 {object} model.Todo
// @Failure 404 {object} map[string]interface{}
// @Router /todos/{id} [put]
// @Security BearerAuth
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var input dto.UpdateTodoDTO
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTodo, err := h.Service.UpdateTodo(uint(id), input.Title, input.Completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

// DeleteTodo godoc
// @Summary Xóa một công việc
// @Description Xóa một công việc khỏi danh sách
// @Tags todos
// @Produce  json
// @Param   id   path      int  true  "ID Công việc"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /todos/{id} [delete]
// @Security BearerAuth
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.Service.DeleteTodo(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}