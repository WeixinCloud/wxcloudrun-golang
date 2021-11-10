package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName = "todo_list"

func (imp *ToDoItemInterfaceImp) GetToDoList() ([]*model.ToDoItemModel, error) {
	var err error
	var toDoList = []*model.ToDoItemModel{}

	cli := db.Get()
	err = cli.Table(tableName).Scan(&toDoList).Error

	return toDoList, err
}

func (imp *ToDoItemInterfaceImp) AddToDoItem(toDoItem *model.ToDoItemModel) error {
	var err error

	cli := db.Get()
	err = cli.Table(tableName).Create(toDoItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (imp *ToDoItemInterfaceImp) DeleteToDoItemById(id int32) error {
	cli := db.Get()
	return cli.Table(tableName).Exec("delete from todo_list where id = ?", id).Error
}

func (imp *ToDoItemInterfaceImp) UpdateToDoItemById(id int32, update *model.ToDoItemModel) error {
	cli := db.Get()
	return cli.Table(tableName).Where("id = ?", id).Updates(update).Error
}

func (imp *ToDoItemInterfaceImp) QueryToDoItemById(id int32) (*model.ToDoItemModel, error) {
	var err error
	var toDoItem = new(model.ToDoItemModel)

	cli := db.Get()
	err = cli.Table(tableName).Where("id = ?", id).First(toDoItem).Error

	return toDoItem, err
}
