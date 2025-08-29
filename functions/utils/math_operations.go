package utils

import "errors"

// Hàm variadic tính tổng của một số lượng số nguyên bất kỳ.
// (Ví dụ hàm variadic)
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Hàm trả về hai giá trị: tổng và hiệu của hai số.
// (Ví dụ hàm trả về nhiều giá trị)
func AddAndSubtract(a, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// Hàm trả về kết quả phép chia và một lỗi nếu có.
// (Ví dụ hàm trả về nhiều giá trị, bao gồm cả lỗi)
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}