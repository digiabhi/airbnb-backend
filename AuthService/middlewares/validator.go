package middlewares

import (
	"AuthService/dto"
	"AuthService/utils"
	"context"
	"fmt"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDTO

		// Read and decode the JSON body into the payload
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			err := utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			if err != nil {
				return
			}
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			err := utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			if err != nil {
				return
			}
			return
		}

		fmt.Println("Payload received for login:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload) // Create a new context with the payload

		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}

func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDTO

		// Read and decode the JSON body into the payload
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			err := utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			if err != nil {
				return
			}
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.Validator.Struct(payload); err != nil {
			err := utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			if err != nil {
				return
			}
			return
		}
		fmt.Println("Payload received for login:", payload)

		ctx := context.WithValue(r.Context(), "payload", payload) // Create a new context with the payload

		next.ServeHTTP(w, r.WithContext(ctx)) // Call the next handler in the chain
	})
}
