package models

import "time"

type Thread struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Detail    string    `gorm:"type:text;not null" json:"detail"`
	Status    string    `gorm:"type:varchar(50);default:'todo'" json:"status"`
	CreatedBy uint      `gorm:"not null" json:"created_by"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedBy uint      `gorm:"not null" json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	Histories []ThreadHistory   `gorm:"foreignKey:ThreadID" json:"histories"`
	Comments  []Comment         `gorm:"foreignKey:ThreadID" json:"comments"`
}

type ThreadHistory struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ThreadID   uint   `gorm:"not null" json:"thread_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Detail    string    `gorm:"type:text;not null" json:"detail"`
	Status    string    `gorm:"type:varchar(50);default:'todo'" json:"status"`
	CreatedBy uint      `gorm:"not null" json:"created_by"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedBy uint      `gorm:"not null" json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ThreadID  uint   `gorm:"not null" json:"thread_id"`
	Content   string `gorm:"type:text;not null" json:"content"`
	CreatedBy uint   `gorm:"not null" json:"created_by"`
	CreatedAt string `gorm:"autoCreateTime" json:"created_at"`
	UpdatedBy uint      `gorm:"not null" json:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at"`
}