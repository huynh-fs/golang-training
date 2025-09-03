# Hệ thống Quản lý Lớp học và Học sinh (CLI)

Đây là một ứng dụng dòng lệnh (CLI - Command-Line Interface) được xây dựng bằng ngôn ngữ Go. Mục tiêu của dự án là quản lý danh sách các lớp học và học sinh, minh họa một cách hiệu quả việc sử dụng `struct`, con trỏ, slice và map để tạo ra một cấu trúc dữ liệu có mối quan hệ rõ ràng và đảm bảo tính toàn vẹn.

Ứng dụng này là một ví dụ thực tế tuyệt vời cho những người đang học Go về cách tổ chức dữ liệu, xử lý input từ người dùng và xây dựng một chương trình tương tác có tính ổn định.

## Tính năng

- **Thêm Lớp học Mới**: Cho phép người dùng thêm nhiều lớp học vào hệ thống.
- **Thêm Học sinh Mới**: Cho phép thêm học sinh và liên kết học sinh đó với một lớp học đã tồn tại.
- **Hiển thị Báo cáo Chi tiết**: In ra danh sách tất cả các lớp, tự động cập nhật sĩ số, và liệt kê danh sách học sinh có trong từng lớp.
- **Giao diện Menu Tương tác**: Cung cấp một menu đơn giản, dễ sử dụng để người dùng lựa chọn các chức năng.
- **Xác thực Dữ liệu (Data Validation)**: Chương trình sẽ tự động kiểm tra và **ngăn chặn** người dùng tạo các lớp học có tên trùng nhau, đảm bảo dữ liệu luôn nhất quán và chính xác.

## Các Khái niệm Kỹ thuật Nổi bật

Dự án này thể hiện rõ các khái niệm quan trọng trong lập trình Go:

- **Struct và Con trỏ (Structs and Pointers)**: Mối quan hệ giữa Học sinh và Lớp học được định nghĩa một cách chặt chẽ. Struct `Students` lưu một **con trỏ** (`*Classes`) trỏ trực tiếp đến đối tượng lớp học, đảm bảo tính toàn vẹn dữ liệu và tạo ra một liên kết hiệu quả.
- **Slices làm Cơ sở dữ liệu trong Bộ nhớ**: `[]Classes` và `[]Students` được sử dụng như một cơ sở dữ liệu tạm thời (in-memory database) để lưu trữ toàn bộ dữ liệu trong suốt phiên làm việc của ứng dụng.
- **Sử dụng Map để Nhóm dữ liệu**: Trong chức năng hiển thị, một `map[*Classes][]Students` được sử dụng một cách thông minh để nhóm tất cả học sinh theo lớp của họ. Đây là một phương pháp rất hiệu quả để xử lý và trình bày dữ liệu.
- **Toàn vẹn Dữ liệu (Data Integrity)**: Bằng cách kiểm tra trùng lặp tên lớp ngay tại khâu nhập liệu, chương trình đảm bảo mỗi lớp học là một thực thể duy nhất. Điều này tránh được các lỗi logic và sai sót dữ liệu có thể xảy ra khi vận hành.
- **Kiến trúc Ứng dụng CLI**: Hàm `main` được tổ chức với vòng lặp `for` và cấu trúc `switch-case`, một mô hình phổ biến và mạnh mẽ để xây dựng các ứng dụng có menu tương tác.

## Yêu cầu

- Cần cài đặt **Go** (phiên bản 1.18 trở lên) trên máy của bạn. Bạn có thể tải về và cài đặt từ trang chủ: [https://go.dev/dl/](https://go.dev/dl/)

## Hướng dẫn Sử dụng

1.  Lưu toàn bộ mã nguồn trên vào một file có tên `main.go`.
2.  Mở terminal (Command Prompt, PowerShell, hoặc Terminal trên Linux/macOS).
3.  Sử dụng lệnh `cd` để di chuyển đến thư mục mà bạn vừa lưu file `main.go`.
4.  Biên dịch và chạy ứng dụng bằng lệnh sau:
    ```sh
    go run main.go
    ```
5.  Chương trình sẽ khởi động và hiển thị menu chính:
    ```
    ===== MENU HỆ THỐNG QUẢN LÝ =====
    1. Thêm Lớp Học Mới
    2. Thêm Học Sinh Mới
    3. Xuất Báo Cáo Thông Tin
    4. Thoát Chương Trình
    Nhập lựa chọn của bạn:
    ```
6.  Làm theo các hướng dẫn trên màn hình:
    - Bắt đầu bằng cách chọn `1` để tạo một vài lớp học (bạn sẽ không thể tạo 2 lớp trùng tên).
    - Sau đó, chọn `2` để thêm học sinh vào các lớp đã tạo.
    - Chọn `3` bất cứ lúc nào để xem báo cáo tổng hợp.
    - Chọn `4` để kết thúc chương trình.

## Cấu trúc Code

- `main()`: Là điểm khởi đầu và vòng lặp chính của ứng dụng, điều hướng các lựa chọn của người dùng.
- `printMenu()`: Chịu trách nhiệm hiển thị các tùy chọn cho người dùng.
- `inputChoice()`: Đọc và trả về lựa chọn của người dùng.
- `inputClassesInfo()`: Xử lý logic cho việc thêm các lớp học mới, bao gồm cả việc **kiểm tra trùng lặp**.
- `inputStudentInfo()`: Xử lý logic cho việc thêm học sinh mới, bao gồm cả việc xác thực tên lớp.
- `displayInfo()`: Xử lý logic tính toán sĩ số và hiển thị thông tin tổng hợp của tất cả các lớp và học sinh.
