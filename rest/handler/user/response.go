package user

type authResponse struct {
	Token string       `json:"token"`
	User  *userResponse `json:"user"`
}

type userResponse struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
