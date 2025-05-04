package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLite(dataSource string) (*Client, error) {
	if dataSource == "" {
		return nil, fmt.Errorf("Definition Error: Datasource")
	}

	db, err := sql.Open("sqlite3", dataSource)

	if err != nil {
		return nil, err
	}

	return &Client{db}, nil
}
