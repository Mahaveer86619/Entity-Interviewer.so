package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Mahaveer86619/Entity-Interviewer.so/go-server/internal/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

    log.Println("Starting Go server on :6062")
    if err := http.ListenAndServe(":6062", r); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}