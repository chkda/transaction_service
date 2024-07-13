package create

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/chkda/transaction_service/internal/app"
	"github.com/labstack/echo/v4"
)

const (
	route = "/transactions/:id"
)

var (
	ErrReadingRequest = errors.New("unable to read reuest")
)

type Controller struct {
	appHandler *app.Handler
}

func New(appHandler *app.Handler) *Controller {
	return &Controller{
		appHandler: appHandler,
	}
}

func (c *Controller) GetRoute() string {
	return route
}

func (c *Controller) Handler(e echo.Context) error {
	request := &Request{}
	response := &Response{}
	id := e.Param("id")
	err := e.Bind(request)
	if err != nil {
		response.Status = "failed"
		response.Message = ErrReadingRequest.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	if id == "" {
		response.Status = "failed"
		response.Message = ErrReadingRequest.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	txnId, err := strconv.Atoi(id)
	if err != nil {
		response.Message = ErrReadingRequest.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add logic
	ctx := e.Request().Context()
	err = c.appHandler.AddTransaction(ctx, int32(txnId), &request.Transaction)
	if err != nil {
		response.Status = "failed"
		response.Message = ErrReadingRequest.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	response.Status = "ok"
	return e.JSON(http.StatusOK, response)
}
