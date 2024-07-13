package sum

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/chkda/transaction_service/internal/transaction_service/app"
	"github.com/labstack/echo/v4"
)

const (
	route = "/sum/:id"
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
	parentId := e.Param("id")
	response := &Response{}
	if parentId == "" {
		response.Message = ErrReadingInputId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add logic
	txnId, err := strconv.Atoi(parentId)
	if err != nil {
		response.Message = ErrReadingInputId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	ctx := e.Request().Context()
	sum, err := c.appHandler.GetSumForTxnId(ctx, int32(txnId))
	if err != nil {
		response.Message = err.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	response.Sum = &sum
	return e.JSON(http.StatusOK, response)
}
