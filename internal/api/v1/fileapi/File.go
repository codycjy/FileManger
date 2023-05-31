package fileapi

import (
	"filemanger/internal/models"
	"filemanger/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func UploadFile(c *gin.Context) {

}

func UpdateFile(c *gin.Context) {

}

func DeleteFile(c *gin.Context) {

}

func GetFolderByID(c *gin.Context) {
	var reqFolder models.Folder
	err:=c.ShouldBindJSON(&reqFolder)
	if err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	folder,err:=services.GetFolderByID(reqFolder.ID)
	// TODO: return a list of files and folders
	c.JSON(200, gin.H{
		"status": 0,
		"folder": folder,
	})
}




