# Go Zoo Management - Advanced Project Layout

Đây là một dự án demo bằng Go, được xây dựng để minh họa một cấu trúc dự án có khả năng mở rộng, bảo trì và kiểm thử cao, tuân thủ các nguyên tắc thiết kế phần mềm hiện đại.

Mục tiêu chính của dự án không phải là xây dựng một ứng dụng phức tạp, mà là cung cấp một **template (khuôn mẫu)** vững chắc cho các dự án Go trong thực tế, thể hiện rõ ràng các khái niệm:

-   Kiến trúc phân lớp (Layered Architecture).
-   Nguyên tắc đảo ngược phụ thuộc (Dependency Inversion Principle).
-   Tách biệt mối quan tâm (Separation of Concerns).
-   Cấu trúc thư mục dự án Go chuẩn.

## ✨ Các khái niệm chính

### Các khái niệm trong Go
-   **Interfaces**: Định nghĩa các "hợp đồng" hành vi (`animal.Creature`).
-   **Struct Embedding**: Tái sử dụng code và thúc đẩy composition over inheritance.
-   **Packages**: Tổ chức code thành các đơn vị logic độc lập.

### Các nguyên tắc kiến trúc
-   **Layered Architecture**: Code được tổ chức thành các lớp riêng biệt (`handler`, `service`) với luồng phụ thuộc một chiều.
-   **Dependency Injection**: Các thành phần phụ thuộc (dependencies) được "tiêm" vào từ bên ngoài (ví dụ: `service` được tiêm vào `handler` trong `main.go`), giúp tăng khả năng kiểm thử.
-   **Clean Architecture**: Lớp nghiệp vụ cốt lõi (`service`) không phụ thuộc vào các chi tiết bên ngoài như giao diện người dùng (`handler`) hay cơ sở dữ liệu.

## 📂 Cấu trúc dự án

Dự án được tổ chức theo một cấu trúc chi tiết để tối ưu hóa việc bảo trì và mở rộng.

```
interface-embedding/
├── cmd/
│   └── zoo/
│       └── main.go         # Entrypoint: Lắp ráp dependencies và khởi chạy ứng dụng.
├── internal/
│   ├── app/
│   │   └── zoo/
│   │       ├── handler/    # Lớp giao tiếp: Xử lý input/output (CLI, HTTP, gRPC...).
│   │       │   └── cli_handler.go
│   │       └── service/    # Lớp nghiệp vụ: Chứa business logic chính.
│   │           └── zoo_service.go
│   └── pkg/                # Các thư viện private, chỉ dùng trong nội bộ dự án.
|
├── pkg/
│   ├── animal/             # Package định nghĩa các interface công khai.
│   │   └── creature.go
│   └── creatures/          # Package chứa các hiện thực cụ thể của interface.
│       ├── bird.go
│       └── dog.go
│       └── snake.go
└── go.mod
```

-   **/cmd**: Điểm khởi đầu của các file thực thi. `main.go` tại đây đóng vai trò là **Composition Root**, nơi các thành phần của ứng dụng được khởi tạo và kết nối với nhau.
-   **/pkg**: Các thư viện **công khai**, an toàn để các dự án khác có thể import và sử dụng. Đây là nơi lý tưởng để định nghĩa các `interface` chung và các `struct` hiện thực chúng.
-   **/internal**: Logic nghiệp vụ và các thư viện **riêng tư** của dự án. Go sẽ ngăn không cho các dự án khác import code từ đây.
    -   `/app`: Chứa code cho các ứng dụng cụ thể.
        -   `handler`: Lớp ngoài cùng, chịu trách nhiệm giao tiếp với thế giới bên ngoài (ví dụ: nhận lệnh từ CLI, xử lý request HTTP). Nó gọi đến `service` để thực hiện công việc.
        -   `service`: Trái tim của ứng dụng, chứa toàn bộ business logic. Nó không biết gì về `handler` và hoàn toàn độc lập.

## 🚀 Bắt đầu

### Yêu cầu

-   Go phiên bản 1.18 trở lên.

### Kết quả mong đợi

Output trên terminal của bạn sẽ là:

```
CHÀO MỪNG ĐẾN VỚI SỞ THÚ GO!
==============================
--- Báo cáo về Chó Mực ---
Tiếng kêu: Gâu gâu!
Cách di chuyển: Chạy bằng bốn chân.

--- Báo cáo về Rắn Hổ Mang ---
Tiếng kêu: Xì xì!
Cách di chuyển: Lê trên mặt đất

--- Báo cáo về Chim Chích Chòe ---
Tiếng kêu: Chíp chíp!
Cách di chuyển: Bay bằng đôi cánh

==============================
```