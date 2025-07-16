package bootstrap

import (
	"alertify/internal/db"
	"alertify/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitServices(pool *pgxpool.Pool) *service.Service {
	queries := db.New(pool)
	productSvc := service.NewProductService(pool)
	// alertSvc:=
	alertSvc := service.NewAlertService(pool, productSvc)
	svc, _ := service.NewService(*productSvc, alertSvc)
	return svc
}
