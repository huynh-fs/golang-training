package main

import (
	"fmt"
	"strings"

	"github.com/huynh-fs/order-manager-project/internal/order"
	"github.com/huynh-fs/order-manager-project/internal/utils"
)

func processGenericInput(input interface{}) {
	fmt.Println("\n--- Ví dụ Type Switch ---")
	switch v := input.(type) {
	case int:
		fmt.Printf("Đầu vào là một số nguyên: %d\n", v)
		if v < 0 {
			fmt.Println("Giá trị âm được phát hiện.")
		}
	case string:
		fmt.Printf("Đầu vào là một chuỗi: \"%s\"\n", v)
		if len(v) > 10 {
			fmt.Println("Chuỗi dài.")
		}
	case bool:
		fmt.Printf("Đầu vào là một boolean: %t\n", v)
	case *order.Order: // Có thể kiểm tra kiểu *order.Order vì main.go đã import gói "order"
		fmt.Printf("Đầu vào là một con trỏ đến Order: ID=%s, Trạng thái=%s\n", v.ID, v.Status)
	default:
		fmt.Printf("Đầu vào có kiểu không xác định: %T (Giá trị: %v)\n", v, v)
	}
	fmt.Println("---------------------------------------")
}

func main() {
	fmt.Println("Chào mừng đến với Hệ thống quản lý đơn hàng E-commerce Go!")

	orderManager := order.NewOrderManager()
	orderManager.AddSampleOrders()

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Tạo đơn hàng mới")
		fmt.Println("2. Xem chi tiết đơn hàng")
		fmt.Println("3. Cập nhật trạng thái đơn hàng")
		fmt.Println("4. Kiểm tra và gợi ý hành động cho đơn hàng")
		fmt.Println("5. Minh họa Type Switch")
		fmt.Println("q. Thoát")
		choice := utils.ReadInput("Nhập lựa chọn của bạn: ")

		switch strings.ToLower(choice) {
		case "1":
			orderManager.CreateOrder()
		case "2":
			orderManager.ViewOrder()
		case "3":
			orderManager.UpdateOrderStatus()
		case "4":
			orderManager.CheckEligibilityAndSuggestions()
		case "5":
			// Minh họa Type Switch với các kiểu dữ liệu khác nhau
			processGenericInput(123)
			processGenericInput("hello world from type switch")
			processGenericInput(true)
			sampleOrder, found := orderManager.GetOrder("ORD001")
			if found {
				processGenericInput(sampleOrder)
			} else {
				fmt.Println("Không tìm thấy đơn hàng mẫu ORD001 cho minh họa Type Switch.")
			}
			processGenericInput(3.14)
		case "q":
			fmt.Println("Đang thoát chương trình. Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng thử lại.")
		}
	}
}