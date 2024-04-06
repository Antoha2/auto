package transport

import (
	"auto/pkg/logger/sl"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"auto/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (a *apiImpl) StartHTTP() error {
	router := gin.Default()
	router.GET("/CarInfo/:id", a.getCarHandler)    //get Car
	router.GET("/CarInfo/", a.getCarsHandler)      //get Car
	router.POST("/CarInfo/", a.addCarHandler)      //add Car
	router.DELETE("/CarInfo/:id", a.delCarHandler) //del Car
	router.PUT("/CarInfo/:id", a.updateCarHandler) //update Car

	router.GET("/CarInfo/GetCarInfo", a.getCarInfoHandler) //get CarInfo

	err := router.Run(fmt.Sprintf(":%s", a.cfg.HTTP.HostPort))
	if err != nil {
		return errors.Wrap(err, "ocurred error StartHTTP")
	}
	return nil
}

func (a *apiImpl) Stop() {
	if err := a.server.Shutdown(context.TODO()); err != nil {
		panic(errors.Wrap(err, "ocurred error Stop"))
	}
}

// get Car
func (a *apiImpl) getCarHandler(c *gin.Context) {

	const op = "getCar"
	log := a.log.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Info("run get Car by ID", sl.Atr("id", id))

	Car, err := a.service.GetCar(c, id)
	if err != nil {
		a.log.Error("occurred error for GetCar", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("get User successfully", sl.Atr("respCar", Car))

	c.JSON(http.StatusOK, Car)
}

// get Cars
func (a *apiImpl) getCarsHandler(c *gin.Context) {

	const op = "getCars"
	log := a.log.With(slog.String("op", op))

	var err error

	limit := service.DefaultPropertyLimit
	offset := service.DefaultPropertyOffset

	q := c.Request.URL.Query()

	qOffset := q.Get(OFFSET)
	if qOffset != "" {
		offset, err = strconv.Atoi(qOffset)
		if err != nil {
			a.log.Error("offset not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	qLimit := q.Get(LIMIT)
	if qLimit != "" {
		limit, err = strconv.Atoi(qLimit)
		if err != nil {
			a.log.Error("limit not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}
	CarsQuery := &service.QueryFilter{
		RegNum: q.Get(REGNUM),
		Mark:   q.Get(MARK),
		Model:  q.Get(MODEL),
		Owner:  q.Get(OWNER),
		Offset: offset,
		Limit:  limit,
	}
	log.Info("run get Cars", sl.Atr("filter", CarsQuery))

	Cars, err := a.service.GetCars(c, CarsQuery)
	if err != nil {
		a.log.Error("occurred error Get Cars", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("get Cars successfully", sl.Atr("respCars", Cars))

	c.JSON(http.StatusOK, Cars)
}

// add
func (a *apiImpl) addCarHandler(c *gin.Context) {

	const op = "addCars"
	log := a.log.With(slog.String("op", op))

	nums := &service.RegNums{}
	if err := c.BindJSON(&nums); err != nil {
		log.Error("cant unmarshall", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	///////////// CarInfo

	log.Info("run add Cars", sl.Atr("RegNums", nums))

	respCar, err := a.service.AddCar(c, nums)

	if err != nil {
		a.log.Error("occurred error for run add Car", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("add Car successfully", sl.Atr("respCar", respCar))

	c.JSON(http.StatusCreated, respCar)
}

// del
func (a *apiImpl) delCarHandler(c *gin.Context) {

	const op = "delCar"
	log := a.log.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}

	log.Info("run del Car by ID", sl.Atr("id", id))

	Car, err := a.service.DeleteCar(c, id)
	if err != nil {
		a.log.Error("occurred error del Car", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("del Car successfully", sl.Atr("respCar", Car))

	c.JSON(http.StatusOK, Car)
}

// update
func (a *apiImpl) updateCarHandler(c *gin.Context) {

	const op = "updateCar"
	log := a.log.With(slog.String("op", op))

	Car := &service.Car{}

	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}

	if err := c.BindJSON(&Car); err != nil {
		a.log.Error("cant unmarshall update Car", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	Car.Id = id

	log.Info("run update Car", sl.Atr("Car", Car))

	respCar, err := a.service.UpdateCar(c, Car)
	if err != nil {
		a.log.Error("occurred error update Car", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("update Car successfully", sl.Atr("respCar", respCar))

	c.JSON(http.StatusCreated, respCar)
}

// get CarInfo
func (a *apiImpl) getCarInfoHandler(c *gin.Context) {
	fmt.Println("44444444 ")
	const op = "getCarInfo"
	log := a.log.With(slog.String("op", op))
	nums := &service.RegNums{}
	if err := c.BindJSON(&nums); err != nil {
		log.Error("cant unmarshall", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//cars:=[]
	car := service.Car{}
	cars := []service.Car{}
	//cars := make(car, len(nums.Nums))

	for i := 0; i < len(nums.Nums); i++ {
		car.Mark = string(i)
		car.Model = string(i)
		car.Owner = string(i)
		car.RegNum = nums.Nums[i]
		cars = append(cars, car)
	}

	fmt.Println("44444444 ")

	// car := &service.Car{
	// 	RegNum: "!!!!!!!!!!!!!!!", //nums.Nums[1],
	// 	Mark:   "mark",
	// 	Model:  "model",
	// 	Owner:  "owner",
	// }

	c.JSON(http.StatusOK, cars)
}
