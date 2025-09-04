# Hệ thống Quản lý Lớp học và Học sinh (CLI)

Đây là một ứng dụng dòng lệnh (CLI - Command-Line Interface) được xây dựng bằng ngôn ngữ Go. Dự án này là một ví dụ thực hành về việc áp dụng **Kiến trúc Phân lớp (Layered Architecture)**, thể hiện rõ ràng việc phân tách trách nhiệm (Separation of Concerns) giữa các thành phần khác nhau của hệ thống.

Mục tiêu chính không chỉ là giải quyết bài toán quản lý, mà còn là để minh họa một cấu trúc dự án có khả năng bảo trì, mở rộng và kiểm thử tốt.

## Kiến trúc

Dự án được cấu trúc theo mô hình Kiến trúc Phân lớp, bao gồm 3 tầng chính. Luồng điều khiển (Control Flow) của ứng dụng đi theo một hướng duy nhất từ ngoài vào trong:

**`main` -> `Handler Layer` -> `Service Layer` -> `Domain Layer`**

1.  **Domain Layer (Tầng Miền)**: Lớp trong cùng, là nền tảng của toàn bộ ứng dụng. Nó chỉ chứa định nghĩa về các đối tượng nghiệp vụ cốt lõi dưới dạng các cấu trúc dữ liệu thuần túy (`struct` không có method). Ví dụ: `Classes`, `Students`.
2.  **Service Layer (Tầng Dịch vụ/Nghiệp vụ)**: Là trái tim của ứng dụng, chứa toàn bộ logic nghiệp vụ. Nó điều phối các đối tượng `Domain` để thực hiện các tác vụ. Tầng này chịu trách nhiệm đảm bảo tính toàn vẹn và nhất quán của dữ liệu (ví dụ: tự cập nhật sĩ số lớp khi thêm học sinh).
3.  **Handler Layer (Tầng Xử lý/Giao tiếp)**: Lớp ngoài cùng, là cầu nối giữa người dùng và ứng dụng. Trong dự án này, nó là một `CLI Handler` xử lý input/output từ console. Nó nhận yêu cầu, gọi đến `Service` tương ứng, và trình bày kết quả cho người dùng.

## Các Quyết định Thiết kế Quan trọng

1.  **Một `SchoolService` duy nhất**: Thay vì tách thành nhiều service nhỏ (`ClassService`, `StudentService`), dự án sử dụng một `SchoolService` duy nhất để quản lý toàn bộ dữ liệu (`classes` và `students`). Quyết định này đảm bảo có một **Nguồn Chân lý Duy nhất (Single Source of Truth)**, tránh được các vấn đề phức tạp về đồng bộ dữ liệu giữa các service khác nhau.
2.  **`Domain` Models là các Cấu trúc Dữ liệu Thuần túy**: Các `struct` trong tầng `domain` (`Classes`, `Students`) không chứa logic nghiệp vụ (không có methods). Chúng chỉ đóng vai trò là các "data container". Toàn bộ hành vi và logic được chuyển lên tầng `Service`, giúp phân tách rõ ràng giữa "dữ liệu là gì" và "bạn có thể làm gì với dữ liệu".

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
│   │           └── school_service.go # Tầng Nghiệp vụ: Chứa business logic.
│   └── domain/
│       ├── class.go             # Tầng Miền: Định nghĩa struct Classes.
│       └── student.go           # Tầng Miền: Định nghĩa struct Students.
└── go.mod
```

-   **`/cmd`**: Chứa điểm khởi đầu của ứng dụng. `main.go` có nhiệm vụ thiết lập và kết nối các tầng lại với nhau (Dependency Injection).
-   **`/internal/domain`**: Định nghĩa các thực thể cốt lõi.
-   **`/internal/app/school-manager/service`**: Thực thi các quy trình nghiệp vụ và quản lý trạng thái của ứng dụng.
-   **`/internal/app/school-manager/handler`**: Dịch các lệnh từ console thành các lời gọi đến `service`.

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