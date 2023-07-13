package repository

import (
	"Erply-api-test-project/models"
)

type DatabaseRepo interface {
	DeleteCustomer(customerID string) error
	GetCustomers() ([]models.Customer, error)
	AddCustomer(customerID, firstName, lastName, email string) error
}
