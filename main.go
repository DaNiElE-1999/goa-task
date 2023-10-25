package main

import (
	"database/sql"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db, err := sql.Open("mysql", "root@/books")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
