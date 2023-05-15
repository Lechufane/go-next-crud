package requestHelper

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Lechufane/go-next-crud/pkg/domain/innerDomain/response"
)

func GetRequest(url string, headers http.Header) (*http.Response, response.Status) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	if headers != nil {
		req.Header = headers
	}
	req.Header.Set("Content-Type", "application/json")
	result, err := client.Do(req)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	return result, response.SuccesfulRequest
}
func PostRequest(url string, body interface{}, headers http.Header) (*http.Response, response.Status) {
	// Parse object to json
	parsedObject, err := json.Marshal(body)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(parsedObject))
	if err != nil {
		return nil, response.ErrorInRequest
	}
	if headers != nil {
		req.Header = headers
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	return res, response.SuccesfulRequest
}
func PostRequestBearerToken(url string, body interface{}, bearerToken string, headers http.Header) (*http.Response, response.Status) {
	// Parse object to json
	parsedObject, err := json.Marshal(body)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	// Set authorization header
	bearer := "Bearer " + bearerToken
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(parsedObject))
	if headers != nil {
		req.Header = headers
	}
	req.Header.Add("Authorization", bearer)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	return res, response.SuccesfulRequest
}

// Body must be a marshalled object
func PutRequest(url string, body interface{}, headers http.Header) (*http.Response, response.Status) {
	// Parse object to json
	parsedObject, err := json.Marshal(body)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(parsedObject))
	if err != nil {
		return nil, response.ErrorInRequest
	}
	if headers != nil {
		req.Header = headers
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	return res, response.SuccesfulRequest
}

func DeleteRequest(url string, headers http.Header) (*http.Response, response.Status) {
	// Create a HTTP post request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	if headers != nil {
		req.Header = headers
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, response.ErrorInRequest
	}
	return res, response.SuccesfulRequest
}
