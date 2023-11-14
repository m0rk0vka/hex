package db

import (
	"database/sql"
	"log"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
		return nil, err
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failur: %v", err)
	}

	return &Adapter{db: db}, nil
}
