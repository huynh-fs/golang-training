package main

import "fmt"

func practiceDataTypes() {
    fmt.Println("\n--- Thực hành về Kiểu dữ liệu (Data Types) ---")

    // Kiểu Boolean (bool)
    var isLearning bool = true
    fmt.Printf("Bạn đang học Go? %t (Kiểu: %T)\n", isLearning, isLearning)

    // Kiểu Chuỗi (string)
    var greeting string = "Xin chào Go!"
    fmt.Printf("Chuỗi: %s (Kiểu: %T)\n", greeting, greeting)

    // Kiểu Số nguyên (int)
    var myInt int = -100 // Kích thước phụ thuộc hệ thống
    var myUint uint = 100 // Không dấu
    var myByte byte = 'A' // Bí danh cho uint8
    fmt.Printf("Số nguyên: %d (Kiểu: %T)\n", myInt, myInt)
    fmt.Printf("Số nguyên không dấu: %d (Kiểu: %T)\n", myUint, myUint)
    fmt.Printf("Byte (ký tự A): %d, dưới dạng ký tự: %c (Kiểu: %T)\n", myByte, myByte, myByte)


    // Kiểu Rune (int32) - cho ký tự Unicode
    var myRune rune = '♥'
    fmt.Printf("Rune: %c, mã Unicode: %U (Kiểu: %T)\n", myRune, myRune, myRune)

    // Kiểu Số thực (float)
    var pi_float32 float32 = 3.14
    var pi_float64 float64 = 3.1415926535
    fmt.Printf("Float32: %f (Kiểu: %T)\n", pi_float32, pi_float32)
    fmt.Printf("Float64: %f (Kiểu: %T)\n", pi_float64, pi_float64)

    // Kiểu Số phức (complex)
    var c complex128 = complex(5, 10) // 5 + 10i
    fmt.Printf("Số phức: %v (Kiểu: %T)\n", c, c)
}