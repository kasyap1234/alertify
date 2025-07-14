package main

import (
	"context"
	"flag"
	"os"

	"alertify/config"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
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
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("error loading config")
	}
	log.Debug().Msgf("config db %v %v ", cfg.DB_Host, cfg.DB_Name)
	//unc SetConnectionString(host string, user string, password string, name string, port in
	dsn := config.SetConnectionString(cfg.DB_Host, cfg.DB_User, cfg.DB_Password, cfg.DB_Name, cfg.DB_Port)
	_, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Error().Err(err).Msg("connection failed ")
		return
	}
	log.Info().Msgf("connecting to database")

}
