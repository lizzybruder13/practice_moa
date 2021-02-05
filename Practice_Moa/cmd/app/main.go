package main

import (
	"net/http"

	web "git.target.com/TargetOSS/avalanche-framework/web/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	web.Gin.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hi",
		})
	})

	web.StartServer()
}
