package main

import (
	"filemanger/internal/models"
	"filemanger/internal/repositories/mysql"
	"fmt"
)

type Content interface {
	GetSize() int64
	GetID() uint
}

func DropTables() {
	var tables = []string{"user_files", "user_folders",}
	db:=mysql.GetDB()
	for _, table := range tables {
		db.Exec("DROP TABLE IF EXISTS " + table)
	}
	db.Migrator().DropTable(&models.User{}, &models.File{}, &models.Folder{})
}

type User models.User
type File models.File
type Folder models.Folder
func main() {
	// Implement GORM connection and initialization
	db:=mysql.GetDB()

	DropTables()
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
	nestedFolder := models.Folder{
		Name:     "Nested",
		Path:     "/Root/Nested",
		UserID:   user.ID,
		ParentID: &rootFolder.ID,
	}

	db.Create(&nestedFolder)
	nestedFolder2 := models.Folder{
		Name:     "Nested2",
		Path:     "/Root/Nested2",
		UserID:   user.ID,
		ParentID: &rootFolder.ID,
	}
	db.Create(&nestedFolder2)
	nestedFolder3 := models.Folder{
		Name:     "Nested3",
		Path:     "/Root/Nested2/Nested3",
		UserID:   user.ID,
		ParentID: &nestedFolder2.ID,
	}


	db.Create(&nestedFolder3)

	rootFolder.Children = append(rootFolder.Children, &nestedFolder,&nestedFolder2)
	nestedFolder2.Children=append(nestedFolder2.Children, &nestedFolder3)
	db.Save(&rootFolder)
	db.Save(&nestedFolder2)
	// Create a file inside the nested folder
	file := File{
		FileName: "testfile.txt",
		Path:     "/Root/Nested/testfile.txt",
		Size:     1024,
		UserID:   user.ID,
		FolderID: nestedFolder2.ID,
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


