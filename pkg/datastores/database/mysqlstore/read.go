package mysqlstore

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func (c *Conn) Read(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	client := c.Client
	rows, err := client.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}
