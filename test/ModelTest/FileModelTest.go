package main

import (
	"filemanger/internal/models"
	"filemanger/internal/repositories/mysql"
	"fmt"

	"gorm.io/gorm"
)

type Content interface {
	GetSize() int64
	GetID() uint
}

type User models.User
type File models.File
type Folder models.Folder

func dropTables(db *gorm.DB) {

	var table=[]string{"user_folders","user_files","users","files","folders",}
	for _,v:=range table{
		db.Exec("DROP TABLE IF EXISTS "+v)
	}

}

func main() {
	// Implement GORM connection and initialization
	db:=mysql.GetDB()

	// Drop all tables
	dropTables(db)

	db.AutoMigrate(&models.User{}, &models.File{}, &models.Folder{},)

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
	user.Folders = append(user.Folders, models.Folder(rootFolder), models.Folder(nestedFolder))
	db.Save(&user)

	// Create a file inside the nested folder
	file := File{
		FileName: "testfile.txt",
		Path:     "/Root/Nested/testfile.txt",
		Size:     1024,
		UserID:   user.ID,
		FolderID: nestedFolder.ID,
	}

	db.Create(&file)
	// Add the file to the user's Files field
	user.Files = append(user.Files, models.File(file))
	db.Save(&user)
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


