package register

import (
	"github.com/rbnacharya/trafficinsights-go/internal/core/models/customer"
)

func (ctrl Register) CheckCustomerExists(customerID int) bool {
	// SQL query to check if customer exists

	customerRepo := customer.CustomerRepository{
		Global: ctrl.Global,
	}

	return customerRepo.FindCustomerByID(customerID) != nil
}
