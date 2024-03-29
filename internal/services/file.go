package services

import (
	"errors"
	"filemanger/internal/constants"
	"filemanger/internal/models"
	"filemanger/internal/repositories"
	"filemanger/internal/repositories/mysql"
	"strconv"

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
	file, err := repositories.GetFileByID(id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CreateFolder(folder *models.Folder, user *models.User) error {
	err := GetUserById(user, constants.Folder)
	if err != nil {
		return err
	}
	db := mysql.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(folder).Error; err != nil {
			return err
		}
		user.Folders = append(user.Folders, *folder)
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UploadFile(file *models.File, user *models.User) error {
	err := GetUserById(user, constants.File)
	if err != nil {
		return err
	}

	db := mysql.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(file).Error; err != nil {
			return err
		}
		user.Files = append(user.Files, *file)
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}


func DeleteFolder(folder *models.Folder) error {
	db := mysql.GetDB()

	db.Preload("Children").Preload("Files").First(folder)
	err := db.Transaction(func(tx *gorm.DB) error {

		for _,child:=range folder.Children{
			if err:=DeleteFolder(child);err!=nil{
				return err
			}
		}

		for _, file := range folder.Files{
			if err := tx.Delete(&file).Error; err != nil {
					return err
			}
		}

		if err := tx.Delete(folder).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func DeleteFile(file *models.File) error {
	db := mysql.GetDB()
	db.First(file)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(file).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func GetFolderByUserid(id uint) (*models.Folder, error) {
	var folder models.Folder
	db := mysql.GetDB()
	path := "/" + strconv.FormatUint(uint64(id), 10)
	if err := db.Preload("Files").Preload("Children").Where("path = ?", path).First(&folder).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func SearchContent(keyword string,user *models.User)([]models.File,[]models.Folder,error){
	// search file
	var files []models.File
	var folders []models.Folder
	db:=mysql.GetDB()
	if err:=db.Where("file_name LIKE ? AND user_id = ?",keyword+"%",user.ID).Find(&files).Error;err!=nil{
		return nil,nil,err
	}
	if err:=db.Where("name LIKE ? AND user_id = ?",keyword+"%",user.ID).Find(&folders).Error;err!=nil{
		return nil,nil,err
	}
	return files,folders,nil
}

func ShareFile(file *models.File) string{
	return "http://localhost:8080/api/v1/share/"+strconv.FormatUint(uint64(file.ID),10)
}
