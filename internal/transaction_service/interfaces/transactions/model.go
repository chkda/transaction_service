package transactions

type Transaction struct {
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"`
	ParentId *int32  `json:"parent_id"`
}
