package restdb

// Database: PostgreSQL
//
// Functions to support the interaction with the

import (
	_ "github.com/lib/pq"
)

func init() {
	ConnectPostgres()
}
