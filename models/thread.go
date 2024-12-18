package models

import "time"

type Thread struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Detail    string    `gorm:"type:text;not null" json:"detail"`
	Status    string    `gorm:"type:varchar(50);default:'todo'" json:"status"`
	CreatedBy uint      `gorm:"not null" json:"created_by"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
