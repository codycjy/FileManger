package model

import "gorm.io/gorm"

type Content interface {
	GetSize() int64
	GetID() uint
}

type File struct {
	gorm.Model
	FileName string
	Path     string
	Size     int64
	UserID   uint
	Folders  []Folder `gorm:"many2many:folder_files;"`
}

type Folder struct {
	gorm.Model
	Name     string
	Path     string
	Size     int64
	UserID   uint
	ParentID *uint
	Children []*Folder `gorm:"foreignkey:ParentID"`
	Files    []File    `gorm:"many2many:folder_files;"`
}
