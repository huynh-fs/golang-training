// internal/loggers/logrus.go
package loggers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// DemonstrateLogrus shows structured logging with logrus.
func DemonstrateLogrus() {
	fmt.Println("--- Bắt đầu minh họa Logrus ---")
	// Cấu hình logrus để output dưới dạng JSON
	log.SetFormatter(&log.JSONFormatter{})

	// Logrus sử dụng cú pháp WithFields để thêm các trường dữ liệu
	log.WithFields(log.Fields{
		"event":    "user_registration",
		"username": "gopher",
	}).Info("Một người dùng mới vừa đăng ký.")

	log.WithFields(log.Fields{
		"order_id": "xyz-123",
		"amount":   99.99,
	}).Warn("Phát hiện giao dịch đáng ngờ.")

	// Reset lại formatter để không ảnh hưởng các log khác nếu cần
	log.SetFormatter(&log.TextFormatter{})
	fmt.Println("Logrus là một trong những thư viện structured log phổ biến và linh hoạt nhất.")
	fmt.Println("--- Kết thúc minh họa Logrus ---")
}