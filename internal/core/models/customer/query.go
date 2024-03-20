package customer

func (ctrl CustomerRepository) FindCustomerByID(customerID int) *Customer {
	db := ctrl.Global.DB
	customer := new(Customer)
	query := "SELECT * FROM customer WHERE id = $1"
	row := db.QueryRow(query, customerID)
	if err := row.Scan(&customer.ID, &customer.Name, &customer.Active); err != nil {
		println(err.Error())
		return nil
	}
	return customer
}
