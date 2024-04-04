package service

import "context"

func (s *servImpl) GetAutos(ctx context.Context, filter *QueryFilter) ([]*Auto, error) {

	return nil, nil
}

func (s *servImpl) GetAuto(ctx context.Context, id int) (*Auto, error) {
	auto := new(Auto)
	return auto, nil
}

func (s *servImpl) DeleteAuto(ctx context.Context, id int) (*Auto, error) {
	auto := new(Auto)
	return auto, nil
}

func (s *servImpl) AddAuto(ctx context.Context, auto *Auto) (*Auto, error) {

	return auto, nil
}

func (s *servImpl) UpdateAuto(ctx context.Context, auto *Auto) (*Auto, error) {

	return auto, nil
}
