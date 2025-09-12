package handler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type CLIHandler struct {
	reader io.Reader
	writer io.Writer
}

func NewCLIHandler() *CLIHandler {
	return &CLIHandler{
		reader: os.Stdin,
		writer: os.Stdout,
	}
}

func (h *CLIHandler) Run() {
	scanner := bufio.NewScanner(h.reader)

	for {
		h.showMenu()
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		if choice == "5" {
			fmt.Fprintln(h.writer, "Tạm biệt!")
			return
		}

		h.handleChoice(choice, scanner)
		fmt.Fprintln(h.writer, "\nNhấn Enter để tiếp tục...")
		scanner.Scan() 
	}
}

func (h *CLIHandler) showMenu() {
	fmt.Fprintln(h.writer, "\n======================================")
	fmt.Fprintln(h.writer, "   Trình Chạy Unit Test Tương Tác")
	fmt.Fprintln(h.writer, "======================================")
	fmt.Fprintln(h.writer, "1. Chạy tất cả các test (go test ./...)")
	fmt.Fprintln(h.writer, "2. Chạy test với output chi tiết (go test -v ./...)")
	fmt.Fprintln(h.writer, "3. Chạy test và xem độ bao phủ code (go test -cover ./...)")
	fmt.Fprintln(h.writer, "4. Chạy một test cụ thể (go test -run 'TestName' ./...)")
	fmt.Fprintln(h.writer, "5. Thoát")
	fmt.Fprint(h.writer, "Nhập lựa chọn của bạn: ")
}

func (h *CLIHandler) handleChoice(choice string, scanner *bufio.Scanner) {
	switch choice {
	case "1":
		fmt.Fprintln(h.writer, "\n--- Chạy tất cả các test ---")
		h.runGoTest()
	case "2":
		fmt.Fprintln(h.writer, "\n--- Chạy test với output chi tiết (-v) ---")
		h.runGoTest("-v")
	case "3":
		fmt.Fprintln(h.writer, "\n--- Chạy test và xem độ bao phủ code (-cover) ---")
		h.runGoTest("-cover")
	case "4":
		fmt.Fprint(h.writer, "Nhập tên test cần chạy (ví dụ: TestCreateTask): ")
		scanner.Scan()
		testName := strings.TrimSpace(scanner.Text())
		fmt.Fprintf(h.writer, "\n--- Chạy test với tên '%s' (-run) ---\n", testName)
		h.runGoTest("-v", "-run", testName)
	default:
		fmt.Fprintln(h.writer, "Lựa chọn không hợp lệ, vui lòng thử lại.")
	}
}

// thực thi lệnh go test với các cờ được cung cấp.
func (h *CLIHandler) runGoTest(flags ...string) {
	args := []string{"test"}
	args = append(args, flags...)
	args = append(args, "./...")

	cmd := exec.Command("go", args...)
	cmd.Stdout = h.writer
	cmd.Stderr = h.writer

	fmt.Fprintf(h.writer, "Đang thực thi: go %s\n", strings.Join(args, " "))
	fmt.Fprintln(h.writer, "--------------------------------------")

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(h.writer, "Lỗi khi thực thi lệnh: %v\n", err)
	}

	fmt.Fprintln(h.writer, "--------------------------------------")
	fmt.Fprintln(h.writer, "Thực thi hoàn tất.")
}