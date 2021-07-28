package dto

type AnswerCreateDTO struct {
	Content    string `json:"content" form:"content" binding:"required"`
	UserID     uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	QuestionID uint64 `json:"question_id,omitempty" form:"question_id,omitempty"`
}

type AnswerUpdateDTO struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	UserID  uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
