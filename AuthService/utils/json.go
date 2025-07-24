package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJSONResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteJSONSuccessResponse(w http.ResponseWriter, status int, message string, data any) error {
	response := map[string]any{}
	response["data"] = data
	response["message"] = message
	response["status"] = "success"

	return WriteJSONResponse(w, status, response)
}

func WriteJSONErrorResponse(w http.ResponseWriter, status int, message string, err error) error {
	response := map[string]any{}
	response["error"] = err.Error()
	response["message"] = message
	response["status"] = "error"

	return WriteJSONResponse(w, status, response)
}

func ReadJSONBody(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Prevents decoding of unknown fields
	return decoder.Decode(result)
}
