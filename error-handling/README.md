# Máy tính Dòng lệnh (Calculator CLI)

Đây là một dự án máy tính dòng lệnh đơn giản được xây dựng bằng ngôn ngữ Go. Mục đích chính của dự án là để thực hành và minh họa các khái niệm cốt lõi về xử lý lỗi (`error`, `defer`, `panic`, `recover`) và cấu trúc dự án chuẩn trong Go.

## Tính năng

- Thực hiện bốn phép tính cơ bản: cộng (`+`), trừ (`-`), nhân (`*`), và chia (`/`).
- Nhận đầu vào trực tiếp từ dòng lệnh.
- Xử lý lỗi đầu vào và lỗi logic một cách rõ ràng.
- Tuân theo cấu trúc dự án chuẩn của Go (Standard Go Project Layout).

## Cấu trúc dự án

Dự án tuân theo cấu trúc chuẩn để tách biệt rõ ràng giữa logic khởi chạy và logic nghiệp vụ.

```
error-handling/
├── cmd/
│   └── calculator/
│       └── main.go         # Điểm vào của ứng dụng, xử lý input/output
│
├── internal/
│   └── service/
│       └── calculator.go   # Logic nghiệp vụ cốt lõi (phép tính)
│
├── go.mod                  # File quản lý module
└── README.md
```

- **`cmd/calculator/main.go`**: Đây là điểm vào (entry point) của ứng dụng. Nó chịu trách nhiệm đọc và xác thực các tham số từ dòng lệnh, gọi đến lớp `service` để thực hiện công việc, và cuối cùng là in kết quả hoặc lỗi ra cho người dùng.
- **`internal/service/calculator.go`**: Đây là lớp dịch vụ (service layer), chứa "bộ não" của ứng dụng. Toàn bộ logic nghiệp vụ về cách thực hiện một phép tính được đặt ở đây. Việc đặt trong thư mục `internal` đảm bảo rằng gói này không thể được import bởi các dự án bên ngoài.

## Điều kiện tiên quyết

- **Go**: Phiên bản 1.18 trở lên.

## Cài đặt và Chạy

1.  **Clone repository về máy:**
    ```bash
    git clone https://github.com/huynh-fs/golang-training.git
    cd error-handling
    ```

2.  **Chạy ứng dụng:**
    Sử dụng lệnh `go run` và chỉ định đường dẫn đến gói `main`.

    **Cú pháp:**
    ```bash
    go run ./cmd/calculator <số 1> <toán tử> <số 2>
    ```

### Ví dụ sử dụng

- **Phép tính thành công:**
    ```bash
    go run ./cmd/calculator 25 + 17
    ```
    Kết quả:
    ```
    --- Bắt đầu chương trình máy tính CLI ---
    -> Hoàn thành logic tính toán.
    Kết quả: 25 + 17 = 42
    ```

- **Lỗi toán tử không hợp lệ (Xử lý `error`):**
    ```bash
    go run ./cmd/calculator 10 x 3
    ```
    Kết quả:
    ```
    --- Bắt đầu chương trình máy tính CLI ---
    -> Hoàn thành logic tính toán.
    Lỗi: toán tử không hợp lệ, chỉ chấp nhận +, -, *, /
    ```

- **Lỗi chia cho không (Minh họa `panic` và `recover`):**
    ```bash
    go run ./cmd/calculator 100 / 0
    ```
    Kết quả:
    ```
    --- Bắt đầu chương trình máy tính CLI ---
    -> Hoàn thành logic tính toán.
    Lỗi hệ thống: lỗi nghiêm trọng: không thể chia cho số không
    Chương trình đã dừng đột ngột.
    ```

## Các khái niệm Go được minh họa

Dự án này được thiết kế để làm rõ các cơ chế xử lý lỗi và luồng thực thi trong Go:

- **Xử lý `error`**: Thông qua `if err != nil`. Được sử dụng khi xác thực đầu vào của người dùng (ví dụ: `strconv.Atoi`, toán tử không hợp lệ). Đây là cách xử lý lỗi tiêu chuẩn cho các lỗi có thể lường trước.

- **`defer`**: Được sử dụng trong hàm `Calculate` của `service` để cho thấy một hành động (in ra màn hình) luôn được thực thi ngay trước khi hàm kết thúc, bất kể hàm trả về bình thường hay bị `panic`.

- **`panic`**: Được gọi khi gặp một lỗi lập trình không thể phục hồi—trong trường hợp này là phép chia cho số không. `panic` dừng luồng thực thi thông thường và bắt đầu quá trình "nổi" lên trên call stack.

- **`recover`**: Được đặt trong một hàm `defer` ở `main.go`. Nó hoạt động như một "lưới an toàn", bắt lại `panic` để ngăn chương trình bị sập hoàn toàn và cho phép chúng ta in ra một thông báo lỗi thân thiện hơn trước khi thoát.