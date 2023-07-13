package main

import (
	"Erply-api-test-project/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Repo.Home)
	mux.HandleFunc("/signin", handlers.Repo.SignIn)
	mux.HandleFunc("/signout", handlers.Repo.SignOut)
	mux.HandleFunc("/savecustomer", handlers.Repo.SaveCustomer)
	mux.HandleFunc("/deletecustomer", handlers.Repo.DeleteCustomer)
	return mux
}
