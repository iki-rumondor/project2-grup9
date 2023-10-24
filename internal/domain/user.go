package domain

import (
	"time"

	"github.com/iki-rumondor/project2-grup9/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not_null;varchar(120)"`
	Email    string `gorm:"unique;not_null; varchar(120)"`
	Password string `gorm:"not_null; varchar(120)"`
	Age      uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashPass, err := utils.HashPassword(u.Password)
	if err != nil{
		return err
	}
	u.Password = hashPass
	return nil
}
