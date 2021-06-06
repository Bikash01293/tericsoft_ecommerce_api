package categories

import (
	"encoding/json"
	"net/http"
	"tericsoft/db"
	"tericsoft/middleware"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetCategories(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	err := middleware.ValidateCookie(response, request)
	if err == nil {
		var category []Categories
		var err error
		params := mux.Vars(request)

		if params["id"] != "" {
			err = db.DB.Table("categories").Select("*").Where("category_id=?", params["id"]).Find(&category).Error
		} else {
			err = db.DB.Table("categories").Select("*").Find(&category).Error
		}
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				response.Write([]byte(`{"response":"` + err.Error() + `"}`))
			} else {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(`{"response":"` + err.Error() + `"}`))
			}
		} else {
			json.NewEncoder(response).Encode(category)
		}
	}
}
