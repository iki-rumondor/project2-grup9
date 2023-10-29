package request

type CreatePhoto struct{
	Title string `json:"title" valid:"required~field title is required"`
	Caption string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"required~field photo_url is required"`
}