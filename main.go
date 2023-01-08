package main

import (
	"fmt"
	"go-issues-api/model"
	"go-issues-api/routes"
	"os"
)

func add(a, b int) int {
	return a + b
}

func main() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Taipei",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME_DEV"),
	)
	model.SetupDatabase(dsn)
	router := routes.SetupRouter()
	router.Run(":3000")
}
