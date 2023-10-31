package domain

import "time"

type User struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	Username        string        `gorm:"unique;not null" json:"username"`
	Email           string        `gorm:"unique;not null" json:"email"`
	Password        string        `gorm:"not null" json:"-"`
	Age             int           `json:"age"`
	ProfileImageURL string        `json:"profile_image_url"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Photos          []Photo       `gorm:"foreignKey:UserID"`
	SocialMedia     []SocialMedia `gorm:"foreignKey:UserID"`
}
