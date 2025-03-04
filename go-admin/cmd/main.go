package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Mahaveer86619/Entity-Interviewer.so/go-admin/internal/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/admin/register", handlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/admin/dashboard", handlers.DashboardHandler).Methods("GET")
    r.HandleFunc("/admin/health", handlers.HealthCheckHandler).Methods("GET")

    log.Println("Starting Admin service on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}