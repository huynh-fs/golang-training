package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"     
	"os"
	"strconv"
	"strings"
)

// map toàn cục để lưu trữ dữ liệu học sinh.
var studentGrades = make(map[string][]int)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Hệ Thống Quản Lý Điểm Học Sinh ---")
		fmt.Println("1. Thêm học sinh và điểm (thủ công)")
		fmt.Println("2. Tải dữ liệu học sinh từ CSV")
		fmt.Println("3. Hiển thị tất cả học sinh và điểm trung bình")
		fmt.Println("4. Tìm kiếm học sinh")
		fmt.Println("5. Phân loại học sinh (Đạt/Không Đạt)")
		fmt.Println("6. Thoát")
		fmt.Print("Chọn chức năng: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			addStudentGradesManual(reader) 
		case "2": 
			loadStudentsFromCSV("./students.csv")
		case "3": 
			displayAllStudents()
		case "4": 
			searchStudent(reader)
		case "5": 
			classifyStudents()
		case "6": 
			fmt.Println("Cảm ơn bạn đã sử dụng hệ thống. Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng thử lại.")
		}
	}
}

// Chức năng thêm học sinh và điểm thủ công
func addStudentGradesManual(reader *bufio.Reader) {
	fmt.Print("Nhập tên học sinh: ")
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)

	if name == "" {
		fmt.Println("Tên học sinh không được để trống.")
		return
	}

	fmt.Printf("Nhập số lượng điểm cho %s: ", name)
	numGradesInput, _ := reader.ReadString('\n')
	numGradesStr := strings.TrimSpace(numGradesInput)
	numGrades, err := strconv.Atoi(numGradesStr)
	if err != nil || numGrades <= 0 {
		fmt.Println("Số lượng điểm không hợp lệ. Vui lòng nhập một số nguyên dương.")
		return
	}

	var grades []int
	for i := 0; i < numGrades; i++ {
		fmt.Printf("Nhập điểm thứ %d (0-10): ", i+1)
		gradeInput, _ := reader.ReadString('\n')
		gradeStr := strings.TrimSpace(gradeInput)
		grade, err := strconv.Atoi(gradeStr)

		if err != nil || grade < 0 || grade > 10 {
			fmt.Println("Điểm không hợp lệ. Vui lòng nhập số nguyên từ 0 đến 10.")
			i--
			continue
		}
		grades = append(grades, grade)
	}
	studentGrades[name] = grades
	fmt.Printf("Đã thêm học sinh %s với %d điểm.\n", name, len(grades))
}

// Chức năng mới để đọc dữ liệu từ tệp CSV
func loadStudentsFromCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Lỗi khi mở tệp CSV '%s': %v\n", filePath, err)
		return
	}
	defer file.Close() 

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1 // Cho phép số lượng trường khác nhau trên mỗi dòng (vì số điểm có thể khác nhau)

	fmt.Printf("Đang tải dữ liệu từ '%s'...\n", filePath)
	recordsLoaded := 0

	// Vòng lặp while-style (for {}) để đọc từng dòng trong tệp CSV
	for {
		record, err := csvReader.Read() // Đọc một dòng (record) từ CSV
		if err == io.EOF {
			break // Sử dụng 'break': Thoát khỏi vòng lặp khi đọc đến cuối tệp
		}
		if err != nil {
			fmt.Printf("Lỗi khi đọc dòng từ CSV: %v\n", err)
			// Sử dụng 'continue': Bỏ qua dòng bị lỗi và tiếp tục đọc dòng tiếp theo
			continue
		}

		if len(record) < 2 { // Cần ít nhất tên và 1 điểm
			fmt.Printf("Bỏ qua dòng không hợp lệ (ít hơn 2 trường): %v\n", record)
			continue
		}

		studentName := record[0]
		var grades []int

		// Vòng lặp 'for-range' trên slice `record` (từ record[1] trở đi) để lấy điểm
		for _, gradeStr := range record[1:] {
			grade, err := strconv.Atoi(gradeStr)
			if err != nil {
				fmt.Printf("Lỗi: Điểm '%s' của học sinh '%s' không phải là số nguyên. Bỏ qua điểm này.\n", gradeStr, studentName)
				// Sử dụng 'continue': Bỏ qua điểm không hợp lệ và tiếp tục với điểm tiếp theo trong cùng dòng
				continue
			}
			if grade < 0 || grade > 10 { // Giả sử điểm từ 0-10
				fmt.Printf("Lỗi: Điểm '%d' của học sinh '%s' nằm ngoài phạm vi (0-10). Bỏ qua điểm này.\n", grade, studentName)
				continue
			}
			grades = append(grades, grade)
		}

		if len(grades) > 0 { // Chỉ thêm học sinh nếu có ít nhất một điểm hợp lệ
			studentGrades[studentName] = grades
			recordsLoaded++
		} else {
			fmt.Printf("Học sinh '%s' không có điểm hợp lệ nào. Bỏ qua.\n", studentName)
		}
	}
	fmt.Printf("Đã tải thành công %d học sinh từ '%s'.\n", recordsLoaded, filePath)
}

// Tính điểm trung bình cho một slice điểm
func calculateAverage(grades []int) float64 {
	if len(grades) == 0 {
		return 0.0
	}
	sum := 0
	// Vòng lặp 'for-range' trên slice để tính tổng điểm
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}

// Hiển thị tất cả học sinh và điểm trung bình của họ
func displayAllStudents() {
	if len(studentGrades) == 0 {
		fmt.Println("Chưa có học sinh nào trong hệ thống.")
		return
	}
	fmt.Println("\n--- Danh Sách Học Sinh ---")
	// Vòng lặp 'for-range' trên map để duyệt qua từng học sinh
	for name, grades := range studentGrades {
		avg := calculateAverage(grades)
		fmt.Printf("Học sinh: %s, Điểm trung bình: %.2f\n", name, avg)
	}
}

// Tìm kiếm học sinh theo tên
func searchStudent(reader *bufio.Reader) {
	fmt.Print("Nhập tên học sinh cần tìm: ")
	nameInput, _ := reader.ReadString('\n')
	searchName := strings.TrimSpace(nameInput)

	found := false
	// Vòng lặp 'for-range' trên map để tìm kiếm học sinh
	for name, grades := range studentGrades {
		if strings.EqualFold(name, searchName) { // So sánh không phân biệt chữ hoa/thường
			avg := calculateAverage(grades)
			fmt.Printf("Tìm thấy học sinh: %s, Điểm trung bình: %.2f\n", name, avg)
			found = true
			// Sử dụng 'break': Thoát khỏi vòng lặp ngay lập tức sau khi tìm thấy học sinh
			break
		}
	}

	if !found {
		fmt.Printf("Không tìm thấy học sinh '%s'.\n", searchName)
	}
}

// Phân loại học sinh Đạt/Không Đạt
func classifyStudents() {
	if len(studentGrades) == 0 {
		fmt.Println("Chưa có học sinh nào trong hệ thống để phân loại.")
		return
	}

	fmt.Println("\n--- Học Sinh Đạt ---")
	foundPassing := false
	// Vòng lặp 'for-range' trên map để duyệt và phân loại
	for name, grades := range studentGrades {
		avg := calculateAverage(grades)
		if avg >= 5.0 { // Điều kiện Đạt
			fmt.Printf("Học sinh: %s, Điểm trung bình: %.2f\n", name, avg)
			foundPassing = true
		}
	}
	if !foundPassing {
		fmt.Println("Không có học sinh nào đạt.")
	}

	fmt.Println("\n--- Học Sinh Không Đạt ---")
	foundFailing := false
	// Vòng lặp 'for-range' trên map để duyệt và phân loại
	for name, grades := range studentGrades {
		avg := calculateAverage(grades)
		if avg < 5.0 { // Điều kiện Không Đạt
			fmt.Printf("Học sinh: %s, Điểm trung bình: %.2f\n", name, avg)
			foundFailing = true
		}
	}
	if !foundFailing {
		fmt.Println("Không có học sinh nào không đạt.")
	}
}