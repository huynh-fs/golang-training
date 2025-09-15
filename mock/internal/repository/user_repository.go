package repository

import "github.com/huynh-fs/gin-api/internal/model"

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	Create(user *model.User) error
}

type RefreshTokenRepository interface {
	Create(token *model.RefreshToken) error
	FindByToken(token string) (*model.RefreshToken, error)
	Delete(token string) error
}