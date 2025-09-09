package loggers

import (
	"fmt"
	"go.uber.org/zap"
)

func DemonstrateZap() {
	fmt.Println("--- Bắt đầu minh họa Zap ---")
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // đảm bảo tất cả log được ghi ra trước khi exit

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