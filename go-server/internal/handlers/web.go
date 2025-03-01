package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func WebHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.templ", gin.H{
        "title": "Main website",
    })
}