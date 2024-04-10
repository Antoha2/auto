package provider

import (
	"auto/internal/service"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type carInfoImpl struct {
	URL string
}

func NewGetCarInfo(URL string) *carInfoImpl {
	return &carInfoImpl{
		URL: URL,
	}
}

type Car struct {
	Id     int    `json:"id"`
	RegNum string `json:"regnum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Owner  string `json:"owner"`
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type RegNums struct {
	Nums []string `json:"regnums"`
}

// @BasePath  /
// @host asdasdasd

// @Summary get additional car info from external sources
// @Schemes
// @Description get additional car info from external sources
// @Accept json
// @Produce json
// @Param regNums body provider.RegNums true "parameters of the record being updated"
// @Success 200 {object} []provider.Car "received records"
// @Failure 400  400  {object}  httputil.HTTPError
// @Failure 500   500  {object}  httputil.HTTPError

func (s *carInfoImpl) GetCarInfo(ctx context.Context, regNums []string) ([]*service.Car, error) {

	nums := &RegNums{Nums: regNums}

	payload, _ := json.Marshal(nums)

	req, err := http.NewRequest("GET", s.URL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "cant wraps NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "cant get the requested data")
	}
	defer resp.Body.Close()

	cars := []*service.Car{}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "cant read CarsInfo ")
	}

	err = json.Unmarshal(body, &cars)
	if err != nil {
		return nil, errors.Wrap(err, "cant unmarshal CarsInfo ")
	}

	return cars, nil
}
