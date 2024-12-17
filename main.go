package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunsunskibiz/robinhood/config"
	"github.com/sunsunskibiz/robinhood/handlers"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	r := gin.Default()

	r.POST("/login", handlers.LoginHandler(db))

	authRoutes := r.Group("/api", handlers.JWTMiddleware())
	authRoutes.GET("/user", handlers.UserHandler(db))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to connect to database after retries: %v", err)
	}

	log.Println("Hash Password: ", string(hashedPassword))
	log.Println("Start robinhood server!")

	r.Run(":8080")
}
