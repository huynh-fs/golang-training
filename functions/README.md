# Golang Functions Demo

Dự án này là một ví dụ minh họa về cách sử dụng các khái niệm hàm cơ bản và nâng cao trong Golang, bao gồm khai báo hàm, tham số, giá trị trả về (multiple returns), hàm variadic, hàm ẩn danh (anonymous functions) và closure. Dự án được cấu trúc thành các gói riêng biệt để thể hiện cách tổ chức mã trong một ứng dụng Go thực tế.

## Cấu trúc thư mục

```
functions/
├── main.go
├── go.mod
├── README.md
├── logger/
│   ├── logger.go
│   └── formatter.go
└── utils/
    └── math_operations.go
```

## Các khái niệm Golang được minh họa

### 1. Khai báo hàm, tham số, giá trị trả về

*   **`utils/math_operations.go`**:
    *   `AddAndSubtract(a, b int) (int, int)`: Minh họa hàm trả về nhiều giá trị (tổng và hiệu).
    *   `Divide(a, b int) (int, error)`: Minh họa hàm trả về nhiều giá trị, bao gồm cả giá trị lỗi theo chuẩn Go.
*   **`logger/logger.go`**:
    *   `Info(message string)`, `Warning(message string)`, `Error(message string)`: Các hàm cơ bản với một tham số đầu vào.
    *   `NewLogger(options ...LoggerOption) *Logger`: Hàm khởi tạo nhận tham số là các `LoggerOption`.

### 2. Hàm Variadic (Variadic Functions)

*   **`utils/math_operations.go`**:
    *   `Sum(nums ...int) int`: Minh họa cách một hàm có thể chấp nhận một số lượng đối số thay đổi (`...int`).
*   **`logger/logger.go`**:
    *   `Logf(level LogLevel, format string, args ...interface{})`: Minh họa một hàm variadic để ghi log với định dạng tùy chỉnh (tương tự `fmt.Printf`).
    *   `NewLogger(options ...LoggerOption) *Logger`: Sử dụng variadic để chấp nhận nhiều `LoggerOption` (mẫu "Functional Options").

### 3. Hàm ẩn danh (Anonymous Functions) và Closure

*   **`logger/formatter.go`**:
    *   `type Formatter func(level string, message string) string`: Định nghĩa một kiểu hàm, cho phép chúng ta gán các hàm ẩn danh cho các biến kiểu `Formatter`.
    *   `DefaultFormatter() Formatter`: Trả về một hàm ẩn danh làm `Formatter` mặc định.
    *   `CreatePrefixFormatter(prefix string) Formatter`: Minh họa **Closure**. Hàm này trả về một hàm ẩn danh (Formatter) mà hàm ẩn danh đó "nhớ" giá trị của biến `prefix` từ môi trường tạo ra nó.
*   **`logger/logger.go`**:
    *   `WithFormatter(f Formatter) LoggerOption`: Nhận một `Formatter` (thường là hàm ẩn danh hoặc closure) và trả về một `LoggerOption` (cũng là một hàm ẩn danh).
*   **`main.go`**:
    *   Trong `main`, bạn sẽ thấy cách truyền trực tiếp các hàm ẩn danh làm đối số cho `logger.WithFormatter` và cách sử dụng closure được tạo bởi `logger.CreatePrefixFormatter`.

## Cách chạy dự án

Để chạy và khám phá dự án này, hãy làm theo các bước sau:

1.  **Sao chép kho lưu trữ:**
    ```bash
    git clone https://github.com/huynh-fs/golang-training.git
    cd functions
    ```
2.  **Chạy ứng dụng:**
    ```bash
    go run .
    ```

### Output dự kiến

Bạn sẽ thấy một loạt các thông báo log và kết quả tính toán được in ra console, minh họa từng khái niệm về hàm.

<img width="821" height="749" alt="image" src="https://github.com/user-attachments/assets/aba17f91-6e73-421f-9ad4-ce1505ebf811" />
