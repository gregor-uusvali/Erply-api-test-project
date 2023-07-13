package main

import (
	"Erply-api-test-project/handlers"
	"net/http"

	_ "github.com/swaggo/http-swagger"
)

func routes() http.Handler {
	mux := http.NewServeMux()

	// Serve the Swagger UI documentation at /swagger/
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs/swagger-ui/dist"))))

	mux.HandleFunc("/", handlers.Repo.Home)
	mux.HandleFunc("/signin", handlers.Repo.SignIn)
	mux.HandleFunc("/signout", handlers.Repo.SignOut)
	mux.HandleFunc("/savecustomer", handlers.Repo.SaveCustomer)
	mux.HandleFunc("/deletecustomer", handlers.Repo.DeleteCustomer)
	return mux
}
