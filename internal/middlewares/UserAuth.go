package middlewares

import (
	model "filemanger/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// finish a jwt middleware

func getUserFromDB(userID int) *model.User {
	// TODO: Finish this function
	// NOTE: Use redis to cache user info
	return &model.User{}
}
func Authorization(allowedLevels ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 2: Extract authentication information (e.g., a JWT token or session ID)
		// For demonstration purposes, we will use a user ID from the header //TODO: Use jwt
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Step 3: Retrieve user information from the database or another storage
		user := getUserFromDB(1) // Replace 1 with the actual user ID
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Step 4: Check if the user level is allowed
		isAllowed := false
		for _, level := range allowedLevels {
			if user.Level == level {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			// Step 5: If the user level is not allowed, return an error response
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		// If everything is OK, set the user in the context and continue processing the request
		c.Set("user", user)
		c.Next()
	}
}

// NOTE: if public let pass 
// if private check if user own the file
