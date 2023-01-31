package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func PropertyRoutes(r *mux.Router) {
	propertyRepository := repositories.RepositoryProperty(mysql.DB)
	h := handlers.HandlerProperty(propertyRepository)

	// sementara
	// r.HandleFunc("/property", middleware.Auth(h.AddProperty)).Methods("POST")
	// this
	r.HandleFunc("/properties", middleware.Auth(h.FindProperties)).Methods("GET")
	r.HandleFunc("/property/{id}", h.GetProperty).Methods("GET")
	r.HandleFunc("/property", middleware.Auth(middleware.UploadFile(h.AddProperty))).Methods("POST")
	r.HandleFunc("/property/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProperty))).Methods("PATCH")
	r.HandleFunc("/property/{id}", middleware.Auth(h.DeleteProperty)).Methods("DELETE")
	// r.HandleFunc("/houses", h.FindHouses).Methods("GET")
	// r.HandleFunc("/house/{id}", h.GetHouse).Methods("GET")
}
