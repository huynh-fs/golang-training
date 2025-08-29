package inventory

import (
	"errors"
	"fmt"
	"github.com/huynh-fs/go_pointer_project/models"
)

// Thêm một số lượng hàng vào sản phẩm.
// Nhận một con trỏ (*models.Product) để có thể thay đổi giá trị Quantity của sản phẩm gốc. 
// Nếu không, nó sẽ chỉ thay đổi trên một bản sao và sản phẩm gốc sẽ không bị ảnh hưởng.
func AddStock(p *models.Product, amount int) {
	if p == nil {
		return
	}
	p.Quantity += amount // Go tự động "dereference" con trỏ khi truy cập trường của struct
	fmt.Printf("   [INFO] Đã thêm %d sản phẩm. Tồn kho mới: %d\n", amount, p.Quantity)
}

// Xóa một số lượng hàng khỏi sản phẩm.
// Tương tự như AddStock, hàm này cũng cần con trỏ để sửa đổi giá trị gốc.
func RemoveStock(p *models.Product, amount int) error {
	if p == nil {
		return errors.New("sản phẩm không tồn tại")
	}
	if p.Quantity < amount {
		return fmt.Errorf("không đủ hàng tồn kho. Chỉ còn %d, cần %d", p.Quantity, amount)
	}
	p.Quantity -= amount
	fmt.Printf("   [INFO] Đã bớt %d sản phẩm. Tồn kho mới: %d\n", amount, p.Quantity)
	return nil
}

// Ví dụ về truyền tham trị (pass-by-value).
// Hàm này nhận một bản sao của Product. 
// Mọi thay đổi trên `p` trong hàm này sẽ không ảnh hưởng đến biến Product gốc được truyền vào.
func TryToUpdateNameByValue(p models.Product, newName string) {
	p.Name = newName
	fmt.Printf("   [INFO] Tên sản phẩm bên trong hàm đã đổi thành: '%s'\n", p.Name)
}

// Hiển thị thông tin sản phẩm.
// Hàm này không thay đổi dữ liệu nên không cần sử dụng con trỏ.
func DisplayProduct(p models.Product) {
	fmt.Printf("--- Thông tin sản phẩm ---\nID: %s\nTên: %s\nGiá: %.2f\nTồn kho: %d\n-------------------------\n",
		p.ID, p.Name, p.Price, p.Quantity)
}