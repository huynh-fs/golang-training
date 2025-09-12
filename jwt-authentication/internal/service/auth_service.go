package service

import (
	"errors"
	"time"

	"github.com/huynh-fs/gin-api/internal/model"
	"github.com/huynh-fs/gin-api/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	db *gorm.DB
	config *config.Config
}

func NewAuthService(db *gorm.DB, config *config.Config) *AuthService {
	return &AuthService{db: db, config: config}
}

func (s *AuthService) Register(username, password string) (*model.User, error) {
	var existingUser model.User
	if err := s.db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	user := model.User{Username: username}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) Login(username, password string) (accessToken string, refreshToken string, err error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if err := user.CheckPassword(password); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	return s.generateTokens(user.ID)
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
	if err := s.db.Create(&dbToken).Error; err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshToken(tokenString string) (newAccessToken string, newRefreshToken string, err error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTRefreshSecret), nil
	})

	if err != nil || !token.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	var dbToken model.RefreshToken
	if err := s.db.Where("token = ?", tokenString).First(&dbToken).Error; err != nil {
		return "", "", errors.New("refresh token has been revoked or is invalid")
	}

	s.db.Delete(&dbToken)

	return s.generateTokens(claims.UserID)
}


func (s *AuthService) Logout(tokenString string) error {
    result := s.db.Where("token = ?", tokenString).Delete(&model.RefreshToken{})
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("token not found or already revoked")
    }
    return nil
}