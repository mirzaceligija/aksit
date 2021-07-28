package dto

type QuestionUpdateDTO struct {
	ID     uint64 `json:"id" form:"id" binding:"required"`
	Title  string `json:"title" form:"title" binding:"required"`
	UserID uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type QuestionCreateDTO struct {
	Title  string `json:"title" form:"title" binding:"required"`
	UserID uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
