package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Config struct{
	DB *gorm.DB
}

func InitDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	var db *gorm.DB
    var err error

    for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err == nil {
			sqlDB, err := db.DB()
			if err == nil {
                break
            }

            err = sqlDB.Ping()
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