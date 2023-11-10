package domain

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not_null; varchar(120)"`
	SocialMediaURL string `gorm:"not_null; varchar(120)"`
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {

	if err := tx.First(&User{}, "id = ?", s.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := tx.First(&SocialMedia{}, "id = ? AND user_id = ?", s.ID, s.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (s *SocialMedia) BeforeDelete(tx *gorm.DB) (err error) {
	if err := tx.First(&SocialMedia{}, "id = ? AND user_id = ?", s.ID, s.UserID).Error; err != nil {
		return err
	}

	return nil
}
