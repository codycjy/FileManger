package api

import "github.com/gin-gonic/gin"

func Router(){
	r := gin.Default()	
	r.Run(":8080")
}
