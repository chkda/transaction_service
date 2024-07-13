package read

import "github.com/chkda/transaction_service/internal/transaction_service/interfaces/transactions"

type Response struct {
	transactions.Transaction
	Message string `json:"message,omitempty"`
}
