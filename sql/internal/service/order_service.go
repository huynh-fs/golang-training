package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"github.com/huynh-fs/sql/internal/model"
	"github.com/huynh-fs/sql/internal/repository"
)

var ErrInsufficientStock = errors.New("số lượng sản phẩm không đủ")	

type OrderService struct {
	db *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrder(ctx context.Context, productID int64, quantity int) (*model.Order, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("không thể bắt đầu transaction: %w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			log.Printf("Lỗi khi rollback transaction: %v", err)
		}
	}()

	productRepo := repository.NewProductRepository(tx)
	orderRepo := repository.NewOrderRepository(tx)

	product, err := productRepo.GetProductByID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy sản phẩm: %w", err)
	}
	if product.Quantity < quantity {
		return nil, ErrInsufficientStock
	}
	if err := productRepo.DecreaseInStock(ctx, productID, quantity); err != nil {
		return nil, fmt.Errorf("không thể giảm số lượng sản phẩm: %w", err)
	}

	order := &model.Order{
		ProductID:       productID,
		QuantityOrdered: quantity,
		OrderStatus:    "COMPLETED",
	}

	if err := orderRepo.CreateOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("không thể tạo đơn hàng: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("không thể commit transaction: %w", err)
	}

	return order, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, id int64) (*model.Order, error) {
	repo := repository.NewOrderRepository(s.db)
	order, err := repo.GetOrderByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy đơn hàng: %w", err)
	}
	return order, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]*model.Order, error) {
	repo := repository.NewOrderRepository(s.db)
	orders, err := repo.GetAllOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy danh sách đơn hàng: %w", err)
	}
	return orders, nil
}