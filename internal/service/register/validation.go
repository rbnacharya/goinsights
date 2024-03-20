package register

import (
	"github.com/labstack/echo/v4"
)

type RegisterValidation struct {
	CustomerIdInvalid bool
	ValidationError   bool
	BindingError      bool
	IpBlacklisted     bool
	UaBlacklisted     bool
}

func (ctrl Register) ValidateInput(ec echo.Context) (*RegisterValidation, *RegisterClickInput) {
	input := new(RegisterClickInput)
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

	return validation, input
}
