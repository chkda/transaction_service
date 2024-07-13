package badgercache

import (
	"context"

	"github.com/dgraph-io/badger/v4"
)

func (c *Conn) Read(ctx context.Context, key string) ([]byte, error) {
	var value []byte
	client := c.Client
	err := client.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	return value, err
}
