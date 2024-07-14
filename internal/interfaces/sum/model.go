package sum

type Response struct {
	Sum     *float64 `json:"sum,omitempty"`
	Message string   `json:"message,omitempty"`
}
