package main

import (
	"backend/api"
	"backend/database"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
)

func main() {
	godotenv.Load("./.env")
	DB_USER := flag.String("DB_USER", "", "username for db connecting/authentication")
	DB_PASS := flag.String("DB_PASS", "", "Password for db connecting/authentication")
	DB_NAME := flag.String("DB_NAME", "", "Name of db to connect to")
	INSTANCE_CONNECTION_NAME := flag.String("INSTANCE_CONNECTION_NAME", "", "Instance ID of cloud sql")
	flag.Parse()

	fmt.Print("DB USER IS: ", *DB_USER)

	// setting up database connection
	db, err := database.ConnectWithConnector(DB_USER, DB_PASS, DB_NAME, INSTANCE_CONNECTION_NAME)
	if err != nil {
		log.Fatal("Error connecting database: ", err)
	}

	// Creating router and endpoints
	// TODO:
	// pass in an instance of busniess logic/database queries to getRouter
	// b := b.New()
	// router := api.GetRouter(b)
	c := &api.Client{Database: db}
	router := c.GetRouter()

	router.Run(":8080")
}
