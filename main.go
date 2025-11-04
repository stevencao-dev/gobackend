package main

import (
	"fmt"
	"net/http"
	"github/stevencao-dev/gobackend/internal/handlers"
)

func main() {
	http.HandleFunc("/products", handlers.GetProducts)

	fmt.Println("Backend running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
	