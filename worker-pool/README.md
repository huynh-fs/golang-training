# Go Worker Pool

Dự án này là một ví dụ thực tế về việc triển khai pattern **Worker Pool** trong Go, được cấu trúc theo một kiến trúc phân lớp sạch sẽ, lấy cảm hứng từ "Standard Go Project Layout". Mục tiêu là minh họa cách xây dựng một ứng dụng Go không chỉ hiệu quả về mặt đồng thời mà còn dễ bảo trì, dễ kiểm thử và dễ mở rộng.

## ✨ Tính năng nổi bật

- **Pattern Worker Pool**: Giới hạn số lượng goroutine chạy đồng thời để xử lý các tác vụ một cách hiệu quả và kiểm soát tài nguyên.
- **Kiến trúc phân lớp (Layered Architecture)**: Tách biệt rõ ràng các mối quan tâm (Separation of Concerns) giữa logic nghiệp vụ (service), truy cập dữ liệu (repository) và điểm khởi động (cmd).
- **Dependency Injection**: Các thành phần phụ thuộc (dependencies) được "tiêm" vào từ bên ngoài, giúp mã nguồn linh hoạt và cực kỳ dễ dàng cho việc viết unit test.
- **Cấu hình động**: Các thông số quan trọng như số lượng worker được quản lý trong file cấu hình YAML, không cần hard-code.
- **Logging có cấu trúc**: Sử dụng package `log` tiêu chuẩn để cung cấp thông tin hữu ích trong quá trình chạy ứng dụng.

## 📂 Cấu trúc Dự án

Dự án tuân theo một cấu trúc thư mục rõ ràng để phân tách các thành phần logic:

```
/worker-pool
├── /cmd/task-processor/main.go         # Điểm vào chính của ứng dụng, nơi khởi tạo và kết nối các thành phần.
├── /internal/
│   ├── repository/             # Lớp truy cập dữ liệu (Data Access Layer).
│   │   └── task_repository.go
│   └── service/                # Lớp logic nghiệp vụ (Business Logic Layer).
│       └── task_service.go
├── /pkg/
│   ├── config/                 # Thư viện đọc và quản lý cấu hình.
│   │   └── config.go
│   └── logger/                 # Thư viện helper cho việc logging.
│       └── logger.go
├── /configs/
│   └── config.yaml             # File cấu hình của ứng dụng.
├── go.mod                      # File quản lý module và các dependency.
└── README.md                   # File mô tả dự án
```

- **/cmd**: Chứa các file `main` của ứng dụng. Mỗi thư mục con là một ứng dụng có thể thực thi.
- **/internal**: Chứa logic cốt lõi của ứng dụng. Các package trong này không thể được import bởi các project bên ngoài.
  - `repository`: Chịu trách nhiệm truy cập dữ liệu (ví dụ: từ database, cache, API...).
  - `service`: Chịu trách nhiệm thực thi các quy trình nghiệp vụ chính của ứng dụng.
- **/pkg**: Chứa các thư viện, helper có thể được tái sử dụng và an toàn để các project khác import.
- **/configs**: Chứa các file cấu hình tĩnh.

## 🚀 Bắt đầu

### Yêu cầu

- [Go](https://golang.org/dl/) (khuyến nghị phiên bản 1.18 trở lên)

### Cài đặt

1.  **Clone repository:**
    ```bash
    git clone https://github.com/huynh-fs/golang-traning.git
    cd worker-pool
    ```

2.  **Tải các dependency:**
    ```bash
    go mod tidy
    ```

### Cấu hình

Bạn có thể thay đổi hoạt động của ứng dụng bằng cách chỉnh sửa file `configs/config.yaml`.

```yaml
worker_pool:
  workers: 4  # Số lượng worker goroutine chạy đồng thời.
  tasks: 20   # Tổng số tác vụ cần xử lý.
```

### Chạy ứng dụng

Để khởi chạy ứng dụng, chạy lệnh sau từ thư mục gốc của dự án:

```bash
go run ./cmd/task-processor
```

Bạn sẽ thấy output tương tự như sau:

```
INFO: 2025/09/08 10:42:26 main.go:16: Bắt đầu ứng dụng...
INFO: 2025/09/08 10:42:26 main.go:23: Cấu hình đã được tải: 4 workers, 20 tasks.
INFO: 2025/09/08 10:42:26 task_service.go:41: Khởi tạo 4 workers...
INFO: 2025/09/08 10:42:26 task_service.go:60: Worker 4 đã bắt đầu
Bắt đầu xử lý tác vụ 1: Data for task 1
INFO: 2025/09/08 10:42:26 task_service.go:60: Worker 1 đã bắt đầu
Bắt đầu xử lý tác vụ 2: Data for task 2
INFO: 2025/09/08 10:42:26 task_service.go:60: Worker 2 đã bắt đầu
Bắt đầu xử lý tác vụ 3: Data for task 3
INFO: 2025/09/08 10:42:26 task_service.go:60: Worker 3 đã bắt đầu
Bắt đầu xử lý tác vụ 4: Data for task 4
>> Hoàn thành xử lý tác vụ 1
Bắt đầu xử lý tác vụ 5: Data for task 5
>> Hoàn thành xử lý tác vụ 4
...
>> Hoàn thành xử lý tác vụ 19
INFO: 2025/09/08 10:42:31 task_service.go:64: Worker 2 đã kết thúc
>> Hoàn thành xử lý tác vụ 20
INFO: 2025/09/08 10:42:31 task_service.go:64: Worker 1 đã kết thúc
INFO: 2025/09/08 10:42:31 main.go:38: Tất cả các tác vụ đã được xử lý.
Chương trình chạy trong: 5.009015s
```

## 💡 Các khái niệm cốt lõi

### Worker Pool

Đây là một pattern quản lý đồng thời, nơi một số lượng worker (goroutine) cố định được tạo ra để xử lý các tác vụ từ một hàng đợi chung (channel). Lợi ích chính là ngăn chặn việc tạo ra vô số goroutine có thể làm cạn kiệt tài nguyên hệ thống khi có lượng lớn công việc cần xử lý.

### Kiến trúc phân lớp & Dependency Injection

- **Luồng phụ thuộc**: `main.go` -> `service` -> `repository`.
- **Nguyên tắc**: Lớp `service` không biết chi tiết về cách `repository` lấy dữ liệu (là từ bộ nhớ, database, hay API). Nó chỉ làm việc với `interface` của repository.
- **Lợi ích**:
  - **Dễ kiểm thử**: Khi viết test cho `TaskService`, chúng ta có thể "mock" (giả lập) `TaskRepository` để cung cấp dữ liệu giả mà không cần kết nối database thật.
  - **Dễ thay thế**: Nếu bạn muốn đổi từ việc lấy task trong bộ nhớ sang lấy từ PostgreSQL, bạn chỉ cần tạo một `PostgresTaskRepository` mới và thay đổi một dòng trong `main.go` để "tiêm" dependency mới này vào `TaskService`. Toàn bộ business logic không cần thay đổi.