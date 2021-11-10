package dao

import (
	"wxcloudrun-golang/db/model"
)

// ToDoItemInterface ...
type ToDoItemInterface interface {
	GetToDoList() ([]*model.ToDoItemModel, error)
	AddToDoItem(*model.ToDoItemModel) error
	DeleteToDoItemById(int32) error
	UpdateToDoItemById(int32, *model.ToDoItemModel) error
	QueryToDoItemById(int32) (*model.ToDoItemModel, error)
}

// ToDoItemInterfaceImp 实现结构
type ToDoItemInterfaceImp struct{}

// Imp 实现实例
var Imp ToDoItemInterface = &ToDoItemInterfaceImp{}
