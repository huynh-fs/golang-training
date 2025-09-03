# Hệ thống Quản lý Sinh viên bằng Go (CLI)

Đây là một dự án ứng dụng dòng lệnh (CLI) đơn giản được xây dựng bằng ngôn ngữ Go. Mục tiêu chính của dự án là để minh họa cách sử dụng `struct` và `struct` lồng nhau (nested structs) trong Go để quản lý một danh sách đối tượng.

Dự án này là một công cụ học tập tuyệt vời cho những ai mới bắt đầu với Go và muốn hiểu rõ hơn về cách cấu trúc dữ liệu trong một ứng dụng thực tế.

## Tính năng

-   **Thêm sinh viên thủ công**: Cho phép người dùng nhập thông tin chi tiết cho từng sinh viên qua các câu lệnh nhắc.
-   **Hiển thị danh sách**: In ra màn hình thông tin của tất cả sinh viên hiện có trong hệ thống.
-   **Tìm kiếm sinh viên**: Tìm một sinh viên cụ thể dựa trên ID duy nhất của họ.
-   **Nhập hàng loạt từ file CSV**: Thêm nhanh nhiều sinh viên vào hệ thống bằng cách đọc dữ liệu từ một file `sinhvien.csv`.

## Các khái niệm Go được sử dụng

Dự án này tập trung vào việc minh họa các khái niệm cốt lõi của Go:

-   **Struct**: Sử dụng `struct SinhVien` để định nghĩa một mẫu (template) cho đối tượng sinh viên, nhóm các dữ liệu liên quan như ID, Tên, Tuổi lại với nhau.
-   **Struct lồng nhau (Nested Structs)**: Sử dụng `struct ThongTinLienLac` bên trong `struct SinhVien` để tổ chức dữ liệu một cách logic và rõ ràng hơn.
-   **Slice của Struct (`[]SinhVien`)**: Dùng một slice để hoạt động như một cơ sở dữ liệu trong bộ nhớ (in-memory database), lưu trữ danh sách các đối tượng sinh viên.
-   **Xử lý File và CSV**: Sử dụng các thư viện chuẩn của Go là `os` và `encoding/csv` để đọc và phân tích cú pháp file dữ liệu.
-   **Input/Output cơ bản**: Dùng các gói `fmt` và `bufio` để tương tác với người dùng qua terminal.

## Yêu cầu

-   Cần cài đặt **Go** (phiên bản 1.18 trở lên) trên máy của bạn. Bạn có thể tải về tại [https://go.dev/dl/](https://go.dev/dl/).

## Cài đặt và Thiết lập

1.  **Clone repository (hoặc tải mã nguồn)**
    Nếu dự án nằm trên Git, bạn có thể clone nó. Nếu không, chỉ cần tạo một thư mục và đặt file `main.go` vào đó.

2.  **Tạo file dữ liệu `sinhvien.csv`**
    Trong cùng thư mục với `main.go`, tạo một file tên là `sinhvien.csv`. Đây là file mà chương trình sẽ đọc để nhập dữ liệu hàng loạt.

    **Quan trọng:** File phải có cấu trúc 5 cột, với dòng đầu tiên là tiêu đề.

    *Nội dung mẫu cho `sinhvien.csv`:*
    ```csv
    Ho,Ten,Tuoi,Email,SoDienThoai
    Nguyen,Van A,21,nva@example.com,0901112222
    Tran,Thi B,22,ttb@example.com,0902223333
    Le,Huu C,20,lhc@example.com,0903334444
    Pham,Gia D,23,pgd@example.com,0904445555
    ```

## Cách sử dụng

1.  Mở terminal hoặc command prompt.
2.  Di chuyển đến thư mục chứa dự án.
3.  Chạy ứng dụng bằng lệnh sau:
    ```sh
    go run main.go
    ```
4.  Chương trình sẽ hiển thị một menu với các lựa chọn:
    ```
    --- HỆ THỐNG QUẢN LÝ SINH VIÊN ---
    1. Thêm sinh viên mới (thủ công)
    2. Hiển thị danh sách sinh viên
    3. Tìm sinh viên theo ID
    4. Thêm sinh viên từ file CSV
    5. Thoát
    Nhập lựa chọn của bạn:
    ```
5.  Nhập số tương ứng với chức năng bạn muốn sử dụng và nhấn Enter.
    -   Chọn `4` để nhập dữ liệu từ file `sinhvien.csv`.
    -   Chọn `2` để xem kết quả.
    -   Chọn `1` để thêm một sinh viên khác bằng tay.