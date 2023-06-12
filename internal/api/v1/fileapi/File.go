package fileapi

import (
	"filemanger/internal/config"
	"filemanger/internal/models"
	"filemanger/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DownloadFile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	file, err := services.DownloadFile(uint(id))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.File(file.Path)
}

type uploadFileRequest struct {
	File models.File `json:"file" binding:"required"`
	User models.User `json:"user" binding:"required"`
}

func UploadFile(c *gin.Context) {
	var reqFile uploadFileRequest
	file, err := c.FormFile("fileupload")

	userid, err := strconv.ParseUint(c.PostForm("userid"), 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	folderid, err := strconv.ParseUint(c.PostForm("folderid"), 10, 64)
	path := c.PostForm("path")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	reqFile.File.FileName = file.Filename
	reqFile.File.Size = file.Size
	reqFile.File.Path = path + "/" + file.Filename
	reqFile.File.UserID = uint(userid)
	reqFile.File.FolderID = uint(folderid)
	reqFile.User.ID = uint(userid)
	c.SaveUploadedFile(file, config.STORE_PATH+path+"/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.UploadFile(&reqFile.File, &reqFile.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
	})

}
func UpdateFile(c *gin.Context) {

}

func DeleteFile(c *gin.Context) {

}

type createFolderRequest struct {
	Folder models.Folder `json:"folder" binding:"required"`
	User   models.User   `json:"user" binding:"required"`
}

func CreateFolder(c *gin.Context) {
	var reqFolder createFolderRequest

	err := c.ShouldBindJSON(&reqFolder)
	parentId := *reqFolder.Folder.ParentID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.CreateFolder(&reqFolder.Folder, &reqFolder.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	currentFolder, err := services.GetFolderByID(parentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
		"folder": currentFolder,
	})

}

func GetFolderByID(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.ParseUint(id, 10, 64)
	reqFolder := models.Folder{
		Model: gorm.Model{
			ID: uint(num),
		},
	}
	folder, err := services.GetFolderByID(reqFolder.ID)
	// TODO: return a list of files and folders
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
		"folder": folder,
	})
}
