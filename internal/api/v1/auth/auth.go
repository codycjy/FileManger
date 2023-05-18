package auth

import (
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
	Token string `json:"token"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest // TODO: replace it with user model
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Validate the username and password
	// ...

	// Create a new token with the user ID as the claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": "123",
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
}

func LogoutHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
