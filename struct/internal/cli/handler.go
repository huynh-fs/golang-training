package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/huynh-fs/struct/internal/models"
)

func PrintMenu() {
	fmt.Println("===== MENU HỆ THỐNG QUẢN LÝ =====")
	fmt.Println("1. Thêm Lớp Học Mới")
	fmt.Println("2. Thêm Học Sinh Mới")
	fmt.Println("3. Xuất Báo Cáo Thông Tin")
	fmt.Println("4. Thoát Chương Trình")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func InputChoice() string {
	reader := bufio.NewReader(os.Stdin)
	luaChon, _ := reader.ReadString('\n')
	return strings.TrimSpace(luaChon)
}

func InputClassesInfo(currentClasses []models.Classes) []models.Classes {
	classes := currentClasses
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- NHẬP THÔNG TIN LỚP HỌC ---")
	fmt.Println("(Nhập tên lớp, sau đó nhấn Enter. Bỏ trống và nhấn Enter để kết thúc)")

	for {
		fmt.Print("Nhập tên lớp: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			break
		}

		var isExist bool
		for _, class := range classes {
			if class.Name == name {
				isExist = true
				break
			}
		}

		if isExist {
			fmt.Printf("Lớp học %s đã tồn tại. Vui lòng nhập lớp học khác.\n", name)
			continue
		}

		class := models.Classes{Name: name, NumOfStudents: 0}
		classes = append(classes, class)
		fmt.Printf("Đã thêm thành công lớp học: %s\n", name)
	}
	fmt.Println("--- NHẬP XONG THÔNG TIN LỚP HỌC ---")
	return classes
}

func InputStudentInfo(classes []models.Classes, currentStudent []models.Students) []models.Students {
	if len(classes) == 0 {
		fmt.Println("Lỗi: Chưa có lớp học nào. Vui lòng thêm lớp học trước khi thêm sinh viên.")
		return currentStudent
	}

	students := currentStudent
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("--- NHẬP THÔNG TIN SINH VIÊN ---")
	fmt.Println("(Nhập tên sinh viên, sau đó nhấn Enter. Bỏ trống và nhấn Enter để kết thúc)")

	for {
		fmt.Print("Nhập tên sinh viên: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			break
		}

		var pClass *models.Classes
		for {
			fmt.Print("Nhập tên lớp học của sinh viên: ")
			className, _ := reader.ReadString('\n')
			className = strings.TrimSpace(className)
			if className == "" {
				break
			}

			var found = false
			for i := range classes {
				if classes[i].Name == className {
					pClass = &classes[i]
					found = true
					break
				}
			}

			if found {
				break
			} else {
				fmt.Printf("Lỗi: Lớp học `%s` không tồn tại. Vui lòng nhập lại.\n", className)
			}
		}

		student := models.Students{Name: name, Class: pClass}
		students = append(students, student)
	}

	return students
}

func DisplayInfo(classes []models.Classes, students []models.Students) {
	studentGroupByClass := make(map[*models.Classes][]models.Students)
	if len(classes) == 0 {
		fmt.Println("Chưa có lớp học nào để hiển thị.")
		return
	}
	for _, student := range students {
		studentGroupByClass[student.Class] = append(studentGroupByClass[student.Class], student)
	}

	for i := range classes {
		class := &classes[i]
		
		studentsInClass := studentGroupByClass[class]
		class.NumOfStudents = len(studentsInClass)

		fmt.Printf("-----Lớp: %s | Số lượng sinh viên: %d -----\n", class.Name, class.NumOfStudents)
		if class.NumOfStudents == 0 {
			fmt.Println("Lớp chưa có sinh viên nào.")
		} else {
			for i, student := range studentsInClass {
				fmt.Printf("%d. %s\n", i+1, student.Name)
			}
		}
	}
}