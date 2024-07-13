package transactions

type Transaction struct {
	Amount   string  `json:"amount"`
	Type     string  `json:"type"`
	ParentId float64 `json:"parent_id"`
}
