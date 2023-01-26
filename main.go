package main

import (
	"go-issues-api/database"

	_r "go-issues-api/routes"
)

func main() {
	dbConn := database.Connect("dev")

	s := _r.Router{
		DBConn: dbConn,
	}
	s.Start()
}
