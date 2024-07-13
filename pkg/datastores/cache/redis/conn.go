package rediscache

import "github.com/go-redis/redis/v8"

type Conn struct {
	Client *redis.Client
}

func New(config *Config) *Conn {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
	return &Conn{
		Client: client,
	}
}
