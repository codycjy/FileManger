package userapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context){
	userID:=c.Param("id")
	c.JSON(http.StatusOK,gin.H {"status":"ok","userID":userID})


}
