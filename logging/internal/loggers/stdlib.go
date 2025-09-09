// internal/loggers/stdlib.go
package loggers

import (
	"log"
	"os"
)

// DemonstrateStdLib shows logging using the standard library's log package.
func DemonstrateStdLib() {
	log.Println("--- Bắt đầu minh họa Standard Lib Log ---")

	// Log mặc định ra standard error
	log.Println("Đây là một log mặc định.")

	// Cấu hình log để ghi thêm file và dòng code
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Log này có thêm thông tin file và dòng code.")

	// Chuyển hướng output của log vào một file
	file, err := os.OpenFile("stdlib.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Không thể mở file:", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Log này được ghi vào file 'stdlib.log'.")

	// Reset lại output về standard error để không ảnh hưởng các log khác
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags) // Reset cờ
	log.Println("--- Kết thúc minh họa Standard Lib Log ---")
}