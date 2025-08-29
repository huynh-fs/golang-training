package main

import (
	"fmt"
	"log"
	"time"

	"github.com/huynh-fs/golang-functions-demo/logger"
	"github.com/huynh-fs/golang-functions-demo/utils"
)

func main() {
	fmt.Println("--- Demo Logger ---")

	// 1. Sử dụng Logger với formatter mặc định
	fmt.Println("\n-- Logger với formatter mặc định --")
	defaultLogger := logger.NewLogger()
	defaultLogger.Info("Đây là một thông tin quan trọng")
	defaultLogger.Warning("Cảnh báo: Có thể có lỗi")
	defaultLogger.Error("Lỗi nghiêm trọng đã xảy ra!")
	defaultLogger.Logf(logger.LevelInfo, "Giá trị của PI là %.2f", 3.14159) // Sử dụng Logf (variadic)

	// 2. Sử dụng Logger với hàm ẩn danh làm formatter tùy chỉnh
	fmt.Println("\n-- Logger với hàm ẩn danh làm formatter --")
	customFormatterLogger := logger.NewLogger(
		logger.WithFormatter(func(level, message string) string { // Hàm ẩn danh được truyền làm tham số
			return fmt.Sprintf("[CUSTOM-FORMATTER][%s] %s -- %s", level, message, time.Now().Format("15:04:05"))
		}),
	)
	customFormatterLogger.Info("Thông báo từ formatter tùy chỉnh")

	// 3. Sử dụng Logger với formatter dựa trên closure
	fmt.Println("\n-- Logger với formatter dựa trên closure --")
	// Tạo một formatter closure với tiền tố "APP-LOG"
	appLogFormatter := logger.CreatePrefixFormatter("APP-LOG") // CreatePrefixFormatter trả về một closure
	closureLogger := logger.NewLogger(
		logger.WithFormatter(appLogFormatter),
	)
	closureLogger.Warning("Thông báo cảnh báo từ closure formatter")
	closureLogger.Error("Lỗi từ closure formatter!")

	fmt.Println("\n--- Demo Utils: Math Operations ---")

	// 1. Sử dụng hàm Variadic (Sum)
	fmt.Println("\n-- Hàm Variadic: Sum --")
	nums1 := []int{10, 20, 30}
	fmt.Printf("Tổng của %v là: %d\n", nums1, utils.Sum(nums1...)) // Truyền slice vào hàm variadic
	fmt.Printf("Tổng của 1, 2, 3, 4, 5 là: %d\n", utils.Sum(1, 2, 3, 4, 5))

	// 2. Sử dụng hàm trả về nhiều giá trị (AddAndSubtract)
	fmt.Println("\n-- Hàm trả về nhiều giá trị: AddAndSubtract --")
	sum, diff := utils.AddAndSubtract(25, 10)
	fmt.Printf("Cho 25 và 10: Tổng = %d, Hiệu = %d\n", sum, diff)

	// 3. Sử dụng hàm trả về nhiều giá trị (bao gồm error) (Divide)
	fmt.Println("\n-- Hàm trả về nhiều giá trị (có error): Divide --")
	result1, err1 := utils.Divide(100, 4)
	if err1 != nil {
		log.Printf("Lỗi khi chia: %v\n", err1)
	} else {
		fmt.Printf("100 chia 4 = %d\n", result1)
	}

	result2, err2 := utils.Divide(100, 0)
	if err2 != nil {
		log.Printf("Lỗi khi chia: %v\n", err2) // Sẽ in lỗi chia cho 0
	} else {
		fmt.Printf("100 chia 0 = %d\n", result2)
	}
}