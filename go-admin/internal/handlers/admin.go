package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/cpu"
)

var startTime = time.Now()

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Handle service registration
	w.Write([]byte("Service registered"))
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/dashboard.html")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).Seconds()
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Percent(0, false)

	data := map[string]interface{}{
		"uptime": uptime,
		"cpu":    c[0],
		"ram":    v.UsedPercent,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
