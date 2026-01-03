package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status_code"`
	Error string `json:"error"`
}

const (
	StatusOk = "ok"
	StatusError = "error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errors []string
	for _, err := range errs {
		switch err.ActualTag(){
			case "required":
				errors = append(errors, fmt.Sprintf("%s is required", err.Field()))
			case "email":
				errors = append(errors, fmt.Sprintf("%s is not a valid email", err.Field()))
			case "min":
				errors = append(errors, fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param()))
			case "max":
				errors = append(errors, fmt.Sprintf("%s must be less than %s", err.Field(), err.Param()))
			default:
				errors = append(errors, fmt.Sprintf("%s is not valid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error: strings.Join(errors, ", "),
	}
}