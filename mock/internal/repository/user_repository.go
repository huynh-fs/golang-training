package repository

import "github.com/huynh-fs/gin-api/internal/model"

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	Create(user *model.User) error
}
