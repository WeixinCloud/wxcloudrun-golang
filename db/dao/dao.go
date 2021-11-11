package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName = "todo_list"

// GetToDoList 获取todo list
func (imp *ToDoItemInterfaceImp) GetToDoList() ([]*model.ToDoItemModel, error) {
	var err error
	var toDoList = []*model.ToDoItemModel{}

	cli := db.Get()
	err = cli.Table(tableName).Scan(&toDoList).Error

	return toDoList, err
}

// AddToDoItem 添加todo项
func (imp *ToDoItemInterfaceImp) AddToDoItem(toDoItem *model.ToDoItemModel) error {
	var err error

	cli := db.Get()
	err = cli.Table(tableName).Create(toDoItem).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteToDoItemById 根据id删除todo项
func (imp *ToDoItemInterfaceImp) DeleteToDoItemById(id int32) error {
	cli := db.Get()
	return cli.Table(tableName).Exec("delete from todo_list where id = ?", id).Error
}

// UpdateToDoItemById 根据id更新todo项
func (imp *ToDoItemInterfaceImp) UpdateToDoItemById(id int32, update *model.ToDoItemModel) error {
	cli := db.Get()
	return cli.Table(tableName).Where("id = ?", id).Updates(update).Error
}

// QueryToDoItemById 根据id查询todo项
func (imp *ToDoItemInterfaceImp) QueryToDoItemById(id int32) (*model.ToDoItemModel, error) {
	var err error
	var toDoItem = new(model.ToDoItemModel)

	cli := db.Get()
	err = cli.Table(tableName).Where("id = ?", id).First(toDoItem).Error

	return toDoItem, err
}
