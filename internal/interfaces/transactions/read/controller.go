package read

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
	ErrReadingInputId = errors.New("unable to read input id")
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
	id := e.Param("id")
	response := &Response{}
	if id == "" {
		response.Message = ErrReadingInputId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add Logic
	ctx := e.Request().Context()
	txnId, err := strconv.Atoi(id)
	if err != nil {
		response.Message = ErrReadingInputId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	txn, err := c.appHandler.GetTransaction(ctx, int32(txnId))
	if err != nil {
		response.Message = err.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	response.Transaction = txn
	return e.JSON(http.StatusOK, response)
}
