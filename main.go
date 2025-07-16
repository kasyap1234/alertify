package main

import (
	"context"
	"flag"
	"os"
	"time"

	config2 "alertify/internal/config"
	"alertify/internal/db"
	"alertify/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Setup logging

	// Load config
	cfg, err := config2.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("error loading config")
	}

	defer pool.Close()
	queries := db.New(pool)

}
