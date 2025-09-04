package handler

import (
	"bufio"
	"fmt"
	"os"
	"github.com/huynh-fs/struct/internal/app/school-manager/service"
	"strings"
)

type CLIHandler struct {
	classService *service.ClassService 
}

func NewCLIHandler(cs *service.ClassService) *CLIHandler {
	return &CLIHandler{
		classService: cs,
	}
}

func (h *CLIHandler) Run() {
	for {
		h.printMenu()
		choice := h.inputChoice()

		switch choice {
		case "1":
			h.handleAddClass()
		case "2":
			h.handleAddStudent()
		case "3":
			h.handleDisplayInfo()
		case "4":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ.")
		}
		fmt.Println()
	}
}

func (h *CLIHandler) handleAddClass() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- THÊM LỚP HỌC MỚI ---")
	fmt.Print("Nhập tên lớp: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if err := h.classService.CreateClass(name); err != nil {
		fmt.Printf("Lỗi: %v\n", err)
	} else {
		fmt.Printf("-> Đã thêm thành công lớp '%s'.\n", name)
	}
}

func (h *CLIHandler) handleAddStudent() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- THÊM HỌC SINH MỚI ---")
	fmt.Print("Nhập tên học sinh: ")
	studentName, _ := reader.ReadString('\n')
	studentName = strings.TrimSpace(studentName)

	fmt.Printf("Nhập lớp cho học sinh '%s': ", studentName)
	className, _ := reader.ReadString('\n')
	className = strings.TrimSpace(className)

	if err := h.classService.AddStudentToClass(studentName, className); err != nil {
		fmt.Printf("Lỗi: %v\n", err)
	} else {
		fmt.Printf("-> Đã thêm thành công học sinh '%s' vào lớp '%s'.\n", studentName, className)
	}
}

func (h *CLIHandler) handleDisplayInfo() {
	classes := h.classService.GetAllClasses()
	fmt.Println("\n======= TỔNG HỢP THÔNG TIN =======")
	if len(classes) == 0 {
		fmt.Println("Chưa có dữ liệu.")
		return
	}
	for _, class := range classes {
		class.Display() 
	}
}

func (h *CLIHandler) printMenu() {
	fmt.Println("===== MENU HỆ THỐNG QUẢN LÝ =====")
	fmt.Println("1. Thêm Lớp Học Mới")
	fmt.Println("2. Thêm Học Sinh Mới")
	fmt.Println("3. Xuất Báo Cáo Thông Tin")
	fmt.Println("4. Thoát Chương Trình")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func (h *CLIHandler) inputChoice() string {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}