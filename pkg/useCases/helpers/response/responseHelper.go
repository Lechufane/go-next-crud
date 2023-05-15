package responseHelper

import (
	"encoding/json"
	"net/http"

	"github.com/Lechufane/go-next-crud/pkg/domain/innerDomain/response"
)

func ResponseBuilder(status int, message string, data interface{}) ([]byte, error) {
	response := response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return marshalResponse, nil
}

func ResponseStatusChecker(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		return
	}
}

// Parse response data object to any struct (use struct as a pointer to parse content into it)
func ParseResponseStruct(v any, res *http.Response, statusToCompare response.Status) response.Status {
	var auxResponse response.Response
	err := json.NewDecoder(res.Body).Decode(&auxResponse)
	if err != nil {
		return response.InternalServerError
	}
	defer res.Body.Close()

	if auxResponse.Message != statusToCompare.String() {
		return response.GetStatusByName(auxResponse.Message)
	}
	parsedObj, err := json.Marshal(auxResponse.Data)
	if err != nil {
		return response.InternalServerError
	}
	err = json.Unmarshal(parsedObj, &v)
	if err != nil {
		return response.InternalServerError
	}

	return response.GetStatusByName(auxResponse.Message)
}
