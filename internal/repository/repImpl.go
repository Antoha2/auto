package repository

import (
	"context"
)

func (r *Rep) GetCar(ctx context.Context, id int) (*RepCar, error) {
	Car := new(RepCar)
	return Car, nil
}

func (r *Rep) GetCars(ctx context.Context, filter *RepQueryFilter) ([]*RepCar, error) {

	return nil, nil
}

func (r *Rep) AddCar(ctx context.Context, Car *RepCar) (*RepCar, error) {

	return Car, nil
}

func (r *Rep) DeleteCar(ctx context.Context, id int) (*RepCar, error) {

	Car := new(RepCar)

	return Car, nil
}

func (r *Rep) UpdateCar(ctx context.Context, Car *RepCar) (*RepCar, error) {

	return Car, nil
}
