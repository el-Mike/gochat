package main

import (
	"log"
	"os"

	"github.com/el-Mike/gochat/routing"

	"github.com/el-Mike/gochat/persist"

	"github.com/joho/godotenv"
)

func main() {
	debugDbHost := os.Getenv("POSTGRES_HOST")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDBname := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_PORT")

	pgHost := debugDbHost

	if pgHost == "" {
		pgHost = os.Getenv("POSTGRES_HOST")
	}

	_, err := persist.InitDatabase(pgUser, pgPassword, pgDBname, pgHost, pgPort)

	if err != nil {
		log.Fatal("Database connection failed")
	}

	err = persist.AutoMigrate()

	if err != nil {
		log.Fatal(err)
	}

	routing.InitRouting()
}
