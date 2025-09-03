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

type ThongTinLienLac struct {
	Email       string
	SoDienThoai string
}

type SinhVien struct {
	ID      int
	Ho      string
	Ten     string
	Tuoi    int
	LienLac ThongTinLienLac
}

var danhSachSV []SinhVien
var idTiepTheo int = 1


func main() {
	for {
		hienThiMenu() 
		luaChon := nhapLuaChon()

		switch luaChon {
		case "1":
			themSinhVien()
		case "2":
			themTuCSV("sinhvien.csv")
		case "3":
			hienThiTatCaSinhVien()
		case "4":
			timSinhVien()
		case "5": 
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
		}
		fmt.Println()
	}
}

// hienThiMenu: Cập nhật để thêm lựa chọn mới
func hienThiMenu() {
	fmt.Println("--- HỆ THỐNG QUẢN LÝ SINH VIÊN ---")
	fmt.Println("1. Thêm sinh viên mới (thủ công)")
	fmt.Println("2. Thêm sinh viên từ file CSV") 
	fmt.Println("3. Hiển thị danh sách sinh viên")
	fmt.Println("4. Tìm sinh viên theo ID")
	fmt.Println("5. Thoát")                     
	fmt.Print("Nhập lựa chọn của bạn: ")
}

// nhapLuaChon: Hàm nhập lựa chọn của người dùng
func nhapLuaChon() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// themTuCSV: Hàm thêm sinh viên từ file CSV
func themTuCSV(tenFile string) {
	fmt.Printf("Đang đọc dữ liệu từ file %s...\n", tenFile)

	file, err := os.Open(tenFile)
	if err != nil {
		fmt.Printf("Lỗi: Không thể mở file '%s'. Vui lòng kiểm tra lại file có tồn tại không.\n", tenFile)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		fmt.Printf("Lỗi: Không thể đọc dòng tiêu đề từ file '%s'.\n", tenFile)
		return
	}

	soLuongThem := 0
	dongSo := 1 
	for {
		dongSo++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Lỗi khi đọc file CSV ở dòng %d: %v\n", dongSo, err)
			continue 
		}

		if len(record) != 5 {
			fmt.Printf("Cảnh báo: Dòng %d có định dạng sai (cần 5 cột), bỏ qua.\n", dongSo)
			continue
		}

		tuoi, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			fmt.Printf("Cảnh báo: Dòng %d có tuổi không hợp lệ ('%s'), bỏ qua.\n", dongSo, record[2])
			continue
		}

		sv := SinhVien{
			ID:   idTiepTheo,
			Ho:   strings.TrimSpace(record[0]),
			Ten:  strings.TrimSpace(record[1]),
			Tuoi: tuoi,
			LienLac: ThongTinLienLac{
				Email:       strings.TrimSpace(record[3]),
				SoDienThoai: strings.TrimSpace(record[4]),
			},
		}

		danhSachSV = append(danhSachSV, sv)
		idTiepTheo++
		soLuongThem++
	}

	if soLuongThem > 0 {
		fmt.Printf("=> Đã thêm thành công %d sinh viên từ file CSV!\n", soLuongThem)
	} else {
		fmt.Println("Không có sinh viên nào được thêm. Vui lòng kiểm tra lại nội dung file CSV.")
	}
}

// themSinhVien: Hàm thêm sinh viên mới
func themSinhVien() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập họ sinh viên: ")
	ho, _ := reader.ReadString('\n')

	fmt.Print("Nhập tên sinh viên: ")
	ten, _ := reader.ReadString('\n')

	fmt.Print("Nhập tuổi sinh viên: ")
	tuoiStr, _ := reader.ReadString('\n')
	tuoi, err := strconv.Atoi(strings.TrimSpace(tuoiStr))
	if err != nil {
		fmt.Println("Tuổi không hợp lệ. Vui lòng nhập số.")
		return
	}

	fmt.Print("Nhập email: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Nhập số điện thoại: ")
	sdt, _ := reader.ReadString('\n')

	sv := SinhVien{
		ID:   idTiepTheo,
		Ho:   strings.TrimSpace(ho),
		Ten:  strings.TrimSpace(ten),
		Tuoi: tuoi,
		LienLac: ThongTinLienLac{
			Email:       strings.TrimSpace(email),
			SoDienThoai: strings.TrimSpace(sdt),
		},
	}

	danhSachSV = append(danhSachSV, sv)
	idTiepTheo++

	fmt.Println("=> Đã thêm sinh viên thành công!")
}

// hienThiTatCaSinhVien: Hàm hiển thị tất cả sinh viên
func hienThiTatCaSinhVien() {
	fmt.Println("\n--- DANH SÁCH TOÀN BỘ SINH VIÊN ---")
	if len(danhSachSV) == 0 {
		fmt.Println("Chưa có sinh viên nào trong danh sách.")
		return
	}

	for _, sv := range danhSachSV {
		fmt.Printf("ID: %d\n", sv.ID)
		fmt.Printf("  Họ và tên: %s %s\n", sv.Ho, sv.Ten)
		fmt.Printf("  Tuổi: %d\n", sv.Tuoi)
		fmt.Printf("  Email: %s\n", sv.LienLac.Email)
		fmt.Printf("  Số điện thoại: %s\n", sv.LienLac.SoDienThoai)
		fmt.Println("--------------------")
	}
}

// timSinhVien: Hàm tìm kiếm sinh viên theo ID
func timSinhVien() {
	fmt.Print("Nhập ID sinh viên cần tìm: ")
	idStr := nhapLuaChon()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID không hợp lệ.")
		return
	}

	timThay := false
	for _, sv := range danhSachSV {
		if sv.ID == id {
			fmt.Println("\n--- THÔNG TIN SINH VIÊN TÌM THẤY ---")
			fmt.Printf("ID: %d\n", sv.ID)
			fmt.Printf("  Họ và tên: %s %s\n", sv.Ho, sv.Ten)
			fmt.Printf("  Tuổi: %d\n", sv.Tuoi)
			fmt.Printf("  Email: %s\n", sv.LienLac.Email)
			fmt.Printf("  Số điện thoại: %s\n", sv.LienLac.SoDienThoai)
			fmt.Println("--------------------")
			timThay = true
			break
		}
	}

	if !timThay {
		fmt.Printf("Không tìm thấy sinh viên có ID = %d.\n", id)
	}
}