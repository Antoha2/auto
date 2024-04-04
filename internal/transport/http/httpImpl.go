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
	router.GET("/auto/:id", a.getAutoHandler)    //get Auto
	router.GET("/auto/", a.getAutosHandler)      //get auto
	router.POST("/auto/", a.addAutoHandler)      //add Auto
	router.DELETE("/auto/:id", a.delAutoHandler) //del Auto
	router.PUT("/auto/:id", a.updateAutoHandler) //update Auto

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

// get Auto
func (a *apiImpl) getAutoHandler(c *gin.Context) {

	const op = "getAuto"
	log := a.log.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Info("run get Auto by ID", sl.Atr("id", id))

	auto, err := a.service.GetAuto(c, id)
	if err != nil {
		a.log.Error("occurred error for GetAuto", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("get User successfully", sl.Atr("respAuto", auto))

	c.JSON(http.StatusOK, auto)
}

// get Autos
func (a *apiImpl) getAutosHandler(c *gin.Context) {

	const op = "getAutos"
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
	autosQuery := &service.QueryFilter{
		RegNum: q.Get(REGNUM),
		Mark:   q.Get(MARK),
		Model:  q.Get(MODEL),
		Owner:  q.Get(OWNER),
		Offset: offset,
		Limit:  limit,
	}
	log.Info("run get Autos", sl.Atr("filter", autosQuery))

	autos, err := a.service.GetAutos(c, autosQuery)
	if err != nil {
		a.log.Error("occurred error Get Autos", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("get Autos successfully", sl.Atr("respAutos", autos))

	c.JSON(http.StatusOK, autos)
}

// add
func (a *apiImpl) addAutoHandler(c *gin.Context) {

	const op = "addAutos"
	log := a.log.With(slog.String("op", op))

	auto := &service.Auto{}
	if err := c.BindJSON(&auto); err != nil {
		log.Error("cant unmarshall", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Info("run add User", sl.Atr("Auto", auto))

	respAuto, err := a.service.AddAuto(c, auto)
	if err != nil {
		a.log.Error("occurred error for run add Auto", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("add Auto successfully", sl.Atr("respAuto", respAuto))

	c.JSON(http.StatusCreated, respAuto)
}

// del
func (a *apiImpl) delAutoHandler(c *gin.Context) {

	const op = "delAuto"
	log := a.log.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}

	log.Info("run del Auto by ID", sl.Atr("id", id))

	auto, err := a.service.DeleteAuto(c, id)
	if err != nil {
		a.log.Error("occurred error del Auto", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("del Auto successfully", sl.Atr("respAuto", auto))

	c.JSON(http.StatusOK, auto)
}

// update
func (a *apiImpl) updateAutoHandler(c *gin.Context) {

	const op = "updateAuto"
	log := a.log.With(slog.String("op", op))

	auto := service.Auto{}

	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}

	if err := c.BindJSON(&auto); err != nil {
		a.log.Error("cant unmarshall update Auto", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	auto.Id = id

	log.Info("run update user", sl.Atr("User", auto))

	respAuto, err := a.service.UpdateAuto(c, &auto)
	if err != nil {
		a.log.Error("occurred error update Auto", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("update Auto successfully", sl.Atr("respUser", respAuto))

	c.JSON(http.StatusCreated, respAuto)
}
