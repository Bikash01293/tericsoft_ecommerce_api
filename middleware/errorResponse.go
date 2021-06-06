package middleware

import (
	"encoding/json"
	"net/http"
)

func ReturnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage ErrorResponse) {
	httpResponse := &ErrorResponse{Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusUnauthorized)
	response.Write(jsonResponse)
}
