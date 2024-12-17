package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
    var db *sql.DB
    var err error

    for i := 0; i < 10; i++ {
        db, err = sql.Open("mysql", dsn)
        if err == nil {
            err = db.Ping()
            if err == nil {
                break
            }
        }
        log.Printf("Database not ready, retrying in 2 seconds... (%d/10)", i+1)
        time.Sleep(2 * time.Second)
    }
    if err != nil {
        log.Fatalf("Failed to connect to database after retries: %v", err)
    }
    log.Println("Successfully connected to the database!")

	return db
}