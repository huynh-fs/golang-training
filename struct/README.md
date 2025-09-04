# Hệ thống Quản lý Lớp học và Học sinh (CLI)

Đây là một ứng dụng dòng lệnh (CLI - Command-Line Interface) được xây dựng bằng ngôn ngữ Go. Dự án này là một ví dụ thực hành về việc áp dụng **Kiến trúc Phân lớp (Layered Architecture)** vào một ứng dụng Go, thể hiện rõ ràng việc phân tách trách nhiệm (Separation of Concerns) giữa các thành phần khác nhau của hệ thống.

Mục tiêu chính không chỉ là giải quyết bài toán quản lý lớp học, mà còn là để minh họa một cấu trúc dự án có khả năng bảo trì, mở rộng và kiểm thử tốt, sẵn sàng cho các yêu cầu phức tạp trong tương lai.

## Kiến trúc

Dự án được cấu trúc theo mô hình Kiến trúc Phân lớp, bao gồm 3 tầng chính:

1.  **Domain Layer (Tầng Miền)**: Lớp trong cùng, chứa định nghĩa về các đối tượng nghiệp vụ cốt lõi (`Classes`, `Students`) và các hành vi (methods) gắn liền với chúng. Tầng này không phụ thuộc vào bất kỳ tầng nào khác.
2.  **Service Layer (Tầng Dịch vụ/Nghiệp vụ)**: Chứa logic nghiệp vụ của ứng dụng. Nó điều phối các đối tượng `Domain` để thực hiện các tác vụ được yêu cầu từ tầng bên ngoài. Tầng này không biết về cách dữ liệu được trình bày (CLI, HTTP, etc.).
3.  **Handler Layer (Tầng Xử lý/Giao tiếp)**: Lớp ngoài cùng, chịu trách nhiệm tương tác với "thế giới bên ngoài". Trong dự án này, nó là một `CLI Handler` xử lý input/output từ console. Tầng này nhận yêu cầu, gọi đến `Service` tương ứng, và trình bày kết quả cho người dùng.

Luồng điều khiển (Control Flow) của ứng dụng đi theo một hướng duy nhất:
`main` -> `Handler Layer` -> `Service Layer` -> `Domain Layer`

## Cấu trúc Thư mục

```
struct/
├── cmd/
│   └── school-manager/
│       └── main.go              # Điểm khởi đầu. Khởi tạo và "tiêm" các phụ thuộc.
├── internal/
│   ├── app/
│   │   └── school-manager/
│   │       ├── handler/
│   │       │   └── cli_handler.go # Tầng Giao tiếp: Tương tác với console.
│   │       └── service/
│   │           └── class_service.go # Tầng Nghiệp vụ: Chứa business logic.
│   └── domain/
│       ├── class.go             # Tầng Miền: Định nghĩa struct Classes và methods.
│       └── student.go           # Tầng Miền: Định nghĩa struct Students.
└── go.mod
```

-   **`/cmd`**: Chứa điểm khởi đầu (entrypoint) của ứng dụng. Vai trò của `main.go` ở đây là thiết lập và kết nối các tầng lại với nhau (Dependency Injection).
-   **`/internal/domain`**: Định nghĩa các thực thể cốt lõi của bài toán.
-   **`/internal/app/school-manager/service`**: Thực thi các quy trình nghiệp vụ. Ví dụ: logic kiểm tra lớp tồn tại, logic thêm học sinh vào lớp.
-   **`/internal/app/school-manager/handler`**: Cầu nối giữa người dùng và ứng dụng. Nó dịch các lệnh từ console thành các lời gọi đến `service`.

## Lợi ích của Kiến trúc này

-   **Phân tách Trách nhiệm Rõ ràng**: Mỗi tầng có một nhiệm vụ duy nhất, giúp code dễ hiểu và dễ quản lý.
-   **Khả năng Kiểm thử (Testability)**: Tầng `service` và `domain` có thể được kiểm thử (unit test) một cách độc lập mà không cần đến giao diện người dùng, giúp đảm bảo logic nghiệp vụ luôn chính xác.
-   **Khả năng Thay thế và Mở rộng**: Nếu trong tương lai muốn chuyển từ CLI sang một giao diện Web API, chúng ta chỉ cần tạo một `http_handler` mới và cho nó gọi đến cùng `ClassService` đã có. Toàn bộ logic nghiệp vụ ở tầng `service` và `domain` được **tái sử dụng hoàn toàn**.

## Cài đặt và Chạy chương trình

### Yêu cầu
-   Cần cài đặt **Go** (phiên bản 1.18 trở lên). Bạn có thể tải về tại [https://go.dev/dl/](https://go.dev/dl/).

### Hướng dẫn
1.  Clone hoặc tải về toàn bộ thư mục dự án này.
2.  Mở terminal (Command Prompt, PowerShell, hoặc Terminal trên Linux/macOS).
3.  Sử dụng lệnh `cd` để di chuyển vào thư mục gốc của dự án (`school-manager`).
4.  Chạy ứng dụng bằng lệnh sau:
    ```sh
    go run ./cmd/school-manager
    ```
5.  Chương trình sẽ khởi động và hiển thị menu để bạn bắt đầu tương tác.