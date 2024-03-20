package controller

import (
	"github.com/labstack/echo/v4"
	hourlystats "github.com/rbnacharya/trafficinsights-go/internal/core/models/hourly_stats"
	"github.com/rbnacharya/trafficinsights-go/internal/core/stats"
	"github.com/rbnacharya/trafficinsights-go/internal/errors"
)


func (ctrl Controllers) FetchStats(ec echo.Context) error {
	// Fetch stats
	hourlyStatsRepository := hourlystats.HourlyStatsRepository{Global: ctrl.Global}

	statsCore := stats.Stats{
		Global: ctrl.Global,
	}

	// Validate input
	statsInput, booms := statsCore.ValidateInput(ec)

	// Check for errors
	if booms != nil && booms.HasBooms() {
		errors.HandleError(booms, ec)
		return nil
	}

	// Fetch hourly stats
	stats, err2 := hourlyStatsRepository.FindHourlyStatsByCustomerID(statsInput.CustomerID, statsInput.StartTime, statsInput.EndTime)

	if err2 != nil {
		booms := errors.NewBooms()
		booms.AddBoom(errors.NewBoom("DatabaseError", "Error fetching hourly stats", err2))
		errors.HandleError(booms, ec)
		return nil
	}

	if stats == nil {
		return ec.JSON(404, ResponseWithMessage{Message: "No stats found"})
	}

	return ec.JSON(200, ResponseWithMessage{Message: "Success", Data: stats})
}
