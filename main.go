package main

import (
	"go-issues-api/database"
	"go-issues-api/routes"
)

func main() {
	database.Connect("dev")
	router := routes.SetupRouter()
	router.Run(":3000")
}
