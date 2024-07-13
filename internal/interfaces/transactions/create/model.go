package create

import "github.com/chkda/transaction_service/internal/interfaces/transactions"

type Request struct {
	transactions.Transaction
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
