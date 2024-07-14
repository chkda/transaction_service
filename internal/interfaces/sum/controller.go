package sum

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/chkda/transaction_service/internal/app"
	"github.com/labstack/echo/v4"
)

const (
	route = "/sum/:id"
)

var (
	ErrMissingTransactionId        = errors.New("transaction id missing in url path")
	ErrTransationIdNotInteger      = errors.New("transaction id not integer")
	ErrTransationIdNegativeInteger = errors.New("transaction id is negative integer")
	ErrFetchingSumForTransactionId = errors.New("failed to fetch sum")
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
		response.Message = ErrMissingTransactionId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}

	txnId, err := strconv.Atoi(parentId)
	if err != nil {
		response.Message = ErrTransationIdNotInteger.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	if txnId < 0 {
		response.Message = ErrTransationIdNegativeInteger.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	ctx := e.Request().Context()
	sum, err := c.appHandler.GetSumForTxnId(ctx, int32(txnId))
	if err != nil {
		log.Println("[ERROR]:interfaces:sum:", err.Error())
		response.Message = ErrFetchingSumForTransactionId.Error()
		return e.JSON(http.StatusBadRequest, response)
	}
	response.Sum = &sum
	return e.JSON(http.StatusOK, response)
}
