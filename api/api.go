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
			return fmt.Errorf("failed to authenticate")
		}
	case *models.GetSessionKeyUserResponse:
		if resp.Status.ErrorCode != 0 {
			return fmt.Errorf("failed to authenticate")
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
