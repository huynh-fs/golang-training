## Hệ thống Xác thực JWT cho dự án Todo List

Dự án này tích hợp một hệ thống xác thực chuyên nghiệp sử dụng JSON Web Tokens (JWT), được thiết kế để cân bằng giữa bảo mật và trải nghiệm người dùng. Kiến trúc này dựa trên việc sử dụng hai loại token: **Access Token** và **Refresh Token**.

### Các thành phần chính

#### 1. Access Token

- **Mục đích:** Dùng để xác thực các request truy cập tài nguyên được bảo vệ (ví dụ: các API trong nhóm `/api/v1/todos/`).
- **Đặc điểm:**
  - **Thời gian sống ngắn (ví dụ: 15 phút)** để giảm thiểu rủi ro nếu token bị lộ.
  - **Stateless:** Server xác thực token chỉ bằng chữ ký bí mật (`JWT_ACCESS_SECRET`) mà không cần truy vấn database, giúp tăng hiệu năng.
- **Sử dụng:** Được gửi kèm trong header `Authorization` của mỗi request API theo định dạng `Bearer <access_token>`.

#### 2. Refresh Token

- **Mục đích:** Dùng để yêu cầu một cặp token mới (cả Access và Refresh) khi `Access Token` đã hết hạn.
- **Đặc điểm:**
  - **Thời gian sống dài (ví dụ: 7 ngày)**, cho phép người dùng duy trì phiên đăng nhập mà không cần nhập lại mật khẩu.
  - **Stateful & Có thể thu hồi:** Refresh Token được lưu trữ trong cơ sở dữ liệu. Điều này là mấu chốt để thực hiện chức năng đăng xuất và thu hồi.
- **Sử dụng:** Chỉ được gửi đến một endpoint duy nhất là `/auth/refresh`.

### Các tính năng bảo mật nâng cao

#### Thu hồi Token (Token Revocation)

- **Cách hoạt động:** Khi người dùng đăng xuất (`/auth/logout`), refresh token tương ứng sẽ bị **xóa khỏi database**.
- **Lợi ích:** Điều này ngay lập tức vô hiệu hóa khả năng tạo token mới từ refresh token đó, kể cả khi nó chưa hết hạn. Đây là một biện pháp bảo mật quan trọng để xử lý các trường hợp như người dùng bị mất thiết bị hoặc muốn đăng xuất khỏi tất cả các thiết bị.

#### Xoay vòng Token (Token Rotation)

- **Cách hoạt động:** Mỗi khi client sử dụng một refresh token hợp lệ để gọi `/auth/refresh`, một cặp Access và Refresh token **hoàn toàn mới** sẽ được tạo ra. Refresh token cũ sẽ bị vô hiệu hóa (xóa khỏi DB).
- **Lợi ích:** Tính năng này giúp ngăn chặn việc tái sử dụng refresh token đã bị lộ. Nếu kẻ tấn công đánh cắp và sử dụng refresh token, người dùng thật sẽ ngay lập tức phát hiện ra khi refresh token của họ không còn hợp lệ.

### Luồng xác thực (Authentication Flow)

1.  **Đăng ký (`POST /auth/register`):** Người dùng tạo tài khoản với `username` và `password`.
2.  **Đăng nhập (`POST /auth/login`):**
    - Client gửi `username` và `password`.
    - Server xác thực, sau đó tạo và trả về một cặp `access_token` và `refresh_token`.
3.  **Truy cập API (`GET /api/v1/todos`):**
    - Client đính kèm `access_token` vào header `Authorization: Bearer <access_token>`.
    - Server xác thực token và trả về dữ liệu.
4.  **Làm mới Token (`POST /auth/refresh`):**
    - Khi `access_token` hết hạn (server trả lỗi `401 Unauthorized`), client gửi `refresh_token` trong body request.
    - Server kiểm tra `refresh_token` trong DB:
      - Nếu hợp lệ, server tạo một cặp token **mới**, xóa token cũ khỏi DB, và trả về cặp token mới cho client.
      - Nếu không hợp lệ (đã bị thu hồi), server trả lỗi.
5.  **Đăng xuất (`POST /auth/logout`):**
    - Client gửi `refresh_token` muốn thu hồi.
    - Server tìm và xóa `refresh_token` đó khỏi DB. Mọi nỗ lực sử dụng token này trong tương lai sẽ thất bại.

### Cách sử dụng và kiểm thử

#### Sử dụng cURL

```bash
# 1. Đăng nhập để lấy token
TOKENS=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"username": "testuser", "password": "password123"}')

ACCESS_TOKEN=$(echo $TOKENS | jq -r '.access_token')

# 2. Gọi API được bảo vệ với Access Token
curl -X GET http://localhost:8080/api/v1/todos \
  -H "Authorization: Bearer $ACCESS_TOKEN"
```

#### Sử dụng Swagger UI

1.  Truy cập [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).
2.  Sử dụng endpoint `/auth/login` để lấy `access_token` và `refresh_token`.
3.  Nhấn vào nút **`Authorize`** ở góc trên bên phải.
4.  Trong hộp thoại, tại mục `BearerAuth`, nhập giá trị theo định dạng `Bearer <your_access_token>`.
5.  Nhấn **Authorize**. Giờ đây, bạn có thể kiểm thử tất cả các API được bảo vệ.
