package products

import (
	"encoding/json"
	"net/http"
	"tericsoft/db"
	"tericsoft/middleware"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Update(response http.ResponseWriter, request *http.Request) {
	err := middleware.ValidateCookie(response, request)
	if err == nil {
		var UpdatedProduct Products
		params := mux.Vars(request)
		err := json.NewDecoder(request.Body).Decode(&UpdatedProduct)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{"response":"` + err.Error() + `"}`))
		}
		err = db.DB.Table("products").Where("product_id = ?", params["id"]).Update(&UpdatedProduct).Error

		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				response.WriteHeader(http.StatusNotFound)
				response.Write([]byte(`{"response":"` + err.Error() + `"}`))
			} else {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(`{"response":"` + "oops! something went wrong! Try again" + `"}`))
			}
		} else {
			response.Header().Set("Content-Type", "application/json")
			json.NewEncoder(response).Encode(UpdatedProduct)
		}
	}
}
