package main

import (
	"log"
	"github.com/huynh-fs/gin-api/pkg/database"
	"github.com/huynh-fs/gin-api/internal/router"
	"github.com/huynh-fs/gin-api/internal/handler"
	"github.com/huynh-fs/gin-api/internal/service"
	_ "github.com/huynh-fs/gin-api/docs" 
)

// @title           Todo List API
// @version         1.0
// @description     Một API đơn giản để quản lý danh sách công việc.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	// kết nối database
	database.Connect()

	// khởi tạo handler với service
	todoService := service.NewTodoService(database.DB)
	todoHandler := handler.NewTodoHandler(todoService)

	// thiết lập router
	r := router.Setup(todoHandler)

	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}