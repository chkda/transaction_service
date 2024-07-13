package types

import (
	"errors"
	"net/http"

	"github.com/chkda/transaction_service/internal/transaction_service/app"
	"github.com/labstack/echo/v4"
)

const (
	route = "/types/:type"
)

var (
	ErrReadingInputType = errors.New("unable to read input type")
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
	inpType := e.Param("type")
	response := &Response{}
	if inpType == "" {
		response.Message = ErrReadingInputType.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add logic
	ctx := e.Request().Context()
	ids, err := c.appHandler.GetTransactionsWithSameType(ctx, inpType)
	if err != nil {
		response.Message = err.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	response.Ids = ids
	return e.JSON(http.StatusOK, response)
}
