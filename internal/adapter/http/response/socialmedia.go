package response

import "time"

// type CreateSocialmedia struct {
// 	ID             uint   `gorm:"primaryKey" json:"id"`
// 	Name           string `json:"name"`
// 	SocialMediaURL string `json:"social_media_url"`
// 	UserID         uint   `json:"user_id"`
// 	CreatedAd      time.Time
// }

type UserProfiles struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}

type Socialmedia struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           UserProfiles
}

type Sosmed struct {
	SocialMedia []*Socialmedia
}

// type UpdateSocialmedia struct{

// }
