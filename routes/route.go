package routes

import (
	"net/http"
	"tericsoft/auth"
	"tericsoft/categories"
	"tericsoft/products"

	"github.com/gorilla/mux"
)

// Routes function is defined for intializing all the end points
func Routes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/register", auth.Register).Methods("POST")
	router.HandleFunc("/api/login", auth.Login).Methods("POST")
	router.HandleFunc("/api/logout", auth.Logout).Methods("POST")

	router.HandleFunc("/api/categories/create", categories.Create).Methods("POST")
	router.HandleFunc("/api/categories/show", categories.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories/show/{id}", categories.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories/update/{id}", categories.Update).Methods("PUT")
	router.HandleFunc("/api/categories/delete/{id}", categories.Delete).Methods("DELETE")

	router.HandleFunc("/api/products/create", products.Create).Methods("POST")
	router.HandleFunc("/api/products/show", products.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/show/{id}", products.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/update/{id}", products.Update).Methods("PUT")
	router.HandleFunc("/api/products/delete/{id}", products.Delete).Methods("DELETE")
	router.HandleFunc("/api/categories/{id}/products", categories.Delete).Methods("GET")

	http.ListenAndServe(":8000", router)
}
