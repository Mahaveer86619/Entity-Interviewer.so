package services

import (
    "net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    // Perform health check logic
    w.Write([]byte("OK"))
}