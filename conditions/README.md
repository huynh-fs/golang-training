# Hệ thống quản lý trạng thái đơn hàng E-commerce (Go CLI)

[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://golang.org/)

## Mô tả dự án

Đây là một ứng dụng dòng lệnh (CLI) được phát triển bằng Go để minh họa việc quản lý trạng thái đơn hàng trong một hệ thống E-commerce đơn giản. Mục tiêu chính của dự án là củng cố kiến thức và thực hành việc áp dụng các mệnh đề điều kiện (`if`, `else if`, `else`, `switch`, `type switch`, `fallthrough`) trong Go một cách hiệu quả và có cấu trúc.

Ứng dụng cho phép người dùng thực hiện các thao tác cơ bản như tạo đơn hàng, xem chi tiết, cập nhật trạng thái và nhận các gợi ý hành động dựa trên trạng thái hiện tại của đơn hàng.

## Tính năng chính

*   **Tạo đơn hàng mới:** Gán ID tự động, trạng thái ban đầu (`Chờ xử lý`/`Pending`), các mặt hàng, tổng giá trị và trạng thái thanh toán.
*   **Xem chi tiết đơn hàng:** Hiển thị tất cả thông tin liên quan đến một đơn hàng cụ thể.
*   **Cập nhật trạng thái đơn hàng:** Chuyển đổi trạng thái đơn hàng theo các quy tắc nghiệp vụ (ví dụ: `Chờ xử lý` -> `Đang xử lý`, nhưng không thể `Đã giao hàng` -> `Đang xử lý`).
*   **Kiểm tra và gợi ý hành động:** Cung cấp các thông báo hoặc khuyến nghị dựa trên trạng thái và các điều kiện khác của đơn hàng (ví dụ: cảnh báo đơn hàng chưa thanh toán đang được xử lý).
*   **Minh họa Type Switch:** Một chức năng riêng để bạn có thể thử nghiệm `type switch` với các kiểu dữ liệu khác nhau, làm nổi bật khả năng xử lý kiểu động của Go.
*   **Menu tương tác CLI:** Giao diện người dùng dựa trên văn bản đơn giản để dễ dàng tương tác.
*   **Xác thực đầu vào:** Kiểm tra tính hợp lệ của dữ liệu nhập vào từ người dùng.
*   **Hỗ trợ nhập trạng thái bằng tiếng Việt và tiếng Anh:** Để tăng tính linh hoạt và trải nghiệm người dùng.

## Cấu trúc dự án

Dự án được cấu trúc theo cách được khuyến nghị cho các ứng dụng Go, sử dụng các gói và thư mục riêng biệt để phân tách trách nhiệm:

```
conditions/
├── cmd/
│   └── main.go                  // Điểm vào chính của ứng dụng, xử lý menu CLI và khởi tạo manager.
├── internal/
│   ├── order/                   // Gói chứa các định nghĩa liên quan đến đơn hàng và logic nghiệp vụ.
│   │   ├── order.go             // Định nghĩa struct Order và các hằng số trạng thái.
│   │   └── manager.go           // Chứa struct OrderManager và các phương thức quản lý đơn hàng.
│   └── utils/                   // Gói chứa các hàm tiện ích chung.
│       └── input.go             // Hàm đọc đầu vào từ console.
├── go.mod                       // Định nghĩa module Go của dự án.
└── README.md                    // File tài liệu dự án này.
```

## Cách chạy dự án

Để chạy ứng dụng này, hãy làm theo các bước sau:

1.  **Clone repository (nếu có):**
    ```bash
    git clone https://github.com/huynh-fs/golang-training.git
    cd conditions
    ```
    (Nếu bạn chỉ có các file cục bộ, hãy điều hướng đến thư mục gốc của dự án `conditions/`).

2.  **Khởi tạo Go Module (nếu chưa):**
    Đảm bảo file `go.mod` đã được tạo. Nếu không, hoặc nếu bạn thay đổi đường dẫn module, chạy:
    ```bash
    go mod init github.com/yourusername/order-manager-project # Thay 'yourusername' bằng username GitHub của bạn
    ```

3.  **Tải xuống các phụ thuộc và dọn dẹp module:**
    ```bash
    go mod tidy
    ```

4.  **Chạy ứng dụng:**
    ```bash
    go run ./cmd
    ```

## Hướng dẫn sử dụng

Sau khi chạy ứng dụng, bạn sẽ thấy một menu CLI. Nhập số tương ứng với lựa chọn của bạn và nhấn Enter.

### Các lựa chọn:

1.  **Tạo đơn hàng mới:** Nhập các thông tin cần thiết như mặt hàng, tổng giá trị và trạng thái thanh toán (`Paid` hoặc `Unpaid`).
2.  **Xem chi tiết đơn hàng:** Nhập ID đơn hàng (ví dụ: `ORD001`) để xem thông tin chi tiết.
3.  **Cập nhật trạng thái đơn hàng:** Nhập ID đơn hàng và trạng thái mới.
    *   **Đầu vào trạng thái được chấp nhận:** Bạn có thể nhập cả tiếng Việt hoặc tiếng Anh (không phân biệt chữ hoa/thường):
        *   `Chờ xử lý` hoặc `Pending`
        *   `Đang xử lý` hoặc `Processing`
        *   `Đã vận chuyển` hoặc `Shipped`
        *   `Đã giao hàng` hoặc `Delivered`
        *   `Đã hủy` hoặc `Cancelled`
    *   **Quy tắc chuyển đổi trạng thái:**
        *   Từ `Chờ xử lý`: chỉ có thể chuyển sang `Đang xử lý` hoặc `Đã hủy`.
        *   Từ `Đang xử lý`: chỉ có thể chuyển sang `Đã vận chuyển` hoặc `Đã hủy`.
        *   Từ `Đã vận chuyển`: chỉ có thể chuyển sang `Đã giao hàng`.
        *   `Đã giao hàng` hoặc `Đã hủy`: không thể thay đổi trạng thái nữa.
4.  **Kiểm tra và gợi ý hành động cho đơn hàng:** Nhập ID đơn hàng để nhận các gợi ý hoặc cảnh báo dựa trên trạng thái và các điều kiện đặc biệt khác (ví dụ: đơn hàng đã vận chuyển lâu nhưng chưa giao).
5.  **Minh họa Type Switch:** Chức năng này sẽ tự động chạy một loạt các ví dụ `type switch` với các kiểu dữ liệu khác nhau (int, string, bool, *order.Order, float64).
6.  **Thoát (q):** Thoát khỏi ứng dụng.

## Các mệnh đề điều kiện được áp dụng

Dự án này được thiết kế để minh họa cách sử dụng tất cả các mệnh đề điều kiện trong Go:

*   **`if`, `else if`, `else`**: Được sử dụng rộng rãi để xác thực đầu vào, kiểm tra điều kiện đơn giản và điều khiển luồng logic chuyển đổi trạng thái.
    *   Ví dụ: Kiểm tra tổng giá trị đơn hàng, kiểm tra `found` khi tìm kiếm đơn hàng, kiểm tra `daysSinceShipped`.
*   **`switch` với biến/biểu thức (Tagged Switch)**: Được sử dụng để xử lý các lựa chọn menu chính và kiểm soát việc chuyển đổi trạng thái dựa trên trạng thái hiện tại của đơn hàng.
    *   Ví dụ: `switch choice` trong `main.go`, `switch order.Status` và `switch newStatus` trong `manager.go`.
*   **`switch` không có biến (Switch true)**: Được sử dụng để kiểm tra nhiều điều kiện boolean độc lập, hoạt động như một chuỗi `if-else if-else` có cấu trúc hơn.
    *   Ví dụ: Trong `CheckEligibilityAndSuggestions()` để đưa ra các gợi ý phức tạp dựa trên nhiều thuộc tính của đơn hàng.
*   **Khai báo biến ngắn trong `if`**: Cho phép khai báo và sử dụng biến chỉ trong phạm vi của câu lệnh `if` và các nhánh `else` liên quan.
    *   Ví dụ: `if daysSinceShipped := time.Since(order.ShippedDate).Hours() / 24; ...`.
*   **`fallthrough`**: Được sử dụng để buộc `switch` tiếp tục thực thi khối mã của `case` tiếp theo, bỏ qua điều kiện của nó.
    *   Ví dụ: Trong phần "Kiểm tra bổ sung" của `CheckEligibilityAndSuggestions()` để minh họa chuỗi hành động liên tiếp.
*   **`case` với nhiều giá trị**: Cho phép một `case` khớp với nhiều giá trị khác nhau.
    *   Ví dụ: `case StatusDelivered, StatusCancelled:` để xử lý cả hai trạng thái này cùng một lúc.
*   **`type switch`**: Được sử dụng để xác định và xử lý kiểu động của một giá trị `interface{}`.
    *   Ví dụ: Trong hàm `processGenericInput` (được gọi từ `main.go`) để minh họa cách xử lý các kiểu dữ liệu khác nhau.

