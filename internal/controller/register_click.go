package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	hourlystats "github.com/rbnacharya/trafficinsights-go/internal/core/models/hourly_stats"
	"github.com/rbnacharya/trafficinsights-go/internal/errors"
	"github.com/rbnacharya/trafficinsights-go/internal/service/register"
)

func (ctrl Controllers) RegisterClick(ec echo.Context) error {
	// Validate input

	register := register.Register{
		Global: ctrl.Global,
	}

	hourlyStatsRepo := hourlystats.HourlyStatsRepository{
		Global: ctrl.Global,
	}

	validation, input := register.ValidateInput(ec)

	if validation != nil && validation.ValidationError {
		booms := errors.NewBooms()
		if validation.BindingError {
			booms.AddBoom(errors.NewBoom("InvalidInput", "Invalid input", nil))
		} else if validation.CustomerIdInvalid {
			booms.AddBoom(errors.NewBoom("InvalidInput", "Customer not found", nil))
		} else {
			if validation.IpBlacklisted {
				booms.AddBoom(errors.NewBoom("InvalidInput", "IP address is blacklisted", nil))
			}
			if validation.UaBlacklisted {
				booms.AddBoom(errors.NewBoom("InvalidInput", "User-Agent is blacklisted", nil))
			}

			if !validation.IpBlacklisted && !validation.UaBlacklisted {
				booms.AddBoom(errors.NewBoom("InvalidInput", "Invalid input", nil))
			}
			// booms.AddBoom(errors.NewBoom("InvalidInput", "Invalid input", nil))
			hourlyStatsRepo.AddEntry(input.CustomerID, input.Timestamp, true)
		}

		errors.HandleError(booms, ec)
		return nil
	}

	// Add entry to hourly stats
	err := hourlyStatsRepo.AddEntry(input.CustomerID, input.Timestamp, false)

	if err != nil {
		booms := errors.NewBooms()
		booms.AddBoom(errors.NewBoom("DatabaseError", "Error adding entry to hourly stats", err))
		errors.HandleError(booms, ec)
		return nil
	}

	return ec.JSON(http.StatusCreated, input)
}
