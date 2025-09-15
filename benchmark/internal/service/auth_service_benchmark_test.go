package service_test

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	"github.com/stretchr/testify/mock"

	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/internal/repository/mocks"
	"github.com/huynh-fs/gin-api/internal/service"
	"github.com/huynh-fs/gin-api/pkg/config"
)

// BenchmarkAuthService_Login đo lường hiệu năng của hàm Login với mock repository.
func BenchmarkAuthService_Login(b *testing.B) {
	// --- SETUP (Thực hiện một lần bên ngoài vòng lặp) ---
	mockUserRepo := new(mocks.UserRepository)
	mockTokenRepo := new(mocks.RefreshTokenRepository)
	cfg := &config.Config{
		JWTAccessSecret:  "benchmark-secret-access",
		JWTRefreshSecret: "benchmark-secret-refresh",
	}
	authService := service.NewAuthService(mockUserRepo, mockTokenRepo, cfg)

	// Chuẩn bị dữ liệu mẫu
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	sampleUser := &model.User{
		Username:     "testuser",
		PasswordHash: string(hashedPassword),
	}
	sampleUser.ID = 1

	// Định nghĩa hành vi của mock (chỉ cần định nghĩa một lần)
	// Giả lập rằng việc tìm user và tạo refresh token luôn thành công.
	mockUserRepo.On("FindByUsername", sampleUser.Username).Return(sampleUser, nil)
	mockTokenRepo.On("Create", mock.Anything).Return(nil)
	
	b.ReportAllocs() // Báo cáo số lần cấp phát bộ nhớ
	b.ResetTimer()   // Bắt đầu tính giờ từ đây, bỏ qua thời gian setup

	// --- VÒNG LẶP BENCHMARK ---
	// Framework sẽ tự động điều chỉnh giá trị của b.N để chạy trong khoảng 1 giây.
	for i := 0; i < b.N; i++ {
		// Hàm cần benchmark được gọi lặp đi lặp lại ở đây.
		_, _, err := authService.Login(sampleUser.Username, password)
		if err != nil {
			b.Fatalf("Login failed during benchmark: %v", err)
		}
	}
}