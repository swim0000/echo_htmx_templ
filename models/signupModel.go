package models

type Signup struct {
	ID       int64
	Email    string `form:"email" validate:"required,email,max=255,"`
	Password string `form:"password" validate:"required,min=8,max=32,"`
}
