// internal/cli/cli.go
package cli

import (
	"fmt"
	"os"
	// Thay đổi "github.com/the_your_username/go-logging-demo" cho đúng với module của bạn
	"github.com/huynh-fs/logging/internal/loggers"
)

// Run starts the command-line interface.
func Run() {
	for {
		printMenu()
		choice := getUserInput()

		switch choice {
		case 1:
			loggers.DemonstrateStdLib()
		case 2:
			loggers.DemonstrateZerolog()
		case 3:
			loggers.DemonstrateZap()
		case 4:
			loggers.DemonstrateLogrus()
		case 5: // Thêm case mới cho Slog
			loggers.DemonstrateSlog()
		case 6: // Cập nhật số để thoát
			fmt.Println("Tạm biệt!")
			os.Exit(0)
		default:
			fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
		}
		fmt.Println("\n========================================")
	}
}

func printMenu() {
	fmt.Println("Chọn một thư viện để minh họa ghi log:")
	fmt.Println("1. Standard Library (log)")
	fmt.Println("2. Zerolog")
	fmt.Println("3. Zap")
	fmt.Println("4. Logrus")
	fmt.Println("5. Slog (Standard Library, Go 1.21+)") // Thêm lựa chọn Slog
	fmt.Println("6. Thoát") // Cập nhật lựa chọn thoát
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func getUserInput() int {
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return 0 // Trả về 0 nếu input không phải là số
	}
	return choice
}