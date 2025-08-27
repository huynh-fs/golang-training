package main

import "fmt"

func practiceVariables() {
    fmt.Println("\n--- Thực hành về Biến (Variables) ---")

    // 1. Khai báo với `var`
    var age int
    fmt.Println("Giá trị zero của biến 'age' (int):", age) // Sẽ in ra 0

    var name string
    fmt.Println("Giá trị zero của biến 'name' (string):", name) // Sẽ in ra chuỗi rỗng ""

    var isStudent bool
    fmt.Println("Giá trị zero của biến 'isStudent' (bool):", isStudent) // Sẽ in ra false

    // Khai báo và khởi tạo giá trị
    var language string = "Go"
    fmt.Println("Ngôn ngữ lập trình:", language)

    // Khai báo nhiều biến cùng lúc
    var bien1, bien2 = 10, "hello"
    fmt.Println("Biến 1:", bien1, "| Biến 2:", bien2)

    // 2. Khai báo ngắn gọn với `:=` (chỉ dùng trong hàm)
    country := "Việt Nam"
    year := 2025
    fmt.Println("Quốc gia:", country, "| Năm:", year)

    // Tái khai báo (redeclare) với `:=`
    // `err` là biến mới, `year` đã tồn tại
    year, err := 2026, "no error"
    fmt.Println("Năm (sau khi tái khai báo):", year, "| Lỗi:", err)
}