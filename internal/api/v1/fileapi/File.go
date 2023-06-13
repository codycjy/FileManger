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
	path:=config.STORE_PATH+file.Path
	c.File(path)
}

type uploadFileRequest struct {
	File models.File `json:"file" binding:"required"`
	User models.User `json:"user" binding:"required"`
}

func UploadFile(c *gin.Context) {
	var reqFile uploadFileRequest
	file,err := c.FormFile("fileupload")
	userid,err:=strconv.ParseUint(c.PostForm("userid"), 10, 64)
	folderid,err:=strconv.ParseUint(c.PostForm("folderid"), 10, 64)
	path:=c.PostForm("path")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	reqFile.File.FileName = file.Filename
	reqFile.File.Size = file.Size
	reqFile.File.Path = path+"/"+file.Filename
	reqFile.File.UserID = uint(userid)
	reqFile.File.FolderID = uint(folderid)
	reqFile.User.ID = uint(userid)
	c.SaveUploadedFile(file,config.STORE_PATH+path+"/"+file.Filename)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	err=services.UploadFile(&reqFile.File,&reqFile.User)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,gin.H{
		"status":0,
	})

}
func UpdateFile(c *gin.Context) {

}

func DeleteFile(c *gin.Context) {
	id:=c.Param("id")
	num,err:=strconv.ParseUint(id, 10, 64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	var reqFile models.File
	reqFile.ID=uint(num)
	err=services.DeleteFile(&reqFile)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	folder,err:=services.GetFolderByID(reqFile.FolderID)
	c.JSON(200,gin.H{
		"status":0,
		"folder":folder,
	})

}
func DeleteFolder(c *gin.Context) {
	id:=c.Param("id")
	num,err:=strconv.ParseUint(id, 10, 64)
	var reqFolder models.Folder
	reqFolder.ID=uint(num)
	err=services.DeleteFolder(&reqFolder)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	if reqFolder.ParentID==nil{
		c.JSON(200,gin.H{
			"status":0,
		})
		return
	}
	folder,err:=services.GetFolderByID(*reqFolder.ParentID)

	c.JSON(200,gin.H{
		"status":0,
		"folder":folder,
	})


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

type searchRequest struct {
	Keyword string `json:"keyword" binding:"required"`
	User    models.User `json:"user" binding:"required"`
}
func SearchContent(c *gin.Context){
	var req searchRequest
	err:=c.ShouldBindJSON(&req)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	files,folders,err:=services.SearchContent(req.Keyword,&req.User)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,gin.H{
		"status":0,
		"files":files,
		"folders":folders,
	})

}

func ShareFile(c *gin.Context){
	var reqFile models.File
	c.ShouldBindJSON(&reqFile)
	res:=services.ShareFile(&reqFile)
	c.JSON(200,gin.H{
		"status":0,
		"url":res,
	})

}

func DownloadShareFile(c *gin.Context) {
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
