package services

import (
	"errors"
	"filemanger/internal/models"
	"filemanger/internal/repositories"
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
