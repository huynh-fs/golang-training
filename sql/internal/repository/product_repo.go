package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huynh-fs/sql/internal/model"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type ProductRepo struct {
	db DBTX
}

func NewProductRepository(db DBTX) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) CreateProduct(ctx context.Context, p *model.Product) error {
	query := `INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := r.db.QueryRowContext(ctx, query, p.Name, p.Price, p.Quantity).Scan(&p.ID, &p.CreatedAt)
	if err != nil {
		return fmt.Errorf("không thể tạo sản phẩm: %w", err)
	}
	return nil
}

func (r *ProductRepo) GetProductByID(ctx context.Context, id int64) (*model.Product, error) {
	query := `SELECT id, name, price, quantity, created_at FROM products WHERE id = $1`
	p := &model.Product{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Quantity, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy sản phẩm: %w", err)
	}
	return p, nil
}

func (r *ProductRepo) GetAllProducts(ctx context.Context) ([]*model.Product, error) {
	query := `SELECT id, name, price, quantity, created_at FROM products`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("không thể lấy danh sách sản phẩm: %w", err)
	}
	defer rows.Close()

	products := []*model.Product{}
	for rows.Next() {
		p := &model.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("không thể quét sản phẩm: %w", err)
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("không thể lấy danh sách sản phẩm: %w", err)
	}
	return products, nil
}

func (r *ProductRepo) UpdateProduct(ctx context.Context, p *model.Product) error {
	query := `UPDATE products SET name = $1, price = $2, quantity = $3 WHERE id = $4`
	result, err := r.db.ExecContext(ctx, query, p.Name, p.Price, p.Quantity, p.ID)
	if err != nil {
		return fmt.Errorf("không thể cập nhật sản phẩm: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("không thể lấy số dòng bị ảnh hưởng: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("không tìm thấy sản phẩm để cập nhật")
	}
	return nil
}

func (r *ProductRepo) DeleteProduct(ctx context.Context, id int64) error {
	query := `DELETE FROM products WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("không thể xóa sản phẩm: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("không thể lấy số dòng bị ảnh hưởng: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("không tìm thấy sản phẩm để xóa")
	}
	return nil
}

func (r *ProductRepo) DecreaseInStock(ctx context.Context, id int64, quantity int) error {
	query := `UPDATE products SET quantity = quantity - $1 WHERE id = $2 AND quantity >= $1`
	result, err := r.db.ExecContext(ctx, query, quantity, id)
	if err != nil {
		return fmt.Errorf("không thể giảm số lượng sản phẩm: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("không thể lấy số dòng bị ảnh hưởng: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("không đủ số lượng sản phẩm trong kho hoặc sản phẩm không tồn tại")
	}
	return nil
}
