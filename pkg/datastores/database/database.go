package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Writer interface {
	Write(context.Context, string, ...any) error
}

type Reader interface {
	Read(context.Context, string, ...any) (*sqlx.Rows, error)
}

type ReaderWriter interface {
	Reader
	Writer
}
