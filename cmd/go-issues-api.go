package main

import (
	"go-issues-api/database"
	"go-issues-api/internal/server"
)

func main() {
	dbConn := database.Connect("dev")

	s := server.Server{
		DBConn: dbConn,
	}
	s.Start()
}
