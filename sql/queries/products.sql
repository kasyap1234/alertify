-- name: AddProduct :one
INSERT INTO products (id, name, sku, stock_quantity, threshold) VALUES ($1,$2,$3,$4,$5) RETURNING  *;


-- name: GetProduct :one
SELECT * FROM products WHERE id=$1;

-- name: AddProducts :copyfrom
INSERT INTO products (id,name,sku,stock_quantity,threshold) VALUES ($1,$2,$3,$4,$5);

-- name: GetProductsBySKU :many
SELECT * FROM products WHERE sku=$1;
