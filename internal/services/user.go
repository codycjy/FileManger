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

func GetUserById(user *models.User) error {
	db := mysql.GetDB()
	err := db.Where("id=?", user.Model.ID).First(&user).Error
	return err
}
