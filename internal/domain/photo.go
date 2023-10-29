package domain

import "time"

type Photo struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not_null;varchar(120)"`
	Caption     string `gorm:"not_null;varchar(120)"`
	PhotoUrl    string `gorm:"not_null; varchar(120)"`
	UserID      uint
	UserProfile User `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
