# Hệ thống Quản lý Lớp học và Học sinh (CLI)

Đây là một ứng dụng dòng lệnh (CLI - Command-Line Interface) được xây dựng bằng ngôn ngữ Go. Dự án này không chỉ là một công cụ quản lý đơn giản mà còn là một ví dụ điển hình về cách cấu trúc một dự án Go theo layout chuẩn mực, thể hiện rõ ràng việc tách biệt các thành phần (separation of concerns).

Ứng dụng này là một tài liệu tham khảo tuyệt vời cho các lập trình viên Go muốn xây dựng các ứng dụng có khả năng bảo trì và mở rộng tốt.

## Tính năng

- **Thêm Lớp học Mới**: Cho phép người dùng thêm nhiều lớp học vào hệ thống.
- **Thêm Học sinh Mới**: Cho phép thêm học sinh và liên kết học sinh đó với một lớp học đã tồn tại.
- **Hiển thị Báo cáo Chi tiết**: In ra danh sách tất cả các lớp, tự động cập nhật sĩ số, và liệt kê danh sách học sinh có trong từng lớp.
- **Giao diện Menu Tương tác**: Cung cấp một menu đơn giản, dễ sử dụng để người dùng lựa chọn các chức năng.
- **Xác thực Dữ liệu (Data Validation)**: Chương trình tự động kiểm tra và **ngăn chặn** việc tạo các lớp học có tên trùng nhau, đảm bảo dữ liệu luôn nhất quán.

## Cấu trúc Thư mục

Dự án được tổ chức theo một cấu trúc thư mục rõ ràng và được khuyến nghị trong cộng đồng Go:

```
school-manager/
├── cmd/
│   └── school-manager/
│       └── main.go         # Điểm khởi đầu (entrypoint) của ứng dụng
├── internal/
│   ├── cli/
│   │   └── handler.go      # Chứa logic nghiệp vụ và tương tác người dùng
│   └── models/
│       └── models.go       # Định nghĩa các struct dữ liệu (Classes, Students)
├── go.mod                  # File quản lý module của Go
└── README.md               # Tài liệu hướng dẫn
```

- **`cmd/school-manager/main.go`**: Đây là điểm khởi đầu duy nhất của chương trình. Nhiệm vụ chính của nó là khởi tạo các biến dữ liệu và gọi các hàm xử lý từ package `cli` trong một vòng lặp menu.
- **`internal/models/`**: Package này chứa định nghĩa cho các cấu trúc dữ liệu cốt lõi của ứng dụng (`Classes`, `Students`). Việc tách riêng models giúp dữ liệu trở nên độc lập và dễ dàng tái sử dụng.
- **`internal/cli/`**: Chứa toàn bộ logic nghiệp vụ, bao gồm việc hiển thị menu, nhận và xử lý input từ người dùng, và in kết quả ra màn hình.
- **`go.mod`**: File định nghĩa module Go, quản lý các dependency của dự án.

## Các Khái niệm Kỹ thuật Nổi bật

- **Tổ chức Package (Package Organization)**: Dự án thể hiện rõ cách chia nhỏ code thành các package có trách nhiệm riêng biệt (`models` cho dữ liệu, `cli` cho logic), giúp code dễ đọc và bảo trì.
- **Struct và Con trỏ**: Mối quan hệ giữa Học sinh và Lớp học được định nghĩa chặt chẽ bằng một con trỏ (`*Classes`) trong struct `Students`, đảm bảo tính toàn vẹn dữ liệu.
- **Toàn vẹn Dữ liệu (Data Integrity)**: Bằng cách kiểm tra trùng lặp tên lớp ngay tại khâu nhập liệu, chương trình đảm bảo mỗi lớp học là một thực thể duy nhất.
- **Slices và Maps**: Sử dụng slice như một cơ sở dữ liệu trong bộ nhớ và map để nhóm dữ liệu một cách hiệu quả khi hiển thị báo cáo.

## Cài đặt và Chạy chương trình

### Yêu cầu

- Cần cài đặt **Go** (phiên bản 1.18 trở lên). Bạn có thể tải về tại [https://go.dev/dl/](https://go.dev/dl/).

### Hướng dẫn

1.  Clone hoặc tải về toàn bộ thư mục dự án này.
2.  Mở terminal (Command Prompt, PowerShell, hoặc Terminal trên Linux/macOS).
3.  Sử dụng lệnh `cd` để di chuyển vào thư mục gốc của dự án (`school-manager`).
4.  Chạy ứng dụng bằng lệnh sau:

    ```sh
    go run ./cmd/school-manager
    ```

    **Lưu ý:** Lệnh này chỉ cho Go biết điểm khởi đầu của ứng dụng nằm ở đâu, sau đó Go sẽ tự động biên dịch tất cả các package liên quan trong `internal`.

5.  Chương trình sẽ khởi động và hiển thị menu để bạn bắt đầu tương tác.
