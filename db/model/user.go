package model

import "time"

// UserModel ...
type UserModel struct {
	Id          int32     `gorm:"column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Age         int32     `gorm:"column:age" json:"age"`
	Email       string    `gorm:"column:email" json:"email"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	Description string    `gorm:"column:description" json:"description"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}
