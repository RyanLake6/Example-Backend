package main

import (
	"backend/api"
	"backend/database"
	"backend/logic"
	"log"
)

func main() {
	// Loading in all the environment variables
	// godotenv.Load("./.env")
	// DB_USER := flag.String("DB_USER", "", "username for db connecting/authentication")
	// DB_PASS := flag.String("DB_PASS", "", "Password for db connecting/authentication")
	// DB_NAME := flag.String("DB_NAME", "", "Name of db to connect to")
	// INSTANCE_CONNECTION_NAME := flag.String("INSTANCE_CONNECTION_NAME", "", "Instance ID of cloud sql")
	// flag.Parse()

	// setting up database connection
	db, err := database.DatabaseConnect()
	if err != nil {
		log.Fatal("Error connecting database: ", err)
	}

	// Creating router and endpoints
	c := &api.Client{Logic: &logic.Logic{Database: db}}
	router := c.CreateRouter()

	// Running the backend server
	router.Run(":8080")
}
