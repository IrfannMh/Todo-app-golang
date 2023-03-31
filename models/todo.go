package models

import "time"

type Todo struct {
	TodoId          uint      `gorm:"primaryKey" json:"id"`
	Title           string    `gorm:"not null" json:"title" form:"title"`
	ActivityGroupId string    `gorm:"not null" json:"activity_group_id" form:"activity_group_id"`
	IsActive        bool      `gorm:"not null" json:"is_active" form:"is_active"`
	Priority        string    `json:"priority"  form:"priority" gorm:"default:very-high"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
