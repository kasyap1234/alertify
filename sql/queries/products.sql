-- name: AddProduct :one
INSERT INTO products (id, name, sku, stock_quantity, threshold) VALUES ($1,$2,$3,$4,$5) RETURNING  *;
