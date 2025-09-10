# Golang & PostgreSQL: CRUD và Transaction CLI Application

Đây là một dự án ứng dụng CLI (Command-Line Interface) đơn giản được xây dựng bằng Golang, dùng để minh họa các phương pháp tốt nhất khi tương tác với cơ sở dữ liệu quan hệ (PostgreSQL). Dự án tập trung vào việc sử dụng package `database/sql` tiêu chuẩn, cấu trúc code theo lớp (repository, service) và xử lý các giao dịch (transaction) một cách an toàn.

Ứng dụng quản lý hai thực thể chính: **Sản phẩm (Products)** và **Đơn hàng (Orders)**.

## Tính năng chính

- **Quản lý Sản phẩm (CRUD):**
  - Liệt kê danh sách tất cả sản phẩm.
  - Tìm kiếm một sản phẩm theo ID.
  - Thêm một sản phẩm mới vào kho.
  - Cập nhật số lượng tồn kho của sản phẩm.
  - Xóa một sản phẩm.
- **Quản lý Đơn hàng:**
  - Liệt kê danh sách tất cả đơn hàng đã tạo.
  - Tìm kiếm một đơn hàng theo ID.
- **Giao dịch (Transaction):**
  - Thực hiện chức năng **Đặt hàng**, một ví dụ kinh điển của transaction:
    1.  Kiểm tra số lượng tồn kho.
    2.  Trừ số lượng tồn kho của sản phẩm.
    3.  Tạo một bản ghi đơn hàng mới.
    - Toàn bộ quá trình sẽ được `ROLLBACK` nếu có bất kỳ lỗi nào xảy ra, đảm bảo tính toàn vẹn dữ liệu.

## Cấu trúc dự án

```text
sql/
├── cmd/
│   └── cli/
│       └── main.go           # Entrypoint: Khởi tạo dependencies và chạy CLI
├── internal/
│   ├── cli/
│   │   └── menu.go           # Logic giao diện người dùng (menu, I/O)
│   ├── config/
│   │   └── config.go         # Tải cấu hình (chuỗi kết nối DB)
│   ├── database/
│   │   └── postgres.go       # Logic kết nối đến PostgreSQL
│   ├── model/
│   │   ├── product.go        # Struct Product
│   │   └── order.go          # Struct Order
│   ├── repository/
│   │   ├── product_repo.go   # Lớp truy cập dữ liệu cho Product (SQL queries)
│   │   └── order_repo.go     # Lớp truy cập dữ liệu cho Order (SQL queries)
│   └── service/
│       ├── product_service.go # Logic nghiệp vụ cho Product
│       └── order_service.go   # Logic nghiệp vụ cho Order (bao gồm transaction)
├── go.mod
├── go.sum
├── docker-compose.yml        # File để chạy PostgreSQL container
└── init.sql                  # Script SQL để tạo bảng và dữ liệu mẫu
```

## Các khái niệm cốt lõi được minh họa

- **Kết nối CSDL:** Sử dụng package `database/sql` và driver `lib/pq` cho PostgreSQL.
- **CRUD Operations:** Thực thi các câu lệnh `SELECT`, `INSERT`, `UPDATE`, `DELETE`.
- **Transactions:** Đảm bảo tính toàn vẹn dữ liệu cho các hoạt động phức tạp thông qua `db.BeginTx()`, `tx.Commit()`, và `tx.Rollback()`.
- **Phòng chống SQL Injection:** Sử dụng câu lệnh tham số hóa (parameterized queries) với `$1`, `$2`,...
- **Clean Architecture (Simplified):** Tách biệt rõ ràng các lớp trách nhiệm:
  - `cli`: Tương tác với người dùng.
  - `service`: Xử lý logic nghiệp vụ.
  - `repository`: Chịu trách nhiệm duy nhất về việc truy vấn CSDL.

## Yêu cầu

- [Go](https://go.dev/dl/) (phiên bản 1.18 trở lên)
- [Docker](https://www.docker.com/get-started/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Hướng dẫn Cài đặt và Chạy

### 1. Clone Repository

```bash
git clone https://github.com/huynh-fs/golang-training.git
cd sql
```

### 2. Khởi động Cơ sở dữ liệu

Dự án sử dụng Docker Compose để khởi tạo một container PostgreSQL một cách nhanh chóng.

```bash
docker-compose up -d
```

Lệnh này sẽ:

- Tải image PostgreSQL.
- Tạo và chạy một container tên là `go_postgres_db`.
- Tự động thực thi file `init.sql` để tạo các bảng (`products`, `orders`) và chèn dữ liệu mẫu.

### 3. Chạy ứng dụng CLI

Mở một terminal khác tại thư mục gốc của dự án và chạy:

```bash
go run ./cmd/cli
```

## Cách sử dụng

Sau khi chạy lệnh trên, một menu tương tác sẽ xuất hiện trên terminal của bạn:

```text
--- MENU QUẢN LÝ SẢN PHẨM ---
1. Liệt kê tất cả sản phẩm
2. Tìm sản phẩm theo ID
3. Thêm sản phẩm mới
4. Cập nhật số lượng sản phẩm
5. Xóa sản phẩm
6. Đặt hàng (Transaction)
7. Lấy thông tin đơn hàng theo ID
8. Lấy danh sách tất cả đơn hàng
0. Thoát
Nhập lựa chọn của bạn:
```

Bạn chỉ cần nhập số tương ứng với chức năng mình muốn và làm theo hướng dẫn.

### Ví dụ về kịch bản đặt hàng:
```
1.  Nhập `1` để xem danh sách sản phẩm và số lượng tồn kho hiện tại.
2.  Nhập `6` để bắt đầu quá trình đặt hàng.
3.  Nhập ID sản phẩm và số lượng bạn muốn mua.
    -   Nếu số lượng hợp lệ, đơn hàng sẽ được tạo thành công.
    -   Nếu số lượng vượt quá tồn kho, bạn sẽ nhận được thông báo lỗi và không có thay đổi nào được ghi vào CSDL.
4.  Nhập `1` một lần nữa để thấy số lượng tồn kho của sản phẩm đã được cập nhật chính xác.
5.  Nhập `8` để xem đơn hàng vừa được tạo.
```
