package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Level    int
	Files    []File `gorm:"many2many:user_files;"`
	Folders  []Folder `gorm:"many2many:user_folders;"`
}

