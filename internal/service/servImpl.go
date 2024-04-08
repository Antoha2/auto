package service

import (
	"auto/internal/repository"
	"context"

	"github.com/pkg/errors"
)

// get Cars
func (s *servImpl) GetCars(ctx context.Context, filter *QueryFilter) ([]*Car, error) {

	readFilter := &repository.RepQueryFilter{
		Id:         filter.Id,
		RegNum:     filter.RegNum,
		Mark:       filter.Mark,
		Model:      filter.Model,
		Year:       filter.Year,
		Name:       filter.Owner.Name,
		Surname:    filter.Owner.Surname,
		Patronymic: filter.Owner.Patronymic,
		Limit:      filter.Limit,
		Offset:     filter.Offset,
	}

	repCars, err := s.rep.GetCars(ctx, readFilter)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error GetCars")
	}

	Cars := make([]*Car, len(repCars))
	for index, repCar := range repCars {
		Cars[index] = s.ServTypeMaping(repCar)
	}
	return Cars, nil
}

// get car
func (s *servImpl) GetCar(ctx context.Context, id int) (*Car, error) {

	repCar, err := s.rep.GetCar(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error GetCar")
	}

	return s.ServTypeMaping(repCar), nil
}

// delete car
func (s *servImpl) DeleteCar(ctx context.Context, id int) (*Car, error) {
	repCar, err := s.rep.DeleteCar(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error DeleteCar")
	}

	return s.ServTypeMaping(repCar), nil
}

//add car
func (s *servImpl) AddCar(ctx context.Context, nums *RegNums) ([]*Car, error) {

	if len(nums.Nums) == 0 {
		return nil, errors.Wrap(errors.New("No add data"), "occurred error AddCar")
	}

	Cars, err := s.carInfoClient.GetCarInfo(ctx, nums.Nums)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error AddCar")
	}

	repCars := make([]*repository.RepCar, len(Cars))
	for i, car := range Cars {
		repCars[i] = s.RepTypeMaping(car)
	}

	respCars, err := s.rep.AddCar(ctx, repCars)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error AddCar")
	}

	for i := 0; i < len(Cars); i++ {
		Cars[i].Id = respCars[i].Id
	}

	return Cars, nil
}

//update car
func (s *servImpl) UpdateCar(ctx context.Context, car *Car) (*Car, error) {

	respCar, err := s.rep.UpdateCar(ctx, s.RepTypeMaping(car))
	if err != nil {
		return nil, errors.Wrap(err, "occurred error UpdateCar")
	}

	return s.ServTypeMaping(respCar), nil
}

//change type service to repository
func (s *servImpl) RepTypeMaping(car *Car) *repository.RepCar {

	repCar := &repository.RepCar{}

	repCar.Id = car.Id
	repCar.RegNum = car.RegNum
	repCar.Mark = car.Mark
	repCar.Model = car.Model
	repCar.Year = car.Year
	repCar.Name = car.Owner.Name
	repCar.Surname = car.Owner.Surname
	repCar.Patronymic = car.Owner.Patronymic

	return repCar
}

//change type repository to service
func (s *servImpl) ServTypeMaping(respCar *repository.RepCar) *Car {

	car := &Car{}

	car.Id = respCar.Id
	car.RegNum = respCar.RegNum
	car.Mark = respCar.Mark
	car.Model = respCar.Model
	car.Year = respCar.Year
	car.Owner.Name = respCar.Name
	car.Owner.Surname = respCar.Surname
	car.Owner.Patronymic = respCar.Patronymic

	return car
}
