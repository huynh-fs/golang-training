package main

import (
	"bufio"
	"fmt"
	"github.com/huynh-fs/method/internal/cli"
	"github.com/huynh-fs/method/internal/model"
	"os"
	"strconv"
	"strings"
)

func printMenu() {
	fmt.Println("\n--- MENU NGÂN HÀNG ---")
	fmt.Println("1. Xem thông tin tài khoản (tài khoản mặc định ID: acc001)")
	fmt.Println("2. Nạp tiền")
	fmt.Println("3. Rút tiền")
	fmt.Println("4. Thoát")
	fmt.Print("Vui lòng chọn một chức năng: ")
}

func main() {
	model.CreateAccount(&model.BankAccount{
		ID:            "acc001",
		OwnerName:     "Bob",
		AccountNumber: "999888777",
		Balance:       5000.0,
	})

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			cli.HandleGetAccount("acc001")

		case "2":
			fmt.Print("Nhập số tiền cần nạp: ")
			scanner.Scan()
			amountStr := strings.TrimSpace(scanner.Text())
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Lỗi: Số tiền không hợp lệ.")
				continue 
			}
			cli.HandleDeposit("acc001", amount)

		case "3":
			fmt.Print("Nhập số tiền cần rút: ")
			scanner.Scan()
			amountStr := strings.TrimSpace(scanner.Text())
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Lỗi: Số tiền không hợp lệ.")
				continue 
			}
			cli.HandleWithdraw("acc001", amount)

		case "4":
			fmt.Println("Cảm ơn bạn đã sử dụng dịch vụ. Tạm biệt!")
			os.Exit(0) 

		default:
			fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
		}
	}
}