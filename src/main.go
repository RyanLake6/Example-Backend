package main

import (
	"backend/api"
	"backend/database"
)

func main() {
	// setting up database connection
	database.ConnectDatabase()

	// Calling api endpoints
	api.GetRouter()
}
