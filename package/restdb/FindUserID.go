package restdb

import (
	"fmt"
	"log"
)

// FindUserID is for returning a user record defined by ID
func FindUserID(ID int) User {
	db := ConnectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return User{}
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM users where ID = $1\n", ID)
	if err != nil {
		log.Println("Query:", err)
		return User{}
	}

	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return User{}
		}

		u = User{c1, c2, c3, c4, c5, c6}
		log.Println("Found user:", u)
	}

	return u
}
