package main

import (
	mysql "filemanger/internal/repositories/Mysql"
	"fmt"

	"gorm.io/gorm"
)

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

type User struct {
	gorm.Model
	Username string
	Password string
	Level    int
	Files    []File `gorm:"many2many:user_files;"`
	Folders  []Folder `gorm:"many2many:user_folders;"`
}

func main() {
	// Implement GORM connection and initialization
	db:=mysql.GetDB()
	db.AutoMigrate(&User{}, &File{}, &Folder{},)

	// Create a user
	user := User{
		Username: "testuser",
		Password: "testpassword",
		Level:    1,
	}

	db.Create(&user)

	// Create a root folder
	rootFolder := Folder{
		Name:   "Root",
		Path:   "/",
		UserID: user.ID,
	}

	db.Create(&rootFolder)

	// Create a nested folder inside the root folder
	nestedFolder := Folder{
		Name:     "Nested",
		Path:     "/Root/Nested",
		UserID:   user.ID,
		ParentID: &rootFolder.ID,
	}

	db.Create(&nestedFolder)

	// Create a file inside the nested folder
	file := File{
		FileName: "testfile.txt",
		Path:     "/Root/Nested/testfile.txt",
		Size:     1024,
		UserID:   user.ID,
		Folders:  []Folder{nestedFolder},
	}

	db.Create(&file)

	// Query the database and print the results
	var folders []Folder
	db.Preload("Children").Preload("Children.Files").Where("user_id = ?", user.ID).Where("parent_id IS NULL").Find(&folders)
	fmt.Println("Folders and their content:")
	for _, folder := range folders {
		fmt.Printf("Folder ID: %d, Name: %s, ParentID: %v\n", folder.ID, folder.Name, folder.ParentID)
		fmt.Println("  - Subfolders:")
		for _, subfolder := range folder.Children {
			fmt.Printf("    - Folder ID: %d, Name: %s\n", subfolder.ID, subfolder.Name)
			fmt.Println("      - Files:")
			for _, file := range subfolder.Files {
				fmt.Printf("        - File ID: %d, Name: %s\n", file.ID, file.FileName)
			}
		}
	}
}


