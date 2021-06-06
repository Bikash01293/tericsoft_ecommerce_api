package products

import (
	"encoding/json"
	"net/http"

	// "tericsoft/categories"

	"tericsoft/db"
	"tericsoft/middleware"
	"tericsoft/products/productsfetch"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetProducts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	err := middleware.ValidateCookie(response, request)
	if err == nil {

		var products []productsfetch.Products
		params := mux.Vars(request)
		var err error
		if params["id"] != "" {
			err = db.DB.Preload("ProductCategory.Category").Where("product_id=?", params["id"]).Find(&products).Error
		} else {
			err = db.DB.Preload("ProductCategory.Category").Find(&products).Error
		}
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				response.Write([]byte(`{"response":"` + err.Error() + `"}`))
			} else {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(`{"response":"` + "oops! something went wrong! Try again" + `"}`))
			}
		} else {
			res := make(map[string]interface{})
			res["products"] = &products
			des, _ := json.Marshal(res)
			response.Write([]byte(des))
		}
	}

}
