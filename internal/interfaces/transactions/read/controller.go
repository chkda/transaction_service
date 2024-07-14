package read

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/chkda/transaction_service/internal/app"
	"github.com/labstack/echo/v4"
)

const (
	route = "/transactions/:id"
)

var (
	ErrMissingTransactionId        = errors.New("transaction id missing in url path")
	ErrTransactionReadFailue       = errors.New("failed to read transaction")
	ErrTransationIdNotInteger      = errors.New("transaction id not integer")
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
	id := e.Param("id")
	response := &Response{}
	if id == "" {
		response.Message = ErrMissingTransactionId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}

	ctx := e.Request().Context()
	txnId, err := strconv.Atoi(id)
	if err != nil {
		response.Message = ErrTransationIdNotInteger.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	if txnId < 0 {
		response.Message = ErrTransationIdNegativeInteger.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	txn, err := c.appHandler.GetTransaction(ctx, int32(txnId))
	if err != nil {
		log.Println("[ERROR]:interfaces:transactions:read:", err.Error())
		response.Message = ErrTransactionReadFailue.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	response.Transaction = txn
	return e.JSON(http.StatusOK, response)
}
