package main

import (
	"database/sql"
	"log"

	"github.com/huynh-fs/sql/internal/cli"
	"github.com/huynh-fs/sql/internal/config"
	"github.com/huynh-fs/sql/internal/database"
	"github.com/huynh-fs/sql/internal/service"
)

func main() {
	// 1. Tải cấu hình
	cfg := config.LoadConfig()

	// 2. Kết nối CSDL
	db := database.ConnectDB(*cfg)
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Printf("Lỗi khi đóng kết nối DB: %v", err)
		}
	}(db)

	// 3. Khởi tạo Service (Lớp nghiệp vụ)
	productService := service.NewProductService(db)
	orderService := service.NewOrderService(db)

	// 4. Khởi tạo và chạy giao diện CLI (Menu)
	// Tất cả logic menu và I/O đều nằm trong đối tượng cliApp.
	cliApp := cli.NewCLI(productService, orderService)

	log.Println("Khởi động ứng dụng CLI...")
	cliApp.Run()
}