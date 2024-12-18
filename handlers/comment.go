package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunsunskibiz/robinhood/config"
	"github.com/sunsunskibiz/robinhood/models"
	"gorm.io/gorm"
)

type createCommentInput struct {
	ThreadID uint   `json:"thread_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

func CreateCommentHandler(c *gin.Context) {
	var input createCommentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var thread models.Thread
	if err := config.Config.DB.First(&thread, input.ThreadID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Thread not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check thread"})
		}
		return
	}

	userID := c.MustGet("userID").(int)
	now := time.Now()
	comment := models.Comment{
		ThreadID:  input.ThreadID,
		Content:   input.Content,
		CreatedBy: uint(userID),
		UpdatedBy: uint(userID),
		UpdatedAt: &now,
	}
	if err := config.Config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		"comment": comment,
	})
}

type editCommentInput struct {
	Content string `json:"content" binding:"required"`
}

func EditCommentHandler(c *gin.Context) {
	commentID := c.Param("id")
	var input editCommentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var comment models.Comment
	if err := config.Config.DB.First(&comment, commentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comment"})
		}
		return
	}

	userID := c.MustGet("userID").(int)
	if comment.CreatedBy != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to edit this comment"})
		return
	}

	comment.Content = input.Content
	comment.UpdatedBy = uint(userID)
	now := time.Now()
	comment.UpdatedAt = &now
	if err := config.Config.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
		"comment": comment,
	})
}

func DeleteCommentHandler(c *gin.Context) {
	commentID := c.Param("id")

	var comment models.Comment
	if err := config.Config.DB.First(&comment, commentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comment"})
		}
		return
	}

	userID := c.MustGet("userID").(int)
	if comment.CreatedBy != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to edit this comment"})
		return
	}

	if err := config.Config.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
	})
}
