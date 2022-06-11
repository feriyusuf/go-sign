package helpers

import (
	"github.com/feriyusuf/go-sign/app/models_mongo"
	"github.com/feriyusuf/go-sign/app/models_pg/commons"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TokenAuthentication(PGDB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("token")

		// There's no headers' token
		if headerToken == "" {
			c.JSON(401, gin.H{"message": "Token is required"})
			return
		}

		username, err := DecodeToken(headerToken)

		// Unrecognized token
		if err != nil {
			c.JSON(401, gin.H{"message": "Unknown token"})
			c.Abort()
			return
		}

		isSessionActive, _ := models_mongo.IsActiveSession(headerToken)

		if !isSessionActive {
			// Set all active session to false if any
			models_mongo.DestroySession(username)
			c.JSON(401, gin.H{"message": "Unknown Token"})
			c.Abort()
			return
		}

		// Check user to database
		user := commons.ComUser{}
		PGDB.Where("username = ?", username).First(&user)

		if user.ID == 0 {
			c.JSON(401, gin.H{"message": "User not found"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Set("user_id", user.ID)
	}
}
