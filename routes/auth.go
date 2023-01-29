package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/login", h.Login).Methods("POST")
}
