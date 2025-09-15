package service_test

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	"github.com/stretchr/testify/mock"
	
	"github.com/huynh-fs/gin-api/internal/dto"
	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/internal/repository"
	"github.com/huynh-fs/gin-api/internal/repository/mocks"
	"github.com/huynh-fs/gin-api/internal/service"
)

// --- UNIT BENCHMARK (dùng mock) ---

func BenchmarkAuthService_Login_Unit(b *testing.B) {
	// Setup (không phụ thuộc vào TestMain)
	mockUserRepo := new(mocks.UserRepository)
	mockTokenRepo := new(mocks.RefreshTokenRepository)
	authService := service.NewAuthService(mockUserRepo, mockTokenRepo, testCfg)

	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	sampleUser := &model.User{Username: "testuser", PasswordHash: string(hashedPassword)}
	sampleUser.ID = 1

	mockUserRepo.On("FindByUsername", sampleUser.Username).Return(sampleUser, nil)
	mockTokenRepo.On("Create", mock.Anything).Return(nil)
	
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := authService.Login(sampleUser.Username, password)
		if err != nil {
			b.Fatalf("Login failed during benchmark: %v", err)
		}
	}
}

// --- INTEGRATION BENCHMARK (dùng DB thật từ TestMain) ---

// Hàm helper để chạy benchmark
func benchmarkCreateTodo(b *testing.B, todoService *service.TodoService, userID uint) {
	createReq := &dto.CreateTodoRequest{
		Title:       "Benchmark Todo",
		Description: "This is created during benchmark",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := todoService.CreateTodo(createReq, userID)
		if err != nil {
			b.Fatalf("CreateTodo failed during benchmark: %v", err)
		}
	}
}

// Hàm benchmark chính, chỉ lo việc setup
func BenchmarkTodoService_CreateTodo_Integration(b *testing.B) {
	if testDB == nil {
		b.Skip("Skipping integration benchmark: database not connected")
	}

	// --- SETUP MỘT LẦN DUY NHẤT CHO TOÀN BỘ BENCHMARK NÀY ---
	testUser := model.User{Username: "benchmark_user"}
	testUser.SetPassword("password")
	
	// Xóa user cũ nếu có để đảm bảo môi trường sạch
	testDB.Unscoped().Where("username = ?", testUser.Username).Delete(&model.User{})

	// Tạo user mới
	if err := testDB.Create(&testUser).Error; err != nil {
		b.Fatalf("Failed to create seed user for benchmark: %v", err)
	}
	
	// Setup service
	todoRepo := repository.NewGormTodoRepository(testDB)
	todoService := service.NewTodoService(todoRepo)

	// --- Chạy các sub-benchmark ---
	// Chúng ta có thể chạy nhiều kịch bản benchmark khác nhau ở đây
	b.Run("CreateSingleTodo", func(b *testing.B) {
		benchmarkCreateTodo(b, todoService, testUser.ID)
	})

	// ... có thể thêm b.Run("CreateBatchTodo", ...) ở đây trong tương lai ...
}