package main

import (
	"fmt"
	"go-issues-api/database"
	"go-issues-api/internal/server"
)

func main() {
	fmt.Print("hey")
	dbConn := database.Connect("dev")

	s := server.Server{
		DBConn: dbConn,
	}
	s.Start()
}
