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
	RepTypeMaping(car *Car) *repository.RepCar
	ServTypeMaping(respCar *repository.RepCar) *Car
}

type CarInfoProvider interface {
	GetCarInfo(ctx context.Context, regNums []string) ([]*Car, error)
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

// Car model info
// @Description Car info information
// @Description with regNum
type Car struct {
	Id     int    `json:"id"  example:"1"`
	RegNum string `json:"regnum" example:"X123XX150"`
	Mark   string `json:"mark" example:"Lada"`
	Model  string `json:"model" example:"Vesta"`
	Year   int    `json:"year" example:"2020"`
	Owner  People `json:"owner"`
}

type People struct {
	Name       string `json:"name" example:"name"`
	Surname    string `json:"surname" example:"surname"`
	Patronymic string `json:"patronymic" example:"patronymic"`
}

type Cars struct {
	Cars []Car
}

// CarQuery model info
// @Description Car info information
// @Description with Limit and Offset
type QueryFilter struct {
	Id     int    `json:"id" example:"1"`
	RegNum string `json:"regNum" example:"X123XX150"`
	Mark   string `json:"mark" example:"Lada"`
	Model  string `json:"model" example:"Vesta"`
	Year   int    `json:"year" example:"2020"`
	Owner  People `json:"owner"`
	Offset int    `json:"offset" example:"0"`
	Limit  int    `json:"limit" example:"100"`
}

// RegNum model info
// @Description regNum list
type RegNums struct {
	Nums []string `json:"regNums" example:"X123XX150"`
}
