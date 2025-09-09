# Ứng dụng Bingo Game bằng Golang

Đây là một ứng dụng dòng lệnh (CLI) mô phỏng trò chơi Bingo 75 bóng được viết bằng ngôn ngữ Go. Chương trình sẽ tự động tạo một tấm vé, gọi số ngẫu nhiên, kiểm tra điều kiện thắng và xuất toàn bộ kết quả ra file CSV.

## Tính năng chính

-   **Tạo vé tự động**: Tự động tạo một tấm vé Bingo 5x5 hợp lệ theo đúng luật chơi (cột B: 1-15, I: 16-30, N: 31-45, G: 46-60, O: 61-75).
-   **Ô Miễn Phí (Free Space)**: Tuân thủ luật chơi với ô trung tâm là ô miễn phí, được đánh dấu sẵn.
-   **Gọi số ngẫu nhiên**: Tự động gọi số ngẫu nhiên không trùng lặp trong khoảng sau mỗi 2 giây.
-   **Kiểm tra thắng**: Kiểm tra điều kiện thắng (BINGO) trên hàng ngang, cột dọc và hai đường chéo.
-   **Tự động dừng**: Chương trình sẽ tự động kết thúc khi phát hiện có người thắng.
-   **Xuất kết quả**: Toàn bộ diễn biến và kết quả của ván chơi được ghi chi tiết vào file `bingo_result.csv`.

## Cấu trúc dự án

Dự án tuân theo cấu trúc thư mục chuẩn của Go để phân tách rõ ràng các nghiệp vụ:

```
file/
├── cmd/
│   └── bingo/
│       └── main.go         # Điểm khởi chạy chính của ứng dụng
├── internal/
│   ├── bingo/              # Logic xử lý tấm vé Bingo (tạo, kiểm tra thắng)
│   │   └── ticket.go
│   ├── service/               # Logic điều khiển luồng chính của trò chơi
│   │   └── game.go
│   ├── output/             # Logic ghi kết quả ra file CSV
│   │   └── writer.go
│   └── random/             # Logic tạo số ngẫu nhiên không trùng lặp
│       └── generator.go
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

Chương trình sẽ bắt đầu chạy trên terminal, hiển thị tấm vé ban đầu và lần lượt các số được gọi ra. Khi có người thắng, chương trình sẽ dừng lại và thông báo đã ghi kết quả vào file `bingo-result.csv`.

## Mô tả file `bingo-result.csv`

File `bingo-result.csv` chứa toàn bộ thông tin của ván chơi, được cấu trúc như sau:

-   **Dòng 1 đến 5**: Trạng thái ban đầu của tấm vé Bingo 5x5.
-   **Dòng 6**: Danh sách tất cả các số đã được gọi, được ngăn cách nhau bởi dấu cách (` `).
-   **Dòng 7**: Thông báo về dòng đã thắng (ví dụ: `BINGO theo hàng ngang 3` hoặc `BINGO theo đường chéo chính (\)`).
-   **Dòng 8 đến 12**: Trạng thái cuối cùng của tấm vé Bingo. Tại đây, những số đã được gọi sẽ được thay thế bằng số `0`.

### Ví dụ nội dung file `output.csv`

```csv
12,28,34,59,70
7,27,45,48,69
5,17,0,51,75
6,18,38,60,63
10,20,33,52,64
57 34 26 33 73 49 10 36 8 22 40 45 52 59 19 65 29 71 20 24 35 9 31 55 50 30 60 4 27 37 43 56 68 61 46 66 38
BINGO theo hàng dọc 3
12,28,0,0,70
7,0,0,48,69
5,17,0,51,75
6,18,0,0,63
0,0,0,0,64
```