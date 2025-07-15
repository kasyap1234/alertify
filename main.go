package main

import (
	"context"
	"flag"
	"os"
	"time"

	config2 "alertify/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Setup logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("production mode enabled")
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("debug mode enabled")
	}

	// Load config
	cfg, err := config2.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("error loading config")
	}

	log.Debug().Msgf("config db host=%v name=%v", cfg.DB_Host, cfg.DB_Name)

	// DSN & pgxpool config
	dsn := config2.SetConnectionString(cfg.DB_Host, cfg.DB_User, cfg.DB_Password, cfg.DB_Name, cfg.DB_Port)
	settings, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse DB config")
	}

	settings.MinConns = int32(cfg.MinConns)
	settings.MaxConns = int32(cfg.MaxConns)
	settings.MaxConnLifetime = cfg.MaxConnLifeTime
	settings.MaxConnIdleTime = cfg.MaxConnIdleTime
	settings.HealthCheckPeriod = cfg.HealthCheckPeriod

	// Setup connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, settings)
	if err != nil {
		log.Fatal().Err(err).Msg("pgx pool creation failed")
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatal().Err(err).Msg("ping to the DB failed")
	}

	log.Info().Msg("connected to database")
	defer pool.Close()

}
