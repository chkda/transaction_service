package cache

import (
	"context"
	"time"
)

type Payload struct {
	Key   string
	Value interface{}
	TTL   time.Duration
}

type Writer interface {
	Write(context.Context, *Payload) error
}

type Reader interface {
	Read(context.Context, string) error
}
