package model

import "time"

// ToDoItemModel
type ToDoItemModel struct {
	Id         int32     `gorm:"column:id" json:"id"`
	Title      string    `gorm:"column:title" json:"title"`
	Status     string    `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}
