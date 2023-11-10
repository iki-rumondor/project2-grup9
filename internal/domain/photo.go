package domain

import (
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
	if err := tx.First(&User{}, "id = ?", p.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := tx.First(&Photo{}, "id = ? AND user_id = ?", p.ID, p.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (p *Photo) BeforeDelete(tx *gorm.DB) (err error) {
	if err := tx.First(&Photo{}, "id = ? AND user_id = ?", p.ID, p.UserID).Error; err != nil {
		return err
	}

	return nil
}
