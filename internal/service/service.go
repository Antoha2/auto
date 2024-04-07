package service

import (
	"auto/internal/config"
	"auto/internal/repository"
	"context"
	"log/slog"
)

type Repository interface {
	GetCar(ctx context.Context, id int) error
	GetCars(ctx context.Context, filter *repository.RepQueryFilter) ([]*repository.RepCar, error)
	AddCar(ctx context.Context, Car []*repository.RepCar) ([]*repository.RepCar, error)
	DeleteCar(ctx context.Context, id int) (*repository.RepCar, error)
	UpdateCar(ctx context.Context, Car *repository.RepCar) (*repository.RepCar, error)
}

type CarInfoProvider interface {
	GetCarInfo(ctx context.Context, regNums []string) ([]Car, error)
}

type servImpl struct {
	cfg           *config.Config
	log           *slog.Logger
	rep           *repository.Rep
	carInfoClient CarInfoProvider
}

func NewServ(
	cfg *config.Config,
	log *slog.Logger,
	rep *repository.Rep,
	carsInfoClient CarInfoProvider,
) *servImpl {
	return &servImpl{
		rep:           rep,
		log:           log,
		cfg:           cfg,
		carInfoClient: carsInfoClient,
	}
}

// type Cars struct {
// 	Cars []*Car `json:"cars"`
// }

// type Car struct {
// 	Id     int    `json:"id"`
// 	RegNum string `json:"regnum"`
// 	Mark   string `json:"mark"`
// 	Model  string `json:"model"`
// 	Owner  string `json:"owner"`
// }

type Car struct {
	Id     int    `json:"id"`
	RegNum string `json:"regnum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  People `json:"owner"`
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type QueryFilter struct {
	Id     int    `json:"id"`
	RegNum string `json:"regNum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  People `json:"owner"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

type RegNums struct {
	Nums []string `json:"regNums"`
}
