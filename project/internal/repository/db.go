package repository

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func NewPostgresDB(connStr string) (*sql.DB, error) {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to open db: %w", err)
    }
    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping db: %w", err)
    }
    return db, nil
}
