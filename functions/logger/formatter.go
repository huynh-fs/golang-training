package logger

import (
	"fmt"
	"time"
)


// Nhận cấp độ log (string) và thông báo (string), trả về chuỗi đã định dạng.
type Formatter func(level string, message string) string

// Trả về một Formatter mặc định.
func DefaultFormatter() Formatter {
	return func(level, message string) string {
		return fmt.Sprintf("[%s][%s] %s", time.Now().Format(time.DateTime), level, message)
	}
}

// Trả về một Formatter được tùy chỉnh với một tiền tố.
// Hàm ẩn danh được trả về "nhớ" giá trị của 'prefix'.
func CreatePrefixFormatter(prefix string) Formatter {
	return func(level, message string) string { // Hàm ẩn danh là một closure
		return fmt.Sprintf("[%s][%s][%s] %s", prefix, time.Now().Format(time.TimeOnly), level, message)
	}
}