package loggers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func DemonstrateLogrus() {
	fmt.Println("--- Bắt đầu minh họa Logrus ---")
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"event":    "user_registration",
		"username": "gopher",
	}).Info("Một người dùng mới vừa đăng ký.")

	log.WithFields(log.Fields{
		"order_id": "xyz-123",
		"amount":   99.99,
	}).Warn("Phát hiện giao dịch đáng ngờ.")

	log.SetFormatter(&log.TextFormatter{})
	fmt.Println("Logrus là một trong những thư viện structured log phổ biến và linh hoạt nhất.")
	fmt.Println("--- Kết thúc minh họa Logrus ---")
}