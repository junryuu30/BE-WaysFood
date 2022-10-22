package routes

import (
	"waysfood/handlers"
	"waysfood/pkg/middleware"
	"waysfood/pkg/mysql"
	"waysfood/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)

	h := handlers.HandlerProduct(productRepository)

	// r.HandleFunc("/products", h.FindProduct).Methods("GET")
	// r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	// r.HandleFunc("/product", h.CreateProduct).Methods("POST")
	// r.HandleFunc("/product/{id}", h.UpdateProduct).Methods("PATCH")
	// r.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")
	// r.HandleFunc("/restosproduct/{userid}", h.GetProductByResto).Methods("GET")

	r.HandleFunc("/products", middleware.Auth(h.FindProduct)).Methods("GET")
	r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH")
	r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")

}
