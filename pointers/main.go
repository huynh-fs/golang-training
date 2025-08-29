package main

import (
	"fmt"
	"github.com/huynh-fs/go_pointer_project/inventory"
	"github.com/huynh-fs/go_pointer_project/models"
)

func main() {
	// 1. Khởi tạo một sản phẩm
	laptop := models.Product{
		ID:       "MBP-01",
		Name:     "Macbook Pro 14",
		Price:    2000.00,
		Quantity: 20,
	}

	fmt.Println(">>> TRẠNG THÁI BAN ĐẦU <<<")
	inventory.DisplayProduct(laptop)

	// 2. Minh họa Pass-by-Value (Truyền tham trị)
	// Gọi hàm TryToUpdateNameByValue.
	// Tên của `laptop` sẽ không thay đổi sau khi hàm này kết thúc.
	fmt.Println("\n>>> THỬ THAY ĐỔI TÊN SẢN PHẨM (PASS-BY-VALUE) <<<")
	inventory.TryToUpdateNameByValue(laptop, "Macbook Pro 16 Siêu Cấp")
	fmt.Printf("Tên sản phẩm bên ngoài hàm vẫn là: '%s'\n", laptop.Name)
	inventory.DisplayProduct(laptop)

	// 3. Minh họa dùng con trỏ để thay đổi giá trị gốc
	// Truyền ĐỊA CHỈ BỘ NHỚ của `laptop` bằng toán tử `&`.
	// Hàm AddStock nhận một con trỏ và có thể thay đổi giá trị tại địa chỉ đó.
	fmt.Println("\n>>> THÊM HÀNG VÀO KHO (DÙNG CON TRỎ) <<<")
	inventory.AddStock(&laptop, 10)
	inventory.DisplayProduct(laptop)

	// 4. Tiếp tục dùng con trỏ để thay đổi giá trị
	fmt.Println("\n>>> BỚT HÀNG KHỎI KHO (DÙNG CON TRỎ) <<<")
	err := inventory.RemoveStock(&laptop, 5)
	if err != nil {
		fmt.Printf("Lỗi: %v\n", err)
	}
	inventory.DisplayProduct(laptop)

	// 5. Minh họa xử lý lỗi khi dùng con trỏ
	fmt.Println("\n>>> THỬ BỚT HÀNG QUÁ SỐ LƯỢNG TỒN KHO <<<")
	err = inventory.RemoveStock(&laptop, 100)
	if err != nil {
		fmt.Printf("Lỗi đã được xử lý: %v\n", err)
	}
	inventory.DisplayProduct(laptop)
}