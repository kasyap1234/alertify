-- name: CreateAlert :one
INSERT INTO alerts (product_id,alert_message,alert_type,status) VALUES ($1,$2,$3,$4) RETURNING  *;

-- name: ListAllAlerts :many
SELECT * FROM alerts ORDER BY DESC created_at;