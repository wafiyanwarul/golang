package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ApiResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ValidationError struct {
	Key    string `json:"key"`
	Errors string `json:"errors"`
}

type ValidationErrorResponse struct {
	Data []ValidationError `json:"data"`
}

func Response(code int, message string, data interface{}) ApiResponse {
	if code == http.StatusBadRequest {
		validationErrorsString := data.(string)
		validationErrors := parseValidationErrors(validationErrorsString)

		jsonData, _ := json.Marshal(validationErrors)
		return ApiResponse{
			Meta: Meta{
				Code:    code,
				Message: message,
			},
			Data: json.RawMessage(jsonData),
		}

	}

	return ApiResponse{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}

func parseValidationErrors(errorsString string) []ValidationError {
	lines := strings.Split(errorsString, "\n")
	validationErrors := make([]ValidationError, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		key := parts[1]
		errorMessage := strings.Join(parts[3:], " ")
		validationErrors = append(validationErrors, ValidationError{
			Key:    key,
			Errors: errorMessage,
		})
	}

	return validationErrors
}
