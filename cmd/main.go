package main

import (
	"auto/internal/config"
	provider "auto/internal/providers/carInfo"
	"auto/internal/repository"
	"auto/internal/service"
	transport "auto/internal/transport/http"
	"auto/pkg/logger"
	"auto/pkg/logger/sl"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// @Param mark query string true "mark"
// @Param model query string true "model"
// @Param year query int true "year"
// @Param owner.name query string true "name"
// @Param owner.surname query string true "surname"
// @Param owner.patronymic query string true "patronymic"

func main() {
	Run()
}

func Run() {
	cfg := config.MustLoad()
	slogger := logger.SetupLogger(cfg.Env)
	dbx := MustInitDb(cfg)

	rep := repository.NewRep(slogger, dbx)

	pGetCarsInfo := provider.NewGetCarInfo(cfg.URLGetCarInfo)

	serv := service.NewServ(cfg, slogger, rep, pGetCarsInfo)
	trans := transport.NewApi(cfg, slogger, serv)

	go trans.StartHTTP()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	trans.Stop()
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
