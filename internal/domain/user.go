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
	Photos []Photo
  SocialMedia []SocialMedia

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPass
	return nil

}
