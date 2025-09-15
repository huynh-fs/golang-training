package handler

import (
	"net/http"
	"github.com/huynh-fs/gin-api/internal/dto"
	"github.com/huynh-fs/gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Đăng kí người dùng mới
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User registration info"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user, err := h.authService.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user_id": user.ID})
}


// Login godoc
// @Summary Đăng nhập
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "User credentials"
// @Success 200 {object} dto.TokenResponse
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    var req dto.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    accessToken, refreshToken, err := h.authService.Login(req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, dto.TokenResponse{
        AccessToken: accessToken,
        RefreshToken: refreshToken,
    })
}

// RefreshToken godoc
// @Summary Refresh access token
// @Tags auth
// @Accept json
// @Produce json
// @Param token body dto.TokenRequest true "Refresh token"
// @Success 200 {object} dto.TokenResponse
// @Failure 401 {object} map[string]string
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
    var req dto.TokenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    newAccessToken, newRefreshToken, err := h.authService.RefreshToken(req.RefreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

     c.JSON(http.StatusOK, dto.TokenResponse{
        AccessToken: newAccessToken,
        RefreshToken: newRefreshToken,
    })
}

// Logout godoc
// @Summary Đăng xuất và thu hồi refresh token
// @Tags auth
// @Accept json
// @Param token body dto.TokenRequest true "Refresh token to revoke"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
    var req dto.TokenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    
    if err := h.authService.Logout(req.RefreshToken); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}