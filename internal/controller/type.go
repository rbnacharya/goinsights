package controller

import "github.com/rbnacharya/trafficinsights-go/internal/core"

type Controllers struct {
	Global *core.Global
}

type ResponseWithMessage struct {
	Message string `json:"message"`
	success bool
	Data    interface{} `json:"data"`
}
