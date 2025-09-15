# Go Gin API - Advanced JWT Authentication & Clean Architecture

Dự án này là một API RESTful hoàn chỉnh, được xây dựng bằng Go và Gin, thể hiện một kiến trúc ứng dụng production-ready. Nó cung cấp các chức năng CRUD cho một ứng dụng "Todo List" và tích hợp một hệ thống **xác thực JWT (Access & Refresh Token)** đầy đủ với khả năng thu hồi token.

Điểm nhấn của dự án là việc áp dụng các nguyên tắc **Clean Architecture**, **Dependency Injection (DI)**, và một **chiến lược kiểm thử tập trung** để đảm bảo chất lượng, độ tin cậy và khả năng bảo trì của mã nguồn.

- **Ngôn ngữ:** Go
- **Framework:** Gin
- **ORM:** GORM
- **Cơ sở dữ liệu:** PostgreSQL
- **Xác thực:** JWT (Access & Refresh Tokens)
- **Kiểm thử:** Unit Testing chuyên sâu với Mockery & Testify
- **Triển khai:** Docker & Docker Compose
- **Tài liệu API:** Swagger

## Kiến trúc và Các tính năng chính

- **Kiến trúc phân tầng (Handler, Service, Repository):** Tách biệt rõ ràng các mối quan tâm, giúp mã nguồn linh hoạt và dễ kiểm thử.
- **Dependency Injection (DI):** Không sử dụng biến toàn cục. Tất cả các phụ thuộc được "tiêm" từ `main.go`.
- **Hệ thống xác thực JWT chuyên nghiệp:** Bao gồm Access/Refresh Token, cơ chế thu hồi (Revocation) và xoay vòng (Rotation).
- **Môi trường Dockerized hoàn chỉnh:** Triển khai nhất quán và đáng tin cậy.

---

## Chiến lược Kiểm thử (Testing)

Dự án áp dụng một chiến lược kiểm thử tập trung và hiệu quả, đặt trọng tâm vào nơi chứa nhiều logic nghiệp vụ phức tạp nhất: **Tầng Service**. Bằng cách đảm bảo tầng Service được kiểm thử 100%, chúng ta có thể tự tin rằng cốt lõi của ứng dụng hoạt động chính xác trong mọi tình huống.

### Trụ cột chính: Unit Testing Tầng Service

Mục tiêu của Unit Test là kiểm tra logic nghiệp vụ (business logic) của từng `Service` một cách **hoàn toàn cô lập**, không phụ thuộc vào database hay các thành phần bên ngoài khác. Điều này giúp bộ test chạy cực kỳ nhanh, đáng tin cậy và có thể được tích hợp dễ dàng vào các quy trình CI/CD.

#### Các nguyên tắc và công cụ chính:

- **Kiến trúc hướng Interface (Interface-Driven Architecture):**
  Các `Service` không phụ thuộc trực tiếp vào GORM, mà phụ thuộc vào các `interface` (`UserRepository`, `TodoRepository`). Đây là nền tảng cốt lõi cho phép chúng ta áp dụng mocking một cách hiệu quả.

- **Mocking toàn diện với `mockery`:**
  Chúng ta sử dụng `mockery` để tự động sinh ra các đối tượng "giả" (mocks) từ các `interface` repository. Trong các bài test, chúng ta có thể ra lệnh cho các mock này trả về bất kỳ dữ liệu hoặc lỗi nào mong muốn. Điều này cho phép chúng ta dễ dàng kiểm tra tất cả các luồng logic, bao gồm cả các trường hợp lỗi khó tái tạo (ví dụ: lỗi kết nối database, dữ liệu không tìm thấy, v.v.).

- **Table-Driven Tests:**
  Tất cả các bài test cho Service đều được viết theo pattern Table-Driven. Mỗi hàm test định nghĩa một "bảng" các trường hợp kiểm thử, sau đó chạy chúng trong một vòng lặp duy nhất. Phương pháp này giúp code test:
  - **Ngắn gọn và không lặp lại (DRY).**
  - **Dễ đọc:** Tất cả các kịch bản được liệt kê rõ ràng.
  - **Cực kỳ dễ mở rộng:** Thêm một trường hợp test mới chỉ đơn giản là thêm một phần tử vào slice.

### Về việc Kiểm thử Tầng Handler

Tầng Handler trong kiến trúc này được thiết kế để **mỏng nhất có thể**. Trách nhiệm chính của nó chỉ bao gồm:

1.  Đọc và xác thực dữ liệu đầu vào (binding DTO).
2.  Gọi phương thức tương ứng ở tầng Service.
3.  Chuyển đổi kết quả (dữ liệu hoặc lỗi) từ Service thành một HTTP response thích hợp.

Với bộ unit test toàn diện cho tầng Service, việc viết các bài test tích hợp đầy đủ cho tầng Handler có thể dẫn đến sự trùng lặp trong việc kiểm tra logic nghiệp vụ. Do đó, chiến lược của dự án là tập trung nguồn lực vào việc đảm bảo tầng Service - nơi chứa đựng bộ não của ứng dụng - hoạt động một cách hoàn hảo.

### Cách chạy bộ Unit Test

Các bài test này không yêu cầu bất kỳ phụ thuộc bên ngoài nào (như database) phải đang chạy.

```bash
# Chạy tất cả các unit test cho tầng service
go test -v ./internal/service/...
```

Để thực thi tất cả các file test có trong dự án, bạn có thể chạy từ thư mục gốc:

```bash
go test -v ./...
```
