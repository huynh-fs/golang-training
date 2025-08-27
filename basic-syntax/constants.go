package main

import "fmt"

func practiceConstants() {
    fmt.Println("\n--- Thực hành về Hằng số (Constants) ---")

    // 1. Khai báo với `const`
    const Pi float64 = 3.14159
    fmt.Println("Giá trị của Pi là:", Pi)

    // Hằng số không thể thay đổi giá trị
    // Pi = 3.14 // Dòng này sẽ gây lỗi biên dịch: cannot assign to Pi

    // Khai báo nhiều hằng số
    const a, b, c = 3, 4, "foo"
    fmt.Println("a:", a, "| b:", b, "| c:", c)

    // Hằng số không định kiểu (Untyped Constants)
    const untypedInt = 100
    var myInt int = untypedInt           // Gán cho int
    var myFloat float64 = untypedInt    // Gán cho float64
    fmt.Println("Untyped Int trong Int:", myInt, "| Untyped Int trong Float64:", myFloat)

    // 2. Sử dụng `iota`
    const (
        c0 = iota // c0 == 0
        c1 = iota // c1 == 1
        c2 = iota // c2 == 2
    )
    fmt.Println("Sử dụng iota:", c0, c1, c2)

    const (
        Read = 1 << iota // 1 << 0 = 1
        Write = 1 << iota // 1 << 1 = 2
        Execute = 1 << iota // 1 << 2 = 4
    )
    fmt.Println("Quyền truy cập (bit shifting): Read =", Read, "| Write =", Write, "| Execute =", Execute)
}