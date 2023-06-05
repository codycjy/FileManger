package services

import (
	"errors"
	"filemanger/internal/models"
	"filemanger/internal/repositories"
	"filemanger/internal/repositories/mysql"

	"gorm.io/gorm"
)

func GetContentInfo(content *models.Content) models.Content {
	if file, ok := (*content).(*models.File); ok {
		// Content is a file
		file, err := repositories.GetFileByID(file.ID)
		if err != nil {
			return nil
		}
		return file
	} else if folder, ok := (*content).(*models.Folder); ok {
		// Content is a folder
		folder, err := repositories.GetFolderByID(folder.ID)
		if err != nil {
			return nil
		}
		return folder
	} else {
		// Invalid content type
		return nil
	}
}

func AddContent(content models.Content) error {
	if file, ok := content.(*models.File); ok {
		// Content is a file
		err := repositories.CreateFile(file)
		if err != nil {
			return err
		}
	} else if folder, ok := content.(*models.Folder); ok {
		// Content is a folder
		err := repositories.CreateFolder(folder)
		if err != nil {
			return err
		}
	} else {
		// Invalid content type
		return errors.New("invalid content type")
	}
	return nil
}

func GetFolderByID(id uint) (*models.Folder, error) {
	return repositories.GetFolderByID(id)
}

func DownloadFile(id uint) (*models.File, error) {
	file,err:=repositories.GetFileByID(id)
	if err!=nil{
		return nil,err
	}
	return file,nil


}

func CreateFolder(folder *models.Folder, user *models.User) error {
    err := GetUserById(user)
    if err != nil {
        return err
    }
    user.Folders = append(user.Folders, *folder)
    db := mysql.GetDB()
    err = db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Save(user).Error; err != nil {
            return err
        }
        if err := tx.Create(folder).Error; err != nil {
            return err
        }
        return nil
    })
    if err != nil {
        return err
    }
    return nil
}
