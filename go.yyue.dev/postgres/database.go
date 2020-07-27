package postgres

import (
	"database/sql"
	"sync"

	// postgres sql
	_ "github.com/lib/pq"
)

// DB is postgres object
type DB struct {
	DB *sql.DB
	tx         *sql.Tx
	mu         sync.Mutex
	maxRetries int // max times a single transaction was retried
}

// Open open sql link
func Open (dataSourceName  string)(*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err!= nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return New(db), nil 
}

// Close closes the database connection.
func (db *DB) Close() error {
	return db.DB.Close()
}


// New return DB struct
func New (db *sql.DB) *DB {
	return &DB{DB:db}
}