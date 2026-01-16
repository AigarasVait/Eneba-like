package db

import (
	"database/sql"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var DB *sql.DB

func Connect() {
	connString := "sqlserver://user:pass@localhost:1433?database=MyDB"

	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	log.Println("Connected to SQL Server")
}
