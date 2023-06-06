package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Level    int `gorm:"not null" json:"level"`
	Files    []File `gorm:"many2many:user_files;" json:"files"`
	Folders  []Folder `gorm:"many2many:user_folders;" json:"folders"`
}

