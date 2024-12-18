package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sunsunskibiz/robinhood/config"
	"github.com/sunsunskibiz/robinhood/models"
	"gorm.io/gorm"
)

type CreateThreadInput struct {
	Name   string `json:"name" binding:"required"`
	Detail string `json:"detail" binding:"required"`
}

func CreateThreadHandler(c *gin.Context) {
	var input CreateThreadInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := config.Config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	userID := c.MustGet("userID").(int)
	thread := models.Thread{
		Name:      input.Name,
		Detail:    input.Detail,
		Status:    "todo", // TODO: Change to enum
		CreatedBy: uint(userID),
	}

	if err := tx.Create(&thread).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	history := models.ThreadHistory{
		ThreadID:  thread.ID,
		Detail:    thread.Detail,
		Status:    thread.Status,
		CreatedBy: uint(userID),
	}
	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread history"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Thread created successfully",
		"thread":  thread,
	})
}

func GetThreadListHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	var threads []models.Thread
	if err := config.Config.DB.Limit(limit).Offset(offset).Order("created_at DESC").Find(&threads).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch threads"})
		return
	}

	var total int64
	if err := config.Config.DB.Model(&models.Thread{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count threads"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": threads,
		"pagination": gin.H{
			"current_page": page,
			"per_page":     limit,
			"total_pages":  (total + int64(limit) - 1) / int64(limit),
			"total_items":  total,
		},
	})
}

func GetThreadDetailHandler(c *gin.Context) {
	threadID := c.Param("id")

	var thread models.Thread

	if err := config.Config.DB.Preload("Histories").Preload("Comments").First(&thread, threadID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Thread not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch thread details"})
		}
		return
	}

	c.JSON(http.StatusOK, thread)
}
