package domain

import (
	"fmt"
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
	if result := tx.First(&Photo{}, "id = ?", c.PhotoID).RowsAffected; result == 0 {
		return fmt.Errorf("photo with id %d is not found", c.PhotoID)
	}

	if result := tx.First(&User{}, "id = ?", c.UserID).RowsAffected; result == 0 {
		return fmt.Errorf("user with id %d is not found", c.UserID)
	}

	return nil
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	if result := tx.First(&Comment{ID: c.ID}).RowsAffected; result == 0 {
		return fmt.Errorf("comment with id %d is not found", c.ID)
	}

	return nil
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	if result := tx.First(&Comment{ID: c.ID}).RowsAffected; result == 0 {
		return fmt.Errorf("comment with id %d is not found", c.ID)
	}

	return nil
}
