# Log Processor - Công cụ phân tích Log file song song

`Log Processor` là một công cụ dòng lệnh (CLI) hiệu suất cao được viết bằng Go, được thiết kế để phân tích và tổng hợp các file log lớn một cách nhanh chóng. Dự án này sử dụng các mẫu lập trình đồng thời (concurrency patterns) của Go, bao gồm Goroutines và Channels, để xử lý dữ liệu theo một pipeline song song.

Dự án này cũng là một ví dụ điển hình về cách cấu trúc một ứng dụng Go theo các tiêu chuẩn thực hành tốt nhất, tách biệt rõ ràng các mối quan tâm (separation of concerns) để dễ dàng bảo trì và mở rộng.

## Tính năng

-   **Xử lý song song:** Sử dụng một nhóm các worker goroutine để phân tích đồng thời nhiều dòng log, tận dụng tối đa CPU đa lõi.
-   **Pipeline xử lý dữ liệu:** Dữ liệu chảy qua một chuỗi các giai đoạn (Đọc -> Phân tích -> Tổng hợp) thông qua các channel, giảm thiểu tranh chấp và không cần sử dụng khóa (locks).
-   **Cấu hình linh hoạt:** Số lượng worker phân tích có thể được cấu hình dễ dàng thông qua file `configs/config.yaml`.
-   **Báo cáo tiến độ:** Cung cấp phản hồi thời gian thực về tiến trình xử lý đối với các file log lớn.
-   **Cấu trúc dự án chuyên nghiệp:** Tuân theo các nguyên tắc của "Standard Go Project Layout" để mã nguồn luôn sạch sẽ, dễ đọc và dễ bảo trì.

## Cấu trúc dự án

Dự án được tổ chức theo cấu trúc sau:

```
channels/
├── cmd/log-processor/        # Entry point của ứng dụng CLI
├── internal/             # Mã nguồn private của ứng dụng
│   ├── app/log-processor/    # Logic nghiệp vụ cốt lõi
│   │   ├── aggregator/   # Logic tổng hợp kết quả
│   │   ├── parser/       # Logic phân tích một dòng log
│   │   └── service/      # Đóng gói và điều phối pipeline
│   └── config/           # Logic đọc file cấu hình
├── configs/              # Các file cấu hình 
├── scripts/              # Các script tiện ích 
└── go.mod                # File quản lý module và dependencies
```

## Yêu cầu

-   [Go](https://golang.org/dl/) (phiên bản 1.18 trở lên)

## Cài đặt và Sử dụng

1.  **Clone repository:**
    ```bash
    git clone https://github.com/huynh-fs/golang-training.git
    cd channels
    ```

2.  **Cấu hình (Tùy chọn):**
    Mở file `configs/config.yaml` và điều chỉnh số lượng `parser_workers` để phù hợp với CPU của bạn. Giá trị mặc định là `4`.
    ```yaml
    processor:
      parser_workers: 8 # Ví dụ: thay đổi thành 8 workers
    ```

3.  **Tạo file log mẫu (Tùy chọn):**
    Để kiểm tra hiệu năng của công cụ, bạn có thể tạo một file log lớn bằng script được cung cấp.
    ```bash
    # Cấp quyền thực thi cho script (chỉ cần làm một lần)
    chmod +x scripts/generate_log.sh

    # Chạy script để tạo file large.log
    ./scripts/generate_log.sh
    ```

4.  **Chạy ứng dụng:**
    Sử dụng lệnh `go run` để biên dịch và chạy ứng dụng. Cung cấp đường dẫn đến file log bạn muốn phân tích làm đối số.
    ```bash
    go run ./cmd/log-processor 
    ```

## Các khái niệm kỹ thuật được áp dụng

Dự án này là một minh chứng thực tế cho các khái niệm cốt lõi trong lập trình đồng thời với Go:

-   **Goroutines:** Mỗi worker phân tích và aggregator đều chạy trong goroutine riêng, cho phép xử lý song song thực sự.
-   **Channels:** Được sử dụng làm "đường ống" để giao tiếp an toàn giữa các giai đoạn của pipeline (`Reader` -> `Parsers` -> `Aggregator`). Triết lý "chia sẻ bộ nhớ bằng cách giao tiếp" được áp dụng triệt để.
-   **Buffered Channels:** Giúp tách rời (decouple) các giai đoạn, cho phép giai đoạn producer (ví dụ: đọc file) có thể chạy trước mà không cần chờ consumer xử lý xong, giúp tăng thông lượng (throughput).
-   **`select` Statement:** Được sử dụng trong aggregator để xử lý đa nhiệm: vừa nhận dữ liệu mới để tổng hợp, vừa in ra báo cáo tiến độ theo một lịch trình định sẵn (`time.Ticker`).
-   **`sync.WaitGroup`:** Đảm bảo các goroutine chính đợi cho đến khi tất cả các worker hoàn thành nhiệm vụ trước khi đóng các channel tiếp theo, tránh tình trạng kết thúc sớm hoặc deadlock.
