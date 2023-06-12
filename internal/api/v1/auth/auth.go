package auth

import (
	"filemanger/internal/models"
	"filemanger/internal/repositories"
	"filemanger/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token  string        `json:"token"`
	User   models.User   `json:"user"`
	Folder models.Folder `json:"folder"`
}

func LoginHandler(c *gin.Context) {
	var req models.User
	var folder *models.Folder
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ...Validate
	err := repositories.LoginUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong username or password"})
		return
	}
	// Create a new token with the user ID as the claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": req.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	folder, err = services.GetFolderByUserid(req.ID)
	c.JSON(http.StatusOK, LoginResponse{Token: tokenString, User: req, Folder: *folder})
}

func LogoutHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
