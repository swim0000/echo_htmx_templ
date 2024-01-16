package models

type Signup struct {
	ID       int    `form:"id"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}
