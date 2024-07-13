package types

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	route = "/types/:type"
)

var (
	ErrReadingInputType = errors.New("unable to read input type")
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetRoute() string {
	return route
}

func (c *Controller) Handler(e echo.Context) error {
	inpType := e.Param("type")
	response := &Response{}
	if inpType == "" {
		response.Message = ErrReadingInputType.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add logic
	return e.JSON(http.StatusOK, response)
}
