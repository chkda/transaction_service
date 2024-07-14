package create

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chkda/transaction_service/internal/app"
	"github.com/labstack/echo/v4"
)

const (
	route = "/transactions/:id"
)

var (
	ErrReadingRequest              = errors.New("unable to read reuest")
	ErrMissingTransactionId        = errors.New("transaction id missing in url path")
	ErrTransationIdNotInteger      = errors.New("transaction id not integer")
	ErrTransactionCreationFailue   = errors.New("failed to create transaction")
	ErrTransationIdNegativeInteger = errors.New("transaction id is negative integer")
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
		response.Message = ErrMissingTransactionId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	txnId, err := strconv.Atoi(id)
	if err != nil {
		response.Message = ErrTransationIdNotInteger.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	if txnId < 0 {
		response.Message = ErrTransationIdNegativeInteger.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	// TODO: Add logic
	ctx := e.Request().Context()
	err = c.appHandler.AddTransaction(ctx, int32(txnId), &request.Transaction)
	if err != nil {
		fmt.Println(err)
		response.Status = "failed"
		response.Message = ErrTransactionCreationFailue.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	response.Status = "ok"
	return e.JSON(http.StatusOK, response)
}
