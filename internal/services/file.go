package services

import (
	"errors"
	model "filemanger/internal/models"
	repositories "filemanger/internal/repositories"
	mysql "filemanger/internal/repositories/Mysql"
)

// GetFileInfo is used to get the file information
func GetFileInfoByID(id uint) (*model.File, error) {
	file, err := repositories.GetFileByID(id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// AddFile is used to add a file
func AddFile(content *model.File) error {
	if (*content).GetName() == "" || (*content).GetPath() == "" {
		return errors.New("name and path are required")
	}
	return repositories.AddFile(content)
}

// DeleteFile is used to delete a file
func DeleteFile(content *model.Content) error {
	if (*content).GetID() == 0 {
		return errors.New("id is required")
	}
	return repositories.DeleteFile((*content).GetID())
}

func GetFolderByID(id uint) (*model.Folder, error) {
	folder, err := repositories.GetFolderByID(id)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

// AddFolder is used to add a folder
func AddFolder(content *model.Folder) error {
	if (*content).GetName() == "" || (*content).GetPath() == "" {
		return errors.New("name and path are required")
	}
	return repositories.AddFolder(content)
}

// DeleteFolder is used to delete a folder
func DeleteFolder(id uint) error {
	db := mysql.GetDB()

	// Retrieve the folder and its children
	var folder model.Folder
	err := db.Preload("Children").Preload("Files").First(&folder, id).Error
	if err != nil {
		return err
	}

	// Delete all the files in the folder
	for _, file := range folder.Files {
		err = repositories.DeleteFile(file.ID)
		if err != nil {
			return err
		}
	}

	// Recursively delete all the subfolders
	for _, child := range folder.Children {
		err = DeleteFolder(child.ID)
		if err != nil {
			return err
		}
	}

	// Delete the folder itself
	result := db.Delete(&folder)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("folder not found")
	}

	return nil
}
