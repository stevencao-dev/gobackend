package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/stevencao-dev/indie-store-backend/internal/handlers"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/products", handlers.GetProducts)

    handler := cors.AllowAll().Handler(mux)

    log.Println("Backend running on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
	