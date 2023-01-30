package main

import (
	"fmt"
	"housy/database"
	"housy/pkg/mysql"
	"housy/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SERVER_NAME := os.Getenv("SERVER_NAME")
	PORT := os.Getenv("SERVER_PORT")
	VERSION := os.Getenv("API_VERSION")

	// Database
	mysql.DatabaseInit()

	// Migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/" + VERSION).Subrouter())

	fmt.Println("Server is running on http://" + SERVER_NAME + ":" + PORT + "/api/" + VERSION)
	http.ListenAndServe(SERVER_NAME+":"+PORT, r)
}
