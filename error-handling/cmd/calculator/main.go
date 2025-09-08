package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/huynh-fs/error-handling/internal/service"
)

func main() {
	// để bắt các panic từ toàn bộ ứng dụng.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Lỗi hệ thống:", r)
			fmt.Println("Chương trình đã dừng đột ngột.")
			os.Exit(1)
		}
	}()

	fmt.Println("--- Bắt đầu chương trình máy tính CLI ---")

	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Sử dụng: go run ./cmd/calculator <số 1> <toán tử> <số 2>")
		return
	}

	// idiomatic go (if err != nil): xử lý lỗi khi parse input
	a, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Lỗi input: '%s' không phải là một số hợp lệ.\n", args[0])
		return
	}

	b, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Lỗi input: '%s' không phải là một số hợp lệ.\n", args[2])
		return
	}

	operator := args[1]

	result, err := calculator.Calculate(a, b, operator)
	if err != nil {
		fmt.Println("Lỗi:", err)
		return
	}

	fmt.Printf("Kết quả: %d %s %d = %d\n", a, operator, b, result)
}