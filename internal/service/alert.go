package service

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AlertService struct {
	db             *pgxpool.Pool
	productService ProductService
}

func NewAlertService(pool *pgxpool.Pool, productSvc ProductService) *AlertService {
	return &AlertService{
		db:             pool,
		productService: productSvc,
	}
}

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	Threshold int       `json:"threshold"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Alert struct {
	ID           int       `json:"id"`
	ProductID    int       `json:"product_id"`
	AlertMessage string    `json:"alert_message"`
	CreatedAt    time.Time `json:"created_at"`
}
