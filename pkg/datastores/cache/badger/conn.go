package badgercache

import "github.com/dgraph-io/badger/v4"

type Conn struct {
	Client *badger.DB
}

func New() (*Conn, error) {
	client, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		return nil, err
	}
	return &Conn{
		Client: client,
	}, nil
}
