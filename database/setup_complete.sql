-- 1. Create Categories Table
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- 2. Create Products Table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    category_id INT,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE SET NULL
);

-- 3. Insert Dummy Data (Optional)
INSERT INTO categories (name, description) VALUES 
('Makanan', 'Aneka makanan ringan dan berat'),
('Minuman', 'Aneka minuman segar');

INSERT INTO products (name, price, stock, category_id) VALUES 
('Indomie Goreng', 3500, 100, 1),
('Teh Botol', 5000, 50, 2);
