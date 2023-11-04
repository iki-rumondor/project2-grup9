package domain

import "time"

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not_null; varchar(120)"`
	SocialMediaURL string `gorm:"not_null; varchar(120)"`
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           User
}
