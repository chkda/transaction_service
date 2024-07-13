package cache

import (
	"context"
	"time"
)

type Payload struct {
	Key   string
	Value []byte
	TTL   time.Duration
}

type Writer interface {
	Write(context.Context, *Payload) error
}

type Reader interface {
	Read(context.Context, string) ([]byte, error)
}
