package errors

import (
	"fmt"
)

const (
	CustomerNotFound = "Customer not found"
	InvalidInput     = "Invalid input"
)

type Boom struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func NewBoom(code string, message string, details interface{}) *Boom {
	return &Boom{
		Code:    code,
		Message: message,
		Details: details,
	}
}

type Booms struct {
	Errors []Boom `json:"errors"`

	HasBooms func() bool `json:"-"`
}

func (b *Booms) AddBoom(boom *Boom) {
	b.Errors = append(b.Errors, *boom)
	b.HasBooms = func() bool {
		return len(b.Errors) > 0
	}
}

func NewBooms() *Booms {
	booms := &Booms{}
	booms.HasBooms = func() bool {
		return len(booms.Errors) > 0
	}
	return booms
}

type RequestError struct {
	Message    string
	StatusCode int
}

func (iie *RequestError) Error() string {
	return fmt.Sprintf("Invalid input")
}
