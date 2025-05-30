package utils

import (
	"encoding/json"
	"net/http"
)

// Response represents the standard API response structure
type Response struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

// JSONResponse sends a JSON response with the given status code and data
func JSONResponse(w http.ResponseWriter, status int, message string, data interface{}) {
    response := Response{
        Status:  status,
        Message: message,
        Data:    data,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)
}

// ErrorResponse sends a JSON error response
func ErrorResponse(w http.ResponseWriter, status int, message string, err error) {
    response := Response{
        Status:  status,
        Message: message,
        Error:   err.Error(),
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(response)
} 