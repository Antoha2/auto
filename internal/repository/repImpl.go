package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

func (r *Rep) GetCar(ctx context.Context, id int) (*RepCar, error) {

	car := new(RepCar)

	query := "SELECT id, regnum, mark, model, owners FROM cars WHERE id = $1"
	row := r.DB.QueryRowContext(ctx, query, id)
	if err := row.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Owner); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql select Car failed, query: %s", query))
	}

	return car, nil
}

func (r *Rep) GetCars(ctx context.Context, filter *RepQueryFilter) ([]*RepCar, error) {

	return nil, nil
}

func (r *Rep) AddCar(ctx context.Context, Cars []*RepCar) ([]RepCar, error) {

	s := buildQueryAddConstrain(Cars)
	query := fmt.Sprintf("INSERT INTO cars (regnum, mark, model, owners) VALUES %s RETURNING id, regnum, mark, model, owners", s)

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql add Cars failed, query: %s", query))
	}

	defer rows.Close()

	repCars := make([]RepCar, 0, len(Cars))
	car := RepCar{}

	for rows.Next() {
		err := rows.Scan(&car.Id, &car.Mark, &car.Model, &car.Owner, &car.RegNum)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("sql add Cars failed, query: %s", query))
		}
		repCars = append(repCars, car)
	}
	log.Println(repCars)
	return repCars, nil
}

func (r *Rep) DeleteCar(ctx context.Context, id int) (*RepCar, error) {

	car := new(RepCar)

	query := "DELETE FROM cars WHERE id = $1 RETURNING id, regnum, mark, model, owners"
	row := r.DB.QueryRowContext(ctx, query, id)
	if err := row.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Owner); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql delete car failed, query: %s", query))
	}

	return car, nil
}

func (r *Rep) UpdateCar(ctx context.Context, car *RepCar) (*RepCar, error) {

	repCar := &RepCar{}

	query := "UPDATE cars SET regnum=$1, mark=$2, model=$3, ownres=$4 WHERE id=$5 RETURNING id, regnum, mark, model, owners"

	row := r.DB.QueryRowContext(ctx, query, &car.RegNum, &car.Mark, &car.Model, &car.Owner, car.Id)
	if err := row.Scan(&repCar.Id, &repCar.RegNum, &repCar.Mark, &repCar.Model, &repCar.Owner); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql update car failed, query: %s", query))
	}

	return repCar, nil
}

//build query string
func buildQueryAddConstrain(cars []*RepCar) string {
	constrains := make([]string, 0, len(cars))
	for _, car := range cars {
		s := fmt.Sprintf("('%s','%s','%s','%s')", car.RegNum, car.Mark, car.Model, car.Owner)
		constrains = append(constrains, s)
	}

	queryConstrain := strings.Join(constrains, ",")

	return queryConstrain
}
