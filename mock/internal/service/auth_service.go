package service

import (
	"errors"
	"time"

	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/internal/repository"
	"github.com/huynh-fs/gin-api/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	userRepo repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
	config *config.Config
}

func NewAuthService(userRepo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository, config *config.Config) *AuthService {
	return &AuthService{userRepo: userRepo, refreshTokenRepo: refreshTokenRepo, config: config}
}

func (s *AuthService) Register(username, password string) (*model.User, error) {
	_, err := s.userRepo.FindByUsername(username)
	if err == nil {
		return nil, errors.New("username already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err 
	}

	user := model.User{Username: username}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) Login(username, password string) (accessToken string, refreshToken string, err error) {
    user, err := s.userRepo.FindByUsername(username)
    if err != nil {
        return "", "", errors.New("invalid credentials")
    }

    if err = user.CheckPassword(password); err != nil {
        return "", "", errors.New("invalid credentials")
    }

    return s.generateTokens(user.ID)
}

func (s *AuthService) RefreshToken(tokenString string) (newAccessToken string, newRefreshToken string, err error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.config.JWTRefreshSecret), nil
    })
    if err != nil || !token.Valid {
        return "", "", errors.New("invalid refresh token")
    }

    dbToken, err := s.refreshTokenRepo.FindByToken(tokenString)
    if err != nil {
        return "", "", errors.New("refresh token has been revoked or is invalid")
    }

    if err := s.refreshTokenRepo.Delete(dbToken.Token); err != nil {
        return "", "", err
    }

    return s.generateTokens(claims.UserID)
}

func (s *AuthService) Logout(tokenString string) error {
    	err := s.refreshTokenRepo.Delete(tokenString)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("token not found or already revoked")
		}
		return err 
	}
	return nil
}

func (s *AuthService) generateTokens(userID uint) (accessToken string, refreshToken string, err error) {
	accessExpiresAt := time.Now().Add(15 * time.Minute)
	accessClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiresAt),
		},
	}
	accessTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessTokenJwt.SignedString([]byte(s.config.JWTAccessSecret))
	if err != nil {
		return "", "", err
	}

	refreshExpiresAt := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiresAt),
		},
	}
	refreshTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTokenJwt.SignedString([]byte(s.config.JWTRefreshSecret))
	if err != nil {
		return "", "", err
	}

	dbToken := model.RefreshToken{
		UserID:    userID,
		Token:     refreshToken, 
		ExpiresAt: refreshExpiresAt,
	}
	if err := s.refreshTokenRepo.Create(&dbToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
