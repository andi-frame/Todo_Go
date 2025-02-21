package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup() {
	r := gin.Default()

	r.GET("/health-check", HealthCheck)
	r.Run(":8080")
}
