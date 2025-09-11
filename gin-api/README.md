# Go Todo API 

Đây là một dự án API RESTful để quản lý danh sách công việc (Todo List), được xây dựng bằng Go. Dự án này là một ví dụ điển hình về việc áp dụng **kiến trúc phân tầng (Layered Architecture)** và các "best practice" trong việc cấu trúc một ứng dụng Go production-ready.

- **Ngôn ngữ:** Go
- **Framework:** Gin
- **ORM:** GORM
- **Cơ sở dữ liệu:** PostgreSQL
- **Triển khai:** Docker & Docker Compose
- **Tài liệu API:** Swagger

## Kiến trúc dự án

Dự án tuân thủ một cấu trúc rõ ràng, tách biệt các mối quan tâm (Separation of Concerns) để tăng cường khả năng bảo trì, kiểm thử và mở rộng.

- **Handler Layer:** Tầng ngoài cùng, chịu trách nhiệm xử lý chu kỳ request/response HTTP. Nó chỉ parse request, validate dữ liệu đầu vào (thông qua DTO) và gọi đến Service layer.
- **Service Layer:** Chứa toàn bộ business logic (logic nghiệp vụ) của ứng dụng. Tầng này không biết gì về HTTP, giúp logic cốt lõi có thể được tái sử dụng ở nhiều nơi khác (ví dụ: gRPC, CLI).
- **Dependency Injection:** Các phụ thuộc (dependencies) được "tiêm" từ `main.go` vào các tầng thấp hơn (ví dụ: DB được inject vào Service, Service được inject vào Handler).

## Cấu trúc thư mục

```
/
├── cmd/api/
│   └── main.go           # Entry point: khởi tạo dependencies và "chắp nối" các tầng
├── internal/
│   ├── dto/              # Data Transfer Objects: Định nghĩa cấu trúc dữ liệu cho API request/response
│   ├── handler/          # HTTP Handlers: Xử lý request, gọi service, trả về response
│   ├── model/            # GORM Models: Định nghĩa cấu trúc cho bảng trong database
│   ├── router/           # Định nghĩa các API endpoint và gán chúng với handler tương ứng
│   └── service/          # Business Logic: Chứa logic nghiệp vụ cốt lõi của ứng dụng
├── pkg/
│   ├── config/           # Các package có thể tái sử dụng: Quản lý config từ .env
│   └── database/         # Các package có thể tái sử dụng: Xử lý kết nối database
├── scripts/
│   └── wait-for-it.sh    # Chứa các script hỗ trợ
├── .env.example          # File mẫu cho biến môi trường
├── docker-compose.yml    # Định nghĩa các service (app, db) cho Docker
├── Dockerfile            # Cấu hình để build Docker image cho ứng dụng
├── go.mod & go.sum       # Quản lý dependencies
└── README.md
```

## Yêu cầu cài đặt

- [Docker](https://www.docker.com/products/docker-desktop/) & Docker Compose
- [Go](https://golang.org/dl/) (Phiên bản 1.22+ được khuyến nghị)
- [Swag](https://github.com/swaggo/swag) (Để sinh tài liệu Swagger)

## Hướng dẫn cài đặt và khởi chạy

### 1. Clone Repository

```bash
git clone https://github.com/huynh-fs/golang-training.git
cd gin-api
```

### 2. Cấu hình môi trường

Tạo một file `.env` từ file mẫu. Các giá trị mặc định đã được cấu hình để hoạt động với Docker Compose.

```bash
cp .env.example .env
```

### 3. Cài đặt các công cụ Go (nếu chưa có)

```bash
# Cài đặt Swag để sinh docs
go install github.com/swaggo/swag/cmd/swag@latest
```

### 4. Khởi chạy dự án với Docker Compose (Khuyến nghị)

Cách này sẽ tự động dựng và quản lý cả ứng dụng và cơ sở dữ liệu.

```bash
docker-compose up --build -d
```

- `--build`: Build lại image nếu có thay đổi trong `Dockerfile`.
- `-d`: Chạy ở chế độ detached (chạy ngầm), giải phóng terminal của bạn.

Ứng dụng sẽ chạy tại `http://localhost:8080`.

## Quy trình phát triển (Chạy trực tiếp trên máy)

Để có chu kỳ phát triển nhanh hơn mà không cần build lại Docker image mỗi lần, bạn có thể chạy ứng dụng Go trực tiếp trên máy của mình.

```bash
# Bước 1: Chỉ khởi động database bằng Docker
docker-compose up -d db

# Bước 2: Chạy ứng dụng Go từ terminal
go run cmd/api/main.go
```

**Lưu ý:** Mỗi khi bạn thay đổi mã nguồn, bạn cần dừng ứng dụng (`Ctrl + C`) và chạy lại lệnh `go run cmd/api/main.go` để thấy được các thay đổi.

## Tài liệu API (Swagger)

Sau khi khởi chạy ứng dụng, bạn có thể truy cập vào giao diện Swagger UI để xem tài liệu chi tiết và thử nghiệm các API:

**URL:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Để cập nhật tài liệu sau khi thay đổi các comment trong code, bạn **phải** chạy lại lệnh sau:

```bash
swag init -g main.go -d ./cmd/api,./internal/handler,./internal/model,./internal/dto
```

Sau khi chạy lệnh trên, hãy khởi động lại ứng dụng của bạn để áp dụng tài liệu mới.

## Các Endpoint của API

Tất cả các endpoint đều có tiền tố là `/api/v1`.

| Phương thức | Endpoint      | Mô tả                                |
| :---------- | :------------ | :----------------------------------- |
| `GET`       | `/todos`      | Lấy danh sách tất cả công việc.      |
| `POST`      | `/todos`      | Tạo một công việc mới.               |
| `GET`       | `/todos/{id}` | Lấy thông tin một công việc theo ID. |
| `PUT`       | `/todos/{id}` | Cập nhật một công việc theo ID.      |
| `DELETE`    | `/todos/{id}` | Xóa một công việc theo ID.           |

## Dừng ứng dụng

Để dừng tất cả các container đang chạy (app và db), sử dụng lệnh:

```bash
docker-compose down
```

Lệnh này sẽ dừng và xóa các container, nhưng giữ lại data của database trong volume.
