package handlers

import (
    "net/http"
    "log"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("HealthCheckHandler called")
    w.Write([]byte("OK"))
}