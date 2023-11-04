package request

type Comment struct {
	Message string `json:"message" valid:"required~message is required"`
	PhotoID uint   `json:"photo_id" valid:"required~photo_id is required"`
}

type CommentID struct {
	Message string `json:"message" valid:"required~message is required"`
}
