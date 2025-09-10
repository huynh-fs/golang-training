package service

import (
	"database/sql"
	"fmt"
	"context"

	"github.com/huynh-fs/sql/internal/model"
	"github.com/huynh-fs/sql/internal/repository"
)

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) CreateProduct(ctx context.Context, name string, price float64, quantity int) (*model.Product, error) {
	repo := repository.NewProductRepository(s.db)
	product := &model.Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
	if err := repo.CreateProduct(ctx, product); err != nil {
		return nil, fmt.Errorf("không thể tạo sản phẩm: %w", err)
	}
	return product, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id int64) (*model.Product, error) {
	repo := repository.NewProductRepository(s.db)
	product, err := repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy sản phẩm: %w", err)
	}
	return product, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]*model.Product, error) {
	repo := repository.NewProductRepository(s.db)
	products, err := repo.GetAllProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy danh sách sản phẩm: %w", err)
	}
	return products, nil
}

func (s *ProductService) UpdateQuantity(ctx context.Context, id int64, quantity int) error {
	repo := repository.NewProductRepository(s.db)
	product , err := repo.GetProductByID(ctx, id)
	if err != nil {
		return fmt.Errorf("không thể lấy sản phẩm: %w", err)
	}
	if product == nil {
		return fmt.Errorf("sản phẩm với ID %d không tồn tại", id)
	}
	product.Quantity += quantity
	if err := repo.UpdateProduct(ctx, product); err != nil {
		return fmt.Errorf("không thể cập nhật số lượng sản phẩm: %w", err)
	}
	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	repo := repository.NewProductRepository(s.db)
	if err := repo.DeleteProduct(ctx, id); err != nil {
		return fmt.Errorf("không thể xóa sản phẩm: %w", err)
	}
	return nil
}
