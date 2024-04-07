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

		car := &Car{}

		car.Id = repCar.Id
		car.RegNum = repCar.RegNum
		car.Mark = repCar.Mark
		car.Model = repCar.Model
		car.Year = repCar.Year
		car.Owner.Name = repCar.Name
		car.Owner.Surname = repCar.Surname
		car.Owner.Patronymic = repCar.Patronymic

		Cars[index] = car
	}
	return Cars, nil
}

// get car
func (s *servImpl) GetCar(ctx context.Context, id int) (*Car, error) {
	repCar, err := s.rep.GetCar(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error GetCar")
	}

	car := &Car{}

	car.Id = repCar.Id
	car.RegNum = repCar.RegNum
	car.Mark = repCar.Mark
	car.Model = repCar.Model
	car.Year = repCar.Year
	car.Owner.Name = repCar.Name
	car.Owner.Surname = repCar.Surname
	car.Owner.Patronymic = repCar.Patronymic

	return car, nil
}

// delete car
func (s *servImpl) DeleteCar(ctx context.Context, id int) (*Car, error) {
	repCar, err := s.rep.DeleteCar(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error DeleteCar")
	}
	car := &Car{}

	car.Id = repCar.Id
	car.RegNum = repCar.RegNum
	car.Mark = repCar.Mark
	car.Model = repCar.Model
	car.Year = repCar.Year
	car.Owner.Name = repCar.Name
	car.Owner.Surname = repCar.Surname
	car.Owner.Patronymic = repCar.Patronymic

	return car, nil
}

//add car
func (s *servImpl) AddCar(ctx context.Context, nums *RegNums) ([]Car, error) {

	Cars, err := s.carInfoClient.GetCarInfo(ctx, nums.Nums)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error AddCar")
	}

	repCars := make([]*repository.RepCar, len(Cars))

	for i, car := range Cars {

		repCar := &repository.RepCar{
			RegNum:     car.RegNum,
			Mark:       car.Mark,
			Model:      car.Model,
			Year:       car.Year,
			Name:       car.Owner.Name,
			Surname:    car.Owner.Surname,
			Patronymic: car.Owner.Patronymic,
		}

		repCars[i] = repCar
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

	repCar := &repository.RepCar{}

	repCar.Id = car.Id
	repCar.RegNum = car.RegNum
	repCar.Mark = car.Mark
	repCar.Model = car.Model
	repCar.Year = car.Year
	repCar.Name = car.Owner.Name
	repCar.Surname = car.Owner.Surname
	repCar.Patronymic = car.Owner.Patronymic

	respCar, err := s.rep.UpdateCar(ctx, repCar)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error UpdateUser")
	}

	car.Id = respCar.Id
	car.RegNum = respCar.RegNum
	car.Mark = respCar.Mark
	car.Model = respCar.Model
	car.Year = respCar.Year
	car.Owner.Name = respCar.Name
	car.Owner.Surname = respCar.Surname
	car.Owner.Patronymic = respCar.Patronymic
	return car, nil
}
