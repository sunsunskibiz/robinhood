package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(int)

		var name, email string
		query := "SELECT name, email FROM users WHERE id = ?"
		err := db.QueryRow(query, userID).Scan(&name, &email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": userID, "name": name, "email": email})
	}
}