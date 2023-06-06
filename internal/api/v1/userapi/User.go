package userapi

import (
	"filemanger/internal/models"
	"filemanger/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Finish these function
func GetUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err})
		return
	}
	services.GetUserById(&user,0)
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": user})

}

func AddUser(c *gin.Context) {

}

type UpdateUserRequest struct {
	models.User
	OldPW string `json:"old_pw"`

}
// TEST: Make test later
func UpdateUser(c *gin.Context) {
	var  quser models.User
	var req UpdateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err})
		return
	}
	quser.ID = req.ID
	services.GetUserById(&quser,0)
	if req.OldPW!= quser.Password {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": "password incorrect"})
		return
	}else{
		err:=services.UpdateUser(&req.User)
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": 0, "data": req.User})

	}





}

func DeleteUser(c *gin.Context) {

}
