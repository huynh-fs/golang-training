// cmd/go-logging-demo/main.go
package main

import (
	// Thay đổi "github.com/the_your_username/go-logging-demo" cho đúng với module của bạn
	"github.com/huynh-fs/logging/internal/cli"
)

func main() {
	// Khởi chạy giao diện dòng lệnh
	cli.Run()
}