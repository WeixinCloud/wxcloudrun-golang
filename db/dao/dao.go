package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const tableName = "user"

func (imp *UserInterfaceImp) AddUser(user *model.UserModel) (int32, error) {
	var err error

	cli := db.Get()
	err = cli.Table(tableName).Create(user).Error
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (imp *UserInterfaceImp) DeleteUserById(id int32) error {
	cli := db.Get()
	return cli.Table(tableName).Exec("delete from user where id = ?", id).Error
}

func (imp *UserInterfaceImp) UpdateUserById(id int32, update *model.UserModel) error {
	cli := db.Get()
	return cli.Table(tableName).Where("id = ?", id).Updates(update).Error
}

func (imp *UserInterfaceImp) QueryUserById(id int32) (*model.UserModel, error) {
	var err error
	var user = new(model.UserModel)

	cli := db.Get()
	err = cli.Table(tableName).Where("id = ?", id).First(user).Error

	return user, err
}
