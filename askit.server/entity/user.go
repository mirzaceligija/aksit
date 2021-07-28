package entity

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string    `gorm:"type:varchar(255);null" json:"firstName"`
	LastName  string    `gorm:"type:varchar(255);null" json:"lastName"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string    `gorm:"->;<-;not null" json:"-"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time `gorm:"<-:create" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
