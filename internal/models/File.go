package models

import "gorm.io/gorm"

type Content interface {
	GetSize() int64
	GetID() uint
	GetName() string
	GetPath() string
	IsFolder() bool
}

type File struct {
    gorm.Model
    FileName  string
    Path      string
    Size      int64
    UserID    uint
    FolderID  uint
}

type Folder struct {
    gorm.Model
    Name      string
    Path      string
    Size      int64
    UserID    uint  
	ParentID  *uint
    Children  []*Folder `gorm:"foreignkey:ParentID"`
    Files     []File    `gorm:"foreignKey:FolderID"`
}

func (f File) GetSize() int64 {
	return f.Size
}

func (f File) GetID() uint {
	return f.ID
}

func (f File) GetName() string {
	return f.FileName
}

func (f File) GetPath() string {
	return f.Path
}

func (f File) IsFolder() bool {
	return false
}



func (f *Folder) GetSize() int64 {
	var size int64
	for _, child := range f.Children {
		size += child.GetSize()
	}
	for _, file := range f.Files {
		size += file.Size
	}
	return size
}

func (f Folder) GetID() uint {
	return f.ID
}

func (f *Folder) GetName() string {
	return f.Name
}

func (f *Folder) GetPath() string {
	return f.Path
}

func (f *Folder) IsFolder() bool {
	return true
}
