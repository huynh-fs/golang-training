# Go Concurrent URL Title Fetcher

Một công cụ dòng lệnh (CLI) đơn giản được viết bằng Go để lấy tiêu đề HTML từ một danh sách các URL. Dự án này được xây dựng với mục đích chính là để minh họa các khái niệm cốt lõi về concurrency (tính đồng thời) trong Go, cũng như cấu trúc một dự án Go theo tiêu chuẩn.

## Tính năng

- **Lấy dữ liệu đồng thời**: Sử dụng Goroutines để gửi request đến nhiều URL cùng một lúc, giúp tăng tốc độ đáng kể so với việc xử lý tuần tự.
- **Giao diện tương tác**: Cung cấp một menu đơn giản cho phép người dùng nhập các URL một cách linh hoạt.
- **Nhập liệu thông minh**: Cho phép người dùng nhập từng URL trên mỗi dòng và kết thúc việc nhập bằng một dòng trống.
- **Cấu trúc dự án chuẩn**: Tuân theo cấu trúc "Standard Go Project Layout", phân tách rõ ràng giữa điểm vào ứng dụng (`cmd`), logic nghiệp vụ (`internal/cli`), và các thành phần cốt lõi có thể tái sử dụng (`internal/fetcher`).

## Cấu trúc dự án

Dự án được tổ chức theo cấu trúc sau để đảm bảo tính module hóa và dễ bảo trì:

```
concurrency/
├── cmd/
│   └── urlfetcher/
│       └── main.go         # Điểm vào (entrypoint) tối giản của ứng dụng.
├── internal/
│   ├── cli/
│   │   └── cli.go          # Xử lý logic giao diện dòng lệnh, menu và điều phối.
│   └── fetcher/
│       └── fetcher.go      # "Worker" chứa logic để lấy và xử lý một URL duy nhất.
├── go.mod                  # File quản lý module và các thư viện phụ thuộc.
└── README.md               # Tài liệu này.
```

- **`cmd/`**: Chứa code thực thi. Nhiệm vụ duy nhất của `main.go` là gọi logic trong package `internal/cli`.
- **`internal/`**: Chứa code nội bộ của ứng dụng.
  - **`cli/`**: Đóng vai trò "nhạc trưởng", điều khiển luồng của ứng dụng, tương tác với người dùng và điều phối các goroutine.
  - **`fetcher/`**: Chứa logic "thuần túy" để thực hiện một công việc cụ thể, có thể được tái sử dụng ở nơi khác nếu cần.

## Cách sử dụng

Bạn có thể chạy ứng dụng bằng `go run` hoặc biên dịch ra một file thực thi.

```bash
go run cmd/urlfetcher
```

### Ví dụ một phiên làm việc

Sau khi chạy ứng dụng, bạn sẽ thấy một menu tương tác:

```
--- URL Fetcher Menu ---
1. Lấy tiêu đề từ các URL
2. Thoát
Nhập lựa chọn của bạn: 1

Nhập từng URL và nhấn Enter. Nhấn Enter trên dòng trống để bắt đầu:
URL> golang.org
URL> github.com
URL> invalid-url.xyz
URL> https://runsystem.net/vi
URL> 
(Nhấn Enter tại đây để bắt đầu xử lý)

--- Kết quả ---
Lỗi khi lấy https://invalid-url.xyz: Get "https://invalid-url.xyz": dial tcp: lookup invalid-url.xyz: no such host
https://github.com -> "GitHub · Build and ship software on a single, collaborative platform · GitHub"
https://runsystem.net/vi -> "GMO-Z.com RUNSYSTEM"
https://golang.org -> "The Go Programming Language"
--- Hoàn thành --- 

--- URL Fetcher Menu ---
1. Lấy tiêu đề từ các URL
2. Thoát
Nhập lựa chọn của bạn: 2
Tạm biệt!
```

## Các khái niệm kỹ thuật được áp dụng

Dự án này là một bài thực hành về các khái niệm quan trọng trong Go:

- **Goroutines**: Mỗi URL được xử lý trong một goroutine riêng biệt, được khởi tạo bằng từ khóa `go`. Điều này cho phép các request HTTP được thực hiện đồng thời.
- **`sync.WaitGroup`**: Được sử dụng như một cơ chế đồng bộ hóa mạnh mẽ. Goroutine chính sẽ bị chặn tại `wg.Wait()` cho đến khi tất cả các goroutine worker gọi `wg.Done()`, đảm bảo chương trình không kết thúc trước khi tất cả công việc hoàn thành.
- **Phân tách các mối quan tâm (Separation of Concerns)**: Logic được phân tách rõ ràng: `main.go` chỉ khởi chạy, `cli.go` xử lý tương tác người dùng, và `fetcher.go` thực hiện công việc cốt lõi.
