package main

import (
	"backend/api"
	"backend/database"
	"backend/logic"
	"log"
)

func main() {
	// setting up database connection
	db, err := database.DatabaseConnect()
	if err != nil {
		log.Fatal("Error connecting database: ", err)
	}

	// Creating router and endpoints
	logicClient := &logic.Logic{Database: db}
	apiClient := &api.Client{Logic: logicClient}
	router := apiClient.CreateRouter()

	// Running the backend server
	router.Run(":8080")
}
