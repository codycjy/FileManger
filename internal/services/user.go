package services

import (
	model "filemanger/internal/models"
	mysql "filemanger/internal/repositories/Mysql"
)


func RegisterUser(user *model.User) (error){
	db:=mysql.GetDB()
	err:=db.Create(&user).Error
	return err
}

func LoginUser(user *model.User) (error){
	db:=mysql.GetDB()
	err:=db.Where("username=? and password=?",user.Username,user.Password).First(&user).Error
	return err
}

func GetUserById(user *model.User) (error){
	db:=mysql.GetDB()
	err:=db.Where("id=?",user.Model.ID).First(&user).Error
	return err
}
