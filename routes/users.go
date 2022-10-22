package routes

import (
	"waysfood/handlers"
	"waysfood/pkg/middleware"
	"waysfood/pkg/mysql"
	"waysfood/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)

	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", middleware.Auth(h.FindUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user", middleware.UploadFile(h.CreateUser)).Methods("POST")
	r.HandleFunc("/user/{id}", middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/user/{id}", middleware.Auth(h.DeleteUser)).Methods("DELETE")
}
