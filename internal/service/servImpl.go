package service

import (
	"auto/internal/repository"
	"context"
	"log"

	"github.com/pkg/errors"
)

func (s *servImpl) GetCars(ctx context.Context, filter *QueryFilter) ([]*Car, error) {

	readFilter := &repository.RepQueryFilter{
		Id:     filter.Id,
		RegNum: filter.RegNum,
		Mark:   filter.Mark,
		Model:  filter.Model,
		Owner:  filter.Owner,
		Limit:  filter.Limit,
		Offset: filter.Offset,
	}
	repCars, err := s.rep.GetCars(ctx, readFilter)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error GetCars")
	}

	Cars := make([]*Car, len(repCars))
	for index, car := range repCars {
		t := &Car{
			Id:     car.Id,
			RegNum: car.RegNum,
			Mark:   car.Mark,
			Model:  car.Model,
			Owner:  car.Owner,
		}
		Cars[index] = t
	}
	return Cars, nil
}

func (s *servImpl) GetCar(ctx context.Context, id int) (*Car, error) {
	repCar, err := s.rep.GetCar(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error GetCar")
	}
	Car := &Car{
		Id:     repCar.Id,
		RegNum: repCar.RegNum,
		Mark:   repCar.Mark,
		Model:  repCar.Model,
		Owner:  repCar.Owner,
	}
	return Car, nil
}

func (s *servImpl) DeleteCar(ctx context.Context, id int) (*Car, error) {
	repCar, err := s.rep.DeleteCar(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error DeleteCar")
	}
	car := &Car{
		Id:     repCar.Id,
		RegNum: repCar.RegNum,
		Mark:   repCar.Mark,
		Model:  repCar.Model,
		Owner:  repCar.Owner,
	}
	return car, nil
}

func (s *servImpl) AddCar(ctx context.Context, nums *RegNums) (*Car, error) {

	repCar := &Car{}
	log.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! ", nums.Nums)
	repCar, err := s.carInfoClient.GetCarInfo(ctx, nums.Nums)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error AddCar")
	}

	//

	//log.Println(repCar)

	// 	{
	// 		RegNum: Car.RegNum,
	// 		Mark:   Car.Mark,
	// 		Model:  Car.Model,
	// 		Owner:  Car.Owner,
	// 	}

	// 	repCar, err := s.rep.AddCar(ctx, repCar)
	// 	if err != nil {
	// 		return nil, errors.Wrap(err, "occurred error AddCar")
	// 	}

	// 	respCar := &Car{
	// 		Id:     repCar.Id,
	// 		RegNum: repCar.RegNum,
	// 		Mark:   repCar.Mark,
	// 		Model:  repCar.Model,
	// 		Owner:  repCar.Owner,
	// 	}
	return repCar, nil
}

func (s *servImpl) UpdateCar(ctx context.Context, car *Car) (*Car, error) {

	reposCar := &repository.RepCar{
		Id:     car.Id,
		RegNum: car.RegNum,
		Mark:   car.Mark,
		Model:  car.Model,
		Owner:  car.Owner,
	}
	reposCar, err := s.rep.UpdateCar(ctx, reposCar)
	if err != nil {
		return nil, errors.Wrap(err, "occurred error UpdateUser")
	}
	respCar := &Car{
		Id:     reposCar.Id,
		RegNum: reposCar.RegNum,
		Mark:   reposCar.Mark,
		Model:  reposCar.Model,
		Owner:  reposCar.Owner,
	}
	return respCar, nil
}
