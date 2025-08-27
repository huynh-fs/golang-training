# Go Web Adventure

<img width="1584" height="606" alt="image" src="https://github.com/user-attachments/assets/f6350392-f027-4ecc-96d3-14dc867e90e5" />

## Giới thiệu

**Go Web Adventure** là một trò chơi phiêu lưu văn bản đơn giản được xây dựng bằng Go, sử dụng các khái niệm cơ bản về web server (`net/http`), HTML, CSS và JavaScript. Dự án này minh họa cách sử dụng hai cấu trúc dữ liệu quan trọng nhất trong Go là **Slice** và **Map** để quản lý trạng thái trò chơi, bao gồm thế giới game, các phòng, vật phẩm và hành trang của người chơi.

Người chơi có thể khám phá các phòng, nhặt đồ vật và quản lý hành trang của mình thông qua một giao diện web đơn giản.

## Các tính năng chính

*   **Khám phá thế giới:** Di chuyển giữa các phòng được kết nối.
*   **Nhặt vật phẩm:** Thu thập các vật phẩm từ các phòng và thêm vào hành trang.
*   **Quản lý hành trang:** Xem các vật phẩm người chơi đang mang.
*   **Giao diện Web:** Tương tác với trò chơi thông qua trình duyệt web.
*   **Sử dụng Slice và Map:** Minh họa rõ ràng cách các cấu trúc dữ liệu này được áp dụng trong một ứng dụng thực tế.
*   **Xử lý đồng thời:** Sử dụng `sync.Mutex` để bảo vệ trạng thái game trong môi trường web đa luồng.

## Cách Slice và Map được sử dụng

Dự án này là một ví dụ tuyệt vời về việc ứng dụng Slice và Map để xây dựng một hệ thống quản lý dữ liệu linh hoạt và hiệu quả.

### 1. Map (`map[KeyType]ValueType`)

Map được sử dụng để lưu trữ các tập hợp khóa-giá trị, cho phép truy cập dữ liệu nhanh chóng bằng khóa.

*   **`world map[string]Room`:**
    *   **Mục đích:** Lưu trữ toàn bộ bản đồ trò chơi. Mỗi phòng được định danh bằng một `string` duy nhất (ID phòng) và được liên kết với một `Room` struct chứa thông tin chi tiết về phòng đó.
    *   **Ứng dụng:**
        *   **Truy cập nhanh:** Khi người chơi di chuyển, chúng ta có thể nhanh chóng tìm kiếm thông tin của phòng đích bằng cách sử dụng tên phòng làm khóa (ví dụ: `world[player.CurrentRoom]`).
        *   **Thêm/Cập nhật:** Dễ dàng thêm phòng mới hoặc cập nhật thông tin của một phòng hiện có.
        *   **Kiểm tra sự tồn tại:** Sử dụng cú pháp `value, ok := myMap[key]` để kiểm tra xem một phòng có tồn tại trong thế giới game hay không.

*   **`Exits map[string]string` trong `Room` struct:**
    *   **Mục đích:** Mỗi `Room` có một map riêng để định nghĩa các lối thoát từ phòng đó.
    *   **Ứng dụng:**
        *   **Định tuyến:** Khóa là hướng đi (ví dụ: `"north"`, `"south"`), và giá trị là tên của phòng đích. Điều này cho phép hệ thống dễ dàng xác định phòng mà người chơi sẽ đến khi di chuyển theo một hướng cụ thể.
        *   **Liệt kê lối thoát:** Dễ dàng lặp qua map này để hiển thị tất cả các hướng mà người chơi có thể đi từ phòng hiện tại.

### 2. Slice (`[]Type`)

Slice được sử dụng để lưu trữ các danh sách động của các phần tử cùng kiểu, cho phép thêm, xóa và truy cập phần tử một cách linh hoạt.

*   **`Items []string` trong `Room` struct:**
    *   **Mục đích:** Lưu trữ danh sách các vật phẩm hiện có trong một căn phòng cụ thể.
    *   **Ứng dụng:**
        *   **Khởi tạo:** Được khởi tạo là một slice rỗng hoặc với các vật phẩm ban đầu.
        *   **Hiển thị:** Lặp qua slice này để liệt kê tất cả các vật phẩm mà người chơi có thể nhìn thấy trong phòng.
        *   **Xóa vật phẩm:** Khi người chơi nhặt một vật phẩm, chúng ta tạo một slice `Items` mới cho phòng bằng cách bỏ qua vật phẩm đã nhặt, minh họa cách "xóa" phần tử khỏi slice trong Go.

*   **`Inventory []string` trong `Player` struct:**
    *   **Mục đích:** Lưu trữ danh sách các vật phẩm mà người chơi đang mang theo.
    *   **Ứng dụng:**
        *   **Khởi tạo:** Được khởi tạo là một slice rỗng khi trò chơi bắt đầu.
        *   **Thêm vật phẩm:** Sử dụng hàm `append()` để thêm vật phẩm vào hành trang khi người chơi nhặt chúng.
        *   **Hiển thị:** Lặp qua slice này để hiển thị tất cả các vật phẩm trong hành trang của người chơi.
        *   **Kiểm tra rỗng:** Sử dụng `len(player.Inventory) == 0` để kiểm tra xem hành trang có trống không.

### 3. Đồng thời (Concurrency) với `sync.Mutex`

Trong môi trường web, nhiều yêu cầu từ client có thể đến server cùng lúc. Để đảm bảo rằng trạng thái game (biến `world` và `player` toàn cục) được truy cập và sửa đổi một cách an toàn mà không gây ra lỗi đồng thời (race conditions), chúng ta sử dụng `sync.Mutex`.

*   **`gameMutex sync.Mutex`:**
    *   **Mục đích:** Bảo vệ các biến trạng thái game toàn cục.
    *   **Ứng dụng:**
        *   `gameMutex.Lock()`: Được gọi trước khi bất kỳ thao tác đọc hoặc ghi nào lên `world` hoặc `player` diễn ra. Điều này đảm bảo chỉ một goroutine có thể truy cập các biến này tại một thời điểm.
        *   `defer gameMutex.Unlock()`: Được gọi sau khi các thao tác hoàn tất (hoặc hàm kết thúc), giải phóng khóa để các goroutine khác có thể truy cập.

## Yêu cầu

*   Go 1.16 trở lên

## Cách chạy dự án

Thực hiện theo các bước sau để thiết lập và chạy trò chơi phiêu lưu văn bản trên web của bạn:

1.  **Clone mục dự án:**
    ```bash
    git clone https://github.com/huynh-fs/golang-training.git
    ```

2.  **Chạy server Go:**
    Mở terminal hoặc command prompt, điều hướng đến thư mục `slice-map` và chạy lệnh:
    ```bash
    go run main.go
    ```
    Bạn sẽ thấy thông báo: `Server đang chạy tại http://localhost:8080`

5.  **Truy cập trò chơi:**
    Mở trình duyệt web của bạn và truy cập địa chỉ:
    [http://localhost:8080](http://localhost:8080)

Bây giờ bạn có thể tương tác với trò chơi bằng cách gõ các lệnh vào ô input và nhấn "Gửi" hoặc Enter.

## Các lệnh trong game

*   `go [hướng]` - Di chuyển theo hướng (ví dụ: `go north`, `go east`)
*   `take [vật phẩm]` - Nhặt vật phẩm (ví dụ: `take chìa khóa`, `take ngọn đuốc`)
*   `inventory` / `inv` - Xem hành trang của bạn
*   `look` - Mô tả lại căn phòng hiện tại
*   `help` - Hiển thị danh sách các lệnh
*   `quit` - Thoát trò chơi (chỉ thoát khỏi giao diện web, không tắt server)
