package repository

import (
	"context"
	"time"
	"fmt"
	"github.com/huynh-fs/sql/internal/model"
)

type OrderRepo struct {
	db DBTX
}

func NewOrderRepository(db DBTX) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, o *model.Order) error {
	query := `INSERT INTO orders (product_id, quantity_ordered, order_status, created_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := r.db.QueryRowContext(ctx, query, o.ProductID, o.QuantityOrdered, o.OrderStatus, time.Now()).Scan(&o.ID, &o.CreatedAt)
	if err != nil {
		return fmt.Errorf("không thể tạo đơn hàng: %w", err)
	}
	return nil
}

func (r *OrderRepo) GetOrderByID(ctx context.Context, id int64) (*model.Order, error) {
	query := `SELECT id, product_id, quantity_ordered, order_status, created_at FROM orders WHERE id = $1`
	o := &model.Order{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&o.ID, &o.ProductID, &o.QuantityOrdered, &o.OrderStatus, &o.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy đơn hàng: %w", err)
	}
	return o, nil
}

func (r *OrderRepo) GetAllOrders(ctx context.Context) ([]*model.Order, error) {
	query := `SELECT id, product_id, quantity_ordered, order_status, created_at FROM orders`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy danh sách đơn hàng: %w", err)
	}
	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		o := &model.Order{}
		if err := rows.Scan(&o.ID, &o.ProductID, &o.QuantityOrdered, &o.OrderStatus, &o.CreatedAt); err != nil {
			return nil, fmt.Errorf("không thể quét đơn hàng: %w", err)
		}
		orders = append(orders, o)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("lỗi khi duyệt đơn hàng: %w", err)
	}
	return orders, nil
}
