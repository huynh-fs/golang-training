-- Xóa bảng cũ nếu tồn tại để tạo lại từ đầu
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;

-- Tạo lại bảng products với cột quantity
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT quantity_non_negative CHECK (quantity >= 0)
);

-- Tạo bảng orders
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    quantity_ordered INT NOT NULL,
    order_status VARCHAR(50) NOT NULL, -- e.g., 'COMPLETED', 'FAILED'
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- Thêm dữ liệu mẫu với số lượng tồn kho
INSERT INTO products (name, price, quantity) VALUES ('Laptop XPS 15', 2500.50, 10);
INSERT INTO products (name, price, quantity) VALUES ('Bàn phím cơ Keychron', 150.75, 50);
INSERT INTO products (name, price, quantity) VALUES ('Màn hình Dell UltraSharp', 890.99, 25);