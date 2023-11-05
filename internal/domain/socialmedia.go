package domain

import (
	"fmt"
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

	if result := tx.First(&User{}, "id = ?", s.UserID).RowsAffected; result == 0 {
		return fmt.Errorf("user with id %d is not found", s.UserID)
	}

	return nil
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	if result := tx.First(&SocialMedia{ID: s.ID}, "user_id = ?", s.UserID).RowsAffected; result == 0{
		return fmt.Errorf("your social media with id %d is not found", s.ID)
	}

	return nil
}

func (s *SocialMedia) BeforeDelete(tx *gorm.DB) (err error) {
	if result := tx.First(&SocialMedia{ID: s.ID}, "user_id = ?", s.UserID).RowsAffected; result == 0{
		return fmt.Errorf("your social media with id %d is not found", s.ID)
	}

	return nil
}