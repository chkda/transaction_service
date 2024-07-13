package read

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	route = "/transactions/"
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
	txnId := e.Param("txn_id")
	response := &Response{}
	if txnId == "" {
		response.Message = ErrReadingInputId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add Logic
	return e.JSON(http.StatusOK, response)
}
