package main

import (
	"log"

	_ "github.com/huynh-fs/gin-api/docs"
	"github.com/huynh-fs/gin-api/internal/handler"
	"github.com/huynh-fs/gin-api/internal/repository"
	"github.com/huynh-fs/gin-api/internal/router"
	"github.com/huynh-fs/gin-api/internal/service"
	"github.com/huynh-fs/gin-api/pkg/config"
	"github.com/huynh-fs/gin-api/pkg/database"
)

// @title           Todo List API
// @version         1.0
// @description     Một API đơn giản để quản lý danh sách công việc.
// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "your-secret-api-key" to get access
func main() {
	// tải cấu hình
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Không thể tải file .env: %v", err)
	}

	// kết nối database
	db, err := database.Connect(cfg.DSN())
	if err != nil {
		log.Fatalf("Không thể kết nối database: %v", err)
	}

	// khởi tạo repository
	userRepo := repository.NewGormUserRepository(db)
	refreshTokenRepo := repository.NewGormRefreshTokenRepository(db)
	todoRepo := repository.NewGormTodoRepository(db)

	// khởi tạo service
	todoService := service.NewTodoService(todoRepo)
	authService := service.NewAuthService(userRepo, refreshTokenRepo, cfg)

	// khởi tạo handler
	todoHandler := handler.NewTodoHandler(todoService)
	authHandler := handler.NewAuthHandler(authService)

	// thiết lập router
	r := router.Setup(cfg, todoHandler, authHandler)

	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}