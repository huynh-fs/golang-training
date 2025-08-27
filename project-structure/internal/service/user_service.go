package service

import (
	domain "github.com/huynh-fs/golang-training/user-service/internal/model"
	"github.com/huynh-fs/golang-training/user-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.GetUserByID(id)
}