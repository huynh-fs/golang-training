package cli

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/huynh-fs/sql/internal/service"
)

type CLI struct {
	ProductService *service.ProductService
	OrderService   *service.OrderService
	Scanner        *bufio.Scanner
	Context        context.Context
}

func NewCLI(productSvc *service.ProductService, orderSvc *service.OrderService) *CLI {
	return &CLI{
		ProductService: productSvc,
		OrderService:   orderSvc,
		Scanner:        bufio.NewScanner(os.Stdin),
		Context:        context.Background(),
	}
}

func (c *CLI) Run() {
	for {
		c.printMenu()

		c.Scanner.Scan()
		choiceStr := c.Scanner.Text()
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			log.Println("Lựa chọn không hợp lệ, vui lòng nhập một số.")
			continue
		}

		if !c.handleChoice(choice) {
			break
		}
	}
}

func (c *CLI) handleChoice(choice int) bool {
	switch choice {
	case 1:
		c.handleListProducts()
	case 2:
		c.handleGetProduct()
	case 3:
		c.handleAddProduct()
	case 4:
		c.handleUpdateQuantity()
	case 5:
		c.handleDeleteProduct()
	case 6:
		c.handlePlaceOrder()
	case 7:
		c.handleGetOrder()
	case 8:
		c.handleListOrders()
	case 0:
		fmt.Println("Tạm biệt!")
		return false // Tín hiệu thoát
	default:
		log.Println("Lựa chọn không hợp lệ, vui lòng chọn lại.")
	}
	return true // Tiếp tục chạy
}

func (c *CLI) printMenu() {
	fmt.Println("\n--- MENU QUẢN LÝ SẢN PHẨM ---")
	fmt.Println("1. Liệt kê tất cả sản phẩm")
	fmt.Println("2. Tìm sản phẩm theo ID")
	fmt.Println("3. Thêm sản phẩm mới")
	fmt.Println("4. Cập nhật số lượng sản phẩm")
	fmt.Println("5. Xóa sản phẩm")
	fmt.Println("6. Đặt hàng (Transaction)")
	fmt.Println("7. Lấy thông tin đơn hàng theo ID")
	fmt.Println("8. Lấy danh sách tất cả đơn hàng")
	fmt.Println("0. Thoát")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func (c *CLI) promptAndRead(prompt string) string {
	fmt.Print(prompt)
	c.Scanner.Scan()
	return strings.TrimSpace(c.Scanner.Text())
}

func (c *CLI) handleListProducts() {
	products, err := c.ProductService.GetAllProducts(c.Context)
	if err != nil {
		log.Printf("Lỗi: %v\n", err)
		return
	}
	if len(products) == 0 {
		fmt.Println("Không có sản phẩm nào.")
		return
	}
	fmt.Println("\n--- DANH SÁCH SẢN PHẨM ---")
	for _, p := range products {
		fmt.Printf("ID: %d, Tên: %s, Giá: %.2f, Số lượng: %d, Tạo lúc: %s\n", p.ID, p.Name, p.Price, p.Quantity, p.CreatedAt.Format("2006-01-02"))
	}
}

func (c *CLI) handleGetProduct() {
	idStr := c.promptAndRead("Nhập ID sản phẩm: ")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("ID không hợp lệ.")
		return
	}

	p, err := c.ProductService.GetProductByID(c.Context, id)
	if err != nil {
		log.Printf("Lỗi: %v\n", err)
		return
	}
	fmt.Printf("Sản phẩm tìm thấy - ID: %d, Tên: %s, Giá: %.2f, Số lượng: %d\n", p.ID, p.Name, p.Price, p.Quantity)
}

func (c *CLI) handleAddProduct() {
	name := c.promptAndRead("Nhập tên sản phẩm: ")
	priceStr := c.promptAndRead("Nhập giá sản phẩm: ")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price <= 0 {
		log.Println("Giá không hợp lệ.")
		return
	}
	quantityStr := c.promptAndRead("Nhập số lượng sản phẩm: ")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		log.Println("Số lượng không hợp lệ.")
		return
	}

	p, err := c.ProductService.CreateProduct(c.Context, name, price, quantity)
	if err != nil {
		log.Printf("Lỗi khi tạo sản phẩm: %v\n", err)
		return
	}
	fmt.Printf("Đã tạo sản phẩm mới thành công với ID: %d\n", p.ID)
}

func (c *CLI) handleUpdateQuantity() {
	idStr := c.promptAndRead("Nhập ID sản phẩm cần cập nhật số lượng: ")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("ID không hợp lệ.")
		return
	}

	quantityStr := c.promptAndRead("Nhập số lượng cần cập nhật: ")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		log.Println("Số lượng không hợp lệ.")
		return
	}

	if err := c.ProductService.UpdateQuantity(c.Context, id, quantity); err != nil {
		log.Printf("Lỗi khi cập nhật số lượng: %v\n", err)
		return
	}
	fmt.Printf("Đã cập nhật số lượng cho sản phẩm ID %d thành công!\n", id)
}

func (c *CLI) handleDeleteProduct() {
	idStr := c.promptAndRead("Nhập ID sản phẩm cần xóa: ")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("ID không hợp lệ.")
		return
	}

	err = c.ProductService.DeleteProduct(c.Context, id)
	if err != nil {
		log.Printf("Lỗi khi xóa sản phẩm: %v\n", err)
		return
	}
	fmt.Printf("Đã xóa sản phẩm với ID %d thành công!\n", id)
}

func (c *CLI) handlePlaceOrder() {
	idStr := c.promptAndRead("Nhập ID sản phẩm cần đặt hàng: ")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("ID không hợp lệ.")
		return
	}
	quantityStr := c.promptAndRead("Nhập số lượng cần đặt: ")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity <= 0 {
		log.Println("Số lượng không hợp lệ.")
		return
	}
	order, err := c.OrderService.CreateOrder(c.Context, id, quantity)
	if err != nil {
		log.Printf("Lỗi khi đặt hàng: %v\n", err)
		return
	}
	fmt.Printf("Đặt hàng thành công! Mã đơn hàng: %d, Sản phẩm ID: %d, Số lượng: %d, Ngày đặt: %s\n", order.ID, order.ProductID, order.QuantityOrdered, order.CreatedAt.Format("2006-01-02"))
}

func (c *CLI) handleGetOrder() {
	idStr := c.promptAndRead("Nhập ID đơn hàng: ")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("ID không hợp lệ.")
		return
	}
	order, err := c.OrderService.GetOrderByID(c.Context, id)
	if err != nil {
		log.Printf("Lỗi khi lấy đơn hàng: %v\n", err)
		return
	}
	fmt.Printf("Đơn hàng tìm thấy - Mã đơn hàng: %d, Sản phẩm ID: %d, Số lượng: %d, Trạng thái: %s, Ngày đặt: %s\n", order.ID, order.ProductID, order.QuantityOrdered, order.OrderStatus, order.CreatedAt.Format("2006-01-02"))
}

func (c *CLI) handleListOrders() {
	orders, err := c.OrderService.GetAllOrders(c.Context)
	if err != nil {
		log.Printf("Lỗi: %v\n", err)
		return
	}
	if len(orders) == 0 {
		fmt.Println("Không có đơn hàng nào.")
		return
	}
	fmt.Println("\n--- DANH SÁCH ĐƠN HÀNG ---")
	for _, o := range orders {
		fmt.Printf("Mã đơn hàng: %d, Sản phẩm ID: %d, Số lượng: %d, Trạng thái: %s, Ngày đặt: %s\n", o.ID, o.ProductID, o.QuantityOrdered, o.OrderStatus, o.CreatedAt.Format("2006-01-02"))
	}
}