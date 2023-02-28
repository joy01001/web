package main

import (
	"github.com/joho/godotenv"
	"log"
	"web/db"
	"web/routes"
)

func init() {
	log.Println("Loading environment variables...")

	err := godotenv.Load("tmp/.env")

	if err != nil {
		log.Fatal("Environment variables not loaded!")
	}
}

func main() {
	db.ConnectDB()
	routes.Handler()

}
