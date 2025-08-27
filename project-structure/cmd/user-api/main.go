// file: cmd/user-api/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	// Import các package trong dự án
	"github.com/huynh-fs/golang-training/user-service/internal/handler"
	"github.com/huynh-fs/golang-training/user-service/internal/repository"
	"github.com/huynh-fs/golang-training/user-service/internal/service"
	"github.com/huynh-fs/golang-training/user-service/pkg/config"
)

func main() {
	// 1. Tải cấu hình
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Không thể đọc file cấu hình: %v", err)
	}

	// 2. Khởi tạo các lớp (Dependency Injection)
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 3. Đăng ký handler cho endpoint
	http.HandleFunc("/users/", userHandler.GetUser)

	// 4. Khởi động server
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server đang lắng nghe tại http://localhost%s", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Không thể khởi động server: %v", err)
	}
}