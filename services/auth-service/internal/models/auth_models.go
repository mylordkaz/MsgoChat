package models

type User struct {
	ID			int		`json:"id"`
	Username	string	`json:"username"`
	Password	string	`json:"-"`
	Email		string	`json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}