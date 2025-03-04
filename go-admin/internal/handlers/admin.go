package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
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
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	go func(ws *websocket.Conn) {
		for {
			uptime := time.Since(startTime).Seconds()
			v, _ := mem.VirtualMemory()
			c, _ := cpu.Percent(0, false)
			data := map[string]interface{}{
				"uptime": uptime,
				"cpu":    c[0],
				"ram":    v.UsedPercent,
			}
			ws.WriteJSON(data)
			time.Sleep(2 * time.Second)
		}
	}(ws)
}
