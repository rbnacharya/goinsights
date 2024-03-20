package stats

import "github.com/rbnacharya/trafficinsights-go/internal/core"

type StatsInput struct {
	CustomerID int
	StartTime  int64
	EndTime    int64
}
type Stats struct {
	Global *core.Global
}
