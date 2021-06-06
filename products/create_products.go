package products

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
		var products Products

		err := json.NewDecoder(request.Body).Decode(&products)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{"response":"` + err.Error() + `"}`))
		}
		err = db.DB.Create(&products).Error

		if err != nil {
			response.Write([]byte(`{"response":"` + err.Error() + `"}`))
			return
		}
		json.NewEncoder(response).Encode(products)
	}
}
