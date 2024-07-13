package create

import "github.com/chkda/transaction_service/internal/transaction_service/interfaces/transactions"

type Request struct {
	transactions.Transaction
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
