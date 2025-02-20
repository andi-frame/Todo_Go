package main

import (
	"net/http"

	"github.com/andi-frame/Todo_Go/database"
	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "All Good!"})
}

func main() {
	database.InitDB()

	r := gin.Default()

	r.GET("/health-check", healthCheck)

	r.Run(":8080")
}
