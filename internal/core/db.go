package core

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitDb() (*error, *sql.DB) {
	var db *sql.DB
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName))

	if err != nil {
		// panic(err)
		return &err, nil
	}

	err = db.Ping()

	if err != nil {
		return &err, nil
		// panic(err)
	}

	fmt.Println("Connected to database")

	return nil, db

}
