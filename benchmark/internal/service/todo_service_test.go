package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/huynh-fs/gin-api/internal/dto"
	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/internal/repository/mocks"
	"github.com/huynh-fs/gin-api/internal/service"
)

// --- Test for CreateTodo ---
func TestTodoService_CreateTodo(t *testing.T) {
	testCases := []struct {
		name          string
		inputReq      *dto.CreateTodoRequest
		inputUserID   uint
		mockSetup     func(*mocks.TodoRepository)
		expectedError string
	}{
		{
			name: "Success",
			inputReq: &dto.CreateTodoRequest{
				Title:       "New Todo",
				Description: "A description",
			},
			inputUserID: uint(1),
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				// Mong đợi hàm Create được gọi và không trả về lỗi
				mockRepo.On("Create", mock.AnythingOfType("*model.Todo")).Return(nil).Once()
			},
			expectedError: "",
		},
		{
			name: "Database Error",
			inputReq: &dto.CreateTodoRequest{
				Title:       "Another Todo",
				Description: "This will fail",
			},
			inputUserID: uint(1),
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				// Giả lập lỗi từ database khi tạo
				dbErr := errors.New("db insert error")
				mockRepo.On("Create", mock.AnythingOfType("*model.Todo")).Return(dbErr).Once()
			},
			expectedError: "db insert error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.TodoRepository)
			todoService := service.NewTodoService(mockRepo)
			tc.mockSetup(mockRepo)

			todo, err := todoService.CreateTodo(tc.inputReq, tc.inputUserID)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, todo)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, todo)
				assert.Equal(t, tc.inputReq.Title, todo.Title)
				assert.Equal(t, tc.inputUserID, todo.UserID)
				assert.False(t, todo.Completed) // Mặc định phải là false
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// --- Test for ListTodos ---
func TestTodoService_ListTodos(t *testing.T) {
	sampleUserID := uint(1)
	mockTodos := []model.Todo{
		{Model: gorm.Model{ID: 1}, Title: "Todo 1", UserID: sampleUserID},
		{Model: gorm.Model{ID: 2}, Title: "Todo 2", UserID: sampleUserID},
	}

	testCases := []struct {
		name             string
		inputUserID      uint
		mockSetup        func(*mocks.TodoRepository)
		expectedTodosLen int
		expectedError    string
	}{
		{
			name:        "Success with multiple todos",
			inputUserID: sampleUserID,
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindAllByUserID", sampleUserID).Return(mockTodos, nil).Once()
			},
			expectedTodosLen: 2,
			expectedError:    "",
		},
		{
			name:        "Success with no todos",
			inputUserID: uint(2),
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindAllByUserID", uint(2)).Return([]model.Todo{}, nil).Once()
			},
			expectedTodosLen: 0,
			expectedError:    "",
		},
		{
			name:        "Database Error",
			inputUserID: sampleUserID,
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				dbErr := errors.New("db find all error")
				mockRepo.On("FindAllByUserID", sampleUserID).Return(nil, dbErr).Once()
			},
			expectedTodosLen: 0,
			expectedError:    "db find all error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.TodoRepository)
			todoService := service.NewTodoService(mockRepo)
			tc.mockSetup(mockRepo)

			todos, err := todoService.ListTodos(tc.inputUserID)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, todos)
			} else {
				assert.NoError(t, err)
				assert.Len(t, todos, tc.expectedTodosLen)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// --- Test for GetTodo ---
func TestTodoService_GetTodo(t *testing.T) {
	sampleUserID := uint(1)
	sampleTodoID := uint(10)
	sampleTodo := &model.Todo{
		Model:  gorm.Model{ID: sampleTodoID},
		Title:  "Sample Todo",
		UserID: sampleUserID,
	}

	testCases := []struct {
		name          string
		todoID        uint
		userID        uint
		mockSetup     func(*mocks.TodoRepository)
		expectedTodo  *model.Todo
		expectedError string
	}{
		{
			name:   "Success",
			todoID: sampleTodoID,
			userID: sampleUserID,
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindByIDAndUserID", sampleTodoID, sampleUserID).Return(sampleTodo, nil).Once()
			},
			expectedTodo:  sampleTodo,
			expectedError: "",
		},
		{
			name:   "Todo not found",
			todoID: uint(99),
			userID: sampleUserID,
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindByIDAndUserID", uint(99), sampleUserID).Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedTodo:  nil,
			expectedError: "record not found",
		},
		{
			name:   "User not authorized",
			todoID: sampleTodoID,
			userID: uint(2), // User khác
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindByIDAndUserID", sampleTodoID, uint(2)).Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedTodo:  nil,
			expectedError: "record not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.TodoRepository)
			todoService := service.NewTodoService(mockRepo)
			tc.mockSetup(mockRepo)

			todo, err := todoService.GetTodo(tc.todoID, tc.userID)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedTodo, todo)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// --- Test for UpdateTodo ---
func TestTodoService_UpdateTodo(t *testing.T) {
	sampleUserID := uint(1)
	sampleTodoID := uint(10)

	testCases := []struct {
		name          string
		inputID       uint
		inputUserID   uint
		inputReq      *dto.UpdateTodoRequest
		mockSetup     func(*mocks.TodoRepository)
		expectedError string
	}{
		{
			name:        "Success - Full Update",
			inputID:     sampleTodoID,
			inputUserID: sampleUserID,
			inputReq: &dto.UpdateTodoRequest{
				Title:       stringPtr("Updated Title"),
				Description: stringPtr("Updated Description"),
				Completed:   boolPtr(true),
			},
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				// Trả về todo gốc khi tìm
				existingTodo := &model.Todo{Model: gorm.Model{ID: sampleTodoID}, UserID: sampleUserID, Title: "Original"}
				mockRepo.On("FindByIDAndUserID", sampleTodoID, sampleUserID).Return(existingTodo, nil).Once()
				// Mong đợi việc Update thành công
				mockRepo.On("Update", mock.AnythingOfType("*model.Todo")).Return(nil).Once()
			},
			expectedError: "",
		},
		{
			name:        "Success - Partial Update (only title)",
			inputID:     sampleTodoID,
			inputUserID: sampleUserID,
			inputReq:    &dto.UpdateTodoRequest{Title: stringPtr("Only Title Updated")},
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				existingTodo := &model.Todo{Model: gorm.Model{ID: sampleTodoID}, UserID: sampleUserID}
				mockRepo.On("FindByIDAndUserID", sampleTodoID, sampleUserID).Return(existingTodo, nil).Once()
				mockRepo.On("Update", existingTodo).Return(nil).Once()
			},
			expectedError: "",
		},
		{
			name:        "Todo Not Found",
			inputID:     uint(99),
			inputUserID: sampleUserID,
			inputReq:    &dto.UpdateTodoRequest{Title: stringPtr("Wont update")},
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindByIDAndUserID", uint(99), sampleUserID).Return(nil, gorm.ErrRecordNotFound).Once()
				// Quan trọng: hàm Update không bao giờ được gọi
			},
			expectedError: "record not found",
		},
		{
			name:        "Database Error on Update",
			inputID:     sampleTodoID,
			inputUserID: sampleUserID,
			inputReq:    &dto.UpdateTodoRequest{Title: stringPtr("Will fail")},
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				existingTodo := &model.Todo{Model: gorm.Model{ID: sampleTodoID}, UserID: sampleUserID}
				mockRepo.On("FindByIDAndUserID", sampleTodoID, sampleUserID).Return(existingTodo, nil).Once()
				dbErr := errors.New("db update error")
				mockRepo.On("Update", existingTodo).Return(dbErr).Once()
			},
			expectedError: "db update error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.TodoRepository)
			todoService := service.NewTodoService(mockRepo)
			tc.mockSetup(mockRepo)

			updatedTodo, err := todoService.UpdateTodo(tc.inputID, tc.inputReq, tc.inputUserID)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, updatedTodo)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, updatedTodo)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// Helper functions để tạo con trỏ cho các kiểu dữ liệu cơ bản, dùng trong Update test
func stringPtr(s string) *string { return &s }
func boolPtr(b bool) *bool       { return &b }


// --- Test for DeleteTodo ---
func TestTodoService_DeleteTodo(t *testing.T) {
	sampleUserID := uint(1)
	sampleTodoID := uint(10)
	existingTodo := &model.Todo{Model: gorm.Model{ID: sampleTodoID}, UserID: sampleUserID}

	testCases := []struct {
		name          string
		inputID       uint
		inputUserID   uint
		mockSetup     func(*mocks.TodoRepository)
		expectedError string
	}{
		{
			name:        "Success",
			inputID:     sampleTodoID,
			inputUserID: sampleUserID,
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindByIDAndUserID", sampleTodoID, sampleUserID).Return(existingTodo, nil).Once()
				mockRepo.On("Delete", existingTodo).Return(nil).Once()
			},
			expectedError: "",
		},
		{
			name:        "Todo Not Found",
			inputID:     uint(99),
			inputUserID: sampleUserID,
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindByIDAndUserID", uint(99), sampleUserID).Return(nil, gorm.ErrRecordNotFound).Once()
				// Quan trọng: hàm Delete không bao giờ được gọi
			},
			expectedError: "record not found",
		},
		{
			name:        "Database Error on Delete",
			inputID:     sampleTodoID,
			inputUserID: sampleUserID,
			mockSetup: func(mockRepo *mocks.TodoRepository) {
				mockRepo.On("FindByIDAndUserID", sampleTodoID, sampleUserID).Return(existingTodo, nil).Once()
				dbErr := errors.New("db delete error")
				mockRepo.On("Delete", existingTodo).Return(dbErr).Once()
			},
			expectedError: "db delete error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.TodoRepository)
			todoService := service.NewTodoService(mockRepo)
			tc.mockSetup(mockRepo)

			err := todoService.DeleteTodo(tc.inputID, tc.inputUserID)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}