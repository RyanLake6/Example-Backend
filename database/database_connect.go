package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Database struct {
	DB *sql.DB
}

func DatabaseConnect() (*Database, error) {
	var database Database
	var err error

	godotenv.Load()
	// Capture connection properties.
	cfg := mysql.Config{
		ParseTime: true,
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "household-db",
	}
	// Get a database handle.
	database.DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := database.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to sql!")
	return &database, nil
}

type User struct {
	UUID     int64
	Username string
	Password string
}

func (d *Database) GetUser() ([]User, error) {
	rows, err := d.DB.Query("SELECT * FROM user")
	if err != nil {
		return nil, fmt.Errorf("error running query: %w", err)
		//gctx.Status(http.StatusInternalServerError)
		//return
	}
	users := []User{}
	defer rows.Close()
	for rows.Next() {
		var u User
		rows.Scan(&u.UUID, &u.Username, &u.Password)
		users = append(users, u)
	}
	rows.Err()

	return users, nil

}
