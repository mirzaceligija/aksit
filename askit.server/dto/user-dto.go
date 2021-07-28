package dto

type UserUpdateDTO struct {
	ID        uint64 `json:"id" form:"id"`
	FirstName string `json:"firstName" form:"firstName" binding:"required"`
	LastName  string `json:"lastName" form:"lastName" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password  string `json:"password,omitempty" form:"password,omitempty" validate:"min:5"`
}
