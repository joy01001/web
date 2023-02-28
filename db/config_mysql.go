package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func ConnectDB() *sql.DB {

	log.Println("Initializing connection to database...")

	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))

	if err != nil {
		log.Fatal("Connection database error!", err)
	}

	log.Println("Database connected...")

	return db
}
