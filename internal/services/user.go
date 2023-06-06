package services

import (
	"filemanger/internal/models"
	"filemanger/internal/repositories/mysql"
)

func RegisterUser(user *models.User) error {
	db := mysql.GetDB()
	err := db.Create(&user).Error
	return err
}

func LoginUser(user *models.User) error {
	db := mysql.GetDB()
	err := db.Where("username=? and password=?", user.Username, user.Password).First(&user).Error
	return err
}

func GetUserById(user *models.User,flag int) error {
	db := mysql.GetDB()
	var err error 
	if flag==0{
		err = db.Where("id=?", user.Model.ID).First(&user).Error
	}else{
		preloads := "Files"
		if flag==2{
			preloads ="Folders"
		}
		err = db.Preload(preloads).Where("id=?", user.Model.ID).First(&user).Error
	}
	return err
}

func UpdateUser(user *models.User) error {
	db := mysql.GetDB()
	err := db.Save(&user).Error
	return err
}
