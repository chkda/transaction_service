package badgercache

import (
	"context"

	"github.com/chkda/transaction_service/pkg/datastores/cache"
	"github.com/dgraph-io/badger/v4"
)

func (c *Conn) Write(ctx context.Context, payload *cache.Payload) error {
	client := c.Client
	err := client.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(payload.Key), payload.Value).WithTTL(payload.TTL)
		err := txn.SetEntry(entry)
		return err
	})
	return err
}
