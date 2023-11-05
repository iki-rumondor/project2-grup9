package request

type SocialMedia struct {
	Name           string `json:"name" valid:"required~name is required"`
	SocialMediaURL string `json:"social_media_url" valid:"required~social_media_url is required"`
}
