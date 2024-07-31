package restdb

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID        int
	Username  string
	Password  string
	LastLogin int64
	Admin     int
	Active    int
}

// PostgreSQL Connection details
//
// We are using localhost as hostname because both
// the utility and PostgreSQL run on the same machine

// Hostname = "localhost" TODO
var (
	Hostname = "xbarber-db"
	Port     = 5432
	Username = "mtsouk"
	Password = "pass"
	Database = "restapi"
)

// postgresql://postgres:password@xbarber-db:5432?sslmode=disable

func ConnectPostgres() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname, Port, Username, Password, Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}
