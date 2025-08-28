package order

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/huynh-fs/order-manager-project/internal/utils"
)

// OrderManager để quản lý bộ sưu tập các đơn hàng.
type OrderManager struct {
	orders      map[string]*Order
	nextOrderID int
}

// NewOrderManager tạo và trả về một instance mới của OrderManager.
func NewOrderManager() *OrderManager {
	return &OrderManager{
		orders:      make(map[string]*Order),
		nextOrderID: 1,
	}
}

// GetOrder lấy đơn hàng theo ID.
func (om *OrderManager) GetOrder(id string) (*Order, bool) {
	order, found := om.orders[id]
	return order, found
}

// normalizeStatusInput chuyển đổi đầu vào trạng thái của người dùng thành hằng số trạng thái chuẩn.
// Trả về string rỗng nếu đầu vào không hợp lệ.
func normalizeStatusInput(input string) string {
	lowerInput := strings.ToLower(input)
	switch lowerInput {
	case strings.ToLower(StatusPending), "pending": // Chấp nhận "Chờ xử lý" hoặc "pending"
		return StatusPending
	case strings.ToLower(StatusProcessing), "processing": // Chấp nhận "Đang xử lý" hoặc "processing"
		return StatusProcessing
	case strings.ToLower(StatusShipped), "shipped": // Chấp nhận "Đã vận chuyển" hoặc "shipped"
		return StatusShipped
	case strings.ToLower(StatusDelivered), "delivered": // Chấp nhận "Đã giao hàng" hoặc "delivered"
		return StatusDelivered
	case strings.ToLower(StatusCancelled), "cancelled": // Chấp nhận "Đã hủy" hoặc "cancelled"
		return StatusCancelled
	default:
		return "" // Trả về rỗng nếu không khớp
	}
}


// CreateOrder tạo một đơn hàng mới và thêm vào hệ thống.
func (om *OrderManager) CreateOrder() {
	orderID := fmt.Sprintf("ORD%03d", om.nextOrderID)
	om.nextOrderID++

	itemsInput := utils.ReadInput("Nhập các mặt hàng (phân cách bởi dấu phẩy, ví dụ: Laptop, Mouse): ")
	items := strings.Split(itemsInput, ",")
	for i := range items {
		items[i] = strings.TrimSpace(items[i])
	}

	totalStr := utils.ReadInput("Nhập tổng giá trị đơn hàng (ví dụ: 1200.00): ")
	total, err := strconv.ParseFloat(totalStr, 64)

	if err != nil || total <= 0 {
		fmt.Println("Lỗi: Tổng giá trị không hợp lệ. Đơn hàng không được tạo.")
		return
	}

	paymentStatus := utils.ReadInput("Trạng thái thanh toán (Paid/Unpaid): ")
	if strings.ToLower(paymentStatus) != "paid" && strings.ToLower(paymentStatus) != "unpaid" {
		fmt.Println("Trạng thái thanh toán không hợp lệ, đặt mặc định là 'Unpaid'.")
		paymentStatus = PaymentStatusUnpaid
	} else {
		paymentStatus = strings.ToLower(paymentStatus)
	}

	newOrder := &Order{
		ID:            orderID,
		Status:        StatusPending,
		Items:         items,
		Total:         total,
		PaymentStatus: paymentStatus,
		CreatedDate:   time.Now(),
	}
	om.orders[orderID] = newOrder
	fmt.Printf("Đã tạo đơn hàng mới: %s\n", orderID)
}

// ViewOrder hiển thị chi tiết của một đơn hàng cụ thể.
func (om *OrderManager) ViewOrder() {
	orderID := utils.ReadInput("Nhập ID đơn hàng cần xem (ví dụ: ORD001): ")
	order, found := om.orders[orderID]

	if found {
		fmt.Printf("\n--- Chi tiết đơn hàng %s ---\n", order.ID)
		fmt.Printf("Trạng thái: %s\n", order.Status)
		fmt.Printf("Mặt hàng: %s\n", strings.Join(order.Items, ", "))
		fmt.Printf("Tổng cộng: %.2f\n", order.Total)
		fmt.Printf("Thanh toán: %s\n", order.PaymentStatus)
		fmt.Printf("Ngày tạo: %s\n", order.CreatedDate.Format("02-01-2006 15:04:05"))
		if order.Status == StatusShipped || order.Status == StatusDelivered {
			fmt.Printf("Ngày vận chuyển: %s\n", order.ShippedDate.Format("02-01-2006 15:04:05"))
		}
		fmt.Println("------------------------------")
	} else {
		fmt.Println("Lỗi: Không tìm thấy đơn hàng với ID:", orderID)
	}
}

// UpdateOrderStatus cập nhật trạng thái của một đơn hàng.
func (om *OrderManager) UpdateOrderStatus() {
	orderID := utils.ReadInput("Nhập ID đơn hàng cần cập nhật (ví dụ: ORD001): ")
	order, found := om.orders[orderID]

	if !found {
		fmt.Println("Lỗi: Không tìm thấy đơn hàng với ID:", orderID)
		return
	}

	fmt.Printf("Trạng thái hiện tại của đơn hàng %s là: %s\n", order.ID, order.Status)

	inputStatus := utils.ReadInput("Nhập trạng thái mới (Chờ xử lý/Pending, Đang xử lý/Processing, Đã vận chuyển/Shipped, Đã giao hàng/Delivered, Đã hủy/Cancelled): ")
	newStatus := normalizeStatusInput(inputStatus) // Chuẩn hóa đầu vào

	// Kiểm tra nếu trạng thái mới không hợp lệ sau khi chuẩn hóa
	if newStatus == "" {
		fmt.Println("Lỗi: Trạng thái mới không hợp lệ. Vui lòng chọn từ các trạng thái được liệt kê.")
		return
	}

	switch order.Status {
	case StatusPending:
		switch newStatus {
		case StatusProcessing:
			order.Status = StatusProcessing
			fmt.Println("Đơn hàng đã được chuyển sang trạng thái 'Đang xử lý'.")
		case StatusCancelled:
			order.Status = StatusCancelled
			fmt.Println("Đơn hàng đã bị 'Hủy'.")
		default:
			fmt.Println("Chỉ có thể chuyển từ 'Chờ xử lý' sang 'Đang xử lý' hoặc 'Đã hủy'.")
		}
	case StatusProcessing:
		switch newStatus {
		case StatusShipped:
			order.Status = StatusShipped
			order.ShippedDate = time.Now()
			fmt.Println("Đơn hàng đã được chuyển sang trạng thái 'Đã vận chuyển'.")
		case StatusCancelled:
			fmt.Println("Cảnh báo: Việc hủy đơn hàng đang xử lý có thể yêu cầu xác nhận thêm.")
			order.Status = StatusCancelled
		default:
			fmt.Println("Chỉ có thể chuyển từ 'Đang xử lý' sang 'Đã vận chuyển' hoặc 'Đã hủy'.")
		}
	case StatusShipped:
		switch newStatus {
		case StatusDelivered:
			order.Status = StatusDelivered
			fmt.Println("Đơn hàng đã được chuyển sang trạng thái 'Đã giao hàng'.")
		default:
			fmt.Println("Chỉ có thể chuyển từ 'Đã vận chuyển' sang 'Đã giao hàng'.")
		}
	case StatusDelivered, StatusCancelled:
		fmt.Println("Không thể cập nhật trạng thái của đơn hàng đã 'Đã giao hàng' hoặc 'Đã hủy'.")
	default:
		fmt.Println("Trạng thái đơn hàng không xác định hoặc không thể cập nhật.")
	}
}

// CheckEligibilityAndSuggestions kiểm tra các điều kiện đặc biệt và đưa ra gợi ý/cảnh báo.
func (om *OrderManager) CheckEligibilityAndSuggestions() {
	orderID := utils.ReadInput("Nhập ID đơn hàng để kiểm tra (ví dụ: ORD001): ")
	order, found := om.orders[orderID]

	if !found {
		fmt.Println("Lỗi: Không tìm thấy đơn hàng với ID:", orderID)
		return
	}

	fmt.Printf("\n--- Kiểm tra và Gợi ý cho đơn hàng %s ---\n", order.ID)

	switch {
	case order.Status == StatusPending && order.PaymentStatus == PaymentStatusUnpaid:
		fmt.Println("[Gợi ý] Đơn hàng đang 'Chờ xử lý' và 'Chưa thanh toán'. Hãy gửi nhắc nhở thanh toán cho khách hàng.")
	case order.Status == StatusPending && order.Total > 500:
		fmt.Println("[Cảnh báo] Đơn hàng 'Chờ xử lý' với giá trị cao (trên 500). Cần kiểm tra kỹ lưỡng trước khi xử lý.")
	case order.Status == StatusProcessing && order.PaymentStatus == PaymentStatusUnpaid:
		fmt.Println("[RỦI RO] Đơn hàng 'Đang xử lý' nhưng 'Chưa thanh toán'! Dừng xử lý và liên hệ khách hàng.")
	case order.Status == StatusShipped:
		if daysSinceShipped := time.Since(order.ShippedDate).Hours() / 24; daysSinceShipped > 3 && daysSinceShipped < 7 {
			fmt.Printf("[Gợi ý] Đơn hàng đã 'Đã vận chuyển' %v ngày. Theo dõi tiến độ giao hàng.\n", int(daysSinceShipped))
		} else if daysSinceShipped >= 7 {
			fmt.Printf("[Cảnh báo] Đơn hàng đã 'Đã vận chuyển' %v ngày nhưng chưa 'Đã giao hàng'. Liên hệ đơn vị vận chuyển ngay lập tức!\n", int(daysSinceShipped))
		}
	case order.Status == StatusDelivered:
		fmt.Println("[Thông tin] Đơn hàng đã 'Đã giao hàng' thành công. Yêu cầu phản hồi từ khách hàng.")
	case order.Status == StatusCancelled:
		fmt.Println("[Thông tin] Đơn hàng đã 'Đã hủy'. Không có hành động nào khác.")
	default:
		fmt.Println("[Thông tin] Không có gợi ý đặc biệt nào cho trạng thái này của đơn hàng.")
	}

	fmt.Println("\n--- Kiểm tra bổ sung (với fallthrough để chạy nhiều case) ---")
	switch order.Status {
	case StatusPending:
		fmt.Println("- Cần xác nhận thông tin đơn hàng ban đầu.")
		fallthrough
	case StatusProcessing:
		fmt.Println("- Đảm bảo có đủ hàng trong kho và bắt đầu đóng gói.")
		fallthrough
	case StatusShipped:
		fmt.Println("- Kiểm tra lại thông tin vận chuyển và mã theo dõi.")
	default:
		fmt.Println("- Đơn hàng ở trạng thái cuối cùng hoặc không liên quan đến chu trình chính.")
	}
	fmt.Println("-------------------------------------------")
}

// AddSampleOrders để thêm một số đơn hàng mẫu vào hệ thống.
func (om *OrderManager) AddSampleOrders() {
	om.orders["ORD001"] = &Order{ID: "ORD001", Status: StatusPending, Items: []string{"Laptop", "Mouse"}, Total: 1200.00, PaymentStatus: PaymentStatusPaid, CreatedDate: time.Now().Add(-48 * time.Hour)}
	om.orders["ORD002"] = &Order{ID: "ORD002", Status: StatusProcessing, Items: []string{"Keyboard"}, Total: 80.00, PaymentStatus: PaymentStatusUnpaid, CreatedDate: time.Now().Add(-24 * time.Hour)}
	om.orders["ORD003"] = &Order{ID: "ORD003", Status: StatusShipped, Items: []string{"Monitor"}, Total: 300.00, PaymentStatus: PaymentStatusPaid, CreatedDate: time.Now().Add(-72 * time.Hour), ShippedDate: time.Now().Add(-72 * time.Hour)}
	om.orders["ORD004"] = &Order{ID: "ORD004", Status: StatusDelivered, Items: []string{"Webcam"}, Total: 50.00, PaymentStatus: PaymentStatusPaid, CreatedDate: time.Now().Add(-120 * time.Hour)}
	om.nextOrderID = 5
}