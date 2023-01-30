package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func HouseRoutes(r *mux.Router) {
	houseRepository := repositories.RepositoryHouse(mysql.DB)
	h := handlers.HandlerHouse(houseRepository)

	r.HandleFunc("/house", h.AddHouse).Methods("POST")
	// r.HandleFunc("/houses", h.FindHouses).Methods("GET")
	// r.HandleFunc("/house/{id}", h.GetHouse).Methods("GET")
}
