# Dự án Thực hành Go Cơ bản

Đây là một dự án đơn giản được tạo ra để thực hành và minh họa các khái niệm cơ bản nhất trong ngôn ngữ lập trình Go. Dự án này phù hợp cho những người mới bắt đầu học Go và muốn có một ví dụ trực quan về cách sử dụng biến, hằng số, các kiểu dữ liệu và toán tử.

## 🎯 Các Khái Niệm Được Đề Cập

Dự án này bao gồm các file riêng biệt, mỗi file tập trung vào một nhóm khái niệm cụ thể:

**1. Biến (Variables)** - `variables.go`
    *   Khai báo biến với từ khóa `var`.
    *   Khai báo ngắn gọn với toán tử `:=`.
    *   Khái niệm "Giá trị Zero" (Zero Value) của các kiểu dữ liệu.
    *   Tái khai báo (redeclaration) trong cùng một khối lệnh.

**2. Hằng số (Constants)** - `constants.go`
    *   Khai báo hằng số với từ khóa `const`.
    *   Hằng số không định kiểu (Untyped Constants) và tính linh hoạt của chúng.
    *   Sử dụng `iota` để tạo các hằng số tăng dần.

**3. Kiểu dữ liệu cơ bản (Data Types)** - `datatypes.go`
    *   Kiểu **Boolean** (`bool`).
    *   Kiểu **Chuỗi** (`string`) và tính bất biến (immutable).
    *   Các kiểu **Số** (Numeric Types):
        *   Số nguyên (`int`, `uint`, `byte`, `rune`).
        *   Số thực (`float32`, `float64`).
        *   Số phức (`complex128`).

**4. Toán tử (Operators)** - `operators.go`
    *   Toán tử số học (`+`, `-`, `*`, `/`, `%`).
    *   Toán tử so sánh (`==`, `!=`, `<`, `>`).
    *   Toán tử logic (`&&`, `||`, `!`).

## 📂 Cấu Trúc Dự Án

Dự án được tổ chức một cách rõ ràng để dễ dàng theo dõi và mở rộng.
```text
basic-syntax/
├── go.mod # File quản lý module của Go
├── main.go # Điểm khởi đầu của chương trình, gọi các hàm thực hành
├── variables.go # Mã thực hành về biến
├── constants.go # Mã thực hành về hằng số
├── datatypes.go # Mã thực hành về các kiểu dữ liệu
├── operators.go # Mã thực hành về các toán tử
└── README.md # Tài liệu hướng dẫn dự án
```
## 🚀 Hướng Dẫn Chạy Dự Án

Để chạy dự án này, bạn cần [cài đặt Go](https://go.dev/doc/install) trên máy tính của mình.

#### 1.  **Clone hoặc tải về dự án này.**
#### 2.  **Mở Terminal hoặc Command Prompt.**
#### 3.  **Di chuyển đến thư mục gốc của dự án (`basic-syntax`).**
```bash
cd /đường/dẫn/đến/go_practice
```
#### 4.  **Thực thi lệnh sau:**
```bash
go run .
```
Lệnh `go run .` sẽ tự động biên dịch và chạy tất cả các file `.go` trong thư mục hiện tại.

## 📋 Kết Quả Đầu Ra Mẫu

Sau khi chạy thành công, bạn sẽ thấy kết quả tương tự như sau trên màn hình console:
```
--- Bắt đầu thực hành Go Lang ---
--- Thực hành về Biến (Variables) ---
Giá trị zero của biến 'age' (int): 0
Giá trị zero của biến 'name' (string):
Giá trị zero của biến 'isStudent' (bool): false
Ngôn ngữ lập trình: Go
Biến 1: 10 | Biến 2: hello
Quốc gia: Việt Nam | Năm: 2025
Năm (sau khi tái khai báo): 2026 | Lỗi: no error
--- Thực hành về Hằng số (Constants) ---
Giá trị của Pi là: 3.14159
a: 3 | b: 4 | c: foo
Untyped Int trong Int: 100 | Untyped Int trong Float64: 100
Sử dụng iota: 0 1 2
Quyền truy cập (bit shifting): Read = 1 | Write = 2 | Execute = 4
--- Thực hành về Kiểu dữ liệu (Data Types) ---
Bạn đang học Go? true (Kiểu: bool)
Chuỗi: Xin chào Go! (Kiểu: string)
Số nguyên: -100 (Kiểu: int)
Số nguyên không dấu: 100 (Kiểu: uint)
Byte (ký tự A): 65, dưới dạng ký tự: A (Kiểu: uint8)
Rune: ♥, mã Unicode: U+2665 (Kiểu: int32)
Float32: 3.140000 (Kiểu: float32)
Float64: 3.141593 (Kiểu: float64)
Số phức: (5+10i) (Kiểu: complex128)
--- Thực hành về Toán tử (Operators) ---
a + b = 13
a - b = 7
a * b = 30
a / b = 3
a % b = 1
x == y: false
x != y: true
x < y: true
x >= y: false
isTrue && isFalse: false
isTrue || isFalse: true
!isTrue: false
--- Kết thúc thực hành ---
```

<img width="669" height="761" alt="image" src="https://github.com/user-attachments/assets/9cd182dc-d5c9-4dbd-aa43-9d360479bbd4" />
