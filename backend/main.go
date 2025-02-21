package main

import (
	"github.com/andi-frame/Todo_Go/database"
	"github.com/andi-frame/Todo_Go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()
	routes.Setup(r)
	r.Run(":8080")
}
