package handlers

import (
	"net/http"
	"time"

	"github.com/Mahaveer86619/Entity-Interviewer.so/internal/services"
	"github.com/gin-gonic/gin"
)

type AdminData struct {
	DBStatus        string
	Uptime          string
	AdminerLink     string
	ServerStartTime string
}

var serverStartTime = time.Now()

func AdminHandler(c *gin.Context) {
	dbStatus := "Connected"
	uptime, err := services.GetUptime()
	if err != nil {
		dbStatus = "Disconnected"
	}

	adminData := AdminData{
		DBStatus:        dbStatus,
		Uptime:          uptime.String(),
		AdminerLink:     "http://localhost:6061",
		ServerStartTime: serverStartTime.Format(time.RFC1123),
	}

	c.HTML(http.StatusOK, "admin.templ", adminData)
}
