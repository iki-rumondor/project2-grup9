package domain

import "time"

type Photo struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not_null;varchar(120)"`
	Caption     string `gorm:"not_null;varchar(120)"`
	PhotoUrl    string `gorm:"not_null; varchar(120)"`
	UserProfile User   `gorm:"foreignKey:UserID"`
	UserID      uint
	Comments    []Comment

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdatePhoto struct {
	ID       uint
	Title    string
	Caption  string
	PhotoUrl string

	UpdatedAt time.Time
}
