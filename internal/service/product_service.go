package service

import (
	"context"
	"errors"
	"strings"

	"alertify/internal/db"
	"alertify/internal/utils"

	"github.com/rs/zerolog/log"
)

type ProductService struct {
	q *db.Queries
}

func NewProductService(q *db.Queries) *ProductService {
	return &ProductService{q: q}
}

type CreateProductInput struct {
	Name          string
	Sku           string
	StockQuantity int32
	Threshold     int32
}

func (s *ProductService) AddProduct(ctx context.Context, input CreateProductInput) (*db.Product, error) {
	if strings.TrimSpace(input.Name) == "" {
		err := errors.New("product name is required")
		log.Error().Err(err).Msg("validation failed")
		return nil, err
	}
	if input.StockQuantity < 0 || input.Threshold < 0 {
		err := errors.New("quantity should be a positive number")
		log.Error().Err(err).Msg("invalid quantity or threshold")
		return nil, err

	}
	params := db.AddProductParams{
		Name:          input.Name,
		Sku:           utils.ToPgText(input.Sku),
		StockQuantity: input.StockQuantity,
		Threshold:     input.Threshold,
	}

	product, err := s.q.AddProduct(ctx, params)
	if err != nil {
		log.Error().Err(err).Msg("failed to add product")
		return nil, err
	}
	return &product, nil
}
