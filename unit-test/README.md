# Go Unit Testing Demo với CLI Tương Tác

Dự án này là một minh họa toàn diện về cách triển khai unit testing trong Go. Nó tuân theo Cấu trúc dự án Go tiêu chuẩn (`Standard Go Project Layout`) và bao gồm một Giao diện Dòng lệnh (CLI) tương tác để chạy và khám phá các bài test.

Mục tiêu chính của dự án là cung cấp một ví dụ thực tế, dễ hiểu cho các lập trình viên muốn nắm vững các kỹ thuật kiểm thử trong Go.

## ✨ Tính Năng Nổi Bật

- **Cấu Trúc Dự Án Tiêu Chuẩn:** Sử dụng layout `cmd` và `internal` để tách biệt rõ ràng các mối quan tâm.
- **Unit Tests Cơ Bản:** Sử dụng package `testing` tích hợp sẵn của Go để viết các bài test.
- **Table-Driven Tests:** Minh họa cách viết các bài test theo mẫu thiết kế "Table-Driven", một phương pháp hay nhất (best practice) trong cộng đồng Go.
- **Tách Biệt Trách Nhiệm:** Phân chia logic rõ ràng giữa `model` (cấu trúc dữ liệu), `service` (logic nghiệp vụ), và `handler` (logic giao diện).
- **CLI Tương Tác:** Cung cấp một công cụ dòng lệnh thân thiện để thực thi các lệnh `go test` với các cờ khác nhau (`-v`, `-cover`, `-run`) mà không cần gõ lệnh thủ công.

## 📂 Cấu Trúc Dự Án

Dự án được tổ chức theo cấu trúc được khuyến nghị để dễ dàng bảo trì và mở rộng.

```
/
├── cmd/
│   └── cli/
│       └── main.go              # Điểm khởi đầu của ứng dụng, chỉ gọi CLI handler.
├── internal/
│   ├── handler/
│   │   └── cli_handler.go           # Chứa tất cả logic cho CLI tương tác.
│   ├── model/
│   │   ├── task.go              # Định nghĩa struct dữ liệu Task.
│   │   └── task_test.go         # Unit test cho model.
│   └── service/
│       ├── task_service.go      # Chứa logic nghiệp vụ chính (thêm, sửa, xóa task).
│       └── task_service_test.go # Unit test cho service, sử dụng table-driven tests.
├── go.mod                       # File quản lý module và dependency.
└── README.md                    # Tài liệu hướng dẫn này.
```

## 🚀 Bắt Đầu

Để chạy dự án này trên máy của bạn, hãy làm theo các bước sau.

### Yêu Cầu

- [Go](https://go.dev/doc/install) (khuyến nghị phiên bản 1.18 trở lên).

### Hướng Dẫn

1.  **Clone repository về máy:**

    ```bash
    git clone https://github.com/huynh-fs/golang-training.git
    ```

2.  **Di chuyển vào thư mục dự án:**

    ```bash
    cd unit-test
    ```

3.  **Chạy ứng dụng CLI:**

    ```bash
    go run ./cmd/cli
    ```

4.  **Tương tác với menu:**
    Sau khi chạy lệnh trên, bạn sẽ thấy một menu tương tác. Hãy nhập lựa chọn của bạn và nhấn Enter.
    ```
    ======================================
       Trình Chạy Unit Test Tương Tác
    ======================================
    1. Chạy tất cả các test (go test ./...)
    2. Chạy test với output chi tiết (go test -v ./...)
    3. Chạy test và xem độ bao phủ code (go test -cover ./...)
    4. Chạy một test cụ thể (go test -run 'TestName' ./...)
    5. Thoát
    Nhập lựa chọn của bạn:
    ```

## 💡 Các Khái Niệm Cốt Lõi Được Minh Họa

### 1. Unit Testing với package `testing`

Dự án sử dụng gói `testing` chuẩn của Go. Các bài test được đặt trong các file `_test.go` và tuân theo quy ước đặt tên `func TestXxx(t *testing.T)`.

### 2. Table-Driven Tests

Trong `internal/service/task_service_test.go`, chúng tôi sử dụng mẫu thiết kế table-driven để kiểm tra hàm `CreateTask` với nhiều trường hợp đầu vào khác nhau một cách ngắn gọn và dễ bảo trì.

_Ví dụ một phần:_

```go
// internal/service/task_service_test.go
func TestCreateTask(t *testing.T) {
	s := NewTaskService()
	testCases := []struct {
		name        string
		title       string
		expectError bool
	}{
		{"Tạo thành công", "Task 1", false},
		{"Tiêu đề trống", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ... logic kiểm thử ...
		})
	}
}
```

### 3. Tách Biệt Logic Giao Diện

Thay vì viết tất cả code vào hàm `main`, chúng tôi đã tạo một package `handler`. Package này chịu trách nhiệm hoàn toàn cho việc hiển thị menu và xử lý input của người dùng. Hàm `main` chỉ có nhiệm vụ khởi tạo và chạy handler này, giúp code trở nên sạch sẽ và tuân thủ Nguyên tắc Đơn trách nhiệm.

## 🛠️ Chạy Test Thủ Công

Ngoài việc sử dụng CLI tương tác, bạn cũng có thể chạy các bài test một cách thủ công bằng các lệnh `go test` tiêu chuẩn từ thư mục gốc của dự án.

- **Chạy tất cả các test:**

  ```bash
  go test ./...
  ```

- **Chạy test với output chi tiết (verbose):**

  ```bash
  go test -v ./...
  ```

- **Chạy test và kiểm tra độ bao phủ code (code coverage):**
  ```bash
  go test -cover ./...
  ```
