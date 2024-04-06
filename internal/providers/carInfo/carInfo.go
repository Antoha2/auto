package provider

import (
	"auto/internal/service"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
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

// type response struct {
// 	cars []*Car
// }

type response struct {
	Id     int    `json:"id"`
	RegNum string `json:"regnum"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Owner  string `json:"owner"`
}

// type response1 struct {
// 	Id int `json:"id"`
// }

// type Car struct {
// 	regNum string `json:"regnum"`
// 	mark   string `json:"mark"`
// 	model  string `json:"model"`
//year   int    `json:"year"`
//Owner  People `json:"owner"`
// }

// type People struct {
// 	name       string `json:"name"`
// 	surname    string `json:"surname"`
// 	patronymic string `json:"patronymic"`
// }

type Nums struct {
	Nums []string `json:"regnums"`
}

func (s *carInfoImpl) GetCarInfo(ctx context.Context, regNums []string) (*service.Car, error) {

	nums := &Nums{Nums: regNums}
	//nums.Nums=regNums
	payload, _ := json.Marshal(nums)
	req, err := http.NewRequest("GET", s.URL, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	car := &response{} //&service.Car{}
	res := &service.Car{}
	// resp, err := http.Get(s.URL)
	// if err != nil {
	// 	return res, errors.Wrap(err, "cant get resp CarsInfo")
	// }

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, errors.Wrap(err, "cant read CarsInfo ")
	}

	defer resp.Body.Close()

	err = json.Unmarshal(body, &car)
	if err != nil {
		return res, errors.Wrap(err, "cant unmarshal CarsInfo ")
	}

	log.Println("@!!!!!!!!!!!!!!!!!22222!!!!!!!!!!!!!", car)

	return res, nil
}

// 	query := fmt.Sprintf("%s?regNum=%s", s.URL, regNums[i])
// 	resp, err := http.Get(query)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "cant get resp Age ")
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "cant read Age ")
// 	}

// 	defer resp.Body.Close()

// 	res := response{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "cant unmarshal Age ")
// 	}

// }
// return nil, nil
// }