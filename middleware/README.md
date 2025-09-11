## Cập nhật và Cải tiến cho dự án `gin-api`

Dự án này được nâng cấp từ nền tảng `gin-api` ban đầu, bổ sung các tính năng quan trọng để tăng cường tính bảo mật, khả năng giám sát và độ tin cậy khi triển khai.

### Các cải tiến chính:

#### 1. Tích hợp hệ thống Middleware

Một thư mục mới `internal/middleware` đã được thêm vào để quản lý các tác vụ xuyên suốt (cross-cutting concerns) một cách tách biệt:

- **`LoggerMiddleware`**: Tự động ghi lại thông tin chi tiết của mỗi request (phương thức, đường dẫn, status code, độ trễ). Giúp việc theo dõi và gỡ lỗi trở nên dễ dàng hơn.
- **`AuthMiddleware`**: Bảo vệ tất cả các API endpoint trong nhóm `/api/v1/todos`. Mọi request đến các endpoint này đều phải cung cấp một `Bearer Token` hợp lệ trong header `Authorization`.

#### 2. Xác thực API bằng Bearer Token

- Toàn bộ API giờ đây đã được bảo mật. Client cần gửi kèm token được cấu hình trong file `.env` để có thể truy cập.
- **Ví dụ gọi API với cURL:**
  ```bash
  curl -X GET 'http://localhost:8080/api/v1/todos' \
    --header 'Authorization: Bearer 123456789'
  ```

#### 3. Cải tiến Swagger với Hỗ trợ Xác thực

- Tài liệu Swagger đã được cập nhật để hỗ trợ kiểm thử các endpoint được bảo vệ.
- **Cách sử dụng:**
  1.  Trên giao diện Swagger, nhấn nút **`Authorize`**.
  2.  Trong ô `value`, nhập token theo định dạng: `Bearer 123456789` (token được hardcode).
  3.  Sau khi xác thực, bạn có thể sử dụng tính năng "Try it out" cho tất cả API, Swagger sẽ tự động đính kèm header `Authorization` vào request.
