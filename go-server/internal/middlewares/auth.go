package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token != "valid_token" {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }
    c.Next()
}