package restdb

import (
	"fmt"
	"log"
)

// ListLogged is for returning all logged users
// This was created by mistake - the server uses
// ReturnLoggedUsers() instead!
func ListLogged() []User {
	db := ConnectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return []User{}
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE active = 1\n")
	if err != nil {
		log.Println(err)
		return []User{}
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		_ = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		temp := User{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)
	}

	log.Println("All:", all)
	return all
}
