package db

import (
	"database/sql"
	"fmt"
)

type Client struct {
	*sql.DB
}

func NewClient(driver string, name string) (client *Client, err error) {
	switch driver {
	case "sqlite":
		client, err = NewSQLite(name)
	case "postgres":
		client, err = NewPostgres(name)
	default:
		err = fmt.Errorf("Definition Error: driver")
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c Client) Close() {
	c.Close()
}
