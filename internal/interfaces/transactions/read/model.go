package read

import "github.com/chkda/transaction_service/internal/interfaces/transactions"

type Response struct {
	*transactions.Transaction
	Message string `json:"message,omitempty"`
}
