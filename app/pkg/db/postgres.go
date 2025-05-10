package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgres(dataSource string) (*Client, error) {
	if dataSource == "" {
		return nil, fmt.Errorf("Definition Error: Datasource")
	}

	db, err := sql.Open("pgx", dataSource)

	if err != nil {
		return nil, err
	}

	return &Client{db}, nil
}
