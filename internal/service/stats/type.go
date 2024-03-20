package stats

import "github.com/rbnacharya/goinsights/internal/core"

type StatsInput struct {
	CustomerID int
	StartTime  int64
	EndTime    int64
}
type Stats struct {
	Global *core.Global
}
