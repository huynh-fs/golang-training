package loggers

import (
	"log"
	"os"
)

func DemonstrateStdLib() {
	log.Println("--- Bắt đầu minh họa Standard Lib Log ---")

	log.Println("Đây là một log mặc định.")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Log này có thêm thông tin file và dòng code.")

	file, err := os.OpenFile("stdlib.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Không thể mở file:", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Log này được ghi vào file 'stdlib.log'.")

	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags)
	log.Println("--- Kết thúc minh họa Standard Lib Log ---")
}