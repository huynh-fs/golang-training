# User Service API

![Go Version](https://img.shields.io/badge/go-1.18%2B-blue.svg) ![License](https://img.shields.io/badge/license-MIT-green.svg)

Một dự án API Service đơn giản được xây dựng bằng Go, nhằm mục đích minh họa cách áp dụng cấu trúc thư mục tiêu chuẩn từ [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

Service này cung cấp một endpoint cơ bản để truy xuất thông tin người dùng qua ID.

## Giới thiệu

Dự án này không chỉ là một API hoạt động được mà còn là một ví dụ thực tế về cách tổ chức mã nguồn trong một dự án Go có quy mô vừa và lớn. Việc tuân theo một cấu trúc chuẩn giúp dự án dễ dàng bảo trì, mở rộng và giúp các lập trình viên mới nhanh chóng nắm bắt được luồng hoạt động.

## Cấu trúc thư mục

Dự án này tuân theo cấu trúc `project-layout` được cộng đồng Go công nhận rộng rãi. Dưới đây là vai trò của các thư mục chính được sử dụng:

-   **/cmd**: Chứa các ứng dụng chính (điểm khởi đầu - `main.go`). Mỗi thư mục con là một ứng dụng có thể thực thi.
-   **/internal**: Chứa toàn bộ mã nguồn riêng tư của ứng dụng. Mã nguồn trong thư mục này không thể được import bởi các dự án bên ngoài.
    -   /internal/model: Định nghĩa các đối tượng dữ liệu cốt lõi (structs).
    -   /internal/repository: Lớp chịu trách nhiệm truy cập dữ liệu (ví dụ: tương tác với database).
    -   /internal/service: Lớp chứa logic nghiệp vụ (business logic).
    -   /internal/handler: Lớp xử lý các request HTTP, gọi đến service và trả về response.
-   **/pkg**: Chứa các gói mã nguồn có thể tái sử dụng và an toàn để các dự án khác import.
-   **/api**: Chứa các tệp định nghĩa, hợp đồng API (ví dụ: OpenAPI/Swagger).
-   **/configs**: Chứa các tệp cấu hình cho ứng dụng.
-   **/scripts**: Chứa các script tiện ích để tự động hóa các tác vụ (ví dụ: build, deploy).
-   **/bin**: Chứa các tệp thực thi sau khi được build (thường được tạo ra bởi script).

## Yêu cầu

-   [Go](https://golang.org/dl/) phiên bản 1.18 trở lên.
-   [curl](https://curl.se/) để kiểm tra các endpoint.

## Hướng dẫn cài đặt và sử dụng

### 1. Clone Repository

```bash
git clone https://github.com/huynh-fs/user-service.git
cd user-service
```

### 2. Cài đặt Dependencies
Dự án sử dụng Go Modules để quản lý các gói phụ thuộc. Lệnh sau sẽ tự động tải về các thư viện cần thiết được định nghĩa trong go.mod.
code
```bash
go mod tidy
```

### 3. Chạy ứng dụng (Development)
Để chạy ứng dụng trực tiếp mà không cần build, sử dụng lệnh go run:
code
```bash
go run ./cmd/user-api
```
Server sẽ khởi động và lắng nghe tại http://localhost:8080 (cấu hình trong configs/config.yaml).
### 4. Build ứng dụng
Để biên dịch ứng dụng ra một tệp thực thi duy nhất, bạn có thể sử dụng script đã được cung cấp:
code
```bash
# Cấp quyền thực thi cho script (chỉ cần làm một lần)
chmod +x scripts/build.sh

# Chạy script build
./scripts/build.sh
```
Tệp thực thi sẽ được tạo tại bin/user-service.
5. Chạy ứng dụng (Production)
Sau khi build, bạn có thể chạy tệp thực thi trực tiếp:
code
```bash
./bin/user-service
```
## API Endpoints
Dưới đây là danh sách các endpoint hiện có.
### Lấy thông tin người dùng qua ID
- **Endpoint**: GET /users/{id}
- **Mô tả**: Trả về thông tin chi tiết của một người dùng dựa trên ID được cung cấp.
- **Ví dụ (thành công)**:
```bash
curl http://localhost:8080/users/1
```      
**Phản hồi:**
```json
{
  "id": "1",
  "name": "Alice",
  "email": "alice@example.com"
}
```
- **Ví dụ (thất bại - không tìm thấy user):**
```bash
curl http://localhost:8080/users/99
```
**Phản hồi:*
```text
user not found
```
### Cấu hình
Ứng dụng sử dụng tệp configs/config.yaml để quản lý cấu hình. Hiện tại, chỉ có cấu hình cho port của server.
```yaml
server:
  port: 8080
```
### Kết quả
Khi chạy, chương trình sẽ in ra màn hình:

<img width="892" height="598" alt="image" src="https://github.com/user-attachments/assets/5c4fcdb1-9108-4dc9-a819-770f2f76d9b7" />

### Đóng góp
Mọi đóng góp đều được chào đón! Vui lòng tạo một Pull Request để đóng góp vào dự án.
### Giấy phép
Dự án này được cấp phép theo Giấy phép MIT.
