package categories

import (
	"net/http"
	"tericsoft/db"
	"tericsoft/middleware"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Delete(response http.ResponseWriter, request *http.Request) {
	err := middleware.ValidateCookie(response, request)
	if err == nil {
		var DeleteCategories Categories
		params := mux.Vars(request)

		err := db.DB.Table("categories").Where("category_id = ?", params["id"]).Delete(&DeleteCategories).Error

		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(`{"response":"` + err.Error() + `"}`))

			} else {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(`{"response":"` + "oops! something went wrong! Try again" + `"}`))
			}
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write([]byte(`{"response" : "successfully deleted!"}`))
		}
	}
}
