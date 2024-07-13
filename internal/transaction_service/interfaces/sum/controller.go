package sum

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	route = "/sum/:parent_id"
)

var (
	ErrReadingInputId = errors.New("unable to read input id")
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
	parentId := e.Param("parent_id")
	response := &Response{}
	if parentId == "" {
		response.Message = ErrReadingInputId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add logic
	return e.JSON(http.StatusOK, response)
}
