package service

import (
	"auto/internal/config"
	"auto/internal/repository"
	"context"
	"log/slog"
)

const DefaultPropertyAge = 0
const DefaultPropertyOffset = 0
const DefaultPropertyLimit = 100

type Repository interface {
	GetAuto(ctx context.Context, id int) error
	GetAutos(ctx context.Context, filter *repository.RepQueryFilter) ([]*repository.RepAuto, error)
	AddAuto(ctx context.Context, auto *repository.RepAuto) (int, error)
	DeleteAuto(ctx context.Context, id int) (*repository.RepAuto, error)
	UpdateAuto(ctx context.Context, auto *repository.RepAuto) (*repository.RepAuto, error)
}

type servImpl struct {
	cfg *config.Config
	log *slog.Logger
	rep *repository.Rep
}

func NewServ(
	cfg *config.Config,
	log *slog.Logger,
	rep *repository.Rep,
) *servImpl {
	return &servImpl{
		rep: rep,
		log: log,
		cfg: cfg,
	}
}

type Auto struct {
	Id     int    `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Owner  string `json:"owner"`
}

type QueryFilter struct {
	Id     int    `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Owner  string `json:"owner"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}
