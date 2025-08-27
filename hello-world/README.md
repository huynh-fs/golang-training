
# Hello World in Go

Đây là chương trình **Hello World** cơ bản được viết bằng ngôn ngữ lập trình **Go (Golang)**.  
Mục đích là để làm quen với cú pháp cơ bản và cách chạy một chương trình Go.

---

## 📌 Cấu trúc chương trình

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
## 🚀 Cách chạy chương trình
### 1. Cài đặt Go
Tải và cài đặt Go tại: https://go.dev/dl/

Kiểm tra cài đặt thành công bằng lệnh:
```bash
go version
```

### 2. Chạy chương trình trực tiếp
```bash
go run main.go
```

3. Biên dịch và chạy file thực thi
```bash
go build main.go
./main       # Linux/MacOS
main.exe     # Windows
```
## 🎯 Kết quả
Khi chạy, chương trình sẽ in ra màn hình:

<img width="623" height="140" alt="image" src="https://github.com/user-attachments/assets/34b8b4ad-280b-4277-9be6-e7793c73aad5" />

## 📖 Tham khảo
- [Tài liệu chính thức Golang](https://go.dev/)
- [Tour of Go](https://go.dev/tour/list) 

