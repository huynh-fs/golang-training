# Go Method Demo: Ứng dụng CLI Quản lý Ngân hàng

Đây là một dự án Go đơn giản được xây dựng để minh họa các khái niệm cốt lõi về **method** (phương thức) và cách tổ chức **cấu trúc dự án** một cách rõ ràng. Ứng dụng mô phỏng một giao diện dòng lệnh (CLI) để tương tác với một tài khoản ngân hàng, tuân theo nguyên tắc phân tách rõ ràng giữa **Model** (logic nghiệp vụ), **CLI Handler** (logic xử lý giao diện) và **Main** (điểm khởi chạy).

## Tính năng chính

- Xem thông tin chi tiết của một tài khoản ngân hàng.
- Nạp tiền vào tài khoản.
- Rút tiền khỏi tài khoản (có kiểm tra số dư).
- Giao diện menu tương tác, thân thiện với người dùng.
- Thao tác trên một tài khoản mặc định (`acc001`) để đơn giản hóa việc demo.

## Kiến trúc & Triết lý thiết kế

Mục tiêu chính của dự án này là thể hiện sức mạnh của việc **Tách biệt các mối quan tâm (Separation of Concerns)**.

1.  **`internal/model` (Lõi nghiệp vụ)**

    - Chứa định nghĩa `struct BankAccount` và toàn bộ logic nghiệp vụ.
    - Các method như `Deposit()` và `Withdraw()` chứa các quy tắc (ví dụ: không được rút quá số dư).
    - **Quan trọng**: Package này hoàn toàn không biết gì về giao diện người dùng. Nó không in ra menu, cũng không đọc input. Nó chỉ làm việc với dữ liệu.

2.  **`internal/cli` (Lớp xử lý giao diện)**

    - Đóng vai trò là cầu nối giữa `main` và `model`.
    - Các hàm như `HandleGetAccount()`, `HandleDeposit()` nhận các dữ liệu đã được xử lý (ví dụ: ID tài khoản, số tiền), sau đó gọi các method tương ứng trong `model`.
    - Chịu trách nhiệm định dạng kết quả và in ra console cho người dùng.

3.  **`cmd/bank/main.go` (Điểm khởi chạy)**
    - Chỉ có một nhiệm vụ duy nhất: hiển thị menu, đọc và xác thực input của người dùng.
    - Sau khi có được input hợp lệ, nó sẽ gọi đến các hàm trong package `cli` để thực thi chức năng.
    - File này không chứa bất kỳ logic nghiệp vụ nào.

=> **Lợi ích**: Với kiến trúc này, nếu trong tương lai bạn muốn thay đổi giao diện từ CLI sang API Web, bạn chỉ cần viết một `handler` mới cho web và một `main.go` mới mà **không cần sửa một dòng code nào** trong `internal/model`.

## Cấu trúc thư mục

```
method/
├── go.mod
├── internal/
│   ├── model/
│   │   └── account.go
│   └── cli/
│       └── account_cli.go
└── cmd/
    └── bank/
        └── main.go
```

## Yêu cầu

- Go (phiên bản 1.18 trở lên được khuyến nghị).

## Ví dụ sử dụng

Sau khi chạy ứng dụng, bạn sẽ thấy một menu tương tác. Dưới đây là một luồng sử dụng mẫu:

```
--- MENU NGÂN HÀNG ---
1. Xem thông tin tài khoản (tài khoản mặc định ID: acc001)
2. Nạp tiền
3. Rút tiền
4. Thoát
Vui lòng chọn một chức năng: 1
--- Thông Tin Tài Khoản ---
ID: acc001
Chủ tài khoản: Bob
Số dư: 5000.00
---------------------------

--- MENU NGÂN HÀNG ---
1. Xem thông tin tài khoản (tài khoản mặc định ID: acc001)
2. Nạp tiền
3. Rút tiền
4. Thoát
Vui lòng chọn một chức năng: 2
Nhập số tiền cần nạp: 1500
Nạp tiền thành công 1500.00.
Số dư mới: 6500.00

--- MENU NGÂN HÀNG ---
1. Xem thông tin tài khoản (tài khoản mặc định ID: acc001)
2. Nạp tiền
3. Rút tiền
4. Thoát
Vui lòng chọn một chức năng: 3
Nhập số tiền cần rút: 9000
lỗi: số dư không đủ

--- MENU NGÂN HÀNG ---
1. Xem thông tin tài khoản (tài khoản mặc định ID: acc001)
2. Nạp tiền
3. Rút tiền
4. Thoát
Vui lòng chọn một chức năng: 4
Cảm ơn bạn đã sử dụng dịch vụ. Tạm biệt!
```

## Các khái niệm Go được áp dụng

- **Methods**: Gắn hàm vào một kiểu dữ liệu (`struct`).
- **Pointer Receivers**: Sử dụng `*BankAccount` để cho phép các method thay đổi trạng thái của đối tượng.
- **Structs**: Định nghĩa kiểu dữ liệu `BankAccount`.
- **Packages**: Tổ chức code thành các module có thể tái sử dụng (`model`, `cli`).
- **Project Layout**: Quy ước cấu trúc thư mục với `cmd` và `internal`.
- **Input/Output (I/O)**: Sử dụng `fmt` để in và `bufio.Scanner` để đọc input từ người dùng.
- **Error Handling**: Kiểm tra và xử lý các lỗi cơ bản (ví dụ: chuyển đổi chuỗi sang số).
