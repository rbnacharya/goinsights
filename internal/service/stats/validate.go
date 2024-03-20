package stats

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rbnacharya/trafficinsights-go/internal/core/models/customer"
	"github.com/rbnacharya/trafficinsights-go/internal/errors"
)

func validateAndGenerateDates(dateStr string) (int64, int64, *errors.Booms) {
	booms := errors.NewBooms()

	// Validate date
	if dateStr == "" {
		booms.AddBoom(errors.NewBoom("Invalid Input", "date is required", ""))
		return 0, 0, booms
	} else {
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			println(err.Error())
			booms.AddBoom(errors.NewBoom("Invalid Input", "Invalid date format", ""))
			return 0, 0, booms
		} else {
			startTime := parsedDate.Unix()
			endTime := startTime + 24*60*60

			return startTime, endTime, nil
		}

	}
}

func (ctrl Stats) ValidateInput(ec echo.Context) (*StatsInput, *errors.Booms) {

	customerId := ec.QueryParam("customerID")
	date := ec.QueryParam("date")

	booms := errors.NewBooms()

	// Validate customerID and date
	if customerId == "" {
		booms.AddBoom(errors.NewBoom("Invalid Input", "customerID is required", ""))
	} else if date == "" {
		booms.AddBoom(errors.NewBoom("Invalid Input", "date is required", ""))
	} else {
	}

	customerIdInt, err := strconv.Atoi(customerId)

	// Check if customerId is integer
	if err != nil && customerId != "" {
		booms.AddBoom(errors.NewBoom("Invalid Input", "customerID should be an integer", ""))
	} else if customerId != "" {
		customerRepo := customer.CustomerRepository{
			Global: ctrl.Global,
		}
		// Check if customer exists
		if customerRepo.FindCustomerByID(customerIdInt) == nil {
			booms.AddBoom(errors.NewBoom("Invalid Input", "Customer not found", ""))
		}
	}

	if !booms.HasBooms() {
		startTime, endTime, booms := validateAndGenerateDates(date)
		if booms != nil {
			return nil, booms
		}

		return &StatsInput{
			CustomerID: customerIdInt,
			StartTime:  startTime,
			EndTime:    endTime,
		}, nil
	} else {
		return nil, booms
	}
}
