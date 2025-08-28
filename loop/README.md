# Hệ Thống Quản Lý Điểm Học Sinh Đơn Giản

Ứng dụng console bằng Go để quản lý điểm của một nhóm học sinh, hỗ trợ nhập liệu thủ công và tải dữ liệu từ tệp CSV. Dự án này được thiết kế để minh họa việc sử dụng các cấu trúc điều khiển vòng lặp (`for` classic, `while-style for`, `for-range` trên slice và map) cùng với các câu lệnh `break` và `continue` trong Go.

## Mục lục

*   [Tính năng](#tính-năng)
*   [Cấu trúc dữ liệu](#cấu-trúc-dữ-liệu)
*   [Cách chạy ứng dụng](#cách-chạy-ứng-dụng)
*   [Định dạng tệp CSV](#định-dạng-tệp-csv)
*   [Sử dụng ứng dụng](#sử-dụng-ứng-dụng)
*   [Các khái niệm Go đã áp dụng](#các-khái-niệm-go-đã-áp-dụng)

## Tính năng

*   **Thêm học sinh và điểm (thủ công):** Nhập tên học sinh và danh sách các điểm số trực tiếp từ console.
*   **Tải dữ liệu học sinh từ CSV:** Đọc thông tin học sinh và điểm từ một tệp CSV.
*   **Hiển thị tất cả học sinh và điểm trung bình:** Liệt kê tên của tất cả học sinh cùng với điểm trung bình của họ.
*   **Tìm kiếm học sinh:** Tìm kiếm học sinh theo tên và hiển thị điểm trung bình của họ.
*   **Phân loại học sinh:** Hiển thị danh sách các học sinh "Đạt" (điểm trung bình >= 5.0) và "Không Đạt" (điểm trung bình < 5.0).
*   **Thoát chương trình:** Kết thúc ứng dụng.

## Cấu trúc dữ liệu

Dữ liệu học sinh được lưu trữ trong một `map` toàn cục:
*   **Key:** Tên học sinh (kiểu `string`).
*   **Value:** Một `slice` các điểm số (kiểu `[]int`).

```go
var studentGrades = make(map[string][]int)
```

## Cách chạy ứng dụng

1.  **Yêu cầu:** Đảm bảo bạn đã cài đặt Go (phiên bản 1.16 trở lên được khuyến nghị).
2.  **Lưu mã nguồn:** Lưu đoạn mã Go vào một tệp có tên `main.go`.
3.  **Tạo tệp CSV (tùy chọn):** Để sử dụng chức năng tải từ CSV, tạo một tệp `students.csv` trong cùng thư mục với `main.go`. Xem [Định dạng tệp CSV](#định-dạng-tệp-csv) bên dưới để biết cấu trúc.
4.  **Mở Terminal/Command Prompt:** Điều hướng đến thư mục chứa tệp `main.go`.
5.  **Chạy ứng dụng:**
    ```bash
    go run main.go
    ```

## Định dạng tệp CSV

Tệp CSV dự kiến có định dạng như sau: mỗi dòng đại diện cho một học sinh, với tên học sinh ở cột đầu tiên, theo sau là các điểm số được phân tách bằng dấu phẩy.

**Ví dụ `students.csv`:**
```csv
Alice,8,7,9
Bob,4,6,5
Charlie,9,9,10,8
David,3,2
Eve,7,8,7,9,6
InvalidStudent,abc,7,8
AnotherStudent,10,9,invalid_grade,8
```
*Lưu ý: Ứng dụng sẽ xử lý các điểm không hợp lệ (không phải số hoặc ngoài phạm vi 0-10) bằng cách bỏ qua chúng.*

## Sử dụng ứng dụng

Sau khi chạy `go run main.go`, bạn sẽ thấy một menu tương tác:

```
--- Hệ Thống Quản Lý Điểm Học Sinh ---
1. Thêm học sinh và điểm (thủ công)
2. Tải dữ liệu học sinh từ CSV
3. Hiển thị tất cả học sinh và điểm trung bình
4. Tìm kiếm học sinh
5. Phân loại học sinh (Đạt/Không Đạt)
6. Thoát
Chọn chức năng:
```

*   **Chọn `1`** để nhập tên và điểm của học sinh thủ công.
*   **Chọn `2`** để tải dữ liệu từ tệp CSV. Bạn cần phải có file `students.csv` trong thư mục.
*   **Chọn `3`** để xem danh sách tất cả học sinh và điểm trung bình của họ.
*   **Chọn `4`** để tìm kiếm một học sinh cụ thể theo tên.
*   **Chọn `5`** để xem danh sách học sinh được phân loại là "Đạt" hoặc "Không Đạt".
*   **Chọn `6`** để thoát khỏi chương trình.

## Các khái niệm Go đã áp dụng

Dự án này là một ví dụ thực tế về việc sử dụng các cấu trúc điều khiển cơ bản trong Go:

*   **Vòng lặp `for` (While-style):** Vòng lặp chính của ứng dụng (`for {}` trong `main`) hoạt động như một vòng lặp `while(true)`, liên tục hiển thị menu cho đến khi người dùng chọn thoát.
*   **Vòng lặp `for` (Classic):** Được sử dụng trong hàm `addStudentGradesManual` để lặp một số lần cố định khi yêu cầu người dùng nhập từng điểm số.
*   **Vòng lặp `for-range` trên `slice`:**
    *   Trong hàm `calculateAverage` để tính tổng các điểm số.
    *   Trong hàm `loadStudentsFromCSV` để duyệt qua các chuỗi điểm trong một dòng CSV.
*   **Vòng lặp `for-range` trên `map`:**
    *   Trong các hàm `displayAllStudents`, `searchStudent`, `classifyStudents` để duyệt qua tất cả các học sinh đã lưu trữ.
*   **`break`:**
    *   Trong hàm `main`, `return` được sử dụng để thoát khỏi hàm `main` (và vòng lặp chính) khi người dùng chọn thoát.
    *   Trong hàm `searchStudent`, `break` được sử dụng để dừng tìm kiếm ngay lập tức sau khi tìm thấy học sinh.
    *   Trong hàm `loadStudentsFromCSV`, `break` được sử dụng để thoát khỏi vòng lặp đọc CSV khi đạt đến `io.EOF` (cuối tệp).
*   **`continue`:**
    *   Trong hàm `addStudentGradesManual`, `continue` được sử dụng để yêu cầu người dùng nhập lại điểm nếu điểm không hợp lệ, mà không chuyển sang điểm tiếp theo.
    *   Trong hàm `loadStudentsFromCSV`, `continue` được sử dụng để bỏ qua các dòng hoặc điểm không hợp lệ trong tệp CSV và tiếp tục xử lý các dữ liệu khác.
*   **`map` và `slice`:** Các kiểu dữ liệu cơ bản của Go được sử dụng để lưu trữ và quản lý dữ liệu học sinh.
*   **Gói `encoding/csv`:** Để đọc và phân tích cú pháp dữ liệu từ tệp CSV một cách hiệu quả.
*   **Xử lý lỗi:** Ứng dụng bao gồm các kiểm tra lỗi cơ bản khi đọc input và xử lý tệp.
