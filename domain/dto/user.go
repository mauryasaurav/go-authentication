package dto

/* User Validation*/
type UserValidator struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserUpdateValidator struct {
	FirstName string `json:"first_name" binding:"min=1"`
	LastName  string `json:"last_name" binding:"min=1"`
	Email     string `json:"email" binding:"min=5"`
}

/* User Validation*/
type UserLoginValidator struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
