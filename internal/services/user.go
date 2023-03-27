package services

import (
	model "filemanger/internal/models"
	mysql "filemanger/internal/repositories/Mysql"
)


func RegisterUser(user model.User) (model.User,error){
	db:=mysql.GetDB()
	err:=db.Create(&user).Error
	return user,err
}

func LoginUser(user model.User) (model.User,error){
	db:=mysql.GetDB()
	var u model.User
	err:=db.Where("username=? and password=?",user.Username,user.Password).First(&u).Error
	return u,err
}

