package transport

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (a *apiImpl) StartHTTP() error {
	router := gin.Default()
	router.GET("/auto/:id", a.getAutoHandler)    //get Auto
	router.GET("/auto/", a.getAutoHandler)       //get auto
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

func (a *apiImpl) getAutoHandler(c *gin.Context) {
}

// get Autos
func (a *apiImpl) getAutosHandler(c *gin.Context) {
}

// add
func (a *apiImpl) addAutoHandler(c *gin.Context) {
}

// del
func (a *apiImpl) delAutoHandler(c *gin.Context) {
}

// update
func (a *apiImpl) updateAutoHandler(c *gin.Context) {
}
