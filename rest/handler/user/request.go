package user

type registerRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
