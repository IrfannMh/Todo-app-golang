package models

import "time"

type Activity struct {
	ActivityId uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"not null" json:"title" form:"title" `
	Email      string    `gorm:"not null" json:"email" form:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
