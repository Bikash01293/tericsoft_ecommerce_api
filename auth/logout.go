package auth

import (
	"net/http"
	"tericsoft/middleware"
)

//logout user
func Logout(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	err := middleware.ValidateCookie(response, request)
	if err == nil {
		middleware.RemoveCookie(response, request)
		response.Write([]byte(`{"message":"Successfully logged out"}`))
	}
}
