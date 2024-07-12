package Metrics

type Writer interface {
	IncCounter(string, float64)
	Summary(string, float64)
}
