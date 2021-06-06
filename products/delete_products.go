package products

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
		var DeleteProduct Products
		var DeleteCategoryProduct CategoryProduct
		params := mux.Vars(request)

		errPro := db.DB.Table("products").Where("product_id = ?", params["id"]).Delete(&DeleteProduct).Error
		errProCat := db.DB.Table("category_products").Where("product_id = ?", params["id"]).Delete(&DeleteCategoryProduct).Error

		if errPro != nil && errProCat != nil {
			if gorm.IsRecordNotFoundError(errPro) {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(`{"response":"` + errPro.Error() + `"}`))

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
