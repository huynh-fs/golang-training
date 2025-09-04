package cli

import (
	"fmt"
	"github.com/huynh-fs/method/internal/model"
)

func HandleGetAccount(id string) {
	account, err := model.GetAccountByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("--- Thông Tin Tài Khoản ---")
	fmt.Printf("ID: %s\n", account.ID)
	fmt.Printf("Chủ tài khoản: %s\n", account.OwnerName)
	fmt.Printf("Số dư: %.2f\n", account.Balance)
	fmt.Println("---------------------------")
}

func HandleDeposit(id string, amount float64) {
	account, err := model.GetAccountByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := account.Deposit(amount); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Nạp tiền thành công %.2f.\n", amount)
	fmt.Printf("Số dư mới: %.2f\n", account.Balance)
}

func HandleWithdraw(id string, amount float64) {
	account, err := model.GetAccountByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := account.Withdraw(amount); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Rút tiền thành công %.2f.\n", amount)
	fmt.Printf("Số dư mới: %.2f\n", account.Balance)
}