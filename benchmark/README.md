# Go Gin API - High-Performance & Test-Driven Development

Dự án này là một API RESTful hiệu năng cao, được xây dựng bằng Go và Gin, thể hiện một kiến trúc ứng dụng được tối ưu và kiểm thử toàn diện. Ngoài việc cung cấp các chức năng CRUD và hệ thống **xác thực JWT (Access & Refresh Token)**, dự án này đặt một trọng tâm đặc biệt vào việc **đo lường và đảm bảo hiệu năng** thông qua một bộ benchmark chuyên sâu.

Kiến trúc dự án dựa trên các nguyên tắc **Clean Architecture** và **Dependency Injection**, tạo nền tảng vững chắc cho việc viết mã nguồn không chỉ đúng về mặt logic mà còn nhanh về mặt thực thi.

- **Ngôn ngữ:** Go
- **Framework:** Gin
- **ORM:** GORM
- **Cơ sở dữ liệu:** PostgreSQL
- **Xác thực:** JWT (Access & Refresh Tokens)
- **Kiểm thử:** Unit Testing & Benchmarking chuyên sâu
- **Triển khai:** Docker & Docker Compose
- **Tài liệu API:** Swagger

## Kiến trúc và Các tính dung chính

- **Kiến trúc phân tầng (Handler, Service, Repository):** Tách biệt rõ ràng các mối quan tâm, giúp mã nguồn linh hoạt và dễ kiểm thử.
- **Dependency Injection (DI):** Không sử dụng biến toàn cục. Tất cả các phụ thuộc được "tiêm" từ `main.go`.
- **Hệ thống xác thực JWT chuyên nghiệp:** Bao gồm Access/Refresh Token, cơ chế thu hồi (Revocation) và xoay vòng (Rotation).
- **Môi trường Dockerized hoàn chỉnh:** Triển khai nhất quán và đáng tin cậy.

---

## Chiến lược Hiệu năng & Benchmark

Chất lượng của một API không chỉ nằm ở tính đúng đắn của logic mà còn ở tốc độ phản hồi. Dự án này áp dụng một chiến lược benchmark nghiêm ngặt để xác định các điểm nghẽn tiềm tàng, đo lường hiệu suất của các thành phần cốt lõi và đảm bảo ứng dụng hoạt động hiệu quả dưới tải.

Chúng tôi sử dụng công cụ benchmark tích hợp của Go, kết hợp với các kỹ thuật setup nâng cao để có được kết quả đo lường chính xác và đáng tin cậy.

### Hai trụ cột của Benchmarking

#### 1. Unit Benchmarks (Sử dụng Mock)

- **Mục tiêu:** Đo lường hiệu năng của **thuật toán và logic Go thuần túy** trong tầng Service. Các benchmark này trả lời câu hỏi: "Logic của tôi (ví dụ: hash password, tạo JWT, xử lý dữ liệu) có nhanh không, và nó cấp phát bao nhiêu bộ nhớ?"
- **Phương pháp:**
  - Sử dụng các đối tượng "giả" (mocks) được tạo bởi `mockery` để loại bỏ hoàn toàn độ trễ của database và network.
  - Kết quả đo lường cực kỳ ổn định, lặp lại được và chỉ phản ánh hiệu suất của mã lệnh.
  - Sử dụng cờ `-benchmem` để phân tích chi tiết số lần và dung lượng bộ nhớ được cấp phát (`allocs/op` và `B/op`), giúp tối ưu hóa việc sử dụng bộ nhớ và giảm tải cho Garbage Collector.

#### 2. Integration Benchmarks (Sử dụng Database thật)

- **Mục tiêu:** Đo lường hiệu năng của **toàn bộ luồng hoạt động từ service đến database**. Các benchmark này trả lời câu hỏi: "Mất bao lâu để thực hiện một thao tác hoàn chỉnh, bao gồm cả việc GORM dịch và thực thi câu lệnh SQL?"
- **Phương pháp:**
  - Kết nối đến một container PostgreSQL **thật** và **riêng biệt** (`db-test`) được quản lý bởi Docker Compose.
  - Sử dụng `TestMain` và các hàm benchmark helper để **setup và dọn dẹp dữ liệu mẫu (seed data)** một cách an toàn, đảm bảo môi trường benchmark luôn sạch sẽ và nhất quán.
  - Cung cấp một cái nhìn thực tế về hiệu năng của các câu query và chi phí I/O.

### Cách đọc kết quả

Output sẽ có dạng:
`BenchmarkAuthService_Login_Unit-8 13479 87882 ns/op 375 B/op 10 allocs/op`

- `13479`: Số lần vòng lặp đã chạy.
- `87882 ns/op`: **Thời gian trung bình** cho một lần thực thi (~87 microsecond).
- `375 B/op`: **Số byte bộ nhớ** được cấp phát mỗi lần chạy.
- `10 allocs/op`: **Số lần cấp phát bộ nhớ** mỗi lần chạy.

Bằng cách thường xuyên chạy và phân tích các kết quả này, chúng ta có thể tự tin đưa ra các quyết định tối ưu hóa dựa trên dữ liệu thực tế, đảm bảo ứng dụng luôn duy trì hiệu năng cao khi phát triển và mở rộng.
