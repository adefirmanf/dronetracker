package repository

import "errors"

var (
	ErrInsert = errors.New("sql: insert statement returned error")
)
