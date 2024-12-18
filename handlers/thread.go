package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sunsunskibiz/robinhood/config"
	"github.com/sunsunskibiz/robinhood/models"
)

func CreateThreadHandler(c *gin.Context) {
	var thread models.Thread

	// Bind JSON to the thread struct // TODO: Remove comment
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default values
	if thread.Status == "" {
		thread.Status = "active"
	}

	// Save thread to the database
	if err := config.Config.DB.Create(&thread).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	// Return the created thread
	c.JSON(http.StatusCreated, thread)
}

func GetThreadListHandler(c *gin.Context) {
	// Default pagination parameters // TODO: Remove comment
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Calculate offset
	offset := (page - 1) * limit

	// Get threads from the database with pagination
	var threads []models.Thread
	if err := config.Config.DB.Limit(limit).Offset(offset).Order("created_at DESC").Find(&threads).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch threads"})
		return
	}

	// Count total threads for pagination metadata
	var total int64
	if err := config.Config.DB.Model(&models.Thread{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count threads"})
		return
	}

	// Response with threads and pagination metadata
	c.JSON(http.StatusOK, gin.H{
		"data": threads,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total_pages":  (total + int64(limit) - 1) / int64(limit), // Ceiling of total/limit
			"total_items":  total,
		},
	})
}