package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"tericsoft/db"
	"tericsoft/middleware"

	"golang.org/x/crypto/bcrypt"
)

//login user
func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var user, dbUser User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	err = db.DB.Table("users").Select("id, name, email, password").
		Where("email = ?", user.Email).
		Find(&dbUser).Error
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		log.Println(passErr)
		response.WriteHeader(http.StatusForbidden)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	jwtToken, err := middleware.GenerateJWT(user.Email)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	middleware.AddCookie(response, "Bearer", jwtToken)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"message":"` + "login successfully" + `"}`))
}
