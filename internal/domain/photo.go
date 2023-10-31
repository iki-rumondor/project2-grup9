package domain

import "time"

type Photo struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Tittle      string    `json:"tittle"`
	Caption     string    `json:"caption"`
	PhotoURL    string    `json:"photo_url"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Comments    []Comment
	UserProfile User `gorm:"foreignKey:UserID"`
}
