package api

import (
	"filemanger/internal/api/v1/fileapi"
	"filemanger/internal/api/v1/userapi"

	"github.com/gin-gonic/gin"
)

func Router(){
    r := gin.Default()

    v1 := r.Group("/api/v1")
    {
        userGroup := v1.Group("/users")
        {
            userGroup.GET("/:id", userapi.GetUser)
            userGroup.POST("", userapi.AddUser)
            userGroup.PUT("/:id", userapi.UpdateUser)
            userGroup.DELETE("/:id", userapi.DeleteUser)
        }

        fileGroup := v1.Group("/files")
        {
            fileGroup.GET("/:id", fileapi.GetFile)
            fileGroup.POST("", fileapi.AddFile)
            fileGroup.PUT("/:id", fileapi.UpdateFile)
            fileGroup.DELETE("/:id", fileapi.DeleteFile)
        }
    }

    r.Run(":8080")
}
