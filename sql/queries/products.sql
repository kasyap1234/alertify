-- name: AddProduct :one
INSERT INTO products (name, sku, stock_quantity, threshold) VALUES ($1,$2,$3,$4) RETURNING  *;


-- name: GetProduct :one
SELECT * FROM products WHERE id=$1;

-- name: AddProducts :copyfrom
INSERT INTO products (id,name,sku,stock_quantity,threshold) VALUES ($1,$2,$3,$4,$5);

-- name: GetProductsBySKU :one
SELECT * FROM products WHERE sku=$1;

-- name: ListProducts :many
SELECT * FROM products ORDER BY created_at DESC;

-- name: UpdateProductStock :exec
UPDATE  products SET stock_quantity=$1 WHERE id=$2;

-- name: GetProductQuantity :one
SELECT products.stock_quantity FROM products WHERE id=$1;

-- name: GetProductsByName :many
SELECT * FROM products WHERE name=$1;

-- name: SearchProductsByName :many
SELECT * FROM products WHERE name ILIKE $1 || '%';


-- name: GetLowStockProducts :many
SELECT * FROM products WHERE stock_quantity < threshold;

-- name: IncrementProductStock :exec
UPDATE products SET stock_quantity=stock_quantity +$1 WHERE id=$2;

-- name: DecrementProductStock :exec
UPDATE products SET stock_quantity=stock_quantity - $1 WHERE id=$2;


-- name: ListProductsPaginated :many
SELECT  * FROM products ORDER BY created_at DESC LIMIT $1 OFFSET $2;