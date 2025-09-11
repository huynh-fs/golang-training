# Ứng dụng Bingo Game bằng Golang

Đây là một ứng dụng dòng lệnh (CLI) mô phỏng trò chơi Bingo 75 bóng được viết bằng ngôn ngữ Go. Chương trình sẽ tự động tạo một tấm vé, gọi số ngẫu nhiên, kiểm tra điều kiện thắng và xuất toàn bộ kết quả ra file CSV.

Dự án được xây dựng theo **kiến trúc phân lớp (Layered Architecture)** với các tầng `handler`, `service`, và `model` để đảm bảo mã nguồn có tổ chức, rõ ràng, dễ bảo trì và mở rộng.

## Tính năng chính

-   **Tạo vé tự động**: Tự động tạo một tấm vé Bingo 5x5 hợp lệ theo đúng luật chơi (cột B: 1-15, I: 16-30, N: 31-45, G: 46-60, O: 61-75).
-   **Ô Miễn Phí (Free Space)**: Tuân thủ luật chơi với ô trung tâm là ô miễn phí, được đánh dấu sẵn.
-   **Gọi số ngẫu nhiên**: Tự động gọi số ngẫu nhiên không trùng lặp trong khoảng sau mỗi 2 giây.
-   **Kiểm tra thắng**: Kiểm tra điều kiện thắng (BINGO) trên hàng ngang, cột dọc và hai đường chéo.
-   **Tự động dừng**: Chương trình sẽ tự động kết thúc khi phát hiện có người thắng.
-   **Xuất kết quả**: Toàn bộ diễn biến và kết quả của ván chơi được ghi chi tiết vào file `bingo_result.csv`.

## Kiến trúc Dự án

Dự án áp dụng kiến trúc phân lớp để tách biệt rõ ràng các mối quan tâm (separation of concerns):

-   **`model`**: Lớp thấp nhất, chỉ chứa các định nghĩa cấu trúc dữ liệu (structs) của ứng dụng, không có logic nghiệp vụ.
-   **`service`**: Là nơi chứa toàn bộ logic nghiệp vụ (business logic). Lớp này điều phối các `model` để thực hiện các tác vụ cốt lõi của trò chơi.
-   **`handler`**: Lớp giao tiếp với thế giới bên ngoài. Nó nhận yêu cầu (trong trường hợp này là từ `main`), gọi các `service` tương ứng và xử lý việc xuất dữ liệu (ghi file).

Cấu trúc thư mục chi tiết:

```
/
├── cmd/
│   └── bingo/
│       └── main.go         # Điểm khởi chạy ứng dụng, điều phối service và handler
├── internal/
│   ├── handler/
│   │   └── file_writer.go  # Xử lý các tác vụ I/O như ghi file CSV
│   ├── model/
│   │   ├── ticket.go         # Định nghĩa cấu trúc dữ liệu cho tấm vé
│   │   └── result.go       # Định nghĩa cấu trúc dữ liệu cho kết quả
│   └── service/
│       ├── game_service.go # Logic điều khiển luồng chính của trò chơi
│       ├── ticket_service.go # Logic tạo và kiểm tra tấm vé Bingo
│       └── random_service.go # Logic tạo số ngẫu nhiên
├── go.mod                  # File định nghĩa module của dự án
└── bingo_result.csv              # File kết quả được tạo ra sau khi chạy
```

## Yêu cầu

-   Cần cài đặt **Go** (phiên bản 1.18 trở lên).

## Cài đặt và Chạy

1.  **Clone repository** về máy của bạn:
    ```sh
    git clone https://github.com/huynh-fs/golang-training.git
    ```

2.  Di chuyển vào thư mục dự án:
    ```sh
    cd file
    ```

3.  Chạy ứng dụng:
    ```sh
    go run ./cmd/bingo
    ```

Chương trình sẽ bắt đầu chạy trên terminal, hiển thị tấm vé ban đầu và lần lượt các số được gọi ra. Khi có người thắng, chương trình sẽ dừng lại và thông báo đã ghi kết quả vào file `bingo_result.csv`.

## Mô tả file `bingo_result.csv`

File `bingo_result.csv` chứa toàn bộ thông tin của ván chơi, được cấu trúc như sau:

-   **Dòng 1 đến 5**: Trạng thái ban đầu của tấm vé Bingo 5x5.
-   **Dòng 6**: Danh sách tất cả các số đã được gọi, được ngăn cách nhau bởi dấu cách (` `).
-   **Dòng 7**: Thông báo về dòng đã thắng (ví dụ: `BINGO theo hàng ngang 3` hoặc `BINGO theo đường chéo chính (\)`).
-   **Dòng 8 đến 12**: Trạng thái cuối cùng của tấm vé Bingo. Tại đây, những số đã được gọi sẽ được thay thế bằng số `0`.

### Ví dụ nội dung file `bingo_result.csv`

```csv
12,22,45,55,61
3,18,33,48,72
9,28,0,51,68
15,16,31,60,75
1,25,40,46,64
51 28 9 68 32 1 12 15 3 40 46
BINGO theo cột dọc 1
0,22,45,55,61
0,18,33,48,0
0,0,0,0,0
0,16,31,60,75
0,25,0,0,64
```