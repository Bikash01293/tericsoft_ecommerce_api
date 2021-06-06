package categories

import (
	"encoding/json"
	"net/http"
	"tericsoft/db"
	"tericsoft/middleware"
)

func Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	err := middleware.ValidateCookie(response, request)
	if err == nil {
		var category Categories

		err := json.NewDecoder(request.Body).Decode(&category)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{"message":"` + err.Error() + `"}`))
			return
		}

		err = db.DB.Create(&category).Error
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message":"` + err.Error() + `"}`))
			return
		}
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(category)
	}
}
