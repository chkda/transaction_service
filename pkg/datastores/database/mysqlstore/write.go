package mysqlstore

import "context"

func (c *Conn) Write(ctx context.Context, query string, args ...any) error {
	client := c.Client
	_, err := client.ExecContext(ctx, query, args...)
	return err
}
