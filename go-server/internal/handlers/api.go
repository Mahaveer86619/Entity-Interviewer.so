package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func APIHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "API Endpoint",
    })
}