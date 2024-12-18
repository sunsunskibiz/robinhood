package models

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ThreadID  uint   `gorm:"not null" json:"thread_id"`
	Content   string `gorm:"type:text;not null" json:"content"`
	CreatedBy uint   `gorm:"not null" json:"created_by"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}