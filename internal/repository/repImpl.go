package repository

import (
	"context"
)

func (r *Rep) GetAuto(ctx context.Context, id int) (*RepAuto, error) {
	auto := new(RepAuto)
	return auto, nil
}

func (r *Rep) GetAutos(ctx context.Context, filter *RepQueryFilter) ([]*RepAuto, error) {

	return nil, nil
}

func (r *Rep) AddAuto(ctx context.Context, auto *RepAuto) (*RepAuto, error) {

	return auto, nil
}

func (r *Rep) DeleteAuto(ctx context.Context, id int) (*RepAuto, error) {

	auto := new(RepAuto)

	return auto, nil
}

func (r *Rep) UpdateAuto(ctx context.Context, auto *RepAuto) (*RepAuto, error) {

	return auto, nil
}
