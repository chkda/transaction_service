package mysqlstore

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Conn struct {
	Client *sqlx.DB
}

func New(config *Config) (*Conn, error) {
	sqlConfig := &mysql.Config{
		Net:    "tcp",
		Addr:   config.Host,
		DBName: config.DB,
		User:   config.Username,
		Passwd: config.Password,
	}
	client, err := sqlx.Open("mysql", sqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &Conn{
		Client: client,
	}, nil
}
