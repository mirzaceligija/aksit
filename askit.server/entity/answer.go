package entity

import "time"

type Answer struct {
	ID         uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Content    string    `gorm:"type:varchar(255);not null" json:"content"`
	UserID     uint64    `gorm:"not null" json:"-"`
	User       User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	QuestionID uint64    `gorm:"<-:create;not null" json:"-"`
	Question   Question  `gorm:"foreignkey:QuestionID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	CreatedAt  time.Time `gorm:"<-:create" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
