package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Mahaveer86619/Entity-Interviewer.so/go-server/internal/handlers"
    "github.com/Mahaveer86619/Entity-Interviewer.so/go-server/internal/config"
)

func main() {
    // Initialize the database connection
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatalf("could not connect to the database: %v\n", err)
    }
    defer config.CloseDBConnection(db)

    // Ensure database tables are created or verified
    if err := config.CreateTables(db); err != nil {
        log.Fatalf("could not create/verify database tables: %v\n", err)
    }

    r := mux.NewRouter()
    r.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

    log.Println("Starting Go server on :6062")
    if err := http.ListenAndServe(":6062", r); err != nil {
        log.Fatalf("could not start server: %v\n", err)
    }
}