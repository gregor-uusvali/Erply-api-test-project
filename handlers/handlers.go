package handlers

import (
	"Erply-api-test-project/api"
	"Erply-api-test-project/driver"
	"Erply-api-test-project/helpers"
	"Erply-api-test-project/models"
	"Erply-api-test-project/render"
	"Erply-api-test-project/repository"
	"Erply-api-test-project/repository/dbrepo"
	"fmt"
	"net/http"
	"strconv"
)

var Repo *Repository

// APIHandler handles the API endpoints
type Repository struct {
	DB repository.DatabaseRepo
}

// NewAPIHandler creates a new APIHandler instance
func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewSQLiteRepo(db.SQL),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.ClientError(w, http.StatusNotFound)
		return
	}
	if !helpers.CheckForSession() {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
	// customers, err := m.DB.GetCustomers()
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// 	return
	// }
	data := make(map[string]interface{})
	customers, err := api.GetCustumers(helpers.Session.User.ClientCode, helpers.Session.SessionKey)
	if err != nil {
		fmt.Println(err)
	}
	data["customers"] = *customers

	skUser, _ := api.GetSessionKeyUser(helpers.Session.User.ClientCode, helpers.Session.SessionKey)
	if render.Template(w, r, "home.html", &models.TemplateData{
		Message: skUser.Records[0].UserName,
		Data:    data,
	}) != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
}

func (m *Repository) SignIn(w http.ResponseWriter, r *http.Request) {
	if helpers.CheckForSession() {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	switch r.Method {
	case "GET":
		if render.Template(w, r, "index.html", &models.TemplateData{}) != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
			return
		}
	case "POST":
		clientCode := r.FormValue("clientCode")
		username := r.FormValue("username")
		password := r.FormValue("password")

		resp, err := api.VerifyUser(clientCode, username, password)
		if err != nil || resp == nil {
			if render.Template(w, r, "error.html", &models.TemplateData{
				Error: "Invalid credentials",
			}) != nil {
				http.Redirect(w, r, "/", http.StatusInternalServerError)
				return
			}
			return
		}
		// Create a new session
		session := models.Session{
			User: models.User{
				ClientCode: clientCode,
				Username:   username,
			},
			SessionKey: resp.Records[0].SessionKey,
		}
		helpers.Session = session
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (m *Repository) SignOut(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		helpers.Session.User.ClientCode = ""
		helpers.Session.User.Username = ""
		helpers.Session.SessionKey = ""
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	}
}

func (m *Repository) SaveCustomer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		firstName := r.FormValue("firstname")
		lastName := r.FormValue("lastname")
		email := r.FormValue("email")
		resp, err := api.SaveCustumer(helpers.Session.User.ClientCode, helpers.Session.SessionKey, firstName, lastName, email)
		if err != nil || resp == nil {
			fmt.Println(err)
			if render.Template(w, r, "error.html", &models.TemplateData{
				Error: "Failed to add customer",
			}) != nil {
				http.Redirect(w, r, "/", http.StatusInternalServerError)
				return
			}
			return
		}
		err = m.DB.AddCustomer(strconv.Itoa(resp.Records[0].CustomerID), firstName, lastName, email)
		if err != nil {
			helpers.ServerError(w, err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (m *Repository) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("deletecustomer")
		err := api.DeleteCustomer(helpers.Session.User.ClientCode, helpers.Session.SessionKey, id)
		if err != nil {
			if render.Template(w, r, "error.html", &models.TemplateData{
				Error: "Could not delete customer",
			}) != nil {
				http.Redirect(w, r, "/", http.StatusInternalServerError)
				return
			}
			return
		}
		err = m.DB.DeleteCustomer(id)
		if err != nil {
			helpers.ServerError(w, err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
