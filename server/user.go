package server

import (
	"simple-todo/model"
	"simple-todo/utils"
)

func GetUser(uid string) (user []model.Todo) {
	d.Where("uid = ?", uid).Find(&user)
	return user
}

func CheckLogin(uid string, pwd string) error {
	newpwd := utils.EncryMd5(pwd)
	var user []model.User
	err := d.Where("uid = ? AND password = ?", uid, newpwd).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
