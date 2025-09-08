package calculator

import (
	"errors"
	"fmt"
)

func Calculate(a, b int, operator string) (int, error) {
	// defer thực thi trước khi hàm kết thúc.
	defer fmt.Println("-> Hoàn thành logic tính toán.")

	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			// lỗi nghiêm trọng
			panic("lỗi nghiêm trọng: không thể chia cho số không")
		}
		return a / b, nil
	default:
		// lỗi có thể dự đoán được
		return 0, errors.New("toán tử không hợp lệ, chỉ chấp nhận +, -, *, /")
	}
}