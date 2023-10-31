package domain

import "time"

type Comment struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"user_id"`
	PhotoID      uint      `json:"photo_id"`
	Message      string    `json:"message"`
	Tittle       string    `json:"tittle"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserProfile  User      `gorm:"foreignKey:UserID"`
	PhotoProfile Photo     `gorm:"foreignKey:PhotoID"`
}

type UpdateComment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Tittle    string    `json:"tittle"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
