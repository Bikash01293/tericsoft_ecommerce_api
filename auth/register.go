package auth

import (
	"encoding/json"
	"net/http"
	"tericsoft/db"
)

//register user
func Register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var user User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	user.Password = getHash([]byte(user.Password))
	
	err = db.DB.Create(&user).Error
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.WriteHeader(http.StatusCreated)
	response.Write([]byte(`{"response":"Successfully registered"}`))

}
