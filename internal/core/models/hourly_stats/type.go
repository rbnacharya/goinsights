package hourlystats

import "github.com/rbnacharya/trafficinsights-go/internal/core"

type HourlyStats struct {
	RequestCount int `json:"request_count"`
	InvalidCount int `json:"invalid_count"`
	CustomerID   int `json:"customer_id"`
}

type HourlyStatsRepository struct {
	Global *core.Global
}
