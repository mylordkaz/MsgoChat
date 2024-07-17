package models

type User struct {
	ID			int		`json:"id"`
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Password	string	`json:"-"`
	GoogleID 	string	`json:"-"`
	GithubID 	string	`json:"-"`
	Provider 	string	`json:"provider"` // local, google, github, etc.
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}