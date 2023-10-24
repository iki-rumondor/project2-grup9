package response

type User struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type JWT struct {
	Token string `json:"token"`
}
