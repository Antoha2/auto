package main

import (
	"auto/internal/config"
	"auto/pkg/logger/sl"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	Run()
}

func Run() {

}

func MustInitDb(cfg *config.Config) *sqlx.DB {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Dbname,
		cfg.DBConfig.Sslmode,
	)

	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		slog.Warn("failed to parse config", sl.Err(err))
		os.Exit(1)
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		slog.Warn("failed to create connection db", sl.Err(err))
		os.Exit(1)
	}

	err = dbx.Ping()
	if err != nil {
		slog.Warn("error to ping connection pool", sl.Err(err))
		os.Exit(1)
	}
	slog.Info(fmt.Sprintf("Подключение к базе данных на http://127.0.0.1:%v\n", cfg.DBConfig.Port))
	return dbx
}
