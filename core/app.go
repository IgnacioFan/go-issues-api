package main

import (
	"go-issues-api/core/server"
	"go-issues-api/database"
)

func main() {
	dbConn := database.Connect("dev")

	s := server.Server{
		DBConn: dbConn,
	}
	s.Start()
}
