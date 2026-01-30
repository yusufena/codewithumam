-- Add category_id to products table
ALTER TABLE products ADD COLUMN IF NOT EXISTS category_id INT;

-- Add foreign key constraint
ALTER TABLE products 
ADD CONSTRAINT fk_category 
FOREIGN KEY (category_id) 
REFERENCES categories (id)
ON DELETE SET NULL;
