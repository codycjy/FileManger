package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Level    int `gorm:"not null"`
	Files    []File `gorm:"many2many:user_files;"`
	Folders  []Folder `gorm:"many2many:user_folders;"`
}

