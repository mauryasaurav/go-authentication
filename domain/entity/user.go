package entity

import "time"

/* User Validation*/
type UserSchema struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
