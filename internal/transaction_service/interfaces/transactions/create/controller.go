package create

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	route = "/transactions/"
)

var (
	ErrReadingRequest = errors.New("unable to read reuest")
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
	request := &Request{}
	response := &Response{}
	err := e.Bind(request)
	if err != nil {
		response.Status = "failed"
		response.Message = ErrReadingRequest.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add logic

	response.Status = "ok"
	return e.JSON(http.StatusOK, response)
}