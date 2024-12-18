package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunsunskibiz/robinhood/config"
	"github.com/sunsunskibiz/robinhood/handlers"
)

func main() {
	db := config.InitDB()
	config.Config.DB = db

	r := gin.Default()

	r.POST("/login", handlers.LoginHandler())

	authRoutes := r.Group("/api", handlers.JWTMiddleware())
	authRoutes.POST("/threads", handlers.CreateThreadHandler)
	authRoutes.GET("/threads", handlers.GetThreadListHandler)

	log.Println("Start robinhood server!")

	// TODO: Add rate limit
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start the server")
	}
}
