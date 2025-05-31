package routes

import (
	"net/http"

	"github.com/dhuharohim/golang-crud-api/controllers"
	"github.com/gorilla/mux"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
    router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
    router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

    router.PathPrefix("/").Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.WriteHeader(http.StatusOK)
    })
}
