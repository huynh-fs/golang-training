# Go Zoo Management Demo

Đây là một dự án demo đơn giản được viết bằng Go, dùng để minh họa các khái niệm cốt lõi về thiết kế phần mềm linh hoạt và dễ mở rộng. Dự án mô phỏng một hệ thống quản lý sở thú nhỏ, tập trung vào cách cấu trúc code hơn là các tính năng phức tạp.

Mục tiêu chính của dự án này là để thực hành và hiểu rõ:

- Interfaces trong Go.
- Struct Embedding để tái sử dụng code.
- Kiến trúc phân lớp (Layered Architecture).
- Cấu trúc thư mục chuẩn trong một dự án Go.

## ✨ Các khái niệm cốt lõi được áp dụng

Dự án này được xây dựng xung quanh các khái niệm mạnh mẽ nhất của Go:

1.  **Interfaces để định nghĩa hành vi**: Các `interface` như `Creature`, `Speaker`, `Mover` định nghĩa các "hợp đồng" về hành vi mà không cần quan tâm đến cách hiện thực cụ thể.
2.  **Hiện thực Interface ngầm**: Các `struct` như `Dog`, `Snake` và `Bird` tự động thỏa mãn các `interface` chỉ bằng cách hiện thực các phương thức được yêu cầu, không cần từ khóa `implements`.
3.  **Struct Embedding (Composition over Inheritance)**: `Dog`, `Snake` và `Bird` "thừa hưởng" các thuộc tính từ `struct` cơ sở `animal` thông qua embedding, giúp tái sử dụng code hiệu quả.
4.  **Kiến trúc phân lớp (Layered Architecture)**: Code được tách biệt rõ ràng thành các lớp với trách nhiệm riêng biệt (`domain`, `service`, `handler`), giúp tăng khả năng bảo trì và kiểm thử.
5.  **Cấu trúc dự án chuẩn (Standard Go Project Layout)**: Dự án tuân theo cấu trúc thư mục `cmd`, `internal`, `pkg` phổ biến trong cộng đồng Go.

## 📂 Cấu trúc thư mục

Dự án được tổ chức theo cấu trúc chuẩn để đảm bảo sự rõ ràng và tách biệt các mối quan tâm.

```
zoo-management/
├── cmd/
│   └── zoo/
│       └── main.go         # Entrypoint: Lắp ráp và chạy ứng dụng.
├── internal/
│   ├── domain/             # Lớp nghiệp vụ cốt lõi, định nghĩa các interface trừu tượng.
│   │   └── creature.go
│   └── service/            # Lớp chứa business logic, làm việc với các đối tượng domain.
│       └── zoo_service.go
├── pkg/
│   └── creatures/          # Các hiện thực cụ thể, có thể được chia sẻ cho dự án khác.
│       ├── interfaces.go
│       ├── dog.go
│       ├── snake.go
│       └── bird.go
└── go.mod                  # File quản lý module và các dependency.
```

- **`cmd/`**: Chứa điểm khởi đầu (entrypoint) của ứng dụng.
- **`internal/`**: Chứa logic nghiệp vụ riêng của dự án này và không thể được import bởi các dự án khác.
  - `domain`: Định nghĩa các `interface` và `struct` cốt lõi nhất của nghiệp vụ.
  - `service`: Điều phối và thực thi các business logic.
- **`pkg/`**: Chứa các thư viện và `struct` cụ thể có thể tái sử dụng một cách an toàn ở các dự án khác.

## 🚀 Bắt đầu

### Yêu cầu

- Go phiên bản 1.18 trở lên.

### Chạy ứng dụng

Để chạy ứng dụng, thực thi file `main.go` trong thư mục `cmd/zoo` từ thư mục gốc của dự án:

```sh
go run ./cmd/zoo
```

### Kết quả mong đợi

Bạn sẽ thấy output sau trên terminal:

```
=======Báo cáo sở thú: ======
----Báo cáo về Mực----
Tiếng kêu: Gâu gâu!
Cách di chuyển: Chạy bằng bốn chân

----Báo cáo về Lê----
Tiếng kêu: Xì xì!
Cách di chuyển: Lê trên mặt đất

----Báo cáo về Ổi----
Tiếng kêu: Chíp chíp!
Cách di chuyển: Bay bằng đôi cánh

```
