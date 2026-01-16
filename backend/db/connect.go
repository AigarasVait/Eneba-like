package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {

	var err error
	DB, err = sql.Open("sqlite", "file:mydb.sqlite")
	if err != nil {
		log.Fatal("Error creating connection pool:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	log.Println("Connected to SQL Server")
}
