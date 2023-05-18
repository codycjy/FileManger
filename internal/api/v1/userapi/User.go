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
	services.GetUserById(&user)
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": user})

}

func AddUser(c *gin.Context) {

}

// NOTE: maybe useless do it later
func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
