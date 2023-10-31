package response

import "time"

type UserProfile struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}

type PhotoProfile struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Tittle   string `json:"tittle"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type CreateComment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateComment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Tittle    string    `json:"tittle"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      UserProfile
	Photo     PhotoProfile
}

type Comments struct {
	Comments []*Comment
}
