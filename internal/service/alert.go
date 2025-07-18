package service

import (
	"context"
	"time"

	"alertify/internal/db"
	"alertify/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type AlertService struct {
	db             *pgxpool.Pool
	productService ProductService
	queries        *db.Queries
}

func NewAlertService(pool *pgxpool.Pool, productSvc ProductService, queries *db.Queries) *AlertService {
	return &AlertService{
		db:             pool,
		productService: productSvc,
		queries:        queries,
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
	AlertType    string    `json:"alert_type"`
	CreatedAt    time.Time `json:"created_at"`
}

func (s *AlertService) CreateAlert(ctx context.Context, product_id int32, alert_message string, alert_type string, status string) error {
	// check if product exists or not
	_, err := s.productService.queries.GetProductByID(ctx, product_id)
	if err != nil {
		return err
	}
	args := db.CreateAlertParams{
		ProductID:    utils.ToPgInt4(product_id),
		AlertMessage: alert_message,
		AlertType:    alert_type,
		Status:       status,
	}
	//return s.queries.CreateAlert(ctx, args)
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Error().Err(err).Msg("failed to begin transaction")
		return err
	}
	defer tx.Rollback(ctx)
	qtx := db.New(tx)
	err = qtx.CreateAlert(ctx, args)
	if err := tx.Commit(ctx); err != nil {
		log.Error().Err(err).Msg("failed to commit transaction")
	}
	return err
}

func (s *AlertService) GetAlertsByStatus(ctx context.Context, status string) ([]db.Alert, error) {
	alerts, err := s.queries.GetAlertsByStatus(ctx, status)
	if err != nil {
		log.Error().Err(err).Msg("failed to get alert by status")
		return nil, err
	}
	return alerts, nil
}

func (s *AlertService) GetAllAlerts(ctx context.Context) ([]db.Alert, error) {
	alerts, err := s.queries.GetAllAlerts(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get all alerts ")
		return nil, err
	}
	return alerts, nil
}

func (s *AlertService) GetPendingAlerts(ctx context.Context) ([]db.Alert, error) {
	alerts, err := s.queries.GetPendingAlerts(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get pending alerts ")
		return nil, err
	}
	return alerts, nil
}

func (s *AlertService) UpdateAlert(ctx context.Context, id string, status string) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Error().Err(err).Msg("update alert begin transaction failed ")
		return err
	}
	defer tx.Rollback(ctx)
	qtx := db.New(tx)
	uuidID := uuid.Scan()
	args := db.UpdateAlertParams{
		ID:     id,
		Status: status,
	}
	err := qtx.UpdateAlert(ctx, args)

}
