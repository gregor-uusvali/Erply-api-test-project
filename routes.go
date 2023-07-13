// package main

// import (
// 	"Erply-api-test-project/handlers"
// 	"net/http"

// 	_ "github.com/swaggo/http-swagger"
// )

// func routes() http.Handler {
// 	mux := http.NewServeMux()

// 	// Serve the Swagger UI documentation at /swagger/
// 	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs/swagger-ui/dist"))))

//		mux.HandleFunc("/", handlers.Repo.Home)
//		mux.HandleFunc("/signin", handlers.Repo.SignIn)
//		mux.HandleFunc("/signout", handlers.Repo.SignOut)
//		mux.HandleFunc("/savecustomer", handlers.Repo.SaveCustomer)
//		mux.HandleFunc("/deletecustomer", handlers.Repo.DeleteCustomer)
//		return mux
//	}
package main

import (
	"Erply-api-test-project/handlers"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func routes() http.Handler {
	r := mux.NewRouter()

	// Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), 
	))

	r.HandleFunc("/", handlers.Repo.Home)
	r.HandleFunc("/signin", handlers.Repo.SignIn)
	r.HandleFunc("/signout", handlers.Repo.SignOut)
	r.HandleFunc("/savecustomer", handlers.Repo.SaveCustomer)
	r.HandleFunc("/deletecustomer", handlers.Repo.DeleteCustomer)

	return r
}
