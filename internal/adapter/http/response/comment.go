package response

import "time"

type PhotoProfile struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
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
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string      `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
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
