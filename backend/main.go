package main

import (
	"github.com/andi-frame/Todo_Go/database"
	"github.com/andi-frame/Todo_Go/routes"
)

func main() {
	database.InitDB()
	routes.Setup()
}
