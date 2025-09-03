package main

import (
	"fmt"
	"github.com/huynh-fs/struct/internal/cli"
	"github.com/huynh-fs/struct/internal/models"
)



func main() {
	var classes []models.Classes
	var students []models.Students

	for {
		cli.PrintMenu()
		luaChon := cli.InputChoice()

		switch luaChon {
		case "1":
			classes = cli.InputClassesInfo(classes)
		case "2":
			students = cli.InputStudentInfo(classes, students)
		case "3":
			cli.DisplayInfo(classes, students)
		case "4":
			fmt.Println("Tạm biệt!")
			return 
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng nhập một số từ 1 đến 4.")
		}
		fmt.Println()
	}
}

