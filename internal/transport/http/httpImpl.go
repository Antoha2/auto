package transport

import (
	"auto/pkg/logger/sl"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"unicode/utf8"

	"auto/internal/service"

	_ "auto/cmd/docs"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Car info
// @version 0.0.1

// @contact.name   Lebedev A.S.
// @contact.email  9112441775@mail.ru

//@title Car Info

// @BasePath  /info/
// @host http://127.0.0.1:80

// @Description  для получения данных из внешного источника необходимо изменить значение  переменной URL_GETCARINFO в .env

func (a *apiImpl) StartHTTP() error {

	router := gin.Default()
	methods := router.Group("/info/")
	{
		methods.GET(":id", a.GetCarHandler)    //get Car
		methods.GET("", a.getCarsHandler)      //get Car
		methods.POST("", a.addCarHandler)      //add Car
		methods.DELETE(":id", a.delCarHandler) //del Car
		methods.PUT(":id", a.updateCarHandler) //update Car
	}
	router.GET("/info/get_carinfo", a.getCarInfoHandler) // external data source emulator

	URLSwagger := ginSwagger.URL("http://127.0.0.1:80/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, URLSwagger))

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

// GetCarHandler godoc
// @Summary get car info by id from the database
// @Description  get car info by id from the database
// @Tags         methods
// @Accept json
// @Produce json
// @Param  id     query    int     true        "ID"
// @Success 200 {object} service.Car "read record"
// @Failure 400  400  {object}  httputil.HTTPError
// @Failure 500   500  {object}  httputil.HTTPError
// @Router /info/:id [get]
func (a *apiImpl) GetCarHandler(c *gin.Context) {

	const op = "getCar"
	log := a.log.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param(a.cfg.ApiConst.ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Info("run get Car by ID", sl.Atr("id", id))

	car, err := a.service.GetCar(c, id)
	if err != nil {
		a.log.Error("occurred error for GetCar", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if car.Id == 0 {
		log.Info("get Car successfully", sl.Atr("respCar", car))
		c.JSON(http.StatusNotFound, "record with this id not found")
	} else {

		log.Info("get Car successfully", sl.Atr("respCar", car))
		c.JSON(http.StatusOK, car)
	}
}

// GetCarsHandler godoc
// @Summary get Cars info from the database
// @Tags         methods
// @Description get Cars info from the database with a search filter
// @Accept json
// @Produce json
// @Param regNum query string false "regNum"
// @Param mark query string false "mark"
// @Param year query int false "year"
// @Param name query string false "owner.name"
// @Param surname query string false "owner.surname"
// @Param patronymic query string false "owner.patronymic"
// @Success 200 {object} []service.Car "read records"
// @Failure 400  400  {object}  httputil.HTTPError
// @Failure 500   500  {object}  httputil.HTTPError
// @Router /info/ [get]
func (a *apiImpl) getCarsHandler(c *gin.Context) {

	const op = "getCars"
	log := a.log.With(slog.String("op", op))

	var err error

	year := 0
	limit := a.cfg.ServiceConst.DefaultPropertyLimit
	offset := a.cfg.ServiceConst.DefaultPropertyOffset
	q := c.Request.URL.Query()

	qOffset := q.Get(a.cfg.ApiConst.OFFSET)
	if qOffset != "" {
		offset, err = strconv.Atoi(qOffset)
		if err != nil {
			a.log.Error("offset not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	qLimit := q.Get(a.cfg.ApiConst.LIMIT)
	if qLimit != "" {
		limit, err = strconv.Atoi(qLimit)
		if err != nil {
			a.log.Error("limit not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	qYear := q.Get(a.cfg.ApiConst.YEAR)
	if qYear != "" {
		year, err = strconv.Atoi(qYear)
		if err != nil {
			a.log.Error("year not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	CarsQuery := &service.QueryFilter{}

	CarsQuery.RegNum = q.Get(a.cfg.ApiConst.REGNUM)
	CarsQuery.Mark = q.Get(a.cfg.ApiConst.MARK)
	CarsQuery.Model = q.Get(a.cfg.ApiConst.MODEL)
	CarsQuery.Owner.Name = q.Get(a.cfg.ApiConst.NAME)
	CarsQuery.Owner.Surname = q.Get(a.cfg.ApiConst.SURNAME)
	CarsQuery.Owner.Patronymic = q.Get(a.cfg.ApiConst.PATRONYMIC)
	CarsQuery.Year = year
	CarsQuery.Limit = limit
	CarsQuery.Offset = offset

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

// addCarHandler godoc
// @Summary add car info to to database
// @Tags         methods
// @Description add car info to database
// @Accept json
// @Produce json
// @Param regNums body service.RegNums true "slice reg numbers"
// @Success 201 {object} []service.Car "added records"
// @Failure 400  400  {object}  httputil.HTTPError
// @Failure 500   500  {object}  httputil.HTTPError
// @Router /info/ [post]
func (a *apiImpl) addCarHandler(c *gin.Context) {

	const op = "addCars"
	log := a.log.With(slog.String("op", op))

	nums := &service.RegNums{}

	if err := c.BindJSON(&nums); err != nil {
		log.Error("cant unmarshall", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if !inputValidation(nums.Nums) {
		err := errors.New("incorrect input data format")
		log.Error("occurred error for run add Car", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

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

// delCarHandler godoc
// @Summary del car info from the database
// @Tags         methods
// @Description del car info from the database
// @Accept json
// @Produce json
// @Param id query int true "ID delete car"
// @Success 200 {object} service.Car "deleted record"
// @Failure 400  400  {object}  httputil.HTTPError
// @Failure 500   500  {object}  httputil.HTTPError
// @Router /info/:id [delete]
func (a *apiImpl) delCarHandler(c *gin.Context) {

	const op = "delCar"
	log := a.log.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param(a.cfg.ApiConst.ID))
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

// updateCarHandler godoc
// @Summary update car info in the database
// @Tags methods
// @Description update car info in the database
// @Accept json
// @Produce json
// @Param car body service.Car true "parameters of the record being updated"
// @Success 201 {object} service.Car "updated record"
// @Failure 400  400  {object}  httputil.HTTPError
// @Failure 500   500  {object}  httputil.HTTPError
// @Router /info/:id [put]
func (a *apiImpl) updateCarHandler(c *gin.Context) {

	const op = "updateCar"
	log := a.log.With(slog.String("op", op))

	car := &service.Car{}

	id, err := strconv.Atoi(c.Param(a.cfg.ApiConst.ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}

	if err := c.BindJSON(&car); err != nil {
		a.log.Error("cant unmarshall update Car", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if !updateDataCheck(car) {
		err := errors.New("No update data")
		a.log.Error("occurred error update Car", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return

	}

	car.Id = id

	log.Info("run update Car", sl.Atr("Car", car))

	respCar, err := a.service.UpdateCar(c, car)
	if err != nil {
		a.log.Error("occurred error update Car", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("update Car successfully", sl.Atr("respCar", respCar))

	c.JSON(http.StatusCreated, respCar)
}

// external data source emulator
func (a *apiImpl) getCarInfoHandler(c *gin.Context) {

	const op = "getCarInfo"
	log := a.log.With(slog.String("op", op))

	nums := &service.RegNums{}
	if err := c.BindJSON(&nums); err != nil {
		log.Error("cant unmarshall", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cars := []*service.Car{}

	for i := 0; i < len(nums.Nums); i++ {

		car := &service.Car{}
		car.Mark = strconv.Itoa(i + 3)
		car.Model = strconv.Itoa(i * 4)
		car.Year = i + 2000
		car.Owner.Name = strconv.Itoa(i + 3)
		car.Owner.Surname = strconv.Itoa(i + 3)
		car.Owner.Patronymic = strconv.Itoa(i + 3)
		car.RegNum = nums.Nums[i]
		cars = append(cars, car)
	}

	c.JSON(http.StatusOK, cars)
}

//checking the correctness of the entered data
func inputValidation(s []string) bool {

	for i := 0; i < len(s); i++ {
		if !checkRegNum(s[i]) {
			return false
		}
	}
	return true
}

func checkRegNum(str string) bool {

	if utf8.RuneCountInString(str) != 8 && utf8.RuneCountInString(str) != 9 {
		return false
	}

	s := []string{"а", "А", "в", "В", "е", "Е", "к", "К", "м", "М", "н", "Н", "о", "О", "р", "Р", "с", "С", "т", "Т", "у", "У", "х", "Х"}
	num := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	rune := []rune(str)

	for i := 0; i < len(rune); i++ {

		c := fmt.Sprintf("%c", rune[i])

		if i == 0 || i == 4 || i == 5 {
			if !checkSymbol(c, s) {
				return false
			}
		} else {
			if !checkSymbol(c, num) {
				return false
			}
		}
	}
	return true
}

func checkSymbol(c string, mas []string) bool {

	for _, s := range mas {
		if c == s {
			return true
		}
	}
	return false
}

//data availability check
func updateDataCheck(car *service.Car) bool {

	if car.RegNum == "" && car.Mark == "" && car.Model == "" && car.Year == 0 && car.Owner.Name == "" && car.Owner.Surname == "" && car.Owner.Patronymic == "" {
		return false
	}
	return true
}
