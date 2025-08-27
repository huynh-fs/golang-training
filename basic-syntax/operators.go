package main

import "fmt"

func practiceOperators() {
    fmt.Println("\n--- Thực hành về Toán tử (Operators) ---")

    // 1. Toán tử số học
    a, b := 10, 3
    fmt.Println("a + b =", a+b)
    fmt.Println("a - b =", a-b)
    fmt.Println("a * b =", a*b)
    fmt.Println("a / b =", a/b)   // Chia lấy phần nguyên
    fmt.Println("a % b =", a%b)   // Chia lấy dư

    // 2. Toán tử so sánh
    x, y := 5, 8
    fmt.Println("x == y:", x == y)
    fmt.Println("x != y:", x != y)
    fmt.Println("x < y:", x < y)
    fmt.Println("x >= y:", x >= y)

    // 3. Toán tử logic
    isTrue, isFalse := true, false
    fmt.Println("isTrue && isFalse:", isTrue && isFalse) // AND
    fmt.Println("isTrue || isFalse:", isTrue || isFalse) // OR
    fmt.Println("!isTrue:", !isTrue)                      // NOT
}