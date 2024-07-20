package models

import "time"

type User struct {
	ID			string 		`json:"id"`
	Username	string		`json:"username"`
	Email 		string		`json:"email"`
	AvatarURl	string		`json:"avatar_url"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}