package main

import (
	"filemanger/internal/api"
	model "filemanger/internal/models"
	mysql "filemanger/internal/repositories/Mysql"
)

func main(){
	db:=mysql.GetDB()
	db.AutoMigrate(&model.User{},&model.File{},&model.Folder{})
	api.Router()
}
