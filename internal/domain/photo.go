package domain

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

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

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	if result := tx.First(&User{ID: p.UserID}).RowsAffected; result == 0 {
		return fmt.Errorf("user with id %d is not found", p.UserID)
	}

	return nil
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if result := tx.First(&Photo{ID: p.ID}, "user_id = ?", p.UserID).RowsAffected; result == 0 {
		return fmt.Errorf("your photo with id %d is not found", p.ID)
	}

	return nil
}

func (p *Photo) BeforeDelete(tx *gorm.DB) (err error) {
	if result := tx.First(&Photo{ID: p.ID}, "user_id = ?", p.UserID).RowsAffected; result == 0 {
		return fmt.Errorf("your photo with id %d is not found", p.ID)
	}

	return nil
}
