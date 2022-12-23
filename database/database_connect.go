package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Database struct {
	DB *sql.DB
}

//var DB *sql.DB

func ConnectWithConnector(DB_USER *string, DB_PASS *string, DB_NAME *string, INSTANCE_CONNECTION_NAME *string) (*Database, error) {
	var database Database

	// loading in the .env file for environment variables
	err := godotenv.Load("./.env")
	if err != nil {
		return nil, fmt.Errorf("Couldn't load env variables: %v", err)
	}

	var (
		dbUser                 = DB_USER                  // e.g. 'my-db-user'
		dbPwd                  = DB_PASS                  // e.g. 'my-db-password'
		dbName                 = DB_NAME                  // e.g. 'my-database'
		instanceConnectionName = INSTANCE_CONNECTION_NAME // e.g.  maybe /cloudql/ 'project:region:instance'
		usePrivate             = os.Getenv("PRIVATE_IP")
	)

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %v", err)
	}
	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, *instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		*dbUser, *dbPwd, *dbName)

	database.DB, err = sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	return &database, nil
}

type User struct {
	ID    int64
	Name  string
	Phone string
}

func (d *Database) GetUser() ([]User, error) {
	rows, err := d.DB.Query("SELECT * FROM User")
	if err != nil {
		return nil, fmt.Errorf("error running query: %w", err)
		//gctx.Status(http.StatusInternalServerError)
		//return
	}
	users := []User{}
	defer rows.Close()
	for rows.Next() {
		var alb User
		rows.Scan(&alb.ID, &alb.Name, &alb.Phone)
		users = append(users, alb)
	}
	rows.Err()

	return users, nil

}
