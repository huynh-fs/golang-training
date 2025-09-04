package model

import "fmt"

type BankAccount struct {
	ID            string
	OwnerName     string
	AccountNumber string
	Balance       float64
}

func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("lỗi: số tiền nạp phải lớn hơn 0")
	}
	acc.Balance += amount
	return nil
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("lỗi: số tiền rút phải lớn hơn 0")
	}
	if amount > acc.Balance {
		return fmt.Errorf("lỗi: số dư không đủ")
	}
	acc.Balance -= amount
	return nil
}

// Giả sử có một kho lưu trữ tài khoản đơn giản.
var accountStore = make(map[string]*BankAccount)

func CreateAccount(acc *BankAccount) {
	accountStore[acc.ID] = acc
}

func GetAccountByID(id string) (*BankAccount, error) {
	acc, found := accountStore[id]
	if !found {
		return nil, fmt.Errorf("lỗi: tài khoản với ID '%s' không tồn tại", id)
	}
	return acc, nil
}