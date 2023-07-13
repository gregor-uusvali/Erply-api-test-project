package dbrepo

import (
	"Erply-api-test-project/models"
)

func (m *sqliteDBRepo) AddCustomer(customerID, firstName, lastName, email string) error {
	_, err := m.DB.Exec("INSERT INTO customers (customer_id, first_name, last_name, email) VALUES (?, ?, ?, ?)",
		customerID, firstName, lastName, email)
	if err != nil {
		return err
	}

	return nil
}


func (m *sqliteDBRepo) GetCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	rows, err := m.DB.Query("SELECT * FROM customers")
	if err != nil {
		return customers, err
	}

	defer rows.Close()
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(
			&customer.ID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
		)
		if err != nil {
			return customers, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (m *sqliteDBRepo) DeleteCustomer(customerID string) error {
	_, err := m.DB.Exec("DELETE FROM customers WHERE customer_id = ?", customerID)
	if err != nil {
		return err
	}

	return nil
}
