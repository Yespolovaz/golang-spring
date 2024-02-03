package main

import (
	"fmt"
	"net/http"
	"tsis1/internal"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/anime", handlers.GetAnimeList).Methods("GET")
	r.HandleFunc("/anime/{id:[1-6]+}", handlers.GetAnimeDetails).Methods("GET")
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
