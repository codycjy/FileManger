package file

import "gorm.io/gorm"

type FileObject struct {
	gorm.Model
	FileName string
	FileType string
	Path     string
	OwnerID  uint
}

type Folder struct {
	*FileObject
	Files []FileObject
}	

type File struct {
	*FileObject
	Size int64
}
