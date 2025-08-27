package repository

import (
	"errors"

	domain "github.com/huynh-fs/golang-training/user-service/internal/model"
)

// UserRepository xử lý việc truy cập dữ liệu người dùng.
type UserRepository struct {
	// Trong một ứng dụng thực tế, ở đây sẽ có một kết nối database
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetUserByID giả lập việc lấy user từ database.
func (r *UserRepository) GetUserByID(id string) (*domain.User, error) {
	// Giả lập dữ liệu
	if id == "1" {
		return &domain.User{
			ID:    "1",
			Name:  "Alice",
			Email: "alice@example.com",
		}, nil
	}
	return nil, errors.New("user not found")
}