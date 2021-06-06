package middleware

import (
	"errors"
	"net/http"
	"strings"
)

func AddCookie(response http.ResponseWriter, name, value string) {

	cookie := http.Cookie{
		Name:  name,
		Value: value,
	}
	http.SetCookie(response, &cookie)
}

func RemoveCookie(response http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "Bearer",
		Value: "",
	}
	http.SetCookie(response, &cookie)
}

func ValidateCookie(response http.ResponseWriter, request *http.Request) error {
	var errorResponse = ErrorResponse{
		Message: "invalid! token try again or login to continue.",
	}
	bearerToken := request.Header.Get("cookie")
	var authorizationToken string
	authorizationTokenArray := strings.Split(bearerToken, "=")
	if len(authorizationTokenArray) > 1 {
		authorizationToken = authorizationTokenArray[1]
	}
	token, _ := VerifyToken(authorizationToken)
	if token == "" {
		ReturnErrorResponse(response, request, errorResponse)
		return errors.New("error in token")
	}
	return nil
}
