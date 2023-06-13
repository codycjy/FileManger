package repositories

import (
	"filemanger/internal/models"
	"filemanger/internal/repositories/mysql"
)

func GetFileByID(id uint) (*models.File, error) {
	var file models.File
	db := mysql.GetDB()
	if err := db.First(&file, id).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func GetFolderByID(id uint) (*models.Folder, error) {
    var folder models.Folder
    db := mysql.GetDB()
    if err := db.Preload("Files").Preload("Children").First(&folder, id).Error; err != nil {
        return nil, err
    }
    return &folder, nil
}

func CreateFile(file *models.File) error {
	db := mysql.GetDB()
	if err := db.Create(file).Error; err != nil {
		return err
	}
	return nil
}

func CreateFolder(folder *models.Folder) error {
	db := mysql.GetDB()
	if err := db.Create(folder).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFileByID(id uint) error {
	db := mysql.GetDB()
	if err := db.Delete(&models.File{}, id).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFolderByID(id uint) error {
	db := mysql.GetDB()

	// Get the folder with the given ID
	var folder models.Folder
	if err := db.Preload("Files").Preload("Children").First(&folder, id).Error; err != nil {
		return err
	}

	// Delete all files associated with the folder
	for _, file := range folder.Files {
		if err := db.Delete(&file).Error; err != nil {
			return err
		}
	}

	// Delete all subfolders associated with the folder
	for _, subfolder := range folder.Children {
		if err := DeleteFolderByID(subfolder.ID); err != nil {
			return err
		}
	}

	// Delete the folder
	if err := db.Delete(&folder).Error; err != nil {
		return err
	}

	return nil
}

func LoginUser(user *models.User) (error) {
    db := mysql.GetDB()
    if err := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
        return err
    }
    return nil
}
