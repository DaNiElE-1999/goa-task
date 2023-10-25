package main

import (
	"database/sql"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	godotenv.Load(".env")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3333)/books") //toBeFixed with env var
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
