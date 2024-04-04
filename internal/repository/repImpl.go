package repository

import (
	"context"
)

func (r *Rep) GetAuto(ctx context.Context, id int) error {

	return nil
}

func (r *Rep) GetAutos(ctx context.Context, filter RepQueryFilter) ([]RepAuto, error) {

	return nil, nil
}

func (r *Rep) AddAuto(ctx context.Context, auto RepAuto) (int, error) {

	return 0, nil
}

func (r *Rep) DeleteAuto(ctx context.Context, id int) (RepAuto, error) {

	auto := RepAuto{}

	return auto, nil
}

func (r *Rep) UpdateAuto(ctx context.Context, auto RepAuto) (RepAuto, error) {

	return auto, nil
}
