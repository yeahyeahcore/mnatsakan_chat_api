package respond

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendJSON sends JSON response to client
func SendJSON(response http.ResponseWriter, data interface{}, statusCode int) {
	encodedJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(response, fmt.Errorf("could not marshal response: %v", err).Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(statusCode)
	response.Write(encodedJSON)
}

// SendMessage sends message response to client
func SendMessage(response http.ResponseWriter, message string, statusCode int) {
	responseMessage := map[string]string{
		message: message,
	}

	SendJSON(response, responseMessage, statusCode)
}
