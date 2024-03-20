package customer

import "github.com/rbnacharya/goinsights/internal/core"

type CustomerRepository struct {
	Global *core.Global
}

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
