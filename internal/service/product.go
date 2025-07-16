package service

import (
	"context"
	"errors"
	"strings"

	"alertify/internal/db"
	"alertify/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type ProductService struct {
	db      *pgxpool.Pool
	queries *db.Queries
}

func NewProductService(db *pgxpool.Pool) *ProductService {
	return &ProductService{db: db}
}

type CreateProductInput struct {
	Name          string
	Sku           string
	StockQuantity int32
	Threshold     int32
}

func (s *ProductService) AddProduct(ctx context.Context, input CreateProductInput) (*db.Product, error) {
	// Input validation
	if strings.TrimSpace(input.Name) == "" {
		err := errors.New("product name is required")
		log.Error().Err(err).Msg("validation failed")
		return nil, err
	}
	if input.StockQuantity < 0 || input.Threshold < 0 {
		err := errors.New("quantity and threshold must be positive")
		log.Error().Err(err).Msg("invalid quantity or threshold")
		return nil, err
	}

	// Start transaction
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Error().Err(err).Msg("failed to begin transaction")
		return nil, err
	}
	defer tx.Rollback(ctx) // rollback unless committed

	qtx := db.New(tx)

	// Prepare params
	params := db.AddProductParams{
		Name:          input.Name,
		Sku:           utils.ToPgText(input.Sku),
		StockQuantity: input.StockQuantity,
		Threshold:     input.Threshold,
	}

	product, err := qtx.AddProduct(ctx, params)
	if err != nil {
		log.Error().Err(err).Msg("failed to add product")
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Error().Err(err).Msg("failed to commit transaction")
		return nil, err
	}

	// Success
	return &product, nil
}

func (s *ProductService) GetLowStockProducts(ctx context.Context) ([]db.Product, error) {
	products, err := s.queries.GetLowStockProducts(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get low stock products")
	}
	return products, nil
}
