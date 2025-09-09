// internal/loggers/slog.go
package loggers

import (
	"fmt"
	"log/slog"
	"os"
)

// DemonstrateSlog shows structured logging using the new standard library's slog package.
func DemonstrateSlog() {
	fmt.Println("--- Bắt đầu minh họa Slog (Go 1.21+) ---")

	// 1. Sử dụng TextHandler để ghi log dễ đọc ra console
	fmt.Println("\n>>> Minh họa TextHandler:")
	textLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	textLogger.Debug("Đang kết nối tới database...", "host", "localhost:5432")
	textLogger.Info("Request đã được xử lý thành công", "method", "POST", "latency_ms", 78)

	// 2. Sử dụng JSONHandler để ghi log dạng JSON, phù hợp cho máy móc
	fmt.Println("\n>>> Minh họa JSONHandler:")
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Warn(
		"API key sắp hết hạn",
		slog.String("key_id", "a1b2c3d4"),
		slog.Int("days_left", 3),
	)
	jsonLogger.Error(
		"Xác thực thất bại",
		slog.Group("request_details", // Nhóm các thuộc tính lại
			slog.String("ip_address", "192.168.1.100"),
			slog.String("user_agent", "Go-http-client/1.1"),
		),
	)

	fmt.Println("\nSlog là giải pháp structured logging chính thức, mạnh mẽ và có sẵn trong Go.")
	fmt.Println("--- Kết thúc minh họa Slog ---")
}