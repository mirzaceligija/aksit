package dto

type RegisterDTO struct {
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
	Email     string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password  string `json:"password" form:"password" binding:"required" validate:"min:5"`
}
