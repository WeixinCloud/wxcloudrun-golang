package dao

import (
	"wxcloudrun-golang/db/model"
)

// UserInterface ...
type UserInterface interface {
	AddUser(*model.UserModel) (int32, error)
	DeleteUserById(int32) error
	UpdateUserById(int32, *model.UserModel) error
	QueryUserById(int32) (*model.UserModel, error)
}

// UserInterfaceImp 实现结构
type UserInterfaceImp struct{}

// Imp 实现实例
var Imp UserInterface = &UserInterfaceImp{}
