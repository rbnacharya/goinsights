package core

import (
	"database/sql"
)

type Global struct {
	DB *sql.DB
}

func NewGlobal(db *sql.DB) *Global {
	return &Global{DB: db}
}
