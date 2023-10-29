package response

import "time"

type User struct {
	Age       uint      `json:"age"`
	Email     string    `json:"email"`
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatedUser struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserProfile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type JWT struct {
	Token string `json:"token"`
}
