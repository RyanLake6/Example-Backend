package main

import (
	"backend/api"
)

func main() {
	// setting up database connection
	//database.ConnectDatabase()

	// Calling api endpoints
	api.GetRouter()
}
