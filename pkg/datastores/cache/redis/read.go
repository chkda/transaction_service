package rediscache

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

var (
	ErrKeyDoesNotExist = errors.New("key does not exist")
)

func (c *Conn) Read(ctx context.Context, key string) ([]byte, error) {
	client := c.Client
	value, err := client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, ErrKeyDoesNotExist
	}

	if err != nil {
		return nil, err
	}

	return value, nil

}
