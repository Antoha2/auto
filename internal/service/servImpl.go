package service

import (
	"auto/internal/repository"
	"context"

	"github.com/pkg/errors"
)

func (s *servImpl) GetAutos(ctx context.Context, filter *QueryFilter) ([]*Auto, error) {

	readFilter := &repository.RepQueryFilter{
		Id:     filter.Id,
		RegNum: filter.RegNum,
		Mark:   filter.Mark,
		Model:  filter.Model,
		Owner:  filter.Owner,
		Limit:  filter.Limit,
		Offset: filter.Offset,
	}
	repAutos, err := s.rep.GetAutos(ctx, readFilter)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error GetAutos")
	}

	autos := make([]*Auto, len(repAutos))
	for index, auto := range repAutos {
		t := &Auto{
			Id:     auto.Id,
			RegNum: auto.RegNum,
			Mark:   auto.Mark,
			Model:  auto.Model,
			Owner:  auto.Owner,
		}
		autos[index] = t
	}
	return autos, nil
}

func (s *servImpl) GetAuto(ctx context.Context, id int) (*Auto, error) {
	repAuto, err := s.rep.GetAuto(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error GetAuto")
	}
	auto := &Auto{
		Id:     repAuto.Id,
		RegNum: repAuto.RegNum,
		Mark:   repAuto.Mark,
		Model:  repAuto.Model,
		Owner:  repAuto.Owner,
	}
	return auto, nil
}

func (s *servImpl) DeleteAuto(ctx context.Context, id int) (*Auto, error) {
	repAuto, err := s.rep.DeleteAuto(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error DeleteAuto")
	}
	auto := &Auto{
		Id:     repAuto.Id,
		RegNum: repAuto.RegNum,
		Mark:   repAuto.Mark,
		Model:  repAuto.Model,
		Owner:  repAuto.Owner,
	}
	return auto, nil
}

func (s *servImpl) AddAuto(ctx context.Context, auto *Auto) (*Auto, error) {

	repAuto := &repository.RepAuto{
		RegNum: auto.RegNum,
		Mark:   auto.Mark,
		Model:  auto.Model,
		Owner:  auto.Owner,
	}

	repAuto, err := s.rep.AddAuto(ctx, repAuto)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error AddAuto")
	}

	respAuto := &Auto{
		Id:     repAuto.Id,
		RegNum: repAuto.RegNum,
		Mark:   repAuto.Mark,
		Model:  repAuto.Model,
		Owner:  repAuto.Owner,
	}
	return respAuto, nil
}

func (s *servImpl) UpdateAuto(ctx context.Context, auto *Auto) (*Auto, error) {

	reposAuto := &repository.RepAuto{
		Id:     auto.Id,
		RegNum: auto.RegNum,
		Mark:   auto.Mark,
		Model:  auto.Model,
		Owner:  auto.Owner,
	}
	reposAuto, err := s.rep.UpdateAuto(ctx, reposAuto)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error UpdateUser")
	}
	respAuto := &Auto{
		Id:     reposAuto.Id,
		RegNum: reposAuto.RegNum,
		Mark:   reposAuto.Mark,
		Model:  reposAuto.Model,
		Owner:  reposAuto.Owner,
	}
	return respAuto, nil
}
