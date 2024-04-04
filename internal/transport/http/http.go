package transport

import (
	"auto/internal/config"
	"auto/internal/service"
	"context"
	"log/slog"
	"net/http"
)

type Service interface {
	GetAutos(ctx context.Context, filter *service.QueryFilter) ([]*service.Auto, error)
	GetAuto(ctx context.Context, id int) (*service.Auto, error)
	DeleteAuto(ctx context.Context, id int) (*service.Auto, error)
	AddAuto(ctx context.Context, auto *service.Auto) (*service.Auto, error)
	UpdateAuto(ctx context.Context, auto *service.Auto) (*service.Auto, error)
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
