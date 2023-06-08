package main

import (
	"filemanger/internal/api"
	"filemanger/internal/models"
	"filemanger/internal/repositories/mysql"
)

func main() {
	db := mysql.GetDB()
	db.AutoMigrate(&models.User{}, &models.File{}, &models.Folder{})
	api.Router()
}
