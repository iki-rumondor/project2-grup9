package domain

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Message   string `gorm:"not_null;varchar(120)"`
	UserID    uint
	PhotoID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Photo     Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if err := tx.First(&Photo{}, "id = ?", c.PhotoID).Error; err != nil {
		return err
	}

	if err := tx.First(&User{}, "id = ?", c.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := tx.First(&Comment{}, "id = ? AND user_id = ?", c.ID, c.UserID).Error; err != nil {
		return err
	}

	return nil
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	if err := tx.First(&Comment{}, "id = ? AND user_id = ?", c.ID, c.UserID).Error; err != nil {
		return err
	}

	return nil
}
