package storage

import "database/sql"

var database *sql.DB

func GetDB(db *sql.DB) {
	database = db
}
