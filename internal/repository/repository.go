package repository

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type Rep struct {
	log *slog.Logger
	DB  *sqlx.DB
}

func NewRep(log *slog.Logger, dbx *sqlx.DB) *Rep {
	return &Rep{
		log: log,
		DB:  dbx,
	}
}

type RepCar struct {
	Id     int
	RegNum string
	Mark   string
	Model  string
	Owner  string
}

type RepQueryFilter struct {
	Id     int
	RegNum string
	Mark   string
	Model  string
	Owner  string
	Offset int
	Limit  int
}
