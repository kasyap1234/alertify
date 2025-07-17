-- name: CreateAlert :exec
INSERT INTO alerts (product_id,alert_message,alert_type,status) VALUES ($1,$2,$3,$4) RETURNING  *;

-- name: ListAllAlerts :many
SELECT * FROM alerts ORDER BY created_at DESC;

-- name: GetPendingAlerts :many
SELECT * FROM alerts WHERE status="pending" ORDER BY created_at DESC;

-- name: UpdateAlert :exec
UPDATE alerts SET status=$2 WHERE id=$1;

-- name: GetAlertsByStatus :many
SELECT * FROM alerts WHERE status=$1 ORDER BY created__at DESC;
