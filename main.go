package main

import (
	"log"
	"net/http"

	"github.com/dhuharohim/golang-crud-api/database"
	"github.com/dhuharohim/golang-crud-api/middleware"
	"github.com/dhuharohim/golang-crud-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// CORS middleware
	r.Use(middleware.CORSMiddleware)

	// Connect DB
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Register routes
	routes.RegisterRoutes(r)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
