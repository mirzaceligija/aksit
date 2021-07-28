package entity

import "time"

type Question struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Likes     uint      `gorm:"defualt:0" json:"likes"`
	UserID    uint64    `gorm:"not null" json:"-"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Answers   *[]Answer `gorm:"foreignkey:QuestionID" json:"answers"`
	CreatedAt time.Time `gorm:"<-:create" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
