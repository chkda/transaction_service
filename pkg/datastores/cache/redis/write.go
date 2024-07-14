package rediscache

import (
	"context"

	"github.com/chkda/transaction_service/pkg/datastores/cache"
)

func (c *Conn) Write(ctx context.Context, payload *cache.Payload) error {
	client := c.Client
	err := client.Set(ctx, payload.Key, payload.Value, payload.TTL).Err()
	return err
}
