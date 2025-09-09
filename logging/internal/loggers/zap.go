// internal/loggers/zap.go
package loggers

import (
	"fmt"
	"go.uber.org/zap"
)

// DemonstrateZap shows structured, high-performance logging with zap.
func DemonstrateZap() {
	fmt.Println("--- Bắt đầu minh họa Zap ---")
	// Zap cung cấp logger cho môi trường production và development
	// Logger cho production có hiệu năng cao, logger cho development dễ đọc hơn
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Đảm bảo tất cả log được ghi ra trước khi chương trình kết thúc

	logger.Info("Bắt đầu xử lý request",
		zap.String("method", "GET"),
		zap.String("path", "/api/users"),
	)

	logger.Error("Không thể tìm thấy tài nguyên",
		zap.Int("resource_id", 99),
		zap.Error(fmt.Errorf("not found in database")),
	)

	fmt.Println("Zap cũng là một thư viện structured log với hiệu năng cực cao.")
	fmt.Println("--- Kết thúc minh họa Zap ---")
}