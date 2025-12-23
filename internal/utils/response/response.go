package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type response struct {
	Status string `json:"Status"`
	Error  string `json:"Error Check,omitempty"`
}

const (
	StatusSuccess = "Successfully created"
	StatusError   = "Error occurred"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) response {
	return response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(err validator.ValidationErrors) response {
	var errMsg []string
	for _, err := range err {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("The %s field is required", err.Field()))
		// Add more cases for different validation tags as needed
		default:
			errMsg = append(errMsg, fmt.Sprintf("The %s field is invalid", err.Field()))
		}
	}

	return response{
		Status: StatusError,
		Error:  strings.Join(errMsg, "; "),
	}
}
