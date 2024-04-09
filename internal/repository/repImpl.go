package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

//get car
func (r *Rep) GetCar(ctx context.Context, id int) (*RepCar, error) {

	car := new(RepCar)

	query := "SELECT id, regnum, mark, model, year, name, surname, patronymic FROM cars WHERE id = $1"
	row := r.DB.QueryRowContext(ctx, query, id)
	if err := row.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Name, &car.Surname, &car.Patronymic); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql select Car failed, query: %s", query))
	}

	return car, nil
}

//get cars
func (r *Rep) GetCars(ctx context.Context, filter *RepQueryFilter) ([]*RepCar, error) {

	cars := make([]*RepCar, 0)
	queryConstrain, args := buildQueryGetCarsConstrain(filter)

	query := fmt.Sprintf("SELECT id, regnum, mark, model, year, name, surname, patronymic FROM cars%s LIMIT $%d OFFSET $%d", queryConstrain, len(args)+1, len(args)+2)
	args = append(args, filter.Limit, filter.Offset)
	//args = append(args, filter.Offset)

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql select Users failed, query: %s", query))
	}

	for rows.Next() {
		car := RepCar{}
		err := rows.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Name, &car.Surname, &car.Patronymic)
		if err != nil {
			return nil, errors.Wrap(err, "sql scan Users failed")
		}
		cars = append(cars, &car)

	}

	return cars, nil
}

// add car
func (r *Rep) AddCar(ctx context.Context, Cars []*RepCar) ([]RepCar, error) {

	s := buildQueryAddCarConstrain(Cars)
	query := fmt.Sprintf("INSERT INTO cars (regnum, mark, model, year, name, surname, patronymic) VALUES %s RETURNING id, regnum, mark, model, year, name, surname, patronymic", s)

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql add Cars failed, query: %s", query))
	}

	defer rows.Close()

	repCars := make([]RepCar, 0, len(Cars))
	car := RepCar{}

	for rows.Next() {
		err := rows.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Name, &car.Surname, &car.Patronymic)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("sql add Cars failed, query: %s", query))
		}
		repCars = append(repCars, car)
	}

	return repCars, nil
}

// delete car
func (r *Rep) DeleteCar(ctx context.Context, id int) (*RepCar, error) {

	car := new(RepCar)

	query := "DELETE FROM cars WHERE id = $1 RETURNING id, regnum, mark, model, year, name, surname, patronymic"
	row := r.DB.QueryRowContext(ctx, query, id)
	if err := row.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Name, &car.Surname, &car.Patronymic); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql delete car failed, query: %s", query))
	}

	return car, nil
}

// update car
func (r *Rep) UpdateCar(ctx context.Context, car *RepCar) (*RepCar, error) {

	repCar := &RepCar{}

	s, args := buildQueryUpdateCarConstrain(car)

	query := fmt.Sprintf("UPDATE cars SET %s RETURNING id, regnum, mark, model, year, name, surname, patronymic", s)

	row := r.DB.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&repCar.Id, &repCar.RegNum, &repCar.Mark, &repCar.Model, &repCar.Year, &repCar.Name, &repCar.Surname, &repCar.Patronymic); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("sql update car failed, query: %s", query))
	}

	return repCar, nil
}

//build query string
func buildQueryAddCarConstrain(cars []*RepCar) string {
	constrains := make([]string, 0, len(cars))
	for _, car := range cars {
		s := fmt.Sprintf("('%s','%s','%s','%d','%s','%s','%s')", car.RegNum, car.Mark, car.Model, car.Year, car.Name, car.Surname, car.Patronymic)
		constrains = append(constrains, s)
	}

	return strings.Join(constrains, ",")
}

func buildQueryGetCarsConstrain(filter *RepQueryFilter) (string, []any) {

	i := 1
	constrains := make([]string, 0, 7)
	args := make([]any, 0, 6)
	if filter.RegNum != "" {
		s := fmt.Sprintf("regnum=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.RegNum)
	}
	if filter.Mark != "" {
		s := fmt.Sprintf("mark=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Mark)
	}
	if filter.Model != "" {
		s := fmt.Sprintf("model=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Model)
	}
	if filter.Year != 0 {
		s := fmt.Sprintf("year=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Year)
	}
	if filter.Name != "" {

		s := fmt.Sprintf("name=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Name)
	}
	if filter.Surname != "" {
		s := fmt.Sprintf("surname=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Surname)
	}
	if filter.Patronymic != "" {
		s := fmt.Sprintf("patronymic=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Patronymic)
	}

	queryConstrain := strings.Join(constrains, " AND ")
	if queryConstrain != "" {

		queryConstrain = fmt.Sprintf(" WHERE %s ORDER BY id ASC", queryConstrain)
	}
	return queryConstrain, args
}

func buildQueryUpdateCarConstrain(filter *RepCar) (string, []any) {

	i := 1
	constrains := make([]string, 0, 7)
	args := make([]any, 0, 6)
	if filter.RegNum != "" {
		s := fmt.Sprintf("regnum=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.RegNum)
	}
	if filter.Mark != "" {
		s := fmt.Sprintf("mark=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Mark)
	}
	if filter.Model != "" {
		s := fmt.Sprintf("model=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Model)
	}
	if filter.Year != 0 {
		s := fmt.Sprintf("year=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Year)
	}
	if filter.Name != "" {

		s := fmt.Sprintf("name=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Name)
	}
	if filter.Surname != "" {
		s := fmt.Sprintf("surname=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Surname)
	}
	if filter.Patronymic != "" {
		s := fmt.Sprintf("patronymic=$%d", i)
		i++

		constrains = append(constrains, s)
		args = append(args, filter.Patronymic)
	}

	queryConstrain := fmt.Sprintf("%s WHERE id=$%d ", strings.Join(constrains, ", "), i)
	args = append(args, filter.Id)

	return queryConstrain, args
}
