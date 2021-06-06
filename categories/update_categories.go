package categories

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
		var UpdatedCategories Categories
		params := mux.Vars(request)
		json.NewDecoder(request.Body).Decode(&UpdatedCategories)
		err := db.DB.Table("categories").Where("category_id = ?", params["id"]).Update(&UpdatedCategories).Error

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
			json.NewEncoder(response).Encode(UpdatedCategories)
		}
	}
}
