package repository

import "github.com/huynh-fs/gin-api/internal/model"

type RefreshTokenRepository interface {
	Create(token *model.RefreshToken) error
	FindByToken(token string) (*model.RefreshToken, error)
	Delete(token string) error
}