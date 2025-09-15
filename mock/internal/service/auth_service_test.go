package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/internal/repository/mocks" 
	"github.com/huynh-fs/gin-api/internal/service"
	"github.com/huynh-fs/gin-api/pkg/config"
)

func generateTestRefreshToken(userID uint, secret string) (string, time.Time) {
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	claims := &service.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenJwt.SignedString([]byte(secret))
	return token, expiresAt
}

func TestAuthService_Register(t *testing.T) {
	testCases := []struct {
		name          string
		username      string
		password      string
		mockSetup     func(*mocks.UserRepository)
		expectedError string
	}{
		{
			name:     "Success",
			username: "newuser",
			password: "password123",
			mockSetup: func(mockUserRepo *mocks.UserRepository) {
				mockUserRepo.On("FindByUsername", "newuser").Return(nil, gorm.ErrRecordNotFound).Once()
				mockUserRepo.On("Create", mock.AnythingOfType("*model.User")).Return(nil).Once()
			},
			expectedError: "",
		},
		{
			name:     "Username already exists",
			username: "existinguser",
			password: "password123",
			mockSetup: func(mockUserRepo *mocks.UserRepository) {
				mockUserRepo.On("FindByUsername", "existinguser").Return(&model.User{}, nil).Once()
			},
			expectedError: "username already exists",
		},
		{
			name:     "Database error on find",
			username: "anotheruser",
			password: "password123",
			mockSetup: func(mockUserRepo *mocks.UserRepository) {
				dbErr := errors.New("unexpected database error")
				mockUserRepo.On("FindByUsername", "anotheruser").Return(nil, dbErr).Once()
			},
			expectedError: "unexpected database error",
		},
		{
			name:     "Database error on create",
			username: "createfailuser",
			password: "password123",
			mockSetup: func(mockUserRepo *mocks.UserRepository) {
				mockUserRepo.On("FindByUsername", "createfailuser").Return(nil, gorm.ErrRecordNotFound).Once()
				dbErr := errors.New("failed to create user in db")
				mockUserRepo.On("Create", mock.AnythingOfType("*model.User")).Return(dbErr).Once()
			},
			expectedError: "failed to create user in db",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUserRepo := new(mocks.UserRepository)
			// Register không cần các repo hay config khác
			authService := service.NewAuthService(mockUserRepo, nil, nil)

			tc.mockSetup(mockUserRepo)

			user, err := authService.Register(tc.username, tc.password)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.username, user.Username)
			}

			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	// Chuẩn bị một user mẫu với mật khẩu đã được hash
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	sampleUser := &model.User{
		Username:     "testuser",
		PasswordHash: string(hashedPassword),
	}
	sampleUser.ID = 1 // Gán ID để sử dụng trong token claims

	// 1. Định nghĩa struct cho test case
	testCases := []struct {
		name          string
		username      string
		password      string
		mockSetup     func(*mocks.UserRepository, *mocks.RefreshTokenRepository) 
		expectedError string
	}{
		// 2. "Bảng" các test case
		{
			name:     "Success",
			username: "testuser",
			password: "password123",
			mockSetup: func(mockUserRepo *mocks.UserRepository, mockTokenRepo *mocks.RefreshTokenRepository) {
				// Mong đợi s.userRepo.FindByUsername được gọi và trả về user mẫu
				mockUserRepo.On("FindByUsername", "testuser").Return(sampleUser, nil).Once()
				// Mong đợi s.tokenRepo.Create được gọi để lưu refresh token và thành công
				mockTokenRepo.On("Create", mock.AnythingOfType("*model.RefreshToken")).Return(nil).Once()
			},
			expectedError: "", // Không có lỗi
		},
		{
			name:     "User not found",
			username: "nonexistent",
			password: "password123",
			mockSetup: func(mockUserRepo *mocks.UserRepository, mockTokenRepo *mocks.RefreshTokenRepository) {
				// Mong đợi s.userRepo.FindByUsername được gọi và trả về lỗi "không tìm thấy"
				mockUserRepo.On("FindByUsername", "nonexistent").Return(nil, errors.New("record not found")).Once()
				// Quan trọng: tokenRepo.Create KHÔNG được gọi
			},
			expectedError: "invalid credentials",
		},
		{
			name:     "Incorrect password",
			username: "testuser",
			password: "wrongpassword",
			mockSetup: func(mockUserRepo *mocks.UserRepository, mockTokenRepo *mocks.RefreshTokenRepository) {
				// Mong đợi s.userRepo.FindByUsername được gọi và trả về user mẫu (để có thể so sánh password)
				mockUserRepo.On("FindByUsername", "testuser").Return(sampleUser, nil).Once()
				// Quan trọng: tokenRepo.Create KHÔNG được gọi
			},
			expectedError: "invalid credentials",
		},
		{
			name:     "Database error on saving refresh token",
			username: "testuser",
			password: "password123",
			mockSetup: func(mockUserRepo *mocks.UserRepository, mockTokenRepo *mocks.RefreshTokenRepository) {
				// User được tìm thấy thành công
				mockUserRepo.On("FindByUsername", "testuser").Return(sampleUser, nil).Once()
				// Nhưng khi lưu refresh token thì gặp lỗi DB
				mockTokenRepo.On("Create", mock.AnythingOfType("*model.RefreshToken")).Return(errors.New("failed to save token")).Once()
			},
			expectedError: "failed to save token",
		},
	}

	// 3. Vòng lặp chạy test
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			mockUserRepo := new(mocks.UserRepository)
			mockTokenRepo := new(mocks.RefreshTokenRepository)
			cfg := &config.Config{
				JWTAccessSecret:  "test-secret",
				JWTRefreshSecret: "test-secret-refresh",
			}
			authService := service.NewAuthService(mockUserRepo, mockTokenRepo, cfg)

			// Áp dụng mock setup cho case hiện tại
			tc.mockSetup(mockUserRepo, mockTokenRepo)

			// Thực thi
			accessToken, refreshToken, err := authService.Login(tc.username, tc.password)

			// Kiểm tra kết quả
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Empty(t, accessToken)
				assert.Empty(t, refreshToken)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, accessToken)
				assert.NotEmpty(t, refreshToken)
			}

			// Xác minh mock
			mockUserRepo.AssertExpectations(t)
			mockTokenRepo.AssertExpectations(t)
		})
	}
}

func TestAuthService_Logout(t *testing.T) {
	testCases := []struct {
		name          string
		inputToken    string
		mockSetup     func(*mocks.RefreshTokenRepository)
		expectedError string
	}{
		{
			name:       "Success",
			inputToken: "valid-token-to-logout",
			mockSetup: func(mockTokenRepo *mocks.RefreshTokenRepository) {
				// Mong đợi việc xóa token thành công và trả về nil (không có lỗi)
				mockTokenRepo.On("Delete", "valid-token-to-logout").Return(nil).Once()
			},
			expectedError: "",
		},
		{
			name:       "Token not found or already revoked",
			inputToken: "non-existent-token",
			mockSetup: func(mockTokenRepo *mocks.RefreshTokenRepository) {
				// Giả lập lỗi trả về từ repository khi không tìm thấy token để xóa
				// Đây là nơi bạn cần đảm bảo implementation GORM của bạn trả về lỗi
				// khi không có dòng nào bị ảnh hưởng.
				mockTokenRepo.On("Delete", "non-existent-token").Return(gorm.ErrRecordNotFound).Once()
			},
			expectedError: "token not found or already revoked",
		},
		{
			name:       "Database error on delete",
			inputToken: "any-token",
			mockSetup: func(mockTokenRepo *mocks.RefreshTokenRepository) {
				// Giả lập lỗi DB bất kỳ khi thực hiện thao tác xóa
				dbErr := errors.New("unexpected db error")
				mockTokenRepo.On("Delete", "any-token").Return(dbErr).Once()
			},
			expectedError: "unexpected db error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockTokenRepo := new(mocks.RefreshTokenRepository)
			// Logout không cần userRepo hay config
			authService := service.NewAuthService(nil, mockTokenRepo, nil)

			tc.mockSetup(mockTokenRepo)

			err := authService.Logout(tc.inputToken)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockTokenRepo.AssertExpectations(t)
		})
	}
}

func TestAuthService_RefreshToken(t *testing.T) {
	cfg := &config.Config{
		JWTAccessSecret:  "test-access-secret",
		JWTRefreshSecret: "test-refresh-secret",
	}
	sampleUserID := uint(1)
	validRefreshToken, _ := generateTestRefreshToken(sampleUserID, cfg.JWTRefreshSecret)

	testCases := []struct {
		name          string
		inputToken    string
		mockSetup     func(*mocks.UserRepository, *mocks.RefreshTokenRepository)
		expectedError string
	}{
		{
			name:       "Success",
			inputToken: validRefreshToken,
			mockSetup: func(_ *mocks.UserRepository, mockTokenRepo *mocks.RefreshTokenRepository) {
				// 1. Mong đợi việc tìm token trong DB thành công
				mockTokenRepo.On("FindByToken", validRefreshToken).Return(&model.RefreshToken{UserID: sampleUserID, Token: validRefreshToken}, nil).Once()
				// 2. Mong đợi việc xóa token cũ thành công
				mockTokenRepo.On("Delete", validRefreshToken).Return(nil).Once()
				// 3. Mong đợi việc tạo token mới thành công
				mockTokenRepo.On("Create", mock.AnythingOfType("*model.RefreshToken")).Return(nil).Once()
			},
			expectedError: "",
		},
		{
			name:       "Token has been revoked (not in DB)",
			inputToken: validRefreshToken,
			mockSetup: func(_ *mocks.UserRepository, mockTokenRepo *mocks.RefreshTokenRepository) {
				// Khi tìm token trong DB, trả về lỗi "không tìm thấy"
				mockTokenRepo.On("FindByToken", validRefreshToken).Return(nil, gorm.ErrRecordNotFound).Once()
			},
			expectedError: "token has been revoked",
		},
		{
			name:          "Malformed or invalid token string",
			inputToken:    "this.is.not.a.valid.token",
			mockSetup:     func(_ *mocks.UserRepository, _ *mocks.RefreshTokenRepository) {
				// Repository sẽ không được gọi vì việc parse token thất bại trước đó
			},
			expectedError: "invalid refresh token",
		},
		{
			name:       "Database error on deleting old token",
			inputToken: validRefreshToken,
			mockSetup: func(_ *mocks.UserRepository, mockTokenRepo *mocks.RefreshTokenRepository) {
				// Tìm token thành công
				mockTokenRepo.On("FindByToken", validRefreshToken).Return(&model.RefreshToken{Token: validRefreshToken}, nil).Once()
				// Nhưng xóa token cũ thất bại
				mockTokenRepo.On("Delete", validRefreshToken).Return(errors.New("db delete error")).Once()
			},
			expectedError: "db delete error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockUserRepo := new(mocks.UserRepository)
			mockTokenRepo := new(mocks.RefreshTokenRepository)
			authService := service.NewAuthService(mockUserRepo, mockTokenRepo, cfg)

			tc.mockSetup(mockUserRepo, mockTokenRepo)

			newAccessToken, newRefreshToken, err := authService.RefreshToken(tc.inputToken)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Empty(t, newAccessToken)
				assert.Empty(t, newRefreshToken)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, newAccessToken)
				assert.NotEmpty(t, newRefreshToken)
			}

			mockUserRepo.AssertExpectations(t)
			mockTokenRepo.AssertExpectations(t)
		})
	}
}

