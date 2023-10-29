package request

type AllUserData struct{
	Age uint `json:"age" valid:"required~field age is required, range(9|99)~minimum age must be 9 years old"`
	Email string `json:"email" valid:"required~field email is required, email, unique(users.email)~the email has already been taken"`
	Password string `json:"password" valid:"required~field password is required, length(6|99)~password at least 6 character"`
	Username string `json:"username" valid:"required~field username is required, unique(users.username)~the username has already been taken"`
}

type UserWithEmail struct{
	Email string `json:"email" valid:"required~field email is required, email"`
	Password string `json:"password" valid:"required~field password is required, length(6|99)~password at least 6 character"`
}
