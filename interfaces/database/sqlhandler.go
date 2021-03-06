package database

import "database/sql"

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
	Begin() (Tx, error)
	Transaction(txFunc func(*sql.Tx) error) error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type Tx interface {
	Rollback() error
	Commit() error
	Exec(string, ...interface{}) (Result, error)
}
