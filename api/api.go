package api

import (
	"Erply-api-test-project/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func makeRequest(clientCode, requestType, sessionKey string, data url.Values, responsePtr interface{}) error {
	if sessionKey != "" {
		data.Set("sessionKey", sessionKey)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s.erply.com/api/", clientCode), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &responsePtr)
	if err != nil {
		return err
	}

	switch resp := responsePtr.(type) {
	case *models.Response:
		if resp.Status.ErrorCode != 0 {
			return fmt.Errorf("authentication failed: %v", resp.Status.ErrorCode)
		}
	case *models.GetSessionKeyInfoResponse:
		if resp.Status.ErrorCode != 0 {
			return fmt.Errorf("failed to get session key info")
		}
	case *models.GetSessionKeyUserResponse:
		if resp.Status.ErrorCode != 0 {
			return fmt.Errorf("failed to get session key user")
		}
	case *models.SaveCustomerResponse:
		if resp.Status.ErrorCode != 0 {
			return fmt.Errorf("failed to save customer")
		}
	case *models.CustomerResponse:
		if resp.Status.ErrorCode != 0 {
			return fmt.Errorf("failed to get customers")
		}
	default:
		return fmt.Errorf("unsupported response type")
	}

	return nil
}

// @Summary Verify user
// @Description Verify user credentials and retrieve a session key.
// @Tags Authentication
// @Accept x-www-form-urlencoded
// @Param clientCode formData string true "Client code"
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Produce json
// @Success 200 {object} models.Response
// @Router /verifyUser [post]
func VerifyUser(clientCode, username, password string) (*models.Response, error) {
	data := url.Values{
		"clientCode":      {clientCode},
		"username":        {username},
		"password":        {password},
		"request":         {"verifyUser"},
		"sendContentType": {"1"},
	}

	response := new(models.Response)
	err := makeRequest(clientCode, "verifyUser", "", data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// @Summary Get session key info
// @Description Get information about the session key.
// @Tags Authentication
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Security SessionKeyAuth
// @Produce json
// @Success 200 {object} models.GetSessionKeyInfoResponse
// @Router /getSessionKeyInfo [post]
func GetSessionKeyInfo(clientCode, sessionKey string) (*models.GetSessionKeyInfoResponse, error) {
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"getSessionKeyInfo"},
		"sendContentType": {"1"},
	}

	response := new(models.GetSessionKeyInfoResponse)
	err := makeRequest(clientCode, "getSessionKeyInfo", sessionKey, data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// @Summary Get session key user
// @Description Get the user associated with the session key.
// @Tags Authentication
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Success 200 {object} models.GetSessionKeyUserResponse
// @Router /getSessionKeyUser [post]
func GetSessionKeyUser(clientCode, sessionKey string) (*models.GetSessionKeyUserResponse, error) {
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"getSessionKeyUser"},
		"sendContentType": {"1"},
	}

	response := new(models.GetSessionKeyUserResponse)
	err := makeRequest(clientCode, "getSessionKeyUser", sessionKey, data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// @Summary Save customer
// @Description Save a customer.
// @Tags Customer
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Param firstName query string true "First name"
// @Param lastName query string true "Last name"
// @Param email query string true "Email"
// @Produce json
// @Success 200 {object} models.SaveCustomerResponse
// @Router /saveCustomer [post]
func SaveCustumer(clientCode, sessionKey, firstName, lastName, email string) (*models.SaveCustomerResponse, error) {
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"saveCustomer"},
		"firstName":       {firstName},
		"lastName":        {lastName},
		"email":           {email},
		"sendContentType": {"1"},
	}

	response := new(models.SaveCustomerResponse)
	err := makeRequest(clientCode, "saveCustomer", sessionKey, data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// @Summary Get customers
// @Description Get a list of customers.
// @Tags Customer
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Security SessionKeyAuth
// @Produce json
// @Success 200 {object} models.CustomerResponse
// @Router /getCustomers [post]
func GetCustumers(clientCode, sessionKey string) (*models.CustomerResponse, error) {
	data := url.Values{
		"clientCode":      {clientCode},
		"sessionKey":      {sessionKey},
		"request":         {"getCustomers"},
		"sendContentType": {"1"},
	}

	response := new(models.CustomerResponse)
	err := makeRequest(clientCode, "getCustomers", sessionKey, data, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// @Summary Delete customer
// @Description Delete a customer.
// @Tags Customer
// @Param clientCode query string true "Client code"
// @Param sessionKey query string true "Session key"
// @Param id query string true "Customer ID"
// @Success 200 {string} string "DELETE request sent successfully"
// @Router /deleteCustomer [delete]
func DeleteCustomer(clientCode, sessionKey, id string) error {
	url := fmt.Sprintf("https://api-crm-eu.erply.com/v1/customers/individuals/%s", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create DELETE request: %v", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("sessionKey", sessionKey)
	req.Header.Set("clientCode", clientCode)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send DELETE request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete customer, status code: %d", resp.StatusCode)
	}

	fmt.Println("DELETE request sent successfully")
	return nil
}
