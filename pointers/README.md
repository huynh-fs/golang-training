# Go: Demo về Con trỏ và Cấu trúc Dự án

Dự án này là một ví dụ đơn giản được viết bằng Go để minh họa các khái niệm cốt lõi về **con trỏ (pointers)**, sự khác biệt giữa **truyền tham trị (pass-by-value)** và cách dùng con trỏ để mô phỏng **truyền tham chiếu (pass-by-reference)**.

Tất cả được đặt trong một cấu trúc thư mục rõ ràng, phân tách các thành phần của ứng dụng thay vì gộp tất cả vào file `main.go`.

## Các khái niệm chính được minh họa

1.  **Con trỏ (`*`) và Địa chỉ (`&`)**:
    *   Cách lấy địa chỉ bộ nhớ của một biến bằng toán tử `&`.
    *   Cách khai báo một biến con trỏ và cách truy cập giá trị mà nó trỏ tới bằng toán tử `*`.

2.  **Truyền tham trị (Pass-by-value)**:
    *   Chứng minh rằng Go luôn truyền tham số cho hàm bằng cách tạo ra một **bản sao (copy)**.
    *   Cho thấy việc thay đổi giá trị của tham số bên trong hàm sẽ **không** ảnh hưởng đến biến gốc bên ngoài.

3.  **Dùng Con trỏ để thay đổi giá trị gốc**:
    *   Cách truyền một con trỏ (địa chỉ bộ nhớ) vào hàm.
    *   Làm thế nào hàm có thể sử dụng con trỏ đó để sửa đổi giá trị của biến gốc, một kỹ thuật cần thiết cho nhiều tác vụ.

4.  **Cấu trúc dự án cơ bản**:
    *   Phân tách `models` (định nghĩa dữ liệu), `services` (logic nghiệp vụ), và `main` (điểm khởi chạy ứng dụng) để mã nguồn dễ đọc, dễ bảo trì hơn.

## Cấu trúc thư mục

```
pointers/
├── go.mod
├── main.go
├── README.md
├── models/
│   └── product.go
└── inventory/
    └── service.go
```

-   **`main.go`**: Điểm khởi đầu của ứng dụng. Nơi khởi tạo dữ liệu và gọi các hàm logic để minh họa.
-   **`models/product.go`**: Định nghĩa cấu trúc dữ liệu `Product`.
-   **`inventory/service.go`**: Chứa các hàm logic nghiệp vụ (services) để thao tác trên dữ liệu `Product`, ví dụ như thêm/bớt hàng tồn kho.

## Yêu cầu

-   Go (phiên bản 1.18 trở lên).

## Cách chạy dự án

1.  Mở terminal và di chuyển đến thư mục gốc của dự án (`go_pointer_project`).
2.  Chạy lệnh sau:
    ```bash
    go run .
    ```

## Luồng hoạt động và giải thích code

Khi chạy, chương trình sẽ thực hiện các bước sau:

1.  **Khởi tạo sản phẩm**: Một biến `laptop` kiểu `models.Product` được tạo trong hàm `main`.

2.  **Minh họa Pass-by-value**:
    -   Hàm `inventory.TryToUpdateNameByValue(laptop, ...)` được gọi.
    -   Hàm này nhận một **bản sao** của `laptop`.
    -   Bên trong hàm, tên của sản phẩm được thay đổi, nhưng vì đây là bản sao nên biến `laptop` gốc trong `main` **không bị ảnh hưởng**.
    -   Output sẽ cho thấy tên sản phẩm không thay đổi sau khi gọi hàm.

3.  **Minh họa dùng con trỏ**:
    -   Hàm `inventory.AddStock(&laptop, ...)` được gọi. Lưu ý toán tử `&` để lấy địa chỉ của `laptop`.
    -   Hàm này nhận vào một **con trỏ** kiểu `*models.Product`.
    -   Bên trong hàm, thông qua con trỏ, nó thay đổi trực tiếp trường `Quantity` của biến `laptop` gốc.
    -   Output sẽ cho thấy số lượng tồn kho đã được cập nhật thành công.

### Điểm nhấn trong Code

-   **`inventory/service.go`**:
    -   `func AddStock(p *models.Product, ...)`: Tham số `p` là một con trỏ. Bất kỳ thay đổi nào trên `p.Quantity` đều ảnh hưởng đến biến gốc.
    -   `func TryToUpdateNameByValue(p models.Product, ...)`: Tham số `p` là một giá trị (bản sao). Thay đổi trên `p.Name` chỉ có tác dụng cục bộ bên trong hàm.

## Kết luận

Dự án này là một bài thực hành nhỏ gọn nhưng hiệu quả để hiểu rõ một trong những khái niệm quan trọng và mạnh mẽ nhất của Go. Nắm vững con trỏ là chìa khóa để viết code hiệu quả, tối ưu và có khả năng sửa đổi trạng thái của chương trình một cách tường minh.