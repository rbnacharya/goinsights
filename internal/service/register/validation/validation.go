package validation

import (
	"github.com/labstack/echo/v4"
	"github.com/rbnacharya/trafficinsights-go/internal/service/register"
)

type RegisterValidation struct {
	CustomerIdInvalid bool
	ValidationError   bool
	BindingError      bool
	IpBlacklisted     bool
	UaBlacklisted     bool
	InvalidTimestamp  bool
}

func (ctrl Validate) ValidateInput(ec echo.Context) (*RegisterValidation, *register.RegisterClickInput) {
	input := new(register.RegisterClickInput)
	validation := new(RegisterValidation)

	validation.ValidationError = false
	validation.BindingError = false
	validation.CustomerIdInvalid = false

	if err := ec.Bind(input); err != nil {
		validation.BindingError = true
		validation.ValidationError = true
		return validation, nil
	}

	if err := ec.Validate(input); err != nil {
		validation.BindingError = false
		validation.ValidationError = true
		return validation, input
	}

	if input.CustomerID == 0 {
		validation.CustomerIdInvalid = true
		validation.ValidationError = true
	} else {
		// Check if customer exists
		if !ctrl.CheckCustomerExists(input.CustomerID) {
			validation.CustomerIdInvalid = true
			validation.ValidationError = true
		}
	}
	if !validation.ValidationError {
		// Check if IP address is in blacklist
		validation.IpBlacklisted = ctrl.ValidateIPBlacklist(input.RemoteIP)
		// Check if User-Agent is in blacklist
		validation.UaBlacklisted = ctrl.ValidateUserAgent(input.UserAgent)

		if validation.IpBlacklisted || validation.UaBlacklisted {
			validation.ValidationError = true
		}
	}

	if !validation.ValidationError {
		// Check timestamp is valid
		if !ctrl.ValidateTimestamp(input.Timestamp) {
			validation.ValidationError = true
			validation.InvalidTimestamp = true
		}
	}

	return validation, input
}
