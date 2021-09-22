package model

import "time"

// UserModel ...
type UserModel struct {
	Id          int32     `gorm:"type:INT(32);column:id`
	Name        string    `gorm:"type:VARCHAR(64);column:name`
	Age         int32     `gorm:"type:INT(16);column:age`
	Email       string    `gorm:"type:VARCHAR(32);column:email`
	Phone       string    `gorm:"type:VARCHAR(32);column:phone`
	Description string    `gorm:"type:VARCHAR(64);column:description`
	CreateTime  time.Time `gorm:"type:TIMESTAMP;column:create_time`
	UpdateTime  time.Time `gorm:"type:TIMESTAMP;column:update_time`
}
