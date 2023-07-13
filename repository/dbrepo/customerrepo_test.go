package dbrepo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	repo := &sqliteDBRepo{DB: db}

	customerID := "123"
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"

	mock.ExpectExec("INSERT INTO customers").
		WithArgs(customerID, firstName, lastName, email).
		WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(nil)

	err = repo.AddCustomer(customerID, firstName, lastName, email)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestDeleteCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	repo := &sqliteDBRepo{DB: db}

	customerID := "123"

	mock.ExpectExec("DELETE FROM customers").
		WithArgs(customerID).
		WillReturnResult(sqlmock.NewResult(1, 1)).
		WillReturnError(nil)

	err = repo.DeleteCustomer(customerID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
