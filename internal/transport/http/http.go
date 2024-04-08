package transport

import (
	"auto/internal/config"
	"auto/internal/service"
	"context"
	"log/slog"
	"net/http"
)

type Service interface {
	GetCars(ctx context.Context, filter *service.QueryFilter) ([]*service.Car, error)
	GetCar(ctx context.Context, id int) (*service.Car, error)
	DeleteCar(ctx context.Context, id int) (*service.Car, error)
	AddCar(ctx context.Context, car *service.RegNums) ([]*service.Car, error)
	UpdateCar(ctx context.Context, car *service.Car) (*service.Car, error)
}

type apiImpl struct {
	cfg     *config.Config
	log     *slog.Logger
	service Service
	server  *http.Server
}

// NewAPI
func NewApi(cfg *config.Config, log *slog.Logger, service Service) *apiImpl {
	return &apiImpl{
		service: service,
		log:     log,
		cfg:     cfg,
	}
}
