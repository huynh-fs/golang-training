package main

import (
	"fmt"
	"os"

	"github.com/huynh-fs/channels/internal/app/log-processor/service"
	"github.com/huynh-fs/channels/internal/config"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		fmt.Printf("Lỗi: Không thể tải cấu hình: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Bắt đầu xử lý file '%s' với %d parser workers...\n", "large.log", cfg.Processor.ParserWorkers)

	processorService := service.New(cfg.Processor.ParserWorkers)
	stats, err := processorService.Run("large.log")
	if err != nil {
		fmt.Printf("Lỗi: Xử lý pipeline thất bại: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n--------- BÁO CÁO CUỐI CÙNG -----------")
	fmt.Printf("Tổng số dòng đã phân tích hợp lệ: %d\n", stats.TotalParsed)
	fmt.Println("Thống kê theo cấp độ:")
	for level, count := range stats.CountByLevel {
		fmt.Printf("- %-5s: %d\n", level, count)
	}
	fmt.Println("---------------------------------------")
}