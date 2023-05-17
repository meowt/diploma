package models

type SignUpInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
