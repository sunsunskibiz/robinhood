package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunsunskibiz/robinhood/config"
	"github.com/sunsunskibiz/robinhood/handlers"
	"golang.org/x/time/rate"
)

func main() {
	db := config.InitDB()
	config.Config.DB = db

	r := gin.Default()

	limiter := rate.NewLimiter(10, 1)
	r.Use(handlers.RateLimitMiddleware(limiter))

	r.POST("/login", handlers.LoginHandler())

	authRoutes := r.Group("/api", handlers.JWTMiddleware())
	authRoutes.POST("/threads", handlers.CreateThreadHandler)
	authRoutes.GET("/threads", handlers.GetThreadListHandler)
	authRoutes.GET("/threads/:id", handlers.GetThreadDetailHandler)
	authRoutes.PUT("/threads/:id", handlers.EditThreadHandler)
	authRoutes.DELETE("/threads/:id", handlers.DeleteThreadHandler)
	authRoutes.POST("/comments", handlers.CreateCommentHandler)
	authRoutes.PUT("/comments/:id", handlers.EditCommentHandler)
	authRoutes.DELETE("/comments/:id", handlers.DeleteCommentHandler)

	log.Println("Start robinhood server!")

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start the server")
	}
}
