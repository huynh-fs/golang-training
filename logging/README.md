# Go Logging Demo CLI

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://go.dev/)

Đây là một ứng dụng Command-Line Interface (CLI) đơn giản được viết bằng Go để minh họa và so sánh các thư viện logging phổ biến. Mục đích của dự án này là giúp các lập trình viên Go có cái nhìn trực quan về cách hoạt động, cú pháp và output của từng thư viện, từ đó đưa ra lựa chọn phù hợp cho dự án của mình.

## Tính năng

- Minh họa việc ghi log bằng thư viện chuẩn `log`.
- Minh họa thư viện structured logging mới nhất của thư viện chuẩn `slog` (yêu cầu Go 1.21+).
- Minh họa các thư viện logging phổ biến của bên thứ ba:
    - **Zerolog**: Hiệu năng cao, zero-allocation.
    - **Zap**: Hiệu năng cực cao, phát triển bởi Uber.
    - **Logrus**: API linh hoạt và thân thiện (hiện đang trong chế độ bảo trì).
- Giao diện menu đơn giản để chọn và xem output của từng thư viện.

## Cấu trúc dự án

```
logging/
├── cmd/
│   └── logging/
│       └── main.go         # Điểm khởi đầu của ứng dụng
├── internal/
│   ├── cli/
│   │   └── cli.go          # Logic xử lý giao diện CLI (menu, input)
│   └── loggers/
│       ├── logrus.go       # Logic minh họa cho Logrus
│       ├── slog.go         # Logic minh họa cho Slog
│       ├── stdlib.go       # Logic minh họa cho package log chuẩn
│       ├── zap.go          # Logic minh họa cho Zap
│       └── zerolog.go      # Logic minh họa cho Zerolog
├── go.mod
├── go.sum
└── README.md
```

## Yêu cầu

- **Go**: Phiên bản `1.21` trở lên (bắt buộc để sử dụng package `slog`).

## Cài đặt và Chạy

1.  **Clone repository:**
    ```bash
    git clone https://github.com/huynh-fs/golang-training.git
    cd logging
    ```

2.  **Tải các dependencies:**
    Lệnh này sẽ tự động tải các thư viện (Zap, Zerolog, Logrus) được định nghĩa trong code.
    ```bash
    go mod tidy
    ```

3.  **Chạy ứng dụng:**
    Từ thư mục gốc của dự án, thực thi lệnh sau:
    ```bash
    go run ./cmd/logging
    ```

## Sử dụng

Sau khi chạy, một menu sẽ hiển thị trên terminal. Bạn chỉ cần nhập số tương ứng với thư viện logging bạn muốn xem và nhấn `Enter`.

**Ví dụ menu:**
```text
Chọn một thư viện để minh họa ghi log:
1. Standard Library (log)
2. Zerolog
3. Zap
4. Logrus
5. Slog (Standard Library, Go 1.21+)
6. Thoát
Nhập lựa chọn của bạn:
```

**Ví dụ output khi chọn Slog:**
```text
--- Bắt đầu minh họa Slog (Go 1.21+) ---

>>> Minh họa TextHandler:
time=2025-09-09T09:30:04.908+07:00 level=DEBUG msg="Đang kết nối tới database..." host=localhost:5432
time=2025-09-09T09:30:04.909+07:00 level=INFO msg="Request đã được xử lý thành công" method=POST latency_ms=78

>>> Minh họa JSONHandler:
{"time":"2025-09-09T09:30:04.9098605+07:00","level":"WARN","msg":"API key sắp hết hạn","key_id":"a1b2c3d4","days_left":3}
{"time":"2025-09-09T09:30:04.9099422+07:00","level":"ERROR","msg":"Xác thực thất bại","request_details":{"ip_address":"192.168.1.100","user_agent":"Go-http-client/1.1"}}

Slog là giải pháp structured logging chính thức, mạnh mẽ và có sẵn trong Go.
--- Kết thúc minh họa Slog ---
```

## Tổng quan về các thư viện

| Thư viện | Đặc điểm chính | Trường hợp sử dụng tốt nhất |
| :--- | :--- | :--- |
| **`log` (Stdlib)** | • Đơn giản, có sẵn<br>• Log không cấu trúc | Script nhỏ, ví dụ nhanh, không cần dependency. |
| **`logrus`** | • API thân thiện, dễ dùng<br>• Có `Hooks` mạnh mẽ<br>• *Đã vào chế độ bảo trì* | Bảo trì các dự án cũ đang sử dụng. **Không khuyến khích cho dự án mới.** |
| **`zap`** | • Hiệu năng cực cao<br>• Tối ưu hóa memory allocation<br>• Structured logging | Các hệ thống yêu cầu hiệu năng logging tối đa, throughput cực lớn. |
| **`zerolog`** | • Hiệu năng rất cao, zero-allocation<br>• API liền mạch (fluent)<br>• Tập trung vào output JSON | Tương tự Zap, khi cần hiệu năng cao và API đơn giản, dễ sử dụng. |
| **`slog` (Stdlib)** | • Có sẵn trong Go 1.21+<br>• Structured logging<br>• Hiệu năng tốt, API linh hoạt | **Lựa chọn mặc định cho hầu hết các dự án mới** để giảm dependency. |