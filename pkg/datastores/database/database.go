package database

import "context"

type Writer interface {
	Write(context.Context, string, ...any) error
}

type Reader interface {
	Read(context.Context, string, ...any) (interface{}, error)
}
