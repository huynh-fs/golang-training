package order

import "time"

const (
	StatusPending    = "Chờ xử lý"
	StatusProcessing = "Đang xử lý"
	StatusShipped    = "Đã vận chuyển"
	StatusDelivered  = "Đã giao hàng"
	StatusCancelled  = "Đã hủy"

	PaymentStatusPaid   = "Đã thanh toán"
	PaymentStatusUnpaid = "Chưa thanh toán"
)

type Order struct {
	ID            string
	Status        string
	Items         []string
	Total         float64
	PaymentStatus string
	CreatedDate   time.Time
	ShippedDate   time.Time
}


