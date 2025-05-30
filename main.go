package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/yourusername/crud-template/database"
    "github.com/yourusername/crud-template/routes"
)

func main() {
    database.Connect()

    router := mux.NewRouter()
    routes.RegisterRoutes(router)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
