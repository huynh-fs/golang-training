package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadInput đọc một dòng từ stdin và loại bỏ ký tự xuống dòng.
func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}