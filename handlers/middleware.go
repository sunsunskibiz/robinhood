package handlers

import (
	"errors"
	"net/http"
	"strings"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/time/rate"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := ExtractToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			log.Println("JWTMiddleware: Error ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", int((*claims)["userID"].(float64)))
		c.Next()
	}
}

func ExtractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	// Check if the header starts with "Bearer "
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}

	return parts[1], nil
}

func RateLimitMiddleware(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
